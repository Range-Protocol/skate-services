package monitor

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"skate.org/avs/contracts/bindings/SkateApp"
	"skate.org/avs/lib/logging"
	"skate.org/avs/lib/network"
	"skate.org/avs/lib/on-chain/backend"
)

type SkateAppMonitor struct {
	ctx           context.Context
	network       network.Network
	contractAddrs map[network.ChainID]common.Address
	backends      map[network.ChainID][]backend.Backend
	logger        *logging.Logger
	verbose       *bool
}

func NewMonitor(
	ctx context.Context, network network.Network,
	skateAppAddrs map[network.ChainID]common.Address,
	backends map[network.ChainID][]backend.Backend,
) (*SkateAppMonitor, *bool) {
	logger := logging.NewLoggerWithConsoleWriter()
	verboseControl := true

	return &SkateAppMonitor{
		ctx:           ctx,
		network:       network,
		contractAddrs: skateAppAddrs,
		backends:      backends,
		logger:        logger,
		verbose:       &verboseControl,
	}, &verboseControl
}

func (m *SkateAppMonitor) SetVerbose(v bool) {
	*m.verbose = v
}

func (m *SkateAppMonitor) Start() {
	watchOpts := &bind.WatchOpts{
		Start:   nil,
		Context: m.ctx,
	}

	for chainID, addr := range m.contractAddrs {
		backends := m.backends[chainID]
		for _, backend := range backends {
			// Attempt to use the current backend
			if err := m.watchWithBackend(addr, backend, watchOpts); err == nil {
				// If successful, move to the next address
				break
			} else {
				// Log the error if verbose mode is enabled
				if *m.verbose {
					m.logger.Error("Error using backend:", err)
				}
			}
		}
	}
}

func (m *SkateAppMonitor) watchWithBackend(addr common.Address, backend backend.Backend, watchOpts *bind.WatchOpts) error {
	contract, err := bindingSkateApp.NewBindingSkateApp(addr, backend)
	if err != nil {
		return err
	}

	sink := make(chan *bindingSkateApp.BindingSkateAppTaskCreated)
	watcher, err := contract.WatchTaskCreated(watchOpts, sink, nil)
	if err != nil {
		return err
	}

	// Handle event notifications in a separate goroutine
	go func() {
		for {
			select {
			case event, ok := <-sink:
				if !ok {
					return
				}
				if *m.verbose {
					m.logger.Info("Received TaskCreated event:", event)
				}
			case err := <-watcher.Err():
				if err != nil && *m.verbose {
					m.logger.Error("Watcher error:", err)
				}
				return
			}
		}
	}()

	return nil
}
