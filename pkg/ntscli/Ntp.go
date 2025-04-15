package ntscli

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func Ntp(cmd *cobra.Command) {

	cmd.Flags().Visit(func(f *pflag.Flag) {

		switch f.Name {
		case "enable":
			writeNtpServerStatus("enable")
		case "disable":
			writeNtpServerStatus("disable")
		case "core":
			jsonData, err := json.MarshalIndent(NtpServerCore, "", " ")
			if err != nil {
				fmt.Println("some json error")
			}
			fmt.Println("NTP SERVER CORE: ", string(jsonData))
		case "status":
			fmt.Println("NTP SERVER STATUS: ", readNtpServerStatus())

		case "reference":
			refId, err := cmd.Flags().GetString(f.Name)
			if err != nil {
				log.Fatal("No such argument for property: ", f.Name, err)
			}
			writeNtpServerReferenceId(refId)
		case "list":
			NtpPrintAll()

		case "instance":
			showNtpServerINSTANCE()

		case "version":
			showNtpServerVERSION()
		default:
			fmt.Println("That does not appear to be a valid flag. Try: ", cmd.UsageString())
		}
	})

}
