package monitor

import (
	"context"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"

	"skatechain.org/lib/logging"
	"skatechain.org/lib/on-chain/backend"
	"skatechain.org/lib/on-chain/network"
)

var (
	Logger  = logging.NewLoggerWithConsoleWriter()
	Verbose = true
  Retries = 2
)

// ProcessLogFunc is a function type for watching events.
type ProcessLogFunc func(addr common.Address, backend backend.Backend, watchOpts *bind.WatchOpts) error

type Monitor struct {
	ctx           map[network.ChainID]context.Context
	contractAddrs map[network.ChainID]common.Address
	backends      map[network.ChainID][]backend.Backend
}

func NewMonitor(
	ctx map[network.ChainID]context.Context,
	skateAppAddrs map[network.ChainID]common.Address,
	backends map[network.ChainID][]backend.Backend,
) *Monitor {
	return &Monitor{
		ctx:           ctx,
		contractAddrs: skateAppAddrs,
		backends:      backends,
	}
}

func (m *Monitor) Start(processLog ProcessLogFunc) {
	var wg sync.WaitGroup
	wg.Add(len(m.contractAddrs))

	for chainID, addr := range m.contractAddrs {
		go func(chainID network.ChainID, addr common.Address) {
			defer wg.Done()

			if Verbose {
				Logger.Info("Listening on chain", "chainID", chainID)
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
				if err := m.watchWithBackend(processLog, addr, backend, watchOpts); err == nil {
					break
				} else {
					if Verbose {
						Logger.Error("Error using backend:", "backendId", id, "chainId", chainID, "rpcUrl")
					}
				}
			}
		}(chainID, addr)
	}

	wg.Wait()
}

func (m *Monitor) watchWithBackend(processLog ProcessLogFunc, addr common.Address, backend backend.Backend, watchOpts *bind.WatchOpts) error {
	retries := Retries
	for {
		// Attempt to connect and watch for events
		if err := processLog(addr, backend, watchOpts); err != nil {
			Logger.Error("Error watching for events:", "error", err, "retry left", retries)
			retries -= 1
			if retries < 0 {
				return errors.New("Retries time out!")
			}
		}

		time.Sleep(time.Second * 1)
	}
}
