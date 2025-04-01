package ntscli

import (
	"fmt"

	"github.com/jowens25/ntscli/pkg/ntscli"
	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:     "set",
	Aliases: []string{"s"},
	Short:   "sets properties",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		res := ntscli.Set(args[0])
		fmt.Println(res)
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
}
