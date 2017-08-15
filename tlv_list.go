package smpp

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
)

type TLVList struct {
	tlvs []*TLV
}

func (t TLVList) String() string {
	var s string = "["
	var space string = ""
	for _, tlv := range t.tlvs {
		s += fmt.Sprintf("%s%s", space, tlv)
		if space == "" {
			space = ", "
		}
	}
	s += "]"
	return s
}

func NewTLVList() *TLVList {
	return &TLVList{make([]*TLV, 0)}
}

func NewTLVListFromBuffer(buf *bytes.Buffer) (*TLVList, error) {
	tlvList := NewTLVList()
	for buf.Len() >= 4 {
		tlv, err := NewTLVFromBuffer(buf)
		if err != nil {
			return nil, err
		}
		tlvList.AddTLV(tlv)
	}
	return tlvList, nil
}

func (t *TLVList) AddTLV(tlv *TLV) {
	t.tlvs = append(t.tlvs, tlv)
}

func (t *TLVList) RemoveTLV(tag Tag) {
	for i := 0; i < len(t.tlvs); i++ {
		if t.tlvs[i].GetTag() == tag {
			t.tlvs = append(t.tlvs[0:i], t.tlvs[i+1:len(t.tlvs)]...)
		}
	}
}

func (t *TLVList) GetTLV(tag Tag) (*TLV, error) {
	for _, c := range t.tlvs {
		if c.GetTag() == tag {
			return c, nil
		}
	}
	return nil, errors.New("TLV NOT FOUND")
}

func (t *TLVList) GetTLVList() *[]*TLV {

	return &t.tlvs
}

func (t *TLVList) GetLength() uint32 {
	var length uint32 = 0
	for _, tlv := range t.tlvs {
		length += uint32(4 + tlv.lenght)
	}
	return length
}

func (t *TLVList) WriteToBuffer(buf *bufio.Writer) error {
	for _, tlv := range t.tlvs {
		if err := tlv.WriteToBuffer(buf); err != nil {
			return err
		}
	}
	return nil
}
