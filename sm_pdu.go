package smpp

import (
	"bufio"
	"bytes"
	"fmt"
)

type SmPDU struct {
	*Header
	serviceType          string
	sourceAddrTon        TON
	sourceAddrNpi        NPI
	sourceAddr           string
	destAddrTon          TON
	destAddrNpi          NPI
	destinationAddr      string
	esmClass             uint8
	protocolId           uint8
	priorityFlag         uint8
	scheduleDeliveryTime string
	validityPeriod       string
	registeredDelivery   uint8
	replaceIfPresentFlag uint8

	dataCoding     uint8
	smDefaultMsgId uint8
	smLength       uint8
	shortMessage   string
	*TLVList
}

func (s SmPDU) String() string {
	return fmt.Sprintf("{Header: %s, serviceType: %s, sourceAddrTon: %s, sourceAddrNpi: %s, sourceAddr: %s, destAddrTon: %s, destAddrNpi: %s, destinationAddr: %s, esmClass: %d, protocolId: %d, priorityFlag: %d, scheduleDeliveryTime: %s, validityPeriod: %s, registeredDelivery: %d, replaceIfPresentFlag: %d, dataCoding: %d, smDefaultMsgId: %d, smLength: %d, shortMessage: %s, TLVList: %s}", s.Header, s.serviceType, s.sourceAddrTon, s.sourceAddrNpi, s.sourceAddr, s.destAddrTon, s.destAddrNpi, s.destinationAddr, s.esmClass, s.protocolId, s.priorityFlag, s.scheduleDeliveryTime, s.validityPeriod, s.registeredDelivery, s.replaceIfPresentFlag, s.dataCoding, s.smDefaultMsgId, s.smLength, s.shortMessage, s.TLVList)
}

func NewSmPDU(commandId CommandId) (*SmPDU, error) {
	return &SmPDU{Header: NewHeader(commandId), TLVList: NewTLVList()}, nil
}

func NewSmPDUFromBuffer(buf *bytes.Buffer) (*SmPDU, error) {
	header, err := NewHeaderFromBuffer(buf)
	if err != nil {
		return nil, err
	}
	return NewSmPDUFromHeaderAndBuffer(header, buf)
}

func NewSmPDUFromHeaderAndBuffer(header *Header, buf *bytes.Buffer) (*SmPDU, error) {

	// fmt.Println("2")
	pdu := &SmPDU{Header: header}
	var err error
	if pdu.serviceType, err = buf.ReadString(0); err != nil {
		return nil, err
	}
	pdu.serviceType = pdu.serviceType[:len(pdu.serviceType)-1]
	var b byte
	if b, err = buf.ReadByte(); err != nil {
		return nil, err
	}
	pdu.sourceAddrTon = TON(b)
	if b, err = buf.ReadByte(); err != nil {
		return nil, err
	}
	pdu.sourceAddrNpi = NPI(b)
	if pdu.sourceAddr, err = buf.ReadString(0); err != nil {
		return nil, err
	}
	pdu.sourceAddr = pdu.sourceAddr[:len(pdu.sourceAddr)-1]
	if b, err = buf.ReadByte(); err != nil {
		return nil, err
	}
	pdu.destAddrTon = TON(b)
	if b, err = buf.ReadByte(); err != nil {
		return nil, err
	}
	pdu.destAddrNpi = NPI(b)
	if pdu.destinationAddr, err = buf.ReadString(0); err != nil {
		return nil, err
	}
	pdu.destinationAddr = pdu.destinationAddr[:len(pdu.destinationAddr)-1]
	if pdu.esmClass, err = buf.ReadByte(); err != nil {
		return nil, err
	}
	if pdu.protocolId, err = buf.ReadByte(); err != nil {
		return nil, err
	}
	if pdu.priorityFlag, err = buf.ReadByte(); err != nil {
		return nil, err
	}
	if pdu.scheduleDeliveryTime, err = buf.ReadString(0); err != nil {
		return nil, err
	}
	pdu.scheduleDeliveryTime = pdu.scheduleDeliveryTime[:len(pdu.scheduleDeliveryTime)-1]
	if pdu.validityPeriod, err = buf.ReadString(0); err != nil {
		return nil, err
	}
	pdu.validityPeriod = pdu.validityPeriod[:len(pdu.validityPeriod)-1]
	if pdu.registeredDelivery, err = buf.ReadByte(); err != nil {
		return nil, err
	}
	if pdu.replaceIfPresentFlag, err = buf.ReadByte(); err != nil {
		return nil, err
	}
	if pdu.dataCoding, err = buf.ReadByte(); err != nil {
		return nil, err
	}
	if pdu.smDefaultMsgId, err = buf.ReadByte(); err != nil {
		return nil, err
	}
	if pdu.smLength, err = buf.ReadByte(); err != nil {
		return nil, err
	}
	var tempBuf bytes.Buffer
	for i := uint8(0); i < pdu.smLength; i++ {
		b, err := buf.ReadByte()
		if err != nil {
			return nil, err
		}
		if err := tempBuf.WriteByte(b); err != nil {
			return nil, err
		}
	}
	pdu.shortMessage = string(tempBuf.Bytes())
	tlvList, err := NewTLVListFromBuffer(buf)
	if err != nil {
		return nil, err
	}
	pdu.TLVList = tlvList
	return pdu, nil
}

func (s *SmPDU) CalculateCommandLength() uint32 {
	return 16 + uint32(len(s.serviceType)+1+1+1+len(s.sourceAddr)+1+1+1+len(s.destinationAddr)+1+1+1+1+len(s.scheduleDeliveryTime)+1+len(s.validityPeriod)+1+1+1+1+1+1+len(s.shortMessage)) + s.TLVList.GetLength()
}

func (s *SmPDU) GetServiceType() string {
	return s.serviceType
}

func (s *SmPDU) SetServiceType(serviceType string) {
	s.serviceType = serviceType
}

func (s *SmPDU) GetSourceAddrTon() TON {
	return s.sourceAddrTon
}

func (s *SmPDU) SetSourceAddrTon(sourceAddrTon TON) {
	s.sourceAddrTon = sourceAddrTon
}

func (s *SmPDU) GetSourceAddrNpi() NPI {
	return s.sourceAddrNpi
}

func (s *SmPDU) SetSourceAddrNpi(sourceAddrNpi NPI) {
	s.sourceAddrNpi = sourceAddrNpi
}

func (s *SmPDU) GetSourceAddr() string {
	return s.sourceAddr
}

func (s *SmPDU) SetSourceAddr(sourceAddr string) {
	s.sourceAddr = sourceAddr
}

func (s *SmPDU) GetDestAddrTon() TON {
	return s.destAddrTon
}

func (s *SmPDU) SetDestAddrTon(destAddrTon TON) {
	s.destAddrTon = destAddrTon
}

func (s *SmPDU) GetDestAddrNpi() NPI {
	return s.destAddrNpi
}

func (s *SmPDU) SetDestAddrNpi(destAddrNpi NPI) {
	s.destAddrNpi = destAddrNpi
}

func (s *SmPDU) GetDestinationAddr() string {
	return s.destinationAddr
}

func (s *SmPDU) SetDestinationAddr(destinationAddr string) {
	s.destinationAddr = destinationAddr
}

func (s *SmPDU) GetEsmClass() uint8 {
	return s.esmClass
}

func (s *SmPDU) SetEsmClass(esmClass uint8) {
	s.esmClass = esmClass
}

func (s *SmPDU) GetProtocolId() uint8 {
	return s.protocolId
}

func (s *SmPDU) SetProtocolId(protocolId uint8) {
	s.protocolId = protocolId
}

func (s *SmPDU) GetPriorityFlag() uint8 {
	return s.priorityFlag
}

func (s *SmPDU) SetPriorityFlag(priorityFlag uint8) {
	s.priorityFlag = priorityFlag
}

func (s *SmPDU) GetScheduleDeliveryTime() string {
	return s.scheduleDeliveryTime
}

func (s *SmPDU) SetScheduleDeliveryTime(scheduleDeliveryTime string) {
	s.scheduleDeliveryTime = scheduleDeliveryTime
}

func (s *SmPDU) GetValidityPeriod() string {
	return s.validityPeriod
}

func (s *SmPDU) SetValidityPeriod(validityPeriod string) {
	s.validityPeriod = validityPeriod
}

func (s *SmPDU) GetRegisteredDelivery() uint8 {
	return s.registeredDelivery
}

func (s *SmPDU) SetRegisteredDelivery(registeredDelivery uint8) {
	s.registeredDelivery = registeredDelivery
}

func (s *SmPDU) GetReplaceIfPresentFlag() uint8 {
	return s.replaceIfPresentFlag
}

func (s *SmPDU) SetReplaceIfPresentFlag(replaceIfPresentFlag uint8) {
	s.replaceIfPresentFlag = replaceIfPresentFlag
}

func (s *SmPDU) GetDataCoding() uint8 {
	return s.dataCoding
}

func (s *SmPDU) SetDataCoding(dataCoding uint8) {
	s.dataCoding = dataCoding
}

func (s *SmPDU) GetSmDefaultMsgId() uint8 {
	return s.smDefaultMsgId
}

func (s *SmPDU) SetSmDefaultMsgId(smDefaultMsgId uint8) {
	s.smDefaultMsgId = smDefaultMsgId
}

func (s *SmPDU) GetSmLength() uint8 {
	return s.smLength
}

func (s *SmPDU) GetShortMessage() string {
	return s.shortMessage
}

func (s *SmPDU) SetShortMessage(shortMessage string) {
	s.smLength = uint8(len(shortMessage))
	s.shortMessage = shortMessage
}

func (s *SmPDU) WriteToBuffer(buf *bufio.Writer) error {
	s.commandLength = s.CalculateCommandLength()

	if err := s.Header.WriteToBuffer(buf); err != nil {
		return err
	}
	if _, err := buf.WriteString(s.serviceType); err != nil {
		return err
	}
	if err := buf.WriteByte(0); err != nil {
		return err
	}
	if err := buf.WriteByte(byte(s.sourceAddrTon)); err != nil {
		return err
	}
	if err := buf.WriteByte(byte(s.sourceAddrNpi)); err != nil {
		return err
	}
	if _, err := buf.WriteString(s.sourceAddr); err != nil {
		return err
	}
	if err := buf.WriteByte(0); err != nil {
		return err
	}
	if err := buf.WriteByte(byte(s.destAddrTon)); err != nil {
		return err
	}
	if err := buf.WriteByte(byte(s.destAddrNpi)); err != nil {
		return err
	}
	if _, err := buf.WriteString(s.destinationAddr); err != nil {
		return err
	}
	if err := buf.WriteByte(0); err != nil {
		return err
	}
	if err := buf.WriteByte(s.esmClass); err != nil {
		return err
	}
	if err := buf.WriteByte(s.protocolId); err != nil {
		return err
	}
	if err := buf.WriteByte(s.priorityFlag); err != nil {
		return err
	}
	if _, err := buf.WriteString(s.scheduleDeliveryTime); err != nil {
		return err
	}
	if err := buf.WriteByte(0); err != nil {
		return err
	}
	if _, err := buf.WriteString(s.validityPeriod); err != nil {
		return err
	}
	if err := buf.WriteByte(0); err != nil {
		return err
	}
	if err := buf.WriteByte(s.registeredDelivery); err != nil {
		return err
	}
	if err := buf.WriteByte(s.replaceIfPresentFlag); err != nil {
		return err
	}
	if err := buf.WriteByte(s.dataCoding); err != nil {
		return err
	}
	if err := buf.WriteByte(s.smDefaultMsgId); err != nil {
		return err
	}
	if err := buf.WriteByte(s.smLength); err != nil {
		return err
	}
	if _, err := buf.WriteString(s.shortMessage); err != nil {
		return err
	}
	if err := s.TLVList.WriteToBuffer(buf); err != nil {
		return err
	}
	return nil
}
