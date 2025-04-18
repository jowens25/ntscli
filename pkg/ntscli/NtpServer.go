package ntscli

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func NtpServer(cmd *cobra.Command) {

	cmd.Flags().Visit(func(f *pflag.Flag) {

		switch f.Name {

		case "stratum":
			value, err := cmd.Flags().GetString(f.Name)
			if err != nil {
				log.Fatal("No such argument for property: ", f.Name, err)
			}
			writeNtpServerStratumValue(value)

		case "poll-interval":
			value, err := cmd.Flags().GetString(f.Name)
			if err != nil {
				log.Fatal("No such argument for property: ", f.Name, err)
			}
			writeNtpServerPollInternal(value)

		case "precision":
			value, err := cmd.Flags().GetString(f.Name)
			if err != nil {
				log.Fatal("No such argument for property: ", f.Name, err)
			}
			writeNtpServerPrecision(value)

		case "reference":
			value, err := cmd.Flags().GetString(f.Name)
			if err != nil {
				log.Fatal("No such argument for property: ", f.Name, err)
			}
			writeNtpServerReferenceId(value)

		case "list":

			showNtpServerSTRATUM()
			showNtpServerPOLLINTERVAL()
			showNtpServerPRECISION()
			showNtpServerREFERENCEID()

		default:
			fmt.Println("That does not appear to be a valid flag. Try: ", cmd.UsageString())
		}
	})
}
