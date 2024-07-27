package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

func RandGenCommand() *cobra.Command {
	if randgenCmd != nil {
		return randgenCmd
	}
	randgenCmd = &cobra.Command{
		Use:   appNameLowerCase,
		Short: "Create and verifiable random files of a given size",
		Run: func(cmd *cobra.Command, args []string) {
			version, _ := cmd.Flags().GetBool(versionFlag.name)
			if version {
				fmt.Println(versionInfo())
			} else {
				cmd.Help()
			}
		},
	}
	randgenCmd.Flags().BoolP(versionFlag.name, versionFlag.shorthand, false, versionFlag.usage)
	randgenCmd.AddCommand(genCommand())
	randgenCmd.AddCommand(verifyCommand())
	randgenCmd.AddCommand(serveCommand())
	randgenCmd.AddCommand(versionCommand())
	return randgenCmd
}
