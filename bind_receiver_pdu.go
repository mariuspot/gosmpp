package smpp

import (
	"bytes"
	"errors"
)

type BindReceiverPDU struct {
	*BindPDU
}

func NewBindReceiverPDU(systemId string, password string, systemType string) (*BindReceiverPDU, error) {
	pdu, err := NewBindPDU(CID_BindReceiver, systemId, password, systemType)
	if err != nil {
		return nil, err
	}
	return &BindReceiverPDU{BindPDU: pdu}, nil
}

func NewBindReceiverPDUFromBuffer(buf *bytes.Buffer) (*BindReceiverPDU, error) {
	pdu, err := NewBindPDUFromBuffer(buf)
	if err != nil {
		return nil, err
	}
	if pdu.GetCommandId() != CID_BindReceiver {
		return nil, errors.New("Not a BindRecevierPDU")
	}
	return &BindReceiverPDU{BindPDU: pdu}, nil
}

func NewBindReceiverPDUFromHeaderAndBuffer(header *Header, buf *bytes.Buffer) (*BindReceiverPDU, error) {
	pdu, err := NewBindPDUFromHeaderAndBuffer(header, buf)
	if err != nil {
		return nil, err
	}
	if pdu.GetCommandId() != CID_BindReceiver {
		return nil, errors.New("Not a BindRecevierPDU")
	}
	return &BindReceiverPDU{BindPDU: pdu}, nil
}
