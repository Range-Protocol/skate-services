package cmd

import (
	"github.com/spf13/cobra"
	"skatechain.org/lib/buildinfo"
	libcmd "skatechain.org/lib/cmd"
)

// New returns a new root cobra command that handles our command line tool.
func New() *cobra.Command {
	return libcmd.NewRootCmd(
		"operator",
		"CLI for Skate Operator",
		registerAvsCmd(),
		registerEigenLayerCmd(),
		depositToAvs(),
		monitorSkateAppCmd(),
		buildinfo.BuildInfoCmd(), // TODO: seperate package info
	)
}
