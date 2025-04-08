package ntscli

import (
	"fmt"
	"log"
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

	if writeRegister(NtpCore.BaseAddrLReg+ntpServer.ControlReg, &tempData) == 0 {
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
		tempData = ip_addr_to_int(addr)
		if writeRegister(NtpCore.BaseAddrLReg+ntpServer.ConfigIpReg, &tempData) == 0 {
			tempData = 0x00000008 // write
			if writeRegister(NtpCore.BaseAddrLReg+ntpServer.ConfigControlReg, &tempData) == 0 {
				showNtpServerIPADDRESS()
			} else {
				fmt.Println("NTP SERVER IP ADDR FAILED")
			}
		} else {
			fmt.Println("NTP SERVER IP ADDR FAILED")

		}
	} else if ipMode == "IPv6" {

		//{
		//	temp_string = ui->NtpServerIpValue->text();
		//	temp_ip6 = QHostAddress(temp_string).toIPv6Address();
		//
		//	temp_data = 0x00000000;
		//	temp_data |= temp_ip6[3] & 0x000000FF;
		//	temp_data = temp_data << 8;
		//	temp_data |= temp_ip6[2] & 0x000000FF;
		//	temp_data = temp_data << 8;
		//	temp_data |= temp_ip6[1] & 0x000000FF;
		//	temp_data = temp_data << 8;
		//	temp_data |= temp_ip6[0] & 0x000000FF;
		//
		//	if (temp_string == "NA")
		//	{
		//		// nothing
		//	}
		//	else if (0 == ucm->com_lib.write_reg(temp_addr + Ucm_NtpServer_ConfigIpReg, temp_data))
		//	{
		//		temp_data = 0x00000000;
		//		temp_data |= temp_ip6[7] & 0x000000FF;
		//		temp_data = temp_data << 8;
		//		temp_data |= temp_ip6[6] & 0x000000FF;
		//		temp_data = temp_data << 8;
		//		temp_data |= temp_ip6[5] & 0x000000FF;
		//		temp_data = temp_data << 8;
		//		temp_data |= temp_ip6[4] & 0x000000FF;
		//
		//		if (temp_string == "NA")
		//		{
		//			// nothing
		//		}
		//		else if (0 == ucm->com_lib.write_reg(temp_addr + Ucm_NtpServer_ConfigIpv61Reg, temp_data))
		//		{
		//			temp_data = 0x00000000;
		//			temp_data |= temp_ip6[11] & 0x000000FF;
		//			temp_data = temp_data << 8;
		//			temp_data |= temp_ip6[10] & 0x000000FF;
		//			temp_data = temp_data << 8;
		//			temp_data |= temp_ip6[9] & 0x000000FF;
		//			temp_data = temp_data << 8;
		//			temp_data |= temp_ip6[8] & 0x000000FF;
		//
		//			if (temp_string == "NA")
		//			{
		//				// nothing
		//			}
		//			else if (0 == ucm->com_lib.write_reg(temp_addr + Ucm_NtpServer_ConfigIpv62Reg, temp_data))
		//			{
		//				temp_data = 0x00000000;
		//				temp_data |= temp_ip6[15] & 0x000000FF;
		//				temp_data = temp_data << 8;
		//				temp_data |= temp_ip6[14] & 0x000000FF;
		//				temp_data = temp_data << 8;
		//				temp_data |= temp_ip6[13] & 0x000000FF;
		//				temp_data = temp_data << 8;
		//				temp_data |= temp_ip6[12] & 0x000000FF;
		//
		//				if (temp_string == "NA")
		//				{
		//					// nothing
		//				}
		//				else if (0 == ucm->com_lib.write_reg(temp_addr + Ucm_NtpServer_ConfigIpv63Reg, temp_data))
		//				{
		//					temp_data = 0x00000008; // write
		//					if (0 == ucm->com_lib.write_reg(temp_addr + Ucm_NtpServer_ConfigControlReg, temp_data))
		//					{
		//						ui->NtpServerIpValue->setText(temp_string);
		//					}
		//					else
		//					{
		//						ui->NtpServerIpValue->setText("NA");
		//					}
		//				}
		//				else
		//				{
		//					ui->NtpServerIpValue->setText("NA");
		//				}
		//			}
		//			else
		//			{
		//				ui->NtpServerIpValue->setText("NA");
		//			}
		//		}
		//		else
		//		{
		//			ui->NtpServerIpValue->setText("NA");
		//		}
		//	}
		//	else
		//	{
		//		ui->NtpServerIpValue->setText("NA");
		//	}
		//}
		//else
		//{
		//	ui->NtpServerIpValue->setText("NA");
		//}

	}
	//	           {
	//	               ui->NtpServerIpValue->setText("NA");
	//	           }
	//	       }
	//	       else
	//	       {
	//	           ui->NtpServerIpValue->setText("NA");
	//	       }
	//	   }
	//	   else if (temp_string == "IPv6")
	//	   {
	//	       temp_string = ui->NtpServerIpValue->text();
	//	       temp_ip6 = QHostAddress(temp_string).toIPv6Address();
	//
	//	       temp_data = 0x00000000;
	//	       temp_data |= temp_ip6[3] & 0x000000FF;
	//	       temp_data = temp_data << 8;
	//	       temp_data |= temp_ip6[2] & 0x000000FF;
	//	       temp_data = temp_data << 8;
	//	       temp_data |= temp_ip6[1] & 0x000000FF;
	//	       temp_data = temp_data << 8;
	//	       temp_data |= temp_ip6[0] & 0x000000FF;
	//
	//	       if (temp_string == "NA")
	//	       {
	//	           // nothing
	//	       }
	//	       else if (0 == ucm->com_lib.write_reg(temp_addr + Ucm_NtpServer_ConfigIpReg, temp_data))
	//	       {
	//	           temp_data = 0x00000000;
	//	           temp_data |= temp_ip6[7] & 0x000000FF;
	//	           temp_data = temp_data << 8;
	//	           temp_data |= temp_ip6[6] & 0x000000FF;
	//	           temp_data = temp_data << 8;
	//	           temp_data |= temp_ip6[5] & 0x000000FF;
	//	           temp_data = temp_data << 8;
	//	           temp_data |= temp_ip6[4] & 0x000000FF;
	//
	//	           if (temp_string == "NA")
	//	           {
	//	               // nothing
	//	           }
	//	           else if (0 == ucm->com_lib.write_reg(temp_addr + Ucm_NtpServer_ConfigIpv61Reg, temp_data))
	//	           {
	//	               temp_data = 0x00000000;
	//	               temp_data |= temp_ip6[11] & 0x000000FF;
	//	               temp_data = temp_data << 8;
	//	               temp_data |= temp_ip6[10] & 0x000000FF;
	//	               temp_data = temp_data << 8;
	//	               temp_data |= temp_ip6[9] & 0x000000FF;
	//	               temp_data = temp_data << 8;
	//	               temp_data |= temp_ip6[8] & 0x000000FF;
	//
	//	               if (temp_string == "NA")
	//	               {
	//	                   // nothing
	//	               }
	//	               else if (0 == ucm->com_lib.write_reg(temp_addr + Ucm_NtpServer_ConfigIpv62Reg, temp_data))
	//	               {
	//	                   temp_data = 0x00000000;
	//	                   temp_data |= temp_ip6[15] & 0x000000FF;
	//	                   temp_data = temp_data << 8;
	//	                   temp_data |= temp_ip6[14] & 0x000000FF;
	//	                   temp_data = temp_data << 8;
	//	                   temp_data |= temp_ip6[13] & 0x000000FF;
	//	                   temp_data = temp_data << 8;
	//	                   temp_data |= temp_ip6[12] & 0x000000FF;
	//
	//	                   if (temp_string == "NA")
	//	                   {
	//	                       // nothing
	//	                   }
	//	                   else if (0 == ucm->com_lib.write_reg(temp_addr + Ucm_NtpServer_ConfigIpv63Reg, temp_data))
	//	                   {
	//	                       temp_data = 0x00000008; // write
	//	                       if (0 == ucm->com_lib.write_reg(temp_addr + Ucm_NtpServer_ConfigControlReg, temp_data))
	//	                       {
	//	                           ui->NtpServerIpValue->setText(temp_string);
	//	                       }
	//	                       else
	//	                       {
	//	                           ui->NtpServerIpValue->setText("NA");
	//	                       }
	//	                   }
	//	                   else
	//	                   {
	//	                       ui->NtpServerIpValue->setText("NA");
	//	                   }
	//	               }
	//	               else
	//	               {
	//	                   ui->NtpServerIpValue->setText("NA");
	//	               }
	//	           }
	//	           else
	//	           {
	//	               ui->NtpServerIpValue->setText("NA");
	//	           }
	//	       }
	//	       else
	//	       {
	//	           ui->NtpServerIpValue->setText("NA");
	//	       }
	//	   }
	//	   else
	//	   {
	//	       ui->NtpServerIpValue->setText("NA");
	//	   }
}

// NtpServer IP MODE
func writeNtpServerIpMode(mode string) {
	tempData = 0x00000000
	readRegister(NtpCore.BaseAddrLReg+ntpServer.ConfigControlReg, &tempData)

	if mode == "ipv4" {
		tempData &= ^0x00000001 // Clear bit 16 (using NOT + AND)
		tempData |= 0x00000001

	} else if mode == "ipv6" {
		tempData &= ^0x01000002 // Clear bit 16 (using NOT + AND)
		tempData |= 0x01000002

	} else {
		tempData |= 0x00000000

	}

	if writeRegister(NtpCore.BaseAddrLReg+ntpServer.ConfigModeReg, &tempData) == 0 {
		tempData = 0x00000001
		if writeRegister(NtpCore.BaseAddrLReg+ntpServer.ConfigControlReg, &tempData) == 0 {
			// nothing
		} else {
			fmt.Println("NTP SERVER IP MODE: UPDATE FAILED")
		}
	} else {
		fmt.Println("NTP SERVER UNICAST: UPDATE FAILED")
	}
	showNtpServerIPMODE()
	showNtpServerIPADDRESS()
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

// NtpServer MULTICAST
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

// NtpServer BROADCAST
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

// NtpServer PRECISION
// NtpServer POLLINTERVAL
// NtpServer STRATUM
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

		if writeRegister(NtpCore.BaseAddrLReg+ntpServer.ConfigReferenceIdReg, &tempData) == 0 {
			tempData = 0x00000010

			if writeRegister(NtpCore.BaseAddrLReg+ntpServer.ConfigControlReg, &tempData) == 0 {
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

//NtpServerUTCSMEARING
//NtpServerUTCLEAP61INPROGRESS
//NtpServerUTCLEAP59INPROGRESS
//NtpServerUTCLEAP61
//NtpServerUTCLEAP59
//NtpServerUTCOFFSETENABLE
//NtpServerUTCOFFSETVALUE
//NtpServerREQUESTCOUNT
//NtpServerRESPONSECOUNT
//NtpServerREQUESTSDROPPED
//NtpServerBROADCASTCOUNT
//NtpServerCOUNTCONTROL
//NtpServerVERSION
