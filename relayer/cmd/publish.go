package cmd

import (
	"context"

	"skatechain.org/lib/logging"
	"skatechain.org/relayer/publish"

	"github.com/spf13/cobra"
	libcmd "skatechain.org/lib/cmd"
)

func publishCmd() *cobra.Command {
	logger := logging.NewLoggerWithConsoleWriter()

	var configFile string
	var overrideSigner string
	var passphrase string

	cmd := &cobra.Command{
		Use:   "retrieve",
		Short: "Listen and retrieve task verification signatures from AVS operators",
		Long: `Only accepts operators who registered with Eigenlayer contracts and the Skate AVS.
    Status of the request is not available yet`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			logger.Info("Publishing to AVS ...")

			config, err := libcmd.ReadYAMLConfig(configFile)
			if err != nil {
				logger.Fatalf("Can't load config file at %s, error = %v", configFile, err)
				return err
			}

			ctx := context.WithValue(context.Background(), "config", config)
			publish.PublishTaskToAVSAndGateway(ctx)

			return nil
		},
	}

	libcmd.BindEnvConfig(cmd, &configFile)
	libcmd.BindSigner(cmd, &overrideSigner)
	libcmd.BindPassphrase(cmd, &passphrase)

	verbose := true
	libcmd.BindVerbose(cmd, &verbose)
	if !verbose {
		publish.Verbose = false
	}

	return cmd
}
