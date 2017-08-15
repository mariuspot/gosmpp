package smpp

const (
	//SubmitSm
	ESM_CLASS_SubmitSm_Default_SMSC_Mode_INVERTED uint8 = 3 //x x x x x x 0 0	(e.g. Store and Forward)
	ESM_CLASS_SubmitSm_Datagram                   uint8 = 1 //x x x x x x 0 1
	ESM_CLASS_SubmitSm_Forward                    uint8 = 2 //x x x x x x 1 0
	ESM_CLASS_SubmitSm_Store_Forward              uint8 = 3 //x x x x x x 1 1
	//(use to select Store and Forward mode if Default SMSC Mode is non Store and Forward)

	//Message Type (bits 5-2)
	ESM_CLASS_SubmitSm_Default_Type                  uint8 = 60 //x x 0 0 0 0 x x	(i.e. normal message)
	ESM_CLASS_SubmitSm_ESME_Delivery_Acknowledgement uint8 = 8  //x x 0 0 1 0 x x
	ESM_CLASS_SubmitSm_ESME_Manual_Acknowledgement   uint8 = 16 //x x 0 1 0 0 x x

	//Acknowledgement
	//GSM Network Specific Features (bits 7-6)
	ESM_CLASS_SubmitSm_No_Features_INVERTED    uint8 = 192 //0 0 x x x x x x
	ESM_CLASS_SubmitSm_UDHI_Indicator          uint8 = 64  //0 1 x x x x x x	(only relevant for MT short messages)
	ESM_CLASS_SubmitSm_Set_Reply_Path          uint8 = 128 //1 0 x x x x x x	(only relevant for GSM network)
	ESM_CLASS_SubmitSm_Set_UDHI_and_Reply_Path uint8 = 196 //1 1 x x x x x x	(only relevant for GSM network)

	//DeliverSm
	ESM_CLASS_DeliverSm_Default_Type_INVERTED              uint8 = 60 //x x 0 0 0 0 x x	(i.e. normal message)
	ESM_CLASS_DeliverSm_SMSC_Delivery_Receipt              uint8 = 4  //x x 0 0 0 1 x x
	ESM_CLASS_DeliverSm_SME_Delivery_Acknowledgement       uint8 = 8  //x x 0 0 1 0 x x
	ESM_CLASS_DeliverSm_reserved                           uint8 = 12 //x x 0 0 1 1 x x
	ESM_CLASS_DeliverSm_SME_Manual_Acknowledgment          uint8 = 16 //x x 0 1 0 0 x x
	ESM_CLASS_DeliverSm_reserved_1                         uint8 = 20 //x x 0 1 0 1 x x
	ESM_CLASS_DeliverSm_Conversation_Abort                 uint8 = 24 //x x 0 1 1 0 x x	(Korean CDMA)
	ESM_CLASS_DeliverSm_reserved_2                         uint8 = 28 //x x 0 1 1 1 x x
	ESM_CLASS_DeliverSm_Intermediate_Delivery_Notification uint8 = 32 //x x 1 0 0 0 x x
)
