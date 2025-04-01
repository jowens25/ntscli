package ntscli

import (
	"encoding/hex"
	"strings"
)

func calculateChecksum(writeData []byte) string {
	var checksum byte = 0
	for i := 1; i < len(writeData); i++ {
		if writeData[i] == byte('*') {
			break
		}
		checksum ^= writeData[i]
	}
	return strings.ToUpper(hex.EncodeToString([]byte{checksum}))
}
