package ecdsa_test

import (
	"crypto/rand"
	"testing"

	"skatechain.org/lib/crypto/ecdsa"
	"github.com/stretchr/testify/assert"
)

func TestECDSA(t *testing.T) {
	// Generate a random private key on K256 curve
	privateKey, err := ecdsa.KeyGen(ecdsa.S256(), rand.Reader)
	if err != nil {
		t.Fatal("Error generating private key:", err)
	}

	// Create a msgHash to sign
  message := []byte("Hello World")
	msgHash := ecdsa.Keccak256(message)

	// Sign the hash
  signature, err := ecdsa.Sign(msgHash[:], privateKey)
	if err != nil {
		t.Fatal("Error signing:", err)
	}

	// Verify the recovered address matches the original address
	originalAddress := ecdsa.PubkeyToAddress(privateKey.PublicKey)
  valid, err := ecdsa.Verify(originalAddress, msgHash[:], signature)
	if err != nil {
		t.Fatal("Error verifying signature:", err)
	}
	assert.True(t, valid, "Signature should be valid")

	// Recover public key from signature
	recoveredAddress, err := ecdsa.EcRecover(msgHash, signature)
	if err != nil {
		t.Fatal("Error recovering public key:", err)
	}
	assert.Equal(t, originalAddress, recoveredAddress, "Recovered address should match original address")
}
