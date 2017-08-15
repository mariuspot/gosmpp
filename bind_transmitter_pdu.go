package smpp

import (
	"bytes"
	"errors"
	"fmt"
)

type BindTransmitterPDU struct {
	*BindPDU
	Response *BindTransmitterResponsePDU
}

func NewBindTransmitterPDU(systemId string, password string, systemType string) (*BindTransmitterPDU, error) {
	pdu, err := NewBindPDU(CID_BindTransmitter, systemId, password, systemType)
	if err != nil {
		return nil, err
	}
	return &BindTransmitterPDU{BindPDU: pdu}, nil
}

func NewBindTransmitterPDUFromBuffer(buf *bytes.Buffer) (*BindTransmitterPDU, error) {
	pdu, err := NewBindPDUFromBuffer(buf)
	if err != nil {
		return nil, err
	}
	if pdu.GetCommandId() != CID_BindTransmitter {
		return nil, errors.New("Not a BindRecevierPDU")
	}
	return &BindTransmitterPDU{BindPDU: pdu}, nil
}

func NewBindTransmitterPDUFromHeaderAndBuffer(header *Header, buf *bytes.Buffer) (*BindTransmitterPDU, error) {
	pdu, err := NewBindPDUFromHeaderAndBuffer(header, buf)
	if err != nil {
		return nil, err
	}
	if pdu.GetCommandId() != CID_BindTransmitter {
		return nil, errors.New("Not a BindRecevierPDU")
	}
	return &BindTransmitterPDU{BindPDU: pdu}, nil
}

func (b *BindTransmitterPDU) SetResponse(response interface{}) error {
	fmt.Println("SetResponse")
	if pdu, ok := response.(*BindTransmitterResponsePDU); ok {
		b.Response = pdu
		return nil
	}
	return errors.New("Unable to set response")
}
