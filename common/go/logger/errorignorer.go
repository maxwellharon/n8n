package logger

import (
	"errors"
	"go.uber.org/zap/zapcore"
)

var ErrUnknownEncoder = errors.New("unknown encoder")

const ErrorIgnorerEncoderName = "errorignorer"

type errorIgnorer struct {
	zapcore.Core
	ignoredErrors []error
}

// NewErrorIgnorerCore creates an core that does not log certain errors.
// Half shamelessly copied and pasted from https://github.com/uber-go/zap/pull/475
func NewErrorIgnorerCore(baseCore zapcore.Core, ignoredErrors []error) zapcore.Core {
	return &errorIgnorer{
		Core:          baseCore,
		ignoredErrors: ignoredErrors,
	}
}

// With implements the With method (part of the zapcore.Core interface) by calling the With of the underlying core.
// We make sure that the ignored errors are cloned instead of copying the slice reference.
// We still don't deep copy the errors, since what we actually need is their address, as these are sentinel errors.
func (core *errorIgnorer) With(fields []zapcore.Field) zapcore.Core {
	return NewErrorIgnorerCore(core.Core.With(fields), append([]error(nil), core.ignoredErrors...))
}

// Check implements the Check method (part of the zapcore.Core interface) by calling the Check of the underlying core.
func (core *errorIgnorer) Check(ent zapcore.Entry, ce *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	if checked := core.Core.Check(ent, nil); checked == nil {
		return ce
	}

	return ce.AddCore(ent, core)
}

// Write returns nil when the log line should be ignored. Otherwise it calls the Write method of the underlying core.
func (core *errorIgnorer) Write(e zapcore.Entry, fields []zapcore.Field) error {
	for _, field := range fields {
		if field.Type == zapcore.ErrorType {
			err, ok := field.Interface.(error)
			if ok && in(err, core.ignoredErrors) {
				return nil
			}
		}
	}

	return core.Core.Write(e, fields)
}

// Sync implements the Sync method (part of the zapcore.Core interface) by calling the Sync of the underlying core.
func (core *errorIgnorer) Sync() error {
	return core.Core.Sync()
}

func in(target error, errs []error) bool {
	for _, e := range errs {
		if errors.Is(target, e) {
			return true
		}
	}

	return false
}
