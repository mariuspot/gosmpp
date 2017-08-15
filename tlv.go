package smpp

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"
)

type TLV struct {
	tag    Tag
	lenght uint16
	value  []byte
}

func (t TLV) String() string {

	return fmt.Sprintf("{tag: %s, length: %d, value: %v}", t.tag, t.lenght, t.value)
}

func NewTLV(tag Tag, value []byte) (*TLV, error) {
	return &TLV{tag, uint16(len(value)), value}, nil
}

func NewTLVString(tag Tag, value string) (*TLV, error) {
	return &TLV{tag, uint16(len(value)), []byte(value)}, nil
}

func NewTLVInt8(tag Tag, value int8) (*TLV, error) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.BigEndian, value); err != nil {
		return nil, err
	}
	return &TLV{tag, 1, buf.Bytes()}, nil
}

func NewTLVUInt8(tag Tag, value uint8) (*TLV, error) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.BigEndian, value); err != nil {
		return nil, err
	}
	return &TLV{tag, 1, buf.Bytes()}, nil
}

func NewTLVInt16(tag Tag, value int16) (*TLV, error) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.BigEndian, value); err != nil {
		return nil, err
	}
	return &TLV{tag, 2, buf.Bytes()}, nil
}

func NewTLVUInt16(tag Tag, value uint16) (*TLV, error) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.BigEndian, value); err != nil {
		return nil, err
	}
	return &TLV{tag, 2, buf.Bytes()}, nil
}

func NewTLVInt32(tag Tag, value int32) (*TLV, error) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.BigEndian, value); err != nil {
		return nil, err
	}
	return &TLV{tag, 4, buf.Bytes()}, nil
}

func NewTLVUInt32(tag Tag, value uint32) (*TLV, error) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.BigEndian, value); err != nil {
		return nil, err
	}
	return &TLV{tag, 4, buf.Bytes()}, nil
}

func NewTLVInt64(tag Tag, value int64) (*TLV, error) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.BigEndian, value); err != nil {
		return nil, err
	}
	return &TLV{tag, 8, buf.Bytes()}, nil
}

func NewTLVUInt64(tag Tag, value uint64) (*TLV, error) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.BigEndian, value); err != nil {
		return nil, err
	}
	return &TLV{tag, 8, buf.Bytes()}, nil
}

func NewTLVFromBuffer(buf *bytes.Buffer) (*TLV, error) {
	tlv := &TLV{}
	if err := binary.Read(buf, binary.BigEndian, &tlv.tag); err != nil {
		return nil, err
	}
	if err := binary.Read(buf, binary.BigEndian, &tlv.lenght); err != nil {
		return nil, err
	}
	var tempBuf bytes.Buffer
	for i := uint16(0); i < tlv.lenght; i++ {
		b, err := buf.ReadByte()
		if err != nil {
			return nil, err
		}
		if err := tempBuf.WriteByte(b); err != nil {
			return nil, err
		}
	}
	tlv.value = tempBuf.Bytes()
	return tlv, nil
}

func (t *TLV) GetTag() Tag {
	return t.tag
}

func (t *TLV) GetLength() uint16 {
	return t.lenght
}

func (t *TLV) GetValue() ([]byte, error) {
	return t.value, nil
}

func (t *TLV) GetValueString() (string, error) {

	return strings.Trim(string(t.value), string([]byte{0})), nil
}

func (t *TLV) GetValueInt8() (int8, error) {
	if t.lenght != 1 {
		return 0, errors.New("TLV length not equal to int8")
	}
	var result int8
	buf := bytes.NewReader(t.value)
	if err := binary.Read(buf, binary.BigEndian, &result); err != nil {
		return 0, err
	}
	return result, nil
}

func (t *TLV) GetValueUInt8() (uint8, error) {
	if t.lenght != 1 {
		return 0, errors.New("TLV length not equal to uint8")
	}
	var result uint8
	buf := bytes.NewReader(t.value)
	if err := binary.Read(buf, binary.BigEndian, &result); err != nil {
		return 0, err
	}
	return result, nil
}

func (t *TLV) GetValueInt16() (int16, error) {
	if t.lenght != 2 {
		return 0, errors.New("TLV length not equal to int16")
	}
	var result int16
	buf := bytes.NewReader(t.value)
	if err := binary.Read(buf, binary.BigEndian, &result); err != nil {
		return 0, err
	}
	return result, nil
}

func (t *TLV) GetValueUInt16() (uint16, error) {
	if t.lenght != 2 {
		return 0, errors.New("TLV length not equal to uint16")
	}
	var result uint16
	buf := bytes.NewReader(t.value)
	if err := binary.Read(buf, binary.BigEndian, &result); err != nil {
		return 0, err
	}
	return result, nil
}

func (t *TLV) GetValueInt32() (int32, error) {
	if t.lenght != 4 {
		return 0, errors.New("TLV length not equal to int32")
	}
	var result int32
	buf := bytes.NewReader(t.value)
	if err := binary.Read(buf, binary.BigEndian, &result); err != nil {
		return 0, err
	}
	return result, nil
}

func (t *TLV) GetValueUInt32() (uint32, error) {
	if t.lenght != 4 {
		return 0, errors.New("TLV length not equal to uint32")
	}
	var result uint32
	buf := bytes.NewReader(t.value)
	if err := binary.Read(buf, binary.BigEndian, &result); err != nil {
		return 0, err
	}
	return result, nil
}

func (t *TLV) GetValueInt64() (int64, error) {
	if t.lenght != 8 {
		return 0, errors.New("TLV length not equal to int64")
	}
	var result int64
	buf := bytes.NewReader(t.value)
	if err := binary.Read(buf, binary.BigEndian, &result); err != nil {
		return 0, err
	}
	return result, nil
}

func (t *TLV) GetValueUInt64() (uint64, error) {
	if t.lenght != 8 {
		return 0, errors.New("TLV length not equal to uint64")
	}
	var result uint64
	buf := bytes.NewReader(t.value)
	if err := binary.Read(buf, binary.BigEndian, &result); err != nil {
		return 0, err
	}
	return result, nil
}

func (t *TLV) WriteToBuffer(buf *bufio.Writer) error {

	if err := binary.Write(buf, binary.BigEndian, t.tag); err != nil {
		return err
	}
	if err := binary.Write(buf, binary.BigEndian, t.lenght); err != nil {
		return err
	}
	if _, err := buf.Write(t.value); err != nil {
		return err
	}
	return nil
}
