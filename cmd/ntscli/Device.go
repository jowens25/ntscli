package ntscli

import (
	"github.com/jowens25/ntscli/pkg/ntscli"
	"github.com/spf13/cobra"
)

func init() {

	rootCmd.AddCommand(device)
	device.AddCommand(coreConnect)
	device.AddCommand(coreList)

}

var device = &cobra.Command{
	Use:     "device",
	Aliases: []string{"d"},
	Short:   "the fpga device",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		ntscli.Device()

	},
}
var coreConnect = &cobra.Command{
	Use:     "connect",
	Aliases: []string{"connect"},
	Short:   "Use this to connect with the core",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ntscli.DeviceConnect(args[0])

	},
}
var coreList = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "show core config",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		ntscli.DeviceList()

	},
}
