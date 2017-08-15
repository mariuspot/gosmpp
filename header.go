package smpp

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
)

type Header struct {
	commandLength  uint32
	commandId      CommandId
	commandStatus  CommandStatus
	sequenceNumber uint32
}

func (h Header) String() string {

	return fmt.Sprintf("{commandLength: %d, commandId: %s, commandStatus: %s, sequenceNumber: %d}", h.commandLength, h.commandId, h.commandStatus, h.sequenceNumber)
}

func NewHeader(commandId CommandId) *Header {
	return &Header{0, commandId, CS_ESME_ROK, 0}
}

func NewHeaderFromBuffer(buf *bytes.Buffer) (*Header, error) {
	header := &Header{}
	if err := binary.Read(buf, binary.BigEndian, &header.commandLength); err != nil {
		return nil, err
	}
	if err := binary.Read(buf, binary.BigEndian, &header.commandId); err != nil {
		return nil, err
	}
	if err := binary.Read(buf, binary.BigEndian, &header.commandStatus); err != nil {
		return nil, err
	}
	if err := binary.Read(buf, binary.BigEndian, &header.sequenceNumber); err != nil {
		return nil, err
	}
	return header, nil
}

func (h *Header) GetCommandLength() uint32 {
	return h.commandLength
}

func (h *Header) GetCommandId() CommandId {
	return h.commandId
}

func (h *Header) GetCommandStatus() CommandStatus {
	return h.commandStatus
}

func (h *Header) SetCommandStatus(commandStatus CommandStatus) {
	h.commandStatus = commandStatus
}

func (h *Header) GetSequenceNumber() uint32 {
	return h.sequenceNumber
}

func (h *Header) SetSequenceNumber(sequenceNumber uint32) {
	h.sequenceNumber = sequenceNumber
}

func (h *Header) WriteToBuffer(buf *bufio.Writer) error {
	if err := binary.Write(buf, binary.BigEndian, h.commandLength); err != nil {
		return err
	}
	if err := binary.Write(buf, binary.BigEndian, h.commandId); err != nil {
		return err
	}
	if err := binary.Write(buf, binary.BigEndian, h.commandStatus); err != nil {
		return err
	}
	if err := binary.Write(buf, binary.BigEndian, h.sequenceNumber); err != nil {
		return err
	}
	return nil
}
