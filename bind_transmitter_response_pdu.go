package smpp

import (
	"bytes"
	"errors"
)

type BindTransmitterResponsePDU struct {
	*BindResponsePDU
}

func NewBindTransmitterResponsePDU(systemId string) (*BindTransmitterResponsePDU, error) {
	pdu, err := NewBindResponsePDU(CID_BindTransmitterResponse, systemId)
	if err != nil {
		return nil, err
	}
	return &BindTransmitterResponsePDU{BindResponsePDU: pdu}, nil
}

func NewBindTransmitterResponsePDUFromBuffer(buf *bytes.Buffer) (*BindTransmitterResponsePDU, error) {
	pdu, err := NewBindResponsePDUFromBuffer(buf)
	if err != nil {
		return nil, err
	}
	if pdu.GetCommandId() != CID_BindTransmitterResponse {
		return nil, errors.New("Not a BindTransmitterResponsePDU")
	}
	return &BindTransmitterResponsePDU{BindResponsePDU: pdu}, nil
}

func NewBindTransmitterResponsePDUFromHeaderAndBuffer(header *Header, buf *bytes.Buffer) (*BindTransmitterResponsePDU, error) {
	pdu, err := NewBindResponsePDUFromHeaderAndBuffer(header, buf)
	if err != nil {
		return nil, err
	}
	if pdu.GetCommandId() != CID_BindTransmitterResponse {
		return nil, errors.New("Not a BindTransmitterResponsePDU")
	}
	return &BindTransmitterResponsePDU{BindResponsePDU: pdu}, nil
}
