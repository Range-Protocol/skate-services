package deposit

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	bindingIERC20 "skatechain.org/contracts/bindings/IERC20"
	bindingIStrategyManager "skatechain.org/contracts/bindings/IStrategyManager"
	"skatechain.org/lib/cmd"
	"skatechain.org/lib/crypto/ecdsa"
	"skatechain.org/lib/logging"
	"skatechain.org/lib/on-chain/backend"
)

var logger = logging.NewLoggerWithConsoleWriter()

func DepositIntoStrategy(
	privateKey *ecdsa.PrivateKey,
	config cmd.EnvironmentConfig,
	strategy common.Address,
	token common.Address,
	amount *big.Int,
	autoApprove bool,
) error {
	be, err := backend.NewBackend(config.HttpRPC)
	if err != nil {
		return err
	}

	strategyManager := common.HexToAddress(config.StrategyManager)
	strategyManagerContract, err := bindingIStrategyManager.NewBindingIStrategyManager(
		strategyManager, be,
	)
	if err != nil {
		return err
	}

	tokenContract, err := bindingIERC20.NewBindingIERC20(token, be)
	allowance, err := tokenContract.Allowance(
		&bind.CallOpts{},
		ecdsa.PubkeyToAddress(privateKey.PublicKey),
		common.HexToAddress(config.StrategyManager),
	)
	if err != nil {
		return err
	}

	chainId := new(big.Int).SetUint64(config.MainChainId)
	transactor, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)

	if amount.Cmp(allowance) > 0 {
		if autoApprove {
			logger.Info("Not enough allowance", "amount", amount, "allowance", allowance, "strategy manager", config.StrategyManager)
			logger.Info("Increasing token approval allowance for strategy manager ...")
			tx, err := tokenContract.Approve(transactor, strategyManager, amount)
			if err != nil {
				logger.Error("Approve failed", "error", errors.Wrap(err, "ERC20.Approve"))
				return err
			}
			_, err = backend.WaitMined(context.Background(), &be, tx)
			logger.Info("Token approved")

		} else {
			logger.Fatal("Not enough allowance", "amount", amount, "allowance", allowance, "strategy manager", config.StrategyManager)
		}
	}

	// Dry run transaction to check for potential rejections
	transactorNoSend := *transactor
	transactorNoSend.NoSend = true
	_, err = strategyManagerContract.DepositIntoStrategy(
		&transactorNoSend,
		strategy,
		token,
		amount,
	)
	if err != nil {
		logger.Error(
			"Transaction simulation failed",
			"error", errors.Wrap(err, "StrategyManager.DepositIntoStrategy"))
		return err
	}

	tx, err := strategyManagerContract.DepositIntoStrategy(
		transactor,
		strategy,
		token,
		amount,
	)
	logger.Info("Transaction sent", "txHash", tx.Hash().Hex())
	logger.Info("Waiting for confirmation ...")

	receipt, err := backend.WaitMined(context.Background(), &be, tx)
	if err != nil {
		logger.Error("Failed to get transaction receipt", "error", errors.Wrap(err, "backend.TransactionReceipt"))
		return err
	}
	logger.Info("Transaction receipt", "receipt", receipt)

	return nil
}
