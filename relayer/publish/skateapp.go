package publish

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	pb "skatechain.org/api/pb/relayer"
	bindingISkateAVS "skatechain.org/contracts/bindings/ISkateAVS"
	libcmd "skatechain.org/lib/cmd"
	"skatechain.org/lib/crypto/ecdsa"
	"skatechain.org/lib/on-chain/avs"
	"skatechain.org/lib/on-chain/backend"
	"skatechain.org/relayer/db/skateapp/disk"
)

func PublishTaskToAVSAndGateway(ctx context.Context) {
	ticker := time.NewTicker(12 * time.Second)
	defer ticker.Stop()

	relayerLogger.Info("Start AVS publisher process...")
	config := ctx.Value("config").(*libcmd.EnvironmentConfig)
	be, err := backend.NewBackend(config.SkateHttpRPC)
	if err != nil {
		relayerLogger.Fatal("AVS rpc error", "rpcUrl", config.SkateHttpRPC)
		return
	}
	avsContract, err := bindingISkateAVS.NewBindingISkateAVS(
		common.HexToAddress(config.SkateAVS), be,
	)
	if err != nil {
		relayerLogger.Fatal("Invalid avs contract", "address", config.SkateAVS)
		return
	}

	signer := ctx.Value("signer").(*libcmd.SignerConfig)
	privateKey, _ := backend.PrivateKeyFromKeystore(common.HexToAddress(signer.Address), signer.Passphrase)

	// Call submitTasks immediately
	submitTasksToAvs(avsContract, &be, config, privateKey)

	for {
		select {
		case <-ctx.Done():
			relayerLogger.Info("AVS publish process stopped!")
			return
		case <-ticker.C:
			submitTasksToAvs(avsContract, &be, config, privateKey)
		}
	}
}

func submitTasksToAvs(avsContract *bindingISkateAVS.BindingISkateAVS, be *backend.Backend, config *libcmd.EnvironmentConfig, privateKey *ecdsa.PrivateKey) {
	tasks := fetchPendingTasks()
	batchTaskId := make([]*big.Int, 0)
	batchMessageData := make([][]byte, 0)
	batchSignatures := make([][]bindingISkateAVS.ISkateAVSSignatureTuple, 0)
	for _, task := range tasks {
		signatures, _ := fetchSignaturesForTask(task)
		operators, _ := avsContract.Operators(&bind.CallOpts{})
		operatorCounts := len(operators)

		if len(signatures)*10_000 > operatorCounts*6_666 {
			if Verbose {
				relayerLogger.Info("Task approved for submission", "task", task)
			}
			taskId := new(big.Int).SetUint64(uint64(task.TaskId))
			messageData := avs.TaskData(task.Message, task.Initiator, pb.ChainType_EVM, task.ChainId)
			batchTaskId = append(batchTaskId, taskId)
			batchMessageData = append(batchMessageData, messageData)
			batchSignatures = append(batchSignatures, signatures)
		}
	}

	if len(batchTaskId) == 0 {
		if Verbose {
			relayerLogger.Info("No task to submit to avs")
		}
		return
	}

	chainId := new(big.Int).SetUint64(config.SkateChainId)
	txOptsWithSigner, _ := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	tx, err := avsContract.BatchSubmitData(
		txOptsWithSigner,
		batchTaskId,
		batchMessageData,
		batchSignatures,
	)
	if err != nil {
		relayerLogger.Error("Failed to submit transaction", "error", errors.Wrap(err, "SkateAVS.BatchSubmitData"))
		return
	}
	relayerLogger.Info("Transaction sent", "txHash", tx.Hash().Hex())

	receipt, err := be.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		relayerLogger.Error("Failed to get transaction receipt", "error", err)
		return
	}
	relayerLogger.Info("Transaction receipt", "status", receipt.Status)
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

func fetchPendingTasks() []disk.SignedTask {
	storedTasks, _ := disk.RetrieveSignedTasks()
	return storedTasks
}
