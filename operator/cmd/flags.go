package cmd

import "github.com/spf13/cobra"

// custom AVS addr, else default
func bindAVSAddress(cmd *cobra.Command, addr *string) {
	cmd.Flags().StringVar(addr, "avs-address", *addr, "Optional address of Skate AVS")
}

// custom rpcURL
func bindRPCURL(cmd *cobra.Command, rpcURL *string) {
	const flagRPCURL = "rpc-url"
	cmd.Flags().StringVar(rpcURL, flagRPCURL, *rpcURL, "RPC endpoint")
	_ = cmd.MarkFlagRequired(flagRPCURL)
}
