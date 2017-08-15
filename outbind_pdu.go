package smpp

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
)

type OutbindPDU struct {
	*Header
	systemId string
	password string
}

func (b OutbindPDU) String() string {
	return fmt.Sprintf("{Header: %s, systemId: %s, password: %s}", b.Header, b.systemId, b.password)
}

func NewOutbindPDU(systemId string, password string, systemType string) (*OutbindPDU, error) {
	if len(systemId) > 16 {
		return nil, errors.New("systemId max length 16")
	}
	if len(password) > 9 {
		return nil, errors.New("password max length 9")
	}
	return &OutbindPDU{Header: NewHeader(CID_Outbind), systemId: systemId, password: password}, nil
}

func NewOutbindPDUFromBuffer(buf *bytes.Buffer) (*OutbindPDU, error) {
	header, err := NewHeaderFromBuffer(buf)
	if err != nil {
		return nil, err
	}
	return NewOutbindPDUFromHeaderAndBuffer(header, buf)
}

func NewOutbindPDUFromHeaderAndBuffer(header *Header, buf *bytes.Buffer) (*OutbindPDU, error) {
	if header.GetCommandId() != CID_Outbind {
		return nil, errors.New("Not a OutbindPDU")
	}
	pdu := &OutbindPDU{Header: header}
	var err error
	if pdu.systemId, err = buf.ReadString(0); err != nil {
		return nil, err
	}
	pdu.systemId = pdu.systemId[:len(pdu.systemId)-1]
	if pdu.password, err = buf.ReadString(0); err != nil {
		return nil, err
	}
	pdu.password = pdu.password[:len(pdu.password)-1]
	return pdu, nil
}

func (b *OutbindPDU) CalculateCommandLength() uint32 {
	return 16 + uint32(len(b.systemId)+1+len(b.password)+1)
}

func (b *OutbindPDU) GetSystemId() string {
	return b.systemId
}

func (b *OutbindPDU) WriteToBuffer(buf *bufio.Writer) error {
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
	n, _ = buf.WriteString(b.password)
	if n != len(b.password) {
		return errors.New("Error writing password to buffer")
	}
	if err := buf.WriteByte(0); err != nil {
		return err
	}
	return nil
}
