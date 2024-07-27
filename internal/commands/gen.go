package commands

import (
	"fmt"

	"dev.shib.me/randgen"
	"github.com/dustin/go-humanize"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func RandGenCommand() *cobra.Command {
	if randgenCmd != nil {
		return randgenCmd
	}
	randgenCmd = &cobra.Command{
		Use:   appNameLowerCase,
		Short: "Generate random files of a given size",
		Run: func(cmd *cobra.Command, args []string) {
			filePath := cmd.Flag(fileFlag.name).Value.String()
			sizeStr := cmd.Flag(sizeFlag.name).Value.String()
			size, err := humanize.ParseBytes(sizeStr)
			if err != nil {
				exitOnError(err)
			}
			secure, _ := cmd.Flags().GetBool(secureFlag.name)
			if err = randgen.CreateFile(filePath, int(size), secure); err != nil {
				exitOnError(err)
			}
			fmt.Printf("File %s created successfully with size %s\n", color.GreenString(filePath), color.GreenString(humanize.Bytes(size)))
		},
	}
	randgenCmd.Flags().StringP(fileFlag.name, fileFlag.shorthand, "", fileFlag.usage)
	randgenCmd.Flags().StringP(sizeFlag.name, sizeFlag.shorthand, "", sizeFlag.usage)
	randgenCmd.Flags().BoolP(secureFlag.name, secureFlag.shorthand, false, secureFlag.usage)
	randgenCmd.MarkFlagRequired(fileFlag.name)
	randgenCmd.MarkFlagRequired(sizeFlag.name)
	randgenCmd.AddCommand(verifyCommand())
	randgenCmd.AddCommand(serveCommand())
	randgenCmd.AddCommand(versionCommand())
	return randgenCmd
}
