package connection

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
	// "encoding/binary"
	// "errors"
	// "fmt"
	"net"

	"github.com/mariuspot/smpp"
)

type SMPPConnection struct {
	connection      net.Conn
	reader          *bufio.Reader
	writer          *bufio.Writer
	sendm           sync.Mutex
	sequenceNumber  uint32
	closeChan       chan bool
	notify          bool
	notifyChan      chan bool
	errorChan       chan error
	pduCommandChan  chan interface{}
	pduResponseChan chan interface{}
	C               <-chan interface{}
	RC              <-chan interface{}
	E               <-chan error
	sendTimer       *time.Timer
	name            string
	sendCounter     *int32

	errorMutex sync.Mutex
	closeMutex sync.Mutex
}

type PDU interface {
	WriteToBuffer(buf *bufio.Writer) error
	GetSequenceNumber() uint32
	SetSequenceNumber(sequenceNumber uint32)
}

func NewSMPPConnection(connection net.Conn, name string, receivers int) *SMPPConnection {
	// connection.SetNoDelay(true)
	smppConnection := &SMPPConnection{
		connection:      connection,
		reader:          bufio.NewReader(connection),
		writer:          bufio.NewWriter(connection),
		sequenceNumber:  0,
		closeChan:       make(chan bool),
		notify:          false,
		notifyChan:      make(chan bool),
		errorChan:       make(chan error, 500),
		pduCommandChan:  make(chan interface{}, 1000),
		pduResponseChan: make(chan interface{}, 1000),
		sendTimer:       time.NewTimer(time.Second),
		name:            name,
		sendCounter:     new(int32),
	}
	smppConnection.C = smppConnection.pduCommandChan
	smppConnection.RC = smppConnection.pduResponseChan
	smppConnection.E = smppConnection.errorChan
	for i := 0; i < receivers; i++ {
		go smppConnection.receiver()
	}
	go smppConnection.buffer()

	return smppConnection
}

func (c *SMPPConnection) Notify() {
	c.notifyChan <- true
}

func (c *SMPPConnection) AutoNotify() {
	c.notify = true
	c.Notify()
}

func (c *SMPPConnection) buffer() {
	defer func() {
		if r := recover(); r != nil {
			err := r.(error)
			fmt.Println("buffer =====", err, "=====")
			time.Sleep(time.Second * 30)
			c.sendm.Unlock()
		}
	}()
	for {
		select {
		case <-c.sendTimer.C:
			if c.writer.Buffered() > 0 && atomic.LoadInt32(c.sendCounter) > 0 {
				c.sendm.Lock()
				c.connection.SetWriteDeadline(time.Now().Add(time.Second * 2))
				if err := c.writer.Flush(); err != nil {
					fmt.Println("Flush error")
					c.error(err)
				}
				atomic.StoreInt32(c.sendCounter, 0)
				c.sendm.Unlock()
			}
			c.sendTimer.Reset(time.Second)
		case <-c.closeChan:
			return
		}
	}
}

func (c *SMPPConnection) Send(pdu smpp.PDU, flush bool) (seqNo uint32, err error) {
	defer func() {
		if r := recover(); r != nil {
			err := r.(error)
			fmt.Println("send =====", err, "=====")
			time.Sleep(time.Second * 30)
			c.sendm.Unlock()
		}
	}()

	c.sendm.Lock()
	select {
	case <-c.closeChan:
		c.sendm.Unlock()
		return 0, ConnectionClosedError
	default:
	}

	c.sequenceNumber++
	if pdu.GetSequenceNumber() == 0 {
		pdu.SetSequenceNumber(c.sequenceNumber)
		seqNo = c.sequenceNumber
	}

	c.connection.SetWriteDeadline(time.Now().Add(time.Second * 2))

	if err = pdu.WriteToBuffer(c.writer); err != nil {
		fmt.Println("Send error")
		c.error(err)
		c.sendm.Unlock()
		return
	}
	atomic.AddInt32(c.sendCounter, 1)
	c.sendTimer.Reset(time.Millisecond * 50)
	if flush || atomic.LoadInt32(c.sendCounter) > 30 { //|| c.sequenceNumber%100 == 0 {
		atomic.StoreInt32(c.sendCounter, 0)
		if err = c.writer.Flush(); err != nil {
			fmt.Println("Send error")
			c.error(err)
		}
	}
	c.sendm.Unlock()
	return
}

func (c *SMPPConnection) Close() {
	c.closeMutex.Lock()
	defer c.closeMutex.Unlock()
	select {
	case <-c.closeChan:
	default:
		close(c.closeChan)
		c.connection.Close()
	}
}

func (c *SMPPConnection) error(err error) {
	c.errorMutex.Lock()
	defer c.errorMutex.Unlock()

	select {
	case <-c.closeChan:
	default:
		c.Close()
		c.errorChan <- err
	}
}

func (c *SMPPConnection) receiver() {
	var (
		byteBuffer []byte = make([]byte, 4096)
		cnt        uint32
		buffer     *bytes.Buffer = &bytes.Buffer{}
	)

	for {
		select {
		case <-c.closeChan:
			return
		case <-c.notifyChan:
		}
		cnt = 0
		for cnt < 4 {
			c.connection.SetReadDeadline(time.Now().Add(time.Second))
			if n, err := c.reader.Read(byteBuffer[cnt:4]); err != nil && n <= 0 {
				if e, ok := err.(*net.OpError); !ok || !e.Timeout() {
					c.error(err)
					return
				}
			} else {
				cnt += uint32(n)
			}
		}
		//
		//

		size := uint32(binary.BigEndian.Uint32(byteBuffer[0:4]))
		if size > 4096 {
			c.error(LargePDUError)
			return
		}

		for cnt < size {
			c.connection.SetReadDeadline(time.Now().Add(time.Second))
			if n, err := c.reader.Read(byteBuffer[cnt:size]); err != nil && n <= 0 {
				if e, ok := err.(*net.OpError); !ok || !e.Timeout() {
					c.error(err)
					return
				}
			} else {
				cnt += uint32(n)
			}
		}
		if c.notify {
			c.Notify()
		}
		buffer.Write(byteBuffer[0:size])

		header, err := smpp.NewHeaderFromBuffer(buffer)
		var pdu interface{}
		if err == nil {
			switch header.GetCommandId() {
			//CID_BindReceiver
			case smpp.CID_BindReceiver:
				pdu, err = smpp.NewBindReceiverPDUFromHeaderAndBuffer(header, buffer)
			//CID_BindTransmitter
			case smpp.CID_BindTransmitter:
				pdu, err = smpp.NewBindTransmitterPDUFromHeaderAndBuffer(header, buffer)
			//CID_QuerySm
			//CID_SubmitSm
			case smpp.CID_SubmitSm:
				pdu, err = smpp.NewSubmitSmPDUFromHeaderAndBuffer(header, buffer)
			//CID_DeliverSm
			case smpp.CID_DeliverSm:
				pdu, err = smpp.NewDeliverSmPDUFromHeaderAndBuffer(header, buffer)
			//CID_Unbind
			case smpp.CID_Unbind:
				pdu, err = smpp.NewUnbindPDUFromHeaderAndBuffer(header, buffer)
			//CID_ReplaceSm
			//CID_CancelSm
			//CID_BindTransceiver
			case smpp.CID_BindTransceiver:
				pdu, err = smpp.NewBindTransceiverPDUFromHeaderAndBuffer(header, buffer)
			//CID_Outbind
			//CID_EnquireLink
			case smpp.CID_EnquireLink:
				pdu, err = smpp.NewEnquireLinkPDUFromHeaderAndBuffer(header, buffer)
			//CID_SubmitMulti
			//CID_AlertNotification
			//CID_DataSm
			//CID_BroadcastSm
			//CID_QueryBroadcastSm
			//CID_CancelBroadcastSm
			//CID_GenericNack
			case smpp.CID_GenericNack:
				pdu, err = smpp.NewGenericNackPDUFromHeaderAndBuffer(header, buffer)
			//CID_BindReceiverResponse
			case smpp.CID_BindReceiverResponse:
				pdu, err = smpp.NewBindReceiverResponsePDUFromHeaderAndBuffer(header, buffer)
			//CID_BindTransmitterResponse
			case smpp.CID_BindTransmitterResponse:
				pdu, err = smpp.NewBindTransmitterResponsePDUFromHeaderAndBuffer(header, buffer)
			//CID_QuerySmResponse
			//CID_SubmitSmResponse
			case smpp.CID_SubmitSmResponse:
				pdu, err = smpp.NewSubmitSmResponsePDUFromHeaderAndBuffer(header, buffer)
			//CID_DeliverSmResponse
			case smpp.CID_DeliverSmResponse:
				pdu, err = smpp.NewDeliverSmResponsePDUFromHeaderAndBuffer(header, buffer)
			//CID_UnbindResponse
			case smpp.CID_UnbindResponse:
				pdu, err = smpp.NewUnbindResponsePDUFromHeaderAndBuffer(header, buffer)
			//CID_ReplaceSmResponse
			//CID_CancelSmResponse
			//CID_BindTransceiverResponse
			case smpp.CID_BindTransceiverResponse:
				pdu, err = smpp.NewBindTransceiverResponsePDUFromHeaderAndBuffer(header, buffer)
			//CID_EnquireLinkResponse
			case smpp.CID_EnquireLinkResponse:
				pdu, err = smpp.NewEnquireLinkResponsePDUFromHeaderAndBuffer(header, buffer)
			//CID_SubmitMultiResponse
			//CID_DataSmResponse
			//CID_BroadcastSmResponse
			//CID_QueryBroadcastSmResponse
			//CID_CancelBroadcastSmResponse
			default:
				err = smpp.InvalidCommandIDError
			}
		}
		if err != nil {
			c.error(err)
			return
		} else {
			if header.GetCommandId() > smpp.CID_GenericNack {
				select {
				case <-c.closeChan:
				default:
					c.pduResponseChan <- pdu
				}
			} else {
				select {
				case <-c.closeChan:
				default:
					c.pduCommandChan <- pdu
				}
			}
		}
		buffer.Reset()
	}
}

func (s *SMPPConnection) RemoteAddr() net.Addr {
	return s.connection.RemoteAddr()
}
