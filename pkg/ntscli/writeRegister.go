package ntscli

import (
	"fmt"
	"log"
	"strings"
	"time"

	"go.bug.st/serial"
)

func writeRegister(addr int64, data *int64) int {
	write_data := make([]byte, 0, 32)
	read_data := make([]byte, 32)
	temp_data := make([]byte, 0, 32)

	//log.Println("VERBOSE: Write Register")

	mode := &serial.Mode{
		BaudRate: BaudRate,
	}

	port, err := serial.Open(FileDescriptor, mode)
	if err != nil {
		port.Close()
		log.Fatal("serial open err: ", err)
	}

	port.SetReadTimeout(time.Millisecond)

	write_data = append(write_data, "$WC,"...)
	hexAddr := fmt.Sprintf("0x%08x", addr)
	write_data = append(write_data, hexAddr...)
	write_data = append(write_data, ',')

	hexData := fmt.Sprintf("0x%08x", *data)
	write_data = append(write_data, hexData...)

	checksum := calculateChecksum(write_data)
	write_data = append(write_data, '*')
	write_data = append(write_data, checksum...)
	write_data = append(write_data, '\r')
	write_data = append(write_data, '\n')

	//log.Print("VERBOSE write: ", string(write_data))
	//fmt.Printf("write: % #x \n", write_data)

	n, err := port.Write(write_data)

	if err != nil {
		log.Fatal("write error: ", err)
	}

	if n == 0 {
		log.Fatal("response: none")
	}

	n, err = port.Read(read_data)
	read_data = read_data[:n]
	port.Close()
	read_string := string(read_data)
	//log.Print("verbose read: ", string(read_data))
	//fmt.Printf("read: % #x \n", read_data)

	checksum = calculateChecksum(read_data)
	temp_data = append(temp_data, '*')
	temp_data = append(temp_data, checksum...)
	temp_data = append(temp_data, '\r')
	temp_data = append(temp_data, '\n')

	//fmt.Printf("% #x \n", temp_data)

	temp_string := string(temp_data)
	if err != nil {
		log.Fatal("read error: ", err)
	}

	if n == 0 {
		log.Fatal("response: none")
	}

	// check response
	if !strings.HasPrefix(read_string, "$WR") {
		log.Fatal("w No correct response received")
	}

	// check checksum
	if !strings.HasSuffix(read_string, temp_string) {
		log.Println("checksum wrong")
	}

	return 0
}
