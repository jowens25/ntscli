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
	//device.AddCommand(coreConnect)
	//device.AddCommand(coreList)
	device.AddCommand(pullConfig)
	// ntp server section
	ntp.Flags().SortFlags = false
	rootCmd.AddCommand(ntp)

	ntp.Flags().BoolP("version", "v", false, "show core version")
	ntp.Flags().BoolP("instance", "i", false, "show core instance number")
	ntp.Flags().BoolP("core", "c", false, "show core address")
	ntp.Flags().BoolP("status", "s", false, "show ntp server status")
	ntp.Flags().BoolP("list", "l", false, "list ntp attributes")
	ntp.Flags().BoolP("enable", "e", false, "set the status of the ntp server (enabled, disabled)")
	ntp.Flags().BoolP("disable", "d", false, "set the status of the ntp server (enabled, disabled)")

	ntp.AddCommand(ntpMac)
	ntpMac.Flags().BoolP("list", "l", false, "list mac address")
	ntpMac.Flags().StringP("addr", "a", "", "set mac address")

	ntp.AddCommand(ntpVlan)
	ntpVlan.Flags().BoolP("list", "l", false, "show the vlan configuration")
	ntpVlan.Flags().BoolP("address", "a", false, "set the address or value of the vlan")
	ntpVlan.Flags().BoolP("status", "s", false, "set the status of the vlan (enabled or disabled)")

	ntp.AddCommand(ntpIp)
	ntpIp.Flags().BoolP("list", "l", false, "show the ip configuration")
	ntpIp.Flags().BoolP("mode", "m", false, "set the ip mode (IPv4, IPv6)")
	ntpIp.Flags().BoolP("address", "a", false, "set the ip address (0.0.0.0)")

	ntp.AddCommand(ntpMode)
	ntpMode.Flags().BoolP("list", "l", false, "list modes")
	ntpMode.Flags().StringP("unicast", "u", "", "unicast mode")
	ntpMode.Flags().StringP("multicast", "m", "", "multicast mode")
	ntpMode.Flags().StringP("broadcast", "b", "", "broadcast mode")
	ntpMode.Flags().BoolP("disable-all", "d", false, "disable all modes")
	ntpMode.Flags().BoolP("enable-all", "e", false, "enable all modes")

	ntp.AddCommand(ntpServer)
	ntpServer.Flags().StringP("stratum", "s", "", "set the stratum of the ntp server")
	ntpServer.Flags().StringP("poll-interval", "i", "", "set the poll interval of the ntp server")
	ntpServer.Flags().StringP("precision", "p", "", "set the precision of the ntp server")
	ntpServer.Flags().StringP("reference", "r", "", "set the reference of the ntp server")
	ntpServer.Flags().BoolP("list", "l", false, "show ntp server configuration")

	ntp.AddCommand(ntpUtc)
	ntpUtc.Flags().StringP("smearing", "s", "", "enable UTC smearing")
	ntpUtc.Flags().String("leap61", "", "enable UTC leap 61")
	ntpUtc.Flags().String("leap59", "", "enable UTC leap 59")
	ntpUtc.Flags().StringP("enable-offset", "e", "", "enable UTC offset")
	ntpUtc.Flags().StringP("offset", "o", "", "set UTC offset value")

	ntp.AddCommand(ntpStats)
	ntpStats.Flags().BoolP("requests", "q", false, "show ntp request count")
	ntpStats.Flags().BoolP("responses", "r", false, "show ntp responses count")
	ntpStats.Flags().BoolP("dropped", "d", false, "show ntp dropped count")
	ntpStats.Flags().BoolP("broadcasts", "b", false, "show ntp broadcast count")
	ntpStats.Flags().BoolP("all", "a", false, "show all ntp stats")

	ntp.AddCommand(ntpClear)

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

	PersistentPreRun: func(cmd *cobra.Command, args []string) {

		ntscli.ReadDeviceConfig()

	},
}

var device = &cobra.Command{
	Use:     "device",
	Aliases: []string{"d"},
	Short:   "the fpga device",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		//ntscli.Device()

	},
}

var pullConfig = &cobra.Command{
	Use:     "pull-config",
	Aliases: []string{"pull"},
	Short:   "get config from device and store in temp file",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		//ntscli.DevicePullConfig()

	},
}

var ntp = &cobra.Command{
	Use:     "ntp",
	Aliases: []string{"ntp"},
	Short:   "high performance ntp server",
	Args:    cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		ntscli.Ntp(cmd)
	},

	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("ntp pre run")
		if ntscli.DeviceHasNtpServer() != 0 {
			log.Fatal("No NTP Core Found")
		}

	},
}

var ntpMac = &cobra.Command{
	Use:     "mac",
	Aliases: []string{"mac"},
	Short:   "mac address",
	Args:    cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		ntscli.NtpMac(cmd)
	},
}

var ntpVlan = &cobra.Command{
	Use:     "vlan",
	Aliases: []string{"vlan"},
	Short:   "ntp server virtual network",
	Args:    cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		ntscli.NtpVlan(cmd)
	},
}

var ntpIp = &cobra.Command{
	Use:     "ip",
	Aliases: []string{"ip"},
	Short:   "ip address",
	Args:    cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		ntscli.NtpIp(cmd)
	},
}

var ntpMode = &cobra.Command{
	Use:     "mode",
	Aliases: []string{"mode"},
	Short:   "m",
	Args:    cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		ntscli.NtpMode(cmd)

	},
}

var ntpServer = &cobra.Command{

	Use:     "server",
	Aliases: []string{"server"},
	Short:   "ntp server configuration",
	Args:    cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		ntscli.NtpServer(cmd)

	},
}

var ntpUtc = &cobra.Command{

	Use:     "utc",
	Aliases: []string{"utc"},
	Short:   "Coordinated Universal Time",
	Args:    cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		ntscli.NtpUtc(cmd)

	},
}

var ntpStats = &cobra.Command{

	Use:     "stats",
	Aliases: []string{"stats"},
	Short:   "NTP server stats",
	Args:    cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		ntscli.NtpStats(cmd)

	},
}

var ntpClear = &cobra.Command{
	Use:     "clear",
	Aliases: []string{"clear"},
	Short:   "clear counters",
	Args:    cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		ntscli.NtpClear()

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
		fmt.Println("ptpoc pre run")
		if ntscli.DeviceHasPtpOc() != 0 {
			log.Fatal("No PTP OC Core Found")
		}

	},
}

var ptpOcVlan

var ptpOcProfile

var ptpOcLayer 

var ptpOcMode

var ptpOcIp



func Execute() {

	// detect :TODO
	// connect :TODO
	ntscli.ReadDeviceConfig()

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
