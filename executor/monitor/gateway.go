package monitor

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"skate.org/avs/lib/logging"
	"skate.org/avs/lib/network"
	"skate.org/avs/lib/on-chain/backend"
)

type SkateGatewayMonitor struct {
	ctx           context.Context
	network       network.Network
	contractAddrs map[network.ChainID]common.Address
	backends      map[network.ChainID][]backend.Backend
	logger        *logging.Logger
	verbose       *bool
}

func NewSkateGatewayMonitor(
	ctx context.Context, network network.Network, skateAppAddrs map[network.ChainID]common.Address,
	backends map[network.ChainID][]backend.Backend,
) (*SkateGatewayMonitor, *bool) {
	logger := logging.NewLoggerWithConsoleWriter()
	verboseControl := true

	return &SkateGatewayMonitor{
		ctx:           ctx,
		network:       network,
		contractAddrs: skateAppAddrs,
		backends:      backends,
		logger:        logger,
		verbose:       &verboseControl,
	}, &verboseControl
}

func (m *SkateGatewayMonitor) SetVerbose(v bool) {
	*m.verbose = v
}

func (m *SkateGatewayMonitor) Start() {
	watchOpts := &bind.WatchOpts{
		Start:   nil,
		Context: m.ctx,
	}

	for chainID, addr := range m.contractAddrs {
		backends := m.backends[chainID]
		for _, backend := range backends {
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

func (m *SkateGatewayMonitor) watchWithBackend(addr common.Address, backend backend.Backend, watchOpts *bind.WatchOpts) error {
	return nil
}
