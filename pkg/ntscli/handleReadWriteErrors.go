package ntscli

import (
	"log"
	"strconv"
	"strings"
)

func handleReadWriteErrors(errorMessage string) string {

	tempString := strings.Trim(errorMessage, "$ER,")
	codeParts := strings.Split(tempString, "*")
	errorCode, err := strconv.ParseInt(codeParts[0], 0, 64)
	if err != nil {
		log.Println("Its a good day when you get an error parsing error... ", err)
	}
	//checkSum := codeParts[1]

	switch errorCode {

	case 0x00000000:
		return "Checksum error"
	case 0x00000001:
		return "Unknown command (or error in command)"
	case 0x00000002:
		return "Read error on AX"
	case 0x00000003:
		return "Write error on AXI"
	case 0x00000004:
		return "Access timeout error on AXI (illegal address, no answer)"
	default:
		return "Uknown error code"
	}

}
