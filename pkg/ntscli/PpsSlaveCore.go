package ntscli

var ppsSlave = RegisterSet{

	"ControlReg":    0x00000000,
	"StatusReg":     0x00000004,
	"PolarityReg":   0x00000008,
	"VersionReg":    0x0000000C,
	"PulseWidthReg": 0x00000010,
	"CableDelayReg": 0x00000020,
}
