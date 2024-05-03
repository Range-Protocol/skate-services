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
		buildinfo.BuildInfoCmd(),
	)
}

func startAllCmd() *cobra.Command {
	logger := logging.NewLoggerWithConsoleWriter()

	cmd := &cobra.Command{
		Use:   "retrieve",
		Short: "Listen and retrieve task verification signatures from AVS operators",
		Long: `Only accepts operators who registered with Eigenlayer contracts and the Skate AVS.
    Status of the request is not available yet`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			logger.Info("Start signature retrieval service ...")

			return nil
		},
	}

	return cmd
}
