package cmd

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// loggerを生成
func newLogger(c *Config) *zap.Logger {
	// EncoderConfigの生成
	var encoderConfig zapcore.EncoderConfig
	if c.DevMode {
		encoderConfig = zap.NewDevelopmentEncoderConfig()
	} else {
		encoderConfig = zap.NewProductionEncoderConfig()
	}

	// Coreの生成
	var core zapcore.Core
	if c.DevMode {
		core = zapcore.NewCore(
			zapcore.NewConsoleEncoder(encoderConfig),
			zapcore.AddSync(os.Stdout),
			zapcore.DebugLevel,
		)
	} else {
		core = zapcore.NewCore(
			zapcore.NewConsoleEncoder(encoderConfig),
			zapcore.AddSync(os.Stdout),
			zapcore.ErrorLevel,
		)
	}

	// Loggerの生成
	return zap.New(core)
}
