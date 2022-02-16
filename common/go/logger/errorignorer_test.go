package logger

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
	"testing"
)

func Test_ErrorIgnorerCore(t *testing.T) {
	next, logs := observer.New(zapcore.ErrorLevel)

	skipError := errors.New("test error")

	core := NewErrorIgnorerCore(next, []error{skipError})

	zap.New(core).Error("nuked", zap.String("a", "b"), zap.String("b", "c"), zap.Error(skipError))

	require.Equal(t, 0, len(logs.All()))
}

func Test_ErrorIgnorerCore_MoreThanOne(t *testing.T) {
	next, logs := observer.New(zapcore.ErrorLevel)

	skipError := errors.New("test error")

	core := NewErrorIgnorerCore(next, []error{skipError})

	logger := zap.New(core)
	logger.Error("nuked", zap.String("a", "b"), zap.String("b", "c"), zap.Error(skipError))
	logger.Error("not nuked", zap.String("a", "b"), zap.String("b", "c"))

	require.Len(t, logs.All(), 1)
	fields := logs.All()[0].Context

	assert.Len(t, fields, 2)
}

func Test_ErrorIgnorerCore_Tee(t *testing.T) {
	core1, logs1 := observer.New(zapcore.ErrorLevel)
	core2, logs2 := observer.New(zapcore.ErrorLevel)

	skipError := errors.New("test error")

	core := zapcore.NewTee(NewErrorIgnorerCore(core1, []error{skipError}), NewErrorIgnorerCore(core2, []error{skipError}))

	logger := zap.New(core)
	logger.Error("nuked", zap.String("a", "b"), zap.String("b", "c"), zap.Error(skipError))
	logger.Error("not nuked", zap.String("a", "b"), zap.String("b", "c"))

	require.Len(t, logs1.All(), 1, "logs1:\n%v", prettyPrintLogs(logs1.All()))
	fields1 := logs1.All()[0].Context
	assert.Len(t, fields1, 2)

	require.Lenf(t, logs2.All(), 1, "logs2:\n%v", prettyPrintLogs(logs2.All()))
	fields2 := logs2.All()[0].Context
	assert.Len(t, fields2, 2)
}

func Test_ErrorIgnorerCore_Tee_DifferentLeveles(t *testing.T) {
	core1, logs1 := observer.New(zapcore.ErrorLevel)
	core2, logs2 := observer.New(zapcore.ErrorLevel)

	skipError := errors.New("test error")

	core := zapcore.NewTee(NewErrorIgnorerCore(core1, []error{skipError}), NewErrorIgnorerCore(core2, []error{skipError}))

	logger := zap.New(core)
	logger.Debug("nuked", zap.String("a", "b"), zap.String("b", "c"), zap.Error(skipError))
	logger.Debug("not nuked", zap.String("a", "b"), zap.String("b", "c"))

	require.Len(t, logs1.All(), 0, "logs1:\n%v", prettyPrintLogs(logs1.All()))

	require.Lenf(t, logs2.All(), 0, "logs2:\n%v", prettyPrintLogs(logs2.All()))
}

func Test_ErrorIgnorerCore_Tee_DifferentCoreMinLevel(t *testing.T) {
	core1, logs1 := observer.New(zapcore.DebugLevel)
	core2, logs2 := observer.New(zapcore.ErrorLevel)

	skipError := errors.New("test error")

	core := zapcore.NewTee(NewErrorIgnorerCore(core1, []error{skipError}), NewErrorIgnorerCore(core2, []error{skipError}))

	logger := zap.New(core)

	logger.Error("errSkipped", zap.String("a", "b"), zap.String("b", "c"), zap.Error(skipError))
	logger.Error("errValid", zap.String("a", "b"), zap.String("b", "c"))

	logger.Debug("dbgSkipped", zap.String("a", "b"), zap.String("b", "c"), zap.Error(skipError))
	logger.Debug("dbgValid", zap.String("a", "b"), zap.String("b", "c"))

	// logs1 should have 2 entries, `errValid` and `dbgValid`. Both `*Skipped` values shouldn't be present because
	// they contain an ignored error
	require.Lenf(t, logs1.All(), 2, "logs1:\n%v", prettyPrintLogs(logs1.All()))
	require.Contains(t, prettyPrintLogs(logs1.All()), "errValid")
	require.Contains(t, prettyPrintLogs(logs1.All()), "dbgValid")
	require.NotContains(t, prettyPrintLogs(logs1.All()), "errSkipped")
	require.NotContains(t, prettyPrintLogs(logs1.All()), "dbgSkipped")

	// logs2 should have only `errValid`. `dbg*` logs should be ignored due to min level
	// `errSkipped` should be skipped because it contains an ignored error
	require.Len(t, logs2.All(), 1, "logs2:\n%v", prettyPrintLogs(logs2.All()))
	require.Contains(t, prettyPrintLogs(logs2.All()), "errValid")
	require.NotContains(t, prettyPrintLogs(logs2.All()), "dbgValid")
	require.NotContains(t, prettyPrintLogs(logs2.All()), "errSkipped")
	require.NotContains(t, prettyPrintLogs(logs2.All()), "dbgSkipped")
}

func prettyPrintLogs(entries []observer.LoggedEntry) string {
	var prettyPrinted string

	for _, e := range entries {
		var fields string
		for _, f := range e.Context {
			fields = fmt.Sprintf("%v %+v", fields, f)
		}

		prettyPrinted = fmt.Sprintf("%v(Level: %v)\nFields:\n[%v]\n%v\n\n", prettyPrinted, e.Level, fields, e.Message)
	}

	return prettyPrinted
}
