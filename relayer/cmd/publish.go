package cmd

import (
	"context"

	"skatechain.org/lib/logging"
	"skatechain.org/relayer/publish"

	"github.com/spf13/cobra"
)

func publishCmd() *cobra.Command {
	logger := logging.NewLoggerWithConsoleWriter()

	cmd := &cobra.Command{
		Use:   "retrieve",
		Short: "Listen and retrieve task verification signatures from AVS operators",
		Long: `Only accepts operators who registered with Eigenlayer contracts and the Skate AVS.
    Status of the request is not available yet`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			logger.Info("Publishing to AVS ...")
      publish.PublishTaskToAVS(context.Background())

			return nil
		},
	}

	return cmd
}
