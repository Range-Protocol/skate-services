package logging_test

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/pkg/errors"
	"skate.org/avs/lib/logging"
	"skate.org/avs/lib/tutil"
)

//go:generate go test . -golden -clean

func TestSimpleLogs(t *testing.T) {
	t.Parallel()

	AssertLogging(t, func(t *testing.T, logger logging.Logger) {
		t.Helper()

		{ // Test some basic logging
			logger.Info("info message", "with", "args")
			logger.Debug("debug this code for me please", "number", 1)
			logger.Warn("watch out!", os.ErrExist)
			logger.Error("something went wrong", io.EOF, "float", 1.234)
		}
		{ // Test errors wrapping and logging
			err := errors.New("first")
			logger.Warn("err1", err)
			err = errors.Wrap(err, "second")
			logger.Error("err2", err)
		}
		{ // Test attributes in context
			loggerCtx1 := logger.With("ctx_key1", "ctx_value1")
			loggerCtx1.Debug("ctx1 debug message", "debug_key1", "debug_value1")
			loggerCtx2 := logger.With("ctx_key2", "ctx_value2")
			loggerCtx2.Info("ctx2 info message", "info_key2", "info_value2")
		}
		{ // Test wrapping errors in pkg errors
			err := errors.New("new")
			err = errors.Wrap(err, "wrap")
			logger.Warn("Pkg wrapped error", err)
		}
	})
}

// AssertLogging returns a function that will assert all loggers' output against
// golden test files.
func AssertLogging(t *testing.T, testFunc func(*testing.T, logging.Logger)) {
	t.Helper()

	// For simplicity, we're only using the logger in this test, so we directly create it here.
	var buf bytes.Buffer
	logger := logging.NewLogger(&buf) // Using a buffer for testing

	testFunc(t, *logger)

	tutil.RequireGoldenBytes(t, buf.Bytes())
}
