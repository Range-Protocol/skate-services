package main

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"skatechain.org/contracts/bindings/IERC20"
	"skatechain.org/lib/monitor"
	"skatechain.org/lib/on-chain/backend"
	"skatechain.org/lib/on-chain/network"
)

// NOTE: example of USDC monitor service on [mainnet, polygon]
func main() {
	mainnet := network.ChainID(1)
	mainnet_backend0, _ := backend.NewBackend("wss://ethereum-rpc.publicnode.com")
	mainnet_USDC := common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48")

	polygon := network.ChainID(137)
	polygon_backend0, _ := backend.NewBackend("wss://polygon-bor-rpc.publicnode.com")
	polygon_USDC := common.HexToAddress("0x3c499c542cEF5E3811e1192ce70d8cC03d5c3359")

	contractAddrs := map[network.ChainID]common.Address{
		mainnet: mainnet_USDC,
		polygon: polygon_USDC,
	}

	backends := map[network.ChainID][]backend.Backend{
		mainnet: {mainnet_backend0},
		polygon: {polygon_backend0},
	}

	ctx := map[network.ChainID]context.Context{
		mainnet: context.Background(),
		polygon: context.Background(),
	}

	monitor := monitor.NewMonitor(ctx, contractAddrs, backends)
	monitor.Start(processLog)
}

func processLog(addr common.Address, backend backend.Backend, watchOpts *bind.WatchOpts) error {
	contract, err := bindingIERC20.NewBindingIERC20(addr, backend)
	if err != nil {
		monitor.Logger.Error("Contract binding error: ", "error", err)
		return err
	}

	sink := make(chan *bindingIERC20.BindingIERC20Transfer)
	watcher, err := contract.WatchTransfer(watchOpts, sink, nil, nil)
	if err != nil {
		if monitor.Verbose {
			monitor.Logger.Error("Watcher initialization error: ", "error", err)
		}
		return err
	}

	// Event handler
	go func() {
		for {
			select {
			case event, ok := <-sink:
				if !ok {
					return
				}
				if monitor.Verbose {
					chainId, _ := backend.ChainID(context.Background())
					monitor.Logger.Info("Received Transfer event:", "from", event.From, "to", event.To, "amount", event.Value, "chain", chainId)
				}
			case err := <-watcher.Err():
				if err != nil && monitor.Verbose {
					monitor.Logger.Error("Watcher received error: ", "error", err)
				}
				return
			}
		}
	}()

	// Wait for the watcher to be closed or an error to occur
	<-watcher.Err()
	return nil
}
