package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerConfig struct {
	Debug bool
}

func NewLogger(cfg *LoggerConfig, options ...zap.Option) (*zap.Logger, error) {
	mergedOptions := []zap.Option{
		zap.WithCaller(true),
		zap.AddStacktrace(zapcore.ErrorLevel),
	}
	copy(mergedOptions, options)

	if cfg.Debug == true {
		return zap.NewDevelopment(options...)
	} else {
		return zap.NewProduction(options...)
	}
}
