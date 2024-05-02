package main

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"skatechain.org/lib/logging"
	"skatechain.org/lib/on-chain/backend"
	diskDb "skatechain.org/relayer/db/skateapp/disk"

	bindingISkateAVS "skatechain.org/contracts/bindings/ISkateAVS"
)

func main() {
	logger := logging.NewLoggerWithConsoleWriter()
	tasks, _ := diskDb.SkateApp_RetriveSignedTasks("SELECT * FROM skateapp_signed_tasks")
	for _, task := range tasks {
		logger.Info("Task", "task", task)
	}

	backend, _ := backend.NewBackend("wss://holesky.drpc.org")
	avsAddress := common.HexToAddress("0x2a0D46ED3D9D13F6a9b5B0D3274675143c803071")

	contract, _ := bindingISkateAVS.NewBindingISkateAVS(avsAddress, backend)
	result, _ := contract.GetRestakeableStrategies(&bind.CallOpts{})
	for _, addr := range result {
		logger.Info("AVS allowed strategy", "address", addr)
	}
}
