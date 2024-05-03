package monitor

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pbCommon "skatechain.org/api/pb/common"
	pb "skatechain.org/api/pb/relayer"

	bindingSkateApp "skatechain.org/contracts/bindings/SkateApp"
	libcmd "skatechain.org/lib/cmd"
	"skatechain.org/lib/crypto/ecdsa"
	"skatechain.org/lib/monitor"
	"skatechain.org/lib/on-chain/avs"
	"skatechain.org/lib/on-chain/backend"
	operatorDb "skatechain.org/operator/db"
)

// WARNING: Must run with wss rpc.
//
// Server subscription only available on websocket endpoints.
func SubscribeSkateApp(addr common.Address, be backend.Backend, ctx context.Context) error {
	contract, err := bindingSkateApp.NewBindingSkateApp(addr, be)
	if err != nil {
		monitor.Logger.Error("Contract binding error: ", "error", err)
		return err
	}

	latest, _ := be.BlockNumber(ctx)
	watchOpts := &bind.WatchOpts{
		Start:   &latest,
		Context: ctx,
	}

	sink := make(chan *bindingSkateApp.BindingSkateAppTaskCreated)
	watcher, err := contract.WatchTaskCreated(watchOpts, sink, nil)
	if err != nil {
		if monitor.Verbose {
			monitor.Logger.Error("Watcher initialization error: ", "error", err)
		}
		return err
	}

	signer := ctx.Value("signer").(*libcmd.SignerConfig)

	var privateKey *ecdsa.PrivateKey
	if signer == nil {
		monitor.Logger.Info("No private key provided, run with watch-only mode")
	} else {
		privateKey, _ = backend.PrivateKeyFromKeystore(common.HexToAddress(signer.Address), signer.Passphrase)
	}

	// Event handler
	go func() {
		for {
			select {
			case task, ok := <-sink:
				if !ok {
					monitor.Logger.Error("Sink error, go-eth related")
					return
				}
				if monitor.Verbose {
					monitor.Logger.Info("Received TaskCreated event:", "task", task)
				}
				err = PostProcessLog(privateKey, task)
				if err != nil && monitor.Verbose {
					monitor.Logger.Error("Postprocess log failed", "error", err)
				}
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

func PostProcessLog(privateKey *ecdsa.PrivateKey, bindingTask *bindingSkateApp.BindingSkateAppTaskCreated) error {
	err := dumpLog(bindingTask)
	if err != nil {
		return err
	}
	if privateKey != nil {
		err = signAndBroadcastLog(privateKey, bindingTask)
		if err != nil {
			return err
		}
	}
	return nil
}

func signAndBroadcastLog(privateKey *ecdsa.PrivateKey, bindingTask *bindingSkateApp.BindingSkateAppTaskCreated) error {
	// Step 1: sign the log
	digestHash := avs.TaskDigest(
		uint32(bindingTask.TaskId.Int64()), bindingTask.Message, bindingTask.Signer.Hex(),
		pb.ChainType_EVM, bindingTask.Chain,
	)
	signature, err := ecdsa.Sign(digestHash, privateKey)
	if err != nil {
		return err
	}

	// Step 2: broad cast log over grpc server
	conn, err := grpc.Dial(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		if monitor.Verbose {
			monitor.Logger.Fatal("Relayer server not found", "error", err)
		}
		return err
	}
	defer conn.Close()
	c := pb.NewSubmissionClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Create a new Task
	task := &pb.Task{
		TaskId:    uint32(bindingTask.TaskId.Uint64()),
		Msg:       bindingTask.Message,
		ChainId:   bindingTask.Chain,
		Initiator: bindingTask.Signer.Hex(),
	}

	opAddr := ecdsa.PubkeyToAddress(privateKey.PublicKey).Hex()
	// Create a new SignedMessage
	signedMessage := &pbCommon.OperatorSignature{
		Signature: signature[:],
		Address:   opAddr,
	}

	// Create a new TaskSubmitRequest
	request := &pb.TaskSubmitRequest{
		Task:      task,
		Signature: signedMessage,
	}

	r, err := c.SubmitTask(ctx, request)
	if err != nil {
		if monitor.Verbose {
			monitor.Logger.Fatal("Could not submit task", "error", err)
		}
		return err
	}

	if monitor.Verbose {
		monitor.Logger.Info("Response: %s", r.Result.String())
	}
	return nil
}

func dumpLog(bindingTask *bindingSkateApp.BindingSkateAppTaskCreated) error {
	return operatorDb.SkateApp_InsertTask(bindingTask)
}

// NOTE: work over both https and wss, polling every 2 seconds by default
// To be implemented (logic not up to date)
func PollSkateApp(addr common.Address, backend backend.Backend, ctx context.Context) error {
	contract, err := bindingSkateApp.NewBindingSkateApp(addr, backend)
	if err != nil {
		monitor.Logger.Error("Contract binding error: ", "error", err)
		return err
	}

	// NOTE: Polling interval = 12s
	ticker := time.NewTicker(12 * time.Second)
	defer ticker.Stop()

	// Retrieve the latest block number as the starting point for the filter
	latest, err := backend.BlockNumber(ctx)
	if err != nil {
		if monitor.Verbose {
			monitor.Logger.Error("Error retrieving latest block number: ", "error", err)
		}
		return err
	}

	// Filter options
	queryOpts := &bind.FilterOpts{
		Start:   latest,
		Context: ctx,
	}

	// Function to process events
	processEvents := func() error {
		it, err := contract.FilterTaskCreated(queryOpts, nil)
		if err != nil {
			if monitor.Verbose {
				monitor.Logger.Error("Error fetching events: ", "error", err)
			}
			return err
		}

		// Process all events since the last poll
		for it.Next() {
			event := it.Event
			if monitor.Verbose {
				monitor.Logger.Info("Retrieved TaskCreated event:", "task", event)
			}
		}

		// Update the starting block for the next query to be the block number of the last fetched event
		if it.Event != nil {
			queryOpts.Start = it.Event.Raw.BlockNumber
		}
		return nil
	}

	// Immediately process events once before starting the ticker
	if err := processEvents(); err != nil {
		return err
	}

	// Polling loop
	go func() {
		for {
			select {
			case <-ticker.C:
				if err := processEvents(); err != nil {
					continue
				}
			case <-ctx.Done():
				if monitor.Verbose {
					monitor.Logger.Info("Polling stopped due to context cancellation")
				}
				return
			}
		}
	}()
	// Keep the function alive until the context is cancelled
	<-ctx.Done()
	return nil
}
