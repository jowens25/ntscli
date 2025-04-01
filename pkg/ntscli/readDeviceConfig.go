package ntscli

import (
	"log"
)

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

//var deviceConfig = DeviceConfig{
//	BlockSize:       16,
//	TypeInstanceReg: 0x00000000,
//	BaseAddrLReg:    0x00000004,
//	BaseAddrHReg:    0x00000008,
//	IrqMaskReg:      0x0000000C,
//	Cores:           make([]CoreConfig, 64),
//}

func readDeviceConfig(tempDeviceConfig *DeviceConfig) int {

	tempDeviceConfig.BlockSize = 16
	tempDeviceConfig.TypeInstanceReg = 0x00000000
	tempDeviceConfig.BaseAddrLReg = 0x00000004
	tempDeviceConfig.BaseAddrHReg = 0x00000008
	tempDeviceConfig.IrqMaskReg = 0x0000000C

	var tempData int64 = 0x00000000
	var tempCore CoreConfig

	for i := int64(0); i < 256; i++ {
		//tempCore := Core{}

		type_addr := (0x00000000 + ((i * tempDeviceConfig.BlockSize) + tempDeviceConfig.TypeInstanceReg))
		//log.Println("i: ", i)
		if readRegister(type_addr, &tempData) == 0 {
			//log.Println(tempData)
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

		low_addr := (0x00000000 + ((i * tempDeviceConfig.BlockSize) + tempDeviceConfig.BaseAddrLReg))
		if readRegister(low_addr, &tempData) == 0 {
			tempCore.BaseAddrLReg = tempData
		} else {
			break
		}

		high_addr := (0x00000000 + ((i * tempDeviceConfig.BlockSize) + tempDeviceConfig.BaseAddrHReg))
		if readRegister(high_addr, &tempData) == 0 {
			tempCore.BaseAddrHReg = tempData
		} else {
			break
		}

		interrupt_mask := (0x00000000 + ((i * tempDeviceConfig.BlockSize) + tempDeviceConfig.IrqMaskReg))
		if readRegister(interrupt_mask, &tempData) == 0 {
			tempCore.IrqMaskReg = tempData
		} else {
			break
		}

		tempDeviceConfig.Cores = append(tempDeviceConfig.Cores, tempCore)
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
