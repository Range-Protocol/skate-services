package main

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	// "skatechain.org/lib/crypto/ecdsa"
	"skatechain.org/lib/logging"
	"skatechain.org/lib/on-chain/backend"

	bindingISkateAVS "skatechain.org/contracts/bindings/ISkateAVS"
)

func main() {
	// tasks, _ := diskDb.SkateApp_RetriveSignedTasks("SELECT * FROM skateapp_signed_tasks")
	// for _, task := range tasks {
	// 	logger.Info("Task", "task", task)
	// }

	logger := logging.NewLoggerWithConsoleWriter()
	be, _ := backend.NewBackend("wss://holesky.drpc.org")
	avsAddress := common.HexToAddress("0x2a0D46ED3D9D13F6a9b5B0D3274675143c803071")

	contract, _ := bindingISkateAVS.NewBindingISkateAVS(avsAddress, be)
	result, _ := contract.GetRestakeableStrategies(&bind.CallOpts{})
	for _, addr := range result {
		logger.Info("AVS allowed strategy", "address", addr)
	}

	// backend.DumpECDSAPrivateKeyToKeystore(
	// 	"c693966183ef1b7f062bd4bbde280dbc7e5d310ee7537dbad7374263a9860a05",
	// 	"hello_world",
	// )

  address := common.HexToAddress("0x786775c9ecB916bd7f5a59c150491871fCfCEe86")
  privKey, _ := backend.PrivateKeyFromKeystore(address, "hello_world")
  logger.Info("Private key", "private key", privKey)
}
