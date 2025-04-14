package ntscli

import (
	"fmt"
	"log"
	"net"
	"slices"
	"strconv"
	"strings"
)

// NtpServerSTATUS
func writeNtpServerStatus(status string) {

	if status == "enable" {
		tempData = 0x00000001
	} else if status == "disable" {
		tempData = 0x00000000
	} else {
		log.Fatal("Please enter a valid status (enabled or disabled)")
	}

	if writeRegister(NtpServerCore.BaseAddrLReg+ntpServer["ControlReg"], &tempData) == 0 {
		showNtpServerSTATUS()
	} else {
		log.Fatal(" VERBOSE ERROR WRITING NTP")
	}

}

// NtpServer INSTANCE
// NtpServer IP ADDRESS

func writeNtpServerIpAddr(addr string) {
	ipMode := readNtpServerIpMode()
	if ipMode == "IPv4" {
		ip4_addr_to_register_value(addr)
	} else if ipMode == "IPv6" {
		ip6_addr_to_register_value(addr)
	}
}

// NtpServer IP MODE
func writeNtpServerIpMode(mode string) {
	currentAddr := readNtpServerIpAddress()
	tempData = 0x00000000
	readRegister(NtpServerCore.BaseAddrLReg+ntpServer["ConfigModeReg"], &tempData)

	if mode == "ipv4" {
		tempData &= ^0x0000002 // Clear bit 16 (using NOT + AND)
		tempData |= 0x00000001

	} else if mode == "ipv6" {
		tempData &= ^0x00000001 // Clear bit 16 (using NOT + AND)
		tempData |= 0x01000002

	} else {
		log.Fatal("ip mode error")

	}

	if writeRegister(NtpServerCore.BaseAddrLReg+ntpServer["ConfigModeReg"], &tempData) == 0 {
		tempData = 0x00000001
		if writeRegister(NtpServerCore.BaseAddrLReg+ntpServer["ConfigControlReg"], &tempData) == 0 {
			writeNtpServerIpAddr(currentAddr)
			showNtpServerIPMODE()
			//showNtpServerIPADDRESS()
		} else {
			fmt.Println("NTP SERVER IP MODE: UPDATE FAILED")
		}
	} else {
		fmt.Println("NTP SERVER IP MODE: UPDATE FAILED")
	}

}

// NtpServer MAC ADDRESS
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

	if writeRegister(NtpServerCore.BaseAddrLReg+ntpServer["ConfigMac1Reg"], &tempData) == 0 { // first bit writes success
		//fmt.Println("after first write")
		tempData = 0x00000000
		tempData |= (mac >> 0) & 0x000000FF
		tempData = tempData << 8
		tempData |= (mac >> 8) & 0x000000FF

		if writeRegister(NtpServerCore.BaseAddrLReg+ntpServer["ConfigMac2Reg"], &tempData) == 0 { // second bit write success
			//fmt.Println("after second write")
			tempData = 0x00000004                                                                        // write
			if writeRegister(NtpServerCore.BaseAddrLReg+ntpServer["ConfigControlReg"], &tempData) == 0 { //update control to show change
				showNtpServerMACADDRESS()
			} else {
				fmt.Println("NTP SERVER MAC ADDRESS UPDATE FAILED")
			}
		} else {
			fmt.Println("NTP SERVER MAC ADDRESS UPDATE FAILED")
		}
	} else {
		fmt.Println("NTP SERVER MAC ADDRESS UPDATE FAILED")
	}
}

// NtpServer VLAN ENABLED
func writeNtpServerVlanEnable(mode string) {

	var currentSettings int64
	readRegister(NtpServerCore.BaseAddrLReg+ntpServer["ConfigVlanReg"], &currentSettings)
	currentSettings &= 0x0000FFFF

	if mode == "enable" {
		tempData = 0x00010000 | currentSettings
	} else if mode == "disable" {
		tempData = 0x00000000 | currentSettings
	} else {
		tempData = 0x00000000 | currentSettings
	}

	if writeRegister(NtpServerCore.BaseAddrLReg+ntpServer["ConfigVlanReg"], &tempData) == 0 {

		tempData = 0x00000002
		if writeRegister(NtpServerCore.BaseAddrLReg+ntpServer["ConfigControlReg"], &tempData) == 0 {
			showNtpServerVLANENABLED()
		} else {
			fmt.Println("vlan enabled: false")

		}

	} else {
		fmt.Println("vlan enabled: false")

	}

}

// NtpServer VLAN ENABLED
func writeVlanEnable(baseAddr int64, registers map[string]int64, mode string) {

	var currentSettings int64
	readRegister(baseAddr+registers["ConfigVlanReg"], &currentSettings)
	currentSettings &= 0x0000FFFF

	if mode == "enable" {
		tempData = 0x00010000 | currentSettings
	} else if mode == "disable" {
		tempData = 0x00000000 | currentSettings
	} else {
		tempData = 0x00000000 | currentSettings
	}

	if writeRegister(baseAddr+registers["ConfigVlanReg"], &tempData) == 0 {

		tempData = 0x00000002
		if writeRegister(baseAddr+registers["ConfigControlReg"], &tempData) == 0 {
			showNtpServerVLANENABLED()
		} else {
			fmt.Println("vlan enabled: false")

		}

	} else {
		fmt.Println("vlan enabled: false")

	}

}

// NtpServer VLAN VALUE
func writeNtpServerVlanValue(value string) {
	value, _ = strings.CutPrefix(value, "0x")
	readRegister(NtpServerCore.BaseAddrLReg+ntpServer["ConfigVlanReg"], &tempData)
	v, err := strconv.ParseInt(value, 16, 64)
	if err != nil {
		log.Fatal("error")
	}
	tempData &= 0xFFFF0000 // keep the current value in the upper part of the register
	tempData |= v

	if writeRegister(NtpServerCore.BaseAddrLReg+ntpServer["ConfigVlanReg"], &tempData) == 0 {
		tempData = 0x00000002
		if writeRegister(NtpServerCore.BaseAddrLReg+ntpServer["ConfigControlReg"], &tempData) == 0 {
			showNtpServerVLANVALUE()
		} else {
			fmt.Println("vlan enabled: false")
		}

	} else {
		fmt.Println("vlan enabled: false")
	}

}

// NtpServer UNICAST
func writeNtpServerUnicastMode(en string) {
	tempData = 0x00000000
	readRegister(NtpServerCore.BaseAddrLReg+ntpServer["ConfigModeReg"], &tempData)
	tempData &= ^0x00000010 // Clear bit 16 (using NOT + AND)
	if en == "enabled" {
		tempData |= 0x00000010 // set unicast on
	}

	writeRegister(NtpServerCore.BaseAddrLReg+ntpServer["ConfigModeReg"], &tempData)
	tempData = 0x00000001
	if writeRegister(NtpServerCore.BaseAddrLReg+ntpServer["ConfigControlReg"], &tempData) == 0 {
		showNtpServerUNICAST()
	}
}

// NtpServer MULTICAST
func writeNtpServerMulticastMode(en string) {
	tempData = 0x00000000
	readRegister(NtpServerCore.BaseAddrLReg+ntpServer["ConfigModeReg"], &tempData)
	tempData &= ^0x00000020 // Clear bit 16 (using NOT + AND)
	if en == "enabled" {
		tempData |= 0x00000020 // set multicast on
	}

	writeRegister(NtpServerCore.BaseAddrLReg+ntpServer["ConfigModeReg"], &tempData)
	tempData = 0x00000001
	if writeRegister(NtpServerCore.BaseAddrLReg+ntpServer["ConfigControlReg"], &tempData) == 0 {
		showNtpServerMULTICAST()
	}
}

// NtpServer BROADCAST
func writeNtpServerBroadcastMode(en string) {
	tempData = 0x00000000
	readRegister(NtpServerCore.BaseAddrLReg+ntpServer["ConfigModeReg"], &tempData)
	tempData &= ^0x00000040 // set broadcast off
	if en == "enabled" {
		tempData |= 0x00000040 // set broadcast on
	}

	writeRegister(NtpServerCore.BaseAddrLReg+ntpServer["ConfigModeReg"], &tempData)
	tempData = 0x00000001
	if writeRegister(NtpServerCore.BaseAddrLReg+ntpServer["ConfigControlReg"], &tempData) == 0 {
		showNtpServerBROADCAST()
	}

}

// NtpServer PRECISION
func writeNtpServerPrecision(value string) {
	tempData = 0x00000000
	readRegister(NtpServerCore.BaseAddrLReg+ntpServer["ConfigModeReg"], &tempData)

	val, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		log.Fatal("invalid precision value")
	}
	tempData &= ^((0xFFFFFFFF & 0x000000FF) << 8)
	tempData |= ((val & 0x000000FF) << 8)

	writeRegister(NtpServerCore.BaseAddrLReg+ntpServer["ConfigModeReg"], &tempData)
	tempData = 0x00000001
	if writeRegister(NtpServerCore.BaseAddrLReg+ntpServer["ConfigControlReg"], &tempData) == 0 {
		showNtpServerPRECISION()
	}
}

// NtpServer POLL INTERVAL
func writeNtpServerPollInternal(value string) {
	tempData = 0x00000000
	readRegister(NtpServerCore.BaseAddrLReg+ntpServer["ConfigModeReg"], &tempData)

	val, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		log.Fatal("invalid poll interval value")
	}
	tempData &= ^((0xFFFFFFFF & 0x000000FF) << 16)

	tempData |= ((val & 0x000000FF) << 16)

	writeRegister(NtpServerCore.BaseAddrLReg+ntpServer["ConfigModeReg"], &tempData)
	tempData = 0x00000001
	if writeRegister(NtpServerCore.BaseAddrLReg+ntpServer["ConfigControlReg"], &tempData) == 0 {
		showNtpServerPOLLINTERVAL()
	}
}

// NtpServer STRATUM
func writeNtpServerStratumValue(value string) {
	tempData = 0x00000000
	readRegister(NtpServerCore.BaseAddrLReg+ntpServer["ConfigModeReg"], &tempData)

	val, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		log.Fatal("invalid stratum value")
	}
	tempData &= ^((0xFFFFFFFF & 0x000000FF) << 24)

	tempData |= ((val & 0x000000FF) << 24)

	writeRegister(NtpServerCore.BaseAddrLReg+ntpServer["ConfigModeReg"], &tempData)
	tempData = 0x00000001
	if writeRegister(NtpServerCore.BaseAddrLReg+ntpServer["ConfigControlReg"], &tempData) == 0 {
		showNtpServerSTRATUM()
	}

}

// NtpServer REFERENCE ID
func writeNtpServerReferenceId(id string) {
	references := []string{"NTP", "NULL", "LOCL", "CESM", "RBDM", "PPS", "IRIG", "ACTS", "USNO", "PTB", "TDF", "DCF", "MSF", "WWV", "WWVB", "WWVH", "CHU", "LORC", "OMEG", "GPS"}

	if slices.Contains(references, id) {

		if len(id) == 1 {
			tempData |= int64(id[0])
			tempData = tempData << 24
		} else if len(id) == 2 {
			tempData |= int64(id[0])
			tempData = tempData << 8
			tempData |= int64(id[1])
			tempData = tempData << 16
		} else if len(id) == 3 {
			tempData |= int64(id[0])
			tempData = tempData << 8
			tempData |= int64(id[1])
			tempData = tempData << 8
			tempData |= int64(id[2])
			tempData = tempData << 8
		} else if len(id) >= 4 {
			tempData |= int64(id[0])
			tempData = tempData << 8
			tempData |= int64(id[1])
			tempData = tempData << 8
			tempData |= int64(id[2])
			tempData = tempData << 8
			tempData |= int64(id[3])
		}

		fmt.Println("DEBUG: ref id write: ", tempData)

		if writeRegister(NtpServerCore.BaseAddrLReg+ntpServer["ConfigReferenceIdReg"], &tempData) == 0 {
			tempData = 0x00000010

			if writeRegister(NtpServerCore.BaseAddrLReg+ntpServer["ConfigControlReg"], &tempData) == 0 {
				// do nothing if both write successfully
			} else {

				fmt.Println("REF ID NOT CHANGED")
			}
		} else {
			fmt.Println("REF ID NOT CHANGED")

		}
	} else {
		log.Fatal("Please enter one of these valid references: ", references)
	}

}

// NtpServer UTC SMEARING
func writeNtpServerUTCSmearing(en string) {
	tempData = 0x00000000
	readRegister(NtpServerCore.BaseAddrLReg+ntpServer["UtcInfoReg"], &tempData)

	tempData &= ^0x00000100 // Clear bit (using NOT + AND)

	if en == "enabled" {
		tempData |= 0x00000100

	} else if en == "disabled" {
		tempData |= 0x00000000
	}

	writeRegister(NtpServerCore.BaseAddrLReg+ntpServer["UtcInfoReg"], &tempData)
	tempData = 0x00000003 // write utc info and leap
	writeRegister(NtpServerCore.BaseAddrLReg+ntpServer["UtcInfoControlReg"], &tempData)

	showNtpServerUTCSMEARING()

}

//NtpServerUTCLEAP61INPROGRESS READ ONLY :TODO
//NtpServerUTCLEAP59INPROGRESS READ ONLY :TODO

// NtpServer UTC LEAP 61
func writeNtpServerUTCLeap61(en string) {
	tempData = 0x00000000
	readRegister(NtpServerCore.BaseAddrLReg+ntpServer["UtcInfoReg"], &tempData)

	tempData &= ^0x00000800 // Clear bit (using NOT + AND)

	if en == "enabled" {
		tempData |= 0x00000800

	} else if en == "disabled" {
		tempData |= 0x00000000
	}

	writeRegister(NtpServerCore.BaseAddrLReg+ntpServer["UtcInfoReg"], &tempData)
	tempData = 0x00000003 // write utc info and leap
	writeRegister(NtpServerCore.BaseAddrLReg+ntpServer["UtcInfoControlReg"], &tempData)

	showNtpServerUTCLEAP61()
}

// NtpServer UTC LEAP 59
func writeNtpServerUTCLeap59(en string) {
	tempData = 0x00000000
	readRegister(NtpServerCore.BaseAddrLReg+ntpServer["UtcInfoReg"], &tempData)

	tempData &= ^0x00001000 // Clear bit (using NOT + AND)

	if en == "enabled" {
		tempData |= 0x00001000

	} else if en == "disabled" {
		tempData |= 0x00000000
	}

	writeRegister(NtpServerCore.BaseAddrLReg+ntpServer["UtcInfoReg"], &tempData)
	tempData = 0x00000003 // write utc info and leap
	writeRegister(NtpServerCore.BaseAddrLReg+ntpServer["UtcInfoControlReg"], &tempData)

	showNtpServerUTCLEAP59()
}

// NtpServer UTC OFFSET ENABLE
func writeNtpServerUTCOffsetEnable(en string) {
	tempData = 0x00000000
	readRegister(NtpServerCore.BaseAddrLReg+ntpServer["UtcInfoReg"], &tempData)

	tempData &= ^0x00002000 // Clear bit (using NOT + AND)

	if en == "enabled" {
		tempData |= 0x00002000

	} else if en == "disabled" {
		tempData |= 0x00000000
	}

	writeRegister(NtpServerCore.BaseAddrLReg+ntpServer["UtcInfoReg"], &tempData)
	tempData = 0x00000003 // write utc info and leap
	writeRegister(NtpServerCore.BaseAddrLReg+ntpServer["UtcInfoControlReg"], &tempData)

	showNtpServerUTCOFFSETENABLE()
}

// NtpServer UTC OFFSET VALUE

func writeNtpServerUTCOffsetValue(value string) {

	tempData = 0x00000000
	readRegister(NtpServerCore.BaseAddrLReg+ntpServer["UtcInfoReg"], &tempData)
	tempData &= 0x0000FFFF // keep the current value in the lower part of the register

	val, err := strconv.ParseInt(value, 10, 64)

	if err != nil {
		log.Fatal("UTC OFFSET VALUE ERROR")
	} else {

		//temp_data |= ((temp_string.toUInt(nullptr, 10) & 0x0000FFFF) << 16);

		tempData |= ((val & 0x0000FFFF) << 16)

		writeRegister(NtpServerCore.BaseAddrLReg+ntpServer["UtcInfoReg"], &tempData)
		tempData = 0x00000003 // write utc info and leap
		writeRegister(NtpServerCore.BaseAddrLReg+ntpServer["UtcInfoControlReg"], &tempData)

		showNtpServerUTCOFFSETVALUE()
	}

}

// NtpServerREQUESTCOUNT READ ONLY :TODO
// NtpServerRESPONSECOUNT READ ONLY : TODO
// NtpServerREQUESTSDROPPED READ ONLY : TODO
// NtpServerBROADCASTCOUNT READ ONLY : TODO
// NtpServer COUNT CONTROL
func writeNtpServerCountControl() {
	tempData = 0x00000000
	tempData |= 0x00000001 // enable

	writeRegister(NtpServerCore.BaseAddrLReg+ntpServer["CountControlReg"], &tempData)

	showNtpServerBROADCASTCOUNT()
}

//NtpServerVERSION

// NTP SERVER IP ADDRESS HELPER FUNCTIONS

func ip4_addr_to_register_value(addr string) int64 {
	tempAddr := make([]int64, 16)
	byteAddr := net.ParseIP(addr)
	byteAddr = byteAddr.To4()
	if byteAddr != nil {

		//fmt.Println("ipv4 ip addr? : ", byteAddr)

		for i, b := range byteAddr {
			//fmt.Println(b)
			tempAddr[i] = int64(b)
		}

		tempData = 0x00000000
		tempData |= tempAddr[3] & 0x000000FF
		tempData = tempData << 8
		tempData |= tempAddr[2] & 0x000000FF
		tempData = tempData << 8
		tempData |= tempAddr[1] & 0x000000FF
		tempData = tempData << 8
		tempData |= tempAddr[0] & 0x000000FF

		//fmt.Println("ipv4 ip? : ", tempData)

		if writeRegister(NtpServerCore.BaseAddrLReg+ntpServer["ConfigIpReg"], &tempData) == 0 {
			tempData = 0x00000008 // write
			if writeRegister(NtpServerCore.BaseAddrLReg+ntpServer["ConfigControlReg"], &tempData) == 0 {
				showNtpServerIPADDRESS()
			} else {
				fmt.Println("IP UPDATE FAILED")
			}
		} else {
			fmt.Println("IP UPDATE FAILED")
		}

	} else {

		fmt.Println("Unabled to convert IPv6 Address to IPv4 - IP set to 0.0.0.0")
		ip4_addr_to_register_value("0.0.0.0")

		///	log.Fatal("In IPv4 Mode. Please enter an IPv4 address or which modes.")
	}

	return 0

}

func ip6_addr_to_register_value(addr string) int64 {

	tempAddr := make([]int64, 16)
	byteAddr := net.ParseIP(addr)
	byteAddr = byteAddr.To16()

	if byteAddr != nil {

		for i, b := range byteAddr {
			tempAddr[i] = int64(b)
		}

		tempData = 0x00000000
		tempData |= tempAddr[3] & 0x000000FF
		tempData = tempData << 8
		tempData |= tempAddr[2] & 0x000000FF
		tempData = tempData << 8
		tempData |= tempAddr[1] & 0x000000FF
		tempData = tempData << 8
		tempData |= tempAddr[0] & 0x000000FF

		if writeRegister(NtpServerCore.BaseAddrLReg+ntpServer["ConfigIpReg"], &tempData) == 0 {
			tempData = 0x00000000
			tempData |= tempAddr[7] & 0x000000FF
			tempData = tempData << 8
			tempData |= tempAddr[6] & 0x000000FF
			tempData = tempData << 8
			tempData |= tempAddr[5] & 0x000000FF
			tempData = tempData << 8
			tempData |= tempAddr[4] & 0x000000FF

			if writeRegister(NtpServerCore.BaseAddrLReg+ntpServer["ConfigIpv61Reg"], &tempData) == 0 {
				tempData = 0x00000000
				tempData |= tempAddr[11] & 0x000000FF
				tempData = tempData << 8
				tempData |= tempAddr[10] & 0x000000FF
				tempData = tempData << 8
				tempData |= tempAddr[9] & 0x000000FF
				tempData = tempData << 8
				tempData |= tempAddr[8] & 0x000000FF

				if writeRegister(NtpServerCore.BaseAddrLReg+ntpServer["ConfigIpv62Reg"], &tempData) == 0 {
					tempData = 0x00000000
					tempData |= tempAddr[15] & 0x000000FF
					tempData = tempData << 8
					tempData |= tempAddr[14] & 0x000000FF
					tempData = tempData << 8
					tempData |= tempAddr[13] & 0x000000FF
					tempData = tempData << 8
					tempData |= tempAddr[12] & 0x000000FF

					if writeRegister(NtpServerCore.BaseAddrLReg+ntpServer["ConfigIpv63Reg"], &tempData) == 0 {
						tempData = 0x00000008 // write
						if writeRegister(NtpServerCore.BaseAddrLReg+ntpServer["ConfigControlReg"], &tempData) == 0 {

							showNtpServerIPADDRESS()

						} else {
							fmt.Println("IP UPDATE FAILED")
						}
					} else {
						fmt.Println("IP UPDATE FAILED")
					}
				} else {
					fmt.Println("IP UPDATE FAILED")
				}
			} else {
				fmt.Println("IP UPDATE FAILED")
			}
		} else {
			fmt.Println("IP UPDATE FAILED")

		}
	}
	return 0

}
