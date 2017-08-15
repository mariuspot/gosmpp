package smpp

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
)

type EnquireLinkResponsePDU struct {
	*Header
}

func (b EnquireLinkResponsePDU) String() string {
	return fmt.Sprintf("{Header: %s}", b.Header)
}

func NewEnquireLinkResponsePDU() (*EnquireLinkResponsePDU, error) {
	return &EnquireLinkResponsePDU{Header: NewHeader(CID_EnquireLinkResponse)}, nil
}

func NewEnquireLinkResponsePDUFromBuffer(buf *bytes.Buffer) (*EnquireLinkResponsePDU, error) {
	header, err := NewHeaderFromBuffer(buf)
	if err != nil {
		return nil, err
	}
	if header.GetCommandId() != CID_EnquireLinkResponse {
		return nil, errors.New("Not a EnquireLinkResponse PDU")
	}
	return &EnquireLinkResponsePDU{Header: header}, nil
}

func NewEnquireLinkResponsePDUFromHeaderAndBuffer(header *Header, buf *bytes.Buffer) (*EnquireLinkResponsePDU, error) {
	if header.GetCommandId() != CID_EnquireLinkResponse {
		return nil, errors.New("Not a EnquireLinkResponse PDU")
	}
	return &EnquireLinkResponsePDU{Header: header}, nil
}

func (b *EnquireLinkResponsePDU) CalculateCommandLength() uint32 {
	return 16
}

func (b *EnquireLinkResponsePDU) WriteToBuffer(buf *bufio.Writer) error {
	b.commandLength = b.CalculateCommandLength()

	if err := b.Header.WriteToBuffer(buf); err != nil {
		return err
	}
	return nil
}
