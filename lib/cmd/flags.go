package cmd

import (
	"net/url"
	"strings"

	"github.com/spf13/pflag"
)

const homeFlag = "home"

// BindHomeFlag binds the home flag to the given flag set.
// This is generally only required for apps that require multiple config files or persist data to disk.
// Using this flag will result in the viper config directory to be updated from default "." to "<home>/config".
func BindHomeFlag(flags *pflag.FlagSet, homeDir *string) {
	flags.StringVar(homeDir, homeFlag, *homeDir, "The application home directory containing config and data")
}

// redact returns a redacted version of the given flag value. It currently supports redacting
// passwords in valid URLs as well as flags that contains words like "token", "password", "secret", "db" or "key".
func redact(flag, val string) string {
	if strings.Contains(flag, "token") ||
		strings.Contains(flag, "password") ||
		strings.Contains(flag, "secret") ||
		strings.Contains(flag, "db") ||
		strings.Contains(flag, "header") ||
		strings.Contains(flag, "key") {
		return "xxxxx"
	}

	u, err := url.Parse(val)
	if err != nil {
		return val
	}

	return u.Redacted()
}
