package ntscli

import (
	"github.com/jowens25/ntscli/pkg/ntscli"
	"github.com/spf13/cobra"
)

func init() {

	rootCmd.AddCommand(ntp)
	ntp.AddCommand(ntpRead)
	ntp.AddCommand(ntpWrite)
	ntp.AddCommand(ntpList)

	//ntp.AddCommand(ntpEnable)

	ntp.Flags().BoolP("enable", "e", false, "enable ntp server")
	ntp.Flags().BoolP("disable", "d", false, "disable ntp server")
	ntp.Flags().BoolP("show", "s", false, "show core address")

	ntpRead.Flags().Bool("all", false, "show ntp server mac")
	ntpRead.Flags().Bool("control", false, "show ntp control reg")
	ntpRead.Flags().Bool("enable", false, "show ntp server enable")
	ntpRead.Flags().Bool("instance", false, "show ntp server instance")
	ntpRead.Flags().Bool("ip", false, "show ntp server ip address")
	ntpRead.Flags().Bool("ip-mode", false, "show ntp server ip mode")
	ntpRead.Flags().Bool("mac", false, "show ntp server mac address")
	ntpRead.Flags().Bool("vlan", false, "show ntp server vlan enabled")
	ntpRead.Flags().Bool("vlan ", false, "show ntp server vlan value")
	ntpRead.Flags().Bool("unicast", false, "show ntp server unicast")
	ntpRead.Flags().Bool("multicast", false, "show ntp server multicast")
	ntpRead.Flags().Bool("broadcast", false, "show ntp server broadcast")
	ntpRead.Flags().Bool("precision", false, "show ntp server precision")
	ntpRead.Flags().Bool("poll interval", false, "show ntp server poll interval")
	ntpRead.Flags().Bool("stratum", false, "show ntp server stratum")
	ntpRead.Flags().Bool("reference id           ", false, "show ntp server reference id")
	ntpRead.Flags().Bool("utc smearing           ", false, "show ntp server utc smearing")
	ntpRead.Flags().Bool("utc leap 61 in progress", false, "show ntp server utc leap 61 in progress")
	ntpRead.Flags().Bool("utc leap 59 in progress", false, "show ntp server utc leap 59 in progress")
	ntpRead.Flags().Bool("utc leap 61            ", false, "show ntp server utc leap 61")
	ntpRead.Flags().Bool("utc leap 59            ", false, "show ntp server utc leap 59")
	ntpRead.Flags().Bool("utc offset val         ", false, "show ntp server utc offset val")
	ntpRead.Flags().Bool("utc offset value       ", false, "show ntp server utc offset value")
	ntpRead.Flags().Bool("request count          ", false, "show ntp server request count")
	ntpRead.Flags().Bool("response count         ", false, "show ntp server response count ")
	ntpRead.Flags().Bool("requests dropped       ", false, "show ntp server requests dropped ")
	ntpRead.Flags().Bool("broadcast count        ", false, "show ntp server broadcast count")
	ntpRead.Flags().Bool("count control          ", false, "show ntp server count control")
	ntpRead.Flags().Bool("version                ", false, "show ntp server version")

	ntpWrite.Flags().String("ip", "", "set ntp server ip")
	ntpWrite.Flags().String("mac", "", "set ntp server mac")
	ntpWrite.Flags().String("vlan", "", "set ntp server vlan address")

}

var ntp = &cobra.Command{
	Use:     "ntp",
	Aliases: []string{"ntp"},
	Short:   "high performance ntp server",
	Args:    cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		ntscli.Ntp(cmd.Flags())

	},
}

var ntpRead = &cobra.Command{
	Use:     "read",
	Aliases: []string{"r"},
	Short:   "reads properties to stdout",
	Args:    cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {

		ntscli.NtpRead(cmd.Flags())
	},
}

var ntpWrite = &cobra.Command{
	Use:     "write",
	Aliases: []string{"w"},
	Short:   "writes the properties of the ntp server",
	Args:    cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		ntscli.NtpWrite(cmd.Flags())
	},
}

var ntpList = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "show ntp config",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		ntscli.NtpList()

	},
}
