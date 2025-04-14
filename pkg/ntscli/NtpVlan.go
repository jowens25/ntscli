package ntscli

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func NtpVlan(cmd *cobra.Command) {

	cmd.Flags().Visit(func(f *pflag.Flag) {
		switch f.Name {
		case "value":
			value, err := cmd.Flags().GetString(f.Name)
			if err != nil {
				log.Fatal("No such argument for property: ", f.Name, err)
			}
			//fmt.Println("// writeNtpServerAddr(value)", value)
			writeNtpServerVlanValue(value)

		case "enable":
			writeNtpServerVlanEnable("enable")
		case "disable":
			writeNtpServerVlanEnable("disable")
		case "list":
			showNtpServerVLANENABLED()
			showNtpServerVLANVALUE()
		default:
			fmt.Println("That does not appear to be a valid flag. Try: ", cmd.UsageString())
		}
	})
}
