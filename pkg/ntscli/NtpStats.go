package ntscli

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func NtpStats(cmd *cobra.Command) {

	cmd.Flags().Visit(func(f *pflag.Flag) {

		switch f.Name {

		case "requests":
			showNtpServerREQUESTCOUNT()
		case "responses":
			showNtpServerRESPONSECOUNT()
		case "dropped":
			showNtpServerREQUESTSDROPPED()

		case "broadcasts":
			showNtpServerBROADCASTCOUNT()

		case "all":
			showNtpServerREQUESTCOUNT()
			showNtpServerRESPONSECOUNT()
			showNtpServerREQUESTSDROPPED()
			showNtpServerBROADCASTCOUNT()
		default:
			fmt.Println("That does not appear to be a valid flag. Try: ", cmd.UsageString())

		}
	})

}
