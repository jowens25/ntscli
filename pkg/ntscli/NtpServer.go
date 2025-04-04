package ntscli

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/spf13/pflag"
)

type NtpServerType struct {
	ControlReg           int64
	StatusReg            int64
	VersionReg           int64
	CountControlReg      int64
	CountReqReg          int64
	CountRespReg         int64
	CountReqDroppedReg   int64
	CountBroadcastReg    int64
	ConfigControlReg     int64
	ConfigModeReg        int64
	ConfigVlanReg        int64
	ConfigMac1Reg        int64
	ConfigMac2Reg        int64
	ConfigIpReg          int64
	ConfigIpv61Reg       int64
	ConfigIpv62Reg       int64
	ConfigIpv63Reg       int64
	ConfigReferenceIdReg int64
	UtcInfoControlReg    int64
	UtcInfoReg           int64
}

var ntpServer = NtpServerType{

	ControlReg:           0x00000000,
	StatusReg:            0x00000004,
	VersionReg:           0x0000000C,
	CountControlReg:      0x00000010,
	CountReqReg:          0x00000014,
	CountRespReg:         0x00000018,
	CountReqDroppedReg:   0x0000001C,
	CountBroadcastReg:    0x00000020,
	ConfigControlReg:     0x00000080,
	ConfigModeReg:        0x00000084,
	ConfigVlanReg:        0x00000088,
	ConfigMac1Reg:        0x0000008C,
	ConfigMac2Reg:        0x00000090,
	ConfigIpReg:          0x00000094,
	ConfigIpv61Reg:       0x00000098,
	ConfigIpv62Reg:       0x0000009C,
	ConfigIpv63Reg:       0x000000A0,
	ConfigReferenceIdReg: 0x000000A4,
	UtcInfoControlReg:    0x00000100,
	UtcInfoReg:           0x00000104,
}
var NtpCore Core

var tempData int64

func Ntp(properties *pflag.FlagSet) {
	fmt.Println("MAIN NTP FUNCTION CALL")
	if DeviceHasNtpServer(&device) != 0 {
		log.Fatal("device does not have an ntp server")
	}
	NtpCore = device.Cores["NtpServerCoreType"]

	properties.SortFlags = false
	properties.Visit(func(f *pflag.Flag) {

		switch f.Name {
		case "enable":
			EnableNtp()
		case "disable":
			DisableNtp()
		case "show":
			fmt.Println(NtpCore)
		}

	})

}

func NtpWrite(properties *pflag.FlagSet) {
	if DeviceHasNtpServer(&device) != 0 {
		log.Fatal("device does not have an ntp server")
	}
	NtpCore = device.Cores["NtpServerCoreType"]
	properties.SortFlags = false
	properties.Visit(func(f *pflag.Flag) {

		switch f.Name {
		case "ip":
			ip, err := properties.GetString(f.Name)
			if err != nil {
				log.Fatal("No such argument for property: ", f.Name, err)
			} else {
				writeNtpServerMac(ip)

			}

		case "mac":

			mac, err := properties.GetString(f.Name)
			if err != nil {
				log.Fatal("No such argument for property: ", f.Name, err)
			} else {
				updateNtpServerMac(mac)

			}

		case "vlan":

			value, err := properties.GetString(f.Name)
			fmt.Println(value)
			if err != nil {
				log.Fatal("No such argument for property: ", f.Name, err)
			} else {
				writeNtpServerVlanEnable(value)

			}

		}

	})

}

func readNtpServerIpMode() string {
	tempData = 0x00000000
	ipMode := ""
	// mode & server config
	if readRegister(NtpCore.BaseAddrLReg+ntpServer.ConfigModeReg, &tempData) == 0 {

		if ((tempData >> 0) & 0x00000003) == 1 {
			ipMode = "IPv4"
		} else if ((tempData >> 0) & 0x00000003) == 2 {
			ipMode = "IPv6"
		} else {
			ipMode = "NA"
		}
		//return result
	} else {
		ipMode = "NA" // IPv4 IPv6 NA

	}

	return ipMode
}

func readNtpServerIpAddress() string {
	tempData = 0x00000000
	ipAddr := ""

	// ip
	if readNtpServerIpMode() == "IPv4" {
		if readRegister(NtpCore.BaseAddrLReg+ntpServer.ConfigIpReg, &tempData) == 0 {
			var temp_ip int64 = 0x00000000
			temp_ip |= (tempData >> 0) & 0x000000FF
			temp_ip = temp_ip << 8
			temp_ip |= (tempData >> 8) & 0x000000FF
			temp_ip = temp_ip << 8
			temp_ip |= (tempData >> 16) & 0x000000FF
			temp_ip = temp_ip << 8
			temp_ip |= (tempData >> 24) & 0x000000FF

			ipAddr = int_to_ip_addr(temp_ip)

		} else {
			ipAddr = "NA"
		}
	} else if readNtpServerIpMode() == "IPv6" {
		temp_ip6 := make([]int64, 16)
		if readRegister(NtpCore.BaseAddrLReg+ntpServer.ConfigIpReg, &tempData) == 0 {
			temp_ip6[0] = (tempData >> 0) & 0x000000FF
			temp_ip6[1] = (tempData >> 8) & 0x000000FF
			temp_ip6[2] = (tempData >> 16) & 0x000000FF
			temp_ip6[3] = (tempData >> 24) & 0x000000FF

			if readRegister(NtpCore.BaseAddrLReg+ntpServer.ConfigIpv61Reg, &tempData) == 0 {
				temp_ip6[4] = (tempData >> 0) & 0x000000FF
				temp_ip6[5] = (tempData >> 8) & 0x000000FF
				temp_ip6[6] = (tempData >> 16) & 0x000000FF
				temp_ip6[7] = (tempData >> 24) & 0x000000FF

				if readRegister(NtpCore.BaseAddrLReg+ntpServer.ConfigIpv62Reg, &tempData) == 0 {
					temp_ip6[8] = (tempData >> 0) & 0x000000FF
					temp_ip6[9] = (tempData >> 8) & 0x000000FF
					temp_ip6[10] = (tempData >> 16) & 0x000000FF
					temp_ip6[11] = (tempData >> 24) & 0x000000FF

					if readRegister(NtpCore.BaseAddrLReg+ntpServer.ConfigIpv63Reg, &tempData) == 0 {

						temp_ip6[12] = (tempData >> 0) & 0x000000FF
						temp_ip6[13] = (tempData >> 8) & 0x000000FF
						temp_ip6[14] = (tempData >> 16) & 0x000000FF
						temp_ip6[15] = (tempData >> 24) & 0x000000FF

						log.Println("IPv6 Addr: ", temp_ip6)

						ipAddr = int_to_ipv6(temp_ip6)
						//temp_string = QHostAddress(temp_ip6).toString() ????

						//ntpServer.IpAddr = string(temp_ip6)

					} else {
						ipAddr = "NA"
					}

				} else {
					ipAddr = "NA"
				}

			} else {
				ipAddr = "NA"
			}

		} else {
			ipAddr = "NA"
		}
	} else {
		ipAddr = "NA"
	}

	return ipAddr
}

func NtpRead(properties *pflag.FlagSet) {
	if DeviceHasNtpServer(&device) != 0 {
		log.Fatal("device does not have an ntp server")
	}
	NtpCore = device.Cores["NtpServerCoreType"]

	properties.SortFlags = false
	properties.Visit(func(f *pflag.Flag) {
		//log.Println(f.Name)
		//readNtpServerMode()
		switch f.Name {
		case "ip":
			//readNtpServerIpAddress()

		case "mac":
			readNtpServerMac()

		case "all":
			NtpReadPrintAll()

		//case "control":
		//tempData = 0x00000000
		//readRegister(NtpCore.BaseAddrLReg+ntpServer.ConfigControlReg, &tempData)
		//fmt.Print(tempData)

		case "vlan":
			fmt.Println("NTP SERVER VLAN Enabled: ", readNtpServerVlanEnable())
			fmt.Println("NTP SERVER VLAN VALUE: ", readNtpServerVlanValue())
		}

	})
}

func NtpReadPrintAll() {

	fmt.Println("NTP SERVER:                           ", readNtpServerEnable())
	fmt.Println("NTP SERVER INSTANCE:                  ", NtpCore.InstanceNumber)
	fmt.Println("NTP SERVER IP ADDRESS:                ", readNtpServerIpAddress())
	fmt.Println("NTP SERVER IP MODE:                   ", readNtpServerIpMode())
	fmt.Println("NTP SERVER MAC ADDRESS:               ", readNtpServerMac())
	fmt.Println("NTP SERVER VLAN ENABLED:              ", readNtpServerVlanEnable())
	fmt.Println("NTP SERVER VLAN VALUE:                ", readNtpServerVlanValue())
	fmt.Println("NTP SERVER UNICAST:                   ", readNtpServerUnicastMode())
	fmt.Println("NTP SERVER MULTICAST:                 ", readNtpServerMulticastMode())
	fmt.Println("NTP SERVER BROADCAST:                 ", readNtpServerBroadcastMode())
	fmt.Println("NTP SERVER PRECISION:                 ", readNtpServerPrecisionValue())
	fmt.Println("NTP SERVER POLL INTERVAL:             ", readNtpServerPollIntervalValue())
	fmt.Println("NTP SERVER STRATUM:                   ", readNtpServerStratumValue())
	fmt.Println("NTP SERVER REFERENCE ID:              ", readNtpServerReferenceId())
	for k, v := range readNtpServerUTC() {
		fmt.Println("NTP SERVER", k, ":", v)
	}

	fmt.Println("NTP SERVER REQUEST COUNT:             ", readNtpServerRequestCount())
	fmt.Println("NTP SERVER RESPONSE COUNT:            ", readNtpServerResponseCount())
	fmt.Println("NTP SERVER REQUESTS DROPPED:          ", readNtpServerRequestsDropped())
	fmt.Println("NTP SERVER BROADCAST COUNT:           ", readNtpServerBroadcastCount())
	fmt.Println("NTP SERVER COUNT CONTROL:             ", readNtpServerCountControl())
	fmt.Println("NTP SERVER VERSION:                   ", readNtpServerVersion())

	//readNtpServerMode()
}

func NtpList() {

}

func DeviceHasNtpServer(dev *Device) int64 {
	for _, core := range dev.Cores {
		if core.CoreType == types.NtpServerCoreType {
			return 0
		}
	}

	return -1
}

// enable block
func EnableNtp() {

	tempData = 0x00000001
	//tempData |= 0x00000001
	if writeRegister(NtpCore.BaseAddrLReg+ntpServer.ControlReg, &tempData) == 0 {
		fmt.Println("VERBOSE NTP SERVER: ", readNtpServerEnable())
	} else {
		log.Fatal(" VERBOSE ERROR WRITING NTP")
	}
}

func DisableNtp() {

	tempData = 0x00000000
	if writeRegister(NtpCore.BaseAddrLReg+ntpServer.ControlReg, &tempData) == 0 {
		fmt.Println("VERBOSE NTP SERVER: ", readNtpServerEnable())
	} else {
		log.Fatal(" VERBOSE ERROR WRITING NTP")
	}
}

func readNtpServerEnable() string {
	tempData = 0x00000000
	enabled := ""
	if readRegister(NtpCore.BaseAddrLReg+ntpServer.ControlReg, &tempData) == 0 {
		if (tempData & 0x00000001) == 0 {
			enabled = "disabled"
		} else {
			enabled = "enabled"
		}
	} else {
		enabled = "disabled"
	}
	return enabled
}

// enable block

// mac address block

func updateNtpServerMac(macAddr string) {
	writeNtpServerMac(macAddr)
	fmt.Println("VERBOSE NTP SERVER MAC ADDRESS: ", readNtpServerMac())
}

func readNtpServerMac() string {
	tempData = 0x00000000
	mac := ""
	// mac
	this_string := make([]byte, 0, 32)

	if readRegister(NtpCore.BaseAddrLReg+ntpServer.ConfigMac1Reg, &tempData) == 0 {

		this_string = append(this_string, fmt.Sprintf("%02x", (tempData>>0)&0x000000FF)...)
		this_string = append(this_string, ':')

		this_string = append(this_string, fmt.Sprintf("%02x", (tempData>>8)&0x000000FF)...)
		this_string = append(this_string, ':')
		this_string = append(this_string, fmt.Sprintf("%02x", (tempData>>16)&0x000000FF)...)
		this_string = append(this_string, ':')
		this_string = append(this_string, fmt.Sprintf("%02x", (tempData>>24)&0x000000FF)...)
		this_string = append(this_string, ':')

		if readRegister(NtpCore.BaseAddrLReg+ntpServer.ConfigMac2Reg, &tempData) == 0 {
			this_string = append(this_string, fmt.Sprintf("%02x", (tempData>>0)&0x000000FF)...)
			this_string = append(this_string, ':')

			this_string = append(this_string, fmt.Sprintf("%02x", (tempData>>8)&0x000000FF)...)

			mac = string(this_string)
		} else {
			mac = "NA"
		}

	} else {
		mac = "NA"
	}

	return mac
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

//mac address block

func readNtpServerVlanEnable() string {
	tempData = 0x00000000
	vlanMode := ""
	if readRegister(NtpCore.BaseAddrLReg+ntpServer.ConfigVlanReg, &tempData) == 0 {
		if (tempData & 0x00010000) == 0 {
			vlanMode = "disabled"
		} else {
			vlanMode = "enabled"
		}
	} else {
		vlanMode = "disabled"
	}

	return vlanMode
}

func readNtpServerVlanValue() string {
	tempData = 0x00000000
	vlanValue := ""
	if readRegister(NtpCore.BaseAddrLReg+ntpServer.ConfigVlanReg, &tempData) == 0 {
		tempData &= 0x0000FFFF
		vlanValue = fmt.Sprintf("0x%04x", tempData)
	} else {
		vlanValue = "NA"
	}

	return vlanValue
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

	// vlan
	temp_string := value
	mode := "enable"

	//temp_data = temp_string.toUInt(nullptr, 16)
	temp_data, _ := strconv.ParseInt(string(temp_string), 16, 64)

	//temp_data &= 0x0000FFFF
	if mode == "enable" {
		temp_data |= 0x00010000 // enable
	}

	if temp_string == "NA" {
		//nothing

	} else if writeRegister(NtpCore.BaseAddrLReg+ntpServer.ConfigVlanReg, &temp_data) == 0 {

		temp_data &= 0x0000FFFF

		// return the Vlan Value

		temp_data = 0x00000002 // write

		if writeRegister(NtpCore.BaseAddrLReg+ntpServer.ConfigControlReg, &temp_data) == 0 {
			// nothing
		} else {
			fmt.Println("vlan enabled: false")
			fmt.Println("vlan value: NA")
		}

		fmt.Println("VALUE VALUE VALUE: ", fmt.Sprintf("0x%08x", temp_data))

	} else {
		fmt.Println("vlan enabled: false")
		fmt.Println("vlan value: NA")
	}

}

func readNtpServerUnicastMode() string {
	tempData = 0x00000000
	unicastMode := ""
	// mode & server config
	if readRegister(NtpCore.BaseAddrLReg+ntpServer.ConfigModeReg, &tempData) == 0 {
		if (tempData & 0x00000010) == 0 {
			unicastMode = "DISABLED"
		} else {
			unicastMode = "ENABLED"
		}
	} else {
		unicastMode = "DISABLED"
	}
	return unicastMode

}

func readNtpServerMulticastMode() string {
	tempData = 0x00000000
	multicastMode := ""
	// mode & server config
	if readRegister(NtpCore.BaseAddrLReg+ntpServer.ConfigModeReg, &tempData) == 0 {
		if (tempData & 0x00000020) == 0 {
			multicastMode = "DISABLED"
		} else {
			multicastMode = "ENABLED"
		}

	} else {
		multicastMode = "DISABLED"

	}
	return multicastMode

}

func readNtpServerBroadcastMode() string {
	tempData = 0x00000000
	broadcastMode := ""
	// mode & server config
	if readRegister(NtpCore.BaseAddrLReg+ntpServer.ConfigModeReg, &tempData) == 0 {
		if (tempData & 0x00000040) == 0 {
			broadcastMode = "DISABLED"
		} else {
			broadcastMode = "ENABLED"
		}

	} else {
		broadcastMode = "DISABLED"

	}
	return broadcastMode
}

func readNtpServerPrecisionValue() rune {
	tempData = 0x00000000
	var PrecisionValue rune
	// mode & server config
	if readRegister(NtpCore.BaseAddrLReg+ntpServer.ConfigModeReg, &tempData) == 0 {
		if (tempData & 0x00000040) == 0 {
			PrecisionValue = 'N'
		} else {
			PrecisionValue = rune(int8(((tempData >> 8) & 0x000000FF)))
		}

	} else {
		PrecisionValue = 'N'
	}
	return PrecisionValue
}

func readNtpServerPollIntervalValue() string {
	tempData = 0x00000000
	var PollIntervalValue string
	// mode & server config
	if readRegister(NtpCore.BaseAddrLReg+ntpServer.ConfigModeReg, &tempData) == 0 {
		// /fmt.Println(tempData)
		if (tempData & 0x00000040) == 0 {
			PollIntervalValue = "NA"
		} else {
			PollIntervalValue = strconv.FormatInt(((tempData >> 16) & 0x000000FF), 10) // Base 10

		}

	} else {
		PollIntervalValue = "NA"
	}
	return PollIntervalValue
}

func readNtpServerStratumValue() string {
	tempData = 0x00000000
	StratumValue := ""
	// mode & server config
	if readRegister(NtpCore.BaseAddrLReg+ntpServer.ConfigModeReg, &tempData) == 0 {
		if (tempData & 0x00000040) == 0 {
			StratumValue = "NA"
		} else {
			StratumValue = strconv.FormatInt(((tempData >> 24) & 0x000000FF), 10) // Base 10
		}

	} else {
		StratumValue = "NA"
	}
	return StratumValue
}

func readNtpServerReferenceId() string {
	tempData = 0x00000000
	referenceId := ""
	// reference id // no ref on UI??
	if readRegister(NtpCore.BaseAddrLReg+ntpServer.ConfigReferenceIdReg, &tempData) == 0 {
		var temp_string []byte
		temp_string = append(temp_string, byte(((tempData >> 24) & 0x000000FF)))
		temp_string = append(temp_string, byte(((tempData >> 16) & 0x000000FF)))
		temp_string = append(temp_string, byte(((tempData >> 8) & 0x000000FF)))
		temp_string = append(temp_string, byte(((tempData >> 0) & 0x000000FF)))
		referenceId = string(temp_string) // TODO
	} else {
		referenceId = "NA"
	}

	return referenceId
}

func readNtpServerUTC() map[string]string {

	result := make(map[string]string)
	// utc info
	tempData = 0x40000000
	utcSmearing := ""
	utcLeap61InProgress := ""
	utcLeap59InProgress := ""
	utcLeap61 := ""
	utcLeap59 := ""
	utcOffsetVal := ""
	utcOffsetValue := ""

	if writeRegister(NtpCore.BaseAddrLReg+ntpServer.UtcInfoControlReg, &tempData) == 0 {
		for i := range 10 {
			if readRegister(NtpCore.BaseAddrLReg+ntpServer.UtcInfoControlReg, &tempData) == 0 {
				if (tempData & 0x80000000) != 0 {
					if readRegister(NtpCore.BaseAddrLReg+ntpServer.UtcInfoReg, &tempData) == 0 {
						if (tempData & 0x00000100) == 0 {
							utcSmearing = "false"
						} else {
							utcSmearing = "true"

						}

						if (tempData & 0x00000200) == 0 {
							utcLeap61InProgress = "false"
						} else {
							utcLeap61InProgress = "true"
						}

						if (tempData & 0x00000400) == 0 {
							utcLeap59InProgress = "false"
						} else {
							utcLeap59InProgress = "true"
						}

						if (tempData & 0x00000800) == 0 {
							utcLeap61 = "false"
						} else {
							utcLeap61 = "true"
						}

						if (tempData & 0x00001000) == 0 {
							utcLeap59 = "false"
						} else {
							utcLeap59 = "true"
						}

						if (tempData & 0x00002000) == 0 {
							utcOffsetVal = "false"
						} else {
							utcOffsetVal = "true"
						}

						//log.Println("ui->NtpServerUtcOffsetValue->setText(QString::number(((tempData >> 16) & 0x0000FFFF)));")

						utcOffsetValue = strconv.FormatInt((tempData>>16)&0x0000FFFF, 10) // Base 10

						//ntpServer.UtcOffsetValue = string((tempData >> 16) & 0x0000FFFF)
						//log.Println("ntpServer.UtcOffsetValue: ", ntpServer.UtcOffsetValue)
					} else {
						utcSmearing = "false"
						utcLeap61InProgress = "false"
						utcLeap59InProgress = "false"
						utcLeap61 = "false"
						utcLeap59 = "false"
						utcOffsetVal = "false"
						utcOffsetValue = "0"

					}
					break
				} else if i == 9 {
					log.Fatal("utc read incomplete")
					utcSmearing = "false"
					utcLeap61InProgress = "false"
					utcLeap59InProgress = "false"
					utcLeap61 = "false"
					utcLeap59 = "false"
					utcOffsetVal = "false"
					utcOffsetValue = "0"
				}

			} else {
				utcSmearing = "false"
				utcLeap61InProgress = "false"
				utcLeap59InProgress = "false"
				utcLeap61 = "false"
				utcLeap59 = "false"
				utcOffsetVal = "false"
				utcOffsetValue = "0"
			}
		}
	} else {
		utcSmearing = "false"
		utcLeap61InProgress = "false"
		utcLeap59InProgress = "false"
		utcLeap61 = "false"
		utcLeap59 = "false"
		utcOffsetVal = "false"
		utcOffsetValue = "0"
	}

	result["UTC SMEARING"] = "             " + utcSmearing
	result["UTC LEAP 61 IN PROGRESS"] = "  " + utcLeap61InProgress
	result["UTC LEAP 59 IN PROGRESS"] = "  " + utcLeap59InProgress
	result["UTC LEAP 61 "] = "             " + utcLeap61
	result["UTC LEAP 59 "] = "             " + utcLeap59
	result["UTC OFFSET VAL"] = "           " + utcOffsetVal
	result["UTC OFFSET VALUE"] = "         " + utcOffsetValue

	return result
}

func readNtpServerRequestCount() string {
	tempData = 0x00000000
	requests := ""
	if readRegister(NtpCore.BaseAddrLReg+ntpServer.CountReqReg, &tempData) == 0 {
		requests = fmt.Sprintf("%d", tempData)
	} else {
		requests = "NA"
	}
	return requests
}

func readNtpServerResponseCount() string {
	tempData = 0x00000000
	responses := ""
	if readRegister(NtpCore.BaseAddrLReg+ntpServer.CountRespReg, &tempData) == 0 {
		responses = fmt.Sprintf("%d", tempData)

	} else {
		responses = "NA"
	}
	return responses
}

func readNtpServerRequestsDropped() string {
	tempData = 0x00000000
	dropped := ""
	if readRegister(NtpCore.BaseAddrLReg+ntpServer.CountReqDroppedReg, &tempData) == 0 {
		dropped = fmt.Sprintf("%d", tempData)

	} else {
		dropped = "NA"
	}
	return dropped
}

func readNtpServerBroadcastCount() string {
	tempData = 0x00000000
	broadcasts := ""
	if readRegister(NtpCore.BaseAddrLReg+ntpServer.CountBroadcastReg, &tempData) == 0 {
		broadcasts = fmt.Sprintf("%d", tempData)

	} else {
		broadcasts = "NA"
	}
	return broadcasts
}

func readNtpServerCountControl() string {
	tempData = 0x00000000
	clear := ""
	if readRegister(NtpCore.BaseAddrLReg+ntpServer.CountControlReg, &tempData) == 0 {
		if (tempData & 0x00000001) == 0 {
			clear = "false"
		} else {
			clear = "true"

		}
	} else {
		clear = "false"

	}

	return clear
}

func readNtpServerVersion() string {
	tempData = 0x00000000
	// version
	version := ""
	if readRegister(NtpCore.BaseAddrLReg+ntpServer.VersionReg, &tempData) == 0 {
		version = fmt.Sprintf("0x%02x", tempData) // base 16 string format
	} else {
		version = "NA"
	}

	return version

}

func int_to_ip_addr(val int64) string {
	hex_string := fmt.Sprintf("%02x", val) // base 16 string format
	return hex_to_decimal(split_into_ip_addr(hex_string))
}

func hex_to_decimal(hex_parts []string) string {

	ip := ""
	decimalValue, _ := strconv.ParseInt(hex_parts[0], 16, 16)
	ip += fmt.Sprint(decimalValue)

	for _, part := range hex_parts[1:] {

		decimalValue, err := strconv.ParseInt(part, 16, 16)
		if err != nil {
			log.Fatal("IP addr error")
		}
		ip += "." + fmt.Sprint(decimalValue)

	}
	return ip
}

func split_into_ip_addr(hex_string string) []string {
	var parts = []string{hex_string[0:2], hex_string[2:4], hex_string[4:6], hex_string[6:8]}
	return parts
}

func int_to_ipv6(addr []int64) string {
	return "::ffff:0" + fmt.Sprintf("%d", addr[0]) + "." + fmt.Sprintf("%d", addr[1]) + "." + fmt.Sprintf("%d", addr[2]) + "." + fmt.Sprintf("%d", addr[3])
}
