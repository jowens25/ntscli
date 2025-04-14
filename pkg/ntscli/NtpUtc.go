package ntscli

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func NtpUtc(cmd *cobra.Command) {

	cmd.Flags().Visit(func(f *pflag.Flag) {

		switch f.Name {

		case "smearing":
			value, err := cmd.Flags().GetString(f.Name)
			if err != nil {
				log.Fatal("No such argument for property: ", f.Name, err)
			}
			writeNtpServerUTCSmearing(value)

		case "leap61":
			value, err := cmd.Flags().GetString(f.Name)
			if err != nil {
				log.Fatal("No such argument for property: ", f.Name, err)
			}
			writeNtpServerUTCLeap61(value)

		case "leap59":
			value, err := cmd.Flags().GetString(f.Name)
			if err != nil {
				log.Fatal("No such argument for property: ", f.Name, err)
			}
			writeNtpServerUTCLeap59(value)

		case "enable-offset":
			value, err := cmd.Flags().GetString(f.Name)
			if err != nil {
				log.Fatal("No such argument for property: ", f.Name, err)
			}
			writeNtpServerUTCOffsetEnable(value)

		case "offset":
			fmt.Println("offset flag")
			value, err := cmd.Flags().GetString(f.Name)
			if err != nil {
				log.Fatal("No such argument for property: ", f.Name, err)
			}
			writeNtpServerUTCOffsetValue(value)

		default:
			fmt.Println("That does not appear to be a valid flag. Try: ", cmd.UsageString())
		}
	})
}
