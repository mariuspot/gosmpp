package connection

import (
	"errors"
)

var (
	ConnectionClosedError = errors.New("Connection Closed")
	LargePDUError         = errors.New("Large PDU, Size > 4096")
)
