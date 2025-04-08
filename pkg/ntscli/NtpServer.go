package ntscli

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type NtpServerType struct {
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
}

var ntpServer = NtpServerType{

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
}
var NtpCore Core

var tempData int64

func DeviceHasNtpServer(dev *Device) int64 {
	for _, core := range dev.Cores {
		if core.CoreType == types.NtpServerCoreType {
			return 0
		}
	}

	return -1
}

func Ntp(cmd *cobra.Command) {
	fmt.Println("MAIN NTP FUNCTION CALL")
	if DeviceHasNtpServer(&device) != 0 {
		log.Fatal("device does not have an ntp server")
	}

	NtpCore = device.Cores["NtpServerCoreType"]

	switch cmd.Name() {

	case "ntp":
		cmd.Flags().Visit(func(f *pflag.Flag) {

			switch f.Name {
			case "enable":
				writeNtpServerStatus("enable")
			case "disable":
				writeNtpServerStatus("disable")
			case "core":
				jsonData, err := json.MarshalIndent(NtpCore, "", " ")
				if err != nil {
					fmt.Println("some json error")
				}
				fmt.Println("NTP SERVER CORE: ", string(jsonData))
			case "status":
				fmt.Println("NTP SERVER STATUS: ", readNtpServerStatus())

			case "reference":
				refId, err := cmd.Flags().GetString(f.Name)
				if err != nil {
					log.Fatal("No such argument for property: ", f.Name, err)
				}
				writeNtpServerReferenceId(refId)
			case "list":
				NtpPrintAll()
			default:
				fmt.Println("That does not appear to be a valid flag. Try: ", cmd.UsageString())
			}
		})

	case "ip":
		cmd.Flags().Visit(func(f *pflag.Flag) {

			switch f.Name {
			case "mode":
				mode, err := cmd.Flags().GetString(f.Name)
				if err != nil {
					log.Fatal("No such argument for property: ", f.Name, err)
				}
				writeNtpServerIpMode(mode)

			case "addr":
				addr, err := cmd.Flags().GetString(f.Name)
				if err != nil {
					log.Fatal("No such argument for property: ", f.Name, err)
				}
				writeNtpServerIpAddr(addr)

			case "list":
				fmt.Println("IP MODE: ", readNtpServerIpMode())
				fmt.Println("IP ADDR: ", readNtpServerIpAddress())

			case "unicast":
				en, err := cmd.Flags().GetString(f.Name)
				if err != nil {
					log.Fatal("No such argument for property: ", f.Name, err)
				}
				writeNtpServerUnicastMode(en)

			default:
				fmt.Println("That does not appear to be a valid flag. Try: ", cmd.UsageString())
			}
		})

	case "mode":
		cmd.Flags().Visit(func(f *pflag.Flag) {

			switch f.Name {

			case "unicast":
				en, err := cmd.Flags().GetString(f.Name)
				if err != nil {
					log.Fatal("No such argument for property: ", f.Name, err)
				}
				writeNtpServerUnicastMode(en)

			case "multicast":
				en, err := cmd.Flags().GetString(f.Name)
				if err != nil {
					log.Fatal("No such argument for property: ", f.Name, err)
				}
				writeNtpServerMulticastMode(en)

			case "broadcast":
				en, err := cmd.Flags().GetString(f.Name)
				if err != nil {
					log.Fatal("No such argument for property: ", f.Name, err)
				}
				writeNtpServerBroadcastMode(en)

			case "disable-all":
				en := "disabled"
				writeNtpServerBroadcastMode(en)
				writeNtpServerMulticastMode(en)
				writeNtpServerUnicastMode(en)
			case "enable-all":
				en := "enabled"
				writeNtpServerBroadcastMode(en)
				writeNtpServerMulticastMode(en)
				writeNtpServerUnicastMode(en)

			case "list":
				fmt.Println("NTP SERVER UNICAST:                   ", readNtpServerUnicastMode())
				fmt.Println("NTP SERVER MULTICAST:                 ", readNtpServerMulticastMode())
				fmt.Println("NTP SERVER BROADCAST:                 ", readNtpServerBroadcastMode())

			default:
				fmt.Println("That does not appear to be a valid flag. Try: ", cmd.UsageString())
			}
		})

	case "mac":
		cmd.Flags().Visit(func(f *pflag.Flag) {

			switch f.Name {

			case "addr":
				addr, err := cmd.Flags().GetString(f.Name)
				if err != nil {
					log.Fatal("No such argument for property: ", f.Name, err)
				}
				//fmt.Println("// writeNtpServerAddr(addr)", addr)
				writeNtpServerMac(addr)
			case "list":
				fmt.Println("MAC ADDRESS: ", readNtpServerMac())
			default:
				fmt.Println("That does not appear to be a valid flag. Try: ", cmd.UsageString())
			}
		})

	case "vlan":
		cmd.Flags().Visit(func(f *pflag.Flag) {

			switch f.Name {

			case "value":
				value, err := cmd.Flags().GetString(f.Name)
				if err != nil {
					log.Fatal("No such argument for property: ", f.Name, err)
				}
				//fmt.Println("// writeNtpServerAddr(value)", value)
				writeNtpServerVlanValue(value)

			case "enable":
				writeNtpServerVlanEnable("enable")
			case "disable":
				writeNtpServerVlanEnable("disable")
			case "list":
				showNtpServerVLANENABLED()
				showNtpServerVLANVALUE()
				//fmt.Println("NTP SERVER VLAN ENABLE: ", readNtpServerVlanEnable())
				//fmt.Println("NTP SERVER VLAN VALUE: ", readNtpServerVlanValue())
			default:
				fmt.Println("That does not appear to be a valid flag. Try: ", cmd.UsageString())
			}
		})

		//fmt.Println("NTP SERVER VLAN ENABLE: ", readNtpServerVlanEnable())
		//fmt.Println("NTP SERVER VLAN VALUE: ", readNtpServerVlanValue())

	default:
		fmt.Println("default case")
		cmd.Flags().Visit(func(f *pflag.Flag) { fmt.Println(f.Name) })
	}

}

//func updateNtpServerMac(macAddr string) {
//	writeNtpServerMac(macAddr)
//	fmt.Println("VERBOSE NTP SERVER MAC ADDRESS: ", readNtpServerMac())
//}
