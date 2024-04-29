package crypto

import (
	"github.com/pkg/errors"

	secp256k1 "github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcec/v2/ecdsa"
	"github.com/cometbft/cometbft/crypto"

	"github.com/ethereum/go-ethereum/common"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
)

// NOTE: ethereum standard signature, built-in with solidity `sign`, `verify`

// privKeyLen is the length of a secp256k1 private key.
const privKeyLen = 32

// pubkeyCompressedLen is the length of a secp256k1 compressed public key.
const pubkeyCompressedLen = 33

// pubkeyUncompressedLen is the length of a secp256k1 uncompressed public key.
const pubkeyUncompressedLen = 65

// Sign returns a signature from input data.
//
// The produced signature is 65 bytes in the [R || S || V] format where V is 27 or 28.
func Sign(key crypto.PrivKey, input [32]byte) ([65]byte, error) {
	keyBytes := key.Bytes()
	if len(keyBytes) != privKeyLen {
		return [65]byte{}, errors.New("invalid private key length")
	}

	privKey, _ := secp256k1.PrivKeyFromBytes(keyBytes)

	sig, err := ecdsa.SignCompact(privKey, input[:], false)
	if err != nil {
		return [65]byte(sig), errors.Wrap(err, "recover public key")
	}

	// Convert signature from "compact" into "Ethereum R S V" format.
	return [65]byte(append(sig[1:], sig[0])), nil
}

// Note the signature MUST be 65 bytes in the Ethereum [R || S || V] format. Ethereum V={27, 28}
func Verify(address common.Address, hash [32]byte, sig [65]byte) (bool, error) {
	// Adjust V from Ethereum 27/28 to secp256k1 0/1
	const vIdx = 64
	if v := sig[vIdx]; v != 27 && v != 28 {
		return false, errors.New("invalid recovery id (V) format, must be 27 or 28")
	}
	sig[vIdx] -= 27

	pubkey, err := ethcrypto.SigToPub(hash[:], sig[:])
	if err != nil {
		return false, errors.Wrap(err, "recover public key")
	}

	actual := ethcrypto.PubkeyToAddress(*pubkey)

	return actual == address, nil
}
