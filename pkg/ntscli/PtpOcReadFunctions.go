package ntscli

import (
	"fmt"
	"log"
	"net"
)

func readPtpOcVersion() string {

	tempData = 0x00000000
	// version
	version := ""
	if readRegister(PtpOcCore.BaseAddrLReg+ptpOc["VersionReg"], &tempData) == 0 {
		version = fmt.Sprintf("0x%02x", tempData) // base 16 string format
	} else {
		version = "NA"
	}

	return version
}

// read PtpOc STATUS
func readPtpOcStatus() string {
	tempData = 0x00000000
	enabled := ""
	if readRegister(PtpOcCore.BaseAddrLReg+ptpOc["ControlReg"], &tempData) == 0 {
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

// read PtpOc INSTANCE
func readPtpOcInstanceNumber() int64 {
	return PtpOcCore.InstanceNumber
}

// read PtpOc Vlan Enable
func readPtpOcVlanEnable() string {
	tempData = 0x00000000
	vlanMode := ""
	if readRegister(PtpOcCore.BaseAddrLReg+ptpOc["ConfigVlanReg"], &tempData) == 0 {
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

// read ptp oc vlan value
func readPtpOcVlanValue() string {
	tempData = 0x00000000
	vlanValue := ""
	if readRegister(PtpOcCore.BaseAddrLReg+ptpOc["ConfigVlanReg"], &tempData) == 0 {
		tempData &= 0x0000FFFF
		vlanValue = fmt.Sprintf("0x%04x", tempData)
	} else {
		vlanValue = "NA"
	}

	return vlanValue
}

func readPtpOcProfile() string {
	tempData = 0x00000000
	profile := ""
	if readRegister(PtpOcCore.BaseAddrLReg+ptpOc["ConfigProfileReg"], &tempData) == 0 {

		switch tempData & 0x00000007 {
		case 0:
			profile = "Default"
		case 1:
			profile = "Power"
		case 2:
			profile = "Utility"

		case 3:
			profile = "TSN"

		case 4:
			profile = "ITUG8265.1"

		case 5:
			profile = "ITUG8275.1"

		case 6:
			profile = "ITUG8275.2"

		default:
			profile = "NA"

		}
	}
	return profile
}

func readPtpOcDefaultDatasetTwoStep() string {
	tempData = 0x00000000
	twoStep := "DISABLED"
	if readRegister(PtpOcCore.BaseAddrLReg+ptpOc["ConfigProfileReg"], &tempData) == 0 {
		switch (tempData >> 8) & 0x00000001 {
		case 0:
			twoStep = "DISABLED"
		case 1:
			twoStep = "ENABLED"
		default:
			twoStep = "DISABLED"
		}
	}
	return twoStep
}

func readPtpOcDefaultDatasetSignaling() string {
	tempData = 0x00000000
	signaling := "DISABLED"
	if readRegister(PtpOcCore.BaseAddrLReg+ptpOc["ConfigProfileReg"], &tempData) == 0 {
		switch (tempData >> 9) & 0x00000001 {
		case 0:
			signaling = "DISABLED"
		case 1:
			signaling = "ENABLED"
		default:
			signaling = "DISABLED"
		}
	}
	return signaling
}

func readPtpOcLayer() string {
	tempData = 0x00000000
	layer := "NA"
	if readRegister(PtpOcCore.BaseAddrLReg+ptpOc["ConfigProfileReg"], &tempData) == 0 {
		switch (tempData >> 16) & 0x00000003 {
		case 0:
			layer = "Layer 2"
		case 1:
			layer = "Layer 3v4"
		case 2:
			layer = "Layer 3v6"
		default:
			layer = "NA"
		}
	}
	return layer
}

func readPtpOcDefaultDatasetSlaveOnly() string {
	tempData = 0x00000000
	slaveOnly := ""
	if readRegister(PtpOcCore.BaseAddrLReg+ptpOc["ConfigProfileReg"], &tempData) == 0 {
		switch (tempData >> 20) & 0x00000003 {
		case 0:
			slaveOnly = "FALSE"
		case 1:
			slaveOnly = "TRUE"
		case 2:
			slaveOnly = "FALSE"
		default:
			slaveOnly = "FALSE"
		}
	}

	return slaveOnly
}

func readPtpOcDefaultDatasetMasterOnly() string {
	tempData = 0x00000000
	masterOnly := ""
	if readRegister(PtpOcCore.BaseAddrLReg+ptpOc["ConfigProfileReg"], &tempData) == 0 {
		switch (tempData >> 20) & 0x00000003 {
		case 0:
			masterOnly = "FALSE"
		case 1:
			masterOnly = "FALSE"
		case 2:
			masterOnly = "TRUE"
		default:
			masterOnly = "FALSE"
		}
	}
	return masterOnly
}

func readPtpOcDefaultDatasetOffsetCorrectionsEnable() string {
	tempData = 0x00000000
	offsetCorrectEnable := ""
	if readRegister(PtpOcCore.BaseAddrLReg+ptpOc["ConfigProfileReg"], &tempData) == 0 {
		switch (tempData >> 22) & 0x00000001 {
		case 0:
			offsetCorrectEnable = "FALSE"
		case 1:
			offsetCorrectEnable = "TRUE"
		default:
			offsetCorrectEnable = "FALSE"
		}
	}
	return offsetCorrectEnable
}

func readPtpOcDefaultDatasetListedUnicastSlavesOnlyEnable() string {
	tempData = 0x00000000
	ListedUnicastSlavesOnly := ""
	if readRegister(PtpOcCore.BaseAddrLReg+ptpOc["ConfigProfileReg"], &tempData) == 0 {
		switch (tempData >> 23) & 0x00000001 {
		case 0:
			ListedUnicastSlavesOnly = "FALSE"
		case 1:
			ListedUnicastSlavesOnly = "TRUE"
		default:
			ListedUnicastSlavesOnly = "FALSE"
		}
	}
	return ListedUnicastSlavesOnly
}

func readPtpOcDelayMechanismValue() string {
	tempData = 0x00000000
	delayMechanismValue := ""
	if readRegister(PtpOcCore.BaseAddrLReg+ptpOc["ConfigProfileReg"], &tempData) == 0 {
		switch (tempData >> 24) & 0x00000001 {
		case 0:
			delayMechanismValue = "P2P"
		case 1:
			if (tempData & 0x02000000) == 0 {
				delayMechanismValue = "E2E"
			} else {
				delayMechanismValue = "E2E Unicast"
			}
		default:
			delayMechanismValue = "NA"
		}
	}
	return delayMechanismValue
}

func readPtpOcIpMode() string {

	tempData = 0x00000000
	ipMode := ""
	// mode & server config
	if readRegister(PtpOcCore.BaseAddrLReg+ptpOc["ConfigModeReg"], &tempData) == 0 {

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

// read PtpOc IP ADDRESS
func readPtpOcIpAddress() string {
	tempData = 0x00000000
	ipAddr := ""

	// ip
	if readPtpOcLayer() == "Layer 3v4" {
		tempAddr := make([]byte, 4)
		temp_ip4 := make([]int64, 4)
		if readRegister(PtpOcCore.BaseAddrLReg+ptpOc["ConfigIpReg"], &tempData) == 0 {
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
	} else if readPtpOcLayer() == "Layer 3v6" {
		//fmt.Println("IPV6 read")
		tempAddr := make([]byte, 16)
		temp_ip6 := make([]int64, 16)
		if readRegister(PtpOcCore.BaseAddrLReg+ptpOc["ConfigIpReg"], &tempData) == 0 {
			temp_ip6[0] = (tempData >> 0) & 0x000000FF
			temp_ip6[1] = (tempData >> 8) & 0x000000FF
			temp_ip6[2] = (tempData >> 16) & 0x000000FF
			temp_ip6[3] = (tempData >> 24) & 0x000000FF
			fmt.Println("test 1")

			if readRegister(PtpOcCore.BaseAddrLReg+ptpOc["ConfigIpv61Reg"], &tempData) == 0 {
				temp_ip6[4] = (tempData >> 0) & 0x000000FF
				temp_ip6[5] = (tempData >> 8) & 0x000000FF
				temp_ip6[6] = (tempData >> 16) & 0x000000FF
				temp_ip6[7] = (tempData >> 24) & 0x000000FF
				fmt.Println("test 2")

				if readRegister(PtpOcCore.BaseAddrLReg+ptpOc["ConfigIpv62Reg"], &tempData) == 0 {
					temp_ip6[8] = (tempData >> 0) & 0x000000FF
					temp_ip6[9] = (tempData >> 8) & 0x000000FF
					temp_ip6[10] = (tempData >> 16) & 0x000000FF
					temp_ip6[11] = (tempData >> 24) & 0x000000FF
					fmt.Println("test 3")

					if readRegister(PtpOcCore.BaseAddrLReg+ptpOc["ConfigIpv63Reg"], &tempData) == 0 {

						temp_ip6[12] = (tempData >> 0) & 0x000000FF
						temp_ip6[13] = (tempData >> 8) & 0x000000FF
						temp_ip6[14] = (tempData >> 16) & 0x000000FF
						temp_ip6[15] = (tempData >> 24) & 0x000000FF
						fmt.Println("test 4")

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

func readPtpOcDefaultDatasetClockId() string {

	tempData = 0x40000000
	clockId := ""
	this_string := make([]byte, 0, 32)

	if writeRegister(PtpOcCore.BaseAddrLReg+ptpOc["DefaultDsControlReg"], &tempData) == 0 {
		for i := range 10 {
			if i < 9 {
				if readRegister(PtpOcCore.BaseAddrLReg+ptpOc["DefaultDsControlReg"], &tempData) == 0 {
					if tempData&0x80000000 != 0 {
						if readRegister(PtpOcCore.BaseAddrLReg+ptpOc["DefaultDs1Reg"], &tempData) == 0 {
							this_string = append(this_string, fmt.Sprintf("%02x", (tempData>>0)&0x000000FF)...)
							this_string = append(this_string, ':')
							this_string = append(this_string, fmt.Sprintf("%02x", (tempData>>8)&0x000000FF)...)
							this_string = append(this_string, ':')
							this_string = append(this_string, fmt.Sprintf("%02x", (tempData>>16)&0x000000FF)...)
							this_string = append(this_string, ':')
							this_string = append(this_string, fmt.Sprintf("%02x", (tempData>>24)&0x000000FF)...)
							this_string = append(this_string, ':')

						}
						if readRegister(PtpOcCore.BaseAddrLReg+ptpOc["DefaultDs2Reg"], &tempData) == 0 {
							this_string = append(this_string, fmt.Sprintf("%02x", (tempData>>0)&0x000000FF)...)
							this_string = append(this_string, ':')
							this_string = append(this_string, fmt.Sprintf("%02x", (tempData>>8)&0x000000FF)...)
							this_string = append(this_string, ':')
							this_string = append(this_string, fmt.Sprintf("%02x", (tempData>>16)&0x000000FF)...)
							this_string = append(this_string, ':')
							this_string = append(this_string, fmt.Sprintf("%02x", (tempData>>24)&0x000000FF)...)
							clockId = string(this_string)

						} else {
							clockId = "NA"
						}
					} else {
						clockId = "NA"
					}
					break // success so return
				}
			} else if i == 9 {
				log.Fatal("read did not complete")
			} else {
				clockId = "NA"
			}

		}
	}
	return clockId
}

func readPtpOcDefaultDatasetDomain() string {

	tempData = 0x40000000
	domain := ""

	if writeRegister(PtpOcCore.BaseAddrLReg+ptpOc["DefaultDsControlReg"], &tempData) == 0 {
		for i := range 10 {
			if i < 9 {
				if readRegister(PtpOcCore.BaseAddrLReg+ptpOc["DefaultDsControlReg"], &tempData) == 0 {
					if tempData&0x80000000 != 0 {

						if readRegister(PtpOcCore.BaseAddrLReg+ptpOc["DefaultDs3Reg"], &tempData) == 0 {
							domain = fmt.Sprintf("0x%02x", (tempData>>0)&0x000000FF)

						} else {
							domain = "NA"
						}
					} else {
						domain = "NA"
					}
					break // success so return
				}
			} else if i == 9 {
				log.Fatal("read did not complete")
			} else {
				domain = "NA"
			}

		}
	}
	return domain
}

func readPtpOcDefaultDatasetPriority1() string {

	tempData = 0x40000000
	priority1 := ""

	if writeRegister(PtpOcCore.BaseAddrLReg+ptpOc["DefaultDsControlReg"], &tempData) == 0 {
		for i := range 10 {
			if i < 9 {
				if readRegister(PtpOcCore.BaseAddrLReg+ptpOc["DefaultDsControlReg"], &tempData) == 0 {
					if tempData&0x80000000 != 0 {

						if readRegister(PtpOcCore.BaseAddrLReg+ptpOc["DefaultDs3Reg"], &tempData) == 0 {
							priority1 = fmt.Sprintf("0x%02x", (tempData>>24)&0x000000FF)

						} else {
							priority1 = "NA"
						}
					} else {
						priority1 = "NA"
					}
					break // success so return
				}
			} else if i == 9 {
				log.Fatal("read did not complete")
			} else {
				priority1 = "NA"
			}

		}
	}
	return priority1
}

func readPtpOcDefaultDatasetPriority2() string {

	tempData = 0x40000000
	priority2 := ""

	if writeRegister(PtpOcCore.BaseAddrLReg+ptpOc["DefaultDsControlReg"], &tempData) == 0 {
		for i := range 10 {
			if i < 9 {
				if readRegister(PtpOcCore.BaseAddrLReg+ptpOc["DefaultDsControlReg"], &tempData) == 0 {
					if tempData&0x80000000 != 0 {

						if readRegister(PtpOcCore.BaseAddrLReg+ptpOc["DefaultDs3Reg"], &tempData) == 0 {
							priority2 = fmt.Sprintf("0x%02x", (tempData>>16)&0x000000FF)

						} else {
							priority2 = "NA"
						}
					} else {
						priority2 = "NA"
					}
					break // success so return
				}
			} else if i == 9 {
				log.Fatal("read did not complete")
			} else {
				priority2 = "NA"
			}

		}
	}
	return priority2
}

func readPtpOcDefaultDatasetAccuracy() string {

	tempData = 0x40000000
	accuracy := ""

	if writeRegister(PtpOcCore.BaseAddrLReg+ptpOc["DefaultDsControlReg"], &tempData) == 0 {
		for i := range 10 {
			if i < 9 {
				if readRegister(PtpOcCore.BaseAddrLReg+ptpOc["DefaultDsControlReg"], &tempData) == 0 {
					if tempData&0x80000000 != 0 {

						if readRegister(PtpOcCore.BaseAddrLReg+ptpOc["DefaultDs4Reg"], &tempData) == 0 {
							accuracy = fmt.Sprintf("%02d", (tempData>>16)&0x000000FF) // base 10

						} else {
							accuracy = "NA"
						}
					} else {
						accuracy = "NA"
					}
					break // success so return
				}
			} else if i == 9 {
				log.Fatal("read did not complete")
			} else {
				accuracy = "NA"
			}

		}
	}
	return accuracy
}

func readPtpOcDefaultDatasetClass() string {

	tempData = 0x40000000
	class := ""

	if writeRegister(PtpOcCore.BaseAddrLReg+ptpOc["DefaultDsControlReg"], &tempData) == 0 {
		for i := range 10 {
			if i < 9 {
				if readRegister(PtpOcCore.BaseAddrLReg+ptpOc["DefaultDsControlReg"], &tempData) == 0 {
					if tempData&0x80000000 != 0 {

						if readRegister(PtpOcCore.BaseAddrLReg+ptpOc["DefaultDs4Reg"], &tempData) == 0 {
							class = fmt.Sprintf("0x%02d", (tempData>>24)&0x000000FF) // base 10

						} else {
							class = "NA"
						}
					} else {
						class = "NA"
					}
					break // success so return
				}
			} else if i == 9 {
				log.Fatal("read did not complete")
			} else {
				class = "NA"
			}

		}
	}
	return class
}

func readPtpOcDefaultDatasetVariance() string {

	tempData = 0x40000000
	variance := ""

	if writeRegister(PtpOcCore.BaseAddrLReg+ptpOc["DefaultDsControlReg"], &tempData) == 0 {
		for i := range 10 {
			if i < 9 {
				if readRegister(PtpOcCore.BaseAddrLReg+ptpOc["DefaultDsControlReg"], &tempData) == 0 {
					if tempData&0x80000000 != 0 {

						if readRegister(PtpOcCore.BaseAddrLReg+ptpOc["DefaultDs4Reg"], &tempData) == 0 {
							variance = fmt.Sprintf("0x%02x", (tempData>>0)&0x0000FFFF)

						} else {
							variance = "NA"
						}
					} else {
						variance = "NA"
					}
					break // success so return
				}
			} else if i == 9 {
				log.Fatal("read did not complete")
			} else {
				variance = "NA"
			}

		}
	}
	return variance
}

func readPtpOcDefaultDatasetShortId() string {

	tempData = 0x40000000
	shortId := ""

	if writeRegister(PtpOcCore.BaseAddrLReg+ptpOc["DefaultDsControlReg"], &tempData) == 0 {
		for i := range 10 {
			if i < 9 {
				if readRegister(PtpOcCore.BaseAddrLReg+ptpOc["DefaultDsControlReg"], &tempData) == 0 {
					if tempData&0x80000000 != 0 {

						if readRegister(PtpOcCore.BaseAddrLReg+ptpOc["DefaultDs5Reg"], &tempData) == 0 {
							shortId = fmt.Sprintf("0x%04x", tempData)

						} else {
							shortId = "NA"
						}
					} else {
						shortId = "NA"
					}
					break // success so return
				}
			} else if i == 9 {
				log.Fatal("read did not complete")
			} else {
				shortId = "NA"
			}

		}
	}
	return shortId
}

func readPtpOcDefaultDatasetInaccuracy() string {

	tempData = 0x40000000
	inaccuracy := ""

	if writeRegister(PtpOcCore.BaseAddrLReg+ptpOc["DefaultDsControlReg"], &tempData) == 0 {
		for i := range 10 {
			if i < 9 {
				if readRegister(PtpOcCore.BaseAddrLReg+ptpOc["DefaultDsControlReg"], &tempData) == 0 {
					if tempData&0x80000000 != 0 {

						if readRegister(PtpOcCore.BaseAddrLReg+ptpOc["DefaultDs6Reg"], &tempData) == 0 {
							inaccuracy = fmt.Sprintf("%d", tempData)

						} else {
							inaccuracy = "NA"
						}
					} else {
						inaccuracy = "NA"
					}
					break // success so return
				}
			} else if i == 9 {
				log.Fatal("read did not complete")
			} else {
				inaccuracy = "NA"
			}

		}
	}
	return inaccuracy
}

func readPtpOcDefaultDatasetNumPorts() string {

	tempData = 0x40000000
	numPorts := ""

	if writeRegister(PtpOcCore.BaseAddrLReg+ptpOc["DefaultDsControlReg"], &tempData) == 0 {
		for i := range 10 {
			if i < 9 {
				if readRegister(PtpOcCore.BaseAddrLReg+ptpOc["DefaultDsControlReg"], &tempData) == 0 {
					if tempData&0x80000000 != 0 {

						if readRegister(PtpOcCore.BaseAddrLReg+ptpOc["DefaultDs7Reg"], &tempData) == 0 {
							numPorts = fmt.Sprintf("%d", tempData)

						} else {
							numPorts = "NA"
						}
					} else {
						numPorts = "NA"
					}
					break // success so return
				}
			} else if i == 9 {
				log.Fatal("read did not complete")
			} else {
				numPorts = "NA"
			}

		}
	}
	return numPorts
}

func readPtpOcPortDatasetPeerDelay() string {

	tempData = 0x40000000
	var tempDelay int64
	peerDelay := ""

	if writeRegister(PtpOcCore.BaseAddrLReg+ptpOc["PortDsControlReg"], &tempData) == 0 {
		for i := range 10 {
			if i < 9 {
				if readRegister(PtpOcCore.BaseAddrLReg+ptpOc["PortDsControlReg"], &tempData) == 0 {
					if tempData&0x80000000 != 0 {

						if (readPtpOcDelayMechanismValue() == "E2E") || (readPtpOcDelayMechanismValue() == "E2E Unicast") {
							// end to end delay
							peerDelay = "NA"
						} else if readRegister(PtpOcCore.BaseAddrLReg+ptpOc["PortDs1Reg"], &tempData) == 0 {
							tempDelay = tempData
							tempData = tempDelay << 32
							if readRegister(PtpOcCore.BaseAddrLReg+ptpOc["PortDs2Reg"], &tempData) == 0 {
								tempDelay |= tempData
								tempDelay = tempDelay >> 16

								peerDelay = fmt.Sprintf("%d", tempDelay)
							} else {
								peerDelay = "NA"
							}
						} else {
							peerDelay = "NA"
						}

					} else {
						peerDelay = "NA"

					}
					break // success so return
				}
			} else if i == 9 {
				log.Fatal("read did not complete")
			} else {
				peerDelay = "NA"

			}

		}
	}
	return peerDelay
}

func readPtpOcPortDatasetState() string {
	tempData = 0x40000000
	state := ""

	if writeRegister(PtpOcCore.BaseAddrLReg+ptpOc["PortDsControlReg"], &tempData) == 0 {
		for i := range 10 {
			if i < 9 {
				if readRegister(PtpOcCore.BaseAddrLReg+ptpOc["PortDsControlReg"], &tempData) == 0 {
					if tempData&0x80000000 != 0 {

						if readRegister(PtpOcCore.BaseAddrLReg+ptpOc["PortDs3Reg"], &tempData) == 0 {
							switch tempData {

							case 0x0000001:
								state = "INITIALIZING"

							case 0x00000002:
								state = "FAULTY"

							case 0x00000003:
								state = "DISABLED"

							case 0x00000004:
								state = "LISTENING"

							case 0x00000005:
								state = "PREMASTER"

							case 0x00000006:
								state = "MASTER"

							case 0x00000007:
								state = "PASSIVE"

							case 0x00000008:
								state = "UNCALIBRATED"

							case 0x00000009:
								state = "SLAVE"

							default:
								state = "NA"
							}

							break
						} else {
							state = "NA"

						}
					}
				}
			} else if i == 9 {
				log.Fatal("read did not complete")
			} else {
				state = "NA"

			}
		}
	}
	return state
}

func readPtpOcPortDatasetAsymmetry() string {
	tempData = 0x40000000
	asymmetry := ""

	if writeRegister(PtpOcCore.BaseAddrLReg+ptpOc["PortDsControlReg"], &tempData) == 0 {
		for i := range 10 {
			if i < 9 {
				if readRegister(PtpOcCore.BaseAddrLReg+ptpOc["PortDsControlReg"], &tempData) == 0 {
					if tempData&0x80000000 != 0 {

						if readRegister(PtpOcCore.BaseAddrLReg+ptpOc["PortDs7Reg"], &tempData) == 0 {
							asymmetry = fmt.Sprintf("%d", tempData)
							break // success so return

						} else {
							asymmetry = "NA"
						}
					} else {
						asymmetry = "NA"
					}

				} else {
					asymmetry = "NA"

				}
			} else if i == 9 {
				log.Fatal("read did not complete")
			} else {
				asymmetry = "NA"

			}

		}
	}
	return asymmetry
}

func readPtpOcPortDatasetMaxDelay() string {

	tempData = 0x40000000
	maxDelay := ""

	if writeRegister(PtpOcCore.BaseAddrLReg+ptpOc["PortDsControlReg"], &tempData) == 0 {
		for i := range 10 {
			if i < 9 {
				if readRegister(PtpOcCore.BaseAddrLReg+ptpOc["PortDsControlReg"], &tempData) == 0 {
					if tempData&0x80000000 != 0 {

						if readRegister(PtpOcCore.BaseAddrLReg+ptpOc["PortDs8Reg"], &tempData) == 0 {
							maxDelay = fmt.Sprintf("%d", tempData)
							break // success so return

						} else {
							maxDelay = "NA"
						}
					} else {
						maxDelay = "NA"
					}

				} else {
					maxDelay = "NA"

				}
			} else if i == 9 {
				log.Fatal("read did not complete")
			} else {
				maxDelay = "NA"

			}

		}
	}
	return maxDelay
}

func readPtpOcPortDatasetPDelayReqLogMsgInterval() string {
	tempData = 0x40000000
	pDelay := ""

	if writeRegister(PtpOcCore.BaseAddrLReg+ptpOc["PortDsControlReg"], &tempData) == 0 {
		for i := range 10 {
			if i < 9 {
				if readRegister(PtpOcCore.BaseAddrLReg+ptpOc["PortDsControlReg"], &tempData) == 0 {
					if tempData&0x80000000 != 0 {

						if readRegister(PtpOcCore.BaseAddrLReg+ptpOc["PortDs4Reg"], &tempData) == 0 {
							pDelay = fmt.Sprintf("%d", (tempData & 0x000000FF))
							break // success so return

						} else {
							pDelay = "NA"
						}
					} else {
						pDelay = "NA"
					}

				} else {
					pDelay = "NA"

				}
			} else if i == 9 {
				log.Fatal("read did not complete")
			} else {
				pDelay = "NA"

			}

		}
	}
	return pDelay

}

func readPtpOcPortDatasetDelayReqLogMsgInterval() string {
	tempData = 0x40000000
	delay := ""

	if writeRegister(PtpOcCore.BaseAddrLReg+ptpOc["PortDsControlReg"], &tempData) == 0 {
		for i := range 10 {
			if i < 9 {
				if readRegister(PtpOcCore.BaseAddrLReg+ptpOc["PortDsControlReg"], &tempData) == 0 {
					if tempData&0x80000000 != 0 {

						if readRegister(PtpOcCore.BaseAddrLReg+ptpOc["PortDs4Reg"], &tempData) == 0 {
							delay = fmt.Sprintf("%d", ((tempData >> 8) & 0x000000FF))
							break // success so return
						} else {
							delay = "NA"
						}
					} else {
						delay = "NA"
					}
				} else {
					delay = "NA"
				}
			} else if i == 9 {
				log.Fatal("read did not complete")
			} else {
				delay = "NA"
			}
		}
	}
	return delay
}

func readPtpOcPortDatasetDelayReceiptTimeout() string {
	tempData = 0x40000000
	delay := ""

	if writeRegister(PtpOcCore.BaseAddrLReg+ptpOc["PortDsControlReg"], &tempData) == 0 {
		for i := range 10 {
			if i < 9 {
				if readRegister(PtpOcCore.BaseAddrLReg+ptpOc["PortDsControlReg"], &tempData) == 0 {
					if tempData&0x80000000 != 0 {

						if readRegister(PtpOcCore.BaseAddrLReg+ptpOc["PortDs4Reg"], &tempData) == 0 {
							delay = fmt.Sprintf("%d", ((tempData >> 16) & 0x000000FF))
							break // success so return
						} else {
							delay = "NA"
						}
					} else {
						delay = "NA"
					}
				} else {
					delay = "NA"
				}
			} else if i == 9 {
				log.Fatal("read did not complete")
			} else {
				delay = "NA"
			}
		}
	}
	return delay
}

func readPtpOcPortDatasetAnnounceLogMsgInterval() string {
	tempData = 0x40000000
	delay := ""

	if writeRegister(PtpOcCore.BaseAddrLReg+ptpOc["PortDsControlReg"], &tempData) == 0 {
		for i := range 10 {
			if i < 9 {
				if readRegister(PtpOcCore.BaseAddrLReg+ptpOc["PortDsControlReg"], &tempData) == 0 {
					if tempData&0x80000000 != 0 {

						if readRegister(PtpOcCore.BaseAddrLReg+ptpOc["PortDs5Reg"], &tempData) == 0 {
							delay = fmt.Sprintf("%d", (tempData & 0x000000FF))
							break // success so return
						} else {
							delay = "NA"
						}
					} else {
						delay = "NA"
					}
				} else {
					delay = "NA"
				}
			} else if i == 9 {
				log.Fatal("read did not complete")
			} else {
				delay = "NA"
			}
		}
	}
	return delay
}

func readPtpOcPortDatasetAnnounceReceiptTimeout() string {
	tempData = 0x40000000
	delay := ""

	if writeRegister(PtpOcCore.BaseAddrLReg+ptpOc["PortDsControlReg"], &tempData) == 0 {
		for i := range 10 {
			if i < 9 {
				if readRegister(PtpOcCore.BaseAddrLReg+ptpOc["PortDsControlReg"], &tempData) == 0 {
					if tempData&0x80000000 != 0 {

						if readRegister(PtpOcCore.BaseAddrLReg+ptpOc["PortDs5Reg"], &tempData) == 0 {
							delay = fmt.Sprintf("%d", ((tempData >> 8) & 0x000000FF))
							break // success so return
						} else {
							delay = "NA"
						}
					} else {
						delay = "NA"
					}
				} else {
					delay = "NA"
				}
			} else if i == 9 {
				log.Fatal("read did not complete")
			} else {
				delay = "NA"
			}
		}
	}
	return delay
}

func readPtpOcPortDatasetSyncLogMsgInterval() string {
	tempData = 0x40000000
	delay := ""

	if writeRegister(PtpOcCore.BaseAddrLReg+ptpOc["PortDsControlReg"], &tempData) == 0 {
		for i := range 10 {
			if i < 9 {
				if readRegister(PtpOcCore.BaseAddrLReg+ptpOc["PortDsControlReg"], &tempData) == 0 {
					if tempData&0x80000000 != 0 {

						if readRegister(PtpOcCore.BaseAddrLReg+ptpOc["PortDs6Reg"], &tempData) == 0 {
							delay = fmt.Sprintf("%d", (tempData & 0x000000FF))
							break // success so return
						} else {
							delay = "NA"
						}
					} else {
						delay = "NA"
					}
				} else {
					delay = "NA"
				}
			} else if i == 9 {
				log.Fatal("read did not complete")
			} else {
				delay = "NA"
			}
		}
	}
	return delay
}

func readPtpOcPortDatasetSyncReceiptTimeout() string {
	tempData = 0x40000000
	delay := ""

	if writeRegister(PtpOcCore.BaseAddrLReg+ptpOc["PortDsControlReg"], &tempData) == 0 {
		for i := range 10 {
			if i < 9 {
				if readRegister(PtpOcCore.BaseAddrLReg+ptpOc["PortDsControlReg"], &tempData) == 0 {
					if tempData&0x80000000 != 0 {

						if readRegister(PtpOcCore.BaseAddrLReg+ptpOc["PortDs6Reg"], &tempData) == 0 {
							delay = fmt.Sprintf("%d", ((tempData >> 8) & 0x000000FF))
							break // success so return
						} else {
							delay = "NA"
						}
					} else {
						delay = "NA"
					}
				} else {
					delay = "NA"
				}
			} else if i == 9 {
				log.Fatal("read did not complete")
			} else {
				delay = "NA"
			}
		}
	}
	return delay
}

func showPtpOcSTATUS() {
	fmt.Println("PTP OC STATUS:                     ", readPtpOcStatus())
}

func showPtpOcINSTANCE() {
	fmt.Println("PTP OC INSTANCE:                   ", readPtpOcInstanceNumber())
}

func showPtpOcIpAddress() {
	fmt.Println("PTP OC IP ADDRESS:                 ", readPtpOcIpAddress())
}

func showPtpOcVlanEnable() {
}

func showPtpOcAll() {

	fmt.Println("PTP OC VERSION:                                    ", readPtpOcVersion())
	fmt.Println("PTP OC STATUS:                                     ", readPtpOcStatus())
	fmt.Println("PTP OC INSTANCE:                                   ", readPtpOcInstanceNumber())
	fmt.Println("PTP OC IP ADDRESS:                                 ", readPtpOcIpAddress())
	fmt.Println("PTP OC VLAN ENABLE:                                ", readPtpOcVlanEnable())
	fmt.Println("PTP OC VLAN VALUE:                                 ", readPtpOcVlanValue())
	fmt.Println("PTP OC PROFILE:                                    ", readPtpOcProfile())
	fmt.Println("PTP OC LAYER:                                      ", readPtpOcLayer())
	fmt.Println("PTP OC DELAY MECHANISM:                            ", readPtpOcDelayMechanismValue())
	fmt.Println("PTP OC IP ADDRESS:                                 ", readPtpOcIpAddress())
	fmt.Println("PTP OC DEFAULT DATASET CLOCK ID:                   ", readPtpOcDefaultDatasetClockId())
	fmt.Println("PTP OC DEFAULT DATASET DOMAIN:                     ", readPtpOcDefaultDatasetDomain())
	fmt.Println("PTP OC DEFAULT DATASET PRIORITY 1:                 ", readPtpOcDefaultDatasetPriority1())
	fmt.Println("PTP OC DEFAULT DATASET PRIORITY 2:                 ", readPtpOcDefaultDatasetPriority2())
	fmt.Println("PTP OC DEFAULT DATASET ACCURACY:                   ", readPtpOcDefaultDatasetAccuracy())
	fmt.Println("PTP OC DEFAULT DATASET CLASS:                      ", readPtpOcDefaultDatasetClass())
	fmt.Println("PTP OC DEFAULT DATASET VARIANCE:                   ", readPtpOcDefaultDatasetVariance())
	fmt.Println("PTP OC DEFAULT DATASET SHORT ID:                   ", readPtpOcDefaultDatasetShortId())
	fmt.Println("PTP OC DEFAULT DATASET INACCURACY:                 ", readPtpOcDefaultDatasetInaccuracy())
	fmt.Println("PTP OC DEFAULT DATASET NUMBER OF PORTS:            ", readPtpOcDefaultDatasetNumPorts())
	fmt.Println("PTP OC DEFAULT DATASET TWO STEP:                   ", readPtpOcDefaultDatasetTwoStep())
	fmt.Println("PTP OC DEFAULT DATASET SIGNALING:                  ", readPtpOcDefaultDatasetSignaling())
	fmt.Println("PTP OC DEFAULT DATASET MASTER ONLY:                ", readPtpOcDefaultDatasetMasterOnly())
	fmt.Println("PTP OC DEFAULT DATASET SLAVE ONLY:                 ", readPtpOcDefaultDatasetSlaveOnly())
	fmt.Println("PTP OC DEFAULT DATASET OFFSET CORRECTION ENABLED:  ", readPtpOcDefaultDatasetOffsetCorrectionsEnable())
	fmt.Println("PTP OC DEFAULT DATASET LISTED UNICAST SLAVES ONLY: ", readPtpOcDefaultDatasetListedUnicastSlavesOnlyEnable())
	fmt.Println("PTP OC PORT DATASET PEER DELAY:                    ", readPtpOcPortDatasetPeerDelay())
	fmt.Println("PTP OC PORT DATASET STATE:                         ", readPtpOcPortDatasetState())
	fmt.Println("PTP OC PORT DATASET ASYMMETRY:                     ", readPtpOcPortDatasetAsymmetry())
	fmt.Println("PTP OC PORT DATASET MAX DELAY [ns]:                ", readPtpOcPortDatasetMaxDelay())
	fmt.Println("PTP OC PORT DATASET P-DELAY-REQ-LOG-MSG-INTERVAL:  ", readPtpOcPortDatasetPDelayReqLogMsgInterval())
	fmt.Println("PTP OC PORT DATASET DELAY-REQ-LOG-MSG-INTERVAL:    ", readPtpOcPortDatasetDelayReqLogMsgInterval())
	fmt.Println("PTP OC PORT DATASET DELAY RECEIPT TIMEOUT:         ", readPtpOcPortDatasetDelayReceiptTimeout())
	fmt.Println("PTP OC PORT DATASET ANNOUCE LOG MSG INTERVAL:      ", readPtpOcPortDatasetAnnounceLogMsgInterval())
	fmt.Println("PTP OC PORT DATASET ANNOUCE RECEIPT TIMEOUT:       ", readPtpOcPortDatasetAnnounceReceiptTimeout())
	fmt.Println("PTP OC PORT DATASET SYNC LOG MSG INTERVAL:         ", readPtpOcPortDatasetSyncLogMsgInterval())
	fmt.Println("PTP OC PORT DATASET SYNC RECEIPT TIMEOUT:          ", readPtpOcPortDatasetSyncReceiptTimeout())
	fmt.Println("PTP OC PORT DATASET SET CUSTOM INTERVALS:         WRITE ONLY ")
}

/*

void Ucm_PtpOcTab::ptp_oc_read_values(void)
{
    unsigned long long temp_next_jump;
    unsigned long long temp_delay;
    long long temp_signed_delay;
    unsigned long long temp_offset;
    long long temp_signed_offset;
    unsigned int temp_length;
    unsigned long temp_ip;
    int temp_min = 0;
    int temp_max = 0;
    unsigned int temp_data = 0;
    unsigned int temp_addr = 0;
    QString temp_string;
    quint8 temp_ip6[16];

    temp_string = ui->PtpOcInstanceComboBox->currentText();
    temp_data = temp_string.toUInt(nullptr, 10);
    for (int i = 0; i < ucm->core_config.size(); i++)
    {
        if ((ucm->core_config.at(i).core_type == Ucm_CoreConfig_PtpOrdinaryClockCoreType) && (ucm->core_config.at(i).core_instance_nr == temp_data))
        {
            temp_addr = ucm->core_config.at(i).address_range_low;
            break;
        }
        else if (i == (ucm->core_config.size()-1))
        {
            ui->PtpOcVlanValue->setText("NA");
            ui->PtpOcProfileValue->setCurrentText("NA");
            ui->PtpOcLayerValue->setCurrentText("NA");
            ui->PtpOcDelayMechanismValue->setCurrentText("NA");
            ui->PtpOcIpValue->setText("NA");
            ui->PtpOcVersionValue->setText("NA");

            ui->PtpOcDefaultDsClockIdValue->setText("NA");
            ui->PtpOcDefaultDsDomainValue->setText("NA");
            ui->PtpOcDefaultDsPriority1Value->setText("NA");
            ui->PtpOcDefaultDsPriority2Value->setText("NA");
            ui->PtpOcDefaultDsVarianceValue->setText("NA");
            ui->PtpOcDefaultDsAccuracyValue->setText("NA");
            ui->PtpOcDefaultDsClassValue->setText("NA");
            ui->PtpOcDefaultDsShortIdValue->setText("NA");
            ui->PtpOcDefaultDsInaccuracyValue->setText("NA");
            ui->PtpOcDefaultDsNrOfPortsValue->setText("NA");
            ui->PtpOcDefaultDsTwoStepCheckBox->setChecked(false);
            ui->PtpOcDefaultDsSignalingCheckBox->setChecked(false);
            ui->PtpOcDefaultDsSlaveOnlyCheckBox->setChecked(false);
            ui->PtpOcDefaultDsMasterOnlyCheckBox->setChecked(false);
            ui->PtpOcDefaultDsListedUnicastSlavesOnlyCheckBox->setChecked(false);
            ui->PtpOcDefaultDsDisableOffsetCorCheckBox->setChecked(false);

            ui->PtpOcPortDsPeerDelayValue->setText("NA");
            ui->PtpOcPortDsStateValue->setText("NA");
            ui->PtpOcPortDsAsymmetryValue->setText("NA");
            ui->PtpOcPortDsMaxPeerDelayValue->setText("NA");
            ui->PtpOcPortDsPDelayReqLogMsgIntervalValue->setText("NA");
            ui->PtpOcPortDsDelayReceiptTimeoutValue->setText("NA");
            ui->PtpOcPortDsDelayReqLogMsgIntervalValue->setText("NA");
            ui->PtpOcPortDsAnnounceLogMsgIntervalValue->setText("NA");
            ui->PtpOcPortDsAnnounceReceiptTimeoutValue->setText("NA");
            ui->PtpOcPortDsSyncLogMsgIntervalValue->setText("NA");
            ui->PtpOcPortDsSyncReceiptTimeoutValue->setText("NA");

            ui->PtpOcCurrentDsStepsRemovedValue->setText("NA");
            ui->PtpOcCurrentDsOffsetValue->setText("NA");

            ui->PtpOcParentDsParentClockIdValue->setText("NA");
            ui->PtpOcParentDsGmClockIdValue->setText("NA");
            ui->PtpOcParentDsGmPriority1Value->setText("NA");
            ui->PtpOcParentDsGmPriority2Value->setText("NA");
            ui->PtpOcParentDsGmAccuracyValue->setText("NA");
            ui->PtpOcParentDsGmClassValue->setText("NA");
            ui->PtpOcParentDsGmShortIdValue->setText("NA");
            ui->PtpOcParentDsGmInaccuracyValue->setText("NA");
            ui->PtpOcParentDsNwInaccuracyValue->setText("NA");

            ui->PtpOcTimePropertiesDsTimeSourceValue->setText("NA");
            ui->PtpOcTimePropertiesDsPtpTimescaleCheckBox->setChecked(false);
            ui->PtpOcTimePropertiesDsFreqTraceableCheckBox->setChecked(false);
            ui->PtpOcTimePropertiesDsTimeTraceableCheckBox->setChecked(false);
            ui->PtpOcTimePropertiesDsLeap59CheckBox->setChecked(false);
            ui->PtpOcTimePropertiesDsLeap61CheckBox->setChecked(false);
            ui->PtpOcTimePropertiesDsUtcOffsetValCheckBox->setChecked(false);
            ui->PtpOcTimePropertiesDsUtcOffsetValue->setText("NA");
            ui->PtpOcTimePropertiesDsCurrentOffsetValue->setText("NA");
            ui->PtpOcTimePropertiesDsJumpSecondsValue->setText("NA");
            ui->PtpOcTimePropertiesDsNextJumpValue->setText("NA");
            ui->PtpOcTimePropertiesDsDisplayNameValue->setText("NA");
            ui->PtpOcTimePropertiesDsSetLocalPropertiesCheckBox->setChecked(false);

            ui->PtpOcVlanEnableCheckBox->setChecked(false);
            ui->PtpOcVersionValue->setText("NA");
            return;
        }
    }

    // enabled
    if (0 == ucm->com_lib.read_reg(temp_addr + Ucm_PtpOc_ControlReg, temp_data))
    {
        if ((temp_data & 0x00000001) == 0)
        {
            ui->PtpOcEnableCheckBox->setChecked(false);
        }
        else
        {
            ui->PtpOcEnableCheckBox->setChecked(true);
        }
    }
    else
    {
        ui->PtpOcEnableCheckBox->setChecked(false);
    }


    // vlan
    if (0 == ucm->com_lib.read_reg(temp_addr + Ucm_PtpOc_ConfigVlanReg, temp_data))
    {
        if ((temp_data & 0x00010000) == 0)
        {
            ui->PtpOcVlanEnableCheckBox->setChecked(false);
        }
        else
        {
            ui->PtpOcVlanEnableCheckBox->setChecked(true);
        }

        temp_data &= 0x0000FFFF;

        ui->PtpOcVlanValue->setText(QString("0x%1").arg(temp_data, 4, 16, QLatin1Char('0')));
    }
    else
    {
        ui->PtpOcVlanEnableCheckBox->setChecked(false);
        ui->PtpOcVlanValue->setText("NA");
    }


    // profile and layer
    if (0 == ucm->com_lib.read_reg(temp_addr + Ucm_PtpOc_ConfigProfileReg, temp_data))
    {
        switch (temp_data & 0x00000007)
        {
        case 0:
            ui->PtpOcProfileValue->setCurrentText("Default");
            break;
        case 1:
            ui->PtpOcProfileValue->setCurrentText("Power");
            break;
        case 2:
            ui->PtpOcProfileValue->setCurrentText("Utility");
            break;
        case 3:
            ui->PtpOcProfileValue->setCurrentText("TSN");
            break;
        case 4:
            ui->PtpOcProfileValue->setCurrentText("ITUG8265.1");
            break;
        case 5:
            ui->PtpOcProfileValue->setCurrentText("ITUG8275.1");
            break;
        case 6:
            ui->PtpOcProfileValue->setCurrentText("ITUG8275.2");
            break;
        default:
            ui->PtpOcProfileValue->setCurrentText("NA");
            break;
        }

        switch ((temp_data >> 8) & 0x00000001)
        {
        case 0:
            ui->PtpOcDefaultDsTwoStepCheckBox->setChecked(false);
            break;
        case 1:
            ui->PtpOcDefaultDsTwoStepCheckBox->setChecked(true);
            break;
        default:
            ui->PtpOcDefaultDsTwoStepCheckBox->setChecked(false);
            break;
        }

        switch ((temp_data >> 9) & 0x00000001)
        {
        case 0:
            ui->PtpOcDefaultDsSignalingCheckBox->setChecked(false);
            break;
        case 1:
            ui->PtpOcDefaultDsSignalingCheckBox->setChecked(true);
            break;
        default:
            ui->PtpOcDefaultDsSignalingCheckBox->setChecked(false);
            break;
        }

        switch ((temp_data >> 16) & 0x00000003)
        {
        case 0:
            ui->PtpOcLayerValue->setCurrentText("Layer 2");
            break;
        case 1:
            ui->PtpOcLayerValue->setCurrentText("Layer 3v4");
            break;
        case 2:
            ui->PtpOcLayerValue->setCurrentText("Layer 3v6");
            break;
        default:
            ui->PtpOcLayerValue->setCurrentText("NA");
            break;
        }

        switch ((temp_data >> 20) & 0x00000003)
        {
        case 0:
            ui->PtpOcDefaultDsSlaveOnlyCheckBox->setChecked(false);
            ui->PtpOcDefaultDsMasterOnlyCheckBox->setChecked(false);
            break;
        case 1:
            ui->PtpOcDefaultDsSlaveOnlyCheckBox->setChecked(true);
            ui->PtpOcDefaultDsMasterOnlyCheckBox->setChecked(false);
            break;
        case 2:
            ui->PtpOcDefaultDsSlaveOnlyCheckBox->setChecked(false);
            ui->PtpOcDefaultDsMasterOnlyCheckBox->setChecked(true);
            break;
        default:
            ui->PtpOcDefaultDsSlaveOnlyCheckBox->setChecked(false);
            ui->PtpOcDefaultDsMasterOnlyCheckBox->setChecked(false);
            break;
        }

        switch ((temp_data >> 22) & 0x00000001)
        {
        case 0:
            ui->PtpOcDefaultDsDisableOffsetCorCheckBox->setChecked(false);
            break;
        case 1:
            ui->PtpOcDefaultDsDisableOffsetCorCheckBox->setChecked(true);
            break;
        default:
            ui->PtpOcDefaultDsDisableOffsetCorCheckBox->setChecked(false);
            break;
        }

        switch ((temp_data >> 23) & 0x00000001)
        {
        case 0:
            ui->PtpOcDefaultDsListedUnicastSlavesOnlyCheckBox->setChecked(false);
            break;
        case 1:
            ui->PtpOcDefaultDsListedUnicastSlavesOnlyCheckBox->setChecked(true);
            break;
        default:
            ui->PtpOcDefaultDsListedUnicastSlavesOnlyCheckBox->setChecked(false);
            break;
        }

        switch ((temp_data >> 24) & 0x00000001)
        {
        case 0:
            ui->PtpOcDelayMechanismValue->setCurrentText("P2P");
            break;
        case 1:
            if ((temp_data & 0x02000000) == 0)
            {
                ui->PtpOcDelayMechanismValue->setCurrentText("E2E");
            }
            else
            {
                ui->PtpOcDelayMechanismValue->setCurrentText("E2E Unicast");
            }
            break;
        default:
            ui->PtpOcDelayMechanismValue->setCurrentText("NA");
            break;
        }
    }
    else
    {
        ui->PtpOcProfileValue->setCurrentText("NA");
        ui->PtpOcDefaultDsTwoStepCheckBox->setChecked(false);
        ui->PtpOcDefaultDsSignalingCheckBox->setChecked(false);
        ui->PtpOcLayerValue->setCurrentText("NA");
        ui->PtpOcDelayMechanismValue->setCurrentText("NA");
        ui->PtpOcDefaultDsSlaveOnlyCheckBox->setChecked(false);
        ui->PtpOcDefaultDsMasterOnlyCheckBox->setChecked(false);
        ui->PtpOcDefaultDsListedUnicastSlavesOnlyCheckBox->setChecked(false);
        ui->PtpOcDefaultDsDisableOffsetCorCheckBox->setChecked(false);
    }

    // ip
    temp_string = ui->PtpOcLayerValue->currentText();
    if (temp_string == "Layer 3v4")
    {
        temp_string.clear();
        if (0 == ucm->com_lib.read_reg(temp_addr + Ucm_PtpOc_ConfigIpReg, temp_data))
        {
            temp_ip = 0x00000000;
            temp_ip |= (temp_data >> 0) & 0x000000FF;
            temp_ip = temp_ip << 8;
            temp_ip |= (temp_data >> 8) & 0x000000FF;
            temp_ip = temp_ip << 8;
            temp_ip |= (temp_data >> 16) & 0x000000FF;
            temp_ip = temp_ip << 8;
            temp_ip |= (temp_data >> 24) & 0x000000FF;

            temp_string = QHostAddress(temp_ip).toString();

            ui->PtpOcIpValue->setText(temp_string);

        }
        else
        {
            ui->PtpOcIpValue->setText("NA");
        }
    }
    else if (temp_string == "Layer 3v6")
    {
        temp_string.clear();
        if (0 == ucm->com_lib.read_reg(temp_addr + Ucm_PtpOc_ConfigIpReg, temp_data))
        {
            temp_ip6[0] = (temp_data >> 0) & 0x000000FF;
            temp_ip6[1] = (temp_data >> 8) & 0x000000FF;
            temp_ip6[2] = (temp_data >> 16) & 0x000000FF;
            temp_ip6[3] = (temp_data >> 24) & 0x000000FF;

            if (0 == ucm->com_lib.read_reg(temp_addr + Ucm_PtpOc_ConfigIpv61Reg, temp_data))
            {
                temp_ip6[4] = (temp_data >> 0) & 0x000000FF;
                temp_ip6[5] = (temp_data >> 8) & 0x000000FF;
                temp_ip6[6] = (temp_data >> 16) & 0x000000FF;
                temp_ip6[7] = (temp_data >> 24) & 0x000000FF;

                if (0 == ucm->com_lib.read_reg(temp_addr + Ucm_PtpOc_ConfigIpv62Reg, temp_data))
                {
                    temp_ip6[8] = (temp_data >> 0) & 0x000000FF;
                    temp_ip6[9] = (temp_data >> 8) & 0x000000FF;
                    temp_ip6[10] = (temp_data >> 16) & 0x000000FF;
                    temp_ip6[11] = (temp_data >> 24) & 0x000000FF;

                    if (0 == ucm->com_lib.read_reg(temp_addr + Ucm_PtpOc_ConfigIpv63Reg, temp_data))
                    {
                        temp_ip6[12] = (temp_data >> 0) & 0x000000FF;
                        temp_ip6[13] = (temp_data >> 8) & 0x000000FF;
                        temp_ip6[14] = (temp_data >> 16) & 0x000000FF;
                        temp_ip6[15] = (temp_data >> 24) & 0x000000FF;

                        temp_string = QHostAddress(temp_ip6).toString();

                        ui->PtpOcIpValue->setText(temp_string);

                    }
                    else
                    {
                        ui->PtpOcIpValue->setText("NA");
                    }

                }
                else
                {
                    ui->PtpOcIpValue->setText("NA");
                }

            }
            else
            {
                ui->PtpOcIpValue->setText("NA");
            }

        }
        else
        {
            ui->PtpOcIpValue->setText("NA");
        }
    }
    else
    {
        ui->PtpOcIpValue->setText("NA");
    }

    //********************************
    // default dataset
    //********************************
    temp_data = 0x40000000;
    if (0 == ucm->com_lib.write_reg(temp_addr + Ucm_PtpOc_DefaultDsControlReg, temp_data))
    {
        for (int i = 0; i < 10; i++)
        {
            if(0 == ucm->com_lib.read_reg(temp_addr + Ucm_PtpOc_DefaultDsControlReg, temp_data))
            {
                if ((temp_data & 0x80000000) != 0)
                {
                    // clock id
                    temp_string.clear();
                    if (0 == ucm->com_lib.read_reg(temp_addr + Ucm_PtpOc_DefaultDs1Reg, temp_data))
                    {
                        temp_string.append(QString("%1").arg(((temp_data >> 0) & 0x000000FF), 2, 16, QLatin1Char('0')));
                        temp_string.append(":");
                        temp_string.append(QString("%1").arg(((temp_data >> 8) & 0x000000FF), 2, 16, QLatin1Char('0')));
                        temp_string.append(":");
                        temp_string.append(QString("%1").arg(((temp_data >> 16) & 0x000000FF), 2, 16, QLatin1Char('0')));
                        temp_string.append(":");
                        temp_string.append(QString("%1").arg(((temp_data >> 24) & 0x000000FF), 2, 16, QLatin1Char('0')));
                        temp_string.append(":");
                        if (0 == ucm->com_lib.read_reg(temp_addr + Ucm_PtpOc_DefaultDs2Reg, temp_data))
                        {
                            temp_string.append(QString("%1").arg(((temp_data >> 0) & 0x000000FF), 2, 16, QLatin1Char('0')));
                            temp_string.append(":");
                            temp_string.append(QString("%1").arg(((temp_data >> 8) & 0x000000FF), 2, 16, QLatin1Char('0')));
                            temp_string.append(":");
                            temp_string.append(QString("%1").arg(((temp_data >> 16) & 0x000000FF), 2, 16, QLatin1Char('0')));
                            temp_string.append(":");
                            temp_string.append(QString("%1").arg(((temp_data >> 24) & 0x000000FF), 2, 16, QLatin1Char('0')));
                            ui->PtpOcDefaultDsClockIdValue->setText(temp_string);
                        }
                        else
                        {
                            ui->PtpOcDefaultDsClockIdValue->setText("NA");
                        }
                    }
                    else
                    {
                        ui->PtpOcDefaultDsClockIdValue->setText("NA");
                    }

                    // domain, priority 1 & 2
                    if (0 == ucm->com_lib.read_reg(temp_addr + Ucm_PtpOc_DefaultDs3Reg, temp_data))
                    {
                        ui->PtpOcDefaultDsDomainValue->setText(QString("0x%1").arg(((temp_data >> 0) & 0x000000FF), 2, 16, QLatin1Char('0')));
                        ui->PtpOcDefaultDsPriority2Value->setText(QString("0x%1").arg(((temp_data >> 16) & 0x000000FF), 2, 16, QLatin1Char('0')));
                        ui->PtpOcDefaultDsPriority1Value->setText(QString("0x%1").arg(((temp_data >> 24) & 0x000000FF), 2, 16, QLatin1Char('0')));

                    }
                    else
                    {
                        ui->PtpOcDefaultDsDomainValue->setText("NA");
                        ui->PtpOcDefaultDsPriority1Value->setText("NA");
                        ui->PtpOcDefaultDsPriority2Value->setText("NA");
                    }

                    // variance, accuracy ,class
                    if (0 == ucm->com_lib.read_reg(temp_addr + Ucm_PtpOc_DefaultDs4Reg, temp_data))
                    {
                        ui->PtpOcDefaultDsVarianceValue->setText(QString("0x%1").arg(((temp_data >> 0) & 0x0000FFFF), 4, 16, QLatin1Char('0')));
                        ui->PtpOcDefaultDsAccuracyValue->setText(QString::number(((temp_data >> 16) & 0x000000FF)));
                        ui->PtpOcDefaultDsClassValue->setText(QString("0x%1").arg(((temp_data >> 24) & 0x000000FF), 2, 16, QLatin1Char('0')));
                    }
                    else
                    {
                        ui->PtpOcDefaultDsVarianceValue->setText("NA");
                        ui->PtpOcDefaultDsAccuracyValue->setText("NA");
                        ui->PtpOcDefaultDsClassValue->setText("NA");
                    }

                    // short id
                    if (0 == ucm->com_lib.read_reg(temp_addr + Ucm_PtpOc_DefaultDs5Reg, temp_data))
                    {
                        ui->PtpOcDefaultDsShortIdValue->setText(QString("0x%1").arg(temp_data, 4, 16, QLatin1Char('0')));
                    }
                    else
                    {
                        ui->PtpOcDefaultDsShortIdValue->setText("NA");
                    }

                    // inaccuracy
                    if (0 == ucm->com_lib.read_reg(temp_addr + Ucm_PtpOc_DefaultDs6Reg, temp_data))
                    {
                        ui->PtpOcDefaultDsInaccuracyValue->setText(QString::number(temp_data));
                    }
                    else
                    {
                        ui->PtpOcDefaultDsInaccuracyValue->setText("NA");
                    }

                    // nr of ports
                    if (0 == ucm->com_lib.read_reg(temp_addr + Ucm_PtpOc_DefaultDs7Reg, temp_data))
                    {
                        ui->PtpOcDefaultDsNrOfPortsValue->setText(QString::number(temp_data));
                    }
                    else
                    {
                        ui->PtpOcDefaultDsNrOfPortsValue->setText("NA");
                    }

                    break;
                }
                else if (i == 9)
                {
                    cout << "ERROR: " << "read did not complete" << endl;
                    ui->PtpOcDefaultDsClockIdValue->setText("NA");
                    ui->PtpOcDefaultDsDomainValue->setText("NA");
                    ui->PtpOcDefaultDsPriority1Value->setText("NA");
                    ui->PtpOcDefaultDsPriority2Value->setText("NA");
                    ui->PtpOcDefaultDsVarianceValue->setText("NA");
                    ui->PtpOcDefaultDsAccuracyValue->setText("NA");
                    ui->PtpOcDefaultDsClassValue->setText("NA");
                    ui->PtpOcDefaultDsShortIdValue->setText("NA");
                    ui->PtpOcDefaultDsInaccuracyValue->setText("NA");
                    ui->PtpOcDefaultDsNrOfPortsValue->setText("NA");
                }

            }
            else
            {
                ui->PtpOcDefaultDsClockIdValue->setText("NA");
                ui->PtpOcDefaultDsDomainValue->setText("NA");
                ui->PtpOcDefaultDsPriority1Value->setText("NA");
                ui->PtpOcDefaultDsPriority2Value->setText("NA");
                ui->PtpOcDefaultDsVarianceValue->setText("NA");
                ui->PtpOcDefaultDsAccuracyValue->setText("NA");
                ui->PtpOcDefaultDsClassValue->setText("NA");
                ui->PtpOcDefaultDsShortIdValue->setText("NA");
                ui->PtpOcDefaultDsInaccuracyValue->setText("NA");
                ui->PtpOcDefaultDsNrOfPortsValue->setText("NA");
            }
        }
    }
    else
    {
        ui->PtpOcDefaultDsClockIdValue->setText("NA");
        ui->PtpOcDefaultDsDomainValue->setText("NA");
        ui->PtpOcDefaultDsPriority1Value->setText("NA");
        ui->PtpOcDefaultDsPriority2Value->setText("NA");
        ui->PtpOcDefaultDsVarianceValue->setText("NA");
        ui->PtpOcDefaultDsAccuracyValue->setText("NA");
        ui->PtpOcDefaultDsClassValue->setText("NA");
        ui->PtpOcDefaultDsShortIdValue->setText("NA");
        ui->PtpOcDefaultDsInaccuracyValue->setText("NA");
        ui->PtpOcDefaultDsNrOfPortsValue->setText("NA");
    }

    //********************************
    // port dataset
    //********************************
    temp_data = 0x40000000;
    if (0 == ucm->com_lib.write_reg(temp_addr + Ucm_PtpOc_PortDsControlReg, temp_data))
    {
        for (int i = 0; i < 10; i++)
        {
            if(0 == ucm->com_lib.read_reg(temp_addr + Ucm_PtpOc_PortDsControlReg, temp_data))
            {
                if ((temp_data & 0x80000000) != 0)
                {
                    temp_string = ui->PtpOcDelayMechanismValue->currentText();
                    if ((temp_string == "E2E") or (temp_string == "E2E Unicast"))
                    {
                        // end to end delay
                        ui->PtpOcPortDsPeerDelayValue->setText("NA");
                    }
                    else if (0 == ucm->com_lib.read_reg(temp_addr + Ucm_PtpOc_PortDs1Reg, temp_data))
                    {
                        // peer delay
                        temp_delay = temp_data;
                        temp_delay = temp_delay << 32;
                        if (0 == ucm->com_lib.read_reg(temp_addr + Ucm_PtpOc_PortDs2Reg, temp_data))
                        {
                            temp_delay |= temp_data;
                            temp_signed_delay = (long long)temp_delay;
                            temp_signed_delay = temp_signed_delay >> 16;
                            ui->PtpOcPortDsPeerDelayValue->setText(QString::number(temp_signed_delay));

                            if (true == ptp_oc_timer->isActive())
                            {
                                ptp_oc_delay_series->append(ptp_oc_delay_number_of_points, temp_signed_delay);

                                if (ptp_oc_delay_number_of_points < 20)
                                {
                                    ptp_oc_delay_number_of_points++;
                                }
                                else
                                {
                                    for (int j = 1; j < ptp_oc_delay_series->count(); j++)
                                    {
                                        QPointF temp_point = ptp_oc_delay_series->at(j);
                                        ptp_oc_delay_series->replace(j, (j-1), temp_point.y());
                                    }
                                    ptp_oc_delay_series->remove(0);
                                }

                                temp_min = 0;
                                temp_max = 0;
                                for (int j = 0; j < ptp_oc_delay_series->count(); j++)
                                {
                                    QPointF temp_point = ptp_oc_delay_series->at(j);
                                    if (j == 0)
                                    {
                                        temp_min = temp_point.y();
                                        temp_max = temp_point.y();
                                    }
                                    if (temp_min > temp_point.y())
                                    {
                                        temp_min = temp_point.y();
                                    }
                                    if (temp_max < temp_point.y())
                                    {
                                        temp_max = temp_point.y();
                                    }
                                }
                                temp_max = ((temp_max/100)+1)*100;
                                temp_min = ((temp_min/100)-1)*100;
                                //if (temp_min < 0)
                                //{
                                //    temp_min = 0;
                                //}
                                ptp_oc_delay_chart->axisY()->setMin(temp_min);
                                ptp_oc_delay_chart->axisY()->setMax(temp_max);

                                ptp_oc_delay_chart->show();
                            }
                        }
                        else
                        {
                            ui->PtpOcPortDsPeerDelayValue->setText("NA");
                        }
                    }
                    else
                    {
                        ui->PtpOcPortDsPeerDelayValue->setText("NA");
                    }

                    // state
                    if (0 == ucm->com_lib.read_reg(temp_addr + Ucm_PtpOc_PortDs3Reg, temp_data))
                    {
                        switch (temp_data)
                        {
                        case 0x0000001:
                            ui->PtpOcPortDsStateValue->setText("INITIALIZING");
                            break;
                        case 0x00000002:
                            ui->PtpOcPortDsStateValue->setText("FAULTY");
                            break;
                        case 0x00000003:
                            ui->PtpOcPortDsStateValue->setText("DISABLED");
                            break;
                        case 0x00000004:
                            ui->PtpOcPortDsStateValue->setText("LISTENING");
                            break;
                        case 0x00000005:
                            ui->PtpOcPortDsStateValue->setText("PREMASTER");
                            break;
                        case 0x00000006:
                            ui->PtpOcPortDsStateValue->setText("MASTER");
                            break;
                        case 0x00000007:
                            ui->PtpOcPortDsStateValue->setText("PASSIVE");
                            break;
                        case 0x00000008:
                            ui->PtpOcPortDsStateValue->setText("UNCALIBRATED");
                            break;
                        case 0x00000009:
                            ui->PtpOcPortDsStateValue->setText("SLAVE");
                            break;
                        default:
                            ui->PtpOcPortDsStateValue->setText("NA");
                            break;
                        }                    }
                    else
                    {
                        ui->PtpOcPortDsStateValue->setText("NA");
                    }

                    // pdelay and delay req log msg interval and max delay
                    if (0 == ucm->com_lib.read_reg(temp_addr + Ucm_PtpOc_PortDs4Reg, temp_data))
                    {
                        ui->PtpOcPortDsPDelayReqLogMsgIntervalValue->setText(QString::number((signed char)(temp_data & 0x000000FF)));
                        ui->PtpOcPortDsDelayReqLogMsgIntervalValue->setText(QString::number((signed char)((temp_data >> 8) & 0x000000FF)));
                        ui->PtpOcPortDsDelayReceiptTimeoutValue->setText(QString::number(((temp_data >> 16) & 0x000000FF)));
                    }
                    else
                    {
                        ui->PtpOcPortDsPDelayReqLogMsgIntervalValue->setText("NA");
                        ui->PtpOcPortDsDelayReqLogMsgIntervalValue->setText("NA");
                        ui->PtpOcPortDsDelayReceiptTimeoutValue->setText("NA");
                    }

                    // announce log msg interval and announce receipt timeout
                    if (0 == ucm->com_lib.read_reg(temp_addr + Ucm_PtpOc_PortDs5Reg, temp_data))
                    {
                        ui->PtpOcPortDsAnnounceLogMsgIntervalValue->setText(QString::number((signed char)(temp_data & 0x000000FF)));
                        ui->PtpOcPortDsAnnounceReceiptTimeoutValue->setText(QString::number(((temp_data >> 8) & 0x000000FF)));

                    }
                    else
                    {
                        ui->PtpOcPortDsAnnounceLogMsgIntervalValue->setText("NA");
                        ui->PtpOcPortDsAnnounceReceiptTimeoutValue->setText("NA");
                    }

                    // sync log msg interval and sync receipt timeout
                    if (0 == ucm->com_lib.read_reg(temp_addr + Ucm_PtpOc_PortDs6Reg, temp_data))
                    {
                        ui->PtpOcPortDsSyncLogMsgIntervalValue->setText(QString::number((signed char)(temp_data & 0x000000FF)));
                        ui->PtpOcPortDsSyncReceiptTimeoutValue->setText(QString::number(((temp_data >> 8) & 0x000000FF)));
                    }
                    else
                    {
                        ui->PtpOcPortDsSyncLogMsgIntervalValue->setText("NA");
                        ui->PtpOcPortDsSyncReceiptTimeoutValue->setText("NA");
                    }

                    // asymmetry
                    if (0 == ucm->com_lib.read_reg(temp_addr + Ucm_PtpOc_PortDs7Reg, temp_data))
                    {
                        ui->PtpOcPortDsAsymmetryValue->setText(QString::number((signed int)temp_data));

                    }
                    else
                    {
                        ui->PtpOcPortDsAsymmetryValue->setText("NA");
                    }

                    // max pdelay
                    if (0 == ucm->com_lib.read_reg(temp_addr + Ucm_PtpOc_PortDs8Reg, temp_data))
                    {
                        ui->PtpOcPortDsMaxPeerDelayValue->setText(QString::number(temp_data));
                    }
                    else
                    {
                        ui->PtpOcPortDsMaxPeerDelayValue->setText("NA");
                    }

                    break;
                }
                else if (i == 9)
                {
                    cout << "ERROR: " << "read did not complete" << endl;
                    ui->PtpOcPortDsPeerDelayValue->setText("NA");
                    ui->PtpOcPortDsStateValue->setText("NA");
                    ui->PtpOcPortDsAsymmetryValue->setText("NA");
                    ui->PtpOcPortDsMaxPeerDelayValue->setText("NA");
                    ui->PtpOcPortDsPDelayReqLogMsgIntervalValue->setText("NA");
                    ui->PtpOcPortDsDelayReqLogMsgIntervalValue->setText("NA");
                    ui->PtpOcPortDsDelayReceiptTimeoutValue->setText("NA");
                    ui->PtpOcPortDsAnnounceLogMsgIntervalValue->setText("NA");
                    ui->PtpOcPortDsAnnounceReceiptTimeoutValue->setText("NA");
                    ui->PtpOcPortDsSyncLogMsgIntervalValue->setText("NA");
                    ui->PtpOcPortDsSyncReceiptTimeoutValue->setText("NA");
                }

            }
            else
            {
                ui->PtpOcPortDsPeerDelayValue->setText("NA");
                ui->PtpOcPortDsStateValue->setText("NA");
                ui->PtpOcPortDsAsymmetryValue->setText("NA");
                ui->PtpOcPortDsMaxPeerDelayValue->setText("NA");
                ui->PtpOcPortDsPDelayReqLogMsgIntervalValue->setText("NA");
                ui->PtpOcPortDsDelayReqLogMsgIntervalValue->setText("NA");
                ui->PtpOcPortDsDelayReceiptTimeoutValue->setText("NA");
                ui->PtpOcPortDsAnnounceLogMsgIntervalValue->setText("NA");
                ui->PtpOcPortDsAnnounceReceiptTimeoutValue->setText("NA");
                ui->PtpOcPortDsSyncLogMsgIntervalValue->setText("NA");
                ui->PtpOcPortDsSyncReceiptTimeoutValue->setText("NA");
            }
        }
    }
    else
    {
        ui->PtpOcPortDsPeerDelayValue->setText("NA");
        ui->PtpOcPortDsStateValue->setText("NA");
        ui->PtpOcPortDsAsymmetryValue->setText("NA");
        ui->PtpOcPortDsMaxPeerDelayValue->setText("NA");
        ui->PtpOcPortDsPDelayReqLogMsgIntervalValue->setText("NA");
        ui->PtpOcPortDsDelayReqLogMsgIntervalValue->setText("NA");
        ui->PtpOcPortDsDelayReceiptTimeoutValue->setText("NA");
        ui->PtpOcPortDsAnnounceLogMsgIntervalValue->setText("NA");
        ui->PtpOcPortDsAnnounceReceiptTimeoutValue->setText("NA");
        ui->PtpOcPortDsSyncLogMsgIntervalValue->setText("NA");
        ui->PtpOcPortDsSyncReceiptTimeoutValue->setText("NA");
    }

    //********************************
    // current dataset
    //********************************
    temp_data = 0x40000000;
    if (0 == ucm->com_lib.write_reg(temp_addr + Ucm_PtpOc_CurrentDsControlReg, temp_data))
    {
        for (int i = 0; i < 10; i++)
        {
            if(0 == ucm->com_lib.read_reg(temp_addr + Ucm_PtpOc_CurrentDsControlReg, temp_data))
            {
                if ((temp_data & 0x80000000) != 0)
                {

                    // steps removed
                    if (0 == ucm->com_lib.read_reg(temp_addr + Ucm_PtpOc_CurrentDs1Reg, temp_data))
                    {
                        ui->PtpOcCurrentDsStepsRemovedValue->setText(QString::number(temp_data & 0xFFFF));
                    }
                    else
                    {
                        ui->PtpOcCurrentDsStepsRemovedValue->setText("NA");
                    }

                    // offset
                    if (0 == ucm->com_lib.read_reg(temp_addr + Ucm_PtpOc_CurrentDs2Reg, temp_data))
                    {
                        temp_offset = temp_data;
                        temp_offset = temp_offset << 32;
                        if (0 == ucm->com_lib.read_reg(temp_addr + Ucm_PtpOc_CurrentDs3Reg, temp_data))
                        {
                            temp_offset |= temp_data;

                            if ((temp_offset & 0x8000000000000000) != 0)
                            {
                                temp_offset = (0xFFFF000000000000 | (temp_offset >> 16));
                                temp_signed_offset = (long long)temp_offset;
                            }
                            else
                            {
                                temp_offset = (0x0000FFFFFFFFFFFF & (temp_offset >> 16));
                                temp_signed_offset = (long long)temp_offset;
                            }

                            if (temp_signed_offset == -4294967296) // negativ 0
                            {
                                temp_signed_offset = 0;
                            }

                            // limit to one second in display
                            if (temp_signed_offset >= 100000)
                            {
                                temp_signed_offset = 100000;
                            }
                            else if (temp_signed_offset <= -100000)
                            {
                                temp_signed_offset = -100000;
                            }

                            ui->PtpOcCurrentDsOffsetValue->setText(QString::number(temp_signed_offset));

                            if (true == ptp_oc_timer->isActive())
                            {

                                ptp_oc_offset_series->append(ptp_oc_offset_number_of_points, temp_signed_offset);

                                if (ptp_oc_offset_number_of_points < 20)
                                {
                                    ptp_oc_offset_number_of_points++;
                                }
                                else
                                {
                                    for (int j = 1; j < ptp_oc_offset_series->count(); j++)
                                    {
                                        QPointF temp_point = ptp_oc_offset_series->at(j);
                                        ptp_oc_offset_series->replace(j, (j-1), temp_point.y());
                                    }
                                    ptp_oc_offset_series->remove(0);
                                }

                                temp_min = 0;
                                temp_max = 0;
                                for (int j = 0; j < ptp_oc_offset_series->count(); j++)
                                {
                                    QPointF temp_point = ptp_oc_offset_series->at(j);
                                    if (j == 0)
                                    {
                                        temp_min = temp_point.y();
                                        temp_max = temp_point.y();
                                    }
                                    if (temp_min > temp_point.y())
                                    {
                                        temp_min = temp_point.y();
                                    }
                                    if (temp_max < temp_point.y())
                                    {
                                        temp_max = temp_point.y();
                                    }
                                }
                                temp_max = ((temp_max/100)+1)*100;
                                temp_min = ((temp_min/100)-1)*100;
                                if (temp_max > 100000)
                                {
                                    temp_max = 100000;
                                }
                                if (temp_min < -100000)
                                {
                                    temp_min = -100000;
                                }
                                ptp_oc_offset_chart->axisY()->setMin(temp_min);
                                ptp_oc_offset_chart->axisY()->setMax(temp_max);

                                ptp_oc_offset_chart->show();
                            }
                        }
                        else
                        {
                            ui->PtpOcCurrentDsOffsetValue->setText("NA");
                        }
                    }
                    else
                    {
                        ui->PtpOcCurrentDsOffsetValue->setText("NA");
                    }


                    temp_string = ui->PtpOcDelayMechanismValue->currentText();
                    if (temp_string == "P2P")
                    {
                        // peer delay
                        ui->PtpOcCurrentDsDelayValue->setText("NA");
                    }
                    else if (0 == ucm->com_lib.read_reg(temp_addr + Ucm_PtpOc_CurrentDs4Reg, temp_data))
                    {
                        // end to end delay
                        temp_delay = temp_data;
                        temp_delay = temp_delay << 32;
                        if (0 == ucm->com_lib.read_reg(temp_addr + Ucm_PtpOc_CurrentDs5Reg, temp_data))
                        {
                            temp_delay |= temp_data;
                            temp_signed_delay = (long long)temp_delay;
                            temp_signed_delay = temp_signed_delay >> 16;
                            ui->PtpOcCurrentDsDelayValue->setText(QString::number(temp_signed_delay));

                            if (true == ptp_oc_timer->isActive())
                            {
                                ptp_oc_delay_series->append(ptp_oc_delay_number_of_points, temp_signed_delay);

                                if (ptp_oc_delay_number_of_points < 20)
                                {
                                    ptp_oc_delay_number_of_points++;
                                }
                                else
                                {
                                    for (int j = 1; j < ptp_oc_delay_series->count(); j++)
                                    {
                                        QPointF temp_point = ptp_oc_delay_series->at(j);
                                        ptp_oc_delay_series->replace(j, (j-1), temp_point.y());
                                    }
                                    ptp_oc_delay_series->remove(0);
                                }

                                temp_min = 0;
                                temp_max = 0;
                                for (int j = 0; j < ptp_oc_delay_series->count(); j++)
                                {
                                    QPointF temp_point = ptp_oc_delay_series->at(j);
                                    if (j == 0)
                                    {
                                        temp_min = temp_point.y();
                                        temp_max = temp_point.y();
                                    }
                                    if (temp_min > temp_point.y())
                                    {
                                        temp_min = temp_point.y();
                                    }
                                    if (temp_max < temp_point.y())
                                    {
                                        temp_max = temp_point.y();
                                    }
                                }
                                temp_max = ((temp_max/100)+1)*100;
                                temp_min = ((temp_min/100)-1)*100;
                                //if (temp_min < 0)
                                //{
                                //    temp_min = 0;
                                //}
                                ptp_oc_delay_chart->axisY()->setMin(temp_min);
                                ptp_oc_delay_chart->axisY()->setMax(temp_max);

                                ptp_oc_delay_chart->show();
                            }
                        }
                        else
                        {
                            ui->PtpOcCurrentDsDelayValue->setText("NA");
                        }
                    }
                    else
                    {
                        ui->PtpOcCurrentDsDelayValue->setText("NA");
                    }

                    break;
                }
                else if (i == 9)
                {
                    cout << "ERROR: " << "read did not complete" << endl;
                    ui->PtpOcCurrentDsStepsRemovedValue->setText("NA");
                    ui->PtpOcCurrentDsOffsetValue->setText("NA");
                }

            }
            else
            {
                ui->PtpOcCurrentDsStepsRemovedValue->setText("NA");
                ui->PtpOcCurrentDsOffsetValue->setText("NA");
            }
        }
    }
    else
    {
        ui->PtpOcCurrentDsStepsRemovedValue->setText("NA");
        ui->PtpOcCurrentDsOffsetValue->setText("NA");
    }

    //********************************
    // parent dataset
    //********************************
    temp_data = 0x40000000;
    if (0 == ucm->com_lib.write_reg(temp_addr + Ucm_PtpOc_ParentDsControlReg, temp_data))
    {
        for (int i = 0; i < 10; i++)
        {
            if(0 == ucm->com_lib.read_reg(temp_addr + Ucm_PtpOc_ParentDsControlReg, temp_data))
            {
                if ((temp_data & 0x80000000) != 0)
                {

                    // parent clock id and port id
                    temp_string.clear();
                    if (0 == ucm->com_lib.read_reg(temp_addr + Ucm_PtpOc_ParentDs1Reg, temp_data))
                    {
                        temp_string.append(QString("%1").arg(((temp_data >> 0) & 0x000000FF), 2, 16, QLatin1Char('0')));
                        temp_string.append(":");
                        temp_string.append(QString("%1").arg(((temp_data >> 8) & 0x000000FF), 2, 16, QLatin1Char('0')));
                        temp_string.append(":");
                        temp_string.append(QString("%1").arg(((temp_data >> 16) & 0x000000FF), 2, 16, QLatin1Char('0')));
                        temp_string.append(":");
                        temp_string.append(QString("%1").arg(((temp_data >> 24) & 0x000000FF), 2, 16, QLatin1Char('0')));
                        temp_string.append(":");
                        if (0 == ucm->com_lib.read_reg(temp_addr + Ucm_PtpOc_ParentDs2Reg, temp_data))
                        {
                            temp_string.append(QString("%1").arg(((temp_data >> 0) & 0x000000FF), 2, 16, QLatin1Char('0')));
                            temp_string.append(":");
                            temp_string.append(QString("%1").arg(((temp_data >> 8) & 0x000000FF), 2, 16, QLatin1Char('0')));
                            temp_string.append(":");
                            temp_string.append(QString("%1").arg(((temp_data >> 16) & 0x000000FF), 2, 16, QLatin1Char('0')));
                            temp_string.append(":");
                            temp_string.append(QString("%1").arg(((temp_data >> 24) & 0x000000FF), 2, 16, QLatin1Char('0')));
                            temp_string.append(".");
                            if (0 == ucm->com_lib.read_reg(temp_addr + Ucm_PtpOc_ParentDs3Reg, temp_data))
                            {
                                temp_string.append(QString("%1").arg(((temp_data >> 0) & 0x0000FFFF), 4, 16, QLatin1Char('0')));
                                ui->PtpOcParentDsParentClockIdValue->setText(temp_string);
                            }
                            else
                            {
                                ui->PtpOcParentDsParentClockIdValue->setText("NA");
                            }
                        }
                        else
                        {
                            ui->PtpOcParentDsParentClockIdValue->setText("NA");
                        }
                    }
                    else
                    {
                        ui->PtpOcParentDsParentClockIdValue->setText("NA");
                    }

                    // gm clock id
                    temp_string.clear();
                    if (0 == ucm->com_lib.read_reg(temp_addr + Ucm_PtpOc_ParentDs4Reg, temp_data))
                    {
                        temp_string.append(QString("%1").arg(((temp_data >> 0) & 0x000000FF), 2, 16, QLatin1Char('0')));
                        temp_string.append(":");
                        temp_string.append(QString("%1").arg(((temp_data >> 8) & 0x000000FF), 2, 16, QLatin1Char('0')));
                        temp_string.append(":");
                        temp_string.append(QString("%1").arg(((temp_data >> 16) & 0x000000FF), 2, 16, QLatin1Char('0')));
                        temp_string.append(":");
                        temp_string.append(QString("%1").arg(((temp_data >> 24) & 0x000000FF), 2, 16, QLatin1Char('0')));
                        temp_string.append(":");
                        if (0 == ucm->com_lib.read_reg(temp_addr + Ucm_PtpOc_ParentDs5Reg, temp_data))
                        {
                            temp_string.append(QString("%1").arg(((temp_data >> 0) & 0x000000FF), 2, 16, QLatin1Char('0')));
                            temp_string.append(":");
                            temp_string.append(QString("%1").arg(((temp_data >> 8) & 0x000000FF), 2, 16, QLatin1Char('0')));
                            temp_string.append(":");
                            temp_string.append(QString("%1").arg(((temp_data >> 16) & 0x000000FF), 2, 16, QLatin1Char('0')));
                            temp_string.append(":");
                            temp_string.append(QString("%1").arg(((temp_data >> 24) & 0x000000FF), 2, 16, QLatin1Char('0')));
                            ui->PtpOcParentDsGmClockIdValue->setText(temp_string);
                        }
                        else
                        {
                            ui->PtpOcParentDsGmClockIdValue->setText("NA");
                        }
                    }
                    else
                    {
                        ui->PtpOcParentDsGmClockIdValue->setText("NA");
                    }

                    // gm priority 1 & 2
                    if (0 == ucm->com_lib.read_reg(temp_addr + Ucm_PtpOc_ParentDs3Reg, temp_data))
                    {
                        ui->PtpOcParentDsGmPriority2Value->setText(QString("0x%1").arg(((temp_data >> 16) & 0x000000FF), 2, 16, QLatin1Char('0')));
                        ui->PtpOcParentDsGmPriority1Value->setText(QString("0x%1").arg(((temp_data >> 24) & 0x000000FF), 2, 16, QLatin1Char('0')));
                    }
                    else
                    {
                        ui->PtpOcParentDsGmPriority1Value->setText("NA");
                        ui->PtpOcParentDsGmPriority2Value->setText("NA");
                    }

                    // variance, accuracy ,class
                    if (0 == ucm->com_lib.read_reg(temp_addr + Ucm_PtpOc_ParentDs6Reg, temp_data))
                    {
                        ui->PtpOcParentDsGmVarianceValue->setText(QString("0x%1").arg(((temp_data >> 0) & 0x0000FFFF), 4, 16, QLatin1Char('0')));
                        ui->PtpOcParentDsGmAccuracyValue->setText(QString::number(((temp_data >> 16) & 0x000000FF)));
                        ui->PtpOcParentDsGmClassValue->setText(QString("0x%1").arg(((temp_data >> 24) & 0x000000FF), 2, 16, QLatin1Char('0')));
                    }
                    else
                    {
                        ui->PtpOcParentDsGmVarianceValue->setText("NA");
                        ui->PtpOcParentDsGmAccuracyValue->setText("NA");
                        ui->PtpOcParentDsGmClassValue->setText("NA");
                    }

                    // gm short id
                    if (0 == ucm->com_lib.read_reg(temp_addr + Ucm_PtpOc_ParentDs7Reg, temp_data))
                    {
                        ui->PtpOcParentDsGmShortIdValue->setText(QString("0x%1").arg(temp_data, 4, 16, QLatin1Char('0')));
                    }
                    else
                    {
                        ui->PtpOcParentDsGmShortIdValue->setText("NA");
                    }

                    // gm inaccuracy
                    if (0 == ucm->com_lib.read_reg(temp_addr + Ucm_PtpOc_ParentDs8Reg, temp_data))
                    {
                        ui->PtpOcParentDsGmInaccuracyValue->setText(QString::number(temp_data));
                    }
                    else
                    {
                        ui->PtpOcParentDsGmInaccuracyValue->setText("NA");
                    }

                    // nw inaccuracy
                    if (0 == ucm->com_lib.read_reg(temp_addr + Ucm_PtpOc_ParentDs9Reg, temp_data))
                    {
                        ui->PtpOcParentDsNwInaccuracyValue->setText(QString::number(temp_data));
                    }
                    else
                    {
                        ui->PtpOcParentDsNwInaccuracyValue->setText("NA");
                    }

                    break;
                }
                else if (i == 9)
                {
                    cout << "ERROR: " << "read did not complete" << endl;
                    ui->PtpOcParentDsParentClockIdValue->setText("NA");
                    ui->PtpOcParentDsGmClockIdValue->setText("NA");
                    ui->PtpOcParentDsGmPriority1Value->setText("NA");
                    ui->PtpOcParentDsGmPriority2Value->setText("NA");
                    ui->PtpOcParentDsGmAccuracyValue->setText("NA");
                    ui->PtpOcParentDsGmClassValue->setText("NA");
                    ui->PtpOcParentDsGmShortIdValue->setText("NA");
                    ui->PtpOcParentDsGmInaccuracyValue->setText("NA");
                    ui->PtpOcParentDsNwInaccuracyValue->setText("NA");

                }

            }
            else
            {
                ui->PtpOcParentDsParentClockIdValue->setText("NA");
                ui->PtpOcParentDsGmClockIdValue->setText("NA");
                ui->PtpOcParentDsGmPriority1Value->setText("NA");
                ui->PtpOcParentDsGmPriority2Value->setText("NA");
                ui->PtpOcParentDsGmAccuracyValue->setText("NA");
                ui->PtpOcParentDsGmClassValue->setText("NA");
                ui->PtpOcParentDsGmShortIdValue->setText("NA");
                ui->PtpOcParentDsGmInaccuracyValue->setText("NA");
                ui->PtpOcParentDsNwInaccuracyValue->setText("NA");

            }
        }
    }
    else
    {
        ui->PtpOcParentDsParentClockIdValue->setText("NA");
        ui->PtpOcParentDsGmClockIdValue->setText("NA");
        ui->PtpOcParentDsGmPriority1Value->setText("NA");
        ui->PtpOcParentDsGmPriority2Value->setText("NA");
        ui->PtpOcParentDsGmAccuracyValue->setText("NA");
        ui->PtpOcParentDsGmClassValue->setText("NA");
        ui->PtpOcParentDsGmShortIdValue->setText("NA");
        ui->PtpOcParentDsGmInaccuracyValue->setText("NA");
        ui->PtpOcParentDsNwInaccuracyValue->setText("NA");

    }

    //********************************
    // time properties dataset
    //********************************
    ui->PtpOcTimePropertiesDsSetLocalPropertiesCheckBox->setChecked(false);
    temp_data = 0x40000000;
    if (0 == ucm->com_lib.write_reg(temp_addr + Ucm_PtpOc_TimePropertiesDsControlReg, temp_data))
    {
        for (int i = 0; i < 10; i++)
        {
            if(0 == ucm->com_lib.read_reg(temp_addr + Ucm_PtpOc_TimePropertiesDsControlReg, temp_data))
            {
                if ((temp_data & 0x80000000) != 0)
                {

                    // time source, ptp timescale, freq traceable, time traceable, lep61, leap 59, ut offset val, utc offset
                    if (0 == ucm->com_lib.read_reg(temp_addr + Ucm_PtpOc_TimePropertiesDs1Reg, temp_data))
                    {
                        ui->PtpOcTimePropertiesDsTimeSourceValue->setText(QString("0x%1").arg(((temp_data >> 0) & 0x000000FF), 2, 16, QLatin1Char('0')));
                        if ((temp_data & 0x00000100) != 0)
                        {
                            ui->PtpOcTimePropertiesDsPtpTimescaleCheckBox->setChecked(true);
                        }
                        else
                        {
                            ui->PtpOcTimePropertiesDsPtpTimescaleCheckBox->setChecked(false);
                        }
                        if ((temp_data & 0x00000200) != 0)
                        {
                            ui->PtpOcTimePropertiesDsFreqTraceableCheckBox->setChecked(true);
                        }
                        else
                        {
                            ui->PtpOcTimePropertiesDsFreqTraceableCheckBox->setChecked(false);
                        }
                        if ((temp_data & 0x00000400) != 0)
                        {
                            ui->PtpOcTimePropertiesDsTimeTraceableCheckBox->setChecked(true);
                        }
                        else
                        {
                            ui->PtpOcTimePropertiesDsTimeTraceableCheckBox->setChecked(false);
                        }
                        if ((temp_data & 0x00000800) != 0)
                        {
                            ui->PtpOcTimePropertiesDsLeap61CheckBox->setChecked(true);
                        }
                        else
                        {
                            ui->PtpOcTimePropertiesDsLeap61CheckBox->setChecked(false);
                        }
                        if ((temp_data & 0x00001000) != 0)
                        {
                            ui->PtpOcTimePropertiesDsLeap59CheckBox->setChecked(true);
                        }
                        else
                        {
                            ui->PtpOcTimePropertiesDsLeap59CheckBox->setChecked(false);
                        }
                        if ((temp_data & 0x00002000) != 0)
                        {
                            ui->PtpOcTimePropertiesDsUtcOffsetValCheckBox->setChecked(true);
                        }
                        else
                        {
                            ui->PtpOcTimePropertiesDsUtcOffsetValCheckBox->setChecked(false);
                        }

                        ui->PtpOcTimePropertiesDsUtcOffsetValue->setText(QString::number((signed short)((temp_data >> 16) & 0x0000FFFF)));
                    }
                    else
                    {
                        ui->PtpOcTimePropertiesDsTimeSourceValue->setText("NA");
                        ui->PtpOcTimePropertiesDsPtpTimescaleCheckBox->setChecked(false);
                        ui->PtpOcTimePropertiesDsFreqTraceableCheckBox->setChecked(false);
                        ui->PtpOcTimePropertiesDsTimeTraceableCheckBox->setChecked(false);
                        ui->PtpOcTimePropertiesDsLeap59CheckBox->setChecked(false);
                        ui->PtpOcTimePropertiesDsLeap61CheckBox->setChecked(false);
                        ui->PtpOcTimePropertiesDsUtcOffsetValCheckBox->setChecked(false);
                        ui->PtpOcTimePropertiesDsUtcOffsetValue->setText("NA");
                    }

                    // current offset
                    if (0 == ucm->com_lib.read_reg(temp_addr + Ucm_PtpOc_TimePropertiesDs2Reg, temp_data))
                    {
                        ui->PtpOcTimePropertiesDsCurrentOffsetValue->setText(QString::number((int)temp_data));
                    }
                    else
                    {
                        ui->PtpOcTimePropertiesDsCurrentOffsetValue->setText("NA");
                    }

                    // jump seconds
                    if (0 == ucm->com_lib.read_reg(temp_addr + Ucm_PtpOc_TimePropertiesDs3Reg, temp_data))
                    {
                        ui->PtpOcTimePropertiesDsJumpSecondsValue->setText(QString::number((int)temp_data));
                    }
                    else
                    {
                        ui->PtpOcTimePropertiesDsJumpSecondsValue->setText("NA");
                    }

                    // next jump
                    if (0 == ucm->com_lib.read_reg(temp_addr + Ucm_PtpOc_TimePropertiesDs4Reg, temp_data))
                    {
                        temp_next_jump = temp_data;
                        temp_next_jump = temp_next_jump << 32;
                        if (0 == ucm->com_lib.read_reg(temp_addr + Ucm_PtpOc_TimePropertiesDs5Reg, temp_data))
                        {
                            temp_next_jump |= temp_data;
                            ui->PtpOcTimePropertiesDsNextJumpValue->setText(QString::number(temp_next_jump));
                        }
                        else
                        {
                            ui->PtpOcTimePropertiesDsNextJumpValue->setText("NA");
                        }
                    }
                    else
                    {
                        ui->PtpOcTimePropertiesDsNextJumpValue->setText("NA");
                    }

                    // display name
                    temp_string.clear();
                    if (0 == ucm->com_lib.read_reg(temp_addr + Ucm_PtpOc_TimePropertiesDs6Reg, temp_length))
                    {
                        for (int j=0; j<3; j++)
                        {
                            if (0 == ucm->com_lib.read_reg(temp_addr + Ucm_PtpOc_TimePropertiesDs7Reg + (j*4), temp_data))
                            {
                                temp_string.append((QChar)((temp_data >> 0) & 0x000000FF));
                                temp_string.append((QChar)((temp_data >> 8) & 0x000000FF));
                                temp_string.append((QChar)((temp_data >> 16) & 0x000000FF));
                                temp_string.append((QChar)((temp_data >> 24) & 0x000000FF));
                            }
                            else
                            {
                                temp_string.clear();
                                temp_string.append("NA");
                                break;
                            }
                        }
                        temp_string.truncate(temp_length);
                        ui->PtpOcTimePropertiesDsDisplayNameValue->setText(temp_string);
                    }
                    else
                    {
                        ui->PtpOcTimePropertiesDsDisplayNameValue->setText("NA");
                    }
                    break;
                }
                else if (i == 9)
                {
                    cout << "ERROR: " << "read did not complete" << endl;
                    ui->PtpOcTimePropertiesDsTimeSourceValue->setText("NA");
                    ui->PtpOcTimePropertiesDsPtpTimescaleCheckBox->setChecked(false);
                    ui->PtpOcTimePropertiesDsFreqTraceableCheckBox->setChecked(false);
                    ui->PtpOcTimePropertiesDsTimeTraceableCheckBox->setChecked(false);
                    ui->PtpOcTimePropertiesDsLeap59CheckBox->setChecked(false);
                    ui->PtpOcTimePropertiesDsLeap61CheckBox->setChecked(false);
                    ui->PtpOcTimePropertiesDsUtcOffsetValCheckBox->setChecked(false);
                    ui->PtpOcTimePropertiesDsUtcOffsetValue->setText("NA");
                    ui->PtpOcTimePropertiesDsCurrentOffsetValue->setText("NA");
                    ui->PtpOcTimePropertiesDsJumpSecondsValue->setText("NA");
                    ui->PtpOcTimePropertiesDsNextJumpValue->setText("NA");
                    ui->PtpOcTimePropertiesDsDisplayNameValue->setText("NA");
                }

            }
            else
            {
                ui->PtpOcTimePropertiesDsTimeSourceValue->setText("NA");
                ui->PtpOcTimePropertiesDsPtpTimescaleCheckBox->setChecked(false);
                ui->PtpOcTimePropertiesDsFreqTraceableCheckBox->setChecked(false);
                ui->PtpOcTimePropertiesDsTimeTraceableCheckBox->setChecked(false);
                ui->PtpOcTimePropertiesDsLeap59CheckBox->setChecked(false);
                ui->PtpOcTimePropertiesDsLeap61CheckBox->setChecked(false);
                ui->PtpOcTimePropertiesDsUtcOffsetValCheckBox->setChecked(false);
                ui->PtpOcTimePropertiesDsUtcOffsetValue->setText("NA");
                ui->PtpOcTimePropertiesDsCurrentOffsetValue->setText("NA");
                ui->PtpOcTimePropertiesDsJumpSecondsValue->setText("NA");
                ui->PtpOcTimePropertiesDsNextJumpValue->setText("NA");
                ui->PtpOcTimePropertiesDsDisplayNameValue->setText("NA");
            }
        }
    }
    else
    {
        ui->PtpOcTimePropertiesDsTimeSourceValue->setText("NA");
        ui->PtpOcTimePropertiesDsPtpTimescaleCheckBox->setChecked(false);
        ui->PtpOcTimePropertiesDsFreqTraceableCheckBox->setChecked(false);
        ui->PtpOcTimePropertiesDsTimeTraceableCheckBox->setChecked(false);
        ui->PtpOcTimePropertiesDsLeap59CheckBox->setChecked(false);
        ui->PtpOcTimePropertiesDsLeap61CheckBox->setChecked(false);
        ui->PtpOcTimePropertiesDsUtcOffsetValCheckBox->setChecked(false);
        ui->PtpOcTimePropertiesDsUtcOffsetValue->setText("NA");
        ui->PtpOcTimePropertiesDsCurrentOffsetValue->setText("NA");
        ui->PtpOcTimePropertiesDsJumpSecondsValue->setText("NA");
        ui->PtpOcTimePropertiesDsNextJumpValue->setText("NA");
        ui->PtpOcTimePropertiesDsDisplayNameValue->setText("NA");
    }


    // version
    if (0 == ucm->com_lib.read_reg(temp_addr + Ucm_PtpOc_VersionReg, temp_data))
    {
        ui->PtpOcVersionValue->setText(QString("0x%1").arg(temp_data, 8, 16, QLatin1Char('0')));

    }
    else
    {
        ui->PtpOcVersionValue->setText("NA");
    }
}
*/

func showPtpOcStatus() {
	fmt.Println("PTP OC STATUS:                  ", readPtpOcStatus())
}
