package smpp

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
)

type UnbindPDU struct {
	*Header
}

func (b UnbindPDU) String() string {
	return fmt.Sprintf("{Header: %s}", b.Header)
}

func NewUnbindPDU() (*UnbindPDU, error) {
	return &UnbindPDU{Header: NewHeader(CID_Unbind)}, nil
}

func NewUnbindPDUFromBuffer(buf *bytes.Buffer) (*UnbindPDU, error) {
	header, err := NewHeaderFromBuffer(buf)
	if err != nil {
		return nil, err
	}
	return NewUnbindPDUFromHeaderAndBuffer(header, buf)
}

func NewUnbindPDUFromHeaderAndBuffer(header *Header, buf *bytes.Buffer) (*UnbindPDU, error) {
	if header.GetCommandId() != CID_Unbind {
		return nil, errors.New("Not a UnbindPDU")
	}
	return &UnbindPDU{Header: header}, nil
}

func (b *UnbindPDU) CalculateCommandLength() uint32 {
	return 16
}

func (b *UnbindPDU) WriteToBuffer(buf *bufio.Writer) error {
	b.commandLength = b.CalculateCommandLength()

	if err := b.Header.WriteToBuffer(buf); err != nil {
		return err
	}
	return nil
}
