package main

// NOTE:
import (
	"context"

	clicmd "skatechain.org/kms/cmd"
	libcmd "skatechain.org/lib/cmd"
	"skatechain.org/lib/logging"

	figure "github.com/common-nighthawk/go-figure"
)

func main() {
	cmd := clicmd.New()

	fig := figure.NewFigure("Key Stonk", "", true)
	cmd.SetHelpTemplate(fig.String() + "\n" + cmd.HelpTemplate())

	libcmd.SilenceErrUsage(cmd)

	// Create a new Logger instance
	logger := logging.NewLoggerWithConsoleWriter()

	err := cmd.ExecuteContext(context.Background())
	if err != nil {
		logger.Fatal("Fatal error occurred!", "error", err.Error())
	}
}
