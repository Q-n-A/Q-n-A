package logger

import (
	"github.com/Q-n-A/Q-n-A/client"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// logger用設定
type Config struct {
	DevMode bool
}

// zap loggerを生成
func NewZapLogger(c *Config, cli client.BotClient) (*zap.Logger, error) {
	// ログレベルを設定
	logLevel := zap.DebugLevel

	// configの生成
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

	// traQログ投稿フックの生成
	hook := newTraQHook(cli)
	hookOpt := zap.Hooks(hook.Fire)

	// Loggerの生成
	log, err := config.Build(hookOpt)
	if err != nil {
		return nil, err
	}

	return log, nil
}
