package smpp

type NPI uint8

const (
	NPI_Unknown          NPI = iota // 0b00000000
	NPI_ISDN_E163_E164              // 0b00000001
	_                               // 0b00000010
	NPI_Data_X_121                  // 0b00000011
	NPI_Telex_F_69                  // 0b00000100
	_                               // 0b00000101
	NPI_LandMobile_E_212            // 0b00000110
	_                               // 0b00000111
	NPI_National                    // 0b00001000
	NPI_Private                     // 0b00001001
	NPI_ERMES                       // 0b00001010
	_                               // 0b00001011
	_                               // 0b00001100
	_                               // 0b00001101
	NPI_Internet_IP                 // 0b00001110
	_                               // 0b00001111
	_                               // 0b00010000
	_                               // 0b00010001
	NPI_WAPClientId                 // 0b00010010
)

func (n NPI) String() string {
	switch n {
	case NPI_ISDN_E163_E164:
		return "ISDN (E163/E164)"
	case NPI_Data_X_121:
		return "Data (X.121)"
	case NPI_Telex_F_69:
		return "Telex (F.69)"
	case NPI_LandMobile_E_212:
		return "Land Mobile (E212)"
	case NPI_National:
		return "National"
	case NPI_Private:
		return "Private"
	case NPI_ERMES:
		return "ERMES"
	case NPI_Internet_IP:
		return "Internet (IP)"
	case NPI_WAPClientId:
		return "WAP Client Id"
	}
	return "Unknwon"
}
