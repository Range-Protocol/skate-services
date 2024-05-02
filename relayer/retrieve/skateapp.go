package retrieve

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	pb "skatechain.org/api/pb/relayer"
	bindingISkateAVS "skatechain.org/contracts/bindings/ISkateAVS"
	"skatechain.org/lib/crypto/ecdsa"
	"skatechain.org/lib/logging"
	"skatechain.org/lib/on-chain/backend"
	memDb "skatechain.org/relayer/db/skateapp/mem"

	_ "github.com/joho/godotenv/autoload"
)

var (
	relayerLogger     = logging.NewLoggerWithConsoleWriter()
	Verbose           = true
	taskCache         = memDb.NewCache(100 * 1024 * 1024) // 100MB
	operatorCache     = memDb.NewCache(2 * 1024 * 1024)   // 2MB
	holeskyBackend, _ = backend.NewBackend(os.Getenv("HOLESKY_RPC"))
)

type submissionServer struct {
	pb.UnimplementedSubmissionServer
}

func isOperator(address common.Address) (bool, error) {
	contract, err := bindingISkateAVS.NewBindingISkateAVS(address, holeskyBackend)
	if err != nil {
		return false, errors.Wrap(err, "isOperator")
	}

	_, err = contract.Operators(&bind.CallOpts{})
	if err != nil {
		return false, errors.Wrap(err, "isOperator")
	}

	return true, nil
}

func (s *submissionServer) SubmitTask(ctx context.Context, in *pb.TaskSubmitRequest) (*pb.TaskSubmitReply, error) {
	relayerLogger.Info("Got request", "payload", in)

	// Step 1: Verify the operator

	// Step 2: Verify signature
	invalidReply := &pb.TaskSubmitReply{
		Result: pb.TaskStatus_REJECTED,
	}
	signature := [65]byte(in.Signature.Signature)
	valid, err := ecdsa.Verify(
		common.HexToAddress(in.Signature.Address),
		in.Task.Hash,
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
				"Message", in.Task.Msg,
				"TaskId", in.Task.TaskId,
				"ChainType", in.Task.ChainType,
				"ChainId", in.Task.ChainId,
			)
		}
		return invalidReply, nil
	}
	// replyResult = pb.TaskStatus_REJECTED

	// Step 2: Update the db and push to memcache
	msg := memDb.Message{
		TaskId:    in.Task.TaskId,
		ChainId:   in.Task.ChainId,
		ChainType: in.Task.ChainType,
		Message:   in.Task.Msg,
		Initiator: in.Task.Initiator,
	}
	cacheKey := memDb.GenKey(msg)
	taskCache.CacheMessage(cacheKey, msg)

	sig := memDb.Signature{
		Operator:  in.Signature.Address,
		Signature: signature,
	}
	taskCache.AppendSignature(cacheKey, sig)

	// Step 3: Check against EL quorum and Update status of db entry whenever threshold reached
	// bindingISkateAVS.NewBindingISkateAVS()

	return &pb.TaskSubmitReply{
		Result: pb.TaskStatus_ACCEPTED,
	}, nil
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

func NewSubmissionServer() *submissionServer {
	return &submissionServer{}
}
