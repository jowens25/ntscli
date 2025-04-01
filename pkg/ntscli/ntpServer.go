package ntscli

import (
	"fmt"
	"log"
)

var IpAddr bool
var MacAddr bool
var All bool

func NtpWrite(input string) {
	if IpAddr {
		log.Println("set ntp server ip to: ", input)
	}

	if MacAddr {
		log.Println("set ntp server mac to: ", input)
	}
}

func NtpRead() {

	if IpAddr {
		log.Println("Read IP")
	}

	if MacAddr {
		log.Println("Read Mac")
	}

	log.Println("ntp read some stuff: ")
}

func Ntp(input string) {
	log.Println("ntp base command: ")
}

func NtpList() {
	ntpCore := CoreConfig{}
	readDeviceConfig(types.NtpServerCoreType, &ntpCore)
	fmt.Println(ntpCore)
}
