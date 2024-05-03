package publish

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	bindingISkateAVS "skatechain.org/contracts/bindings/ISkateAVS"
	"skatechain.org/lib/logging"
	"skatechain.org/lib/on-chain/avs"
	"skatechain.org/lib/on-chain/backend"
	avsMemcache "skatechain.org/relayer/db/avs/mem"
	"skatechain.org/relayer/db/skateapp/disk"
	skateappMemcache "skatechain.org/relayer/db/skateapp/mem"

	_ "github.com/joho/godotenv/autoload"
)

var (
	relayerLogger = logging.NewLoggerWithConsoleWriter()
	Verbose       = true
	taskCache     = skateappMemcache.NewCache(100 * 1024 * 1024) // 100MB
	operatorCache = avsMemcache.NewCache(2 * 1024 * 1024)        // 2MB

	// NOTE: change with config files
	holeskyBackend, _     = backend.NewBackend("https://holesky.drpc.org")
	holeskyAvsContract, _ = bindingISkateAVS.NewBindingISkateAVS(
		common.HexToAddress("0x2a0D46ED3D9D13F6a9b5B0D3274675143c803071"), holeskyBackend,
	)
	holeskyStrategy = common.HexToAddress("0x2a0D46ED3D9D13F6a9b5B0D3274675143c803071")
)

func PublishTaskToAVS(ctx context.Context) {
	ticker := time.NewTicker(12 * time.Second)
	defer ticker.Stop()
	relayerLogger.Info("Start AVS publisher process...")

  storedTasks := make([]skateappMemcache.Message, 0)

	relayerAddress := ctx.Value("address").(common.Address)
	passphrase := ctx.Value("passphrase").(string)

	// step 1: populate cachedTask with non-completed tasks in the db
  disk.SkateApp_RetriveSignedTasks("SELECT * FROM ?", disk.SignedTaskSchema)

	for {
		select {
		case <-ctx.Done():
			relayerLogger.Info("AVS publish process stopped!")
			return
		case <-ticker.C:
			// TODO: revalidate cachedTask
			for _, task := range storedTasks {
				msgKey := skateappMemcache.GenKey(task)
				cachedSignatures, _ := taskCache.GetSignatures(msgKey)
				operatorCounts, _ := operatorCache.GetOperatorCount()

				// Reference from SkateAVS contract, should change to (op * 2 + 2) / 3
				if uint32(len(cachedSignatures))*10_000 > (*operatorCounts)*6_666 {
					taskId := new(big.Int).SetUint64(uint64(task.TaskId))
					messageData := avs.TaskData(task.Message, task.Initiator, task.ChainType, task.ChainId)
					signatures := make([]bindingISkateAVS.ISkateAVSSignatureTuple, len(cachedSignatures))
					for _, sig := range cachedSignatures {
						signatureTuple := bindingISkateAVS.ISkateAVSSignatureTuple{
							Operator:  common.HexToAddress(sig.Operator),
							Signature: sig.Signature[:],
						}
						signatures = append(signatures, signatureTuple)
					}

					// Create a transactor to submit on-chain
					chainId := new(big.Int).SetUint64(uint64(task.ChainId))
					txOptsWithSigner, _ := backend.TransactorFromKeystore(relayerAddress, passphrase, chainId)
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
