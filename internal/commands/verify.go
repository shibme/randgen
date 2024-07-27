package commands

import (
	"fmt"

	"dev.shib.me/randgen"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func verifyCommand() *cobra.Command {
	if verifyCmd != nil {
		return verifyCmd
	}
	verifyCmd = &cobra.Command{
		Use:     "verify",
		Aliases: []string{"v"},
		Short:   "Verifies the integrity of a file generated by RandGen",
		Run: func(cmd *cobra.Command, args []string) {
			filePath := cmd.Flag(fileFlag.name).Value.String()
			if err := randgen.VerifyFile(filePath); err != nil {
				exitOnError(err)
			}
			fmt.Printf("File %s is verified successfully\n", color.GreenString(filePath))
		},
	}
	verifyCmd.Flags().StringP(fileFlag.name, fileFlag.shorthand, "", fileFlag.usage)
	verifyCmd.MarkFlagRequired(fileFlag.name)
	return verifyCmd
}
