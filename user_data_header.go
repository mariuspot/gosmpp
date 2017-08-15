package smpp

import (
	"errors"
)

//Based on
//  GSM 03.40
//  July 1996
//  Version 5.3.0
var (
	E_INFORMATION_ELEMENT_NOT_FOUND  error = errors.New("InformationElement Not Found")
	E_INFORMATION_ELEMENT_WRONG_TYPE error = errors.New("InformationElement Wrong Type")
)

type InformationElementIdentifier uint8

const (
	IEI_Concatenated_Short_Message InformationElementIdentifier = 0x00 //00	Concatenated short messages
	IEI_Special_SMS_Message                                     = 0x01 //01	Special SMS Message Indication
	//02	Reserved
	//03	Value not used to avoid misinterpretation as <LF> character
	//04 - 7F	Reserved for future use
	//80 - 9F	SME to SME specific use
	//A0 - BF	Reserved for future use
	//C0 - DF	SC specific use
	//E0 - FF	Reserved for future use
)

func (i InformationElementIdentifier) String() string {
	switch i {
	case IEI_Concatenated_Short_Message:
		return "Concatenated Short Message"
	case IEI_Special_SMS_Message:
		return "Special SMS Message"
	default:
		return "Unkown Information Element Identifier"
	}
}

type SpecialMessageIndicationType uint8

const (
	SMIT_Voice_Message_Waiting           SpecialMessageIndicationType = 0x00 //000 0000 Voice Message Waiting
	SMIT_Fax_Message_Waiting                                          = 0x01 //000 0001 Fax Message Waiting
	SMIT_Electronic_Mail_Message_Waiting                              = 0x02 //000 0010 Electronic Mail Message Waiting
	SMIT_Other_Message_Waiting                                        = 0x03 //000 0011 Other Message Waiting (see TS GSM 03.38 for definition of "other")
)

func (s SpecialMessageIndicationType) String() string {
	switch s {
	case SMIT_Voice_Message_Waiting:
		return "Voice Message Waiting"
	case SMIT_Fax_Message_Waiting:
		return "Fax Message Waiting"
	case SMIT_Electronic_Mail_Message_Waiting:
		return "Electronic Mail Message Waiting"
	case SMIT_Other_Message_Waiting:
		return "Other Message Waiting"
	default:
		return "Invalid Special Message Indication Type"
	}
}

type InformationElement struct {
	identifier InformationElementIdentifier
	data       []byte
}

func (i *InformationElement) GetIdentifier() InformationElementIdentifier {
	return i.identifier
}

func (i *InformationElement) GetData() []byte {
	return i.data
}

func (i *InformationElement) GetConcatenatedShortMessageInfo() (reference byte, segments byte, sequence byte, err error) {
	if i.identifier == IEI_Concatenated_Short_Message && len(i.data) == 3 {
		return i.data[0], i.data[1], i.data[2], nil
	}
	return 0, 0, 0, E_INFORMATION_ELEMENT_WRONG_TYPE
}

func (i *InformationElement) GetSpecialSMSMessageIndication() (store bool, indicationType SpecialMessageIndicationType, count byte, err error) {
	if i.identifier == IEI_Special_SMS_Message && len(i.data) == 2 {
		if i.data[0]&0x80 == 0x80 {
			store = true
			indicationType = SpecialMessageIndicationType(i.data[0] - 0x80)
		} else {
			indicationType = SpecialMessageIndicationType(i.data[0])
		}
		count = i.data[1]
		return
	}
	return false, 0, 0, E_INFORMATION_ELEMENT_WRONG_TYPE
}

func (i InformationElement) GetDataString() string {
	return string(i.data)
}

func (i InformationElement) GetDataByte() byte {
	if len(i.data) > 0 {
		return i.data[0]
	}
	return 0
}

type UserDataHeader struct {
	informationElements []*InformationElement
}

func NewUserDataHeader() *UserDataHeader {
	return &UserDataHeader{make([]*InformationElement, 0, 0)}
}

func (u *UserDataHeader) Decode(udh []byte) {
	for i := byte(1); i < udh[0]; i += 2 + udh[i+1] {
		u.informationElements = append(u.informationElements, &InformationElement{InformationElementIdentifier(udh[i]), udh[i+2 : i+2+udh[i+1]]})
	}
}

func (u *UserDataHeader) GetInformationElements() []*InformationElement {
	return u.informationElements
}

func (u *UserDataHeader) GetInformationElement(identifier InformationElementIdentifier) (*InformationElement, error) {
	for _, informationElement := range u.informationElements {
		if informationElement.GetIdentifier() == identifier {
			return informationElement, nil
		}
	}
	return nil, E_INFORMATION_ELEMENT_NOT_FOUND
}
