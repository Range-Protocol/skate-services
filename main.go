package main

import (
	"log"

	"skatechain.org/lib/cmd"
	bindingISkateAVS "skatechain.org/contracts/bindings/ISkateAVS"
	"skatechain.org/relayer/db/skateapp/disk"
)

func main() {
	config, _ := cmd.ReadYAMLConfig("testnet")

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

type SignatureTuple = bindingISkateAVS.ISkateAVSSignatureTuple

func fetchSignaturesForTask(task disk.SignedTask) ([]SignatureTuple, error) {
	var signatures []SignatureTuple

	rows, err := disk.SkateAppDB.Query(
		`SELECT operator, signature FROM ? WHERE 
    taskId = ? AND
    chainId = ? AND
    chainType = ?`,
		disk.SignedTaskSchema,
		task.TaskId,
		task.ChainId,
		task.ChainType,
	)
	if err != nil {
		return signatures, err
	}

	for rows.Next() {
		var sigTuple SignatureTuple
		err := rows.Scan(
			&sigTuple.Signature, &sigTuple.Operator,
		)
		if err != nil {
			return nil, err
		}
		signatures = append(signatures, sigTuple)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return signatures, nil
}
