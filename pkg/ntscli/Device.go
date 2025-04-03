package ntscli

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

var FileDescriptor string = "/dev/ttyUSB0"
var BaudRates = []int{2000000, 1000000, 500000, 460800, 115200}
var BaudRate = 1000000

type Device struct {
	Name            string `json:"Name"`
	BlockSize       int64  `json:"BlockSize"`
	TypeInstanceReg int64  `json:"TypeInstanceReg"`
	BaseAddrLReg    int64  `json:"BaseAddrLReg"`
	BaseAddrHReg    int64  `json:"BaseAddrHReg"`
	IrqMaskReg      int64  `json:"IrqMaskReg"`
	Cores           []Core `json:"Cores"`
}

type Core struct {
	Name           string `json:"Name"`
	CoreType       int64  `json:"CoreType"`
	InstanceNumber int64  `json:"InstanceNumber"`
	BaseAddrLReg   int64  `json:"BaseAddrLReg"`
	BaseAddrHReg   int64  `json:"BaseAddrHReg"`
	IrqMaskReg     int64  `json:"IrqMaskReg"`
}

//func Device() {
//	log.Println("Device base command")
//}

/*

func DeviceConnect(input string) int {

	log.Println("VERBOSE: Serial Connect")

	mode := &serial.Mode{
		BaudRate: BaudRate,
	}

	port, err := serial.Open(FileDescriptor, mode)
	port.SetReadTimeout(time.Millisecond)

	if err != nil {
		port.Close()

		log.Fatal("serial open err: ", err)
	}

	write_data := make([]byte, 0, 32)
	write_data = append(write_data, "$CC"...)
	checksum := calculateChecksum(write_data)
	write_data = append(write_data, '*')
	write_data = append(write_data, checksum...)
	write_data = append(write_data, '\r')
	write_data = append(write_data, '\n')

	log.Println("writing: ", string(write_data))

	n, err := port.Write(write_data)
	if err != nil {
		log.Fatal("write error: ", err)
	}

	read_data := make([]byte, 32)
	n, err = port.Read(read_data)
	if err != nil {
		log.Fatal("read error: ", err)
	}
	if n == 0 {
		log.Fatal("response: none")
	}

	port.Close()
	read_string := string(read_data)

	log.Printf("received: %v", read_string)

	// check response
	if !strings.HasPrefix(read_string, "$CR") {
		log.Fatal("response: incorrect")
	}

	return 0
}

*/

/*
	func DeviceList() {
		readDeviceConfig()
		readDeviceConfigFile()
	}
*/
func DevicePullConfig() {
	readDeviceConfig()
}
func writeDeviceConfigFile(device Device) int {

	jsonData, err := json.MarshalIndent(device, "", " ")
	if err != nil {
		fmt.Println("error marshaling json: ", err)
		return -1
	}

	err = os.WriteFile("deviceConfig.json", jsonData, 0644)
	if err != nil {
		fmt.Println("error writing to file: ", err)
		return -1
	}

	fmt.Println("Struct written to deviceConfig.json successfully!")
	return 0
}

/*
func readDeviceConfigFile() int { // allow passing file name?

	fileBytes, err := os.ReadFile("deviceConfig.json")

	if err != nil {
		fmt.Println("error reading config file: ", err)
		return -1
	}
	device := Device{}
	err = json.Unmarshal(fileBytes, &device)
	if err != nil {
		fmt.Println("error parsing config file: ", err)
		return -1
	}

	fmt.Println("loaded config file")
	fmt.Println(device)
	return 0

}
*/
func readDeviceConfig() int {
	var device Device

	device.Name = "Novus Time Server"

	device.BlockSize = 16
	device.TypeInstanceReg = 0x00000000
	device.BaseAddrLReg = 0x00000004
	device.BaseAddrHReg = 0x00000008
	device.IrqMaskReg = 0x0000000C

	var tempData int64 = 0x00000000
	var tempCore Core

	for i := int64(0); i < 256; i++ {

		type_addr := (0x00000000 + ((i * device.BlockSize) + device.TypeInstanceReg))
		if readRegister(type_addr, &tempData) == 0 {
			if (i == 0) && ((((tempData >> int64(16)) & 0x0000FFFF) != types.ConfSlaveCoreType) || (((tempData >> int64(0)) & 0x0000FFFF) != 1)) {

				log.Println("ERROR: not a conf block at the address expected")
				break

			} else if tempData == 0 {
				break

			} else {

				tempCore.CoreType = ((tempData >> 16) & 0x0000FFFF)
				tempCore.Name = get_name(tempCore.CoreType)
				tempCore.InstanceNumber = ((tempData >> 0) & 0x0000FFFF)
			}

		} else {
			log.Fatal("Error in reading modules config")
		}

		low_addr := (0x00000000 + ((i * device.BlockSize) + device.BaseAddrLReg))
		if readRegister(low_addr, &tempData) == 0 {
			tempCore.BaseAddrLReg = tempData
		} else {
			break
		}

		high_addr := (0x00000000 + ((i * device.BlockSize) + device.BaseAddrHReg))
		if readRegister(high_addr, &tempData) == 0 {
			tempCore.BaseAddrHReg = tempData
		} else {
			break
		}

		interrupt_mask := (0x00000000 + ((i * device.BlockSize) + device.IrqMaskReg))
		if readRegister(interrupt_mask, &tempData) == 0 {
			tempCore.IrqMaskReg = tempData
		} else {
			break
		}

		device.Cores = append(device.Cores, tempCore)
		//fmt.Println(fmt.Sprintf("low 0x%08x", tempCore.BaseAddrLReg), fmt.Sprintf(" high 0x%08x", tempCore.BaseAddrHReg), " ", tempCore, " ", "Core type: ", get_name(tempCore.CoreType))

		//if coreType == tempCore.CoreType {
		//	fmt.Println(fmt.Sprintf("low 0x%08x", tempCore.BaseAddrLReg), fmt.Sprintf(" high 0x%08x", tempCore.BaseAddrHReg), " ", tempCore, " ", "Core type: ", get_name(tempCore.CoreType))
		//	break
		//} else if coreType == 0 {
		//	fmt.Println(fmt.Sprintf("low 0x%08x", tempCore.BaseAddrLReg), fmt.Sprintf(" high 0x%08x", tempCore.BaseAddrHReg), " ", tempCore, " ", "Core type: ", get_name(tempCore.CoreType))
		//}
		////coreConfig.Cores = append(coreConfig.Cores, *tempCore) ?? not sure theres a good reason for this?

		// /read_core_parameters(tempCore)
	}

	writeDeviceConfigFile(device)
	return 0
}
