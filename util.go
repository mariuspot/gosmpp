package smpp

func IsResponsePDU(commandId CommandId) bool {
	if commandId >= CID_GenericNack {
		return true
	}
	return false
}
