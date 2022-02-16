package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func stringToLevel(l string) (zap.AtomicLevel, error) {
	var level zapcore.Level

	err := level.UnmarshalText([]byte(l))
	if err != nil {
		return zap.NewAtomicLevel(), fmt.Errorf("unable to convert level %s to a zap level: %w", l, err)
	}

	return zap.NewAtomicLevelAt(level), nil
}

// NewDefaultConfig creates our default config: it uses the ProductionConfig as a base,
// but it disables sampling and uses "message" for the message key and "time" for the time key,
// as well as encoding time as ISO8691.
func NewDefaultConfig(level string) (zap.Config, error) {
	c := zap.NewProductionConfig()

	c.Sampling = nil // we normally don't use sampling
	c.EncoderConfig.MessageKey = "message"
	c.EncoderConfig.TimeKey = "time"
	c.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	var err error

	c.Level, err = stringToLevel(level)
	if err != nil {
		return c, fmt.Errorf("unable to create level: %w", err)
	}

	return c, err
}

// NewLogger creates a logger with the provided configuration and extra options.
func NewLogger(config zap.Config, options ...Option) (*zap.Logger, error) {
	var opts Options
	for _, o := range options {
		o(&opts)
	}

	logger, err := config.Build()
	if err != nil {
		return nil, fmt.Errorf("building logger: %w", err)
	}

	// this wraps individually every core
	if len(opts.ignoredErrors) > 0 {
		if len(opts.extraCores) > 0 {
			for i, core := range opts.extraCores {
				opts.extraCores[i] = NewErrorIgnorerCore(core, opts.ignoredErrors)
			}
		}

		logger = logger.WithOptions(zap.WrapCore(func(c zapcore.Core) zapcore.Core {
			return NewErrorIgnorerCore(c, opts.ignoredErrors)
		}))
	}

	if len(opts.extraCores) > 0 {
		newCores := append([]zapcore.Core{logger.Core()}, opts.extraCores...)

		logger = logger.WithOptions(zap.WrapCore(func(zapcore.Core) zapcore.Core {
			return zapcore.NewTee(
				newCores...,
			)
		}))
	}

	return logger, nil
}
