package smpp

import (
	"bytes"
	"errors"
)

type SubmitSmResponsePDU struct {
	*SmResponsePDU
}

func NewSubmitSmResponsePDU(messageId string) (*SubmitSmResponsePDU, error) {
	pdu, err := NewSmResponsePDU(CID_SubmitSmResponse, messageId)
	if err != nil {
		return nil, err
	}
	return &SubmitSmResponsePDU{SmResponsePDU: pdu}, nil
}

func NewSubmitSmResponsePDUFromBuffer(buf *bytes.Buffer) (*SubmitSmResponsePDU, error) {
	pdu, err := NewSmResponsePDUFromBuffer(buf)
	if err != nil {
		return nil, err
	}
	if pdu.GetCommandId() != CID_SubmitSmResponse {
		return nil, errors.New("Not a SubmitSmResponsePDU")
	}
	return &SubmitSmResponsePDU{SmResponsePDU: pdu}, nil
}

func NewSubmitSmResponsePDUFromHeaderAndBuffer(header *Header, buf *bytes.Buffer) (*SubmitSmResponsePDU, error) {
	pdu, err := NewSmResponsePDUFromHeaderAndBuffer(header, buf)
	if err != nil {
		return nil, err
	}
	if pdu.GetCommandId() != CID_SubmitSmResponse {
		return nil, errors.New("Not a SubmitSmResponsePDU")
	}
	return &SubmitSmResponsePDU{SmResponsePDU: pdu}, nil
}
