package smpp

type TON uint8

const (
	TON_Unknown          TON = iota // 0b00000000
	TON_International               // 0b00000001
	TON_National                    // 0b00000010
	TON_NetworkSpecific             // 0b00000011
	TON_SubscriberNumber            // 0b00000100
	TON_Alphanumeric                // 0b00000101
	TON_Abbreviated                 // 0b00000110
)

func (t TON) String() string {
	switch t {
	case TON_International:
		return "International"
	case TON_National:
		return "National"
	case TON_NetworkSpecific:
		return "Network Specific"
	case TON_SubscriberNumber:
		return "Subscriber Number"
	case TON_Alphanumeric:
		return "Alphanumeric"
	case TON_Abbreviated:
		return "Abbreviated"
	}
	return "Unknown"
}
