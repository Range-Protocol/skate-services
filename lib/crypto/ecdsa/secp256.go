package ecdsa

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/hex"
	"io"

	"github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/common"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
)

// Sign returns a signature from input data.
//
// Follow EVM standard: [R (32b) || S (32b) || V (1b)] and (V={27, 28}).
func Sign(hash []byte, key *ecdsa.PrivateKey) ([65]byte, error) {
	sig, err := ethcrypto.Sign(hash[:], key)
	if err != nil {
		return [65]byte(sig), err
	}
	// NOTE: COMPATIBILITY REQUIREMENTS: Adjust V from secp256k1 0/1 to Ethereum 27/28
	sig[64] += 27
	return [65]byte(sig), nil
}

// Recover the public key that sign a given hash
//
// WARNING: Not the ecrecover in solidity, i.e. hash IS NOT prefixed with "\x19Ethereum Signed Message\n32".
//
// NOTE: Equivalent to Openzeppelin's ECDSA.recover function: https://github.com/OpenZeppelin/openzeppelin-contracts/blob/master/contracts/utils/cryptography/ECDSA.sol
func EcRecover(hash []byte, sig [65]byte) (common.Address, error) {
	println("Hash", hex.EncodeToString(hash), "Sig", hex.EncodeToString(sig[:]))
	if v := sig[64]; v != 27 && v != 28 {
		return common.Address{}, errors.New("invalid recovery id (V) format, must be 27 or 28")
	}
	// NOTE: COMPATIBILITY REQUIREMENTS: Adjust V from Ethereum 27/28 to secp256k1 0/1
	sig[64] -= 27

	pubkey, err := ethcrypto.SigToPub(hash[:], sig[:])
	if err != nil {
		return common.Address{}, errors.Wrap(err, "Recover public key")
	}

	return ethcrypto.PubkeyToAddress(*pubkey), nil
}

// Verify ethereum signed message
func Verify(address common.Address, hash []byte, sig [65]byte) (bool, error) {
	actual, err := EcRecover(hash, sig)

	return actual == address, err
}

func PubkeyToAddress(publicKey ecdsa.PublicKey) common.Address {
  return ethcrypto.PubkeyToAddress(publicKey)
}

func Keccak256(data ...[]byte) []byte {
	return ethcrypto.Keccak256(data...)
}

func S256() elliptic.Curve {
	return secp256k1.S256()
}

func KeyGen(curve elliptic.Curve, rand io.Reader) (*ecdsa.PrivateKey, error) {
	return ecdsa.GenerateKey(curve, rand)
}
