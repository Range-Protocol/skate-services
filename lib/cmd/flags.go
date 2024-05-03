package cmd

import (
	"strings"

	"github.com/spf13/cobra"
)

// NOTE: Redact sensitive log, e.g. authorization header, password, ...
// To be used in the futures
func Redact(input string) string {
	if strings.Contains(input, "token") ||
		strings.Contains(input, "password") ||
		strings.Contains(input, "secret") ||
		strings.Contains(input, "db") ||
		strings.Contains(input, "header") ||
		strings.Contains(input, "key") {
		return "█████"
	}

	return input
}

// custom AVS addr, else default
func BindVerbose(cmd *cobra.Command, verbose *bool) {
	cmd.Flags().BoolVar(verbose, "verbose", *verbose, "Run with verbose logs")
}
