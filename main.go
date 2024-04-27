package main

import (
	"context"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"

	"skate.org/avs/contracts/bindings/IERC20"
	"skate.org/avs/lib/logging"
	"skate.org/avs/lib/network"
	"skate.org/avs/lib/on-chain/backend"
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

	monitor, control := NewMonitor(ctx, contractAddrs, backends)
	*control = true
	monitor.Start()
}

type ExampleMonitor struct {
	ctx           map[network.ChainID]context.Context
	contractAddrs map[network.ChainID]common.Address
	backends      map[network.ChainID][]backend.Backend
	logger        *logging.Logger
	verbose       *bool
}

func NewMonitor(
	ctx map[network.ChainID]context.Context,
	skateAppAddrs map[network.ChainID]common.Address,
	backends map[network.ChainID][]backend.Backend,
) (*ExampleMonitor, *bool) {
	logger := logging.NewLoggerWithConsoleWriter()
	verboseControl := true

	return &ExampleMonitor{
		ctx:           ctx,
		contractAddrs: skateAppAddrs,
		backends:      backends,
		logger:        logger,
		verbose:       &verboseControl,
	}, &verboseControl
}

func (m *ExampleMonitor) SetVerbose(v bool) {
	*m.verbose = v
}

func (m *ExampleMonitor) Start() {
	var wg sync.WaitGroup
	wg.Add(len(m.contractAddrs))
	for chainID, addr := range m.contractAddrs {
		go func(chainID network.ChainID, addr common.Address) {
			defer wg.Done()

			if *m.verbose {
				m.logger.Info("Listening on chain", "chainID", chainID)
			}
			backends := m.backends[chainID]
			latest := uint64(0)
			ctx := m.ctx[chainID]
			watchOpts := &bind.WatchOpts{
				Start:   &latest,
				Context: ctx,
			}
			for id, backend := range backends {
				// Attempt to use the current backend
				latest, _ = backend.BlockNumber(ctx)
				if err := m.watchWithBackend(addr, backend, watchOpts); err == nil {
					break
				} else {
					if *m.verbose {
						m.logger.Error("Error using backend:", "id", id)
					}
				}
			}
		}(chainID, addr)
	}
	wg.Wait()
}

func (m *ExampleMonitor) watchWithBackend(addr common.Address, backend backend.Backend, watchOpts *bind.WatchOpts) error {
	retries := 2
	for {
		// Attempt to connect and watch for events
		if err := m.watchOnce(addr, backend, watchOpts); err != nil {
			m.logger.Error("Error watching for events:", "error", err, "retry left", retries)
			retries -= 1
			if retries < 0 {
				return errors.New("Retries time out!")
			}
		}

		// Wait before attempting to reconnect
		time.Sleep(time.Second * 3)
	}
}

func (m *ExampleMonitor) watchOnce(addr common.Address, backend backend.Backend, watchOpts *bind.WatchOpts) error {
	contract, err := bindingIERC20.NewBindingIERC20(addr, backend)
	if err != nil {
		m.logger.Error("Contract binding error: ", "error", err)
		return err
	}

	sink := make(chan *bindingIERC20.BindingIERC20Transfer)
	watcher, err := contract.WatchTransfer(watchOpts, sink, nil, nil)
	if err != nil {
		if *m.verbose {
			m.logger.Error("Watcher initialization error: ", "error", err)
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
				if *m.verbose {
					chainId, _ := backend.ChainID(context.Background())
					m.logger.Info("Received Transfer event:", "from", event.From, "to", event.To, "amount", event.Value, "chain", chainId)
				}
			case err := <-watcher.Err():
				if err != nil && *m.verbose {
					m.logger.Error("Watcher received error: ", "error", err)
				}
				return
			}
		}
	}()

	// Wait for the watcher to be closed or an error to occur
	<-watcher.Err()
	return nil
}
