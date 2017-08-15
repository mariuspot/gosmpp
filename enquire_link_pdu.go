package smpp

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
)

type EnquireLinkPDU struct {
	*Header
}

func (b EnquireLinkPDU) String() string {
	return fmt.Sprintf("{Header: %s}", b.Header)
}

func NewEnquireLinkPDU() (*EnquireLinkPDU, error) {
	return &EnquireLinkPDU{Header: NewHeader(CID_EnquireLink)}, nil
}

func NewEnquireLinkPDUFromBuffer(buf *bytes.Buffer) (*EnquireLinkPDU, error) {
	header, err := NewHeaderFromBuffer(buf)
	if err != nil {
		return nil, err
	}
	if header.GetCommandId() != CID_EnquireLink {
		return nil, errors.New("Not a EnquireLinkPDU")
	}
	return &EnquireLinkPDU{Header: header}, nil
}

func NewEnquireLinkPDUFromHeaderAndBuffer(header *Header, buf *bytes.Buffer) (*EnquireLinkPDU, error) {
	if header.GetCommandId() != CID_EnquireLink {
		return nil, errors.New("Not a EnquireLinkPDU")
	}
	return &EnquireLinkPDU{Header: header}, nil
}

func (b *EnquireLinkPDU) CalculateCommandLength() uint32 {
	return 16
}

func (b *EnquireLinkPDU) WriteToBuffer(buf *bufio.Writer) error {
	b.commandLength = b.CalculateCommandLength()

	if err := b.Header.WriteToBuffer(buf); err != nil {
		return err
	}
	return nil
}
