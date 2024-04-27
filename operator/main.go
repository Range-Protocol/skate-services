package main

import (
	"context"
	libcmd "skate.org/avs/lib/cmd"
	"skate.org/avs/lib/logging"
	clicmd "skate.org/avs/operator/cmd"

	figure "github.com/common-nighthawk/go-figure"
)

func main() {
	cmd := clicmd.New()

	fig := figure.NewFigure("Skate OPERATOR", "", true)
	cmd.SetHelpTemplate(fig.String() + "\n" + cmd.HelpTemplate())

	libcmd.SilenceErrUsage(cmd)

	// Create a new Logger instance
	logger := logging.NewLoggerWithConsoleWriter()

	err := cmd.ExecuteContext(context.Background())
	if err != nil {
    logger.Fatal("Fatal error occurred!", "error", err.Error())
	}
}
