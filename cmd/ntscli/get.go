package ntscli

import (
	"fmt"

	"github.com/jowens25/ntscli/cmd/ntscli"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:     "get",
	Aliases: []string{"g"},
	Short:   "gets properties",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		res := ntscli.Get(args[0])
		fmt.Println(res)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
