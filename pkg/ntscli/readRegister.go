package ntscli

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"go.bug.st/serial"
)

func readRegister(addr int64, data *int64) int {

	write_data := make([]byte, 0, 32)
	read_data := make([]byte, 32)
	temp_data := make([]byte, 0, 32)

	//log.Println("VERBOSE: Read Register")

	mode := &serial.Mode{
		BaudRate: BaudRate,
	}

	port, err := serial.Open(FileDescriptor, mode)
	if err != nil {
		port.Close()
		log.Fatal("serial open err: ", err)
	}

	port.SetReadTimeout(time.Millisecond)

	write_data = append(write_data, "$RC,"...)
	hexAddr := fmt.Sprintf("0x%08x", addr)
	write_data = append(write_data, hexAddr...)
	checksum := calculateChecksum(write_data)
	write_data = append(write_data, '*')
	write_data = append(write_data, checksum...)
	write_data = append(write_data, '\r')
	write_data = append(write_data, '\n')

	//fmt.Print("VERBOSE: Read Command: ", string(write_data))

	//fmt.Printf("% #x ", write_data)

	n, err := port.Write(write_data)

	if err != nil {
		log.Fatal("write error: ", err)
	}

	if n == 0 {
		log.Fatal("response: none")
	}

	n, err = port.Read(read_data)
	port.Close()

	if err != nil {
		log.Fatal("read error: ", err)
	}

	if n == 0 {
		log.Fatal("response: none")
	}
	read_data = read_data[:n] // chop off
	read_string := string(read_data)

	//log.Print("VERBOSE: Read Response: ", string(read_data))

	//fmt.Printf("read: ", read_data)

	checksum = calculateChecksum(read_data)
	temp_data = append(temp_data, '*')
	temp_data = append(temp_data, checksum...)
	temp_data = append(temp_data, '\r')
	temp_data = append(temp_data, '\n')

	temp_string := string(temp_data)
	//fmt.Printf("% #x ", temp_data)

	if strings.HasPrefix(read_string, "$ER") {
		log.Println(handleReadWriteErrors(read_string))
		return -1
	}

	// check response
	if !strings.HasPrefix(read_string, "$RR") {
		log.Println("No read response received")
		return -1
	}

	// check checksum
	if !strings.HasSuffix(read_string, temp_string) {
		log.Println("checksum wrong: ", temp_string)
		//log.Fatal("checksum ")
		return -1
	} else {
		//log.Println(read_data)
		data_string := string(read_data[17 : 17+8])
		//log.Println("data_string: ", data_string)
		result, err := strconv.ParseInt(data_string, 16, 64)
		if err == nil {
			*data = int64(result)
		}

		//log.Println("read reg data dude: ", *data)

	}

	return 0
}
