package cmd

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

const homeFlag = "home"

// BindHomeFlag binds the home flag to the given flag set.
// This is generally only required for apps that require multiple config files or persist data to disk.
// Using this flag will result in the viper config directory to be updated from default "." to "<home>/config".
func BindHomeFlag(flags *pflag.FlagSet, homeDir *string) {
	flags.StringVar(homeDir, homeFlag, *homeDir, "The application home directory containing config and data")
}

func Redact(flag string) string {
	if strings.Contains(flag, "token") ||
		strings.Contains(flag, "password") ||
		strings.Contains(flag, "secret") ||
		strings.Contains(flag, "db") ||
		strings.Contains(flag, "header") ||
		strings.Contains(flag, "key") {
		return "█████"
	}

	return flag
}

// custom AVS addr, else default
func BindVerbose(cmd *cobra.Command, verbose *bool) {
	cmd.Flags().BoolVar(verbose, "verbose", *verbose, "Run with verbose logs")
}
