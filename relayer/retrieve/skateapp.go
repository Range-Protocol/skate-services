package retrieve

import (
	"context"
	"log"
	"net"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	pb "skatechain.org/api/pb/relayer"
	bindingISkateAVS "skatechain.org/contracts/bindings/ISkateAVS"
	"skatechain.org/lib/crypto/ecdsa"
	"skatechain.org/lib/logging"
	"skatechain.org/lib/on-chain/avs"
	"skatechain.org/lib/on-chain/backend"
	avsMemcache "skatechain.org/relayer/db/avs/mem"
	skateappDb "skatechain.org/relayer/db/skateapp/disk"
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

type submissionServer struct {
	pb.UnimplementedSubmissionServer
}

func NewSubmissionServer() *submissionServer {
	return &submissionServer{}
}

func (s *submissionServer) Start() {
	grpc_server := grpc.NewServer()
	pb.RegisterSubmissionServer(grpc_server, &submissionServer{})
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	relayerLogger.Info("server listening at ", "addr", lis.Addr())
	if err := grpc_server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *submissionServer) SubmitTask(ctx context.Context, in *pb.TaskSubmitRequest) (*pb.TaskSubmitReply, error) {
	relayerLogger.Info("Got request", "payload", in)
	invalidReply := &pb.TaskSubmitReply{
		Result: pb.TaskStatus_REJECTED,
	}

	// Step 1: Verify the operator
	isValidOperator, err := isOperator(in.Signature.Address)
	if err != nil {
		if Verbose {
			relayerLogger.Error("Validator address format error", "error", err)
		}
		return invalidReply, err
	}
	if !isValidOperator {
		if Verbose {
			relayerLogger.Error("Not an operator", "error", err)
		}
		return invalidReply, err
	}

	// Step 2: Verify signature
	signature := [65]byte(in.Signature.Signature)
	taskDigest := avs.TaskDigest(in.Task.TaskId, in.Task.Msg, in.Task.Initiator, in.Task.ChainType, in.Task.ChainId)
	valid, err := ecdsa.Verify(
		common.HexToAddress(in.Signature.Address),
		taskDigest,
		signature,
	)
	if err != nil {
		if Verbose {
			relayerLogger.Error("Signature format error", "error", err)
		}
		return invalidReply, err
	}
	if !valid {
		if Verbose {
			relayerLogger.Error("Signature is invalid",
				"operator", in.Signature.Address,
				"signature", signature,
				"TaskId", in.Task.TaskId,
				"Message", in.Task.Msg,
				"Initiator", in.Task.Initiator,
				"ChainType", in.Task.ChainType,
				"ChainId", in.Task.ChainId,
			)
		}
		return invalidReply, nil
	}

	// Step 3: Update the db and push to memcache
	msg := skateappMemcache.Message{
		TaskId:    in.Task.TaskId,
		ChainId:   in.Task.ChainId,
		ChainType: in.Task.ChainType,
		Message:   in.Task.Msg,
		Initiator: in.Task.Initiator,
	}

	msgKey := skateappMemcache.GenKey(msg)
	taskCache.CacheMessage(msgKey, msg)
	sig := skateappMemcache.Signature{
		Operator:  in.Signature.Address,
		Signature: signature,
	}
	taskCache.AppendSignature(msgKey, sig)

	signedTask := skateappDb.SignedTask{
		TaskId:    in.Task.TaskId,
		Message:   in.Task.Msg,
		Initiator: in.Task.Initiator,
		ChainId:   in.Task.ChainId,
		ChainType: uint32(in.Task.ChainType),
		Hash:      string(in.Task.Hash),
		Operator:  in.Signature.Address,
		Signature: string(in.Signature.Signature),
	}
	err = skateappDb.SkateApp_InsertSignedTask(signedTask)
	if err != nil && Verbose {
		relayerLogger.Error("Insert signed task to db failed", "error", err)
	}

	return &pb.TaskSubmitReply{
		Result: pb.TaskStatus_PROCESSING,
	}, nil
}

// NOTE: Right now we control the operators list,
// therefore cache revalidation period is set to INF (no expiration).
// Might need to change in the future.
func isOperator(address string) (bool, error) {
	// step 1: look up cache
	cachedOperator, _ := operatorCache.GetOperator(address)
	if cachedOperator != nil {
		return true, nil
	}

	// step 2: populate cache with on-chain data
	operators, err := holeskyAvsContract.Operators(&bind.CallOpts{})
	if err != nil {
		return false, errors.Wrap(err, "isOperator.Operators")
	}

	operatorCache.CacheOperatorCount(uint32(len(operators)))
	for _, op := range operators {
		cacheOp := avsMemcache.Operator{
			Address: op.Addr.Hex(),
		}
		operatorCache.CacheOperator(cacheOp)
		if op.Addr.Hex() == address {
			return true, nil
		}
		// TODO: cache stake amounts as well
	}

	return false, nil
}
