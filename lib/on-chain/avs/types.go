package avs

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	gethcommon "github.com/ethereum/go-ethereum/common"
	sdkTypes "github.com/Layr-Labs/eigensdk-go/types"
	crypto "skatechain.org/lib/crypto"
)

type OperatorID = sdkTypes.OperatorId

// Security and Quorum Parameters

// QuorumID is a unique identifier for a quorum; initially EigenDA wil support upt to 256 quorums
type QuorumID = uint8


// QuorumResult contains the quorum ID and the amount signed for the quorum
type QuorumResult struct {
	QuorumID QuorumID
	// PercentSigned is percentage of the total stake for the quorum that signed for a particular batch.
	PercentSigned uint8
}

// OperatorInfo contains information about an operator which is stored on the blockchain state,
// corresponding to a particular quorum
type OperatorInfo struct {
	// Stake is the amount of stake held by the operator in the quorum
	Stake sdkTypes.StakeAmount
	// Index is the index of the operator within the quorum
	Index uint
}

// OperatorState contains information about the current state of operators which is stored in the blockchain state
type OperatorState struct {
	// Operators is a map from quorum ID to a map from the operators in that quourm to their StoredOperatorInfo. Membership
	// in the map implies membership in the quorum.
	Operators map[QuorumID]map[OperatorID]*OperatorInfo
	// Totals is a map from quorum ID to the total stake (Stake) and total count (Index) of all operators in that quorum
	Totals map[QuorumID]*OperatorInfo
	// BlockNumber is the block number at which this state was retrieved
	BlockNumber uint
}

// IndexedOperatorInfo contains information about an operator which is contained in events from the EigenDA smart contracts. Note that
// this information does not depend on the quorum.
type IndexedOperatorInfo struct {
	// PubKeyG1 and PubKeyG2 are the public keys of the operator, which are retrieved from the EigenDAPubKeyCompendium smart contract
	PubkeyG1 *crypto.G1Point
	PubkeyG2 *crypto.G2Point
	// Socket is the socket address of the operator, in the form "host:port"
	Socket string
}

// IndexedOperatorState contains information about the current state of operators which is contained in events from the EigenDA smart contracts,
// in addition to the information contained in OperatorState
type IndexedOperatorState struct {
	*OperatorState
	// IndexedOperators is a map from operator ID to the IndexedOperatorInfo for that operator.
	IndexedOperators map[OperatorID]*IndexedOperatorInfo
	// AggKeys is a map from quorum ID to the aggregate public key of the operators in that quorum
	AggKeys map[QuorumID]*crypto.G1Point
}

type Transactor interface {
	// GetRegisteredQuorumIdsForOperator returns the quorum ids that the operator is registered in with the given public key.
	GetRegisteredQuorumIdsForOperator(ctx context.Context, operatorID OperatorID) ([]QuorumID, error)

	// RegisterOperator registers a new operator with the given public key and socket with the provided quorum ids.
	// If the operator is already registered with a given quorum id, the transaction will fail (noop) and an error
	// will be returned.
	RegisterOperator(
		ctx context.Context,
		keypair *crypto.KeyPair,
		socket string,
		quorumIds []QuorumID,
		operatorEcdsaPrivateKey *ecdsa.PrivateKey,
		operatorToAvsRegistrationSigSalt [32]byte,
		operatorToAvsRegistrationSigExpiry *big.Int,
	) error

	// DeregisterOperator deregisters an operator with the given public key from the all the quorums that it is
	// registered with at the supplied block number. To fully deregister an operator, this function should be called
	// with the current block number.
	// If the operator isn't registered with any of the specified quorums, this function will return error, and
	// no quorum will be deregistered.
	DeregisterOperator(ctx context.Context, pubkeyG1 *crypto.G1Point, blockNumber uint32, quorumIds []QuorumID) error

	// UpdateOperatorSocket updates the socket of the operator in all the quorums that it is registered with.
	UpdateOperatorSocket(ctx context.Context, socket string) error

	// StakeRegistry returns the address of the stake registry contract.
	StakeRegistry(ctx context.Context) (gethcommon.Address, error)

	// OperatorIDToAddress returns the address of the operator from the operator id.
	OperatorIDToAddress(ctx context.Context, operatorId OperatorID) (gethcommon.Address, error)

	// BatchOperatorIDToAddress returns the addresses of the operators from the operator id.
	BatchOperatorIDToAddress(ctx context.Context, operatorIds []OperatorID) ([]gethcommon.Address, error)

	// GetCurrentQuorumBitmapByOperatorId returns the current quorum bitmap for the operator.
	GetCurrentQuorumBitmapByOperatorId(ctx context.Context, operatorId OperatorID) (*big.Int, error)

	// GetQuorumBitmapForOperatorsAtBlockNumber returns the quorum bitmaps for the operators at the given block number.
	// The result slice will be of same length as "operatorIds", with the i-th entry be the result for the operatorIds[i].
	// If an operator failed to find bitmap, the corresponding result entry will be an empty bitmap.
	GetQuorumBitmapForOperatorsAtBlockNumber(ctx context.Context, operatorIds []OperatorID, blockNumber uint32) ([]*big.Int, error)

	// GetCurrentBlockNumber returns the current block number.
	GetCurrentBlockNumber(ctx context.Context) (uint32, error)
}
