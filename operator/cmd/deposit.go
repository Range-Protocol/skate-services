package cmd

// NOTE: command for operator to deposit token into a strategy
import (
	"math/big"

	"skatechain.org/lib/logging"

	"github.com/spf13/cobra"

	"github.com/ethereum/go-ethereum/common"

	libcmd "skatechain.org/lib/cmd"
	"skatechain.org/lib/monitor"
	"skatechain.org/lib/on-chain/backend"
	"skatechain.org/operator/deposit"
)

func depositToAvs() *cobra.Command {
	logger := logging.NewLoggerWithConsoleWriter()

	var envConfigFile string
	var signerConfigFile string
	var overrideSigner string
	var passphrase string

	// params
	var amount uint64
	var strategy string
	var token string
	var autoApprove bool

	cmd := &cobra.Command{
		Use:   "deposit",
		Short: "Monitor TaskCreated events from Skate AVS, verify and dispatch to relayer",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			signerConfig, err := libcmd.ReadConfig[libcmd.SignerConfig]("/signer/operator", signerConfigFile)
			if overrideSigner != "" {
				signerConfig.Address = overrideSigner
			}
			if passphrase != "" {
				signerConfig.Passphrase = passphrase
			}
			if signerConfig.Address == "" {
				logger.Fatal("No signer provided")
			}
			privateKey, err := backend.PrivateKeyFromKeystore(common.HexToAddress(signerConfig.Address), signerConfig.Passphrase)
			if err != nil {
				logger.Fatal("Invalid keystore for signer", signerConfig)
				return err
			}

			envConfig, err := libcmd.ReadConfig[libcmd.EnvironmentConfig]("/environment", envConfigFile)
			if err != nil {
				logger.Fatalf("Can't load config file at %s, error = %v", envConfigFile, err)
				return err
			}
			if strategy == "" {
				strategy = envConfig.Strategy_stETH
			}
			if token == "" {
				token = envConfig.Token_stETH
			}
			logger.Info("Deposit into Eigenlayer staked strategy...", "token", token, "strategy", strategy, "amount", amount)

			deposit.DepositIntoStrategy(
				privateKey,
				*envConfig,
				common.HexToAddress(strategy),
				common.HexToAddress(token),
				new(big.Int).SetUint64(amount),
				autoApprove,
			)

			return nil
		},
	}

	libcmd.BindEnvConfig(cmd, &envConfigFile)
	libcmd.BindSignerConfig(cmd, &signerConfigFile)
	libcmd.BindSigner(cmd, &overrideSigner)
	libcmd.BindPassphrase(cmd, &passphrase)

	cmd.Flags().Uint64VarP(&amount, "amount", "a", 2, "Amount to deposit")
	cmd.Flags().StringVarP(&strategy, "strategy", "s", "", "Staking strategy to participate in")
	cmd.Flags().StringVarP(&token, "token", "t", "", "Token for staking strategy")
	cmd.Flags().BoolVar(&autoApprove, "auto-approve", false, "Whether to auto set allowance to amount if allowance is insufficient")

	verbose := true
	libcmd.BindVerbose(cmd, &verbose)
	if !verbose {
		monitor.Verbose = false
	}

	return cmd
}
