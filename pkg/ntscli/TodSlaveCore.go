package ntscli

var todSlave = RegisterSet{

	"ControlReg":              0x00000000,
	"StatusReg":               0x00000004,
	"PolarityReg":             0x00000008,
	"VersionReg":              0x0000000C,
	"CorrectionReg":           0x00000010,
	"UartBaudRateReg":         0x00000020,
	"UtcStatusReg":            0x00000030,
	"TimeToLeapSecondReg":     0x00000034,
	"GnssStatus_Reg_Con":      0x00000040,
	"SatelliteNumber_Reg_Con": 0x00000044,
}
