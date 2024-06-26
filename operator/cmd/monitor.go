package cmd

import (
	"context"
	"fmt"

	"skatechain.org/lib/logging"

	"github.com/spf13/cobra"

	"github.com/ethereum/go-ethereum/common"

	libcmd "skatechain.org/lib/cmd"
	"skatechain.org/lib/monitor"
	"skatechain.org/lib/on-chain/backend"
	"skatechain.org/lib/on-chain/network"
	skateappDb "skatechain.org/operator/db/skateapp/disk"
	operatorMonitor "skatechain.org/operator/monitor"
)

// TODO: decouple monitor from signing service
func monitorSkateAppCmd() *cobra.Command {
	logger := logging.NewLoggerWithConsoleWriter()

	var envConfigFile string
	var signerConfigFile string
	var overrideSigner string
	var passphrase string

	cmd := &cobra.Command{
		Use:   "monitor",
		Short: "Monitor TaskCreated events from Skate AVS, verify and dispatch to relayer",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			envConfig, err := libcmd.ReadConfig[libcmd.EnvironmentConfig]("/environment", envConfigFile)
			if err != nil {
				logger.Fatalf("Can't load config file at %s, error = %v", envConfigFile, err)
				return err
			}
			ctx := context.WithValue(context.Background(), "config", envConfig)

			signerConfig, err := libcmd.ReadConfig[libcmd.SignerConfig]("/signer/operator", signerConfigFile)
			if overrideSigner != "" {
				signerConfig.Address = overrideSigner
			}
			if passphrase != "" {
				signerConfig.Passphrase = passphrase
			}

			if signerConfig.Address != "" {
				ctx = context.WithValue(ctx, "signer", signerConfig)

				_, err := backend.PrivateKeyFromKeystore(common.HexToAddress(signerConfig.Address), signerConfig.Passphrase)
				if err != nil {
					logger.Fatal("Invalid keystore for signer", signerConfig)
					return err
				}
				logger.Info("Operator: monitoring and processing tasks ..",
					"signer", signerConfig.Address,
					"fromConfig", fmt.Sprintf("configs/signer/operator/%s.yaml", signerConfigFile),
				)
			} else {
				logger.Info("No signer provided, running with watch-only mode...")
			}

			startMonitor(ctx)

			return nil
		},
	}

	libcmd.BindEnvConfig(cmd, &envConfigFile)
	libcmd.BindSignerConfig(cmd, &signerConfigFile)
	libcmd.BindSigner(cmd, &overrideSigner)
	libcmd.BindPassphrase(cmd, &passphrase)

	verbose := true
	libcmd.BindVerbose(cmd, &verbose)
	if !verbose {
		monitor.Verbose = false
	}

	return cmd
}

// TODO: populate context to the runner
func startMonitor(ctx context.Context) {
	env := ctx.Value("config").(*libcmd.EnvironmentConfig)
	nollie := network.ChainID(env.SkateChainId)
	nollie_backend0, _ := backend.NewBackend(env.SkateWSSRPC)
	nollie_SkateApp := common.HexToAddress(env.SkateApp)

	skateappDb.InitializeSkateApp()

	contractAddrs := map[network.ChainID]common.Address{
		nollie: nollie_SkateApp,
	}

	backends := map[network.ChainID][]backend.Backend{
		nollie: {nollie_backend0},
	}

	ctxs := map[network.ChainID]context.Context{
		nollie: ctx,
	}

	monitor := monitor.NewMonitor(ctxs, contractAddrs, backends)
	monitor.Start(operatorMonitor.SubscribeSkateApp)
}
