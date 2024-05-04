package main

import (
	"log"

	// bindingISkateAVS "skatechain.org/contracts/bindings/ISkateAVS"
	"skatechain.org/lib/cmd"
	// "skatechain.org/relayer/db/skateapp/disk"
	// "skatechain.org/operator/db/skateapp/disk"
)

func main() {
	config, _ := cmd.ReadEnvironmentConfig("testnet")

	// Access config values
	log.Printf("Environment: %s", config.Environment)
	log.Printf("Skate WSS RPC: %s", config.SkateWSSRPC)
	log.Printf("Skate App: %s", config.SkateApp)
	log.Printf("Skate AVS: %s", config.SkateAVS)
	log.Printf("Holesky HTTP RPC: %s", config.HoleskyHTTPRPC)
	log.Printf("Holesky WSS RPC: %s", config.HoleskyWSSRPC)
	log.Printf("Signer: %s", config.Signer)
	log.Printf("wsETH Strategy: %s", config.WsETHStrategy)
	log.Printf("Eigen Metrics IP Port: %s", config.EigenMetricsIPPort)
	log.Printf("Enable Metrics: %t", config.EnableMetrics)
	log.Printf("Node API IP Port: %s", config.NodeAPIIPPort)
	log.Printf("Enable Node API: %t", config.EnableNodeAPI)
}
