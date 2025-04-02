package ntscli

import (
	"log"
	"strings"
	"time"

	"go.bug.st/serial"
)

var FileDescriptor string = "/dev/ttyUSB0"
var BaudRates = []int{2000000, 1000000, 500000, 460800, 115200}
var BaudRate = 1000000

type CoreConfig struct {
	CoreType       int64
	InstanceNumber int64
	BaseAddrLReg   int64
	BaseAddrHReg   int64
	IrqMaskReg     int64
}

type DeviceConfig struct {
	BlockSize       int64
	TypeInstanceReg int64
	BaseAddrLReg    int64
	BaseAddrHReg    int64
	IrqMaskReg      int64
	Cores           []CoreConfig
}

var deviceConfig = DeviceConfig{}

func Device() {
	log.Println("Device base command")
}

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

func DeviceList() {
	log.Println("List out the core config")
	readDeviceConfig()
	log.Println(deviceConfig)
}

func readDeviceConfig() int {

	deviceConfig.BlockSize = 16
	deviceConfig.TypeInstanceReg = 0x00000000
	deviceConfig.BaseAddrLReg = 0x00000004
	deviceConfig.BaseAddrHReg = 0x00000008
	deviceConfig.IrqMaskReg = 0x0000000C

	var tempData int64 = 0x00000000
	var tempCore CoreConfig

	for i := int64(0); i < 256; i++ {

		type_addr := (0x00000000 + ((i * deviceConfig.BlockSize) + deviceConfig.TypeInstanceReg))
		if readRegister(type_addr, &tempData) == 0 {
			if (i == 0) && ((((tempData >> int64(16)) & 0x0000FFFF) != types.ConfSlaveCoreType) || (((tempData >> int64(0)) & 0x0000FFFF) != 1)) {

				log.Println("ERROR: not a conf block at the address expected")
				break

			} else if tempData == 0 {
				break

			} else {
				tempCore.CoreType = ((tempData >> 16) & 0x0000FFFF)
				tempCore.InstanceNumber = ((tempData >> 0) & 0x0000FFFF)
			}

		} else {
			log.Fatal("Error in reading modules config")
		}

		low_addr := (0x00000000 + ((i * deviceConfig.BlockSize) + deviceConfig.BaseAddrLReg))
		if readRegister(low_addr, &tempData) == 0 {
			tempCore.BaseAddrLReg = tempData
		} else {
			break
		}

		high_addr := (0x00000000 + ((i * deviceConfig.BlockSize) + deviceConfig.BaseAddrHReg))
		if readRegister(high_addr, &tempData) == 0 {
			tempCore.BaseAddrHReg = tempData
		} else {
			break
		}

		interrupt_mask := (0x00000000 + ((i * deviceConfig.BlockSize) + deviceConfig.IrqMaskReg))
		if readRegister(interrupt_mask, &tempData) == 0 {
			tempCore.IrqMaskReg = tempData
		} else {
			break
		}

		deviceConfig.Cores = append(deviceConfig.Cores, tempCore)
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

	return 0
}
