package ntscli

import (
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/pflag"
)

type NtpServer struct {
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

var ntpServer = NtpServer{

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

func NtpWrite(properties *pflag.FlagSet) {
	properties.SortFlags = false
	properties.Visit(func(f *pflag.Flag) {

		switch f.Name {
		case "ip":
			ip, err := properties.GetString(f.Name)
			if err != nil {
				log.Fatal("No such argument for property: ", f.Name, err)
			} else {
				ntpServerSetIpAddress(ip)

			}

		case "mac":

			mac, err := properties.GetString(f.Name)
			if err != nil {
				log.Fatal("No such argument for property: ", f.Name, err)
			} else {
				ntpServerSetMacAddress(mac)

			}

		}

	})

}

func NtpRead(properties *pflag.FlagSet) {

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
			fmt.Println("NTP SERVER ENABLED:     ", readNtpServerEnable())
			fmt.Println("NTP SERVER MAC ADDRESS: ", readNtpServerMac())
		}

	})
}

func ntpServerSetIpAddress(input string) {
	fmt.Println("ntp server ip addr: ", input)
}

func ntpServerSetMacAddress(input string) {
	fmt.Println("ntp server mac addr: ", input)
}

func Ntp(properties *pflag.FlagSet) {
	properties.SortFlags = false
	properties.Visit(func(f *pflag.Flag) {

		switch f.Name {
		case "enable":
			EnableNtp()
		case "disable":
			DisableNtp()
		}

	})

}

func NtpList() {
	coreConfig := CoreConfig{}
	readDeviceConfig()
	if deviceHasNtpServer(&deviceConfig, &coreConfig) == 0 {
	} else {
		log.Println("whats going on?")
	}

}

func getNtpCore() *CoreConfig {

	readDeviceConfig()

	coreConfig := CoreConfig{}

	if deviceHasNtpServer(&deviceConfig, &coreConfig) != 0 {
		log.Fatal("Device has no NTP SERVER")
	}

	return &coreConfig

}

func deviceHasNtpServer(deviceConfig *DeviceConfig, coreConfig *CoreConfig) int64 {
	for _, core := range deviceConfig.Cores {
		if core.CoreType == types.NtpServerCoreType {
			*coreConfig = core
			return 0
		}
	}

	return -1
}

func EnableNtp() {
	ntpCore := getNtpCore()
	var tempData int64 = 0x00000000
	tempData |= 0x00000001

	if writeRegister(ntpCore.BaseAddrLReg+ntpServer.ControlReg, &tempData) == 0 {
		if readRegister(ntpCore.BaseAddrLReg+ntpServer.ControlReg, &tempData) == 0 {
			if (tempData & 0x00000001) == 1 {
				fmt.Println("VERBOSE: NTP ENABLED")

			}

		}
	} else {
		log.Fatal(" VERBOSE ERROR WRITING NTP")
	}

}

func DisableNtp() {
	ntpCore := getNtpCore()

	var tempData int64 = 0x00000000

	if writeRegister(ntpCore.BaseAddrLReg+ntpServer.ControlReg, &tempData) == 0 {
		fmt.Println("VERBOSE: NTP DISABLED")
	} else {
		log.Fatal(" VERBOSE ERROR WRITING NTP")
	}
}

func readNtpServerEnable() string {
	ntpCore := getNtpCore()
	var tempData int64 = 0x00000000
	if readRegister(ntpCore.BaseAddrLReg+ntpServer.ControlReg, &tempData) == 0 {
		if (tempData & 0x00000001) == 0 {
			return "disabled"
		} else {
			return "enabled"
		}
	} else {
		return "disabled"
	}

}

func readNtpServerMac() string {
	ntpCore := getNtpCore()
	var tempData int64 = 0x00000000

	// mac
	this_string := make([]byte, 0, 32)

	if readRegister(ntpCore.BaseAddrLReg+ntpServer.ConfigMac1Reg, &tempData) == 0 {

		this_string = append(this_string, fmt.Sprintf("%02x", (tempData>>0)&0x000000FF)...)
		this_string = append(this_string, ':')

		this_string = append(this_string, fmt.Sprintf("%02x", (tempData>>8)&0x000000FF)...)
		this_string = append(this_string, ':')
		this_string = append(this_string, fmt.Sprintf("%02x", (tempData>>16)&0x000000FF)...)
		this_string = append(this_string, ':')
		this_string = append(this_string, fmt.Sprintf("%02x", (tempData>>24)&0x000000FF)...)
		this_string = append(this_string, ':')

		if readRegister(ntpCore.BaseAddrLReg+ntpServer.ConfigMac2Reg, &tempData) == 0 {
			this_string = append(this_string, fmt.Sprintf("%02x", (tempData>>0)&0x000000FF)...)
			this_string = append(this_string, ':')

			this_string = append(this_string, fmt.Sprintf("%02x", (tempData>>8)&0x000000FF)...)

			return string(this_string)
		} else {
			return "NA"
		}

	} else {
		return "NA"
	}

}

/*
	func readNtpServerVlan() {
		ntpCore := getNtpCore()
		var tempData int64 = 0x00000000
		if readRegister(ntpCore.BaseAddrLReg+ntpServer.ConfigVlanReg, &tempData) == 0 {
			if (tempData & 0x00010000) == 0 {
				ntpServer.VlanEnable = false
			} else {
				ntpServer.VlanEnable = true
			}

			tempData &= 0x0000FFFF
			ntpServer.VlanValue = fmt.Sprintf("0x%08x", &tempData)

		} else {
			ntpServer.VlanEnable = false
			ntpServer.VlanValue = "NA"

		}
	}

	func readNtpServerMode() {
		ntpCore := getNtpCore()
		var tempData int64 = 0x00000000

		// mode & server config
		if readRegister(ntpCore.BaseAddrLReg+ntpServer.ConfigModeReg, &tempData) == 0 {
			if ((tempData >> 0) & 0x00000003) == 1 {
				ntpServer.IpMode = "IPv4"
			} else if ((tempData >> 0) & 0x00000003) == 2 {
				ntpServer.IpMode = "IPv6"
			} else {
				ntpServer.IpMode = "NA"
			}

			if (tempData & 0x00000010) == 0 {
				ntpServer.UnicastMode = false
			} else {
				ntpServer.UnicastMode = true
			}

			if (tempData & 0x00000020) == 0 {
				ntpServer.MulticastMode = false
			} else {
				ntpServer.MulticastMode = true
			}

			if (tempData & 0x00000040) == 0 {
				ntpServer.BroadcastMode = false
			} else {
				ntpServer.BroadcastMode = true
			}

			ntpServer.PrecisionValue = rune(int8(((tempData >> 8) & 0x000000FF)))
			ntpServer.PollIntervalValue = string(((tempData >> 16) & 0x000000FF))
			ntpServer.StratumValue = string((tempData >> 24) & 0x000000FF)

		} else {
			ntpServer.IpMode = "NA"
			ntpServer.UnicastMode = false
			ntpServer.MulticastMode = false
			ntpServer.BroadcastMode = false

			ntpServer.PrecisionValue = 'N'
			ntpServer.PollIntervalValue = "NA"
			ntpServer.StratumValue = "NA"

		}
	}

	func readNtpServerReferenceId() {
		ntpCore := getNtpCore()
		var tempData int64 = 0x00000000
		// reference id // no ref on UI??
		if readRegister(ntpCore.BaseAddrLReg+ntpServer.ConfigReferenceIdReg, &tempData) == 0 {
			var temp_string []byte
			temp_string = append(temp_string, byte(((tempData >> 24) & 0x000000FF)))
			temp_string = append(temp_string, byte(((tempData >> 16) & 0x000000FF)))
			temp_string = append(temp_string, byte(((tempData >> 8) & 0x000000FF)))
			temp_string = append(temp_string, byte(((tempData >> 0) & 0x000000FF)))
			ntpServer.ReferenceId = fmt.Sprintf("0x%08x", temp_string) // TODO
		} else {
			ntpServer.ReferenceId = "BANANA"
		}
	}

func readNtpServerIpAddress() {

		ntpCore := getNtpCore()
		var tempData int64 = 0x00000000

		// ip
		if ntpServer.IpMode == "IPv4" {
			if readRegister(ntpCore.BaseAddrLReg+ntpServer.ConfigIpReg, &tempData) == 0 {
				var temp_ip int64 = 0x00000000
				temp_ip |= (tempData >> 0) & 0x000000FF
				temp_ip = temp_ip << 8
				temp_ip |= (tempData >> 8) & 0x000000FF
				temp_ip = temp_ip << 8
				temp_ip |= (tempData >> 16) & 0x000000FF
				temp_ip = temp_ip << 8
				temp_ip |= (tempData >> 24) & 0x000000FF

				ntpServer.IpAddr = int_to_ip_addr(temp_ip)

			} else {
				ntpServer.IpAddr = "NA"
			}
		} else if ntpServer.IpMode == "IPv6" {
			temp_ip6 := make([]int64, 16)
			if readRegister(ntpCore.BaseAddrLReg+ntpServer.ConfigIpReg, &tempData) == 0 {
				temp_ip6[0] = (tempData >> 0) & 0x000000FF
				temp_ip6[1] = (tempData >> 8) & 0x000000FF
				temp_ip6[2] = (tempData >> 16) & 0x000000FF
				temp_ip6[3] = (tempData >> 24) & 0x000000FF

				if readRegister(ntpCore.BaseAddrLReg+ntpServer.ConfigIpv61Reg, &tempData) == 0 {
					temp_ip6[4] = (tempData >> 0) & 0x000000FF
					temp_ip6[5] = (tempData >> 8) & 0x000000FF
					temp_ip6[6] = (tempData >> 16) & 0x000000FF
					temp_ip6[7] = (tempData >> 24) & 0x000000FF

					if readRegister(ntpCore.BaseAddrLReg+ntpServer.ConfigIpv62Reg, &tempData) == 0 {
						temp_ip6[8] = (tempData >> 0) & 0x000000FF
						temp_ip6[9] = (tempData >> 8) & 0x000000FF
						temp_ip6[10] = (tempData >> 16) & 0x000000FF
						temp_ip6[11] = (tempData >> 24) & 0x000000FF

						if readRegister(ntpCore.BaseAddrLReg+ntpServer.ConfigIpv63Reg, &tempData) == 0 {

							temp_ip6[12] = (tempData >> 0) & 0x000000FF
							temp_ip6[13] = (tempData >> 8) & 0x000000FF
							temp_ip6[14] = (tempData >> 16) & 0x000000FF
							temp_ip6[15] = (tempData >> 24) & 0x000000FF

							log.Println("IPv6 Addr: ", temp_ip6)

							ntpServer.IpAddr = int_to_ipv6(temp_ip6)
							//temp_string = QHostAddress(temp_ip6).toString() ????

							//ntpServer.IpAddr = string(temp_ip6)

						} else {
							ntpServer.IpAddr = "NA"
						}

					} else {
						ntpServer.IpAddr = "NA"
					}

				} else {
					ntpServer.IpAddr = "NA"
				}

			} else {
				ntpServer.IpAddr = "NA"
			}
		} else {
			ntpServer.IpAddr = "NA"
		}
	}

	func readNtpServerUTC() {
		ntpCore := getNtpCore()
		var tempData int64 = 0x00000000

		// utc info
		tempData = 0x40000000

		if writeRegister(ntpCore.BaseAddrLReg+ntpServer.UtcInfoControlReg, &tempData) == 0 {
			for i := range 10 {
				if readRegister(ntpCore.BaseAddrLReg+ntpServer.UtcInfoControlReg, &tempData) == 0 {
					if (tempData & 0x80000000) != 0 {
						if readRegister(ntpCore.BaseAddrLReg+ntpServer.UtcInfoReg, &tempData) == 0 {
							if (tempData & 0x00000100) == 0 {
								ntpServer.UtcSmearing = false
							} else {
								ntpServer.UtcSmearing = true

							}

							if (tempData & 0x00000200) == 0 {
								ntpServer.UtcLeap61InProgress = false
							} else {
								ntpServer.UtcLeap61InProgress = true
							}

							if (tempData & 0x00000400) == 0 {
								ntpServer.UtcLeap59InProgress = false
							} else {
								ntpServer.UtcLeap59InProgress = true
							}

							if (tempData & 0x00000800) == 0 {
								ntpServer.UtcLeap61 = false
							} else {
								ntpServer.UtcLeap61 = true
							}

							if (tempData & 0x00001000) == 0 {
								ntpServer.UtcLeap59 = false
							} else {
								ntpServer.UtcLeap59 = true
							}

							if (tempData & 0x00002000) == 0 {
								ntpServer.UtcOffsetVal = false
							} else {
								ntpServer.UtcOffsetVal = true
							}

							//log.Println("ui->NtpServerUtcOffsetValue->setText(QString::number(((tempData >> 16) & 0x0000FFFF)));")

							ntpServer.UtcOffsetValue = (tempData >> 16) & 0x0000FFFF

							//ntpServer.UtcOffsetValue = string((tempData >> 16) & 0x0000FFFF)
							//log.Println("ntpServer.UtcOffsetValue: ", ntpServer.UtcOffsetValue)
						} else {
							ntpServer.UtcSmearing = false
							ntpServer.UtcLeap61InProgress = false
							ntpServer.UtcLeap59InProgress = false
							ntpServer.UtcLeap61 = false
							ntpServer.UtcLeap59 = false
							ntpServer.UtcOffsetVal = false
							ntpServer.UtcOffsetValue = 0

						}
						break
					} else if i == 9 {
						log.Fatal("utc read in complete")
						ntpServer.UtcSmearing = false
						ntpServer.UtcLeap61InProgress = false
						ntpServer.UtcLeap59InProgress = false
						ntpServer.UtcLeap61 = false
						ntpServer.UtcLeap59 = false
						ntpServer.UtcOffsetVal = false
						ntpServer.UtcOffsetValue = 0
					}

				} else {
					ntpServer.UtcSmearing = false
					ntpServer.UtcLeap61InProgress = false
					ntpServer.UtcLeap59InProgress = false
					ntpServer.UtcLeap61 = false
					ntpServer.UtcLeap59 = false
					ntpServer.UtcOffsetVal = false
					ntpServer.UtcOffsetValue = 0
				}
			}
		} else {
			ntpServer.UtcSmearing = false
			ntpServer.UtcLeap61InProgress = false
			ntpServer.UtcLeap59InProgress = false
			ntpServer.UtcLeap61 = false
			ntpServer.UtcLeap59 = false
			ntpServer.UtcOffsetVal = false
			ntpServer.UtcOffsetValue = 0
		}
	}

	func readNtpServerStatus() {
		ntpCore := getNtpCore()
		var tempData int64 = 0x00000000
		// status
		if readRegister(ntpCore.BaseAddrLReg+ntpServer.CountReqReg, &tempData) == 0 {
			ntpServer.RequestsValue = fmt.Sprintf("%d", tempData)
		} else {
			ntpServer.RequestsValue = "NA"
		}

		if readRegister(ntpCore.BaseAddrLReg+ntpServer.CountRespReg, &tempData) == 0 {
			ntpServer.ResponsesValue = fmt.Sprintf("%d", tempData)

		} else {
			ntpServer.ResponsesValue = "NA"
		}

		if readRegister(ntpCore.BaseAddrLReg+ntpServer.CountReqDroppedReg, &tempData) == 0 {
			ntpServer.RequestsDroppedValue = fmt.Sprintf("%d", tempData)
		} else {
			ntpServer.RequestsDroppedValue = "NA"
		}

		if readRegister(ntpCore.BaseAddrLReg+ntpServer.CountBroadcastReg, &tempData) == 0 {
			ntpServer.BroadcastsValue = fmt.Sprintf("%d", tempData)

		} else {
			ntpServer.BroadcastsValue = "NA"
		}

		if readRegister(ntpCore.BaseAddrLReg+ntpServer.CountControlReg, &tempData) == 0 {
			if (tempData & 0x00000001) == 0 {
				ntpServer.ClearCounters = false
			} else {
				ntpServer.ClearCounters = true

			}
		} else {
			ntpServer.ClearCounters = false

		}
	}

	func readNtpServerVersion() {
		ntpCore := getNtpCore()
		var tempData int64 = 0x00000000
		// version
		if readRegister(ntpCore.BaseAddrLReg+ntpServer.VersionReg, &tempData) == 0 {
			ntpServer.VersionValue = fmt.Sprintf("0x%02x", tempData) // base 16 string format
		} else {
			ntpServer.VersionValue = "NA"
		}

}
*/
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
