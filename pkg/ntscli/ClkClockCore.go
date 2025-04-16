package ntscli

var clkClock = RegisterSet{

	"ControlReg":               0x00000000,
	"StatusReg":                0x00000004,
	"SelectReg":                0x00000008,
	"VersionReg":               0x0000000C,
	"TimeValueLReg":            0x00000010,
	"TimeValueHReg":            0x00000014,
	"TimeAdjValueLReg":         0x00000020,
	"TimeAdjValueHReg":         0x00000024,
	"OffsetAdjValueReg":        0x00000030,
	"OffsetAdjIntervalReg":     0x00000034,
	"DriftAdjValueReg":         0x00000040,
	"DriftAdjIntervalReg":      0x00000044,
	"InSyncThresholdReg":       0x00000050,
	"ServoOffsetFactorPReg":    0x00000060,
	"ServoOffsetFactorIReg":    0x00000064,
	"ServoDriftFactorPReg":     0x00000068,
	"ServoDriftFactorIReg":     0x0000006C,
	"StatusOffsetReg":          0x00000070,
	"StatusDriftReg":           0x00000074,
	"StatusOffsetFractionsReg": 0x00000078,
	"StatusDriftFractionsReg":  0x0000007C,
}
