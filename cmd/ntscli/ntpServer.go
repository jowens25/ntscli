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

	ntpRead.Flags().BoolVar(&ntscli.All, "all", false, "show ntp server mac")
	ntpRead.Flags().BoolVar(&ntscli.IpAddr, "ip", false, "show ntp server ip")
	ntpRead.Flags().BoolVar(&ntscli.MacAddr, "mac", false, "show ntp server mac")

	ntpWrite.Flags().BoolVar(&ntscli.IpAddr, "ip", false, "set ntp server ip")
	ntpWrite.Flags().BoolVar(&ntscli.MacAddr, "mac", false, "set ntp server mac")

}

var ntp = &cobra.Command{
	Use:     "ntp",
	Aliases: []string{"ntp"},
	Short:   "high performance ntp server",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ntscli.Ntp(args[0])

	},
}

var ntpRead = &cobra.Command{
	Use:     "read",
	Aliases: []string{"r"},
	Short:   "reads properties to stdout",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		ntscli.NtpRead()
	},
}

var ntpWrite = &cobra.Command{
	Use:     "write",
	Aliases: []string{"w"},
	Short:   "writes properties to the time",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ntscli.NtpWrite(args[0])
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
