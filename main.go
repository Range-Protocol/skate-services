package main

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	bindingISkateAVS "skatechain.org/contracts/bindings/ISkateAVS"
	"skatechain.org/lib/cmd"
	"skatechain.org/lib/crypto/ecdsa"
	"skatechain.org/lib/logging"
	"skatechain.org/lib/on-chain/backend"
	"skatechain.org/relayer/db/skateapp/disk"
	// "skatechain.org/operator/db/skateapp/disk"
)

func main() {
	logger := logging.NewLoggerWithConsoleWriter()
	config, _ := cmd.ReadConfig[cmd.EnvironmentConfig]("/environment", "testnet")

	be, _ := backend.NewBackend(config.HttpRPC)
	avs, _ := bindingISkateAVS.NewBindingISkateAVS(common.HexToAddress(config.SkateAVS), be)

	result, _ := avs.Operators(&bind.CallOpts{})

	logger.Info("Skaate AVS existing operators", "count", len(result), "addresses", result)

  signedTasks, _ := disk.RetrieveSignedTasks()
	logger.Info("Signed tasks", "count", len(signedTasks))

  bytesA := []byte("A")
  bytesB := []byte("B")
  hashA := ecdsa.Keccak256(bytesA, bytesB)
  hashB := ecdsa.Keccak256(append(bytesA, bytesB...))

	logger.Info("Test hash", "hashA", hashA, "hashB", hashB)
}
