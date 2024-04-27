package cmd

// NOTE: command for operator to register with Skate AVS

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/pkg/errors"
	"skate.org/avs/lib/logging"

	geth "github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"

	elCliTypes "github.com/Layr-Labs/eigenlayer-cli/pkg/types"
	elCliUtils "github.com/Layr-Labs/eigenlayer-cli/pkg/utils"
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

func avsAddressOrDefault(avsAddr string, chainID *big.Int) (geth.Address, error) {
	var resp geth.Address
	if avsAddr != "" {
		if !geth.IsHexAddress(avsAddr) {
			return geth.Address{}, errors.New(fmt.Sprintf("invalid avs address: %s", avsAddr))
		}
		resp = geth.HexToAddress(avsAddr)
	} else if addr, ok := skateAvsAddressOnChain(chainID); ok {
		resp = addr
	} else {
		return geth.Address{}, errors.New(fmt.Sprintf("avs address not provided and no default for chain found: %d", chainID.Uint64()))
	}

	return resp, nil
}

// TODO: replace with our AVS address
func skateAvsAddressOnChain(chainID *big.Int) (geth.Address, bool) {
	switch chainID.Int64() {
	case elCliUtils.HoleskyChainId:
		return geth.HexToAddress(""), true
	case elCliUtils.MainnetChainId:
		return geth.HexToAddress(""), true
	default:
		return geth.Address{}, false
	}
}

func registerCmd() *cobra.Command {
	var cfg RegisterConfig
	logger := logging.NewLoggerWithConsoleWriter()

	cmd := &cobra.Command{
		Use:   "register",
		Short: "Register an operator with the Skate AVS",
		Long: `Register command expects a Eigen-Layer yaml config file as an argument to 
    successfully register an operator address with the Skate AVS. 
    Note the operator must already be registered with Eigen-Layer.`,
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
