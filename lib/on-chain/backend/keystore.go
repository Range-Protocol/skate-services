package backend

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
	"skatechain.org/lib/logging"
)

func DumpECDSAPrivateKeyToKeystore(privateKeyHex string, path string, passphrase string) {
	logger := logging.NewLoggerWithConsoleWriter()
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		logger.Fatal("Invalid private key", "err", errors.Wrap(err, "crpyo.HexToECDSA"))
	}

	// Create a new keystore directory
	ks := keystore.NewKeyStore(path, keystore.StandardScryptN, keystore.StandardScryptP)

	// Import the private key into the keystore, encrypting it with the passphrase
	account, err := ks.ImportECDSA(privateKey, passphrase)
	if err != nil {
		logger.Fatal("Import ECDSA key failed", "err", errors.Wrap(err, "keystore.ImportECDSA"))
	}

	// Delete the key in-memory if you imported an existing key
	err = ks.Delete(account, passphrase) // Comment this line if you want to keep the key in keystore
	if err != nil {
		logger.Fatal("Could not delete keystore file", "err", errors.Wrap(err, "keystore.Delete"))
	}

	logger.Info("Keystore created for account:", "account", account.Address.Hex())
}

func LoadTransactorFromKeystore(address common.Address, passphrase string, chainId *big.Int) (*bind.TransactOpts, error) {
	ks := keystore.NewKeyStore("keys", keystore.StandardScryptN, keystore.StandardScryptP)
	account := accounts.Account{Address: address}
	// keyJSON, err := ks.Export(account, passphrase, passphrase)
	ks.TimedUnlock(account, passphrase, time.Minute*0)
	return bind.NewKeyStoreTransactorWithChainID(ks, account, chainId)
}
