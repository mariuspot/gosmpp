package smpp

import (
	"bytes"
	"errors"
	"fmt"
)

type BindTransceiverPDU struct {
	*BindPDU
	Response *BindTransceiverResponsePDU
}

func NewBindTransceiverPDU(systemId string, password string, systemType string) (*BindTransceiverPDU, error) {
	pdu, err := NewBindPDU(CID_BindTransceiver, systemId, password, systemType)
	if err != nil {
		return nil, err
	}
	return &BindTransceiverPDU{BindPDU: pdu}, nil
}

func NewBindTransceiverPDUFromBuffer(buf *bytes.Buffer) (*BindTransceiverPDU, error) {
	pdu, err := NewBindPDUFromBuffer(buf)
	if err != nil {
		return nil, err
	}
	if pdu.GetCommandId() != CID_BindTransceiver {
		return nil, errors.New("Not a BindRecevierPDU")
	}
	return &BindTransceiverPDU{BindPDU: pdu}, nil
}

func NewBindTransceiverPDUFromHeaderAndBuffer(header *Header, buf *bytes.Buffer) (*BindTransceiverPDU, error) {
	pdu, err := NewBindPDUFromHeaderAndBuffer(header, buf)
	if err != nil {
		return nil, err
	}
	if pdu.GetCommandId() != CID_BindTransceiver {
		return nil, errors.New("Not a BindRecevierPDU")
	}
	return &BindTransceiverPDU{BindPDU: pdu}, nil
}

func (b *BindTransceiverPDU) SetResponse(response interface{}) error {
	fmt.Println("SetResponse")
	if pdu, ok := response.(*BindTransceiverResponsePDU); ok {
		b.Response = pdu
		return nil
	}
	return errors.New("Unable to set response")
}
