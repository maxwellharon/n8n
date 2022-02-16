package logger

import "go.uber.org/zap/zapcore"

type Option func(*Options)

type Options struct {
	ignoredErrors []error
	extraCores    []zapcore.Core
}

func WithExtraCores(cores ...zapcore.Core) func(*Options) {
	return func(o *Options) {
		o.extraCores = cores
	}
}

func WithIgnoredErrors(errs ...error) func(*Options) {
	return func(o *Options) {
		o.ignoredErrors = errs
	}
}
