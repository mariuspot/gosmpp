package smpp

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
)

type BindResponsePDU struct {
	*Header
	systemId string
	*TLVList
}

func (b BindResponsePDU) String() string {
	return fmt.Sprintf("{Header: %s, systemId: %s, TLVList: %s}", b.Header, b.systemId, b.TLVList)
}

func NewBindResponsePDU(commandId CommandId, systemId string) (*BindResponsePDU, error) {
	if len(systemId) > 16 {
		return nil, errors.New("systemId max length 16")
	}
	return &BindResponsePDU{Header: NewHeader(commandId), systemId: systemId, TLVList: NewTLVList()}, nil
}

func NewBindResponsePDUFromBuffer(buf *bytes.Buffer) (*BindResponsePDU, error) {
	header, err := NewHeaderFromBuffer(buf)
	if err != nil {
		return nil, err
	}
	return NewBindResponsePDUFromHeaderAndBuffer(header, buf)
}

func NewBindResponsePDUFromHeaderAndBuffer(header *Header, buf *bytes.Buffer) (*BindResponsePDU, error) {
	pdu := &BindResponsePDU{Header: header}
	var err error
	if pdu.systemId, err = buf.ReadString(0); err != nil {
		return nil, err
	}
	pdu.systemId = pdu.systemId[:len(pdu.systemId)-1]
	tlvList, err := NewTLVListFromBuffer(buf)
	if err != nil {
		return nil, err
	}
	pdu.TLVList = tlvList
	return pdu, nil
}

func (b *BindResponsePDU) CalculateCommandLength() uint32 {
	return 16 + uint32(len(b.systemId)+1) + b.TLVList.GetLength()
}

func (b *BindResponsePDU) GetSystemId() string {
	return b.systemId
}

func (b *BindResponsePDU) WriteToBuffer(buf *bufio.Writer) error {
	b.commandLength = b.CalculateCommandLength()

	if err := b.Header.WriteToBuffer(buf); err != nil {
		return err
	}
	n, _ := buf.WriteString(b.systemId)
	if n != len(b.systemId) {
		return errors.New("Error writing systemId to buffer")
	}
	if err := buf.WriteByte(0); err != nil {
		return err
	}
	if err := b.TLVList.WriteToBuffer(buf); err != nil {
		return err
	}

	return nil
}
