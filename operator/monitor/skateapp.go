package monitor

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	// "skatechain.org/lib/logging"
	// "skatechain.org/lib/on-chain/network"
	bindingSkateApp "skatechain.org/contracts/bindings/SkateApp"
	"skatechain.org/lib/monitor"
	"skatechain.org/lib/on-chain/backend"
	"skatechain.org/operator/db"
)

func processSkateAppLog(addr common.Address, backend backend.Backend, watchOpts *bind.WatchOpts) error {
	contract, err := bindingSkateApp.NewBindingSkateApp(addr, backend)
	if err != nil {
		monitor.Logger.Error("Contract binding error: ", "error", err)
		return err
	}

	sink := make(chan *bindingSkateApp.BindingSkateAppTaskCreated)
	watcher, err := contract.WatchTaskCreated(watchOpts, sink, nil)
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
					monitor.Logger.Info("Received TaskCreated event:",
						"id", event.TaskId,
						"message", event.Message,
						"target chain", event.Chain,
						"hash", event.TaskHash,
						"source chain", chainId,
					)
				}
				task := db.SkateAppTaskToDbTask(event)
				db.InsertTask(task)
				// TODO: Operator sign the task and publish to our relayer node
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
