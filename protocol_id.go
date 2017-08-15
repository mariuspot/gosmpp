package smpp

// TP-Protocol-Identifier (TP-PID)
// The TP-Protocol-Identifier parameter serves the purposes indicated in section 3.2.3. It consists of one
// octet, and the bits in the octet are used as follows:
// The MS will not interpret reserved or unsupported values but shall store them as received.
// The SC may reject messages with a TP-Protocol-Identifier containing a reserved value or one which is not
// supported.
//
// bits usage
// 7 6
// 0 0 Assigns bits 0..5 as defined below
// 0 1 Assigns bits 0..5 as defined below
// 1 0 reserved
// 1 1 Assigns bits 0-5 for SC specific use
// In the case where bit 7 = 0 and bit 6 = 0,
// bit 5 indicates telematic interworking:
// value = 0 : no interworking, but SME-to-SME protocol
// value = 1 : telematic interworking
// In the case of telematic interworking, the following five bit patterns in bits 4..0 are used to indicate different
// types of telematic devices:
// 4...0
// 00000 implicit - device type is specific to this SC, or can be concluded on the basis of
// the address
// 00001 telex (or teletex reduced to telex format)
// 00010 group 3 telefax
// 00011 group 4 telefax
// 00100 voice telephone (i.e. conversion to speech)
// 00101 ERMES (European Radio Messaging System)
// 00110 National Paging system (known to the SC)
// 00111 Videotex (T.100/T.101)
// 01000 teletex, carrier unspecified
// 01001 teletex, in PSPDN
// 01010 teletex, in CSPDN
// 01011 teletex, in analogue PSTN
// 01100 teletex, in digital ISDN
// 01101 UCI (Universal Computer Interface, ETSI DE/PS 3 01-3)
// 01110..01111 (reserved, 2 combinations)
// 10000 a message handling facility (known to the SC)
// 10001 any public X.400-based message handling system
// 10010 Internet Electronic Mail
// 10011..10111 (reserved, 5 combinations)
// 11000..11110 values specific to each SC, usage based on mutual agreement between the
//
// SME and the SC (7 combinations available for each SC)
// 11111 A GSM mobile station. The SC converts the SM from the received
// TP-Data-Coding-Scheme to any data coding scheme supported by that MS (e.g.
// the default).
// If bit 5 has value 1 in an SMS-SUBMIT PDU, it indicates that the SME is a telematic device of a type which
// is indicated in bits 4..0, and requests the SC to convert the SM into a form suited for that device type. If
// the destination network is ISDN, the SC must also select the proper service indicators for connecting to a
// device of that type.
// If bit 5 has value 1 in an SMS-DELIVER PDU, it indicates that the SME is a telematic device of a type
// which is indicated in bits 4..0.
// If bit 5 has value 0 in an SMS-DELIVER PDU, the value in bits 4..0 identifies the SM-AL protocol being
// used between the SME and the MS.
// Note that for the straightforward case of simple MS-to-SC short message transfer the Protocol Identifier is
// set to the value 0.
//
// In the case where bit 7 = 0, bit 6 = 1, bits 5..0 are used as defined below
// 5....0
// 000000 Short Message Type 0
// 000001 Replace Short Message Type 1
// 000010 Replace Short Message Type 2
// 000011 Replace Short Message Type 3
// 000100 Replace Short Message Type 4
// 000101 Replace Short Message Type 5
// 000110 Replace Short Message Type 6
// 000111 Replace Short Message Type 7
// 001000..011110 Reserved
// 011111 Return Call Message
// 100000..111110 Reserved
// 111111 SIM Data download
//
// If bit 5 has value 1 in an SMS-SUBMIT PDU, it indicates that the SME is a telematic device of a type which
// is indicated in bits 4..0, and requests the SC to convert the SM into a form suited for that device type. If
// the destination network is ISDN, the SC must also select the proper service indicators for connecting to a
// device of that type.
// If bit 5 has value 1 in an SMS-DELIVER PDU, it indicates that the SME is a telematic device of a type
// which is indicated in bits 4..0.
// If bit 5 has value 0 in an SMS-DELIVER PDU, the value in bits 4..0 identifies the SM-AL protocol being
// used between the SME and the MS.
// Note that for the straightforward case of simple MS-to-SC short message transfer the Protocol Identifier is
// set to the value 0.

const (
	PROTOCOL_ID_Short_Message_Type_0         uint8 = 64 // 01000000 Short Message Type 0
	PROTOCOL_ID_Replace_Short_Message_Type_1 uint8 = 65 // 01000001 Replace Short Message Type 1
	PROTOCOL_ID_Replace_Short_Message_Type_2 uint8 = 66 // 01000010 Replace Short Message Type 2
	PROTOCOL_ID_Replace_Short_Message_Type_3 uint8 = 67 // 01000011 Replace Short Message Type 3
	PROTOCOL_ID_Replace_Short_Message_Type_4 uint8 = 68 // 01000100 Replace Short Message Type 4
	PROTOCOL_ID_Replace_Short_Message_Type_5 uint8 = 69 // 01000101 Replace Short Message Type 5
	PROTOCOL_ID_Replace_Short_Message_Type_6 uint8 = 70 // 01000110 Replace Short Message Type 6
	PROTOCOL_ID_Replace_Short_Message_Type_7 uint8 = 71 // 01000111 Replace Short Message Type 7
	// // 01001000..01011110 Reserved
	PROTOCOL_ID_Retrun_Call_Message uint8 = 95 // 01011111 Return Call Message
	// // 01100000..01111110 Reserved
	PROTOCOL_ID_Sim_Data_Download uint8 = 127 // 01111111 SIM Data download

)
