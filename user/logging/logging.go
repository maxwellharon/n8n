package logging

import (
	"fmt"

	"go.uber.org/zap"

	logger "github.com/crafted-systems/smartcollect-pro/common/go/logger/defaultlogger"
)

// Configure configures the logging for the server.
func Configure(logLevel string) error {
	l, err := logger.NewLogger(logLevel)
	if err != nil {
		return fmt.Errorf("creating new logger: %w", err)
	}

	zap.ReplaceGlobals(l)

	return nil
}
