package ntscli

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func NtpMac(cmd *cobra.Command) {

	cmd.Flags().Visit(func(f *pflag.Flag) {

		switch f.Name {

		case "addr":
			addr, err := cmd.Flags().GetString(f.Name)
			if err != nil {
				log.Fatal("No such argument for property: ", f.Name, err)
			}
			//fmt.Println("// writeNtpServerAddr(addr)", addr)
			writeNtpServerMac(addr)
		case "list":
			fmt.Println("MAC ADDRESS: ", readNtpServerMac())
		default:
			fmt.Println("That does not appear to be a valid flag. Try: ", cmd.UsageString())
		}
	})

}
