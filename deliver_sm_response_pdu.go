package smpp

import (
	"bytes"
	"errors"
)

type DeliverSmResponsePDU struct {
	*SmResponsePDU
}

func NewDeliverSmResponsePDU(messageId string) (*DeliverSmResponsePDU, error) {
	pdu, err := NewSmResponsePDU(CID_DeliverSmResponse, messageId)
	if err != nil {
		return nil, err
	}
	return &DeliverSmResponsePDU{SmResponsePDU: pdu}, nil
}

func NewDeliverSmResponsePDUFromBuffer(buf *bytes.Buffer) (*DeliverSmResponsePDU, error) {
	pdu, err := NewSmResponsePDUFromBuffer(buf)
	if err != nil {
		return nil, err
	}
	if pdu.GetCommandId() != CID_DeliverSmResponse {
		return nil, errors.New("Not a DeliverSmResponsePDU")
	}
	return &DeliverSmResponsePDU{SmResponsePDU: pdu}, nil
}

func NewDeliverSmResponsePDUFromHeaderAndBuffer(header *Header, buf *bytes.Buffer) (*DeliverSmResponsePDU, error) {
	pdu, err := NewSmResponsePDUFromHeaderAndBuffer(header, buf)
	if err != nil {
		return nil, err
	}
	if pdu.GetCommandId() != CID_DeliverSmResponse {
		return nil, errors.New("Not a DeliverSmResponsePDU")
	}
	return &DeliverSmResponsePDU{SmResponsePDU: pdu}, nil
}
