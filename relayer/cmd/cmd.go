package cmd

import (
	"github.com/spf13/cobra"
	"skatechain.org/lib/buildinfo"
	libcmd "skatechain.org/lib/cmd"
	"skatechain.org/lib/logging"
)

// New returns a new root cobra command that handles our command line tool.
func New() *cobra.Command {
	return libcmd.NewRootCmd(
		"relayer",
		"CLI for Skate Relayer",
		retrieveCmd(),
    publishCmd(),
		buildinfo.BuildInfoCmd(),
	)
}

func startAllCmd() *cobra.Command {
	logger := logging.NewLoggerWithConsoleWriter()

	cmd := &cobra.Command{
		Use:   "retrieve",
		Short: "Listen and retrieve task verification signatures from AVS operators",
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			logger.Info("Not implemented!")

			return nil
		},
	}

	return cmd
}
