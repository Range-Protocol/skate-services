package avs

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	// sdkAvsClient "github.com/Layr-Labs/eigensdk-go/chainio/clients/avsregistry"
	sdkElClient "github.com/Layr-Labs/eigensdk-go/chainio/clients/elcontracts"
	sdkTypes "github.com/Layr-Labs/eigensdk-go/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	avsdirectory "skatechain.org/contracts/bindings/AVSDirectory"
	"skatechain.org/lib/logging"
	"skatechain.org/lib/metrics"
	"skatechain.org/lib/on-chain/backend"
)

const (
	AVS_NAME = "skate-avs"
	SEM_VER  = "0.1.0"
)

// NOTE: operator registration interface
//
// struct SignatureWithSaltAndExpiry {
//     // the signature itself, formatted as a single bytes object
//     bytes signature;
//     // the salt used to generate the signature
//     bytes32 salt;
//     // the expiration timestamp (UTC) of the signature
//     uint256 expiry;
// }
//
// function registerOperatorToAVS(
//     address operator,
//     ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature
// ) external;

// NOTE: Sig digest for the avs operator registration
//
// bytes32 public constant OPERATOR_AVS_REGISTRATION_TYPEHASH =
//   keccak256("OperatorAVSRegistration(address operator,address avs,bytes32 salt,uint256 expiry)");
//
// /**
//  * @param operator The account registering as an operator
//  * @param avs The address of the service manager contract for the AVS that the operator is registering to
//  * @param salt A unique and single use value associated with the approver signature.
//  * @param expiry Time after which the approver's signature becomes invalid
//  */
// function calculateOperatorAVSRegistrationDigestHash(
//     address operator,
//     address avs,
//     bytes32 salt,
//     uint256 expiry
// ) public view returns (bytes32) {
//     // calculate the struct hash
//     bytes32 structHash = keccak256(
//         abi.encode(OPERATOR_AVS_REGISTRATION_TYPEHASH, operator, avs, salt, expiry)
//     );
//     // calculate the digest hash
//     bytes32 digestHash = keccak256(
//         abi.encodePacked("\x19\x01", domainSeparator(), structHash)
//     );
//     return digestHash;
// }

func RegisterOperatorWithAVS(
	ctx context.Context, opPrivkey *ecdsa.PrivateKey,
	backend *backend.Backend, avsDirectoryAddr common.Address,
	salt [32]byte, expiry *big.Int,
) (*gethtypes.Receipt, error) {
	opAddr := crypto.PubkeyToAddress(opPrivkey.PublicKey)

	avsDir, err := avsdirectory.NewBindingAVSDirectoryCaller(avsDirectoryAddr, backend)
	if err != nil {
		return nil, err
	}

	digestHash, err := avsDir.CalculateOperatorAVSRegistrationDigestHash(
		&bind.CallOpts{},
		opAddr, avsDirectoryAddr,
		salt, expiry,
	)

	if err != nil {
		return nil, err
	}

	opSig, err := crypto.Sign(digestHash[:], opPrivkey)
	opSig[64] += 27 // INFO: convert v of [r,s,v] into EVM format, originally ecdsa v={0,1}.
	// TODO:call `RegisterOperator` function from AVS.
	return nil, nil
}

func BuildELChainWriter(
	delegationManagerAddr common.Address, avsDirectoryAddr common.Address,
	ethClient backend.Backend, logger logging.Logger, metrics metrics.Metrics, txMgr backend.TxManager,
) (*sdkElClient.ELChainWriter, error) {
	return sdkElClient.BuildELChainWriter(
		delegationManagerAddr, avsDirectoryAddr,
		ethClient, logger, metrics, txMgr,
	)
}

func RegisterOperatorWithEL(ctx context.Context, elWriter sdkElClient.ELWriter, op sdkTypes.Operator) error {
	_, err := elWriter.RegisterAsOperator(ctx, op)
	if err != nil {
		return err
	}
	return nil
}

func DepositIntoStrategy(
	ctx context.Context, logger logging.Logger, elWriter sdkElClient.ELWriter,
	strategy common.Address, amount sdkTypes.StakeAmount,
) error {
	_, err := elWriter.DepositERC20IntoStrategy(ctx, strategy, amount)
	if err != nil {
		return err
	}
	return nil
}
