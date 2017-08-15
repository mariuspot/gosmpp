package smpp

import (
	"bytes"
	"errors"
)

type BindTransceiverResponsePDU struct {
	*BindResponsePDU
}

func NewBindTransceiverResponsePDU(systemId string) (*BindTransceiverResponsePDU, error) {
	pdu, err := NewBindResponsePDU(CID_BindTransceiverResponse, systemId)
	if err != nil {
		return nil, err
	}
	return &BindTransceiverResponsePDU{BindResponsePDU: pdu}, nil
}

func NewBindTransceiverResponsePDUFromBuffer(buf *bytes.Buffer) (*BindTransceiverResponsePDU, error) {
	pdu, err := NewBindResponsePDUFromBuffer(buf)
	if err != nil {
		return nil, err
	}
	if pdu.GetCommandId() != CID_BindTransceiverResponse {
		return nil, errors.New("Not a BindTransceiverResponsePDU")
	}
	return &BindTransceiverResponsePDU{BindResponsePDU: pdu}, nil
}

func NewBindTransceiverResponsePDUFromHeaderAndBuffer(header *Header, buf *bytes.Buffer) (*BindTransceiverResponsePDU, error) {
	pdu, err := NewBindResponsePDUFromHeaderAndBuffer(header, buf)
	if err != nil {
		return nil, err
	}
	if pdu.GetCommandId() != CID_BindTransceiverResponse {
		return nil, errors.New("Not a BindTransceiverResponsePDU")
	}
	return &BindTransceiverResponsePDU{BindResponsePDU: pdu}, nil
}
