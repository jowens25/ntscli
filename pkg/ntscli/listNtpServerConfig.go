package ntscli

import (
	"fmt"
	"log"
	"strconv"
)

func listNtpServerConfig(core *CoreConfig) {

	fmt.Println("list ntp server config: ", *core)
	fmt.Println("core: ", core.BaseAddrHReg)
	var temp_data int64 = 0x00000000
	// enabled
	if readRegister(core.BaseAddrLReg+ntpServer.ControlReg, &temp_data) == 0 {
		if (temp_data & 0x00000001) == 0 {
			ntpServer.Enabled = false
		} else {
			ntpServer.Enabled = true
		}
	} else {
		ntpServer.Enabled = false
		//fmt.Println("NtpServerEnabled: False!")
	}

	// mac
	this_string := make([]byte, 0, 32)

	if readRegister(core.BaseAddrLReg+ntpServer.ConfigMac1Reg, &temp_data) == 0 {

		this_string = append(this_string, fmt.Sprintf("%02x", (temp_data>>0)&0x000000FF)...)
		this_string = append(this_string, ':')

		this_string = append(this_string, fmt.Sprintf("%02x", (temp_data>>8)&0x000000FF)...)
		this_string = append(this_string, ':')
		this_string = append(this_string, fmt.Sprintf("%02x", (temp_data>>16)&0x000000FF)...)
		this_string = append(this_string, ':')
		this_string = append(this_string, fmt.Sprintf("%02x", (temp_data>>24)&0x000000FF)...)
		this_string = append(this_string, ':')

		if readRegister(core.BaseAddrLReg+ntpServer.ConfigMac2Reg, &temp_data) == 0 {
			this_string = append(this_string, fmt.Sprintf("%02x", (temp_data>>0)&0x000000FF)...)
			this_string = append(this_string, ':')

			this_string = append(this_string, fmt.Sprintf("%02x", (temp_data>>8)&0x000000FF)...)

			ntpServer.MacAddr = string(this_string)
			//fmt.Println("NtpServerMacValue: ", string(this_string))
		} else {
			//fmt.Println("NtpServerMacValue: NA")
			ntpServer.MacAddr = "NA"
		}

	} else {
		//fmt.Println("NtpServerMacValue: NA")
		ntpServer.MacAddr = "NA"
	}

	// vlan
	if readRegister(core.BaseAddrLReg+ntpServer.ConfigVlanReg, &temp_data) == 0 {
		if (temp_data & 0x00010000) == 0 {
			//fmt.Println("NtpServerVlanEnable: False")
			ntpServer.VlanEnable = false
		} else {
			ntpServer.VlanEnable = true
			//fmt.Println("NtpServerVlanEnable: True")
		}

		temp_data &= 0x0000FFFF
		//fmt.Println("NtpServerVlanValue: ", fmt.Sprintf("0x%08x",&temp_data))
		ntpServer.VlanValue = fmt.Sprintf("0x%08x", &temp_data)

	} else {
		ntpServer.VlanEnable = false

		///fmt.Println("NtpServerVlanEnable: False")
		ntpServer.VlanValue = "NA"

	}

	// mode & server config
	if readRegister(core.BaseAddrLReg+ntpServer.ConfigModeReg, &temp_data) == 0 {
		if ((temp_data >> 0) & 0x00000003) == 1 {
			ntpServer.IpMode = "IPv4"
		} else if ((temp_data >> 0) & 0x00000003) == 2 {
			ntpServer.IpMode = "IPv6"
		} else {
			ntpServer.IpMode = "NA"
		}

		if (temp_data & 0x00000010) == 0 {
			ntpServer.UnicastMode = false
		} else {
			ntpServer.UnicastMode = true
		}

		if (temp_data & 0x00000020) == 0 {
			ntpServer.MulticastMode = false
		} else {
			ntpServer.MulticastMode = true
		}

		if (temp_data & 0x00000040) == 0 {
			ntpServer.BroadcastMode = false
		} else {
			ntpServer.BroadcastMode = true
		}

		ntpServer.PrecisionValue = rune(int8(((temp_data >> 8) & 0x000000FF)))
		ntpServer.PollIntervalValue = string(((temp_data >> 16) & 0x000000FF))
		ntpServer.StratumValue = string((temp_data >> 24) & 0x000000FF)

	} else {
		ntpServer.IpMode = "NA"
		ntpServer.UnicastMode = false
		ntpServer.MulticastMode = false
		ntpServer.BroadcastMode = false

		ntpServer.PrecisionValue = 'N'
		ntpServer.PollIntervalValue = "NA"
		ntpServer.StratumValue = "NA"

	}

	// reference id // no ref on UI??
	if readRegister(core.BaseAddrLReg+ntpServer.ConfigReferenceIdReg, &temp_data) == 0 {
		var temp_string []byte
		temp_string = append(temp_string, byte(((temp_data >> 24) & 0x000000FF)))
		temp_string = append(temp_string, byte(((temp_data >> 16) & 0x000000FF)))
		temp_string = append(temp_string, byte(((temp_data >> 8) & 0x000000FF)))
		temp_string = append(temp_string, byte(((temp_data >> 0) & 0x000000FF)))
		ntpServer.ReferenceId = fmt.Sprintf("0x%08x", temp_string) // TODO
	} else {
		ntpServer.ReferenceId = "BANANA"
	}

	//log.Println(ntpServer.IpMode)
	//log.Println(ntpServer.IpMode == "IPv4")
	// ip
	if ntpServer.IpMode == "IPv4" {
		if readRegister(core.BaseAddrLReg+ntpServer.ConfigIpReg, &temp_data) == 0 {
			var temp_ip int64 = 0x00000000
			temp_ip |= (temp_data >> 0) & 0x000000FF
			temp_ip = temp_ip << 8
			temp_ip |= (temp_data >> 8) & 0x000000FF
			temp_ip = temp_ip << 8
			temp_ip |= (temp_data >> 16) & 0x000000FF
			temp_ip = temp_ip << 8
			temp_ip |= (temp_data >> 24) & 0x000000FF

			//log.Println(temp_ip)
			//log.Println(int_to_ip_addr(temp_ip))
			//temp_string := string(temp_ip)
			ntpServer.IpAddr = int_to_ip_addr(temp_ip)

		} else {
			ntpServer.IpAddr = "NA"
		}
	} else if ntpServer.IpMode == "IPv6" {
		temp_ip6 := make([]int64, 16)
		if readRegister(core.BaseAddrLReg+ntpServer.ConfigIpReg, &temp_data) == 0 {
			temp_ip6[0] = (temp_data >> 0) & 0x000000FF
			temp_ip6[1] = (temp_data >> 8) & 0x000000FF
			temp_ip6[2] = (temp_data >> 16) & 0x000000FF
			temp_ip6[3] = (temp_data >> 24) & 0x000000FF

			if readRegister(core.BaseAddrLReg+ntpServer.ConfigIpv61Reg, &temp_data) == 0 {
				temp_ip6[4] = (temp_data >> 0) & 0x000000FF
				temp_ip6[5] = (temp_data >> 8) & 0x000000FF
				temp_ip6[6] = (temp_data >> 16) & 0x000000FF
				temp_ip6[7] = (temp_data >> 24) & 0x000000FF

				if readRegister(core.BaseAddrLReg+ntpServer.ConfigIpv62Reg, &temp_data) == 0 {
					temp_ip6[8] = (temp_data >> 0) & 0x000000FF
					temp_ip6[9] = (temp_data >> 8) & 0x000000FF
					temp_ip6[10] = (temp_data >> 16) & 0x000000FF
					temp_ip6[11] = (temp_data >> 24) & 0x000000FF

					if readRegister(core.BaseAddrLReg+ntpServer.ConfigIpv63Reg, &temp_data) == 0 {

						temp_ip6[12] = (temp_data >> 0) & 0x000000FF
						temp_ip6[13] = (temp_data >> 8) & 0x000000FF
						temp_ip6[14] = (temp_data >> 16) & 0x000000FF
						temp_ip6[15] = (temp_data >> 24) & 0x000000FF

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

	// utc info
	temp_data = 0x40000000

	if writeRegister(core.BaseAddrLReg+ntpServer.UtcInfoControlReg, &temp_data) == 0 {
		for i := range 10 {
			if readRegister(core.BaseAddrLReg+ntpServer.UtcInfoControlReg, &temp_data) == 0 {
				if (temp_data & 0x80000000) != 0 {
					if readRegister(core.BaseAddrLReg+ntpServer.UtcInfoReg, &temp_data) == 0 {
						if (temp_data & 0x00000100) == 0 {
							ntpServer.UtcSmearing = false
						} else {
							ntpServer.UtcSmearing = true

						}

						if (temp_data & 0x00000200) == 0 {
							ntpServer.UtcLeap61InProgress = false
						} else {
							ntpServer.UtcLeap61InProgress = true
						}

						if (temp_data & 0x00000400) == 0 {
							ntpServer.UtcLeap59InProgress = false
						} else {
							ntpServer.UtcLeap59InProgress = true
						}

						if (temp_data & 0x00000800) == 0 {
							ntpServer.UtcLeap61 = false
						} else {
							ntpServer.UtcLeap61 = true
						}

						if (temp_data & 0x00001000) == 0 {
							ntpServer.UtcLeap59 = false
						} else {
							ntpServer.UtcLeap59 = true
						}

						if (temp_data & 0x00002000) == 0 {
							ntpServer.UtcOffsetVal = false
						} else {
							ntpServer.UtcOffsetVal = true
						}

						//log.Println("ui->NtpServerUtcOffsetValue->setText(QString::number(((temp_data >> 16) & 0x0000FFFF)));")

						ntpServer.UtcOffsetValue = (temp_data >> 16) & 0x0000FFFF

						//ntpServer.UtcOffsetValue = string((temp_data >> 16) & 0x0000FFFF)
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

	// status
	if readRegister(core.BaseAddrLReg+ntpServer.CountReqReg, &temp_data) == 0 {
		ntpServer.RequestsValue = fmt.Sprintf("%d", temp_data)
	} else {
		ntpServer.RequestsValue = "NA"
	}

	if readRegister(core.BaseAddrLReg+ntpServer.CountRespReg, &temp_data) == 0 {
		ntpServer.ResponsesValue = fmt.Sprintf("%d", temp_data)

	} else {
		ntpServer.ResponsesValue = "NA"
	}

	if readRegister(core.BaseAddrLReg+ntpServer.CountReqDroppedReg, &temp_data) == 0 {
		ntpServer.RequestsDroppedValue = fmt.Sprintf("%d", temp_data)
	} else {
		ntpServer.RequestsDroppedValue = "NA"
	}

	if readRegister(core.BaseAddrLReg+ntpServer.CountBroadcastReg, &temp_data) == 0 {
		ntpServer.BroadcastsValue = fmt.Sprintf("%d", temp_data)

	} else {
		ntpServer.BroadcastsValue = "NA"
	}

	if readRegister(core.BaseAddrLReg+ntpServer.CountControlReg, &temp_data) == 0 {
		if (temp_data & 0x00000001) == 0 {
			ntpServer.ClearCounters = false
		} else {
			ntpServer.ClearCounters = true

		}
	} else {
		ntpServer.ClearCounters = false

	}

	// version
	if readRegister(core.BaseAddrLReg+ntpServer.VersionReg, &temp_data) == 0 {
		ntpServer.VersionValue = fmt.Sprintf("0x%02x", temp_data) // base 16 string format
	} else {
		ntpServer.VersionValue = "NA"
	}

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
