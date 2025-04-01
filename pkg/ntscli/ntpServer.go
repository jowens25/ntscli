package ntscli

import (
	"flag"
	"log"
)

var IpAddr bool
var MacAddr bool
var All bool

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

	Enabled    bool
	MacAddr    string
	VlanEnable bool
	VlanValue  string

	IpMode            string
	UnicastMode       bool
	MulticastMode     bool
	BroadcastMode     bool
	PrecisionValue    rune
	PollIntervalValue string
	StratumValue      string
	ReferenceId       string
	IpAddrMode        string
	IpAddr            string
	Commands          *flag.FlagSet

	UtcSmearing         bool
	UtcLeap61InProgress bool
	UtcLeap59InProgress bool
	UtcLeap61           bool
	UtcLeap59           bool
	UtcOffsetVal        bool
	UtcOffsetValue      int64

	RequestsValue        string
	ResponsesValue       string
	RequestsDroppedValue string
	BroadcastsValue      string

	VersionValue  string
	ClearCounters bool
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

	Enabled:       false,
	MacAddr:       "NA",
	VlanEnable:    false,
	VlanValue:     "NA",
	IpMode:        "NA",
	UnicastMode:   false,
	MulticastMode: false,
	BroadcastMode: false,

	UtcSmearing:         false,
	UtcLeap61InProgress: false,
	UtcLeap59InProgress: false,
	UtcLeap61:           false,
	UtcLeap59:           false,
	UtcOffsetVal:        false,
	UtcOffsetValue:      0,

	RequestsValue:        "NA",
	ResponsesValue:       "NA",
	RequestsDroppedValue: "NA",
	BroadcastsValue:      "NA",

	ClearCounters: false,

	PrecisionValue:    'N',
	PollIntervalValue: "NA",
	StratumValue:      "NA",

	ReferenceId: "NA",
	IpAddrMode:  "NA",
	IpAddr:      "NA",

	Commands: flag.NewFlagSet("ntp", flag.ExitOnError),
}

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

}

func Ntp(input string) {
	log.Println("ntp base command: ")
}

func NtpList() {
	deviceConfig := DeviceConfig{}
	coreConfig := CoreConfig{}
	readDeviceConfig(&deviceConfig)
	if deviceHasNtpServer(&deviceConfig, &coreConfig) == 0 {
		listNtpServerConfig(&coreConfig)
	} else {
		log.Println("whats going on?")
	}

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
