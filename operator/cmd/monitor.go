package cmd

import (
	"context"

	"skatechain.org/lib/logging"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/ethereum/go-ethereum/common"

	libcmd "skatechain.org/lib/cmd"
	"skatechain.org/lib/monitor"
	"skatechain.org/lib/on-chain/backend"
	"skatechain.org/lib/on-chain/network"
	"skatechain.org/operator/db"
	oMonitor "skatechain.org/operator/monitor"
)

func monitorSkateAppCmd() *cobra.Command {
	logger := logging.NewLoggerWithConsoleWriter()

	cmd := &cobra.Command{
		Use:   "monitor-skateapp",
		Short: "Monitor TaskCreated events from Skate AVS, verify and dispatch to relayer",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			logger.Info("listening to Skate App request...")
			nollie_rpc := viper.GetString("holesky_wss_rpc")
			skateapp_addr := viper.GetString("skate_app")
			startMonitor(nollie_rpc, skateapp_addr)

			return nil
		},
	}

	verbose := true
	libcmd.BindVerbose(cmd, &verbose)
	if !verbose {
		monitor.Verbose = false
	}

	return cmd
}

// TODO: bind with config files
func startMonitor(rpc string, skateapp string) {
	nollie := network.ChainID(5051)
	nollie_backend0, _ := backend.NewBackend(rpc)
	nollie_SkateApp := common.HexToAddress(skateapp)

	db.InitializeSkateApp()

	contractAddrs := map[network.ChainID]common.Address{
		nollie: nollie_SkateApp,
	}

	backends := map[network.ChainID][]backend.Backend{
		nollie: {nollie_backend0},
	}

	ctx := map[network.ChainID]context.Context{
		nollie: context.Background(),
	}

	monitor := monitor.NewMonitor(ctx, contractAddrs, backends)
	monitor.Start(oMonitor.SubscribeSkateApp)
}
