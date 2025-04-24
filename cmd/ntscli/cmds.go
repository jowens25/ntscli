package ntscli

import (
	"fmt"
	"log"
	"os"

	"github.com/jowens25/ntscli/pkg/ntscli"
	"github.com/spf13/cobra"
)

func init() {
	// device section
	rootCmd.AddCommand(device)

	rootCmd.AddCommand(run)
	rootCmd.AddCommand(stop)

	device.Flags().StringP("load", "l", "", "load a config file")
	device.Flags().StringP("dump", "d", "", "dump a config file")
	device.Flags().BoolP("connect", "c", false, "attempt to connect to FPGA")
	//device.AddCommand(coreConnect)
	//device.AddCommand(coreList)
	//device.AddCommand(pullConfig)
	// ntp server section

	rootCmd.AddCommand(ptpOc)
	ptpOc.Flags().BoolP("version", "v", false, "show core version")
	ptpOc.Flags().BoolP("instance", "i", false, "show core instance number")
	ptpOc.Flags().BoolP("core", "c", false, "show core address")
	ptpOc.Flags().BoolP("status", "s", false, "show ntp server status")
	ptpOc.Flags().BoolP("list", "l", false, "list ntp attributes")
	ptpOc.Flags().BoolP("enable", "e", false, "set the status of the ntp server (enabled, disabled)")
	ptpOc.Flags().BoolP("disable", "d", false, "set the status of the ntp server (enabled, disabled)")

	//
	//
	////ptpoc.Flags().SortFlags = false
	////rootCmd.AddCommand(ptpoc)
	//
	//ptpoc.AddCommand(vlan)
	//
	//ptpoc.AddCommand(test)
	//// ptpoc - version RO
	//ptpoc.Flags().BoolP("version", "v", false, "show ptp oc version")
	//
	//// ptpoc - instance RO
	//ptpoc.Flags().BoolP("instance", "i", false, "show ptp oc instance")
	//
	//// ptpoc - vlan enable RW
	//ptpoc.Flags().String("vlan", "", "")
	////ptpoc --vlan enable
	////ptpoc --vlan disable
	////ptpoc --vlan value
	//// ptpoc - vlan value RW
	//// ptpoc - profile/layer RW
	//// ptpoc - layer RW
	//// ptpoc - p2p RW
	//// ptpoc - ip address RW
	//// ptpoc - status / enabled RW
	//
	//ptpoc.Flags().BoolP("status", "s", false, "show ptp oc status")
	//ptpoc.Flags().BoolP("enable", "e", false, "enable ptp oc")
	//ptpoc.Flags().BoolP("disable", "d", false, "disable ptp oc")

}

var rootCmd = &cobra.Command{
	Use:     "ntscli",
	Short:   "Novus Time Server configuration tool",
	Long:    "Novus Time Server configuration tool is used for updating the parameters in Novus Power Products Time Servers",
	Version: "0.0.1",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var device = &cobra.Command{
	Use:     "device",
	Aliases: []string{"d"},
	Short:   "the fpga device",
	Args:    cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		ntscli.UpdateDevice(cmd)

	},

	PersistentPreRun: func(cmd *cobra.Command, args []string) {

		ntscli.ReadDeviceConfig()

	},
}

// ptp oc command section

var ptpOc = &cobra.Command{
	Use:     "ptpoc",
	Aliases: []string{"ptpoc"},
	Short:   "high performance ptp oc",
	Args:    cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		ntscli.PtpOc(cmd)
	},

	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		ntscli.ReadDeviceConfig()

		fmt.Println("ptpoc pre run")
		if ntscli.DeviceHasPtpOc() != 0 {
			log.Fatal("No PTP OC Core Found")
		}

	},
}

var run = &cobra.Command{
	Use:     "run",
	Aliases: []string{"run"},
	Short:   "run the go web server",
	Args:    cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		//ntscli.Runner()
		ntscli.RunServers()
	},

	//PersistentPreRun: func(cmd *cobra.Command, args []string) {
	//	ntscli.ReadDeviceConfig()
	//
	//	fmt.Println("ptpoc pre run")
	//	if ntscli.DeviceHasPtpOc() != 0 {
	//		log.Fatal("No PTP OC Core Found")
	//	}
	//
	//},
}

var stop = &cobra.Command{
	Use:     "stop",
	Aliases: []string{"stop"},
	Short:   "stop the go web server",
	Args:    cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		//ntscli.Runner()
		ntscli.StopServers()
	},

	//PersistentPreRun: func(cmd *cobra.Command, args []string) {
	//	ntscli.ReadDeviceConfig()
	//
	//	fmt.Println("ptpoc pre run")
	//	if ntscli.DeviceHasPtpOc() != 0 {
	//		log.Fatal("No PTP OC Core Found")
	//	}
	//
	//},
}

//var ptpOcVlan
//
//var ptpOcProfile
//
//var ptpOcLayer
//
//var ptpOcMode
//
//var ptpOcIp

func Execute() {

	// detect :TODO
	// connect :TODO
	// ntscli.ReadDeviceConfig()

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
