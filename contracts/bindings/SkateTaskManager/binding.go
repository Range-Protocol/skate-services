// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindingSkateTaskManager

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// BN254G1Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G1Point struct {
	X *big.Int
	Y *big.Int
}

// BN254G2Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G2Point struct {
	X [2]*big.Int
	Y [2]*big.Int
}

// IBLSSignatureCheckerNonSignerStakesAndSignature is an auto generated low-level Go binding around an user-defined struct.
type IBLSSignatureCheckerNonSignerStakesAndSignature struct {
	NonSignerQuorumBitmapIndices []uint32
	NonSignerPubkeys             []BN254G1Point
	QuorumApks                   []BN254G1Point
	ApkG2                        BN254G2Point
	Sigma                        BN254G1Point
	QuorumApkIndices             []uint32
	TotalStakeIndices            []uint32
	NonSignerStakeIndices        [][]uint32
}

// IBLSSignatureCheckerQuorumStakeTotals is an auto generated low-level Go binding around an user-defined struct.
type IBLSSignatureCheckerQuorumStakeTotals struct {
	SignedStakeForQuorum []*big.Int
	TotalStakeForQuorum  []*big.Int
}

// ISkateTaskManagerTask is an auto generated low-level Go binding around an user-defined struct.
type ISkateTaskManagerTask struct {
	InputMessage              string
	TaskCreatedBlock          uint32
	QuorumNumbers             []byte
	QuorumThresholdPercentage uint32
}

// ISkateTaskManagerTaskResponse is an auto generated low-level Go binding around an user-defined struct.
type ISkateTaskManagerTaskResponse struct {
	TaskId       uint32
	RelayMessage string
}

// ISkateTaskManagerTaskResponseMetadata is an auto generated low-level Go binding around an user-defined struct.
type ISkateTaskManagerTaskResponseMetadata struct {
	TaskResponsedBlock uint32
	HashOfNonSigners   [32]byte
}

// OperatorStateRetrieverCheckSignaturesIndices is an auto generated low-level Go binding around an user-defined struct.
type OperatorStateRetrieverCheckSignaturesIndices struct {
	NonSignerQuorumBitmapIndices []uint32
	QuorumApkIndices             []uint32
	TotalStakeIndices            []uint32
	NonSignerStakeIndices        [][]uint32
}

// OperatorStateRetrieverOperator is an auto generated low-level Go binding around an user-defined struct.
type OperatorStateRetrieverOperator struct {
	Operator   common.Address
	OperatorId [32]byte
	Stake      *big.Int
}

// BindingSkateTaskManagerMetaData contains all meta data concerning the BindingSkateTaskManager contract.
var BindingSkateTaskManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"_taskResponseWindowBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"TASK_CHALLENGE_WINDOW_BLOCK\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TASK_RESPONSE_WINDOW_BLOCK\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"aggregator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allTaskHashes\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allTaskResponses\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"blsApkRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBLSApkRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkSignatures\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.QuorumStakeTotals\",\"components\":[{\"name\":\"signedStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"},{\"name\":\"totalStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"}]},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createNewTask\",\"inputs\":[{\"name\":\"message\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"generator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCheckSignaturesIndices\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"nonSignerOperatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOperatorStateRetriever.CheckSignaturesIndices\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapsAtBlockNumber\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTaskResponseWindowBlock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_aggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_generator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"latestTaskNum\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"raiseAndResolveChallenge\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structISkateTaskManager.Task\",\"components\":[{\"name\":\"inputMessage\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structISkateTaskManager.TaskResponse\",\"components\":[{\"name\":\"taskId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"relayMessage\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"taskResponseMetadata\",\"type\":\"tuple\",\"internalType\":\"structISkateTaskManager.TaskResponseMetadata\",\"components\":[{\"name\":\"taskResponsedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"pubkeysOfNonSigningOperators\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registryCoordinator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"respondToTask\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structISkateTaskManager.Task\",\"components\":[{\"name\":\"inputMessage\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structISkateTaskManager.TaskResponse\",\"components\":[{\"name\":\"taskId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"relayMessage\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"nonSignerStakesAndSignature\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPauserRegistry\",\"inputs\":[{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setStaleStakesForbidden\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakeRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"staleStakesForbidden\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"taskNumber\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"taskSuccesfullyChallenged\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"trySignatureAndApkVerification\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"apk\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"pairingSuccessful\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"siganatureIsValid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewTaskCreated\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"task\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structISkateTaskManager.Task\",\"components\":[{\"name\":\"inputMessage\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauserRegistrySet\",\"inputs\":[{\"name\":\"pauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaleStakesForbiddenUpdate\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskChallengedSuccessfully\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"challenger\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskChallengedUnsuccessfully\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"challenger\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskCompleted\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskResponded\",\"inputs\":[{\"name\":\"taskResponse\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structISkateTaskManager.TaskResponse\",\"components\":[{\"name\":\"taskId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"relayMessage\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"taskResponseMetadata\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structISkateTaskManager.TaskResponseMetadata\",\"components\":[{\"name\":\"taskResponsedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x6101206040523480156200001257600080fd5b5060405162005e8e38038062005e8e8339810160408190526200003591620001f7565b81806001600160a01b03166080816001600160a01b031681525050806001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200008f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620000b591906200023e565b6001600160a01b031660a0816001600160a01b031681525050806001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200010d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200013391906200023e565b6001600160a01b031660c0816001600160a01b03168152505060a0516001600160a01b031663df5cf7236040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200018d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001b391906200023e565b6001600160a01b031660e052506097805460ff1916600117905563ffffffff16610100525062000265565b6001600160a01b0381168114620001f457600080fd5b50565b600080604083850312156200020b57600080fd5b82516200021881620001de565b602084015190925063ffffffff811681146200023357600080fd5b809150509250929050565b6000602082840312156200025157600080fd5b81516200025e81620001de565b9392505050565b60805160a05160c05160e05161010051615b9e620002f06000396000818161027d0152818161059c015261322b01526000818161056501526124e701526000818161041e01526126c90152600081816104450152818161289f0152612a6101526000818161046c01528181610e0b015281816121d10152818161234a01526125840152615b9e6000f3fe608060405234801561001057600080fd5b506004361061021c5760003560e01c80636d14a98711610125578063b98d0908116100ad578063f2fde38b1161007c578063f2fde38b14610587578063f5c9899d1461059a578063f63c5bab146105c0578063f8c8765e146105c8578063fabc1cbc146105db57600080fd5b8063b98d09081461051f578063cefdc1d41461052c578063d07241f01461054d578063df5cf7231461056057600080fd5b80637afa1eed116100f45780637afa1eed146104c5578063886f1195146104d85780638b00ce7c146104eb5780638da5cb5b146104fb578063b00b6c0a1461050c57600080fd5b80636d14a987146104675780636efb46361461048e578063715018a6146104af57806372d18e8d146104b757600080fd5b80634f739f74116101a85780635c975abb116101775780635c975abb146103db5780635d565c29146103e35780635decc3f5146103f65780635df4594614610419578063683048351461044057600080fd5b80634f739f7414610360578063595c6a67146103805780635ac86ab7146103885780635c155662146103bb57600080fd5b8063245a7bfc116101ef578063245a7bfc146102b45780632cb223d5146102df5780632d89f6fc1461030d5780633563b0d11461032d578063416c7e5e1461034d57600080fd5b806310d67a2f14610221578063136439dd14610236578063171f1d5b146102495780631ad4318914610278575b600080fd5b61023461022f3660046145b7565b6105ee565b005b6102346102443660046145d4565b6106aa565b61025c610257366004614752565b6107e9565b6040805192151583529015156020830152015b60405180910390f35b61029f7f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff909116815260200161026f565b60cd546102c7906001600160a01b031681565b6040516001600160a01b03909116815260200161026f565b6102ff6102ed3660046147c0565b60cb6020526000908152604090205481565b60405190815260200161026f565b6102ff61031b3660046147c0565b60ca6020526000908152604090205481565b61034061033b366004614834565b610973565b60405161026f919061494c565b61023461035b366004614974565b610e09565b61037361036e3660046149d9565b610f7e565b60405161026f9190614add565b6102346116a4565b6103ab610396366004614ba7565b606654600160ff9092169190911b9081161490565b604051901515815260200161026f565b6103ce6103c9366004614be7565b61176b565b60405161026f9190614ca1565b6066546102ff565b6102346103f1366004614d7e565b611933565b6103ab6104043660046147c0565b60cc6020526000908152604090205460ff1681565b6102c77f000000000000000000000000000000000000000000000000000000000000000081565b6102c77f000000000000000000000000000000000000000000000000000000000000000081565b6102c77f000000000000000000000000000000000000000000000000000000000000000081565b6104a161049c366004615024565b611e1e565b60405161026f9291906150e4565b610234612d16565b60c95463ffffffff1661029f565b60ce546102c7906001600160a01b031681565b6065546102c7906001600160a01b031681565b60c95461029f9063ffffffff1681565b6033546001600160a01b03166102c7565b61023461051a36600461512d565b612d2a565b6097546103ab9060ff1681565b61053f61053a3660046151be565b612ea8565b60405161026f9291906151f5565b61023461055b366004615216565b61303a565b6102c77f000000000000000000000000000000000000000000000000000000000000000081565b6102346105953660046145b7565b6134b9565b7f000000000000000000000000000000000000000000000000000000000000000061029f565b61029f606481565b6102346105d636600461529d565b61352f565b6102346105e93660046145d4565b613680565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610641573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061066591906152f9565b6001600160a01b0316336001600160a01b03161461069e5760405162461bcd60e51b815260040161069590615316565b60405180910390fd5b6106a7816137dc565b50565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa1580156106f2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107169190615360565b6107325760405162461bcd60e51b81526004016106959061537d565b606654818116146107ab5760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e70617573653a20696e76616c696420617474656d70742060448201527f746f20756e70617573652066756e6374696f6e616c69747900000000000000006064820152608401610695565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d906020015b60405180910390a250565b60008060007f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000187876000015188602001518860000151600060028110610831576108316153c5565b60200201518951600160200201518a60200151600060028110610856576108566153c5565b60200201518b60200151600160028110610872576108726153c5565b602090810291909101518c518d8301516040516108cf9a99989796959401988952602089019790975260408801959095526060870193909352608086019190915260a085015260c084015260e08301526101008201526101200190565b6040516020818303038152906040528051906020012060001c6108f291906153db565b905061096561090b61090488846138d3565b869061396a565b6109136139fe565b61095b61094c85610946604080518082018252600080825260209182015281518083019092526001825260029082015290565b906138d3565b6109558c613abe565b9061396a565b886201d4c0613b4e565b909890975095505050505050565b60606000846001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa1580156109b5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109d991906152f9565b90506000856001600160a01b0316639e9923c26040518163ffffffff1660e01b8152600401602060405180830381865afa158015610a1b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a3f91906152f9565b90506000866001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa158015610a81573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610aa591906152f9565b9050600086516001600160401b03811115610ac257610ac26145ed565b604051908082528060200260200182016040528015610af557816020015b6060815260200190600190039081610ae05790505b50905060005b8751811015610dfd576000888281518110610b1857610b186153c5565b0160200151604051638902624560e01b815260f89190911c6004820181905263ffffffff8a16602483015291506000906001600160a01b03871690638902624590604401600060405180830381865afa158015610b79573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610ba191908101906153fd565b905080516001600160401b03811115610bbc57610bbc6145ed565b604051908082528060200260200182016040528015610c0757816020015b6040805160608101825260008082526020808301829052928201528252600019909201910181610bda5790505b50848481518110610c1a57610c1a6153c5565b602002602001018190525060005b8151811015610de7576040518060600160405280876001600160a01b03166347b314e8858581518110610c5d57610c5d6153c5565b60200260200101516040518263ffffffff1660e01b8152600401610c8391815260200190565b602060405180830381865afa158015610ca0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cc491906152f9565b6001600160a01b03168152602001838381518110610ce457610ce46153c5565b60200260200101518152602001896001600160a01b031663fa28c627858581518110610d1257610d126153c5565b60209081029190910101516040516001600160e01b031960e084901b168152600481019190915260ff8816602482015263ffffffff8f166044820152606401602060405180830381865afa158015610d6e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d92919061548d565b6001600160601b0316815250858581518110610db057610db06153c5565b60200260200101518281518110610dc957610dc96153c5565b60200260200101819052508080610ddf906154cc565b915050610c28565b5050508080610df5906154cc565b915050610afb565b50979650505050505050565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610e67573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e8b91906152f9565b6001600160a01b0316336001600160a01b031614610f375760405162461bcd60e51b815260206004820152605c60248201527f424c535369676e6174757265436865636b65722e6f6e6c79436f6f7264696e6160448201527f746f724f776e65723a2063616c6c6572206973206e6f7420746865206f776e6560648201527f72206f6620746865207265676973747279436f6f7264696e61746f7200000000608482015260a401610695565b6097805460ff19168215159081179091556040519081527f40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc9060200160405180910390a150565b610fa96040518060800160405280606081526020016060815260200160608152602001606081525090565b6000876001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa158015610fe9573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061100d91906152f9565b905061103a6040518060800160405280606081526020016060815260200160608152602001606081525090565b6040516361c8a12f60e11b81526001600160a01b038a169063c391425e9061106a908b90899089906004016154e7565b600060405180830381865afa158015611087573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526110af9190810190615531565b81526040516340e03a8160e11b81526001600160a01b038316906381c07502906110e1908b908b908b906004016155e8565b600060405180830381865afa1580156110fe573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526111269190810190615531565b6040820152856001600160401b03811115611143576111436145ed565b60405190808252806020026020018201604052801561117657816020015b60608152602001906001900390816111615790505b50606082015260005b60ff81168711156115b5576000856001600160401b038111156111a4576111a46145ed565b6040519080825280602002602001820160405280156111cd578160200160208202803683370190505b5083606001518360ff16815181106111e7576111e76153c5565b602002602001018190525060005b868110156114b55760008c6001600160a01b03166304ec63518a8a85818110611220576112206153c5565b905060200201358e8860000151868151811061123e5761123e6153c5565b60200260200101516040518463ffffffff1660e01b815260040161127b9392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa158015611298573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112bc9190615611565b90506001600160c01b0381166113605760405162461bcd60e51b815260206004820152605c60248201527f4f70657261746f7253746174655265747269657665722e676574436865636b5360448201527f69676e617475726573496e64696365733a206f70657261746f72206d7573742060648201527f6265207265676973746572656420617420626c6f636b6e756d62657200000000608482015260a401610695565b8a8a8560ff16818110611375576113756153c5565b6001600160c01b03841692013560f81c9190911c6001908116141590506114a257856001600160a01b031663dd9846b98a8a858181106113b7576113b76153c5565b905060200201358d8d8860ff168181106113d3576113d36153c5565b6040516001600160e01b031960e087901b1681526004810194909452919091013560f81c60248301525063ffffffff8f166044820152606401602060405180830381865afa158015611429573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061144d919061563a565b85606001518560ff1681518110611466576114666153c5565b6020026020010151848151811061147f5761147f6153c5565b63ffffffff909216602092830291909101909101528261149e816154cc565b9350505b50806114ad816154cc565b9150506111f5565b506000816001600160401b038111156114d0576114d06145ed565b6040519080825280602002602001820160405280156114f9578160200160208202803683370190505b50905060005b8281101561157a5784606001518460ff1681518110611520576115206153c5565b60200260200101518181518110611539576115396153c5565b6020026020010151828281518110611553576115536153c5565b63ffffffff9092166020928302919091019091015280611572816154cc565b9150506114ff565b508084606001518460ff1681518110611595576115956153c5565b6020026020010181905250505080806115ad90615657565b91505061117f565b506000896001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa1580156115f6573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061161a91906152f9565b60405163354952a360e21b81529091506001600160a01b0382169063d5254a8c9061164d908b908b908e90600401615677565b600060405180830381865afa15801561166a573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526116929190810190615531565b60208301525098975050505050505050565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa1580156116ec573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117109190615360565b61172c5760405162461bcd60e51b81526004016106959061537d565b600019606681905560405190815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2565b60606000846001600160a01b031663c391425e84866040518363ffffffff1660e01b815260040161179d9291906156a1565b600060405180830381865afa1580156117ba573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526117e29190810190615531565b9050600084516001600160401b038111156117ff576117ff6145ed565b604051908082528060200260200182016040528015611828578160200160208202803683370190505b50905060005b855181101561192957866001600160a01b03166304ec6351878381518110611858576118586153c5565b602002602001015187868581518110611873576118736153c5565b60200260200101516040518463ffffffff1660e01b81526004016118b09392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa1580156118cd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118f19190615611565b6001600160c01b031682828151811061190c5761190c6153c5565b602090810291909101015280611921816154cc565b91505061182e565b5095945050505050565b600061194260208501856147c0565b63ffffffff8116600090815260cb60205260409020549091506119b15760405162461bcd60e51b815260206004820152602160248201527f5461736b206861736e2774206265656e20726573706f6e64656420746f2079656044820152601d60fa1b6064820152608401610695565b83836040516020016119c4929190615771565b60408051601f19818403018152918152815160209283012063ffffffff8416600090815260cb90935291205414611a635760405162461bcd60e51b815260206004820152603d60248201527f5461736b20726573706f6e736520646f6573206e6f74206d617463682074686560448201527f206f6e65207265636f7264656420696e2074686520636f6e74726163740000006064820152608401610695565b63ffffffff8116600090815260cc602052604090205460ff1615611afb5760405162461bcd60e51b815260206004820152604360248201527f54686520726573706f6e736520746f2074686973207461736b2068617320616c60448201527f7265616479206265656e206368616c6c656e676564207375636365737366756c606482015262363c9760e91b608482015260a401610695565b6064611b0a60208501856147c0565b611b1491906157af565b63ffffffff164363ffffffff161115611b955760405162461bcd60e51b815260206004820152603760248201527f546865206368616c6c656e676520706572696f6420666f72207468697320746160448201527f736b2068617320616c726561647920657870697265642e0000000000000000006064820152608401610695565b6000611ba460208601866157d7565b604051602001611bb592919061581d565b60408051601f198184030181529190528051602090910120611bd787806157d7565b604051602001611be892919061581d565b60408051601f1981840301815291905280516020909101201490506001811415611c4657604051339063ffffffff8416907ffd3e26beeb5967fc5a57a0446914eabc45b4aa474c67a51b4b5160cac60ddb0590600090a35050611e18565b600083516001600160401b03811115611c6157611c616145ed565b604051908082528060200260200182016040528015611c8a578160200160208202803683370190505b50905060005b8451811015611cfc57611ccd858281518110611cae57611cae6153c5565b6020026020010151805160009081526020918201519091526040902090565b828281518110611cdf57611cdf6153c5565b602090810291909101015280611cf4816154cc565b915050611c90565b506000611d0f6040890160208a016147c0565b82604051602001611d21929190615831565b60405160208183030381529060405280519060200120905085602001358114611dcb5760405162461bcd60e51b815260206004820152605060248201527f546865207075626b657973206f66206e6f6e2d7369676e696e67206f7065726160448201527f746f727320737570706c69656420627920746865206368616c6c656e6765722060648201526f30b932903737ba1031b7b93932b1ba1760811b608482015260a401610695565b63ffffffff8416600081815260cc6020526040808220805460ff19166001179055513392917fc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec91a3505050505b50505050565b6040805180820190915260608082526020820152600084611e955760405162461bcd60e51b81526020600482015260376024820152600080516020615b4983398151915260448201527f7265733a20656d7074792071756f72756d20696e7075740000000000000000006064820152608401610695565b60408301515185148015611ead575060a08301515185145b8015611ebd575060c08301515185145b8015611ecd575060e08301515185145b611f375760405162461bcd60e51b81526020600482015260416024820152600080516020615b4983398151915260448201527f7265733a20696e7075742071756f72756d206c656e677468206d69736d6174636064820152600d60fb1b608482015260a401610695565b82515160208401515114611faf5760405162461bcd60e51b815260206004820152604460248201819052600080516020615b49833981519152908201527f7265733a20696e707574206e6f6e7369676e6572206c656e677468206d69736d6064820152630c2e8c6d60e31b608482015260a401610695565b4363ffffffff168463ffffffff161061201e5760405162461bcd60e51b815260206004820152603c6024820152600080516020615b4983398151915260448201527f7265733a20696e76616c6964207265666572656e636520626c6f636b000000006064820152608401610695565b6040805180820182526000808252602080830191909152825180840190935260608084529083015290866001600160401b0381111561205f5761205f6145ed565b604051908082528060200260200182016040528015612088578160200160208202803683370190505b506020820152866001600160401b038111156120a6576120a66145ed565b6040519080825280602002602001820160405280156120cf578160200160208202803683370190505b50815260408051808201909152606080825260208201528560200151516001600160401b03811115612103576121036145ed565b60405190808252806020026020018201604052801561212c578160200160208202803683370190505b5081526020860151516001600160401b0381111561214c5761214c6145ed565b604051908082528060200260200182016040528015612175578160200160208202803683370190505b50816020018190525060006122478a8a8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505060408051639aa1653d60e01b815290516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169350639aa1653d925060048083019260209291908290030181865afa15801561221e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906122429190615879565b613d72565b905060005b8760200151518110156124c35761227288602001518281518110611cae57611cae6153c5565b83602001518281518110612288576122886153c5565b602090810291909101015280156123485760208301516122a9600183615896565b815181106122b9576122b96153c5565b602002602001015160001c836020015182815181106122da576122da6153c5565b602002602001015160001c11612348576040805162461bcd60e51b8152602060048201526024810191909152600080516020615b4983398151915260448201527f7265733a206e6f6e5369676e65725075626b657973206e6f7420736f727465646064820152608401610695565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166304ec63518460200151838151811061238d5761238d6153c5565b60200260200101518b8b6000015185815181106123ac576123ac6153c5565b60200260200101516040518463ffffffff1660e01b81526004016123e99392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa158015612406573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061242a9190615611565b6001600160c01b031683600001518281518110612449576124496153c5565b6020026020010181815250506124af6109046124838486600001518581518110612475576124756153c5565b602002602001015116613e05565b8a602001518481518110612499576124996153c5565b6020026020010151613e3090919063ffffffff16565b9450806124bb816154cc565b91505061224c565b50506124ce83613f14565b60975490935060ff166000816124e5576000612567565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c448feb86040518163ffffffff1660e01b8152600401602060405180830381865afa158015612543573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061256791906158ad565b905060005b8a811015612be55782156126c7578963ffffffff16827f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663249a0c428f8f868181106125c3576125c36153c5565b60405160e085901b6001600160e01b031916815292013560f81c600483015250602401602060405180830381865afa158015612603573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061262791906158ad565b61263191906158c6565b116126c75760405162461bcd60e51b81526020600482015260666024820152600080516020615b4983398151915260448201527f7265733a205374616b6552656769737472792075706461746573206d7573742060648201527f62652077697468696e207769746864726177616c44656c6179426c6f636b732060848201526577696e646f7760d01b60a482015260c401610695565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166368bccaac8d8d84818110612708576127086153c5565b9050013560f81c60f81b60f81c8c8c60a00151858151811061272c5761272c6153c5565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa158015612788573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906127ac91906158de565b6001600160401b0319166127cf8a604001518381518110611cae57611cae6153c5565b67ffffffffffffffff19161461286b5760405162461bcd60e51b81526020600482015260616024820152600080516020615b4983398151915260448201527f7265733a2071756f72756d41706b206861736820696e2073746f72616765206460648201527f6f6573206e6f74206d617463682070726f76696465642071756f72756d2061706084820152606b60f81b60a482015260c401610695565b61289b89604001518281518110612884576128846153c5565b60200260200101518761396a90919063ffffffff16565b95507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c8294c568d8d848181106128de576128de6153c5565b9050013560f81c60f81b60f81c8c8c60c001518581518110612902576129026153c5565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa15801561295e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612982919061548d565b85602001518281518110612998576129986153c5565b6001600160601b039092166020928302919091018201528501518051829081106129c4576129c46153c5565b6020026020010151856000015182815181106129e2576129e26153c5565b60200260200101906001600160601b031690816001600160601b0316815250506000805b8a6020015151811015612bd057612a5a86600001518281518110612a2c57612a2c6153c5565b60200260200101518f8f86818110612a4657612a466153c5565b600192013560f81c9290921c811614919050565b15612bbe577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663f2be94ae8f8f86818110612aa057612aa06153c5565b9050013560f81c60f81b60f81c8e89602001518581518110612ac457612ac46153c5565b60200260200101518f60e001518881518110612ae257612ae26153c5565b60200260200101518781518110612afb57612afb6153c5565b60209081029190910101516040516001600160e01b031960e087901b16815260ff909416600485015263ffffffff92831660248501526044840191909152166064820152608401602060405180830381865afa158015612b5f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612b83919061548d565b8751805185908110612b9757612b976153c5565b60200260200101818151612bab9190615909565b6001600160601b03169052506001909101905b80612bc8816154cc565b915050612a06565b50508080612bdd906154cc565b91505061256c565b505050600080612bff8c868a606001518b608001516107e9565b9150915081612c705760405162461bcd60e51b81526020600482015260436024820152600080516020615b4983398151915260448201527f7265733a2070616972696e6720707265636f6d70696c652063616c6c206661696064820152621b195960ea1b608482015260a401610695565b80612cd15760405162461bcd60e51b81526020600482015260396024820152600080516020615b4983398151915260448201527f7265733a207369676e617475726520697320696e76616c6964000000000000006064820152608401610695565b50506000878260200151604051602001612cec929190615831565b60408051808303601f190181529190528051602090910120929b929a509198505050505050505050565b612d1e613faf565b612d286000614009565b565b60ce546001600160a01b03163314612d8e5760405162461bcd60e51b815260206004820152602160248201527f5461736b2067656e657261746f72206d757374206265207468652063616c6c656044820152603960f91b6064820152608401610695565b60408051608081018252606081830181905286825263ffffffff438116602080850191909152908716918301919091528251601f8501829004820281018201909352838352909190849084908190840183828082843760009201919091525050505060408083019190915251612e0890829060200161597e565b60408051601f19818403018152828252805160209182012060c9805463ffffffff908116600090815260ca90945293909220555416907fb456431ed13b80d91ab280ac82018e48afde7e840d354010364822d75691447990612e6b90849061597e565b60405180910390a260c954612e879063ffffffff1660016157af565b60c9805463ffffffff191663ffffffff929092169190911790555050505050565b6040805160018082528183019092526000916060918391602080830190803683370190505090508481600081518110612ee357612ee36153c5565b60209081029190910101526040516361c8a12f60e11b81526000906001600160a01b0388169063c391425e90612f1f90889086906004016156a1565b600060405180830381865afa158015612f3c573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052612f649190810190615531565b600081518110612f7657612f766153c5565b60209081029190910101516040516304ec635160e01b81526004810188905263ffffffff87811660248301529091166044820181905291506000906001600160a01b038916906304ec635190606401602060405180830381865afa158015612fe2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906130069190615611565b6001600160c01b03169050600061301c8261405b565b90508161302a8a838a610973565b9550955050505050935093915050565b60cd546001600160a01b031633146130945760405162461bcd60e51b815260206004820152601d60248201527f41676772656761746f72206d757374206265207468652063616c6c65720000006044820152606401610695565b60006130a660408501602086016147c0565b90503660006130b860408701876157d7565b909250905060006130cf60808801606089016147c0565b905060ca60006130e260208901896147c0565b63ffffffff1663ffffffff168152602001908152602001600020548760405160200161310e91906159e5565b60405160208183030381529060405280519060200120146131975760405162461bcd60e51b815260206004820152603d60248201527f737570706c696564207461736b20646f6573206e6f74206d617463682074686560448201527f206f6e65207265636f7264656420696e2074686520636f6e74726163740000006064820152608401610695565b600060cb816131a960208a018a6147c0565b63ffffffff1663ffffffff16815260200190815260200160002054146132265760405162461bcd60e51b815260206004820152602c60248201527f41676772656761746f722068617320616c726561647920726573706f6e64656460448201526b20746f20746865207461736b60a01b6064820152608401610695565b6132507f0000000000000000000000000000000000000000000000000000000000000000856157af565b63ffffffff164363ffffffff1611156132c15760405162461bcd60e51b815260206004820152602d60248201527f41676772656761746f722068617320726573706f6e64656420746f207468652060448201526c7461736b20746f6f206c61746560981b6064820152608401610695565b6000866040516020016132d49190615a72565b6040516020818303038152906040528051906020012090506000806132fc8387878a8c611e1e565b9150915060005b858110156133fb578460ff1683602001518281518110613325576133256153c5565b60200260200101516133379190615a85565b6001600160601b0316606484600001518381518110613358576133586153c5565b60200260200101516001600160601b03166133739190615ab4565b10156133e9576040805162461bcd60e51b81526020600482015260248101919091527f5369676e61746f7269657320646f206e6f74206f776e206174206c656173742060448201527f7468726573686f6c642070657263656e74616765206f6620612071756f72756d6064820152608401610695565b806133f3816154cc565b915050613303565b5060408051808201825263ffffffff43168152602080820184905291519091613428918c91849101615ad3565b6040516020818303038152906040528051906020012060cb60008c600001602081019061345591906147c0565b63ffffffff1663ffffffff168152602001908152602001600020819055507fd0e833257cf2514f7f416921130672e0a5099b141e19ad5285b881ec34e9eb548a826040516134a4929190615ad3565b60405180910390a15050505050505050505050565b6134c1613faf565b6001600160a01b0381166135265760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610695565b6106a781614009565b600054610100900460ff161580801561354f5750600054600160ff909116105b806135695750303b158015613569575060005460ff166001145b6135cc5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610695565b6000805460ff1916600117905580156135ef576000805461ff0019166101001790555b6135fa856000614127565b61360384614009565b60cd80546001600160a01b038086166001600160a01b03199283161790925560ce8054928516929091169190911790558015613679576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050505050565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156136d3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906136f791906152f9565b6001600160a01b0316336001600160a01b0316146137275760405162461bcd60e51b815260040161069590615316565b6066541981196066541916146137a55760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e756e70617573653a20696e76616c696420617474656d7060448201527f7420746f2070617573652066756e6374696f6e616c69747900000000000000006064820152608401610695565b606681905560405181815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c906020016107de565b6001600160a01b03811661386a5760405162461bcd60e51b815260206004820152604960248201527f5061757361626c652e5f73657450617573657252656769737472793a206e657760448201527f50617573657252656769737472792063616e6e6f7420626520746865207a65726064820152686f206164647265737360b81b608482015260a401610695565b606554604080516001600160a01b03928316815291831660208301527f6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6910160405180910390a1606580546001600160a01b0319166001600160a01b0392909216919091179055565b60408051808201909152600080825260208201526138ef6144c8565b835181526020808501519082015260408082018490526000908360608460076107d05a03fa905080801561392257613924565bfe5b50806139625760405162461bcd60e51b815260206004820152600d60248201526c1958cb5b5d5b0b59985a5b1959609a1b6044820152606401610695565b505092915050565b60408051808201909152600080825260208201526139866144e6565b835181526020808501518183015283516040808401919091529084015160608301526000908360808460066107d05a03fa90508080156139225750806139625760405162461bcd60e51b815260206004820152600d60248201526c1958cb5859190b59985a5b1959609a1b6044820152606401610695565b613a06614504565b50604080516080810182527f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c28183019081527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed6060830152815281518083019092527f275dc4a288d1afb3cbb1ac09187524c7db36395df7be3b99e673b13a075a65ec82527f1d9befcd05a5323e6da4d435f3b617cdb3af83285c2df711ef39c01571827f9d60208381019190915281019190915290565b604080518082019091526000808252602082015260008080613aee600080516020615b29833981519152866153db565b90505b613afa81614211565b9093509150600080516020615b29833981519152828309831415613b34576040805180820190915290815260208101919091529392505050565b600080516020615b29833981519152600182089050613af1565b604080518082018252868152602080820186905282518084019093528683528201849052600091829190613b80614529565b60005b6002811015613d45576000613b99826006615ab4565b9050848260028110613bad57613bad6153c5565b60200201515183613bbf8360006158c6565b600c8110613bcf57613bcf6153c5565b6020020152848260028110613be657613be66153c5565b60200201516020015183826001613bfd91906158c6565b600c8110613c0d57613c0d6153c5565b6020020152838260028110613c2457613c246153c5565b6020020151515183613c378360026158c6565b600c8110613c4757613c476153c5565b6020020152838260028110613c5e57613c5e6153c5565b6020020151516001602002015183613c778360036158c6565b600c8110613c8757613c876153c5565b6020020152838260028110613c9e57613c9e6153c5565b602002015160200151600060028110613cb957613cb96153c5565b602002015183613cca8360046158c6565b600c8110613cda57613cda6153c5565b6020020152838260028110613cf157613cf16153c5565b602002015160200151600160028110613d0c57613d0c6153c5565b602002015183613d1d8360056158c6565b600c8110613d2d57613d2d6153c5565b60200201525080613d3d816154cc565b915050613b83565b50613d4e614548565b60006020826101808560088cfa9151919c9115159b50909950505050505050505050565b600080613d7e84614293565b9050808360ff166001901b11613dfc5760405162461bcd60e51b815260206004820152603f60248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206269746d61702065786365656473206d61782076616c7565006064820152608401610695565b90505b92915050565b6000805b8215613dff57613e1a600184615896565b9092169180613e2881615b06565b915050613e09565b60408051808201909152600080825260208201526102008261ffff1610613e8c5760405162461bcd60e51b815260206004820152601060248201526f7363616c61722d746f6f2d6c6172676560801b6044820152606401610695565b8161ffff1660011415613ea0575081613dff565b6040805180820190915260008082526020820181905284906001905b8161ffff168661ffff1610613f0957600161ffff871660ff83161c81161415613eec57613ee9848461396a565b93505b613ef6838461396a565b92506201fffe600192831b169101613ebc565b509195945050505050565b60408051808201909152600080825260208201528151158015613f3957506020820151155b15613f57575050604080518082019091526000808252602082015290565b604051806040016040528083600001518152602001600080516020615b298339815191528460200151613f8a91906153db565b613fa290600080516020615b29833981519152615896565b905292915050565b919050565b6033546001600160a01b03163314612d285760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610695565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b606060008061406984613e05565b61ffff166001600160401b03811115614084576140846145ed565b6040519080825280601f01601f1916602001820160405280156140ae576020820181803683370190505b5090506000805b8251821080156140c6575061010081105b1561411d576001811b93508584161561410d578060f81b8383815181106140ef576140ef6153c5565b60200101906001600160f81b031916908160001a9053508160010191505b614116816154cc565b90506140b5565b5090949350505050565b6065546001600160a01b031615801561414857506001600160a01b03821615155b6141ca5760405162461bcd60e51b815260206004820152604760248201527f5061757361626c652e5f696e697469616c697a655061757365723a205f696e6960448201527f7469616c697a6550617573657228292063616e206f6e6c792062652063616c6c6064820152666564206f6e636560c81b608482015260a401610695565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a261420d826137dc565b5050565b60008080600080516020615b298339815191526003600080516020615b2983398151915286600080516020615b29833981519152888909090890506000614287827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f52600080516020615b29833981519152614420565b91959194509092505050565b60006101008251111561431c5760405162461bcd60e51b8152602060048201526044602482018190527f4269746d61705574696c732e6f72646572656442797465734172726179546f42908201527f69746d61703a206f7264657265644279746573417272617920697320746f6f206064820152636c6f6e6760e01b608482015260a401610695565b815161432a57506000919050565b60008083600081518110614340576143406153c5565b0160200151600160f89190911c81901b92505b84518110156144175784818151811061436e5761436e6153c5565b0160200151600160f89190911c1b91508282116144035760405162461bcd60e51b815260206004820152604760248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206f72646572656442797465734172726179206973206e6f74206064820152661bdc99195c995960ca1b608482015260a401610695565b91811791614410816154cc565b9050614353565b50909392505050565b60008061442b614548565b614433614566565b602080825281810181905260408201819052606082018890526080820187905260a082018690528260c08360056107d05a03fa92508280156139225750826144bd5760405162461bcd60e51b815260206004820152601a60248201527f424e3235342e6578704d6f643a2063616c6c206661696c7572650000000000006044820152606401610695565b505195945050505050565b60405180606001604052806003906020820280368337509192915050565b60405180608001604052806004906020820280368337509192915050565b6040518060400160405280614517614584565b8152602001614524614584565b905290565b604051806101800160405280600c906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b6040518060c001604052806006906020820280368337509192915050565b60405180604001604052806002906020820280368337509192915050565b6001600160a01b03811681146106a757600080fd5b6000602082840312156145c957600080fd5b8135613dfc816145a2565b6000602082840312156145e657600080fd5b5035919050565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b0381118282101715614625576146256145ed565b60405290565b60405161010081016001600160401b0381118282101715614625576146256145ed565b604051601f8201601f191681016001600160401b0381118282101715614676576146766145ed565b604052919050565b60006040828403121561469057600080fd5b614698614603565b9050813581526020820135602082015292915050565b600082601f8301126146bf57600080fd5b604051604081018181106001600160401b03821117156146e1576146e16145ed565b80604052508060408401858111156146f857600080fd5b845b81811015613f095780358352602092830192016146fa565b60006080828403121561472457600080fd5b61472c614603565b905061473883836146ae565b815261474783604084016146ae565b602082015292915050565b600080600080610120858703121561476957600080fd5b8435935061477a866020870161467e565b92506147898660608701614712565b91506147988660e0870161467e565b905092959194509250565b63ffffffff811681146106a757600080fd5b8035613faa816147a3565b6000602082840312156147d257600080fd5b8135613dfc816147a3565b60006001600160401b038311156147f6576147f66145ed565b614809601f8401601f191660200161464e565b905082815283838301111561481d57600080fd5b828260208301376000602084830101529392505050565b60008060006060848603121561484957600080fd5b8335614854816145a2565b925060208401356001600160401b0381111561486f57600080fd5b8401601f8101861361488057600080fd5b61488f868235602084016147dd565b92505060408401356148a0816147a3565b809150509250925092565b6000815180845260208085019450848260051b86018286016000805b8681101561493e578484038a52825180518086529087019087860190845b8181101561492957835180516001600160a01b031684528a8101518b8501526040908101516001600160601b031690840152928901926060909201916001016148e5565b50509a87019a945050918501916001016148c7565b509198975050505050505050565b60208152600061495f60208301846148ab565b9392505050565b80151581146106a757600080fd5b60006020828403121561498657600080fd5b8135613dfc81614966565b60008083601f8401126149a357600080fd5b5081356001600160401b038111156149ba57600080fd5b6020830191508360208285010111156149d257600080fd5b9250929050565b600080600080600080608087890312156149f257600080fd5b86356149fd816145a2565b95506020870135614a0d816147a3565b945060408701356001600160401b0380821115614a2957600080fd5b614a358a838b01614991565b90965094506060890135915080821115614a4e57600080fd5b818901915089601f830112614a6257600080fd5b813581811115614a7157600080fd5b8a60208260051b8501011115614a8657600080fd5b6020830194508093505050509295509295509295565b600081518084526020808501945080840160005b83811015614ad257815163ffffffff1687529582019590820190600101614ab0565b509495945050505050565b600060208083528351608082850152614af960a0850182614a9c565b905081850151601f1980868403016040870152614b168383614a9c565b92506040870151915080868403016060870152614b338383614a9c565b60608801518782038301608089015280518083529194508501925084840190600581901b8501860160005b82811015614b8a5784878303018452614b78828751614a9c565b95880195938801939150600101614b5e565b509998505050505050505050565b60ff811681146106a757600080fd5b600060208284031215614bb957600080fd5b8135613dfc81614b98565b60006001600160401b03821115614bdd57614bdd6145ed565b5060051b60200190565b600080600060608486031215614bfc57600080fd5b8335614c07816145a2565b92506020848101356001600160401b03811115614c2357600080fd5b8501601f81018713614c3457600080fd5b8035614c47614c4282614bc4565b61464e565b81815260059190911b82018301908381019089831115614c6657600080fd5b928401925b82841015614c8457833582529284019290840190614c6b565b8096505050505050614c98604085016147b5565b90509250925092565b6020808252825182820181905260009190848201906040850190845b81811015614cd957835183529284019291840191600101614cbd565b50909695505050505050565b600060808284031215614cf757600080fd5b50919050565b600060408284031215614cf757600080fd5b600082601f830112614d2057600080fd5b81356020614d30614c4283614bc4565b82815260069290921b84018101918181019086841115614d4f57600080fd5b8286015b84811015614d7357614d65888261467e565b835291830191604001614d53565b509695505050505050565b60008060008060a08587031215614d9457600080fd5b84356001600160401b0380821115614dab57600080fd5b614db788838901614ce5565b95506020870135915080821115614dcd57600080fd5b614dd988838901614cfd565b9450614de88860408901614cfd565b93506080870135915080821115614dfe57600080fd5b50614e0b87828801614d0f565b91505092959194509250565b600082601f830112614e2857600080fd5b81356020614e38614c4283614bc4565b82815260059290921b84018101918181019086841115614e5757600080fd5b8286015b84811015614d73578035614e6e816147a3565b8352918301918301614e5b565b600082601f830112614e8c57600080fd5b81356020614e9c614c4283614bc4565b82815260059290921b84018101918181019086841115614ebb57600080fd5b8286015b84811015614d735780356001600160401b03811115614ede5760008081fd5b614eec8986838b0101614e17565b845250918301918301614ebf565b60006101808284031215614f0d57600080fd5b614f1561462b565b905081356001600160401b0380821115614f2e57600080fd5b614f3a85838601614e17565b83526020840135915080821115614f5057600080fd5b614f5c85838601614d0f565b60208401526040840135915080821115614f7557600080fd5b614f8185838601614d0f565b6040840152614f938560608601614712565b6060840152614fa58560e0860161467e565b6080840152610120840135915080821115614fbf57600080fd5b614fcb85838601614e17565b60a0840152610140840135915080821115614fe557600080fd5b614ff185838601614e17565b60c084015261016084013591508082111561500b57600080fd5b5061501884828501614e7b565b60e08301525092915050565b60008060008060006080868803121561503c57600080fd5b8535945060208601356001600160401b038082111561505a57600080fd5b61506689838a01614991565b90965094506040880135915061507b826147a3565b9092506060870135908082111561509157600080fd5b5061509e88828901614efa565b9150509295509295909350565b600081518084526020808501945080840160005b83811015614ad25781516001600160601b0316875295820195908201906001016150bf565b60408152600083516040808401526150ff60808401826150ab565b90506020850151603f1984830301606085015261511c82826150ab565b925050508260208301529392505050565b6000806000806060858703121561514357600080fd5b84356001600160401b038082111561515a57600080fd5b818701915087601f83011261516e57600080fd5b61517d888335602085016147dd565b95506020870135915061518f826147a3565b909350604086013590808211156151a557600080fd5b506151b287828801614991565b95989497509550505050565b6000806000606084860312156151d357600080fd5b83356151de816145a2565b92506020840135915060408401356148a0816147a3565b82815260406020820152600061520e60408301846148ab565b949350505050565b60008060006060848603121561522b57600080fd5b83356001600160401b038082111561524257600080fd5b61524e87838801614ce5565b9450602086013591508082111561526457600080fd5b61527087838801614cfd565b9350604086013591508082111561528657600080fd5b5061529386828701614efa565b9150509250925092565b600080600080608085870312156152b357600080fd5b84356152be816145a2565b935060208501356152ce816145a2565b925060408501356152de816145a2565b915060608501356152ee816145a2565b939692955090935050565b60006020828403121561530b57600080fd5b8151613dfc816145a2565b6020808252602a908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526939903ab73830bab9b2b960b11b606082015260800190565b60006020828403121561537257600080fd5b8151613dfc81614966565b60208082526028908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526739903830bab9b2b960c11b606082015260800190565b634e487b7160e01b600052603260045260246000fd5b6000826153f857634e487b7160e01b600052601260045260246000fd5b500690565b6000602080838503121561541057600080fd5b82516001600160401b0381111561542657600080fd5b8301601f8101851361543757600080fd5b8051615445614c4282614bc4565b81815260059190911b8201830190838101908783111561546457600080fd5b928401925b8284101561548257835182529284019290840190615469565b979650505050505050565b60006020828403121561549f57600080fd5b81516001600160601b0381168114613dfc57600080fd5b634e487b7160e01b600052601160045260246000fd5b60006000198214156154e0576154e06154b6565b5060010190565b63ffffffff84168152604060208201819052810182905260006001600160fb1b0383111561551457600080fd5b8260051b8085606085013760009201606001918252509392505050565b6000602080838503121561554457600080fd5b82516001600160401b0381111561555a57600080fd5b8301601f8101851361556b57600080fd5b8051615579614c4282614bc4565b81815260059190911b8201830190838101908783111561559857600080fd5b928401925b828410156154825783516155b0816147a3565b8252928401929084019061559d565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b63ffffffff841681526040602082015260006156086040830184866155bf565b95945050505050565b60006020828403121561562357600080fd5b81516001600160c01b0381168114613dfc57600080fd5b60006020828403121561564c57600080fd5b8151613dfc816147a3565b600060ff821660ff81141561566e5761566e6154b6565b60010192915050565b60408152600061568b6040830185876155bf565b905063ffffffff83166020830152949350505050565b60006040820163ffffffff851683526020604081850152818551808452606086019150828701935060005b818110156156e8578451835293830193918301916001016156cc565b5090979650505050505050565b6000808335601e1984360301811261570c57600080fd5b83016020810192503590506001600160401b0381111561572b57600080fd5b8036038313156149d257600080fd5b60008135615747816147a3565b63ffffffff16835261575c60208301836156f5565b604060208601526156086040860182846155bf565b606081526000615784606083018561573a565b90508235615791816147a3565b63ffffffff8116602084015250602083013560408301529392505050565b600063ffffffff8083168185168083038211156157ce576157ce6154b6565b01949350505050565b6000808335601e198436030181126157ee57600080fd5b8301803591506001600160401b0382111561580857600080fd5b6020019150368190038213156149d257600080fd5b60208152600061520e6020830184866155bf565b63ffffffff60e01b8360e01b1681526000600482018351602080860160005b8381101561586c57815185529382019390820190600101615850565b5092979650505050505050565b60006020828403121561588b57600080fd5b8151613dfc81614b98565b6000828210156158a8576158a86154b6565b500390565b6000602082840312156158bf57600080fd5b5051919050565b600082198211156158d9576158d96154b6565b500190565b6000602082840312156158f057600080fd5b815167ffffffffffffffff1981168114613dfc57600080fd5b60006001600160601b0383811690831681811015615929576159296154b6565b039392505050565b6000815180845260005b818110156159575760208185018101518683018201520161593b565b81811115615969576000602083870101525b50601f01601f19169290920160200192915050565b60208152600082516080602084015261599a60a0840182615931565b9050602084015163ffffffff808216604086015260408601519150601f198584030160608601526159cb8383615931565b925080606087015116608086015250508091505092915050565b6020815260006159f583846156f5565b60806020850152615a0a60a0850182846155bf565b9150506020840135615a1b816147a3565b63ffffffff8082166040860152615a3560408701876156f5565b868503601f190160608801529250615a4e8484836155bf565b93505060608601359150615a61826147a3565b166080939093019290925250919050565b60208152600061495f602083018461573a565b60006001600160601b0380831681851681830481118215151615615aab57615aab6154b6565b02949350505050565b6000816000190483118215151615615ace57615ace6154b6565b500290565b606081526000615ae6606083018561573a565b905063ffffffff8351166020830152602083015160408301529392505050565b600061ffff80831681811415615b1e57615b1e6154b6565b600101939250505056fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47424c535369676e6174757265436865636b65722e636865636b5369676e617475a2646970667358221220997fae8963b3c5757ee9416f6b91bc5e2f212d96bfb637babe373f4da75b0da364736f6c634300080c0033",
}

// BindingSkateTaskManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use BindingSkateTaskManagerMetaData.ABI instead.
var BindingSkateTaskManagerABI = BindingSkateTaskManagerMetaData.ABI

// BindingSkateTaskManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BindingSkateTaskManagerMetaData.Bin instead.
var BindingSkateTaskManagerBin = BindingSkateTaskManagerMetaData.Bin

// DeployBindingSkateTaskManager deploys a new Ethereum contract, binding an instance of BindingSkateTaskManager to it.
func DeployBindingSkateTaskManager(auth *bind.TransactOpts, backend bind.ContractBackend, _registryCoordinator common.Address, _taskResponseWindowBlock uint32) (common.Address, *types.Transaction, *BindingSkateTaskManager, error) {
	parsed, err := BindingSkateTaskManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BindingSkateTaskManagerBin), backend, _registryCoordinator, _taskResponseWindowBlock)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BindingSkateTaskManager{BindingSkateTaskManagerCaller: BindingSkateTaskManagerCaller{contract: contract}, BindingSkateTaskManagerTransactor: BindingSkateTaskManagerTransactor{contract: contract}, BindingSkateTaskManagerFilterer: BindingSkateTaskManagerFilterer{contract: contract}}, nil
}

// BindingSkateTaskManager is an auto generated Go binding around an Ethereum contract.
type BindingSkateTaskManager struct {
	BindingSkateTaskManagerCaller     // Read-only binding to the contract
	BindingSkateTaskManagerTransactor // Write-only binding to the contract
	BindingSkateTaskManagerFilterer   // Log filterer for contract events
}

// BindingSkateTaskManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type BindingSkateTaskManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingSkateTaskManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BindingSkateTaskManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingSkateTaskManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BindingSkateTaskManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingSkateTaskManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BindingSkateTaskManagerSession struct {
	Contract     *BindingSkateTaskManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// BindingSkateTaskManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BindingSkateTaskManagerCallerSession struct {
	Contract *BindingSkateTaskManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// BindingSkateTaskManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BindingSkateTaskManagerTransactorSession struct {
	Contract     *BindingSkateTaskManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// BindingSkateTaskManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type BindingSkateTaskManagerRaw struct {
	Contract *BindingSkateTaskManager // Generic contract binding to access the raw methods on
}

// BindingSkateTaskManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BindingSkateTaskManagerCallerRaw struct {
	Contract *BindingSkateTaskManagerCaller // Generic read-only contract binding to access the raw methods on
}

// BindingSkateTaskManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BindingSkateTaskManagerTransactorRaw struct {
	Contract *BindingSkateTaskManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBindingSkateTaskManager creates a new instance of BindingSkateTaskManager, bound to a specific deployed contract.
func NewBindingSkateTaskManager(address common.Address, backend bind.ContractBackend) (*BindingSkateTaskManager, error) {
	contract, err := bindBindingSkateTaskManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BindingSkateTaskManager{BindingSkateTaskManagerCaller: BindingSkateTaskManagerCaller{contract: contract}, BindingSkateTaskManagerTransactor: BindingSkateTaskManagerTransactor{contract: contract}, BindingSkateTaskManagerFilterer: BindingSkateTaskManagerFilterer{contract: contract}}, nil
}

// NewBindingSkateTaskManagerCaller creates a new read-only instance of BindingSkateTaskManager, bound to a specific deployed contract.
func NewBindingSkateTaskManagerCaller(address common.Address, caller bind.ContractCaller) (*BindingSkateTaskManagerCaller, error) {
	contract, err := bindBindingSkateTaskManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BindingSkateTaskManagerCaller{contract: contract}, nil
}

// NewBindingSkateTaskManagerTransactor creates a new write-only instance of BindingSkateTaskManager, bound to a specific deployed contract.
func NewBindingSkateTaskManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*BindingSkateTaskManagerTransactor, error) {
	contract, err := bindBindingSkateTaskManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BindingSkateTaskManagerTransactor{contract: contract}, nil
}

// NewBindingSkateTaskManagerFilterer creates a new log filterer instance of BindingSkateTaskManager, bound to a specific deployed contract.
func NewBindingSkateTaskManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*BindingSkateTaskManagerFilterer, error) {
	contract, err := bindBindingSkateTaskManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BindingSkateTaskManagerFilterer{contract: contract}, nil
}

// bindBindingSkateTaskManager binds a generic wrapper to an already deployed contract.
func bindBindingSkateTaskManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BindingSkateTaskManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BindingSkateTaskManager *BindingSkateTaskManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BindingSkateTaskManager.Contract.BindingSkateTaskManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BindingSkateTaskManager *BindingSkateTaskManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BindingSkateTaskManager.Contract.BindingSkateTaskManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BindingSkateTaskManager *BindingSkateTaskManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BindingSkateTaskManager.Contract.BindingSkateTaskManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BindingSkateTaskManager *BindingSkateTaskManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BindingSkateTaskManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BindingSkateTaskManager *BindingSkateTaskManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BindingSkateTaskManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BindingSkateTaskManager *BindingSkateTaskManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BindingSkateTaskManager.Contract.contract.Transact(opts, method, params...)
}

// TASKCHALLENGEWINDOWBLOCK is a free data retrieval call binding the contract method 0xf63c5bab.
//
// Solidity: function TASK_CHALLENGE_WINDOW_BLOCK() view returns(uint32)
func (_BindingSkateTaskManager *BindingSkateTaskManagerCaller) TASKCHALLENGEWINDOWBLOCK(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _BindingSkateTaskManager.contract.Call(opts, &out, "TASK_CHALLENGE_WINDOW_BLOCK")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TASKCHALLENGEWINDOWBLOCK is a free data retrieval call binding the contract method 0xf63c5bab.
//
// Solidity: function TASK_CHALLENGE_WINDOW_BLOCK() view returns(uint32)
func (_BindingSkateTaskManager *BindingSkateTaskManagerSession) TASKCHALLENGEWINDOWBLOCK() (uint32, error) {
	return _BindingSkateTaskManager.Contract.TASKCHALLENGEWINDOWBLOCK(&_BindingSkateTaskManager.CallOpts)
}

// TASKCHALLENGEWINDOWBLOCK is a free data retrieval call binding the contract method 0xf63c5bab.
//
// Solidity: function TASK_CHALLENGE_WINDOW_BLOCK() view returns(uint32)
func (_BindingSkateTaskManager *BindingSkateTaskManagerCallerSession) TASKCHALLENGEWINDOWBLOCK() (uint32, error) {
	return _BindingSkateTaskManager.Contract.TASKCHALLENGEWINDOWBLOCK(&_BindingSkateTaskManager.CallOpts)
}

// TASKRESPONSEWINDOWBLOCK is a free data retrieval call binding the contract method 0x1ad43189.
//
// Solidity: function TASK_RESPONSE_WINDOW_BLOCK() view returns(uint32)
func (_BindingSkateTaskManager *BindingSkateTaskManagerCaller) TASKRESPONSEWINDOWBLOCK(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _BindingSkateTaskManager.contract.Call(opts, &out, "TASK_RESPONSE_WINDOW_BLOCK")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TASKRESPONSEWINDOWBLOCK is a free data retrieval call binding the contract method 0x1ad43189.
//
// Solidity: function TASK_RESPONSE_WINDOW_BLOCK() view returns(uint32)
func (_BindingSkateTaskManager *BindingSkateTaskManagerSession) TASKRESPONSEWINDOWBLOCK() (uint32, error) {
	return _BindingSkateTaskManager.Contract.TASKRESPONSEWINDOWBLOCK(&_BindingSkateTaskManager.CallOpts)
}

// TASKRESPONSEWINDOWBLOCK is a free data retrieval call binding the contract method 0x1ad43189.
//
// Solidity: function TASK_RESPONSE_WINDOW_BLOCK() view returns(uint32)
func (_BindingSkateTaskManager *BindingSkateTaskManagerCallerSession) TASKRESPONSEWINDOWBLOCK() (uint32, error) {
	return _BindingSkateTaskManager.Contract.TASKRESPONSEWINDOWBLOCK(&_BindingSkateTaskManager.CallOpts)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_BindingSkateTaskManager *BindingSkateTaskManagerCaller) Aggregator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BindingSkateTaskManager.contract.Call(opts, &out, "aggregator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_BindingSkateTaskManager *BindingSkateTaskManagerSession) Aggregator() (common.Address, error) {
	return _BindingSkateTaskManager.Contract.Aggregator(&_BindingSkateTaskManager.CallOpts)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_BindingSkateTaskManager *BindingSkateTaskManagerCallerSession) Aggregator() (common.Address, error) {
	return _BindingSkateTaskManager.Contract.Aggregator(&_BindingSkateTaskManager.CallOpts)
}

// AllTaskHashes is a free data retrieval call binding the contract method 0x2d89f6fc.
//
// Solidity: function allTaskHashes(uint32 ) view returns(bytes32)
func (_BindingSkateTaskManager *BindingSkateTaskManagerCaller) AllTaskHashes(opts *bind.CallOpts, arg0 uint32) ([32]byte, error) {
	var out []interface{}
	err := _BindingSkateTaskManager.contract.Call(opts, &out, "allTaskHashes", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AllTaskHashes is a free data retrieval call binding the contract method 0x2d89f6fc.
//
// Solidity: function allTaskHashes(uint32 ) view returns(bytes32)
func (_BindingSkateTaskManager *BindingSkateTaskManagerSession) AllTaskHashes(arg0 uint32) ([32]byte, error) {
	return _BindingSkateTaskManager.Contract.AllTaskHashes(&_BindingSkateTaskManager.CallOpts, arg0)
}

// AllTaskHashes is a free data retrieval call binding the contract method 0x2d89f6fc.
//
// Solidity: function allTaskHashes(uint32 ) view returns(bytes32)
func (_BindingSkateTaskManager *BindingSkateTaskManagerCallerSession) AllTaskHashes(arg0 uint32) ([32]byte, error) {
	return _BindingSkateTaskManager.Contract.AllTaskHashes(&_BindingSkateTaskManager.CallOpts, arg0)
}

// AllTaskResponses is a free data retrieval call binding the contract method 0x2cb223d5.
//
// Solidity: function allTaskResponses(uint32 ) view returns(bytes32)
func (_BindingSkateTaskManager *BindingSkateTaskManagerCaller) AllTaskResponses(opts *bind.CallOpts, arg0 uint32) ([32]byte, error) {
	var out []interface{}
	err := _BindingSkateTaskManager.contract.Call(opts, &out, "allTaskResponses", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AllTaskResponses is a free data retrieval call binding the contract method 0x2cb223d5.
//
// Solidity: function allTaskResponses(uint32 ) view returns(bytes32)
func (_BindingSkateTaskManager *BindingSkateTaskManagerSession) AllTaskResponses(arg0 uint32) ([32]byte, error) {
	return _BindingSkateTaskManager.Contract.AllTaskResponses(&_BindingSkateTaskManager.CallOpts, arg0)
}

// AllTaskResponses is a free data retrieval call binding the contract method 0x2cb223d5.
//
// Solidity: function allTaskResponses(uint32 ) view returns(bytes32)
func (_BindingSkateTaskManager *BindingSkateTaskManagerCallerSession) AllTaskResponses(arg0 uint32) ([32]byte, error) {
	return _BindingSkateTaskManager.Contract.AllTaskResponses(&_BindingSkateTaskManager.CallOpts, arg0)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_BindingSkateTaskManager *BindingSkateTaskManagerCaller) BlsApkRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BindingSkateTaskManager.contract.Call(opts, &out, "blsApkRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_BindingSkateTaskManager *BindingSkateTaskManagerSession) BlsApkRegistry() (common.Address, error) {
	return _BindingSkateTaskManager.Contract.BlsApkRegistry(&_BindingSkateTaskManager.CallOpts)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_BindingSkateTaskManager *BindingSkateTaskManagerCallerSession) BlsApkRegistry() (common.Address, error) {
	return _BindingSkateTaskManager.Contract.BlsApkRegistry(&_BindingSkateTaskManager.CallOpts)
}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_BindingSkateTaskManager *BindingSkateTaskManagerCaller) CheckSignatures(opts *bind.CallOpts, msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	var out []interface{}
	err := _BindingSkateTaskManager.contract.Call(opts, &out, "checkSignatures", msgHash, quorumNumbers, referenceBlockNumber, params)

	if err != nil {
		return *new(IBLSSignatureCheckerQuorumStakeTotals), *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(IBLSSignatureCheckerQuorumStakeTotals)).(*IBLSSignatureCheckerQuorumStakeTotals)
	out1 := *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return out0, out1, err

}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_BindingSkateTaskManager *BindingSkateTaskManagerSession) CheckSignatures(msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	return _BindingSkateTaskManager.Contract.CheckSignatures(&_BindingSkateTaskManager.CallOpts, msgHash, quorumNumbers, referenceBlockNumber, params)
}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_BindingSkateTaskManager *BindingSkateTaskManagerCallerSession) CheckSignatures(msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	return _BindingSkateTaskManager.Contract.CheckSignatures(&_BindingSkateTaskManager.CallOpts, msgHash, quorumNumbers, referenceBlockNumber, params)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_BindingSkateTaskManager *BindingSkateTaskManagerCaller) Delegation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BindingSkateTaskManager.contract.Call(opts, &out, "delegation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_BindingSkateTaskManager *BindingSkateTaskManagerSession) Delegation() (common.Address, error) {
	return _BindingSkateTaskManager.Contract.Delegation(&_BindingSkateTaskManager.CallOpts)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_BindingSkateTaskManager *BindingSkateTaskManagerCallerSession) Delegation() (common.Address, error) {
	return _BindingSkateTaskManager.Contract.Delegation(&_BindingSkateTaskManager.CallOpts)
}

// Generator is a free data retrieval call binding the contract method 0x7afa1eed.
//
// Solidity: function generator() view returns(address)
func (_BindingSkateTaskManager *BindingSkateTaskManagerCaller) Generator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BindingSkateTaskManager.contract.Call(opts, &out, "generator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Generator is a free data retrieval call binding the contract method 0x7afa1eed.
//
// Solidity: function generator() view returns(address)
func (_BindingSkateTaskManager *BindingSkateTaskManagerSession) Generator() (common.Address, error) {
	return _BindingSkateTaskManager.Contract.Generator(&_BindingSkateTaskManager.CallOpts)
}

// Generator is a free data retrieval call binding the contract method 0x7afa1eed.
//
// Solidity: function generator() view returns(address)
func (_BindingSkateTaskManager *BindingSkateTaskManagerCallerSession) Generator() (common.Address, error) {
	return _BindingSkateTaskManager.Contract.Generator(&_BindingSkateTaskManager.CallOpts)
}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_BindingSkateTaskManager *BindingSkateTaskManagerCaller) GetCheckSignaturesIndices(opts *bind.CallOpts, registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	var out []interface{}
	err := _BindingSkateTaskManager.contract.Call(opts, &out, "getCheckSignaturesIndices", registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)

	if err != nil {
		return *new(OperatorStateRetrieverCheckSignaturesIndices), err
	}

	out0 := *abi.ConvertType(out[0], new(OperatorStateRetrieverCheckSignaturesIndices)).(*OperatorStateRetrieverCheckSignaturesIndices)

	return out0, err

}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_BindingSkateTaskManager *BindingSkateTaskManagerSession) GetCheckSignaturesIndices(registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	return _BindingSkateTaskManager.Contract.GetCheckSignaturesIndices(&_BindingSkateTaskManager.CallOpts, registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)
}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_BindingSkateTaskManager *BindingSkateTaskManagerCallerSession) GetCheckSignaturesIndices(registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	return _BindingSkateTaskManager.Contract.GetCheckSignaturesIndices(&_BindingSkateTaskManager.CallOpts, registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)
}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_BindingSkateTaskManager *BindingSkateTaskManagerCaller) GetOperatorState(opts *bind.CallOpts, registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	var out []interface{}
	err := _BindingSkateTaskManager.contract.Call(opts, &out, "getOperatorState", registryCoordinator, quorumNumbers, blockNumber)

	if err != nil {
		return *new([][]OperatorStateRetrieverOperator), err
	}

	out0 := *abi.ConvertType(out[0], new([][]OperatorStateRetrieverOperator)).(*[][]OperatorStateRetrieverOperator)

	return out0, err

}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_BindingSkateTaskManager *BindingSkateTaskManagerSession) GetOperatorState(registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	return _BindingSkateTaskManager.Contract.GetOperatorState(&_BindingSkateTaskManager.CallOpts, registryCoordinator, quorumNumbers, blockNumber)
}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_BindingSkateTaskManager *BindingSkateTaskManagerCallerSession) GetOperatorState(registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	return _BindingSkateTaskManager.Contract.GetOperatorState(&_BindingSkateTaskManager.CallOpts, registryCoordinator, quorumNumbers, blockNumber)
}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (address,bytes32,uint96)[][])
func (_BindingSkateTaskManager *BindingSkateTaskManagerCaller) GetOperatorState0(opts *bind.CallOpts, registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	var out []interface{}
	err := _BindingSkateTaskManager.contract.Call(opts, &out, "getOperatorState0", registryCoordinator, operatorId, blockNumber)

	if err != nil {
		return *new(*big.Int), *new([][]OperatorStateRetrieverOperator), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new([][]OperatorStateRetrieverOperator)).(*[][]OperatorStateRetrieverOperator)

	return out0, out1, err

}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (address,bytes32,uint96)[][])
func (_BindingSkateTaskManager *BindingSkateTaskManagerSession) GetOperatorState0(registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	return _BindingSkateTaskManager.Contract.GetOperatorState0(&_BindingSkateTaskManager.CallOpts, registryCoordinator, operatorId, blockNumber)
}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (address,bytes32,uint96)[][])
func (_BindingSkateTaskManager *BindingSkateTaskManagerCallerSession) GetOperatorState0(registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	return _BindingSkateTaskManager.Contract.GetOperatorState0(&_BindingSkateTaskManager.CallOpts, registryCoordinator, operatorId, blockNumber)
}

// GetQuorumBitmapsAtBlockNumber is a free data retrieval call binding the contract method 0x5c155662.
//
// Solidity: function getQuorumBitmapsAtBlockNumber(address registryCoordinator, bytes32[] operatorIds, uint32 blockNumber) view returns(uint256[])
func (_BindingSkateTaskManager *BindingSkateTaskManagerCaller) GetQuorumBitmapsAtBlockNumber(opts *bind.CallOpts, registryCoordinator common.Address, operatorIds [][32]byte, blockNumber uint32) ([]*big.Int, error) {
	var out []interface{}
	err := _BindingSkateTaskManager.contract.Call(opts, &out, "getQuorumBitmapsAtBlockNumber", registryCoordinator, operatorIds, blockNumber)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetQuorumBitmapsAtBlockNumber is a free data retrieval call binding the contract method 0x5c155662.
//
// Solidity: function getQuorumBitmapsAtBlockNumber(address registryCoordinator, bytes32[] operatorIds, uint32 blockNumber) view returns(uint256[])
func (_BindingSkateTaskManager *BindingSkateTaskManagerSession) GetQuorumBitmapsAtBlockNumber(registryCoordinator common.Address, operatorIds [][32]byte, blockNumber uint32) ([]*big.Int, error) {
	return _BindingSkateTaskManager.Contract.GetQuorumBitmapsAtBlockNumber(&_BindingSkateTaskManager.CallOpts, registryCoordinator, operatorIds, blockNumber)
}

// GetQuorumBitmapsAtBlockNumber is a free data retrieval call binding the contract method 0x5c155662.
//
// Solidity: function getQuorumBitmapsAtBlockNumber(address registryCoordinator, bytes32[] operatorIds, uint32 blockNumber) view returns(uint256[])
func (_BindingSkateTaskManager *BindingSkateTaskManagerCallerSession) GetQuorumBitmapsAtBlockNumber(registryCoordinator common.Address, operatorIds [][32]byte, blockNumber uint32) ([]*big.Int, error) {
	return _BindingSkateTaskManager.Contract.GetQuorumBitmapsAtBlockNumber(&_BindingSkateTaskManager.CallOpts, registryCoordinator, operatorIds, blockNumber)
}

// GetTaskResponseWindowBlock is a free data retrieval call binding the contract method 0xf5c9899d.
//
// Solidity: function getTaskResponseWindowBlock() view returns(uint32)
func (_BindingSkateTaskManager *BindingSkateTaskManagerCaller) GetTaskResponseWindowBlock(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _BindingSkateTaskManager.contract.Call(opts, &out, "getTaskResponseWindowBlock")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetTaskResponseWindowBlock is a free data retrieval call binding the contract method 0xf5c9899d.
//
// Solidity: function getTaskResponseWindowBlock() view returns(uint32)
func (_BindingSkateTaskManager *BindingSkateTaskManagerSession) GetTaskResponseWindowBlock() (uint32, error) {
	return _BindingSkateTaskManager.Contract.GetTaskResponseWindowBlock(&_BindingSkateTaskManager.CallOpts)
}

// GetTaskResponseWindowBlock is a free data retrieval call binding the contract method 0xf5c9899d.
//
// Solidity: function getTaskResponseWindowBlock() view returns(uint32)
func (_BindingSkateTaskManager *BindingSkateTaskManagerCallerSession) GetTaskResponseWindowBlock() (uint32, error) {
	return _BindingSkateTaskManager.Contract.GetTaskResponseWindowBlock(&_BindingSkateTaskManager.CallOpts)
}

// LatestTaskNum is a free data retrieval call binding the contract method 0x8b00ce7c.
//
// Solidity: function latestTaskNum() view returns(uint32)
func (_BindingSkateTaskManager *BindingSkateTaskManagerCaller) LatestTaskNum(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _BindingSkateTaskManager.contract.Call(opts, &out, "latestTaskNum")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LatestTaskNum is a free data retrieval call binding the contract method 0x8b00ce7c.
//
// Solidity: function latestTaskNum() view returns(uint32)
func (_BindingSkateTaskManager *BindingSkateTaskManagerSession) LatestTaskNum() (uint32, error) {
	return _BindingSkateTaskManager.Contract.LatestTaskNum(&_BindingSkateTaskManager.CallOpts)
}

// LatestTaskNum is a free data retrieval call binding the contract method 0x8b00ce7c.
//
// Solidity: function latestTaskNum() view returns(uint32)
func (_BindingSkateTaskManager *BindingSkateTaskManagerCallerSession) LatestTaskNum() (uint32, error) {
	return _BindingSkateTaskManager.Contract.LatestTaskNum(&_BindingSkateTaskManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BindingSkateTaskManager *BindingSkateTaskManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BindingSkateTaskManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BindingSkateTaskManager *BindingSkateTaskManagerSession) Owner() (common.Address, error) {
	return _BindingSkateTaskManager.Contract.Owner(&_BindingSkateTaskManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BindingSkateTaskManager *BindingSkateTaskManagerCallerSession) Owner() (common.Address, error) {
	return _BindingSkateTaskManager.Contract.Owner(&_BindingSkateTaskManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_BindingSkateTaskManager *BindingSkateTaskManagerCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _BindingSkateTaskManager.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_BindingSkateTaskManager *BindingSkateTaskManagerSession) Paused(index uint8) (bool, error) {
	return _BindingSkateTaskManager.Contract.Paused(&_BindingSkateTaskManager.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_BindingSkateTaskManager *BindingSkateTaskManagerCallerSession) Paused(index uint8) (bool, error) {
	return _BindingSkateTaskManager.Contract.Paused(&_BindingSkateTaskManager.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_BindingSkateTaskManager *BindingSkateTaskManagerCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BindingSkateTaskManager.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_BindingSkateTaskManager *BindingSkateTaskManagerSession) Paused0() (*big.Int, error) {
	return _BindingSkateTaskManager.Contract.Paused0(&_BindingSkateTaskManager.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_BindingSkateTaskManager *BindingSkateTaskManagerCallerSession) Paused0() (*big.Int, error) {
	return _BindingSkateTaskManager.Contract.Paused0(&_BindingSkateTaskManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_BindingSkateTaskManager *BindingSkateTaskManagerCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BindingSkateTaskManager.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_BindingSkateTaskManager *BindingSkateTaskManagerSession) PauserRegistry() (common.Address, error) {
	return _BindingSkateTaskManager.Contract.PauserRegistry(&_BindingSkateTaskManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_BindingSkateTaskManager *BindingSkateTaskManagerCallerSession) PauserRegistry() (common.Address, error) {
	return _BindingSkateTaskManager.Contract.PauserRegistry(&_BindingSkateTaskManager.CallOpts)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_BindingSkateTaskManager *BindingSkateTaskManagerCaller) RegistryCoordinator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BindingSkateTaskManager.contract.Call(opts, &out, "registryCoordinator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_BindingSkateTaskManager *BindingSkateTaskManagerSession) RegistryCoordinator() (common.Address, error) {
	return _BindingSkateTaskManager.Contract.RegistryCoordinator(&_BindingSkateTaskManager.CallOpts)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_BindingSkateTaskManager *BindingSkateTaskManagerCallerSession) RegistryCoordinator() (common.Address, error) {
	return _BindingSkateTaskManager.Contract.RegistryCoordinator(&_BindingSkateTaskManager.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_BindingSkateTaskManager *BindingSkateTaskManagerCaller) StakeRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BindingSkateTaskManager.contract.Call(opts, &out, "stakeRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_BindingSkateTaskManager *BindingSkateTaskManagerSession) StakeRegistry() (common.Address, error) {
	return _BindingSkateTaskManager.Contract.StakeRegistry(&_BindingSkateTaskManager.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_BindingSkateTaskManager *BindingSkateTaskManagerCallerSession) StakeRegistry() (common.Address, error) {
	return _BindingSkateTaskManager.Contract.StakeRegistry(&_BindingSkateTaskManager.CallOpts)
}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_BindingSkateTaskManager *BindingSkateTaskManagerCaller) StaleStakesForbidden(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BindingSkateTaskManager.contract.Call(opts, &out, "staleStakesForbidden")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_BindingSkateTaskManager *BindingSkateTaskManagerSession) StaleStakesForbidden() (bool, error) {
	return _BindingSkateTaskManager.Contract.StaleStakesForbidden(&_BindingSkateTaskManager.CallOpts)
}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_BindingSkateTaskManager *BindingSkateTaskManagerCallerSession) StaleStakesForbidden() (bool, error) {
	return _BindingSkateTaskManager.Contract.StaleStakesForbidden(&_BindingSkateTaskManager.CallOpts)
}

// TaskNumber is a free data retrieval call binding the contract method 0x72d18e8d.
//
// Solidity: function taskNumber() view returns(uint32)
func (_BindingSkateTaskManager *BindingSkateTaskManagerCaller) TaskNumber(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _BindingSkateTaskManager.contract.Call(opts, &out, "taskNumber")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TaskNumber is a free data retrieval call binding the contract method 0x72d18e8d.
//
// Solidity: function taskNumber() view returns(uint32)
func (_BindingSkateTaskManager *BindingSkateTaskManagerSession) TaskNumber() (uint32, error) {
	return _BindingSkateTaskManager.Contract.TaskNumber(&_BindingSkateTaskManager.CallOpts)
}

// TaskNumber is a free data retrieval call binding the contract method 0x72d18e8d.
//
// Solidity: function taskNumber() view returns(uint32)
func (_BindingSkateTaskManager *BindingSkateTaskManagerCallerSession) TaskNumber() (uint32, error) {
	return _BindingSkateTaskManager.Contract.TaskNumber(&_BindingSkateTaskManager.CallOpts)
}

// TaskSuccesfullyChallenged is a free data retrieval call binding the contract method 0x5decc3f5.
//
// Solidity: function taskSuccesfullyChallenged(uint32 ) view returns(bool)
func (_BindingSkateTaskManager *BindingSkateTaskManagerCaller) TaskSuccesfullyChallenged(opts *bind.CallOpts, arg0 uint32) (bool, error) {
	var out []interface{}
	err := _BindingSkateTaskManager.contract.Call(opts, &out, "taskSuccesfullyChallenged", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TaskSuccesfullyChallenged is a free data retrieval call binding the contract method 0x5decc3f5.
//
// Solidity: function taskSuccesfullyChallenged(uint32 ) view returns(bool)
func (_BindingSkateTaskManager *BindingSkateTaskManagerSession) TaskSuccesfullyChallenged(arg0 uint32) (bool, error) {
	return _BindingSkateTaskManager.Contract.TaskSuccesfullyChallenged(&_BindingSkateTaskManager.CallOpts, arg0)
}

// TaskSuccesfullyChallenged is a free data retrieval call binding the contract method 0x5decc3f5.
//
// Solidity: function taskSuccesfullyChallenged(uint32 ) view returns(bool)
func (_BindingSkateTaskManager *BindingSkateTaskManagerCallerSession) TaskSuccesfullyChallenged(arg0 uint32) (bool, error) {
	return _BindingSkateTaskManager.Contract.TaskSuccesfullyChallenged(&_BindingSkateTaskManager.CallOpts, arg0)
}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_BindingSkateTaskManager *BindingSkateTaskManagerCaller) TrySignatureAndApkVerification(opts *bind.CallOpts, msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	var out []interface{}
	err := _BindingSkateTaskManager.contract.Call(opts, &out, "trySignatureAndApkVerification", msgHash, apk, apkG2, sigma)

	outstruct := new(struct {
		PairingSuccessful bool
		SiganatureIsValid bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PairingSuccessful = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.SiganatureIsValid = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_BindingSkateTaskManager *BindingSkateTaskManagerSession) TrySignatureAndApkVerification(msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	return _BindingSkateTaskManager.Contract.TrySignatureAndApkVerification(&_BindingSkateTaskManager.CallOpts, msgHash, apk, apkG2, sigma)
}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_BindingSkateTaskManager *BindingSkateTaskManagerCallerSession) TrySignatureAndApkVerification(msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	return _BindingSkateTaskManager.Contract.TrySignatureAndApkVerification(&_BindingSkateTaskManager.CallOpts, msgHash, apk, apkG2, sigma)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0xb00b6c0a.
//
// Solidity: function createNewTask(string message, uint32 quorumThresholdPercentage, bytes quorumNumbers) returns()
func (_BindingSkateTaskManager *BindingSkateTaskManagerTransactor) CreateNewTask(opts *bind.TransactOpts, message string, quorumThresholdPercentage uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _BindingSkateTaskManager.contract.Transact(opts, "createNewTask", message, quorumThresholdPercentage, quorumNumbers)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0xb00b6c0a.
//
// Solidity: function createNewTask(string message, uint32 quorumThresholdPercentage, bytes quorumNumbers) returns()
func (_BindingSkateTaskManager *BindingSkateTaskManagerSession) CreateNewTask(message string, quorumThresholdPercentage uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _BindingSkateTaskManager.Contract.CreateNewTask(&_BindingSkateTaskManager.TransactOpts, message, quorumThresholdPercentage, quorumNumbers)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0xb00b6c0a.
//
// Solidity: function createNewTask(string message, uint32 quorumThresholdPercentage, bytes quorumNumbers) returns()
func (_BindingSkateTaskManager *BindingSkateTaskManagerTransactorSession) CreateNewTask(message string, quorumThresholdPercentage uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _BindingSkateTaskManager.Contract.CreateNewTask(&_BindingSkateTaskManager.TransactOpts, message, quorumThresholdPercentage, quorumNumbers)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _pauserRegistry, address initialOwner, address _aggregator, address _generator) returns()
func (_BindingSkateTaskManager *BindingSkateTaskManagerTransactor) Initialize(opts *bind.TransactOpts, _pauserRegistry common.Address, initialOwner common.Address, _aggregator common.Address, _generator common.Address) (*types.Transaction, error) {
	return _BindingSkateTaskManager.contract.Transact(opts, "initialize", _pauserRegistry, initialOwner, _aggregator, _generator)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _pauserRegistry, address initialOwner, address _aggregator, address _generator) returns()
func (_BindingSkateTaskManager *BindingSkateTaskManagerSession) Initialize(_pauserRegistry common.Address, initialOwner common.Address, _aggregator common.Address, _generator common.Address) (*types.Transaction, error) {
	return _BindingSkateTaskManager.Contract.Initialize(&_BindingSkateTaskManager.TransactOpts, _pauserRegistry, initialOwner, _aggregator, _generator)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _pauserRegistry, address initialOwner, address _aggregator, address _generator) returns()
func (_BindingSkateTaskManager *BindingSkateTaskManagerTransactorSession) Initialize(_pauserRegistry common.Address, initialOwner common.Address, _aggregator common.Address, _generator common.Address) (*types.Transaction, error) {
	return _BindingSkateTaskManager.Contract.Initialize(&_BindingSkateTaskManager.TransactOpts, _pauserRegistry, initialOwner, _aggregator, _generator)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_BindingSkateTaskManager *BindingSkateTaskManagerTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _BindingSkateTaskManager.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_BindingSkateTaskManager *BindingSkateTaskManagerSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _BindingSkateTaskManager.Contract.Pause(&_BindingSkateTaskManager.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_BindingSkateTaskManager *BindingSkateTaskManagerTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _BindingSkateTaskManager.Contract.Pause(&_BindingSkateTaskManager.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_BindingSkateTaskManager *BindingSkateTaskManagerTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BindingSkateTaskManager.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_BindingSkateTaskManager *BindingSkateTaskManagerSession) PauseAll() (*types.Transaction, error) {
	return _BindingSkateTaskManager.Contract.PauseAll(&_BindingSkateTaskManager.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_BindingSkateTaskManager *BindingSkateTaskManagerTransactorSession) PauseAll() (*types.Transaction, error) {
	return _BindingSkateTaskManager.Contract.PauseAll(&_BindingSkateTaskManager.TransactOpts)
}

// RaiseAndResolveChallenge is a paid mutator transaction binding the contract method 0x5d565c29.
//
// Solidity: function raiseAndResolveChallenge((string,uint32,bytes,uint32) task, (uint32,string) taskResponse, (uint32,bytes32) taskResponseMetadata, (uint256,uint256)[] pubkeysOfNonSigningOperators) returns()
func (_BindingSkateTaskManager *BindingSkateTaskManagerTransactor) RaiseAndResolveChallenge(opts *bind.TransactOpts, task ISkateTaskManagerTask, taskResponse ISkateTaskManagerTaskResponse, taskResponseMetadata ISkateTaskManagerTaskResponseMetadata, pubkeysOfNonSigningOperators []BN254G1Point) (*types.Transaction, error) {
	return _BindingSkateTaskManager.contract.Transact(opts, "raiseAndResolveChallenge", task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators)
}

// RaiseAndResolveChallenge is a paid mutator transaction binding the contract method 0x5d565c29.
//
// Solidity: function raiseAndResolveChallenge((string,uint32,bytes,uint32) task, (uint32,string) taskResponse, (uint32,bytes32) taskResponseMetadata, (uint256,uint256)[] pubkeysOfNonSigningOperators) returns()
func (_BindingSkateTaskManager *BindingSkateTaskManagerSession) RaiseAndResolveChallenge(task ISkateTaskManagerTask, taskResponse ISkateTaskManagerTaskResponse, taskResponseMetadata ISkateTaskManagerTaskResponseMetadata, pubkeysOfNonSigningOperators []BN254G1Point) (*types.Transaction, error) {
	return _BindingSkateTaskManager.Contract.RaiseAndResolveChallenge(&_BindingSkateTaskManager.TransactOpts, task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators)
}

// RaiseAndResolveChallenge is a paid mutator transaction binding the contract method 0x5d565c29.
//
// Solidity: function raiseAndResolveChallenge((string,uint32,bytes,uint32) task, (uint32,string) taskResponse, (uint32,bytes32) taskResponseMetadata, (uint256,uint256)[] pubkeysOfNonSigningOperators) returns()
func (_BindingSkateTaskManager *BindingSkateTaskManagerTransactorSession) RaiseAndResolveChallenge(task ISkateTaskManagerTask, taskResponse ISkateTaskManagerTaskResponse, taskResponseMetadata ISkateTaskManagerTaskResponseMetadata, pubkeysOfNonSigningOperators []BN254G1Point) (*types.Transaction, error) {
	return _BindingSkateTaskManager.Contract.RaiseAndResolveChallenge(&_BindingSkateTaskManager.TransactOpts, task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BindingSkateTaskManager *BindingSkateTaskManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BindingSkateTaskManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BindingSkateTaskManager *BindingSkateTaskManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _BindingSkateTaskManager.Contract.RenounceOwnership(&_BindingSkateTaskManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BindingSkateTaskManager *BindingSkateTaskManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BindingSkateTaskManager.Contract.RenounceOwnership(&_BindingSkateTaskManager.TransactOpts)
}

// RespondToTask is a paid mutator transaction binding the contract method 0xd07241f0.
//
// Solidity: function respondToTask((string,uint32,bytes,uint32) task, (uint32,string) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_BindingSkateTaskManager *BindingSkateTaskManagerTransactor) RespondToTask(opts *bind.TransactOpts, task ISkateTaskManagerTask, taskResponse ISkateTaskManagerTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _BindingSkateTaskManager.contract.Transact(opts, "respondToTask", task, taskResponse, nonSignerStakesAndSignature)
}

// RespondToTask is a paid mutator transaction binding the contract method 0xd07241f0.
//
// Solidity: function respondToTask((string,uint32,bytes,uint32) task, (uint32,string) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_BindingSkateTaskManager *BindingSkateTaskManagerSession) RespondToTask(task ISkateTaskManagerTask, taskResponse ISkateTaskManagerTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _BindingSkateTaskManager.Contract.RespondToTask(&_BindingSkateTaskManager.TransactOpts, task, taskResponse, nonSignerStakesAndSignature)
}

// RespondToTask is a paid mutator transaction binding the contract method 0xd07241f0.
//
// Solidity: function respondToTask((string,uint32,bytes,uint32) task, (uint32,string) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_BindingSkateTaskManager *BindingSkateTaskManagerTransactorSession) RespondToTask(task ISkateTaskManagerTask, taskResponse ISkateTaskManagerTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _BindingSkateTaskManager.Contract.RespondToTask(&_BindingSkateTaskManager.TransactOpts, task, taskResponse, nonSignerStakesAndSignature)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_BindingSkateTaskManager *BindingSkateTaskManagerTransactor) SetPauserRegistry(opts *bind.TransactOpts, newPauserRegistry common.Address) (*types.Transaction, error) {
	return _BindingSkateTaskManager.contract.Transact(opts, "setPauserRegistry", newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_BindingSkateTaskManager *BindingSkateTaskManagerSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _BindingSkateTaskManager.Contract.SetPauserRegistry(&_BindingSkateTaskManager.TransactOpts, newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_BindingSkateTaskManager *BindingSkateTaskManagerTransactorSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _BindingSkateTaskManager.Contract.SetPauserRegistry(&_BindingSkateTaskManager.TransactOpts, newPauserRegistry)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_BindingSkateTaskManager *BindingSkateTaskManagerTransactor) SetStaleStakesForbidden(opts *bind.TransactOpts, value bool) (*types.Transaction, error) {
	return _BindingSkateTaskManager.contract.Transact(opts, "setStaleStakesForbidden", value)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_BindingSkateTaskManager *BindingSkateTaskManagerSession) SetStaleStakesForbidden(value bool) (*types.Transaction, error) {
	return _BindingSkateTaskManager.Contract.SetStaleStakesForbidden(&_BindingSkateTaskManager.TransactOpts, value)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_BindingSkateTaskManager *BindingSkateTaskManagerTransactorSession) SetStaleStakesForbidden(value bool) (*types.Transaction, error) {
	return _BindingSkateTaskManager.Contract.SetStaleStakesForbidden(&_BindingSkateTaskManager.TransactOpts, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BindingSkateTaskManager *BindingSkateTaskManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BindingSkateTaskManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BindingSkateTaskManager *BindingSkateTaskManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BindingSkateTaskManager.Contract.TransferOwnership(&_BindingSkateTaskManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BindingSkateTaskManager *BindingSkateTaskManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BindingSkateTaskManager.Contract.TransferOwnership(&_BindingSkateTaskManager.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_BindingSkateTaskManager *BindingSkateTaskManagerTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _BindingSkateTaskManager.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_BindingSkateTaskManager *BindingSkateTaskManagerSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _BindingSkateTaskManager.Contract.Unpause(&_BindingSkateTaskManager.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_BindingSkateTaskManager *BindingSkateTaskManagerTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _BindingSkateTaskManager.Contract.Unpause(&_BindingSkateTaskManager.TransactOpts, newPausedStatus)
}

// BindingSkateTaskManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the BindingSkateTaskManager contract.
type BindingSkateTaskManagerInitializedIterator struct {
	Event *BindingSkateTaskManagerInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BindingSkateTaskManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingSkateTaskManagerInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BindingSkateTaskManagerInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BindingSkateTaskManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingSkateTaskManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingSkateTaskManagerInitialized represents a Initialized event raised by the BindingSkateTaskManager contract.
type BindingSkateTaskManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_BindingSkateTaskManager *BindingSkateTaskManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*BindingSkateTaskManagerInitializedIterator, error) {

	logs, sub, err := _BindingSkateTaskManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BindingSkateTaskManagerInitializedIterator{contract: _BindingSkateTaskManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_BindingSkateTaskManager *BindingSkateTaskManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BindingSkateTaskManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _BindingSkateTaskManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingSkateTaskManagerInitialized)
				if err := _BindingSkateTaskManager.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_BindingSkateTaskManager *BindingSkateTaskManagerFilterer) ParseInitialized(log types.Log) (*BindingSkateTaskManagerInitialized, error) {
	event := new(BindingSkateTaskManagerInitialized)
	if err := _BindingSkateTaskManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingSkateTaskManagerNewTaskCreatedIterator is returned from FilterNewTaskCreated and is used to iterate over the raw logs and unpacked data for NewTaskCreated events raised by the BindingSkateTaskManager contract.
type BindingSkateTaskManagerNewTaskCreatedIterator struct {
	Event *BindingSkateTaskManagerNewTaskCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BindingSkateTaskManagerNewTaskCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingSkateTaskManagerNewTaskCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BindingSkateTaskManagerNewTaskCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BindingSkateTaskManagerNewTaskCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingSkateTaskManagerNewTaskCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingSkateTaskManagerNewTaskCreated represents a NewTaskCreated event raised by the BindingSkateTaskManager contract.
type BindingSkateTaskManagerNewTaskCreated struct {
	TaskIndex uint32
	Task      ISkateTaskManagerTask
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewTaskCreated is a free log retrieval operation binding the contract event 0xb456431ed13b80d91ab280ac82018e48afde7e840d354010364822d756914479.
//
// Solidity: event NewTaskCreated(uint32 indexed taskIndex, (string,uint32,bytes,uint32) task)
func (_BindingSkateTaskManager *BindingSkateTaskManagerFilterer) FilterNewTaskCreated(opts *bind.FilterOpts, taskIndex []uint32) (*BindingSkateTaskManagerNewTaskCreatedIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _BindingSkateTaskManager.contract.FilterLogs(opts, "NewTaskCreated", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return &BindingSkateTaskManagerNewTaskCreatedIterator{contract: _BindingSkateTaskManager.contract, event: "NewTaskCreated", logs: logs, sub: sub}, nil
}

// WatchNewTaskCreated is a free log subscription operation binding the contract event 0xb456431ed13b80d91ab280ac82018e48afde7e840d354010364822d756914479.
//
// Solidity: event NewTaskCreated(uint32 indexed taskIndex, (string,uint32,bytes,uint32) task)
func (_BindingSkateTaskManager *BindingSkateTaskManagerFilterer) WatchNewTaskCreated(opts *bind.WatchOpts, sink chan<- *BindingSkateTaskManagerNewTaskCreated, taskIndex []uint32) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _BindingSkateTaskManager.contract.WatchLogs(opts, "NewTaskCreated", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingSkateTaskManagerNewTaskCreated)
				if err := _BindingSkateTaskManager.contract.UnpackLog(event, "NewTaskCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewTaskCreated is a log parse operation binding the contract event 0xb456431ed13b80d91ab280ac82018e48afde7e840d354010364822d756914479.
//
// Solidity: event NewTaskCreated(uint32 indexed taskIndex, (string,uint32,bytes,uint32) task)
func (_BindingSkateTaskManager *BindingSkateTaskManagerFilterer) ParseNewTaskCreated(log types.Log) (*BindingSkateTaskManagerNewTaskCreated, error) {
	event := new(BindingSkateTaskManagerNewTaskCreated)
	if err := _BindingSkateTaskManager.contract.UnpackLog(event, "NewTaskCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingSkateTaskManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BindingSkateTaskManager contract.
type BindingSkateTaskManagerOwnershipTransferredIterator struct {
	Event *BindingSkateTaskManagerOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BindingSkateTaskManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingSkateTaskManagerOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BindingSkateTaskManagerOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BindingSkateTaskManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingSkateTaskManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingSkateTaskManagerOwnershipTransferred represents a OwnershipTransferred event raised by the BindingSkateTaskManager contract.
type BindingSkateTaskManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BindingSkateTaskManager *BindingSkateTaskManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BindingSkateTaskManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BindingSkateTaskManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BindingSkateTaskManagerOwnershipTransferredIterator{contract: _BindingSkateTaskManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BindingSkateTaskManager *BindingSkateTaskManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BindingSkateTaskManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BindingSkateTaskManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingSkateTaskManagerOwnershipTransferred)
				if err := _BindingSkateTaskManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BindingSkateTaskManager *BindingSkateTaskManagerFilterer) ParseOwnershipTransferred(log types.Log) (*BindingSkateTaskManagerOwnershipTransferred, error) {
	event := new(BindingSkateTaskManagerOwnershipTransferred)
	if err := _BindingSkateTaskManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingSkateTaskManagerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the BindingSkateTaskManager contract.
type BindingSkateTaskManagerPausedIterator struct {
	Event *BindingSkateTaskManagerPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BindingSkateTaskManagerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingSkateTaskManagerPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BindingSkateTaskManagerPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BindingSkateTaskManagerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingSkateTaskManagerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingSkateTaskManagerPaused represents a Paused event raised by the BindingSkateTaskManager contract.
type BindingSkateTaskManagerPaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_BindingSkateTaskManager *BindingSkateTaskManagerFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*BindingSkateTaskManagerPausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _BindingSkateTaskManager.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &BindingSkateTaskManagerPausedIterator{contract: _BindingSkateTaskManager.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_BindingSkateTaskManager *BindingSkateTaskManagerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *BindingSkateTaskManagerPaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _BindingSkateTaskManager.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingSkateTaskManagerPaused)
				if err := _BindingSkateTaskManager.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_BindingSkateTaskManager *BindingSkateTaskManagerFilterer) ParsePaused(log types.Log) (*BindingSkateTaskManagerPaused, error) {
	event := new(BindingSkateTaskManagerPaused)
	if err := _BindingSkateTaskManager.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingSkateTaskManagerPauserRegistrySetIterator is returned from FilterPauserRegistrySet and is used to iterate over the raw logs and unpacked data for PauserRegistrySet events raised by the BindingSkateTaskManager contract.
type BindingSkateTaskManagerPauserRegistrySetIterator struct {
	Event *BindingSkateTaskManagerPauserRegistrySet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BindingSkateTaskManagerPauserRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingSkateTaskManagerPauserRegistrySet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BindingSkateTaskManagerPauserRegistrySet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BindingSkateTaskManagerPauserRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingSkateTaskManagerPauserRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingSkateTaskManagerPauserRegistrySet represents a PauserRegistrySet event raised by the BindingSkateTaskManager contract.
type BindingSkateTaskManagerPauserRegistrySet struct {
	PauserRegistry    common.Address
	NewPauserRegistry common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterPauserRegistrySet is a free log retrieval operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_BindingSkateTaskManager *BindingSkateTaskManagerFilterer) FilterPauserRegistrySet(opts *bind.FilterOpts) (*BindingSkateTaskManagerPauserRegistrySetIterator, error) {

	logs, sub, err := _BindingSkateTaskManager.contract.FilterLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return &BindingSkateTaskManagerPauserRegistrySetIterator{contract: _BindingSkateTaskManager.contract, event: "PauserRegistrySet", logs: logs, sub: sub}, nil
}

// WatchPauserRegistrySet is a free log subscription operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_BindingSkateTaskManager *BindingSkateTaskManagerFilterer) WatchPauserRegistrySet(opts *bind.WatchOpts, sink chan<- *BindingSkateTaskManagerPauserRegistrySet) (event.Subscription, error) {

	logs, sub, err := _BindingSkateTaskManager.contract.WatchLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingSkateTaskManagerPauserRegistrySet)
				if err := _BindingSkateTaskManager.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePauserRegistrySet is a log parse operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_BindingSkateTaskManager *BindingSkateTaskManagerFilterer) ParsePauserRegistrySet(log types.Log) (*BindingSkateTaskManagerPauserRegistrySet, error) {
	event := new(BindingSkateTaskManagerPauserRegistrySet)
	if err := _BindingSkateTaskManager.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingSkateTaskManagerStaleStakesForbiddenUpdateIterator is returned from FilterStaleStakesForbiddenUpdate and is used to iterate over the raw logs and unpacked data for StaleStakesForbiddenUpdate events raised by the BindingSkateTaskManager contract.
type BindingSkateTaskManagerStaleStakesForbiddenUpdateIterator struct {
	Event *BindingSkateTaskManagerStaleStakesForbiddenUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BindingSkateTaskManagerStaleStakesForbiddenUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingSkateTaskManagerStaleStakesForbiddenUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BindingSkateTaskManagerStaleStakesForbiddenUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BindingSkateTaskManagerStaleStakesForbiddenUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingSkateTaskManagerStaleStakesForbiddenUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingSkateTaskManagerStaleStakesForbiddenUpdate represents a StaleStakesForbiddenUpdate event raised by the BindingSkateTaskManager contract.
type BindingSkateTaskManagerStaleStakesForbiddenUpdate struct {
	Value bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterStaleStakesForbiddenUpdate is a free log retrieval operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_BindingSkateTaskManager *BindingSkateTaskManagerFilterer) FilterStaleStakesForbiddenUpdate(opts *bind.FilterOpts) (*BindingSkateTaskManagerStaleStakesForbiddenUpdateIterator, error) {

	logs, sub, err := _BindingSkateTaskManager.contract.FilterLogs(opts, "StaleStakesForbiddenUpdate")
	if err != nil {
		return nil, err
	}
	return &BindingSkateTaskManagerStaleStakesForbiddenUpdateIterator{contract: _BindingSkateTaskManager.contract, event: "StaleStakesForbiddenUpdate", logs: logs, sub: sub}, nil
}

// WatchStaleStakesForbiddenUpdate is a free log subscription operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_BindingSkateTaskManager *BindingSkateTaskManagerFilterer) WatchStaleStakesForbiddenUpdate(opts *bind.WatchOpts, sink chan<- *BindingSkateTaskManagerStaleStakesForbiddenUpdate) (event.Subscription, error) {

	logs, sub, err := _BindingSkateTaskManager.contract.WatchLogs(opts, "StaleStakesForbiddenUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingSkateTaskManagerStaleStakesForbiddenUpdate)
				if err := _BindingSkateTaskManager.contract.UnpackLog(event, "StaleStakesForbiddenUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStaleStakesForbiddenUpdate is a log parse operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_BindingSkateTaskManager *BindingSkateTaskManagerFilterer) ParseStaleStakesForbiddenUpdate(log types.Log) (*BindingSkateTaskManagerStaleStakesForbiddenUpdate, error) {
	event := new(BindingSkateTaskManagerStaleStakesForbiddenUpdate)
	if err := _BindingSkateTaskManager.contract.UnpackLog(event, "StaleStakesForbiddenUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingSkateTaskManagerTaskChallengedSuccessfullyIterator is returned from FilterTaskChallengedSuccessfully and is used to iterate over the raw logs and unpacked data for TaskChallengedSuccessfully events raised by the BindingSkateTaskManager contract.
type BindingSkateTaskManagerTaskChallengedSuccessfullyIterator struct {
	Event *BindingSkateTaskManagerTaskChallengedSuccessfully // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BindingSkateTaskManagerTaskChallengedSuccessfullyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingSkateTaskManagerTaskChallengedSuccessfully)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BindingSkateTaskManagerTaskChallengedSuccessfully)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BindingSkateTaskManagerTaskChallengedSuccessfullyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingSkateTaskManagerTaskChallengedSuccessfullyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingSkateTaskManagerTaskChallengedSuccessfully represents a TaskChallengedSuccessfully event raised by the BindingSkateTaskManager contract.
type BindingSkateTaskManagerTaskChallengedSuccessfully struct {
	TaskIndex  uint32
	Challenger common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTaskChallengedSuccessfully is a free log retrieval operation binding the contract event 0xc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec.
//
// Solidity: event TaskChallengedSuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_BindingSkateTaskManager *BindingSkateTaskManagerFilterer) FilterTaskChallengedSuccessfully(opts *bind.FilterOpts, taskIndex []uint32, challenger []common.Address) (*BindingSkateTaskManagerTaskChallengedSuccessfullyIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _BindingSkateTaskManager.contract.FilterLogs(opts, "TaskChallengedSuccessfully", taskIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return &BindingSkateTaskManagerTaskChallengedSuccessfullyIterator{contract: _BindingSkateTaskManager.contract, event: "TaskChallengedSuccessfully", logs: logs, sub: sub}, nil
}

// WatchTaskChallengedSuccessfully is a free log subscription operation binding the contract event 0xc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec.
//
// Solidity: event TaskChallengedSuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_BindingSkateTaskManager *BindingSkateTaskManagerFilterer) WatchTaskChallengedSuccessfully(opts *bind.WatchOpts, sink chan<- *BindingSkateTaskManagerTaskChallengedSuccessfully, taskIndex []uint32, challenger []common.Address) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _BindingSkateTaskManager.contract.WatchLogs(opts, "TaskChallengedSuccessfully", taskIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingSkateTaskManagerTaskChallengedSuccessfully)
				if err := _BindingSkateTaskManager.contract.UnpackLog(event, "TaskChallengedSuccessfully", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTaskChallengedSuccessfully is a log parse operation binding the contract event 0xc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec.
//
// Solidity: event TaskChallengedSuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_BindingSkateTaskManager *BindingSkateTaskManagerFilterer) ParseTaskChallengedSuccessfully(log types.Log) (*BindingSkateTaskManagerTaskChallengedSuccessfully, error) {
	event := new(BindingSkateTaskManagerTaskChallengedSuccessfully)
	if err := _BindingSkateTaskManager.contract.UnpackLog(event, "TaskChallengedSuccessfully", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingSkateTaskManagerTaskChallengedUnsuccessfullyIterator is returned from FilterTaskChallengedUnsuccessfully and is used to iterate over the raw logs and unpacked data for TaskChallengedUnsuccessfully events raised by the BindingSkateTaskManager contract.
type BindingSkateTaskManagerTaskChallengedUnsuccessfullyIterator struct {
	Event *BindingSkateTaskManagerTaskChallengedUnsuccessfully // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BindingSkateTaskManagerTaskChallengedUnsuccessfullyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingSkateTaskManagerTaskChallengedUnsuccessfully)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BindingSkateTaskManagerTaskChallengedUnsuccessfully)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BindingSkateTaskManagerTaskChallengedUnsuccessfullyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingSkateTaskManagerTaskChallengedUnsuccessfullyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingSkateTaskManagerTaskChallengedUnsuccessfully represents a TaskChallengedUnsuccessfully event raised by the BindingSkateTaskManager contract.
type BindingSkateTaskManagerTaskChallengedUnsuccessfully struct {
	TaskIndex  uint32
	Challenger common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTaskChallengedUnsuccessfully is a free log retrieval operation binding the contract event 0xfd3e26beeb5967fc5a57a0446914eabc45b4aa474c67a51b4b5160cac60ddb05.
//
// Solidity: event TaskChallengedUnsuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_BindingSkateTaskManager *BindingSkateTaskManagerFilterer) FilterTaskChallengedUnsuccessfully(opts *bind.FilterOpts, taskIndex []uint32, challenger []common.Address) (*BindingSkateTaskManagerTaskChallengedUnsuccessfullyIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _BindingSkateTaskManager.contract.FilterLogs(opts, "TaskChallengedUnsuccessfully", taskIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return &BindingSkateTaskManagerTaskChallengedUnsuccessfullyIterator{contract: _BindingSkateTaskManager.contract, event: "TaskChallengedUnsuccessfully", logs: logs, sub: sub}, nil
}

// WatchTaskChallengedUnsuccessfully is a free log subscription operation binding the contract event 0xfd3e26beeb5967fc5a57a0446914eabc45b4aa474c67a51b4b5160cac60ddb05.
//
// Solidity: event TaskChallengedUnsuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_BindingSkateTaskManager *BindingSkateTaskManagerFilterer) WatchTaskChallengedUnsuccessfully(opts *bind.WatchOpts, sink chan<- *BindingSkateTaskManagerTaskChallengedUnsuccessfully, taskIndex []uint32, challenger []common.Address) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _BindingSkateTaskManager.contract.WatchLogs(opts, "TaskChallengedUnsuccessfully", taskIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingSkateTaskManagerTaskChallengedUnsuccessfully)
				if err := _BindingSkateTaskManager.contract.UnpackLog(event, "TaskChallengedUnsuccessfully", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTaskChallengedUnsuccessfully is a log parse operation binding the contract event 0xfd3e26beeb5967fc5a57a0446914eabc45b4aa474c67a51b4b5160cac60ddb05.
//
// Solidity: event TaskChallengedUnsuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_BindingSkateTaskManager *BindingSkateTaskManagerFilterer) ParseTaskChallengedUnsuccessfully(log types.Log) (*BindingSkateTaskManagerTaskChallengedUnsuccessfully, error) {
	event := new(BindingSkateTaskManagerTaskChallengedUnsuccessfully)
	if err := _BindingSkateTaskManager.contract.UnpackLog(event, "TaskChallengedUnsuccessfully", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingSkateTaskManagerTaskCompletedIterator is returned from FilterTaskCompleted and is used to iterate over the raw logs and unpacked data for TaskCompleted events raised by the BindingSkateTaskManager contract.
type BindingSkateTaskManagerTaskCompletedIterator struct {
	Event *BindingSkateTaskManagerTaskCompleted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BindingSkateTaskManagerTaskCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingSkateTaskManagerTaskCompleted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BindingSkateTaskManagerTaskCompleted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BindingSkateTaskManagerTaskCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingSkateTaskManagerTaskCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingSkateTaskManagerTaskCompleted represents a TaskCompleted event raised by the BindingSkateTaskManager contract.
type BindingSkateTaskManagerTaskCompleted struct {
	TaskIndex uint32
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTaskCompleted is a free log retrieval operation binding the contract event 0x9a144f228a931b9d0d1696fbcdaf310b24b5d2d21e799db623fc986a0f547430.
//
// Solidity: event TaskCompleted(uint32 indexed taskIndex)
func (_BindingSkateTaskManager *BindingSkateTaskManagerFilterer) FilterTaskCompleted(opts *bind.FilterOpts, taskIndex []uint32) (*BindingSkateTaskManagerTaskCompletedIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _BindingSkateTaskManager.contract.FilterLogs(opts, "TaskCompleted", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return &BindingSkateTaskManagerTaskCompletedIterator{contract: _BindingSkateTaskManager.contract, event: "TaskCompleted", logs: logs, sub: sub}, nil
}

// WatchTaskCompleted is a free log subscription operation binding the contract event 0x9a144f228a931b9d0d1696fbcdaf310b24b5d2d21e799db623fc986a0f547430.
//
// Solidity: event TaskCompleted(uint32 indexed taskIndex)
func (_BindingSkateTaskManager *BindingSkateTaskManagerFilterer) WatchTaskCompleted(opts *bind.WatchOpts, sink chan<- *BindingSkateTaskManagerTaskCompleted, taskIndex []uint32) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _BindingSkateTaskManager.contract.WatchLogs(opts, "TaskCompleted", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingSkateTaskManagerTaskCompleted)
				if err := _BindingSkateTaskManager.contract.UnpackLog(event, "TaskCompleted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTaskCompleted is a log parse operation binding the contract event 0x9a144f228a931b9d0d1696fbcdaf310b24b5d2d21e799db623fc986a0f547430.
//
// Solidity: event TaskCompleted(uint32 indexed taskIndex)
func (_BindingSkateTaskManager *BindingSkateTaskManagerFilterer) ParseTaskCompleted(log types.Log) (*BindingSkateTaskManagerTaskCompleted, error) {
	event := new(BindingSkateTaskManagerTaskCompleted)
	if err := _BindingSkateTaskManager.contract.UnpackLog(event, "TaskCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingSkateTaskManagerTaskRespondedIterator is returned from FilterTaskResponded and is used to iterate over the raw logs and unpacked data for TaskResponded events raised by the BindingSkateTaskManager contract.
type BindingSkateTaskManagerTaskRespondedIterator struct {
	Event *BindingSkateTaskManagerTaskResponded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BindingSkateTaskManagerTaskRespondedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingSkateTaskManagerTaskResponded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BindingSkateTaskManagerTaskResponded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BindingSkateTaskManagerTaskRespondedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingSkateTaskManagerTaskRespondedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingSkateTaskManagerTaskResponded represents a TaskResponded event raised by the BindingSkateTaskManager contract.
type BindingSkateTaskManagerTaskResponded struct {
	TaskResponse         ISkateTaskManagerTaskResponse
	TaskResponseMetadata ISkateTaskManagerTaskResponseMetadata
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterTaskResponded is a free log retrieval operation binding the contract event 0xd0e833257cf2514f7f416921130672e0a5099b141e19ad5285b881ec34e9eb54.
//
// Solidity: event TaskResponded((uint32,string) taskResponse, (uint32,bytes32) taskResponseMetadata)
func (_BindingSkateTaskManager *BindingSkateTaskManagerFilterer) FilterTaskResponded(opts *bind.FilterOpts) (*BindingSkateTaskManagerTaskRespondedIterator, error) {

	logs, sub, err := _BindingSkateTaskManager.contract.FilterLogs(opts, "TaskResponded")
	if err != nil {
		return nil, err
	}
	return &BindingSkateTaskManagerTaskRespondedIterator{contract: _BindingSkateTaskManager.contract, event: "TaskResponded", logs: logs, sub: sub}, nil
}

// WatchTaskResponded is a free log subscription operation binding the contract event 0xd0e833257cf2514f7f416921130672e0a5099b141e19ad5285b881ec34e9eb54.
//
// Solidity: event TaskResponded((uint32,string) taskResponse, (uint32,bytes32) taskResponseMetadata)
func (_BindingSkateTaskManager *BindingSkateTaskManagerFilterer) WatchTaskResponded(opts *bind.WatchOpts, sink chan<- *BindingSkateTaskManagerTaskResponded) (event.Subscription, error) {

	logs, sub, err := _BindingSkateTaskManager.contract.WatchLogs(opts, "TaskResponded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingSkateTaskManagerTaskResponded)
				if err := _BindingSkateTaskManager.contract.UnpackLog(event, "TaskResponded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTaskResponded is a log parse operation binding the contract event 0xd0e833257cf2514f7f416921130672e0a5099b141e19ad5285b881ec34e9eb54.
//
// Solidity: event TaskResponded((uint32,string) taskResponse, (uint32,bytes32) taskResponseMetadata)
func (_BindingSkateTaskManager *BindingSkateTaskManagerFilterer) ParseTaskResponded(log types.Log) (*BindingSkateTaskManagerTaskResponded, error) {
	event := new(BindingSkateTaskManagerTaskResponded)
	if err := _BindingSkateTaskManager.contract.UnpackLog(event, "TaskResponded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingSkateTaskManagerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the BindingSkateTaskManager contract.
type BindingSkateTaskManagerUnpausedIterator struct {
	Event *BindingSkateTaskManagerUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BindingSkateTaskManagerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingSkateTaskManagerUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BindingSkateTaskManagerUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BindingSkateTaskManagerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingSkateTaskManagerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingSkateTaskManagerUnpaused represents a Unpaused event raised by the BindingSkateTaskManager contract.
type BindingSkateTaskManagerUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_BindingSkateTaskManager *BindingSkateTaskManagerFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*BindingSkateTaskManagerUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _BindingSkateTaskManager.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &BindingSkateTaskManagerUnpausedIterator{contract: _BindingSkateTaskManager.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_BindingSkateTaskManager *BindingSkateTaskManagerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *BindingSkateTaskManagerUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _BindingSkateTaskManager.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingSkateTaskManagerUnpaused)
				if err := _BindingSkateTaskManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_BindingSkateTaskManager *BindingSkateTaskManagerFilterer) ParseUnpaused(log types.Log) (*BindingSkateTaskManagerUnpaused, error) {
	event := new(BindingSkateTaskManagerUnpaused)
	if err := _BindingSkateTaskManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
