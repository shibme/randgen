package commands

import (
	"strconv"

	"dev.shib.me/randgen/internal/api"
	"github.com/spf13/cobra"
)

func serveCommand() *cobra.Command {
	if serveCmd != nil {
		return serveCmd
	}
	serveCmd = &cobra.Command{
		Use:     "serve",
		Aliases: []string{"server", "api", "web"},
		Short:   "Serves RandGen as a web service",
		Run: func(cmd *cobra.Command, args []string) {
			port, _ := strconv.ParseUint(cmd.Flag(portFlag.name).Value.String(), 10, 16)
			if err := api.Serve(uint16(port)); err != nil {
				exitOnError(err)
			}
		},
	}
	serveCmd.Flags().StringP(portFlag.name, portFlag.shorthand, "", portFlag.usage)
	return serveCmd
}
