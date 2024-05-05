package buildinfo

import (
	"context"
	"runtime/debug"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// version of the skate repo and all binaries built from this git commit.
// This value is set by goreleaser at build-time and should be the git tag for official releases.
var version = "v0.1.0"

func Version() string {
	return version
}

// Instrument logs the version, git commit hash, and timestamp from the runtime build info.
func Instrument(ctx context.Context) {
	_, commit, timestamp := getBuildInfo()

	versionGauge.WithLabelValues(version).Set(1)
	commitGauge.WithLabelValues(commit).Set(1)

	ts, _ := time.Parse(time.RFC3339, timestamp)
	timestampGauge.Set(float64(ts.Unix()))
}

func BuildInfoCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version information of this binary",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			goversion, commit, timestamp := getBuildInfo()

			var sb strings.Builder
			_, _ = sb.WriteString("Package version: " + version)
			_, _ = sb.WriteString("\n")
			_, _ = sb.WriteString("Go version:      " + goversion)
			_, _ = sb.WriteString("\n")
			_, _ = sb.WriteString("Git Commit:      " + commit)
			_, _ = sb.WriteString("\n")
			_, _ = sb.WriteString("Git Timestamp:   " + timestamp)
			_, _ = sb.WriteString("\n")

			cmd.Printf(sb.String())
		},
	}
}

// getBuildInfo returns the git commit hash and timestamp from the runtime build info.
// TODO: vcs info not shown, to be fixed
func getBuildInfo() (goversion, hash string, timestamp string) {
	goversion, hash, timestamp = "█████", "█████", "█████"
	hashLen := 7

	info, ok := debug.ReadBuildInfo()
	if !ok {
		return goversion, hash, timestamp
	}
  goversion = info.GoVersion

	for _, s := range info.Settings {
		if s.Key == "vcs.revision" {
			if len(s.Value) < hashLen {
				hashLen = len(s.Value)
			}
			hash = s.Value[:hashLen]
		} else if s.Key == "vcs.time" {
			timestamp = s.Value
		}
	}

	return goversion, hash, timestamp
}
