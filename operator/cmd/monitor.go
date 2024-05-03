package cmd

import (
	"context"

	"skatechain.org/lib/logging"

	"github.com/spf13/cobra"
	// "github.com/spf13/viper"

	"github.com/ethereum/go-ethereum/common"

	libcmd "skatechain.org/lib/cmd"
	"skatechain.org/lib/monitor"
	"skatechain.org/lib/on-chain/backend"
	"skatechain.org/lib/on-chain/network"
	"skatechain.org/operator/db"
	oMonitor "skatechain.org/operator/monitor"
)

// TODO: decouple monitor from signing service
func monitorSkateAppCmd() *cobra.Command {
	logger := logging.NewLoggerWithConsoleWriter()

	var configFile string
	var overrideSigner string
	var passphrase string

	cmd := &cobra.Command{
		Use:   "monitor",
		Short: "Monitor TaskCreated events from Skate AVS, verify and dispatch to relayer",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			logger.Info("Listening to Skate App request...")

			config, err := libcmd.ReadYAMLConfig(configFile)
			if err != nil {
				logger.Fatalf("Can't load config file at %s, error = %v", configFile, err)
				return err
			}

			ctx := context.WithValue(context.Background(), "config", config)
			signer := config.Signer
			if overrideSigner != "" {
				signer = overrideSigner
			}
			signerConfig := &libcmd.SignerConfig{
				Address:    signer,
				Passphrase: passphrase,
			}
			if signer == "" {
				logger.Info("No signer provided, run with read-only mode")
			} else {
				ctx = context.WithValue(ctx, "signer", signerConfig)

				_, err := backend.PrivateKeyFromKeystore(common.HexToAddress(signer), passphrase)
				if err != nil {
					logger.Fatal("Invalid keystore for signer", "address", signer, "passphrase", passphrase)
					return err
				}
			}

			startMonitor(ctx)

			return nil
		},
	}

	libcmd.BindEnvConfig(cmd, &configFile)
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
	nollie := network.ChainID(env.ChainId)
	nollie_backend0, _ := backend.NewBackend(env.SkateWSSRPC)
	nollie_SkateApp := common.HexToAddress(env.SkateApp)

	db.InitializeSkateApp()

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
	monitor.Start(oMonitor.SubscribeSkateApp)
}
