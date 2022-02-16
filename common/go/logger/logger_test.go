package logger

import (
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"testing"
)

func Test_stringToLevel(t *testing.T) {
	type args struct {
		l string
	}
	tests := []struct {
		name    string
		args    args
		want    zapcore.Level
		wantErr bool
	}{
		{
			name: "not passing level, defaults to info",
			args: args{
				l: "",
			},
			want:    zapcore.InfoLevel,
			wantErr: false,
		},
		{
			name: "unrecognized level",
			args: args{
				l: "asdadsad",
			},
			want:    zapcore.InfoLevel,
			wantErr: true,
		},
		{
			name: "debug",
			args: args{
				l: "debug",
			},
			want:    zapcore.DebugLevel,
			wantErr: false,
		},
		{
			name: "error",
			args: args{
				l: "error",
			},
			want:    zapcore.ErrorLevel,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		testCase := tt
		t.Run(testCase.name, func(t *testing.T) {
			level, err := stringToLevel(testCase.args.l)
			assert.Truef(t, (err != nil) == testCase.wantErr, "wanted err to be %v, but err is %v", testCase.wantErr, err)
			assert.Equal(t, testCase.want, level.Level())
		})
	}
}

func TestNewDefaultConfig(t *testing.T) {
	type args struct {
		level string
	}
	type want struct {
		messageKey       string
		timeKey          string
		level            zapcore.Level
		sampling         *zap.SamplingConfig
		outputPaths      []string
		errorOutputPaths []string
	}
	tests := []struct {
		name    string
		args    args
		want    want
		wantErr bool
	}{
		{
			name: "defaults created appropriately",
			args: args{
				level: "error",
			},
			want: want{
				messageKey:       "message",
				timeKey:          "time",
				level:            zapcore.ErrorLevel,
				sampling:         nil,
				outputPaths:      []string{"stderr"},
				errorOutputPaths: []string{"stderr"},
			},
			wantErr: false,
		},
		{
			name: "error occurs",
			args: args{
				level: "some weird level",
			},
			want: want{
				messageKey:       "message",
				timeKey:          "time",
				level:            zapcore.InfoLevel,
				sampling:         nil,
				outputPaths:      []string{"stderr"},
				errorOutputPaths: []string{"stderr"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		testCase := tt
		t.Run(testCase.name, func(t *testing.T) {
			config, err := NewDefaultConfig(testCase.args.level)
			assert.Truef(t, (err != nil) == testCase.wantErr, "wanted err to be %v, but err is %v", testCase.wantErr, err)
			assert.Equal(t, testCase.want.messageKey, config.EncoderConfig.MessageKey)
			assert.Equal(t, testCase.want.timeKey, config.EncoderConfig.TimeKey)
			assert.Equal(t, testCase.want.level, config.Level.Level())
			assert.Equal(t, testCase.want.sampling, config.Sampling)
			assert.ElementsMatch(t, testCase.want.outputPaths, config.OutputPaths)
			assert.ElementsMatch(t, testCase.want.errorOutputPaths, config.ErrorOutputPaths)
		})
	}
}
