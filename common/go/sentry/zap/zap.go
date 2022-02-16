package zap

import (
	"github.com/TheZeroSlave/zapsentry"
	"go.uber.org/zap/zapcore"
)

func DefaultConfiguration() zapsentry.Configuration {
	return zapsentry.Configuration{
		Level: zapcore.ErrorLevel,
	}
}
