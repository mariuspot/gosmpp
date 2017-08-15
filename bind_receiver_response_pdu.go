package smpp

import (
	"bytes"
	"errors"
)

type BindReceiverResponsePDU struct {
	*BindResponsePDU
}

func NewBindReceiverResponsePDU(systemId string) (*BindReceiverResponsePDU, error) {
	pdu, err := NewBindResponsePDU(CID_BindReceiverResponse, systemId)
	if err != nil {
		return nil, err
	}
	return &BindReceiverResponsePDU{BindResponsePDU: pdu}, nil
}

func NewBindReceiverResponsePDUFromBuffer(buf *bytes.Buffer) (*BindReceiverResponsePDU, error) {
	pdu, err := NewBindResponsePDUFromBuffer(buf)
	if err != nil {
		return nil, err
	}
	if pdu.GetCommandId() != CID_BindReceiverResponse {
		return nil, errors.New("Not a BindReceiverResponsePDU")
	}
	return &BindReceiverResponsePDU{BindResponsePDU: pdu}, nil
}

func NewBindReceiverResponsePDUFromHeaderAndBuffer(header *Header, buf *bytes.Buffer) (*BindReceiverResponsePDU, error) {
	pdu, err := NewBindResponsePDUFromHeaderAndBuffer(header, buf)
	if err != nil {
		return nil, err
	}
	if pdu.GetCommandId() != CID_BindReceiverResponse {
		return nil, errors.New("Not a BindReceiverResponsePDU")
	}
	return &BindReceiverResponsePDU{BindResponsePDU: pdu}, nil
}
