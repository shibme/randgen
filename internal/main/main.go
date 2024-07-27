package main

import (
	"fmt"
	"os"
	"strings"

	"dev.shib.me/randgen/internal/api"
	"dev.shib.me/randgen/internal/commands"
	"github.com/fatih/color"
)

func exitOnError(err error) {
	fmt.Fprintln(os.Stderr, color.RedString(err.Error()))
	os.Exit(1)
}

const (
	envar_RANDGEN_SERVER_ENABLED = "RANDGEN_SERVER_ENABLED"
)

func main() {
	if strings.ToLower(os.Getenv(envar_RANDGEN_SERVER_ENABLED)) == "true" {
		if err := api.Serve(0); err != nil {
			exitOnError(err)
		}
	} else if err := commands.RandGenCommand().Execute(); err != nil {
		exitOnError(err)
	}
}
