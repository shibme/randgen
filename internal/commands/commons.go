package commands

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

const (
	appNameLowerCase = "randgen"
	website          = "https://dev.shib.me/randgen"
	description      = "RandGen is a simple CLI tool to generate random files of a given size."
	art              = `

██████╗░░█████╗░███╗░░██╗██████╗░░██████╗░███████╗███╗░░██╗
██╔══██╗██╔══██╗████╗░██║██╔══██╗██╔════╝░██╔════╝████╗░██║
██████╔╝███████║██╔██╗██║██║░░██║██║░░██╗░█████╗░░██╔██╗██║
██╔══██╗██╔══██║██║╚████║██║░░██║██║░░╚██╗██╔══╝░░██║╚████║
██║░░██║██║░░██║██║░╚███║██████╔╝╚██████╔╝███████╗██║░╚███║
╚═╝░░╚═╝╚═╝░░╚═╝╚═╝░░╚══╝╚═════╝░░╚═════╝░╚══════╝╚═╝░░╚══╝
`
)

var (
	version    = ""
	commitDate = ""
	fullCommit = ""
	releaseURL = ""
	versionStr *string

	randgenCmd *cobra.Command
	verifyCmd  *cobra.Command
	serveCmd   *cobra.Command
	versionCmd *cobra.Command
)

type flagDef struct {
	name      string
	shorthand string
	usage     string
}

var (
	fileFlag = flagDef{
		name:      "file",
		shorthand: "f",
		usage:     "Name or path of the file to generate",
	}

	sizeFlag = flagDef{
		name:      "size",
		shorthand: "s",
		usage:     "Size of the file to generate",
	}

	secureFlag = flagDef{
		name:  "secure",
		usage: "Use secure random generator",
	}

	portFlag = flagDef{
		name:      "port",
		shorthand: "p",
		usage:     "Port to listen on",
	}
)

func exitOnError(err error) {
	fmt.Fprintln(os.Stderr, color.RedString(err.Error()))
	os.Exit(1)
}
