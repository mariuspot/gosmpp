package smpp

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
)

type UnbindResponsePDU struct {
	*Header
}

func (b UnbindResponsePDU) String() string {
	return fmt.Sprintf("{Header: %s}", b.Header)
}

func NewUnbindResponsePDU() (*UnbindResponsePDU, error) {
	return &UnbindResponsePDU{Header: NewHeader(CID_UnbindResponse)}, nil
}

func NewUnbindResponsePDUFromBuffer(buf *bytes.Buffer) (*UnbindResponsePDU, error) {
	header, err := NewHeaderFromBuffer(buf)
	if err != nil {
		return nil, err
	}
	return NewUnbindResponsePDUFromHeaderAndBuffer(header, buf)
}

func NewUnbindResponsePDUFromHeaderAndBuffer(header *Header, buf *bytes.Buffer) (*UnbindResponsePDU, error) {
	if header.GetCommandId() != CID_UnbindResponse {
		return nil, errors.New("Not a UnbindResponse PDU")
	}
	return &UnbindResponsePDU{Header: header}, nil
}

func (b *UnbindResponsePDU) CalculateCommandLength() uint32 {
	return 16
}

func (b *UnbindResponsePDU) WriteToBuffer(buf *bufio.Writer) error {
	b.commandLength = b.CalculateCommandLength()

	if err := b.Header.WriteToBuffer(buf); err != nil {
		return err
	}
	return nil
}
