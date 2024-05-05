package main

import (
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	bindingISkateAVS "skatechain.org/contracts/bindings/ISkateAVS"
	"skatechain.org/lib/cmd"
	"skatechain.org/lib/logging"
	"skatechain.org/lib/on-chain/backend"
	// "skatechain.org/relayer/db/skateapp/disk"
	// "skatechain.org/operator/db/skateapp/disk"
)

func main() {
	logger := logging.NewLoggerWithConsoleWriter()
	config, _ := cmd.ReadConfig[cmd.EnvironmentConfig]("/environment", "testnet")

	// Access config values
	log.Printf("Environment: %s", config.Environment)
	log.Printf("Skate WSS RPC: %s", config.SkateWSSRPC)
	log.Printf("Skate App: %s", config.SkateApp)
	log.Printf("Skate AVS: %s", config.SkateAVS)
	log.Printf("Holesky HTTP RPC: %s", config.HttpRPC)
	log.Printf("Holesky WSS RPC: %s", config.WsRPC)
	log.Printf("wsETH Strategy: %s", config.WsETHStrategy)
	log.Printf("Eigen Metrics IP Port: %s", config.EigenMetricsIPPort)
	log.Printf("Enable Metrics: %t", config.EnableMetrics)
	log.Printf("Node API IP Port: %s", config.NodeAPIIPPort)
	log.Printf("Enable Node API: %t", config.EnableNodeAPI)

	be, _ := backend.NewBackend(config.HttpRPC)
	avs, _ := bindingISkateAVS.NewBindingISkateAVS(common.HexToAddress(config.SkateAVS), be)

	result, _ := avs.GetRestakeableStrategies(&bind.CallOpts{})

	logger.Info("Result", "addresses", result)
}
