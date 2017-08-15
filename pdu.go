package smpp

import (
	"bufio"
)

type PDU interface {
	GetCommandLength() uint32
	GetCommandId() CommandId
	GetCommandStatus() CommandStatus
	SetCommandStatus(commandStatus CommandStatus)
	GetSequenceNumber() uint32
	SetSequenceNumber(sequenceNumber uint32)
	WriteToBuffer(buf *bufio.Writer) error
}
