package backend

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"skatechain.org/lib/logging"
)

// WaitMined waits for tx to be mined on the blockchain.
// Returns the transaction receipt when the transaction is mined.
func WaitMined(ctx context.Context, backend *Backend, tx *types.Transaction) (*types.Receipt, error) {
	logger := logging.NewLoggerWithConsoleWriter()
	queryTicker := time.NewTicker(time.Second * 3)
	defer queryTicker.Stop()

	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-queryTicker.C:
			receipt, _ := backend.TransactionReceipt(ctx, tx.Hash())
			if receipt != nil {
				return receipt, nil
			}
			logger.Info("Wait for next block...", "transaction_hash", tx.Hash().Hex())
		}
	}
}
