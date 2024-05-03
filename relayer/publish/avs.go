package publish

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	pb "skatechain.org/api/pb/relayer"
	bindingISkateAVS "skatechain.org/contracts/bindings/ISkateAVS"
	libcmd "skatechain.org/lib/cmd"
	"skatechain.org/lib/on-chain/avs"
	"skatechain.org/lib/on-chain/backend"
	"skatechain.org/relayer/db/skateapp/disk"
)

func PublishTaskToAVS(ctx context.Context) {
	ticker := time.NewTicker(12 * time.Second)
	defer ticker.Stop()
	relayerLogger.Info("Start AVS publisher process...")

	config := ctx.Value("config").(*libcmd.EnvironmentConfig)
	holeskyBackend, err := backend.NewBackend(config.SkateHttpRPC)
	if err != nil {
		relayerLogger.Fatal("AVS rpc error", "rpcUrl", config.SkateHttpRPC)
		return
	}
	holeskyAvsContract, err := bindingISkateAVS.NewBindingISkateAVS(
		common.HexToAddress(config.SkateAVS), holeskyBackend,
	)
	if err != nil {
		relayerLogger.Fatal("Invalid avs contract", "address", config.SkateAVS)
		return
	}

	signer := ctx.Value("signer").(*libcmd.SignerConfig)
	privateKey, _ := backend.PrivateKeyFromKeystore(common.HexToAddress(signer.Address), signer.Passphrase)

	for {
		select {
		case <-ctx.Done():
			relayerLogger.Info("AVS publish process stopped!")
			return
		case <-ticker.C:
			tasks := fetchSignedTasks()
			for _, task := range tasks {
				// NOTE: use memcache in future versions
				signatures, _ := fetchSignaturesForTask(task)
				operators, _ := holeskyAvsContract.Operators(&bind.CallOpts{})
				operatorCounts := len(operators)

				// Reference from SkateAVS contract, should change to (op * 2 + 2) / 3
				if len(signatures)*10_000 > operatorCounts*6_666 {
					taskId := new(big.Int).SetUint64(uint64(task.TaskId))
					messageData := avs.TaskData(task.Message, task.Initiator, pb.ChainType_EVM, task.ChainId)

					// Create a transactor to submit on-chain
					chainId := new(big.Int).SetUint64(uint64(task.ChainId))
					txOptsWithSigner, _ := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
					tx, err := holeskyAvsContract.BindingISkateAVSTransactor.SubmitData(
						txOptsWithSigner,
						taskId,
						messageData,
						signatures,
					)
					if err != nil {
						relayerLogger.Error("Failed to submit transaction", "error", err)
						continue
					}
					relayerLogger.Info("Transaction sent", "txHash", tx.Hash().Hex())

					// Collect the transaction receipt
					receipt, err := holeskyBackend.TransactionReceipt(ctx, tx.Hash())
					if err != nil {
						relayerLogger.Error("Failed to get transaction receipt", "error", err)
						continue
					}
					relayerLogger.Info("Transaction receipt", "status", receipt.Status)
				}
			}
		}
	}
}

type SignatureTuple = bindingISkateAVS.ISkateAVSSignatureTuple

func fetchSignaturesForTask(task disk.SignedTask) ([]SignatureTuple, error) {
	var signatures []SignatureTuple

	rows, err := disk.SkateAppDB.Query(
		`SELECT operator, signature FROM `+disk.SignedTaskSchema+` WHERE 
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

func fetchSignedTasks() []disk.SignedTask {
	storedTasks, _ := disk.RetrieveSignedTasks()
	return storedTasks
}
