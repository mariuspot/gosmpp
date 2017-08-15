package smpp

import (
	"fmt"
)

type CommandStatus uint32

const (
	CS_ESME_ROK              CommandStatus = 0x00000000 //No Error
	CS_ESME_RINVMSGLEN                     = 0x00000001 //Message Length is invalid
	CS_ESME_RINVCMDLEN                     = 0x00000002 //Command Length is invalid
	CS_ESME_RINVCMDID                      = 0x00000003 //Invalid Command ID
	CS_ESME_RINVBNDSTS                     = 0x00000004 //Incorrect BIND Status for given com-mand
	CS_ESME_RALYBND                        = 0x00000005 //ESME Already in Bound State
	CS_ESME_RINVPRTFLG                     = 0x00000006 //Invalid Priority Flag
	CS_ESME_RINVREGDLVFLG                  = 0x00000007 //Invalid Registered Delivery Flag
	CS_ESME_RSYSERR                        = 0x00000008 //System Error
	_                                      = 0          // Reserved 0x00000009 Reserved
	CS_ESME_RINVSRCADR                     = 0x0000000A //Invalid Source Address
	CS_ESME_RINVDSTADR                     = 0x0000000B //Invalid Dest Addr
	CS_ESME_RINVMSGID                      = 0x0000000C //Message ID is invalid
	CS_ESME_RBINDFAIL                      = 0x0000000D //Bind Failed
	CS_ESME_RINVPASWD                      = 0x0000000E //Invalid Password
	CS_ESME_RINVSYSID                      = 0x0000000F //Invalid System ID
	_                                      = 0          // Reserved 0x00000010 Reserved
	CS_ESME_RCANCELFAIL                    = 0x00000011 //Cancel SM Failed
	_                                      = 0          // Reserved 0x00000012 Reserved
	CS_ESME_RREPLACEFAIL                   = 0x00000013 //Replace SM Failed
	CS_ESME_RMSGQFUL                       = 0x00000014 //Message Queue Full
	CS_ESME_RINVSERTYP                     = 0x00000015 //Invalid Service Type
	_                                      = 0          // Reserved 0x00000016-0x00000032 Reserved
	CS_ESME_RINVNUMDESTS                   = 0x00000033 //Invalid number of destinations
	CS_ESME_RINVDLNAME                     = 0x00000034 //Invalid Distribution List name
	_                                      = 0          // Reserved 0x00000035-0x0000003F Reserved
	CS_ESME_RINVDESTFLAG                   = 0x00000040 //Destination flag (submit_multi)
	_                                      = 0          // Reserved 0x00000041 Reserved
	CS_ESME_RINVSUBREP                     = 0x00000042 //Invalid ‘submit with replace’ request (i.e. submit_sm with replace_if_present_flag set)
	CS_ESME_RINVESMCLASS                   = 0x00000043 //Invalid esm_class field data
	CS_ESME_RCNTSUBDL                      = 0x00000044 //Cannot Submit to Distribution List
	CS_ESME_RSUBMITFAIL                    = 0x00000045 //submit_sm or submit_multi failed
	_                                      = 0          // Reserved 0x00000046-0x00000047 Reserved
	CS_ESME_RINVSRCTON                     = 0x00000048 //Invalid Source address TON
	CS_ESME_RINVSRCNPI                     = 0x00000049 //Invalid Source address NPI
	CS_ESME_RINVDSTTON                     = 0x00000050 //Invalid Destination address TON
	CS_ESME_RINVDSTNPI                     = 0x00000051 //Invalid Destination address NPI
	_                                      = 0          // Reserved 0x00000052 Reserved
	CS_ESME_RINVSYSTYP                     = 0x00000053 //Invalid system_type field
	CS_ESME_RINVREPFLAG                    = 0x00000054 //Invalid replace_if_present flag
	CS_ESME_RINVNUMMSGS                    = 0x00000055 //Invalid number of messages
	_                                      = 0          // Reserved 0x00000056-0x00000057 Reserved
	CS_ESME_RTHROTTLED                     = 0x00000058 //Throttling error (ESME has exceeded allowed message limits)
	_                                      = 0          // Reserved 0x00000059-0x00000060 Reserved
	CS_ESME_RINVSCHED                      = 0x00000061 //Invalid Scheduled Delivery Time
	CS_ESME_RINVEXPIRY                     = 0x00000062 //Invalid message (Expiry time)
	CS_ESME_RINVDFTMSGID                   = 0x00000063 //Predefined Message Invalid or Not Found
	CS_ESME_RX_T_APPN                      = 0x00000064 //ESME Receiver Temporary App Error Code
	CS_ESME_RX_P_APPN                      = 0x00000065 //ESME Receiver Permanent App Error Code
	CS_ESME_RX_R_APPN                      = 0x00000066 //ESME Receiver Reject Message Error Code
	CS_ESME_RQUERYFAIL                     = 0x00000067 //query_sm request failed
	_                                      = 0          // Reserved 0x00000068-0x000000BF Reserved
	CS_ESME_RINVOPTPARSTREAM               = 0x000000C0 //Error in the optional part of the PDU Body.
	CS_ESME_ROPTPARNOTALLWD                = 0x000000C1 //Optional Parameter not allowed
	CS_ESME_RINVPARLEN                     = 0x000000C2 //Invalid Parameter Length.
	CS_ESME_RMISSINGOPTPARAM               = 0x000000C3 //Expected Optional Parameter missing
	CS_ESME_RINVOPTPARAMVAL                = 0x000000C4 //Invalid Optional Parameter Value
	_                                      = 0          // Reserved 0x000000C5-0x000000FD Reserved
	CS_ESME_RDELIVERYFAILURE               = 0x000000FE //Delivery Failure (used for data_sm_resp)
	CS_ESME_RUNKNOWNERR                    = 0x000000FF //Unknown Error
	_                                                   // Reserved for SMPP extension 0x00000100-0x000003FF Reserved for SMPP extension
	_                                                   // Reserved for SMSC vendorspecific errors 0x00000400-0x000004FF Reserved for SMSC vendor specific errors
	_                                                   // Reserved 0x00000500-0xFFFFFFFF Reserve
)

func (c CommandStatus) String() string {
	switch c {
	case CS_ESME_ROK:
		return "No Error [CS_ESME_ROK]"
	case CS_ESME_RINVMSGLEN:
		return "Message Length is invalid [CS_ESME_RINVMSGLEN]"
	case CS_ESME_RINVCMDLEN:
		return "Command Length is invalid [CS_ESME_RINVCMDLEN]"
	case CS_ESME_RINVCMDID:
		return "Invalid Command ID [CS_ESME_RINVCMDID]"
	case CS_ESME_RINVBNDSTS:
		return "Incorrect BIND Status for given com-mand [CS_ESME_RINVBNDSTS]"
	case CS_ESME_RALYBND:
		return "ESME Already in Bound State [CS_ESME_RALYBND]"
	case CS_ESME_RINVPRTFLG:
		return "Invalid Priority Flag [CS_ESME_RINVPRTFLG]"
	case CS_ESME_RINVREGDLVFLG:
		return "Invalid Registered Delivery Flag [CS_ESME_RINVREGDLVFLG]"
	case CS_ESME_RSYSERR:
		return "System Error [CS_ESME_RSYSERR]"
	// Reserved 0x00000009 Reserved
	case CS_ESME_RINVSRCADR:
		return "Invalid Source Address [CS_ESME_RINVSRCADR]"
	case CS_ESME_RINVDSTADR:
		return "Invalid Dest Addr [CS_ESME_RINVDSTADR]"
	case CS_ESME_RINVMSGID:
		return "Message ID is invalid [CS_ESME_RINVMSGID]"
	case CS_ESME_RBINDFAIL:
		return "Bind Failed [CS_ESME_RBINDFAIL]"
	case CS_ESME_RINVPASWD:
		return "Invalid Password [CS_ESME_RINVPASWD]"
	case CS_ESME_RINVSYSID:
		return "Invalid System ID [CS_ESME_RINVSYSID]"
	// Reserved 0x00000010 Reserved
	case CS_ESME_RCANCELFAIL:
		return "Cancel SM Failed [CS_ESME_RCANCELFAIL]"
	// Reserved 0x00000012 Reserved
	case CS_ESME_RREPLACEFAIL:
		return "Replace SM Failed [CS_ESME_RREPLACEFAIL]"
	case CS_ESME_RMSGQFUL:
		return "Message Queue Full [CS_ESME_RMSGQFUL]"
	case CS_ESME_RINVSERTYP:
		return "Invalid Service Type [CS_ESME_RINVSERTYP]"
	// Reserved 0x00000016-0x00000032 Reserved
	case CS_ESME_RINVNUMDESTS:
		return "Invalid number of destinations [CS_ESME_RINVNUMDESTS]"
	case CS_ESME_RINVDLNAME:
		return "Invalid Distribution List name [CS_ESME_RINVDLNAME]"
	// Reserved 0x00000035-0x0000003F Reserved
	case CS_ESME_RINVDESTFLAG:
		return "Destination flag (submit_multi) [CS_ESME_RINVDESTFLAG]"
	// Reserved 0x00000041 Reserved
	case CS_ESME_RINVSUBREP:
		return "Invalid ‘submit with replace’ request (i.e. submit_sm with replace_if_present_flag set) [CS_ESME_RINVSUBREP]"
	case CS_ESME_RINVESMCLASS:
		return "Invalid esm_class field data [CS_ESME_RINVESMCLASS]"
	case CS_ESME_RCNTSUBDL:
		return "Cannot Submit to Distribution List [CS_ESME_RCNTSUBDL]"
	case CS_ESME_RSUBMITFAIL:
		return "submit_sm or submit_multi failed [CS_ESME_RSUBMITFAIL]"
	// Reserved 0x00000046-0x00000047 Reserved
	case CS_ESME_RINVSRCTON:
		return "Invalid Source address TON [CS_ESME_RINVSRCTON]"
	case CS_ESME_RINVSRCNPI:
		return "Invalid Source address NPI [CS_ESME_RINVSRCNPI]"
	case CS_ESME_RINVDSTTON:
		return "Invalid Destination address TON [CS_ESME_RINVDSTTON]"
	case CS_ESME_RINVDSTNPI:
		return "Invalid Destination address NPI [CS_ESME_RINVDSTNPI]"
	// Reserved 0x00000052 Reserved
	case CS_ESME_RINVSYSTYP:
		return "Invalid system_type field [CS_ESME_RINVSYSTYP]"
	case CS_ESME_RINVREPFLAG:
		return "Invalid replace_if_present flag [CS_ESME_RINVREPFLAG]"
	case CS_ESME_RINVNUMMSGS:
		return "Invalid number of messages [CS_ESME_RINVNUMMSGS]"
	// Reserved 0x00000056-0x00000057 Reserved
	case CS_ESME_RTHROTTLED:
		return "Throttling error (ESME has exceeded allowed message limits) [CS_ESME_RTHROTTLED]"
	// Reserved 0x00000059-0x00000060 Reserved
	case CS_ESME_RINVSCHED:
		return "Invalid Scheduled Delivery Time [CS_ESME_RINVSCHED]"
	case CS_ESME_RINVEXPIRY:
		return "Invalid message (Expiry time) [CS_ESME_RINVEXPIRY]"
	case CS_ESME_RINVDFTMSGID:
		return "Predefined Message Invalid or Not Found [CS_ESME_RINVDFTMSGID]"
	case CS_ESME_RX_T_APPN:
		return "ESME Receiver Temporary App Error Code [CS_ESME_RX_T_APPN]"
	case CS_ESME_RX_P_APPN:
		return "ESME Receiver Permanent App Error Code [CS_ESME_RX_P_APPN]"
	case CS_ESME_RX_R_APPN:
		return "ESME Receiver Reject Message Error Code [CS_ESME_RX_R_APPN]"
	case CS_ESME_RQUERYFAIL:
		return "query_sm request failed [CS_ESME_RQUERYFAIL]"
	// Reserved 0x00000068-0x000000BF Reserved
	case CS_ESME_RINVOPTPARSTREAM:
		return "Error in the optional part of the PDU Body. [CS_ESME_RINVOPTPARSTREAM]"
	case CS_ESME_ROPTPARNOTALLWD:
		return "Optional Parameter not allowed [CS_ESME_ROPTPARNOTALLWD]"
	case CS_ESME_RINVPARLEN:
		return "Invalid Parameter Length. [CS_ESME_RINVPARLEN]"
	case CS_ESME_RMISSINGOPTPARAM:
		return "Expected Optional Parameter missing [CS_ESME_RMISSINGOPTPARAM]"
	case CS_ESME_RINVOPTPARAMVAL:
		return "Invalid Optional Parameter Value [CS_ESME_RINVOPTPARAMVAL]"
	// Reserved 0x000000C5-0x000000FD Reserved
	case CS_ESME_RDELIVERYFAILURE:
		return "Delivery Failure (used for data_sm_resp) [CS_ESME_RDELIVERYFAILURE]"
	case CS_ESME_RUNKNOWNERR:
		return "Unknown Error [CS_ESME_RUNKNOWNERR]"
		// Reserved for SMPP extension 0x00000100-0x000003FF Reserved for SMPP extension
		// Reserved for SMSC vendorspecific errors 0x00000400-0x000004FF Reserved for SMSC vendor specific errors
		// Reserved 0x00000500-0xFFFFFFFF Reserve
	default:
		return fmt.Sprintf("UNKNWON %x", uint32(c))
	}

}
