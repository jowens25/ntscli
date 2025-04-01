package ntscli

import (
	"log"
)

func Core() {
	log.Println("Core base command")
}

func CoreConnect(input string) {

	log.Println("connect to device at: ", input)
}

func CoreList() {
	log.Println("List out the core config")
	coreConfig := CoreConfig{}
	readDeviceConfig(0, &coreConfig)
}
