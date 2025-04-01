package ntscli

import (
	"fmt"
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
}

var deviceConfig = DeviceConfig{
	BlockSize:       16,
	TypeInstanceReg: 0x00000000,
	BaseAddrLReg:    0x00000004,
	BaseAddrHReg:    0x00000008,
	IrqMaskReg:      0x0000000C,
}

func readDeviceConfig(coreType int64, temp_core *CoreConfig) int {

	var temp_data int64 = 0x00000000

	for i := int64(0); i < 256; i++ {
		//temp_core := Core{}

		type_addr := (0x00000000 + ((i * deviceConfig.BlockSize) + deviceConfig.TypeInstanceReg))
		//log.Println("i: ", i)
		if readRegister(type_addr, &temp_data) == 0 {
			//log.Println(temp_data)
			if (i == 0) && ((((temp_data >> int64(16)) & 0x0000FFFF) != types.ConfSlaveCoreType) || (((temp_data >> int64(0)) & 0x0000FFFF) != 1)) {

				log.Println("ERROR: not a conf block at the address expected")
				break

			} else if temp_data == 0 {
				break

			} else {
				temp_core.CoreType = ((temp_data >> 16) & 0x0000FFFF)
				temp_core.InstanceNumber = ((temp_data >> 0) & 0x0000FFFF)
			}

		} else {
			log.Fatal("Error in reading modules config")
		}

		low_addr := (0x00000000 + ((i * deviceConfig.BlockSize) + deviceConfig.BaseAddrLReg))
		if readRegister(low_addr, &temp_data) == 0 {
			temp_core.BaseAddrLReg = temp_data
		} else {
			break
		}

		high_addr := (0x00000000 + ((i * deviceConfig.BlockSize) + deviceConfig.BaseAddrHReg))
		if readRegister(high_addr, &temp_data) == 0 {
			temp_core.BaseAddrHReg = temp_data
		} else {
			break
		}

		interrupt_mask := (0x00000000 + ((i * deviceConfig.BlockSize) + deviceConfig.IrqMaskReg))
		if readRegister(interrupt_mask, &temp_data) == 0 {
			temp_core.IrqMaskReg = temp_data
		} else {
			break
		}

		//fmt.Println(fmt.Sprintf("low 0x%08x", temp_core.BaseAddrLReg), fmt.Sprintf(" high 0x%08x", temp_core.BaseAddrHReg), " ", temp_core, " ", "Core type: ", get_name(temp_core.CoreType))

		if coreType == temp_core.CoreType {
			fmt.Println(fmt.Sprintf("low 0x%08x", temp_core.BaseAddrLReg), fmt.Sprintf(" high 0x%08x", temp_core.BaseAddrHReg), " ", temp_core, " ", "Core type: ", get_name(temp_core.CoreType))
			break
		}
		////coreConfig.Cores = append(coreConfig.Cores, *temp_core) ?? not sure theres a good reason for this?

		// /read_core_parameters(temp_core)
	}

	return 0
}
