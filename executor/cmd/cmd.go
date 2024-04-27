package cmd

import (
	"github.com/spf13/cobra"
	"skate.org/avs/lib/buildinfo"
	libcmd "skate.org/avs/lib/cmd"
)

// New returns a new root cobra command that handles our command line tool.
func New() *cobra.Command {
	return libcmd.NewRootCmd(
		"executor",
		"CLI for Skate Executor",
		buildinfo.BuildInfoCmd(),
	)
}
