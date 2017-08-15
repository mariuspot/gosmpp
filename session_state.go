package smpp

type SessionState byte

const (
	SS_OPEN SessionState = iota
	SS_BOUND_TX
	SS_BOUND_RX
	SS_BOUND_TRX
	SS_CLOSED
)
