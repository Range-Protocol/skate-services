package cmd

// NOTE: command for operator to register with Skate AVS

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/pkg/errors"
	"skatechain.org/lib/logging"

	"github.com/spf13/cobra"

	elCliTypes "github.com/Layr-Labs/eigenlayer-cli/pkg/types"
	"gopkg.in/yaml.v3"
)

const l1BlockPeriod = time.Second * 12

type RegisterConfig struct {
	ConfigFile string
	AVSAddr    string
}

// Register registers the operator with Skate AVS. Assuming that operator address is registered with Eigen-Layer
// and configuration file with ecdsa keystore are present on disk.
// TODO: fill in the Register logic here
func Register(ctx context.Context, cfg RegisterConfig) error {
	// Default dependencies.
	return nil
}

func registerCmd() *cobra.Command {
	var cfg RegisterConfig
	logger := logging.NewLoggerWithConsoleWriter()

	cmd := &cobra.Command{
		Use:   "register",
		Short: "Register an operator with the Skate AVS",
		Long: `Register command expects a Eigen-Layer yaml config file as an argument.
    Note that the operator must already be registered with Eigen-Layer.`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			logger.Info("Registering operator to Skate AVS ...")
			err := Register(cmd.Context(), cfg)
			if err != nil {
				return errors.Wrap(err, "registration failed")
			}

			return nil
		},
	}

	bindRegisterConfig(cmd, &cfg)

	return cmd
}

func bindRegisterConfig(cobraCmd *cobra.Command, cfg *RegisterConfig) {
	const flagConfig = "config-file"
	cobraCmd.Flags().StringVar(&cfg.ConfigFile, flagConfig, cfg.ConfigFile, "Path to the Eigen-Layer yaml configuration file")
	_ = cobraCmd.MarkFlagRequired(flagConfig)
}

// readConfig returns the eigen-layer operator configuration from the given file.
func readConfig(file string) (elCliTypes.OperatorConfigNew, error) {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return elCliTypes.OperatorConfigNew{}, errors.Wrap(err, fmt.Sprintf("eigen config file not found: %s", file))
	}

	file_bytes, err := os.ReadFile(file)
	if err != nil {
		return elCliTypes.OperatorConfigNew{}, errors.Wrap(err, fmt.Sprintf("read eigen config file: %s", file))
	}

	var config elCliTypes.OperatorConfigNew
	if err := yaml.Unmarshal(file_bytes, &config); err != nil {
		return elCliTypes.OperatorConfigNew{}, errors.Wrap(err, "unmarshal eigen config file")
	}

	return config, nil
}
