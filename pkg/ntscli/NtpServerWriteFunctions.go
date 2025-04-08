package ntscli

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func writeNtpServerStatus(status string) {

	if status == "enable" {
		tempData = 0x00000001
	} else if status == "disable" {
		tempData = 0x00000000
	} else {
		log.Fatal("Please enter a valid status (enabled or disabled)")
	}

	if writeRegister(NtpCore.BaseAddrLReg+ntpServer.ControlReg, &tempData) == 0 {
		fmt.Println("VERBOSE NTP SERVER: ", readNtpServerEnable())
	} else {
		log.Fatal(" VERBOSE ERROR WRITING NTP")
	}

}

func writeNtpServerMac(macAddr string) {

	macAddr = strings.Replace(macAddr, ":", "", -1)

	mac, err := strconv.ParseInt(macAddr, 16, 64)
	if err != nil {
		log.Println("Please enter a valid mac address")
		log.Fatal(err)
	}
	// 4e:54:4c:ff:00:00
	//fmt.Println(mac)

	tempData = 0x00000000
	tempData |= (mac >> 16) & 0x000000FF
	tempData = tempData << 8
	tempData |= (mac >> 24) & 0x000000FF
	tempData = tempData << 8
	tempData |= (mac >> 32) & 0x000000FF
	tempData = tempData << 8
	tempData |= (mac >> 40) & 0x000000FF

	if writeRegister(NtpCore.BaseAddrLReg+ntpServer.ConfigMac1Reg, &tempData) == 0 { // first bit writes success
		//fmt.Println("after first write")
		tempData = 0x00000000
		tempData |= (mac >> 0) & 0x000000FF
		tempData = tempData << 8
		tempData |= (mac >> 8) & 0x000000FF

		if writeRegister(NtpCore.BaseAddrLReg+ntpServer.ConfigMac2Reg, &tempData) == 0 { // second bit write success
			//fmt.Println("after second write")
			tempData = 0x00000004                                                               // write
			if writeRegister(NtpCore.BaseAddrLReg+ntpServer.ConfigControlReg, &tempData) == 0 { //update control to show change
				// ui->NtpServerMacValue->setText(temp_string);

				//fmt.Println(readNtpServerMac())

			} else {
				// ui->NtpServerMacValue->setText("NA");
			}

		}

	}

	//return macAddr
}

func writeNtpServerVlanEnable(mode string) {

	var currentSettings int64
	readRegister(NtpCore.BaseAddrLReg+ntpServer.ConfigVlanReg, &currentSettings)
	currentSettings &= 0x0000FFFF

	if mode == "enable" {
		tempData = 0x00010000 | currentSettings
	} else if mode == "disable" {
		tempData = 0x00000000 | currentSettings
	} else {
		tempData = 0x00000000 | currentSettings
	}

	if writeRegister(NtpCore.BaseAddrLReg+ntpServer.ConfigVlanReg, &tempData) == 0 {

		tempData = 0x00000002
		if writeRegister(NtpCore.BaseAddrLReg+ntpServer.ConfigControlReg, &tempData) == 0 {
			// do nothing
		} else {
			fmt.Println("vlan enabled: false")

		}

	} else {
		fmt.Println("vlan enabled: false")

	}

}

func writeNtpServerVlanValue(value string) {
	value, _ = strings.CutPrefix(value, "0x")
	readRegister(NtpCore.BaseAddrLReg+ntpServer.ConfigVlanReg, &tempData)
	v, err := strconv.ParseInt(value, 16, 64)
	if err != nil {
		log.Fatal("error")
	}
	//fmt.Println("v: ", v)
	tempData &= 0xFFFF0000 // keep the current value in the upper part of the register
	tempData |= v

	fmt.Println(tempData)
	if writeRegister(NtpCore.BaseAddrLReg+ntpServer.ConfigVlanReg, &tempData) == 0 {

		tempData = 0x00000002
		if writeRegister(NtpCore.BaseAddrLReg+ntpServer.ConfigControlReg, &tempData) == 0 {
			// do nothing
		} else {
			fmt.Println("vlan enabled: false")

		}

	} else {
		fmt.Println("vlan enabled: false")

	}

}

func writeNtpServerIpMode(mode string) {
	tempData = 0x00000000
	readRegister(NtpCore.BaseAddrLReg+ntpServer.ConfigControlReg, &tempData)

	if mode == "ipv4" {
		tempData |= 0x00000001
	} else if mode == "ipv6" {
		tempData |= 0x01000002
	} else {
		tempData |= 0x00000000

	}

	writeRegister(NtpCore.BaseAddrLReg+ntpServer.ConfigModeReg, &tempData)
}

func writeNtpServerUnicastMode(en string) {
	//fmt.Println("NTP UNICAST: ", readNtpServerUnicastMode())
	tempData = 0x00000000
	readRegister(NtpCore.BaseAddrLReg+ntpServer.ConfigModeReg, &tempData)
	//fmt.Printf("current: 0x%08x\n", tempData) // base 16 string format
	tempData &= ^0x00000010 // Clear bit 16 (using NOT + AND)
	if en == "enabled" {
		tempData |= 0x00000010 // set unicast on
	} else if en == "disabled" {
		tempData |= 0x00000000
	} else {
		tempData |= 0x00000000
	}
	//fmt.Printf("to write : 0x%08x\n", tempData) // base 16 string format

	if writeRegister(NtpCore.BaseAddrLReg+ntpServer.ConfigModeReg, &tempData) == 0 {
		tempData = 0x00000001
		if writeRegister(NtpCore.BaseAddrLReg+ntpServer.ConfigControlReg, &tempData) == 0 {
			// nothing
		} else {
			fmt.Println("NTP SERVER UNICAST: DISABLED")
		}
	} else {
		fmt.Println("NTP SERVER UNICAST: DISABLED")
	}
}

func writeNtpServerMulticastMode(en string) {
	//fmt.Println("NTP UNICAST: ", readNtpServerUnicastMode())
	tempData = 0x00000000
	readRegister(NtpCore.BaseAddrLReg+ntpServer.ConfigModeReg, &tempData)
	//fmt.Printf("current: 0x%08x\n", tempData) // base 16 string format
	tempData &= ^0x00000020 // Clear bit 16 (using NOT + AND)
	if en == "enabled" {
		tempData |= 0x00000020 // set unicast on
	} else if en == "disabled" {
		tempData |= 0x00000000
	} else {
		tempData |= 0x00000000
	}
	//fmt.Printf("to write : 0x%08x\n", tempData) // base 16 string format

	if writeRegister(NtpCore.BaseAddrLReg+ntpServer.ConfigModeReg, &tempData) == 0 {
		tempData = 0x00000001
		if writeRegister(NtpCore.BaseAddrLReg+ntpServer.ConfigControlReg, &tempData) == 0 {
			// nothing
		} else {
			fmt.Println("NTP SERVER UNICAST: DISABLED")
		}
	} else {
		fmt.Println("NTP SERVER UNICAST: DISABLED")
	}
}

func writeNtpServerBroadcastMode(en string) {
	//fmt.Println("NTP UNICAST: ", readNtpServerUnicastMode())
	tempData = 0x00000000
	readRegister(NtpCore.BaseAddrLReg+ntpServer.ConfigModeReg, &tempData)
	//fmt.Printf("current: 0x%08x\n", tempData) // base 16 string format
	tempData &= ^0x00000040 // Clear bit 16 (using NOT + AND)
	if en == "enabled" {
		tempData |= 0x00000040 // set unicast on
	} else if en == "disabled" {
		tempData |= 0x00000000
	} else {
		tempData |= 0x00000000
	}
	//fmt.Printf("to write : 0x%08x\n", tempData) // base 16 string format

	if writeRegister(NtpCore.BaseAddrLReg+ntpServer.ConfigModeReg, &tempData) == 0 {
		tempData = 0x00000001
		if writeRegister(NtpCore.BaseAddrLReg+ntpServer.ConfigControlReg, &tempData) == 0 {
			// nothing
		} else {
			fmt.Println("NTP SERVER BROADCAST: DISABLED")
		}
	} else {
		fmt.Println("NTP SERVER BROADCAST: DISABLED")
	}
}
