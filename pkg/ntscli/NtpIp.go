package ntscli

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func NtpIp(cmd *cobra.Command) {

	cmd.Flags().Visit(func(f *pflag.Flag) {

		switch f.Name {
		case "mode":
			mode, err := cmd.Flags().GetString(f.Name)
			if err != nil {
				log.Fatal("No such argument for property: ", f.Name, err)
			}
			writeNtpServerIpMode(mode)

		case "addr":
			addr, err := cmd.Flags().GetString(f.Name)
			if err != nil {
				log.Fatal("No such argument for property: ", f.Name, err)
			}
			writeNtpServerIpAddr(addr)

		case "list":
			fmt.Println("IP MODE: ", readNtpServerIpMode())
			fmt.Println("IP ADDR: ", readNtpServerIpAddress())

		case "unicast":
			en, err := cmd.Flags().GetString(f.Name)
			if err != nil {
				log.Fatal("No such argument for property: ", f.Name, err)
			}
			writeNtpServerUnicastMode(en)

		default:
			fmt.Println("That does not appear to be a valid flag. Try: ", cmd.UsageString())
		}
	})

}
