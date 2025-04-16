package ntscli

type RegisterSet map[string]int64

type Types struct {
	ConfSlaveCoreType             int64
	ClkClockCoreType              int64
	ClkSignalGeneratorCoreType    int64
	ClkSignalTimestamperCoreType  int64
	IrigSlaveCoreType             int64
	IrigMasterCoreType            int64
	PpsSlaveCoreType              int64
	PpsMasterCoreType             int64
	PtpOrdinaryClockCoreType      int64
	PtpTransparentClockCoreType   int64
	PtpHybridClockCoreType        int64
	RedHsrPrpCoreType             int64
	RtcSlaveCoreType              int64
	RtcMasterCoreType             int64
	TodSlaveCoreType              int64
	TodMasterCoreType             int64
	TapSlaveCoreType              int64
	DcfSlaveCoreType              int64
	DcfMasterCoreType             int64
	RedTsnCoreType                int64
	TsnIicCoreType                int64
	NtpServerCoreType             int64
	NtpClientCoreType             int64
	ClkFrequencyGeneratorCoreType int64
	SynceNodeCoreType             int64
	PpsClkToPpsCoreType           int64
	PtpServerCoreType             int64
	PtpClientCoreType             int64
	PhyConfigurationCoreType      int64
	I2cConfigurationCoreType      int64
	IoConfigurationCoreType       int64
	EthernetTestplatformType      int64
	MinSwitchCoreType             int64
	ConfExtCoreType               int64
}

var types = Types{
	ConfSlaveCoreType:             1,
	ClkClockCoreType:              2,
	ClkSignalGeneratorCoreType:    3,
	ClkSignalTimestamperCoreType:  4,
	IrigSlaveCoreType:             5,
	IrigMasterCoreType:            6,
	PpsSlaveCoreType:              7,
	PpsMasterCoreType:             8,
	PtpOrdinaryClockCoreType:      9,
	PtpTransparentClockCoreType:   10,
	PtpHybridClockCoreType:        11,
	RedHsrPrpCoreType:             12,
	RtcSlaveCoreType:              13,
	RtcMasterCoreType:             14,
	TodSlaveCoreType:              15,
	TodMasterCoreType:             16,
	TapSlaveCoreType:              17,
	DcfSlaveCoreType:              18,
	DcfMasterCoreType:             19,
	RedTsnCoreType:                20,
	TsnIicCoreType:                21,
	NtpServerCoreType:             22,
	NtpClientCoreType:             23,
	ClkFrequencyGeneratorCoreType: 25,
	SynceNodeCoreType:             26,
	PpsClkToPpsCoreType:           27,
	PtpServerCoreType:             28,
	PtpClientCoreType:             29,
	PhyConfigurationCoreType:      10000,
	I2cConfigurationCoreType:      10001,
	IoConfigurationCoreType:       10002,
	EthernetTestplatformType:      10003,
	MinSwitchCoreType:             10004,
	ConfExtCoreType:               20000,
}

func getName(core_type int64) string {

	switch core_type {

	case types.ConfSlaveCoreType:
		return "ConfSlaveCoreType"

	case types.ClkClockCoreType:
		return "ClkClockCoreType"

	case types.ClkSignalGeneratorCoreType:
		return "ClkSignalGeneratorCoreType"

	case types.ClkSignalTimestamperCoreType:
		return "ClkSignalTimestamperCoreType"

	case types.IrigSlaveCoreType:
		return "IrigSlaveCoreType"

	case types.IrigMasterCoreType:
		return "IrigMasterCoreType"

	case types.PpsSlaveCoreType:
		return "PpsSlaveCoreType"

	case types.PpsMasterCoreType:
		return "PpsMasterCoreType"

	case types.PtpOrdinaryClockCoreType:
		return "PtpOrdinaryClockCoreType"

	case types.PtpTransparentClockCoreType:
		return "PtpTransparentClockCoreType"

	case types.PtpHybridClockCoreType:
		return "PtpHybridClockCoreType"

	case types.RedHsrPrpCoreType:
		return "RedHsrPrpCoreType"

	case types.RtcSlaveCoreType:
		return "RtcSlaveCoreType"

	case types.RtcMasterCoreType:
		return "RtcMasterCoreType"

	case types.TodSlaveCoreType:
		return "TodSlaveCoreType"

	case types.TodMasterCoreType:
		return "TodMasterCoreType"

	case types.TapSlaveCoreType:
		return "TapSlaveCoreType"

	case types.DcfSlaveCoreType:
		return "DcfSlaveCoreType"

	case types.DcfMasterCoreType:
		return "DcfMasterCoreType"

	case types.RedTsnCoreType:
		return "RedTsnCoreType"

	case types.TsnIicCoreType:
		return "TsnIicCoreType"

	case types.NtpServerCoreType:
		return "NtpServerCoreType"

	case types.NtpClientCoreType:
		return "NtpClientCoreType"

	case types.ClkFrequencyGeneratorCoreType:
		return "ClkFrequencyGeneratorCoreType"

	case types.SynceNodeCoreType:
		return "SynceNodeCoreType"

	case types.PpsClkToPpsCoreType:
		return "PpsClkToPpsCoreType"

	case types.PtpServerCoreType:
		return "PtpServerCoreType"

	case types.PtpClientCoreType:
		return "PtpClientCoreType"

	case types.PhyConfigurationCoreType:
		return "PhyConfigurationCoreType"

	case types.I2cConfigurationCoreType:
		return "I2cConfigurationCoreType"

	case types.IoConfigurationCoreType:
		return "IoConfigurationCoreType"

	case types.EthernetTestplatformType:
		return "EthernetTestplatformType"

	case types.MinSwitchCoreType:
		return "MinSwitchCoreType"

	case types.ConfExtCoreType:
		return "ConfExtCoreType"

	default:
		return "Core not found error type: " + string(core_type)

	}

}

func getRegistersByType(core_type int64) map[string]int64 {

	switch core_type {

	case types.ConfSlaveCoreType:
		return confSlave
	//
	case types.ClkClockCoreType:
		return clkClock
	//
	//case types.ClkSignalGeneratorCoreType:
	//	return "ClkSignalGeneratorCoreType"
	//
	//case types.ClkSignalTimestamperCoreType:
	//	return "ClkSignalTimestamperCoreType"
	//
	//case types.IrigSlaveCoreType:
	//	return "IrigSlaveCoreType"
	//
	//case types.IrigMasterCoreType:
	//	return "IrigMasterCoreType"
	//
	case types.PpsSlaveCoreType:
		return ppsSlave
	//
	//case types.PpsMasterCoreType:
	//	return "PpsMasterCoreType"
	//
	case types.PtpOrdinaryClockCoreType:
		return ptpOc

		//	case types.PtpTransparentClockCoreType:
		//		return "PtpTransparentClockCoreType"
		//
		//	case types.PtpHybridClockCoreType:
		//		return "PtpHybridClockCoreType"
		//
		//	case types.RedHsrPrpCoreType:
		//		return "RedHsrPrpCoreType"
		//
		//	case types.RtcSlaveCoreType:
		//		return "RtcSlaveCoreType"
		//
		//	case types.RtcMasterCoreType:
		//		return "RtcMasterCoreType"
		//
	case types.TodSlaveCoreType:
		return todSlave
		//
		//	case types.TodMasterCoreType:
		//		return "TodMasterCoreType"
		//
		//	case types.TapSlaveCoreType:
		//		return "TapSlaveCoreType"
		//
		//	case types.DcfSlaveCoreType:
		//		return "DcfSlaveCoreType"
		//
		//	case types.DcfMasterCoreType:
		//		return "DcfMasterCoreType"
		//
		//	case types.RedTsnCoreType:
		//		return "RedTsnCoreType"
		//
		//	case types.TsnIicCoreType:
		//		return "TsnIicCoreType"
		//
	case types.NtpServerCoreType:
		return ntpServer

		//case types.NtpClientCoreType:
		//	return "NtpClientCoreType"
		//
		//case types.ClkFrequencyGeneratorCoreType:
		//	return "ClkFrequencyGeneratorCoreType"
		//
		//case types.SynceNodeCoreType:
		//	return "SynceNodeCoreType"
		//
		//case types.PpsClkToPpsCoreType:
		//	return "PpsClkToPpsCoreType"
		//
		//case types.PtpServerCoreType:
		//	return "PtpServerCoreType"
		//
		//case types.PtpClientCoreType:
		//	return "PtpClientCoreType"
		//
		//case types.PhyConfigurationCoreType:
		//	return "PhyConfigurationCoreType"
		//
		//case types.I2cConfigurationCoreType:
		//	return "I2cConfigurationCoreType"
		//
		//case types.IoConfigurationCoreType:
		//	return "IoConfigurationCoreType"
		//
		//case types.EthernetTestplatformType:
		//	return "EthernetTestplatformType"
		//
		//case types.MinSwitchCoreType:
		//	return "MinSwitchCoreType"
		//
		//case types.ConfExtCoreType:
		//return "ConfExtCoreType"

	default:
		return map[string]int64{
			"failed": 0x00000000}

	}

}
