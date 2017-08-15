package smpp

type CommandId uint32

const (
	CID_BindReceiver              CommandId = 0x00000001
	CID_BindTransmitter                     = 0x00000002
	CID_QuerySm                             = 0x00000003
	CID_SubmitSm                            = 0x00000004
	CID_DeliverSm                           = 0x00000005
	CID_Unbind                              = 0x00000006
	CID_ReplaceSm                           = 0x00000007
	CID_CancelSm                            = 0x00000008
	CID_BindTransceiver                     = 0x00000009
	CID_Outbind                             = 0x0000000B
	CID_EnquireLink                         = 0x00000015
	CID_SubmitMulti                         = 0x00000021
	CID_AlertNotification                   = 0x00000102
	CID_DataSm                              = 0x00000103
	CID_BroadcastSm                         = 0x00000111
	CID_QueryBroadcastSm                    = 0x00000112
	CID_CancelBroadcastSm                   = 0x00000113
	CID_GenericNack                         = 0x80000000
	CID_BindReceiverResponse                = 0x80000001
	CID_BindTransmitterResponse             = 0x80000002
	CID_QuerySmResponse                     = 0x80000003
	CID_SubmitSmResponse                    = 0x80000004
	CID_DeliverSmResponse                   = 0x80000005
	CID_UnbindResponse                      = 0x80000006
	CID_ReplaceSmResponse                   = 0x80000007
	CID_CancelSmResponse                    = 0x80000008
	CID_BindTransceiverResponse             = 0x80000009
	CID_EnquireLinkResponse                 = 0x80000015
	CID_SubmitMultiResponse                 = 0x80000021
	CID_DataSmResponse                      = 0x80000103
	CID_BroadcastSmResponse                 = 0x80000111
	CID_QueryBroadcastSmResponse            = 0x80000112
	CID_CancelBroadcastSmResponse           = 0x80000113
)

func (c CommandId) String() string {
	switch c {
	case CID_BindReceiver:
		return "BindReceiver [0x00000001]"
	case CID_BindTransmitter:
		return "BindTransmitter [0x00000002]"
	case CID_QuerySm:
		return "QuerySm [0x00000003]"
	case CID_SubmitSm:
		return "SubmitSm [0x00000004]"
	case CID_DeliverSm:
		return "DeliverSm [0x00000005]"
	case CID_Unbind:
		return "Unbind [0x00000006]"
	case CID_ReplaceSm:
		return "ReplaceSm [0x00000007]"
	case CID_CancelSm:
		return "CancelSm [0x00000008]"
	case CID_BindTransceiver:
		return "BindTransceiver [0x00000009]"
	case CID_Outbind:
		return "Outbind [0x0000000B]"
	case CID_EnquireLink:
		return "EnquireLink [0x00000015]"
	case CID_SubmitMulti:
		return "SubmitMulti [0x00000021]"
	case CID_AlertNotification:
		return "AlertNotification [0x00000102]"
	case CID_DataSm:
		return "DataSm [0x00000103]"
	case CID_BroadcastSm:
		return "BroadcastSm [0x00000111]"
	case CID_QueryBroadcastSm:
		return "QueryBroadcastSm [0x00000112]"
	case CID_CancelBroadcastSm:
		return "CancelBroadcastSm [0x00000113]"
	case CID_GenericNack:
		return "GenericNack [0x80000000]"
	case CID_BindReceiverResponse:
		return "BindReceiverResponse [0x80000001]"
	case CID_BindTransmitterResponse:
		return "BindTransmitterResponse [0x80000002]"
	case CID_QuerySmResponse:
		return "QuerySmResponse [0x80000003]"
	case CID_SubmitSmResponse:
		return "SubmitSmResponse [0x80000004]"
	case CID_DeliverSmResponse:
		return "DeliverSmResponse [0x80000005]"
	case CID_UnbindResponse:
		return "UnbindResponse [0x80000006]"
	case CID_ReplaceSmResponse:
		return "ReplaceSmResponse [0x80000007]"
	case CID_CancelSmResponse:
		return "CancelSmResponse [0x80000008]"
	case CID_BindTransceiverResponse:
		return "BindTransceiverResponse [0x80000009]"
	case CID_EnquireLinkResponse:
		return "EnquireLinkResponse [0x80000015]"
	case CID_SubmitMultiResponse:
		return "SubmitMultiResponse [0x80000021]"
	case CID_DataSmResponse:
		return "DataSmResponse [0x80000103]"
	case CID_BroadcastSmResponse:
		return "BroadcastSmResponse [0x80000111]"
	case CID_QueryBroadcastSmResponse:
		return "QueryBroadcastSmResponse [0x80000112]"
	case CID_CancelBroadcastSmResponse:
		return "CancelBroadcastSmResponse [0x80000113]"
	}
	return "UNKNWON"
}
