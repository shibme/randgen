package main

import (
	"fmt"
	"os"

	"dev.shib.me/randgen/internal/commands"
	"github.com/fatih/color"
)

func main() {
	if err := commands.RandGenCommand().Execute(); err != nil {
		fmt.Fprintln(os.Stderr, color.RedString(err.Error()))
		os.Exit(1)
	}
}
