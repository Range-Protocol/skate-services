package cmd

import (
	"context"

	"skatechain.org/lib/logging"

	"github.com/spf13/cobra"

	"github.com/ethereum/go-ethereum/common"

	libcmd "skatechain.org/lib/cmd"
	"skatechain.org/lib/monitor"
	"skatechain.org/lib/on-chain/backend"
	"skatechain.org/lib/on-chain/network"
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
			startMonitor()

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
func startMonitor() {
	nollie := network.ChainID(5051)
	nollie_backend0, _ := backend.NewBackend("wss://nollie-rpc.skatechain.org/socket")
	nollie_SkateApp := common.HexToAddress("0x2968C1663B41Cc633540148c679f43136a4644Fc")

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
