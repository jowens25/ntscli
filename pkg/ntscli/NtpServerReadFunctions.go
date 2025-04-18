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
	if readRegister(NtpServerCore.BaseAddrLReg+ntpServer["ControlReg"], &tempData) == 0 {
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
func readNtpServerInstanceNumber() string {
	return string(NtpServerCore.InstanceNumber)
}

// read NtpServer IP ADDRESS
func readNtpServerIpAddress() string {
	tempData = 0x00000000
	ipAddr := ""

	// ip
	if readNtpServerIpMode() == "IPv4" {
		tempAddr := make([]byte, 4)
		temp_ip4 := make([]int64, 4)
		if readRegister(NtpServerCore.BaseAddrLReg+ntpServer["ConfigIpReg"], &tempData) == 0 {
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
		if readRegister(NtpServerCore.BaseAddrLReg+ntpServer["ConfigIpReg"], &tempData) == 0 {
			temp_ip6[0] = (tempData >> 0) & 0x000000FF
			temp_ip6[1] = (tempData >> 8) & 0x000000FF
			temp_ip6[2] = (tempData >> 16) & 0x000000FF
			temp_ip6[3] = (tempData >> 24) & 0x000000FF

			if readRegister(NtpServerCore.BaseAddrLReg+ntpServer["ConfigIpv61Reg"], &tempData) == 0 {
				temp_ip6[4] = (tempData >> 0) & 0x000000FF
				temp_ip6[5] = (tempData >> 8) & 0x000000FF
				temp_ip6[6] = (tempData >> 16) & 0x000000FF
				temp_ip6[7] = (tempData >> 24) & 0x000000FF

				if readRegister(NtpServerCore.BaseAddrLReg+ntpServer["ConfigIpv62Reg"], &tempData) == 0 {
					temp_ip6[8] = (tempData >> 0) & 0x000000FF
					temp_ip6[9] = (tempData >> 8) & 0x000000FF
					temp_ip6[10] = (tempData >> 16) & 0x000000FF
					temp_ip6[11] = (tempData >> 24) & 0x000000FF

					if readRegister(NtpServerCore.BaseAddrLReg+ntpServer["ConfigIpv63Reg"], &tempData) == 0 {

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
	if readRegister(NtpServerCore.BaseAddrLReg+ntpServer["ConfigModeReg"], &tempData) == 0 {

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

	if readRegister(NtpServerCore.BaseAddrLReg+ntpServer["ConfigMac1Reg"], &tempData) == 0 {

		this_string = append(this_string, fmt.Sprintf("%02x", (tempData>>0)&0x000000FF)...)
		this_string = append(this_string, ':')

		this_string = append(this_string, fmt.Sprintf("%02x", (tempData>>8)&0x000000FF)...)
		this_string = append(this_string, ':')
		this_string = append(this_string, fmt.Sprintf("%02x", (tempData>>16)&0x000000FF)...)
		this_string = append(this_string, ':')
		this_string = append(this_string, fmt.Sprintf("%02x", (tempData>>24)&0x000000FF)...)
		this_string = append(this_string, ':')

		if readRegister(NtpServerCore.BaseAddrLReg+ntpServer["ConfigMac2Reg"], &tempData) == 0 {
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
	if readRegister(NtpServerCore.BaseAddrLReg+ntpServer["ConfigVlanReg"], &tempData) == 0 {
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
	if readRegister(NtpServerCore.BaseAddrLReg+ntpServer["ConfigVlanReg"], &tempData) == 0 {
		tempData &= 0x0000FFFF
		vlanValue = fmt.Sprintf("0x%04x", tempData)
	} else {
		vlanValue = "NA"
	}

	return vlanValue
}

// read VLAN VALUE
func readVlanValue(addr int64) string {
	tempData = 0x00000000
	vlanValue := ""
	if readRegister(addr, &tempData) == 0 {

		//if readRegister(NtpServerCore.BaseAddrLReg+ntpServer["ConfigVlanReg"], &tempData) == 0 {
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
	if readRegister(NtpServerCore.BaseAddrLReg+ntpServer["ConfigModeReg"], &tempData) == 0 {
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
	if readRegister(NtpServerCore.BaseAddrLReg+ntpServer["ConfigModeReg"], &tempData) == 0 {
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
	if readRegister(NtpServerCore.BaseAddrLReg+ntpServer["ConfigModeReg"], &tempData) == 0 {
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
func readNtpServerPrecisionValue() string {
	tempData = 0x00000000
	var PrecisionValue string
	if readRegister(NtpServerCore.BaseAddrLReg+ntpServer["ConfigModeReg"], &tempData) == 0 {
		PrecisionValue = string(rune(int8(((tempData >> 8) & 0x000000FF))))

	} else {
		PrecisionValue = "NA"
	}
	return PrecisionValue
}

// read NtpServer POLL INTERVAL
func readNtpServerPollIntervalValue() string {
	tempData = 0x00000000
	PollIntervalValue := ""
	if readRegister(NtpServerCore.BaseAddrLReg+ntpServer["ConfigModeReg"], &tempData) == 0 {
		PollIntervalValue = strconv.FormatInt(((tempData >> 16) & 0x000000FF), 10) // Base 10
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
	if readRegister(NtpServerCore.BaseAddrLReg+ntpServer["ConfigModeReg"], &tempData) == 0 {
		StratumValue = strconv.FormatInt(((tempData >> 24) & 0x000000FF), 10) // Base 10
	} else {
		StratumValue = "NA"
	}
	return StratumValue
}

// read NtpServer REFERENCEID
func readNtpServerReferenceId() string {
	tempData = 0x00000000
	referenceId := ""
	// reference id
	if readRegister(NtpServerCore.BaseAddrLReg+ntpServer["ConfigReferenceIdReg"], &tempData) == 0 {
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

	if writeRegister(NtpServerCore.BaseAddrLReg+ntpServer["UtcInfoControlReg"], &tempData) == 0 {
		for i := range 10 {
			if readRegister(NtpServerCore.BaseAddrLReg+ntpServer["UtcInfoControlReg"], &tempData) == 0 {
				if (tempData & 0x80000000) != 0 {
					if readRegister(NtpServerCore.BaseAddrLReg+ntpServer["UtcInfoReg"], &tempData) == 0 {
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

	if writeRegister(NtpServerCore.BaseAddrLReg+ntpServer["UtcInfoControlReg"], &tempData) == 0 {
		for i := range 10 {
			if readRegister(NtpServerCore.BaseAddrLReg+ntpServer["UtcInfoControlReg"], &tempData) == 0 {
				if (tempData & 0x80000000) != 0 {
					if readRegister(NtpServerCore.BaseAddrLReg+ntpServer["UtcInfoReg"], &tempData) == 0 {
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

	if writeRegister(NtpServerCore.BaseAddrLReg+ntpServer["UtcInfoControlReg"], &tempData) == 0 {
		for i := range 10 {
			if readRegister(NtpServerCore.BaseAddrLReg+ntpServer["UtcInfoControlReg"], &tempData) == 0 {
				if (tempData & 0x80000000) != 0 {
					if readRegister(NtpServerCore.BaseAddrLReg+ntpServer["UtcInfoReg"], &tempData) == 0 {
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

	if writeRegister(NtpServerCore.BaseAddrLReg+ntpServer["UtcInfoControlReg"], &tempData) == 0 {
		for i := range 10 {
			if readRegister(NtpServerCore.BaseAddrLReg+ntpServer["UtcInfoControlReg"], &tempData) == 0 {
				if (tempData & 0x80000000) != 0 {
					if readRegister(NtpServerCore.BaseAddrLReg+ntpServer["UtcInfoReg"], &tempData) == 0 {
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

	if writeRegister(NtpServerCore.BaseAddrLReg+ntpServer["UtcInfoControlReg"], &tempData) == 0 {
		for i := range 10 {
			if readRegister(NtpServerCore.BaseAddrLReg+ntpServer["UtcInfoControlReg"], &tempData) == 0 {
				if (tempData & 0x80000000) != 0 {
					if readRegister(NtpServerCore.BaseAddrLReg+ntpServer["UtcInfoReg"], &tempData) == 0 {
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

	if writeRegister(NtpServerCore.BaseAddrLReg+ntpServer["UtcInfoControlReg"], &tempData) == 0 {
		for i := range 10 {
			if readRegister(NtpServerCore.BaseAddrLReg+ntpServer["UtcInfoControlReg"], &tempData) == 0 {
				if (tempData & 0x80000000) != 0 {
					if readRegister(NtpServerCore.BaseAddrLReg+ntpServer["UtcInfoReg"], &tempData) == 0 {
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

	if writeRegister(NtpServerCore.BaseAddrLReg+ntpServer["UtcInfoControlReg"], &tempData) == 0 {
		for i := range 10 {
			if readRegister(NtpServerCore.BaseAddrLReg+ntpServer["UtcInfoControlReg"], &tempData) == 0 {
				if (tempData & 0x80000000) != 0 {
					if readRegister(NtpServerCore.BaseAddrLReg+ntpServer["UtcInfoReg"], &tempData) == 0 {

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
	if readRegister(NtpServerCore.BaseAddrLReg+ntpServer["CountReqReg"], &tempData) == 0 {
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
	if readRegister(NtpServerCore.BaseAddrLReg+ntpServer["CountRespReg"], &tempData) == 0 {
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
	if readRegister(NtpServerCore.BaseAddrLReg+ntpServer["CountReqDroppedReg"], &tempData) == 0 {
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
	if readRegister(NtpServerCore.BaseAddrLReg+ntpServer["CountBroadcastReg"], &tempData) == 0 {
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
	if readRegister(NtpServerCore.BaseAddrLReg+ntpServer["CountControlReg"], &tempData) == 0 {
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
	if readRegister(NtpServerCore.BaseAddrLReg+ntpServer["VersionReg"], &tempData) == 0 {
		version = fmt.Sprintf("0x%02x", tempData) // base 16 string format
	} else {
		version = "NA"
	}

	return version

}

func formatNtpServerSTATUS() string {
	return "NTP SERVER STATUS:                     " + readNtpServerStatus()
}
func formatNtpServerINSTANCE() string {
	return "NTP SERVER INSTANCE:                   " + readNtpServerInstanceNumber()
}
func formatNtpServerIPADDRESS() string {
	return "NTP SERVER IP ADDRESS:                 " + readNtpServerIpAddress()
}
func formatNtpServerIPMODE() string {
	return "NTP SERVER IP MODE:                    " + readNtpServerIpMode()
}
func formatNtpServerMACADDRESS() string {
	return "NTP SERVER MAC ADDRESS:                " + readNtpServerMac()
}
func formatNtpServerVLANENABLED() string {
	return "NTP SERVER VLAN ENABLED:               " + readNtpServerVlanEnable()
}
func formatNtpServerVLANVALUE() string {
	return "NTP SERVER VLAN VALUE:                 " + readNtpServerVlanValue()
}
func formatNtpServerUNICAST() string {
	return "NTP SERVER UNICAST:                    " + readNtpServerUnicastMode()
}
func formatNtpServerMULTICAST() string {
	return "NTP SERVER MULTICAST:                  " + readNtpServerMulticastMode()
}
func formatNtpServerBROADCAST() string {
	return "NTP SERVER BROADCAST:                  " + readNtpServerBroadcastMode()
}
func formatNtpServerPRECISION() string {
	return "NTP SERVER PRECISION:                  " + readNtpServerPrecisionValue()
}
func formatNtpServerPOLLINTERVAL() string {
	return "NTP SERVER POLL INTERVAL:              " + readNtpServerPollIntervalValue()
}
func formatNtpServerSTRATUM() string {
	return "NTP SERVER STRATUM:                    " + readNtpServerStratumValue()
}
func formatNtpServerREFERENCEID() string {
	return "NTP SERVER REFERENCE ID:               " + readNtpServerReferenceId()
}
func formatNtpServerUTCSMEARING() string {
	return "NTP SERVER UTC SMEARING:               " + readNtpServerUTCSmearing()
}
func formatNtpServerUTCLEAP61INPROGRESS() string {
	return "NTP SERVER UTC LEAP61 IN PROGRESS:     " + readNtpServerUTCLeap61InProgress()
}
func formatNtpServerUTCLEAP59INPROGRESS() string {
	return "NTP SERVER UTC LEAP59 IN PROGRESS:     " + readNtpServerUTCLeap59InProgress()
}
func formatNtpServerUTCLEAP61() string {
	return "NTP SERVER UTC LEAP 61:                " + readNtpServerUTCLeap61()
}
func formatNtpServerUTCLEAP59() string {
	return "NTP SERVER UTC LEAP 59:                " + readNtpServerUTCLeap59()
}
func formatNtpServerUTCOFFSETENABLE() string {
	return "NTP SERVER UTC OFFSET ENABLE:          " + readNtpServerUTCOffsetEnable()
}
func formatNtpServerUTCOFFSETVALUE() string {
	return "NTP SERVER UTC OFFSET VALUE:           " + readNtpServerUTCOffsetValue()
}
func formatNtpServerREQUESTCOUNT() string {
	return "NTP SERVER REQUEST COUNT:              " + readNtpServerRequestCount()
}
func formatNtpServerRESPONSECOUNT() string {
	return "NTP SERVER RESPONSE COUNT:             " + readNtpServerResponseCount()
}
func formatNtpServerREQUESTSDROPPED() string {
	return "NTP SERVER REQUESTS DROPPED:           " + readNtpServerRequestsDropped()
}
func formatNtpServerBROADCASTCOUNT() string {
	return "NTP SERVER BROADCAST COUNT:            " + readNtpServerBroadcastCount()
}
func formatNtpServerCOUNTCONTROL() string {
	return "NTP SERVER COUNT CONTROL:              " + readNtpServerCountControl()
}
func formatNtpServerVERSION() string {
	return "NTP SERVER VERSION:                    " + readNtpServerVersion()
}

func showNtpServerSTATUS() {
	fmt.Println(formatNtpServerSTATUS())
}
func showNtpServerINSTANCE() {
	fmt.Println(formatNtpServerINSTANCE())
}
func showNtpServerIPADDRESS() {
	fmt.Println(formatNtpServerIPADDRESS())
}
func showNtpServerIPMODE() {
	fmt.Println(formatNtpServerIPMODE())
}
func showNtpServerMACADDRESS() {
	fmt.Println(formatNtpServerMACADDRESS())
}
func showNtpServerVLANENABLED() {
	fmt.Println(formatNtpServerVLANENABLED())
}
func showNtpServerVLANVALUE() {
	fmt.Println(formatNtpServerVLANVALUE())
}
func showNtpServerUNICAST() {
	fmt.Println(formatNtpServerUNICAST())
}
func showNtpServerMULTICAST() {
	fmt.Println(formatNtpServerMULTICAST())
}
func showNtpServerBROADCAST() {
	fmt.Println(formatNtpServerBROADCAST())
}
func showNtpServerPRECISION() {
	fmt.Println(formatNtpServerPRECISION())
}
func showNtpServerPOLLINTERVAL() {
	fmt.Println(formatNtpServerPOLLINTERVAL())
}
func showNtpServerSTRATUM() {
	fmt.Println(formatNtpServerSTRATUM())
}
func showNtpServerREFERENCEID() {
	fmt.Println(formatNtpServerREFERENCEID())
}
func showNtpServerUTCSMEARING() {
	fmt.Println(formatNtpServerUTCSMEARING())
}
func showNtpServerUTCLEAP61INPROGRESS() {
	fmt.Println(formatNtpServerUTCLEAP61INPROGRESS())
}
func showNtpServerUTCLEAP59INPROGRESS() {
	fmt.Println(formatNtpServerUTCLEAP59INPROGRESS())
}
func showNtpServerUTCLEAP61() {
	fmt.Println(formatNtpServerUTCLEAP61())
}
func showNtpServerUTCLEAP59() {
	fmt.Println(formatNtpServerUTCLEAP59())
}
func showNtpServerUTCOFFSETENABLE() {
	fmt.Println(formatNtpServerUTCOFFSETENABLE())
}
func showNtpServerUTCOFFSETVALUE() {
	fmt.Println(formatNtpServerUTCOFFSETVALUE())
}
func showNtpServerREQUESTCOUNT() {
	fmt.Println(formatNtpServerREQUESTCOUNT())
}
func showNtpServerRESPONSECOUNT() {
	fmt.Println(formatNtpServerRESPONSECOUNT())
}
func showNtpServerREQUESTSDROPPED() {
	fmt.Println(formatNtpServerREQUESTSDROPPED())
}
func showNtpServerBROADCASTCOUNT() {
	fmt.Println(formatNtpServerBROADCASTCOUNT())
}
func showNtpServerCOUNTCONTROL() {
	fmt.Println(formatNtpServerCOUNTCONTROL())
}
func showNtpServerVERSION() {
	fmt.Println(formatNtpServerVERSION())
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

func addNtpPropertyComment(prop string) string {
	var val string
	switch prop {
	case "ControlReg":
		val = "--"
	case "StatusReg":
		val = "--" + formatNtpServerSTATUS() + "\n"
	case "VersionReg":
		val = "--" + formatNtpServerVERSION() + "\n"
	case "CountControlReg":
		val = "--"
	case "CountReqReg":
		val = "--" + formatNtpServerREQUESTCOUNT() + "\n"

	case "CountRespReg":
		val = "--" + formatNtpServerRESPONSECOUNT() + "\n"
	case "CountReqDroppedReg":
		val = "--" + formatNtpServerREQUESTSDROPPED() + "\n"
	case "CountBroadcastReg":
		val = "--" + formatNtpServerBROADCASTCOUNT() + "\n"
	case "ConfigControlReg":
		val = "--"
	case "ConfigModeReg":
		val = "--"

	case "ConfigVlanReg":
		val =
			"--" + formatNtpServerVLANENABLED() + "\n" +
				"--" + formatNtpServerVLANVALUE() + "\n"

	case "ConfigMac1Reg":
		val = "--" + formatNtpServerMACADDRESS() + "\n"
	case "ConfigMac2Reg":
		val = ""

	case "ConfigIpReg":
		val =
			"--" + formatNtpServerIPMODE() + "\n" +
				"--" + formatNtpServerIPADDRESS() + "\n"
	case "ConfigIpv61Reg":
	case "ConfigIpv62Reg":
	case "ConfigIpv63Reg":
	case "ConfigReferenceIdReg":

	case "UtcInfoControlReg":
		val = ""

	case "UtcInfoReg":
		val = formatNtpServerUTCOFFSETENABLE() + "\n" +
			"--" + formatNtpServerUTCOFFSETVALUE() + "\n" +
			"--" + formatNtpServerUTCLEAP59() + "\n" +
			"--" + formatNtpServerUTCLEAP59INPROGRESS() + "\n" +
			"--" + formatNtpServerUTCLEAP61() + "\n" +
			"--" + formatNtpServerUTCLEAP61INPROGRESS() + "\n" +
			"--" + formatNtpServerUTCSMEARING() + "\n"
	default:
		val = ""
	}

	return val

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
