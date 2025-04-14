package ntscli

import "log"

func writePtpOcStatus(status string) {

	if status == "enable" {
		tempData = 0x00000001
	} else if status == "disable" {
		tempData = 0x00000000
	} else {
		log.Fatal("Please enter a valid status (enabled or disabled)")
	}

	if writeRegister(PtpOcCore.BaseAddrLReg+ptpOc["ControlReg"], &tempData) == 0 {
		showPtpOcStatus()
	} else {
		log.Fatal(" VERBOSE ERROR WRITING PTP OC")
	}
}
