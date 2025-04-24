package ntscli

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var FileDescriptor string = "/dev/ttyUSB0"
var BaudRates = []int{2000000, 1000000, 500000, 460800, 115200}
var BaudRate = 1000000
var tempData int64

type Device struct {
	Name            string          `json:"Name"`
	BlockSize       int64           `json:"BlockSize"`
	TypeInstanceReg int64           `json:"TypeInstanceReg"`
	BaseAddrLReg    int64           `json:"BaseAddrLReg"`
	BaseAddrHReg    int64           `json:"BaseAddrHReg"`
	IrqMaskReg      int64           `json:"IrqMaskReg"`
	Cores           map[string]Core `json:"Cores"`
}

var NovusDevice = Device{Cores: make(map[string]Core)}

type Core struct {
	Name           string      `json:"Name"`
	CoreType       int64       `json:"CoreType"`
	InstanceNumber int64       `json:"InstanceNumber"`
	BaseAddrLReg   int64       `json:"BaseAddrLReg"`
	BaseAddrHReg   int64       `json:"BaseAddrHReg"`
	IrqMaskReg     int64       `json:"IrqMaskReg"`
	Registers      RegisterSet `json:"Registers"`
}

var NtpServerCore Core

func DeviceHasNtpServer() int64 {
	for _, core := range NovusDevice.Cores {
		if core.CoreType == types.NtpServerCoreType {
			NtpServerCore = NovusDevice.Cores["NtpServerCoreType"]
			return 0
		}
	}
	return -1
}

var PtpOcCore Core

func DeviceHasPtpOc() int64 {
	for _, core := range NovusDevice.Cores {
		if core.CoreType == types.NtpServerCoreType {
			PtpOcCore = NovusDevice.Cores["PtpOrdinaryClockCoreType"]
			return 0
		}
	}
	return -1
}

func ReadDeviceConfig() int {

	NovusDevice.Name = "Novus Time Server"

	NovusDevice.BlockSize = 16
	NovusDevice.TypeInstanceReg = 0x00000000
	NovusDevice.BaseAddrLReg = 0x00000004
	NovusDevice.BaseAddrHReg = 0x00000008
	NovusDevice.IrqMaskReg = 0x0000000C

	var tempCore Core

	var tempData int64 = 0x00000000

	for i := int64(0); i < 256; i++ {

		typeAddr := (0x00000000 + ((i * NovusDevice.BlockSize) + NovusDevice.TypeInstanceReg))
		if readRegister(typeAddr, &tempData) == 0 {
			if (i == 0) && ((((tempData >> int64(16)) & 0x0000FFFF) != 1) || (((tempData >> int64(0)) & 0x0000FFFF) != 1)) {

				log.Println("ERROR: not a conf block at the address expected")
				break

			} else if tempData == 0 {
				break

			} else {

				tempCore.CoreType = ((tempData >> 16) & 0x0000FFFF)
				tempCore.Name = getName(tempCore.CoreType)
				tempCore.InstanceNumber = ((tempData >> 0) & 0x0000FFFF)
				tempCore.Registers = getRegistersByType(tempCore.CoreType)
			}

		} else {
			log.Fatal("Error in reading modules config")
		}

		lowAddr := (0x00000000 + ((i * NovusDevice.BlockSize) + NovusDevice.BaseAddrLReg))
		if readRegister(lowAddr, &tempData) == 0 {
			tempCore.BaseAddrLReg = tempData
		} else {
			break
		}

		highAddr := (0x00000000 + ((i * NovusDevice.BlockSize) + NovusDevice.BaseAddrHReg))
		if readRegister(highAddr, &tempData) == 0 {
			tempCore.BaseAddrHReg = tempData
		} else {
			break
		}

		interruptMask := (0x00000000 + ((i * NovusDevice.BlockSize) + NovusDevice.IrqMaskReg))
		if readRegister(interruptMask, &tempData) == 0 {
			tempCore.IrqMaskReg = tempData
		} else {
			break
		}

		NovusDevice.Cores[tempCore.Name] = tempCore

		//device.Cores = append(device.Cores, tempCore)
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
	//fmt.Println(device)
	//writeDeviceConfigFile(device)

	//fmt.Println(NovusDevice)
	return 0
}

func loadDeviceConfig(fileName string) {

	data, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal("file err: ", err)
	}

	for _, line := range strings.Split(string(data), "\n") {
		//fmt.Println(i, line)

		if strings.Contains(line, "--") {
			// a comment
			continue

		} else if strings.Contains(line, "$WC") {

			line = strings.Trim(line, "\r\n")
			lineParts := strings.Split(line, ",")

			//fmt.Println(lineParts[1], lineParts[2])
			addr, err := strconv.ParseInt(lineParts[1], 0, 64)
			if err != nil {
				log.Fatal(err)
			}
			data, err := strconv.ParseInt(lineParts[2], 0, 64)
			if err != nil {
				log.Fatal(err)
			}
			writeRegister(addr, &data)
		}

	}

}

func dumpDeviceConfig(fileName string) {

	// read the device config - get the cores available
	// for the cores available: read all the known registers
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		log.Fatal(err)
	}

	for coreName, core := range NovusDevice.Cores {

		file.WriteString("--" + coreName + "\n")

		for regName, reg := range core.Registers {
			file.WriteString("----" + regName + "\n")
			tempData = 0x00000000
			//readRegister(core.BaseAddrLReg+reg, &tempData)

			file.WriteString(addNtpPropertyComment(regName))

			file.WriteString("----" + fmt.Sprintf("0x%08x", core.BaseAddrLReg+reg) + "," + fmt.Sprintf("0x%08x", tempData) + "\n")
		}
	}

	//addr := NovusDevice.Cores["NtpServerCoreType"].Registers["ControlReg"] + NovusDevice.Cores["NtpServerCoreType"].BaseAddrLReg
	//readRegister(addr, &tempData)
	//
	//hexData := fmt.Sprintf("0x%08x", tempData)
	//hexData = hexToUpper(hexData)
	//fmt.Println("hexdata ", hexData)
	//
	//hexAddr := fmt.Sprintf("0x%08x", addr)
	//hexAddr = hexToUpper(hexAddr)
	//fmt.Println("hexAddr ", hexAddr)

	//for k, addr := range ntpServer {
	//	fmt.Println(k, addr)
	//	fmt.Println(k, NovusDevice.Cores["NtpServerCoreType"].Registers[k])
	//}

	//for k, address := range NovusDevice.Cores["NtpServerCoreType"].Registers {
	//	addr := NovusDevice.Cores["NtpServerCoreType"].BaseAddrLReg + address
	//
	//	tempData = 0x00000000
	//	readRegister(address, &tempData)
	//
	//	hexAddr := fmt.Sprintf("0x%08x", addr)
	//	hexAddr = hexToUpper(hexAddr)
	//	fmt.Print("hexaddr ", hexAddr, "  ")
	//
	//	hexData := fmt.Sprintf("0x%08x", tempData)
	//	hexData = hexToUpper(hexData)
	//	fmt.Println("hexdata ", hexData)
	//
	//	file.WriteString("-- " + k + "\n")
	//	file.WriteString("$WC," + hexAddr + "," + hexData + "\r\n")
	//}

	for _, reg := range ntpServer {
		addrr := NovusDevice.Cores["NtpServerCoreType"].BaseAddrLReg + reg
		readRegister(addrr, &tempData)

		addr := fmt.Sprintf("0x%08x", addrr)
		data := fmt.Sprintf("0x%08x", tempData)

		file.WriteString("$WC," + addr + "," + data + "\r\n")
	}

	//for k, reg := range ntpServer {
	//	tempData = 0x00000000
	//
	//	fmt.Println(k)
	//	addrr := NovusDevice.Cores["NtpServerCoreType"].BaseAddrLReg + reg
	//	readRegister(addrr, &tempData)
	//
	//	addr := fmt.Sprintf("0x%08x", addrr)
	//	data := fmt.Sprintf("0x%08x", tempData)
	//
	//	file.WriteString("$WC," + addr + "," + data + "\r\n")
	//}

	file.Close()
}

func UpdateDevice(cmd *cobra.Command) {
	cmd.Flags().Visit(func(f *pflag.Flag) {

		switch f.Name {
		case "load":
			configFile, err := cmd.Flags().GetString(f.Name)
			if err != nil {
				log.Fatal("No such argument for property: ", f.Name, err)
			}
			loadDeviceConfig(configFile)

		case "dump":
			configFile, err := cmd.Flags().GetString(f.Name)
			if err != nil {
				log.Fatal("No such argument for property: ", f.Name, err)
			}
			dumpDeviceConfig(configFile)

		case "connect":

			connectDevice()

		default:
			fmt.Println("That does not appear to be a valid flag. Try: ", cmd.UsageString())
		}
	})

}
