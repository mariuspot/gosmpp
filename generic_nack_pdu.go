package smpp

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
)

type GenericNackPDU struct {
	*Header
}

func (b GenericNackPDU) String() string {
	return fmt.Sprintf("{Header: %s}", b.Header)
}

func NewGenericNackPDU() (*GenericNackPDU, error) {
	return &GenericNackPDU{Header: NewHeader(CID_GenericNack)}, nil
}

func NewGenericNackPDUFromBuffer(buf *bytes.Buffer) (*GenericNackPDU, error) {
	header, err := NewHeaderFromBuffer(buf)
	if err != nil {
		return nil, err
	}
	if header.GetCommandId() != CID_GenericNack {
		return nil, errors.New("Not a GenericNackPDU")
	}
	return &GenericNackPDU{Header: header}, nil
}

func NewGenericNackPDUFromHeaderAndBuffer(header *Header, buf *bytes.Buffer) (*GenericNackPDU, error) {
	if header.GetCommandId() != CID_GenericNack {
		return nil, errors.New("Not a GenericNackPDU")
	}
	return &GenericNackPDU{Header: header}, nil
}

func (b *GenericNackPDU) CalculateCommandLength() uint32 {
	return 16
}

func (b *GenericNackPDU) WriteToBuffer(buf *bufio.Writer) error {
	b.commandLength = b.CalculateCommandLength()

	if err := b.Header.WriteToBuffer(buf); err != nil {
		return err
	}
	return nil
}
