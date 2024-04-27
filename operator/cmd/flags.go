package cmd

import "github.com/spf13/cobra"

func bindRegisterConfig(cobraCmd *cobra.Command, cfg *RegisterConfig) {
	bindAVSAddress(cobraCmd, &cfg.AVSAddr)

	const flagConfig = "config-file"
	cobraCmd.Flags().StringVar(&cfg.ConfigFile, flagConfig, cfg.ConfigFile, "Path to the Eigen-Layer yaml configuration file")
	_ = cobraCmd.MarkFlagRequired(flagConfig)
}

// NOTE: for futures use
func bindAVSAddress(cmd *cobra.Command, addr *string) {
	cmd.Flags().StringVar(addr, "avs-address", *addr, "Optional address of Skate AVS")
}

// NOTE: for futures use
func bindRPCURL(cmd *cobra.Command, rpcURL *string) {
	const flagRPCURL = "rpc-url"
	cmd.Flags().StringVar(rpcURL, flagRPCURL, *rpcURL, "RPC endpoint")
	_ = cmd.MarkFlagRequired(flagRPCURL)
}
