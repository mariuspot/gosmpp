package smpp

import (
	"bytes"
	"errors"
)

type DeliverSmPDU struct {
	*SmPDU

	messageState *int8
}

func NewDeliverSmPDU() (*DeliverSmPDU, error) {
	pdu, err := NewSmPDU(CID_DeliverSm)
	if err != nil {
		return nil, err
	}
	return &DeliverSmPDU{SmPDU: pdu}, nil
}

func NewDeliverSmPDUFromBuffer(buf *bytes.Buffer) (*DeliverSmPDU, error) {
	pdu, err := NewSmPDUFromBuffer(buf)
	if err != nil {
		return nil, err
	}
	if pdu.GetCommandId() != CID_DeliverSm {
		return nil, errors.New("Not a DeliverSmPDU")
	}
	return &DeliverSmPDU{SmPDU: pdu}, nil
}

func NewDeliverSmPDUFromHeaderAndBuffer(header *Header, buf *bytes.Buffer) (*DeliverSmPDU, error) {
	pdu, err := NewSmPDUFromHeaderAndBuffer(header, buf)
	if err != nil {
		return nil, err
	}
	if pdu.GetCommandId() != CID_DeliverSm {
		return nil, errors.New("Not a DeliverSmPDU")
	}
	return &DeliverSmPDU{SmPDU: pdu}, nil
}

func (d *DeliverSmPDU) IsDeliveryReceipt() bool {
	return d.GetEsmClass()&ESM_CLASS_DeliverSm_SMSC_Delivery_Receipt == ESM_CLASS_DeliverSm_SMSC_Delivery_Receipt
}

func (d *DeliverSmPDU) GetReceiptMessageId() (messageId string, err error) {
	if tlv, err := d.GetTLV(TLV_receipted_message_id); err == nil {
		messageId, err = tlv.GetValueString()
	}
	return
}

func (d *DeliverSmPDU) GetMessageState() (messageState int8, err error) {
	if tlv, err := d.GetTLV(TLV_message_state); err == nil {
		messageState, err = tlv.GetValueInt8()
		d.messageState = &messageState
	}
	return
}

func (d *DeliverSmPDU) IsFinalState() bool {
	if d.messageState == nil {
		d.GetMessageState()
	}
	return d.messageState != nil && *d.messageState != TLVV_message_state_ENROUTE
}

func (d *DeliverSmPDU) IsDeliveredState() bool {
	if d.messageState == nil {
		d.GetMessageState()
	}
	return d.messageState != nil && *d.messageState == TLVV_message_state_DELIVERED
}

func (d *DeliverSmPDU) IsUndeliveredState() bool {
	if d.messageState == nil {
		d.GetMessageState()
	}
	return d.messageState != nil && *d.messageState != TLVV_message_state_DELIVERED && *d.messageState != TLVV_message_state_ENROUTE
}

// func (d *DeliverSmPDU) IsDelivered() bool {

// 	return d.GetMessageState() == TLVV_message_state_DELIVERED
// }

// if messageState != smpp.TLVV_message_state_ENROUTE && messageState != smpp.TLVV_message_state_ACCEPTED && messageState != smpp.TLVV_message_state_UNKNOWN {
// 								delivery := &adworkerMessages.Delivery{
// 									Id:         proto.String(fmt.Sprintf("%d", localMessageIdRoute.messageId)),
// 									Identifier: proto.String(typePdu.GetDestinationAddr()),
// 									Channel:    proto.String(localMessageIdRoute.channel),
// 									AdvertId:   proto.String(localMessageIdRoute.advertId),
// 								}
// 								if messageState == smpp.TLVV_message_state_DELIVERED
