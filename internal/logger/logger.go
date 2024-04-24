package logger

import (
	"go.uber.org/zap"
)

type LoggerConfig struct {
	Debug bool
}

func NewLogger(cfg *LoggerConfig, options ...zap.Option) (*zap.Logger, error) {
	mergedOptions := []zap.Option{
		zap.WithCaller(true),
	}
	copy(mergedOptions, options)

	if cfg.Debug == true {
		return zap.NewDevelopment(options...)
	} else {
		return zap.NewProduction(options...)
	}
}
