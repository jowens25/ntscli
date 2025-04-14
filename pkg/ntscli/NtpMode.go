package ntscli

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func NtpMode(cmd *cobra.Command) {

	cmd.Flags().Visit(func(f *pflag.Flag) {

		switch f.Name {

		case "unicast":
			en, err := cmd.Flags().GetString(f.Name)
			if err != nil {
				log.Fatal("No such argument for property: ", f.Name, err)
			}
			writeNtpServerUnicastMode(en)

		case "multicast":
			en, err := cmd.Flags().GetString(f.Name)
			if err != nil {
				log.Fatal("No such argument for property: ", f.Name, err)
			}
			writeNtpServerMulticastMode(en)

		case "broadcast":
			en, err := cmd.Flags().GetString(f.Name)
			if err != nil {
				log.Fatal("No such argument for property: ", f.Name, err)
			}
			writeNtpServerBroadcastMode(en)

		case "disable-all":
			en := "disabled"
			writeNtpServerBroadcastMode(en)
			writeNtpServerMulticastMode(en)
			writeNtpServerUnicastMode(en)
		case "enable-all":
			en := "enabled"
			writeNtpServerBroadcastMode(en)
			writeNtpServerMulticastMode(en)
			writeNtpServerUnicastMode(en)

		case "list":
			fmt.Println("NTP SERVER UNICAST:                   ", readNtpServerUnicastMode())
			fmt.Println("NTP SERVER MULTICAST:                 ", readNtpServerMulticastMode())
			fmt.Println("NTP SERVER BROADCAST:                 ", readNtpServerBroadcastMode())

		default:
			fmt.Println("That does not appear to be a valid flag. Try: ", cmd.UsageString())
		}
	})

}
