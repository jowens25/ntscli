package ntscli

import (
	"github.com/jowens25/ntscli/pkg/ntscli"
	"github.com/spf13/cobra"
)

func init() {

	rootCmd.AddCommand(core)
	core.AddCommand(coreConnect)
	core.AddCommand(coreList)

}

var core = &cobra.Command{
	Use:     "core",
	Aliases: []string{"core"},
	Short:   "the core of the fpga",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		ntscli.Core()

	},
}
var coreConnect = &cobra.Command{
	Use:     "connect",
	Aliases: []string{"connect"},
	Short:   "Use this to connect with the core",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ntscli.CoreConnect(args[0])

	},
}
var coreList = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "show core config",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		ntscli.CoreList()

	},
}
