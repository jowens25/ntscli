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
	ntp.AddCommand(mode)

	// ntp reads
	ntp.Flags().BoolP("core", "c", false, "show core address")
	ntp.Flags().BoolP("status", "s", false, "show ntp server status")
	ntp.Flags().BoolP("list", "l", false, "list ntp attributes")
	ntp.Flags().BoolP("enable", "e", false, "set the status of the ntp server (enabled, disabled)")
	ntp.Flags().BoolP("disable", "d", false, "set the status of the ntp server (enabled, disabled)")

	ip.Flags().BoolP("list", "l", false, "list ip info")

	ip.Flags().StringP("mode", "m", "", "set ip mode (IPv4, IPv6)")
	ip.Flags().StringP("addr", "a", "", "set ip addr (0.0.0.0)")

	mode.Flags().BoolP("list", "l", false, "list modes")
	mode.Flags().StringP("unicast", "u", "", "unicast mode")
	mode.Flags().StringP("multicast", "m", "", "multicast mode")
	mode.Flags().StringP("broadcast", "b", "", "broadcast mode")
	mode.Flags().BoolP("disable-all", "d", false, "disable all modes")
	mode.Flags().BoolP("enable-all", "e", false, "enable all modes")

	mac.Flags().BoolP("list", "l", false, "list mac address")
	mac.Flags().StringP("addr", "a", "", "set mac address")

	vlan.Flags().BoolP("list", "l", false, "list vlan info")
	vlan.Flags().BoolP("enable", "e", false, "enable vlan")
	vlan.Flags().BoolP("disable", "d", false, "disablevlan")
	vlan.Flags().StringP("value", "v", "", "set vlan value")

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

var mode = &cobra.Command{
	Use:     "mode",
	Aliases: []string{"modes"},
	Short:   "ntp transmission modes",
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
