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
	ntpRead.Flags().Bool("all", false, "show ntp server mac")
	ntpRead.Flags().Bool("ip", false, "show ntp server ip")
	ntpRead.Flags().Bool("mac", false, "show ntp server mac")

	ntpWrite.Flags().String("ip", "", "set ntp server ip")
	ntpWrite.Flags().String("mac", "", "set ntp server mac")

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
