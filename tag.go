package smpp

import (
	"fmt"
)

type Tag uint16

func (t Tag) String() string {
	return fmt.Sprintf("%d", t)
}

const (
	TLV_dest_addr_subunit           Tag = 0x0005 // 0x0005 GSM
	TLV_dest_network_type           Tag = 0x0006 // 0x0006 Generic
	TLV_dest_bearer_type            Tag = 0x0007 // 0x0007 Generic
	TLV_dest_telematics_id          Tag = 0x0008 // 0x0008 GSM
	TLV_source_addr_subunit         Tag = 0x000D // 0x000D GSM
	TLV_source_network_type         Tag = 0x000E // 0x000E Generic
	TLV_source_bearer_type          Tag = 0x000F // 0x000F Generic
	TLV_source_telematics_id        Tag = 0x0010 // 0x0010 GSM
	TLV_qos_time_to_live            Tag = 0x0017 // 0x0017 Generic
	TLV_payload_type                Tag = 0x0019 // 0x0019 Generic
	TLV_additional_status_info_text Tag = 0x001D // 0x001D Generic
	TLV_receipted_message_id        Tag = 0x001E // 0x001E Generic
	TLV_ms_msg_wait_facilities      Tag = 0x0030 // 0x0030 GSM
	TLV_privacy_indicator           Tag = 0x0201 // 0x0201 CDMA, TDMA
	TLV_source_subaddress           Tag = 0x0202 // 0x0202 CDMA, TDMA
	TLV_dest_subaddress             Tag = 0x0203 // 0x0203 CDMA, TDMA
	TLV_user_message_reference      Tag = 0x0204 // 0x0204 Generic
	TLV_user_response_code          Tag = 0x0205 // 0x0205 CDMA, TDMA
	TLV_source_port                 Tag = 0x020A // 0x020A Generic
	TLV_destination_port            Tag = 0x020B // 0x020B Generic
	TLV_sar_msg_ref_num             Tag = 0x020C // 0x020C Generic
	TLV_language_indicator          Tag = 0x020D // 0x020D CDMA, TDMA
	TLV_sar_total_segments          Tag = 0x020E // 0x020E Generic
	TLV_sar_segment_seqnum          Tag = 0x020F // 0x020F Generic
	TLV_SC_interface_version        Tag = 0x0210 // 0x0210 Generic
	TLV_callback_num_pres_ind       Tag = 0x0302 // 0x0302 TDMA
	TLV_callback_num_atag           Tag = 0x0303 // 0x0303 TDMA
	TLV_number_of_messages          Tag = 0x0304 // 0x0304 CDMA
	TLV_callback_num                Tag = 0x0381 // 0x0381 CDMA, TDMA, GSM, iDEN
	TLV_dpf_result                  Tag = 0x0420 // 0x0420 Generic
	TLV_set_dpf                     Tag = 0x0421 // 0x0421 Generic
	TLV_ms_availability_status      Tag = 0x0422 // 0x0422 Generic
	TLV_network_error_code          Tag = 0x0423 // 0x0423 Generic
	TLV_message_payload             Tag = 0x0424 // 0x0424 Generic
	TLV_delivery_failure_reason     Tag = 0x0425 // 0x0425 Generic
	TLV_more_messages_to_send       Tag = 0x0426 // 0x0426 GSM
	TLV_message_state               Tag = 0x0427 // 0x0427 Generic
	TLV_ussd_service_op             Tag = 0x0501 // 0x0501 GSM (USSD)
	TLV_display_time                Tag = 0x1201 // 0x1201 CDMA, TDMA
	TLV_sms_signal                  Tag = 0x1203 // 0x1203 TDMA
	TLV_ms_validity                 Tag = 0x1204 // 0x1204 CDMA, TDMA
	TLV_alert_on_message_delivery   Tag = 0x130C // 0x130C CDMA
	TLV_its_reply_type              Tag = 0x1380 // 0x1380 CDMA
	TLV_its_session_info            Tag = 0x1383 // 0x1383 CDMA
)

const (
	TLVV_message_state_ENROUTE       int8 = 1 //1 The message is in enroute state.
	TLVV_message_state_DELIVERED     int8 = 2 //2 Message is delivered to destination
	TLVV_message_state_EXPIRED       int8 = 3 //3 Message validity period has expired.
	TLVV_message_state_DELETED       int8 = 4 //4 Message has been deleted.
	TLVV_message_state_UNDELIVERABLE int8 = 5 //5 Message is undeliverable
	TLVV_message_state_ACCEPTED      int8 = 6 //6 Message is in accepted state (i.e. has been manually read on behalf of the subscriber by customer service)
	TLVV_message_state_UNKNOWN       int8 = 7 //7 Message is in invalid state
	TLVV_message_state_REJECTED      int8 = 8 //8 Message is in a rejected state
)
