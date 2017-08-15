package smpp

import (
	"bytes"
	"errors"
)

type SubmitSmPDU struct {
	*SmPDU
}

func NewSubmitSmPDU() (*SubmitSmPDU, error) {
	pdu, err := NewSmPDU(CID_SubmitSm)
	if err != nil {
		return nil, err
	}
	return &SubmitSmPDU{SmPDU: pdu}, nil
}

func NewSubmitSmPDUFromBuffer(buf *bytes.Buffer) (*SubmitSmPDU, error) {
	pdu, err := NewSmPDUFromBuffer(buf)
	if err != nil {
		return nil, err
	}
	if pdu.GetCommandId() != CID_SubmitSm {
		return nil, errors.New("Not a SubmitSmPDU")
	}
	return &SubmitSmPDU{SmPDU: pdu}, nil
}

func NewSubmitSmPDUFromHeaderAndBuffer(header *Header, buf *bytes.Buffer) (*SubmitSmPDU, error) {

	// fmt.Println("1")
	pdu, err := NewSmPDUFromHeaderAndBuffer(header, buf)
	if err != nil {
		return nil, err
	}
	if pdu.GetCommandId() != CID_SubmitSm {
		return nil, errors.New("Not a SubmitSmPDU")
	}
	return &SubmitSmPDU{SmPDU: pdu}, nil
}

func (s *SubmitSmPDU) HasUserDataHeader() bool {
	return s.esmClass&ESM_CLASS_SubmitSm_UDHI_Indicator == ESM_CLASS_SubmitSm_UDHI_Indicator
}

func (s *SubmitSmPDU) GetUserDataHeader() (udh *UserDataHeader) {
	// fmt.Println([]byte(s.shortMessage))
	if s.smLength > 0 && len(s.shortMessage) > 0 {
		length := byte(s.shortMessage[0]) + 1
		udh = NewUserDataHeader()
		udh.Decode([]byte(s.shortMessage[0:length]))
	}
	return
}
