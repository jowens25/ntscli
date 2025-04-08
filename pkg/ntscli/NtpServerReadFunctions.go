package ntscli

import (
	"fmt"
	"log"
	"strconv"
)

func NtpReadPrintAll() {

	fmt.Println("NTP SERVER STATUS:                     ", readNtpServerEnable())
	fmt.Println("NTP SERVER INSTANCE:                   ", NtpCore.InstanceNumber)
	fmt.Println("NTP SERVER IP ADDRESS:                 ", readNtpServerIpAddress())
	fmt.Println("NTP SERVER IP MODE:                    ", readNtpServerIpMode())
	fmt.Println("NTP SERVER MAC ADDRESS:                ", readNtpServerMac())
	fmt.Println("NTP SERVER VLAN ENABLED:               ", readNtpServerVlanEnable())
	fmt.Println("NTP SERVER VLAN VALUE:                 ", readNtpServerVlanValue())
	fmt.Println("NTP SERVER UNICAST:                    ", readNtpServerUnicastMode())
	fmt.Println("NTP SERVER MULTICAST:                  ", readNtpServerMulticastMode())
	fmt.Println("NTP SERVER BROADCAST:                  ", readNtpServerBroadcastMode())
	fmt.Println("NTP SERVER PRECISION:                  ", readNtpServerPrecisionValue())
	fmt.Println("NTP SERVER POLL INTERVAL:              ", readNtpServerPollIntervalValue())
	fmt.Println("NTP SERVER STRATUM:                    ", readNtpServerStratumValue())
	fmt.Println("NTP SERVER REFERENCE ID:               ", readNtpServerReferenceId())

	fmt.Println("NTP SERVER UTC SMEARING:               ", readNtpServerUTCSmearing())
	fmt.Println("NTP SERVER UTC LEAP61 IN PROGRESS:     ", readNtpServerUTCLeap61InProgress())
	fmt.Println("NTP SERVER UTC LEAP59 IN PROGRESS:     ", readNtpServerUTCLeap59InProgress())
	fmt.Println("NTP SERVER UTC LEAP 61:                ", readNtpServerUTCLeap61())
	fmt.Println("NTP SERVER UTC LEAP 59:                ", readNtpServerUTCLeap59())
	fmt.Println("NTP SERVER UTC OFFSET ENABLE:          ", readNtpServerUTCOffsetEnable())
	fmt.Println("NTP SERVER UTC OFFSET VALUE:           ", readNtpServerUTCOffsetValue())

	fmt.Println("NTP SERVER REQUEST COUNT:              ", readNtpServerRequestCount())
	fmt.Println("NTP SERVER RESPONSE COUNT:             ", readNtpServerResponseCount())
	fmt.Println("NTP SERVER REQUESTS DROPPED:           ", readNtpServerRequestsDropped())
	fmt.Println("NTP SERVER BROADCAST COUNT:            ", readNtpServerBroadcastCount())
	fmt.Println("NTP SERVER COUNT CONTROL:              ", readNtpServerCountControl())
	fmt.Println("NTP SERVER VERSION:                    ", readNtpServerVersion())

	//readNtpServerMode()
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

func readNtpServerUTCSmearing() string {

	// utc info
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

func readNtpServerUTCLeap61InProgress() string {

	// utc info
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
