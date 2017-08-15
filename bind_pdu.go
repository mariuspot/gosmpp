package smpp

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
)

type BindPDU struct {
	*Header
	systemId         string
	password         string
	systemType       string
	interfaceVersion byte
	addrTon          TON
	addrNpi          NPI
	addressRange     string
}

func (b BindPDU) String() string {
	return fmt.Sprintf("{Header: %s, systemId: %s, password: %s, systemType: %s, interfaceVersion: %d, addrTon: %d, addrNpi: %d, addressRange: %s}", b.Header, b.systemId, b.password, b.systemType, b.interfaceVersion, b.addrTon, b.addrNpi, b.addressRange)
}

func NewBindPDU(commandId CommandId, systemId string, password string, systemType string) (*BindPDU, error) {
	if len(systemId) > 16 {
		return nil, errors.New("systemId max length 16")
	}
	if len(password) > 9 {
		return nil, errors.New("password max length 9")
	}
	if len(systemType) > 13 {
		return nil, errors.New("systemType max length 13")
	}
	return &BindPDU{Header: NewHeader(commandId), systemId: systemId, password: password, systemType: systemType, interfaceVersion: 0x34}, nil
}

func NewBindPDUFromBuffer(buf *bytes.Buffer) (*BindPDU, error) {
	header, err := NewHeaderFromBuffer(buf)
	if err != nil {
		return nil, err
	}
	return NewBindPDUFromHeaderAndBuffer(header, buf)
}
func NewBindPDUFromHeaderAndBuffer(header *Header, buf *bytes.Buffer) (*BindPDU, error) {
	pdu := &BindPDU{Header: header}
	var err error
	if pdu.systemId, err = buf.ReadString(0); err != nil {
		return nil, err
	}
	pdu.systemId = pdu.systemId[:len(pdu.systemId)-1]
	if pdu.password, err = buf.ReadString(0); err != nil {
		return nil, err
	}
	pdu.password = pdu.password[:len(pdu.password)-1]
	if pdu.systemType, err = buf.ReadString(0); err != nil {
		return nil, err
	}
	pdu.systemType = pdu.systemType[:len(pdu.systemType)-1]
	if err = binary.Read(buf, binary.BigEndian, &pdu.interfaceVersion); err != nil {
		return nil, err
	}
	if err = binary.Read(buf, binary.BigEndian, &pdu.addrTon); err != nil {
		return nil, err
	}
	if err = binary.Read(buf, binary.BigEndian, &pdu.addrNpi); err != nil {
		return nil, err
	}
	if pdu.addressRange, err = buf.ReadString(0); err != nil {
		return nil, err
	}
	pdu.addressRange = pdu.addressRange[:len(pdu.addressRange)-1]
	return pdu, nil
}

func (b *BindPDU) CalculateCommandLength() uint32 {
	return 16 + uint32(len(b.systemId)+1+len(b.password)+1+len(b.systemType)+1+1+1+1+len(b.addressRange)+1)
}

func (b *BindPDU) GetSystemId() string {
	return b.systemId
}

func (b *BindPDU) GetPassword() string {
	return b.password
}

func (b *BindPDU) GetSystemType() string {
	return b.systemType
}

func (b *BindPDU) GetInterfaceVersion() byte {
	return b.interfaceVersion
}

func (b *BindPDU) GetAddrTon() TON {
	return b.addrTon
}

func (b *BindPDU) SetAddrTon(addrTon TON) {
	b.addrTon = addrTon
}

func (b *BindPDU) GetAddrNpi() NPI {
	return b.addrNpi
}

func (b *BindPDU) SetAddrNpi(addrNpi NPI) {
	b.addrNpi = addrNpi
}

func (b *BindPDU) GetAddressRange() string {
	return b.addressRange
}

func (b *BindPDU) SetAddressRange(addressRange string) error {
	if len(addressRange) > 41 {
		return errors.New("addressRange max length 41")
	}
	b.addressRange = addressRange
	return nil
}

func (b *BindPDU) WriteToBuffer(buf *bufio.Writer) error {
	b.commandLength = b.CalculateCommandLength()

	if err := b.Header.WriteToBuffer(buf); err != nil {
		return err
	}
	n, _ := buf.WriteString(b.systemId)
	if n != len(b.systemId) {
		return errors.New("Error writing systemId to buffer")
	}
	if err := buf.WriteByte(0); err != nil {
		return err
	}
	n, _ = buf.WriteString(b.password)
	if n != len(b.password) {
		return errors.New("Error writing password to buffer")
	}
	if err := buf.WriteByte(0); err != nil {
		return err
	}
	n, _ = buf.WriteString(b.systemType)
	if n != len(b.systemType) {
		return errors.New("Error writing systemType to buffer")
	}
	if err := buf.WriteByte(0); err != nil {
		return err
	}
	if err := buf.WriteByte(b.interfaceVersion); err != nil {
		return err
	}
	if err := binary.Write(buf, binary.BigEndian, b.addrTon); err != nil {
		return err
	}
	if err := binary.Write(buf, binary.BigEndian, b.addrNpi); err != nil {
		return err
	}
	n, _ = buf.WriteString(b.addressRange)
	if n != len(b.addressRange) {
		return errors.New("Error writing addressRange to buffer")
	}
	if err := buf.WriteByte(0); err != nil {
		return err
	}
	return nil
}
