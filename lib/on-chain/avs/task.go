package avs

import (
	"encoding/binary"
	"encoding/hex"
	"math/big"
	"strings"

	"skatechain.org/lib/crypto/ecdsa"
	// "github.com/ethereum/go-ethereum/accounts/abi"

	pb "skatechain.org/api/pb/relayer"
)

// Calculate the pack encoded data of a TaskData
//
// Format: chainType (2bytes) | chainId (4bytes) | initiator (20bytes for EVM || ? for non-EVM) | msg (x-bytes)
func TaskData(
	msg string,
	initiator string,
	chainType pb.ChainType,
	chainId uint32,
) []byte {
	switch chainType {
	case pb.ChainType_EVM:
		// Convert uint16 chainType to 2 bytes
		chainTypeBytes := make([]byte, 2)
		binary.BigEndian.PutUint16(chainTypeBytes, uint16(chainType))

		// Convert uint32 to 4 bytes
		chainIdBytes := make([]byte, 4)
		binary.BigEndian.PutUint32(chainIdBytes, chainId)

		// Convert the Ethereum address (initiator) from hex string to byte slice
		// Strip the '0x' prefix if present and decode hex to bytes
		if strings.HasPrefix(initiator, "0x") {
			initiator = initiator[2:]
		}
		addressBytes, err := hex.DecodeString(initiator)
		if err != nil {
			panic("Invalid hex in initiator address")
		}

		msgBytes := []byte(msg)

		// Concatenate all byte slices
		result := append(chainTypeBytes, chainIdBytes...)
		result = append(result, addressBytes...)
		result = append(result, msgBytes...)

		return result

	default:
		panic("HashTask for unknown chain")
	}
}

func TaskDigestHash(
	taskId uint32,
	msg string,
	initiator string,
	chainType pb.ChainType,
	chainId uint32,
) []byte {
	switch chainType {
	case pb.ChainType_EVM:
		buf32 := make([]byte, 32) // taskId is uint256 in avs contract
		taskIdBytes := new(big.Int).SetUint64(uint64(taskId)).FillBytes(buf32)
		msgBytes := TaskData(msg, initiator, chainType, chainId)

		return ecdsa.Keccak256Message(taskIdBytes, msgBytes)
	default:
		panic("ecdsa.TaskDigestHash: not implemented!")
	}
}
