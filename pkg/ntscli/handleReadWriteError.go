package ntscli

func handleReadWriteError(errorCode int64) string {

	switch errorCode {

	case 0x00000000:
		return "Check sum error"
	case 0x00000001:
		return "Uknown command or error in command"
	case 0x00000002:
		return "Read Error on AXI"
	case 0x00000004:
		return "Write Error on AXI"
	default:
		return "Uknown error code"
	}

}
