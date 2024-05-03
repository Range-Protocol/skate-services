package cmd

import (
	"skatechain.org/lib/logging"
	"skatechain.org/lib/on-chain/backend"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func storePrivatekeyCmd() *cobra.Command {
	logger := logging.NewLoggerWithConsoleWriter()

	var privateKey string
	var passphrase string
	cmd := &cobra.Command{
		Use:   "store",
		Short: "Store encrypted private key with password guarded",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			if privateKey == "" {
				logger.Error("private key not provided")
				return errors.New("Private key not provided")
			}
			if passphrase == "" {
				logger.Error("Pass phrase not provided")
				return errors.New("Pass phrase not provided")
			}
			if len(passphrase) < 8 {
				logger.Error("Pass phrase should have at least 8 characters")
				return errors.New("Pass phrase should have at least 8 characters")
			}
			backend.DumpECDSAPrivateKeyToKeystore(privateKey, passphrase)
			logger.Info("Private key successfully encrypted in `keystore`")
			return nil
		},
	}

	cmd.Flags().StringVarP(&privateKey, "private-key", "k", "", "Private key")
	cmd.Flags().StringVarP(&passphrase, "passphrase", "p", "", "Passphrase")

	return cmd
}
