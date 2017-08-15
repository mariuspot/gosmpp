package smpp

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
)

type SmResponsePDU struct {
	*Header
	messageId string
}

func (s SmResponsePDU) String() string {
	return fmt.Sprintf("{Header: %s, messageId: %s}", s.Header, s.messageId)
}

func NewSmResponsePDU(commandId CommandId, messageId string) (*SmResponsePDU, error) {
	if len(messageId) > 65 {
		return nil, errors.New("messageId max length 65")
	}
	return &SmResponsePDU{Header: NewHeader(commandId), messageId: messageId}, nil
}

func NewSmResponsePDUFromBuffer(buf *bytes.Buffer) (*SmResponsePDU, error) {
	header, err := NewHeaderFromBuffer(buf)
	if err != nil {
		return nil, err
	}
	return NewSmResponsePDUFromHeaderAndBuffer(header, buf)
}

func NewSmResponsePDUFromHeaderAndBuffer(header *Header, buf *bytes.Buffer) (*SmResponsePDU, error) {
	pdu := &SmResponsePDU{Header: header}

	if header.GetCommandLength() > 16 {
		var err error
		if pdu.messageId, err = buf.ReadString(0); err != nil {
			return nil, err
		}
		pdu.messageId = pdu.messageId[:len(pdu.messageId)-1]
	}

	return pdu, nil
}

func (s *SmResponsePDU) CalculateCommandLength() uint32 {
	return 16 + uint32(len(s.messageId)+1)
}

func (s *SmResponsePDU) GetMessageId() string {
	return s.messageId
}

func (s *SmResponsePDU) WriteToBuffer(buf *bufio.Writer) error {
	s.commandLength = s.CalculateCommandLength()

	if err := s.Header.WriteToBuffer(buf); err != nil {
		return err
	}
	if _, err := buf.WriteString(s.messageId); err != nil {
		return err
	}
	if err := buf.WriteByte(0); err != nil {
		return err
	}
	return nil
}
