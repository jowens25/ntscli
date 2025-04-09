package ntscli

import (
	"fmt"
	"log"
	"net"
	"strconv"
)

// read NtpServer STATUS
func readNtpServerStatus() string {
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

// read NtpServer INSTANCE
func readNtpServerInstanceNumber() int64 {
	return NtpCore.InstanceNumber
}

// read NtpServer IP ADDRESS
func readNtpServerIpAddress() string {
	tempData = 0x00000000
	ipAddr := ""

	// ip
	if readNtpServerIpMode() == "IPv4" {
		tempAddr := make([]byte, 4)
		temp_ip4 := make([]int64, 4)
		if readRegister(NtpCore.BaseAddrLReg+ntpServer.ConfigIpReg, &tempData) == 0 {
			temp_ip4[0] = (tempData >> 0) & 0x000000FF
			temp_ip4[1] = (tempData >> 8) & 0x000000FF
			temp_ip4[2] = (tempData >> 16) & 0x000000FF
			temp_ip4[3] = (tempData >> 24) & 0x000000FF

			//ipAddr = int_to_ip_addr(temp_ip)
			//fmt.Println(temp_ip4)
			for i, intPart := range temp_ip4 {
				tempAddr[i] = byte(intPart)
			}

			ip := net.IP(tempAddr)

			if ip.To16() != nil {
				ipAddr = ip.String()

			} else {
				ipAddr = "::ffff:" + ip.String()

			}

		} else {
			ipAddr = "NA"
		}
	} else if readNtpServerIpMode() == "IPv6" {
		//fmt.Println("IPV6 read")
		tempAddr := make([]byte, 16)
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

						//log.Println("IPv6 Addr: ", temp_ip6)

						for i, intPart := range temp_ip6 {
							tempAddr[i] = byte(intPart)
						}

						ip := net.IP(tempAddr)

						if ip.To4() != nil {
							ipAddr = "::ffff:" + ip.String()
						} else {
							ipAddr = ip.String()
						}

						//log.Println(ipAddr)
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

// read NtpServer IP MODE
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

// read NtpServer MAC ADDRESS
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

// read NtpServer VLAN ENABLED
func readNtpServerVlanEnable() string {
	tempData = 0x00000000
	vlanMode := ""
	if readRegister(NtpCore.BaseAddrLReg+ntpServer.ConfigVlanReg, &tempData) == 0 {
		if (tempData & 0x00010000) == 0 {
			vlanMode = "DISABLED"
		} else {
			vlanMode = "ENABLED"
		}
	} else {
		vlanMode = "DISABLED"
	}

	return vlanMode
}

// read NtpServer VLAN VALUE
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

// read NtpServer UNICAST
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

// read NtpServer MULTICAST

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

// read NtpServer BROADCAST
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

// read NtpServer PRECISION
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

// read NtpServer POLLINTERVAL
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

// read NtpServer STRATUM
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

// read NtpServer REFERENCEID
func readNtpServerReferenceId() string {
	tempData = 0x00000000
	referenceId := ""
	// reference id // no ref on UI??
	if readRegister(NtpCore.BaseAddrLReg+ntpServer.ConfigReferenceIdReg, &tempData) == 0 {
		fmt.Println("DEBUG: ref id hex: ", tempData)
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

// read NtpServer UTC SMEARING
func readNtpServerUTCSmearing() string {

	tempData = 0x40000000
	utcSmearing := "NA"

	if writeRegister(NtpCore.BaseAddrLReg+ntpServer.UtcInfoControlReg, &tempData) == 0 {
		for i := range 10 {
			if readRegister(NtpCore.BaseAddrLReg+ntpServer.UtcInfoControlReg, &tempData) == 0 {
				if (tempData & 0x80000000) != 0 {
					if readRegister(NtpCore.BaseAddrLReg+ntpServer.UtcInfoReg, &tempData) == 0 {
						if (tempData & 0x00000100) == 0 {
							utcSmearing = "FALSE"
						} else {
							utcSmearing = "TRUE"
						}
					} else {
						utcSmearing = "FALSE"
					}
					break
				} else if i == 9 {
					utcSmearing = "FALSE"
					log.Fatal("utc read incomplete")
				}

			} else {
				utcSmearing = "FALSE"
			}
		}
	} else {
		utcSmearing = "FALSE"
	}

	return utcSmearing
}

// read NtpServer UTC LEAP61 INPROGRESS
func readNtpServerUTCLeap61InProgress() string {
	tempData = 0x40000000
	utcLeap61InProgress := "NA"

	if writeRegister(NtpCore.BaseAddrLReg+ntpServer.UtcInfoControlReg, &tempData) == 0 {
		for i := range 10 {
			if readRegister(NtpCore.BaseAddrLReg+ntpServer.UtcInfoControlReg, &tempData) == 0 {
				if (tempData & 0x80000000) != 0 {
					if readRegister(NtpCore.BaseAddrLReg+ntpServer.UtcInfoReg, &tempData) == 0 {
						if (tempData & 0x00000200) == 0 {
							utcLeap61InProgress = "FALSE"
						} else {
							utcLeap61InProgress = "TRUE"
						}
					} else {
						utcLeap61InProgress = "FALSE"
					}
					break
				} else if i == 9 {
					utcLeap61InProgress = "FALSE"
					log.Fatal("utc read incomplete")
				}

			} else {
				utcLeap61InProgress = "FALSE"
			}
		}
	} else {
		utcLeap61InProgress = "FALSE"
	}

	return utcLeap61InProgress
}

// read NtpServer UTC LEAP59 INPROGRESS
func readNtpServerUTCLeap59InProgress() string {

	tempData = 0x40000000
	utcLeap59InProgress := "NA"

	if writeRegister(NtpCore.BaseAddrLReg+ntpServer.UtcInfoControlReg, &tempData) == 0 {
		for i := range 10 {
			if readRegister(NtpCore.BaseAddrLReg+ntpServer.UtcInfoControlReg, &tempData) == 0 {
				if (tempData & 0x80000000) != 0 {
					if readRegister(NtpCore.BaseAddrLReg+ntpServer.UtcInfoReg, &tempData) == 0 {
						if (tempData & 0x00000400) == 0 {
							utcLeap59InProgress = "FALSE"
						} else {
							utcLeap59InProgress = "TRUE"
						}
					} else {
						utcLeap59InProgress = "FALSE"
					}
					break
				} else if i == 9 {
					utcLeap59InProgress = "FALSE"
					log.Fatal("utc read incomplete")
				}

			} else {
				utcLeap59InProgress = "FALSE"
			}
		}
	} else {
		utcLeap59InProgress = "FALSE"
	}

	return utcLeap59InProgress
}

// read NtpServer UTC LEAP61
func readNtpServerUTCLeap61() string {

	tempData = 0x40000000
	utcLeap61 := "NA"

	if writeRegister(NtpCore.BaseAddrLReg+ntpServer.UtcInfoControlReg, &tempData) == 0 {
		for i := range 10 {
			if readRegister(NtpCore.BaseAddrLReg+ntpServer.UtcInfoControlReg, &tempData) == 0 {
				if (tempData & 0x80000000) != 0 {
					if readRegister(NtpCore.BaseAddrLReg+ntpServer.UtcInfoReg, &tempData) == 0 {
						if (tempData & 0x00000800) == 0 {
							utcLeap61 = "FALSE"
						} else {
							utcLeap61 = "TRUE"
						}
					} else {
						utcLeap61 = "FALSE"
					}
					break
				} else if i == 9 {
					utcLeap61 = "FALSE"
					log.Fatal("utc read incomplete")
				}

			} else {
				utcLeap61 = "FALSE"
			}
		}
	} else {
		utcLeap61 = "FALSE"
	}

	return utcLeap61
}

// read NtpServer UTC LEAP59
func readNtpServerUTCLeap59() string {

	tempData = 0x40000000
	utcLeap59 := "NA"

	if writeRegister(NtpCore.BaseAddrLReg+ntpServer.UtcInfoControlReg, &tempData) == 0 {
		for i := range 10 {
			if readRegister(NtpCore.BaseAddrLReg+ntpServer.UtcInfoControlReg, &tempData) == 0 {
				if (tempData & 0x80000000) != 0 {
					if readRegister(NtpCore.BaseAddrLReg+ntpServer.UtcInfoReg, &tempData) == 0 {
						if (tempData & 0x00001000) == 0 {
							utcLeap59 = "FALSE"
						} else {
							utcLeap59 = "TRUE"
						}
					} else {
						utcLeap59 = "FALSE"
					}
					break
				} else if i == 9 {
					utcLeap59 = "FALSE"
					log.Fatal("utc read incomplete")
				}

			} else {
				utcLeap59 = "FALSE"
			}
		}
	} else {
		utcLeap59 = "FALSE"
	}

	return utcLeap59
}

// read NtpServer UTC OFFSET ENABLE
func readNtpServerUTCOffsetEnable() string {

	tempData = 0x40000000
	offsetEnable := "NA"

	if writeRegister(NtpCore.BaseAddrLReg+ntpServer.UtcInfoControlReg, &tempData) == 0 {
		for i := range 10 {
			if readRegister(NtpCore.BaseAddrLReg+ntpServer.UtcInfoControlReg, &tempData) == 0 {
				if (tempData & 0x80000000) != 0 {
					if readRegister(NtpCore.BaseAddrLReg+ntpServer.UtcInfoReg, &tempData) == 0 {
						if (tempData & 0x00002000) == 0 {
							offsetEnable = "FALSE"
						} else {
							offsetEnable = "TRUE"
						}
					} else {
						offsetEnable = "FALSE"
					}
					break
				} else if i == 9 {
					offsetEnable = "FALSE"
					log.Fatal("utc read incomplete")
				}

			} else {
				offsetEnable = "FALSE"
			}
		}
	} else {
		offsetEnable = "FALSE"
	}

	return offsetEnable
}

// read NtpServer UTC OFFSET VALUE
func readNtpServerUTCOffsetValue() string {

	tempData = 0x40000000
	utcOffsetValue := "NA"

	if writeRegister(NtpCore.BaseAddrLReg+ntpServer.UtcInfoControlReg, &tempData) == 0 {
		for i := range 10 {
			if readRegister(NtpCore.BaseAddrLReg+ntpServer.UtcInfoControlReg, &tempData) == 0 {
				if (tempData & 0x80000000) != 0 {
					if readRegister(NtpCore.BaseAddrLReg+ntpServer.UtcInfoReg, &tempData) == 0 {

						utcOffsetValue = strconv.FormatInt((tempData>>16)&0x0000FFFF, 10) // Base 10

					} else {
						utcOffsetValue = "0"
					}
					break
				} else if i == 9 {
					utcOffsetValue = "0"
					log.Fatal("utc read incomplete")
				}

			} else {
				utcOffsetValue = "0"
			}
		}
	} else {
		utcOffsetValue = "0"
	}

	return utcOffsetValue
}

// read NtpServer REQUEST COUNT
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

// read NtpServer RESPONSE COUNT
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

// read NtpServer REQUESTS DROPPED
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

// read NtpServer BROADCAST COUNT
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

// read NtpServer COUNT CONTROL
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

// read NtpServer VERSION
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

func showNtpServerSTATUS() {
	fmt.Println("NTP SERVER STATUS:                     ", readNtpServerStatus())
}
func showNtpServerINSTANCE() {
	fmt.Println("NTP SERVER INSTANCE:                   ", readNtpServerInstanceNumber())
}
func showNtpServerIPADDRESS() {
	fmt.Println("NTP SERVER IP ADDRESS:                 ", readNtpServerIpAddress())
}
func showNtpServerIPMODE() {
	fmt.Println("NTP SERVER IP MODE:                    ", readNtpServerIpMode())
}
func showNtpServerMACADDRESS() {
	fmt.Println("NTP SERVER MAC ADDRESS:                ", readNtpServerMac())
}
func showNtpServerVLANENABLED() {
	fmt.Println("NTP SERVER VLAN ENABLED:               ", readNtpServerVlanEnable())
}
func showNtpServerVLANVALUE() {
	fmt.Println("NTP SERVER VLAN VALUE:                 ", readNtpServerVlanValue())
}
func showNtpServerUNICAST() {
	fmt.Println("NTP SERVER UNICAST:                    ", readNtpServerUnicastMode())
}
func showNtpServerMULTICAST() {
	fmt.Println("NTP SERVER MULTICAST:                  ", readNtpServerMulticastMode())
}
func showNtpServerBROADCAST() {
	fmt.Println("NTP SERVER BROADCAST:                  ", readNtpServerBroadcastMode())
}
func showNtpServerPRECISION() {
	fmt.Println("NTP SERVER PRECISION:                  ", readNtpServerPrecisionValue())
}
func showNtpServerPOLLINTERVAL() {
	fmt.Println("NTP SERVER POLL INTERVAL:              ", readNtpServerPollIntervalValue())
}
func showNtpServerSTRATUM() {
	fmt.Println("NTP SERVER STRATUM:                    ", readNtpServerStratumValue())
}
func showNtpServerREFERENCEID() {
	fmt.Println("NTP SERVER REFERENCE ID:               ", readNtpServerReferenceId())
}
func showNtpServerUTCSMEARING() {
	fmt.Println("NTP SERVER UTC SMEARING:               ", readNtpServerUTCSmearing())
}
func showNtpServerUTCLEAP61INPROGRESS() {
	fmt.Println("NTP SERVER UTC LEAP61 IN PROGRESS:     ", readNtpServerUTCLeap61InProgress())
}
func showNtpServerUTCLEAP59INPROGRESS() {
	fmt.Println("NTP SERVER UTC LEAP59 IN PROGRESS:     ", readNtpServerUTCLeap59InProgress())
}
func showNtpServerUTCLEAP61() {
	fmt.Println("NTP SERVER UTC LEAP 61:                ", readNtpServerUTCLeap61())
}
func showNtpServerUTCLEAP59() {
	fmt.Println("NTP SERVER UTC LEAP 59:                ", readNtpServerUTCLeap59())
}
func showNtpServerUTCOFFSETENABLE() {
	fmt.Println("NTP SERVER UTC OFFSET ENABLE:          ", readNtpServerUTCOffsetEnable())
}
func showNtpServerUTCOFFSETVALUE() {
	fmt.Println("NTP SERVER UTC OFFSET VALUE:           ", readNtpServerUTCOffsetValue())
}
func showNtpServerREQUESTCOUNT() {
	fmt.Println("NTP SERVER REQUEST COUNT:              ", readNtpServerRequestCount())
}
func showNtpServerRESPONSECOUNT() {
	fmt.Println("NTP SERVER RESPONSE COUNT:             ", readNtpServerResponseCount())
}
func showNtpServerREQUESTSDROPPED() {
	fmt.Println("NTP SERVER REQUESTS DROPPED:           ", readNtpServerRequestsDropped())
}
func showNtpServerBROADCASTCOUNT() {
	fmt.Println("NTP SERVER BROADCAST COUNT:            ", readNtpServerBroadcastCount())
}
func showNtpServerCOUNTCONTROL() {
	fmt.Println("NTP SERVER COUNT CONTROL:              ", readNtpServerCountControl())
}
func showNtpServerVERSION() {
	fmt.Println("NTP SERVER VERSION:                    ", readNtpServerVersion())
}

func NtpPrintAll() {

	showNtpServerSTATUS()
	showNtpServerINSTANCE()
	showNtpServerIPADDRESS()
	showNtpServerIPMODE()
	showNtpServerMACADDRESS()
	showNtpServerVLANENABLED()
	showNtpServerVLANVALUE()
	showNtpServerUNICAST()
	showNtpServerMULTICAST()
	showNtpServerBROADCAST()
	showNtpServerPRECISION()
	showNtpServerPOLLINTERVAL()
	showNtpServerSTRATUM()
	showNtpServerREFERENCEID()
	showNtpServerUTCSMEARING()
	showNtpServerUTCLEAP61INPROGRESS()
	showNtpServerUTCLEAP59INPROGRESS()
	showNtpServerUTCLEAP61()
	showNtpServerUTCLEAP59()
	showNtpServerUTCOFFSETENABLE()
	showNtpServerUTCOFFSETVALUE()
	showNtpServerREQUESTCOUNT()
	showNtpServerRESPONSECOUNT()
	showNtpServerREQUESTSDROPPED()
	showNtpServerBROADCASTCOUNT()
	showNtpServerCOUNTCONTROL()
	showNtpServerVERSION()
}

// read Ntp Server IP helper functions
//func int_to_ip_addr(val int64) string {
//	fmt.Println("in to ip addr: ", val)
//	hex_string := fmt.Sprintf("%02x", val) // base 16 string format
//	fmt.Println("in to ip addr hex string: ", hex_string)
//
//	return hex_to_decimal(split_into_ip_addr(hex_string))
//}

// func hex_to_decimal(hex_parts []string) string {
//
// 	ip := ""
// 	decimalValue, _ := strconv.ParseInt(hex_parts[0], 16, 16)
// 	ip += fmt.Sprint(decimalValue)
//
// 	for _, part := range hex_parts[1:] {
//
// 		decimalValue, err := strconv.ParseInt(part, 16, 16)
// 		if err != nil {
// 			log.Fatal("IP addr error")
// 		}
// 		ip += "." + fmt.Sprint(decimalValue)
//
// 	}
// 	return ip
// }

//func split_into_ip_addr(hex_string string) []string {
//	var parts = []string{hex_string[0:2], hex_string[2:4], hex_string[4:6], hex_string[6:8]}
//	return parts
//}
//
//func int_to_ipv6(addr []int64) string {
//	return "::ffff:" + fmt.Sprintf("%d", addr[0]) + "." + fmt.Sprintf("%d", addr[1]) + "." + fmt.Sprintf("%d", addr[2]) + "." + fmt.Sprintf("%d", addr[3])
//}
