package commands

import (
	"fmt"
	"os"

	"dev.shib.me/randgen"
	"github.com/dustin/go-humanize"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func genCommand() *cobra.Command {
	if genCmd != nil {
		return genCmd
	}
	genCmd = &cobra.Command{
		Use:     "gen",
		Aliases: []string{"generate", "g", "create", "make"},
		Short:   "Generate random files of a given size",
		Run: func(cmd *cobra.Command, args []string) {
			filePath := cmd.Flag(fileFlag.name).Value.String()
			sizeStr := cmd.Flag(sizeFlag.name).Value.String()
			size, err := humanize.ParseBytes(sizeStr)
			if err != nil {
				exitOnError(err)
			}
			secure, _ := cmd.Flags().GetBool(secureFlag.name)
			if filePath == "" {
				if err = randgen.WriteRand(os.Stdout, int(size), secure); err != nil {
					exitOnError(err)
				}
			} else {
				if err = randgen.CreateFile(filePath, int(size), secure); err != nil {
					exitOnError(err)
				}
				fmt.Printf("Random file of size %s created at %s\n", humanize.Bytes(uint64(size)), color.GreenString(filePath))
			}
		},
	}
	genCmd.Flags().StringP(fileFlag.name, fileFlag.shorthand, "", fileFlag.usage)
	genCmd.Flags().StringP(sizeFlag.name, sizeFlag.shorthand, "", sizeFlag.usage)
	genCmd.Flags().BoolP(secureFlag.name, secureFlag.shorthand, false, secureFlag.usage)
	genCmd.MarkFlagRequired(sizeFlag.name)
	return genCmd
}
