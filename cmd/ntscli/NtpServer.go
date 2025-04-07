package ntscli

import (
	"github.com/jowens25/ntscli/pkg/ntscli"
	"github.com/spf13/cobra"
)

func init() {
	ntp.Flags().SortFlags = false
	rootCmd.AddCommand(ntp)
	ntp.AddCommand(ip)
	ntp.AddCommand(mac)

	ntp.AddCommand(vlan)
	//ntp.AddCommand(ntpRead)
	//ntp.AddCommand(ntpWrite)
	//ntp.AddCommand(ntpList)

	//ntp.AddCommand(ntpEnable)
	ntp.Flags().BoolP("core", "c", false, "show core address")
	ntp.Flags().BoolP("status", "s", false, "show ntp server status")
	ntp.Flags().BoolP("enable", "e", false, "set the status of the ntp server (enabled, disabled)")
	ntp.Flags().BoolP("disable", "d", false, "set the status of the ntp server (enabled, disabled)")
	ntp.Flags().BoolP("list", "l", false, "list ntp attributes")
	//ntp.Flags().BoolP("list", "l", false, "list all ntp server attributes")
	//
	//ntp.Flags().Bool("ip", false, "show ip address attributes")
	//ntp.Flags().String("set-ip", "", "set the ip address of the ntp server")
	//ntp.Flags().String("set-ip-mode", "", "set the ip address mode of the ntp server")

	ip.Flags().BoolP("list", "l", false, "list ip info")
	ip.Flags().StringP("mode", "m", "", "set ip mode (IPv4, IPv6)")
	ip.Flags().StringP("addr", "a", "", "set ip addr (0.0.0.0)")

	mac.Flags().BoolP("list", "l", false, "list mac address")
	mac.Flags().StringP("addr", "a", "", "set mac address")

	vlan.Flags().BoolP("list", "l", false, "list vlan info")
	vlan.Flags().BoolP("enable", "e", false, "enable vlan")
	vlan.Flags().BoolP("disable", "d", false, "disablevlan")
	vlan.Flags().StringP("value", "v", "", "set vlan value")
	//ntp.Flags().Bool("list-mac", false, "list the mac address")
	//ntp.Flags().String("set-mac", "", "set the mac address of the ntp server")
	//
	//ntp.Flags().Bool("list-vlan", false, "list the properties of the vlan")
	//ntp.Flags().Bool("enable-vlan", false, "enable vlan of the ntp server")
	//ntp.Flags().Bool("disable-vlan", false, "disable vlan of the ntp server")
	//ntp.Flags().String("set-vlan", "", "set the value of the vlan")

	//ntp.Flags().String("ip", "", "set the ip address of the ntp server")

	//ntp.Flags().Bool("all", false, "show ntp server mac")
	//ntp.Flags().Bool("control", false, "show ntp control reg")
	//ntp.Flags().Bool("enable", false, "show ntp server enable")
	//ntp.Flags().Bool("instance", false, "show ntp server instance")
	//ntp.Flags().Bool("ip", false, "show ntp server ip address")
	//ntp.Flags().Bool("ip-mode", false, "show ntp server ip mode")
	//ntp.Flags().Bool("mac", false, "show ntp server mac address")
	//ntp.Flags().Bool("vlan", false, "show ntp server vlan enabled")
	//ntp.Flags().Bool("vlan ", false, "show ntp server vlan value")
	//ntp.Flags().Bool("unicast", false, "show ntp server unicast")
	//ntp.Flags().Bool("multicast", false, "show ntp server multicast")
	//ntp.Flags().Bool("broadcast", false, "show ntp server broadcast")
	//ntp.Flags().Bool("precision", false, "show ntp server precision")
	//ntp.Flags().Bool("poll interval", false, "show ntp server poll interval")
	//ntp.Flags().Bool("stratum", false, "show ntp server stratum")
	//ntp.Flags().Bool("reference id           ", false, "show ntp server reference id")
	//ntp.Flags().Bool("utc smearing           ", false, "show ntp server utc smearing")
	//ntp.Flags().Bool("utc leap 61 in progress", false, "show ntp server utc leap 61 in progress")
	//ntp.Flags().Bool("utc leap 59 in progress", false, "show ntp server utc leap 59 in progress")
	//ntp.Flags().Bool("utc leap 61            ", false, "show ntp server utc leap 61")
	//ntp.Flags().Bool("utc leap 59            ", false, "show ntp server utc leap 59")
	//ntp.Flags().Bool("utc offset val         ", false, "show ntp server utc offset val")
	//ntp.Flags().Bool("utc offset value       ", false, "show ntp server utc offset value")
	//ntp.Flags().Bool("request count          ", false, "show ntp server request count")
	//ntp.Flags().Bool("response count         ", false, "show ntp server response count ")
	//ntp.Flags().Bool("requests dropped       ", false, "show ntp server requests dropped ")
	//ntp.Flags().Bool("broadcast count        ", false, "show ntp server broadcast count")
	//ntp.Flags().Bool("count control          ", false, "show ntp server count control")
	//ntp.Flags().Bool("version                ", false, "show ntp server version")
	//
	//ntpWrite.Flags().String("ip", "", "set ntp server ip")
	//ntpWrite.Flags().String("mac", "", "set ntp server mac")
	//ntpWrite.Flags().String("vlan", "", "set ntp server vlan address")

}

var ntp = &cobra.Command{
	Use:     "ntp",
	Aliases: []string{"ntp"},
	Short:   "high performance ntp server",
	Args:    cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		ntscli.Ntp(cmd)

	},
}

var ip = &cobra.Command{
	Use:     "ip",
	Aliases: []string{"ip"},
	Short:   "ip address",
	Args:    cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		ntscli.Ntp(cmd)

	},
}

var mac = &cobra.Command{
	Use:     "mac",
	Aliases: []string{"mac"},
	Short:   "mac address",
	Args:    cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		ntscli.Ntp(cmd)

	},
}

var vlan = &cobra.Command{

	Use:     "vlan",
	Aliases: []string{"vlan"},
	Short:   "virtual network",
	Args:    cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		ntscli.Ntp(cmd)

	},
}

//var ntpRead = &cobra.Command{
//	Use:     "read",
//	Aliases: []string{"r"},
//	Short:   "reads properties to stdout",
//	Args:    cobra.ArbitraryArgs,
//	Run: func(cmd *cobra.Command, args []string) {
//
//		ntscli.NtpRead(cmd.Flags())
//	},
//}
//
//var ntpWrite = &cobra.Command{
//	Use:     "write",
//	Aliases: []string{"w"},
//	Short:   "writes the properties of the ntp server",
//	Args:    cobra.ArbitraryArgs,
//	Run: func(cmd *cobra.Command, args []string) {
//		ntscli.NtpWrite(cmd.Flags())
//	},
//}

//var ntpList = &cobra.Command{
//	Use:     "list",
//	Aliases: []string{"ls"},
//	Short:   "show ntp config",
//	Args:    cobra.ExactArgs(0),
//	Run: func(cmd *cobra.Command, args []string) {
//		ntscli.NtpList()
//
//	},
//}
//
