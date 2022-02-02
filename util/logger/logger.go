package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// logger用設定
type Config struct {
	DevMode bool
}

// cmd層向けloggerを生成
// traQへのログ送信ナシ
func NewRootLogger(c *Config) (*zap.Logger, error) {
	// ログレベルを設定
	var logLevel zapcore.Level
	if c.DevMode {
		logLevel = zap.DebugLevel
	} else {
		logLevel = zap.ErrorLevel
	}

	// configを生成
	config := &zap.Config{
		Level:    zap.NewAtomicLevelAt(logLevel),
		Encoding: "console",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			EncodeTime:     zapcore.RFC3339TimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}
	if c.DevMode {
		config.Development = true
	}

	// Loggerの生成
	log, err := config.Build()
	if err != nil {
		return nil, err
	}

	return log, nil
}
