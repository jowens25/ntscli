package ntscli

var ntpServer = RegisterSet{
	"ControlReg":           0x00000000,
	"StatusReg":            0x00000004,
	"VersionReg":           0x0000000C,
	"CountControlReg":      0x00000010,
	"CountReqReg":          0x00000014,
	"CountRespReg":         0x00000018,
	"CountReqDroppedReg":   0x0000001C,
	"CountBroadcastReg":    0x00000020,
	"ConfigControlReg":     0x00000080,
	"ConfigModeReg":        0x00000084,
	"ConfigVlanReg":        0x00000088,
	"ConfigMac1Reg":        0x0000008C,
	"ConfigMac2Reg":        0x00000090,
	"ConfigIpReg":          0x00000094,
	"ConfigIpv61Reg":       0x00000098,
	"ConfigIpv62Reg":       0x0000009C,
	"ConfigIpv63Reg":       0x000000A0,
	"ConfigReferenceIdReg": 0x000000A4,
	"UtcInfoControlReg":    0x00000100,
	"UtcInfoReg":           0x00000104,
}
