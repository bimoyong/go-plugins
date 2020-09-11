package zap

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/micro/go-micro/v2/logger"
)

type Options struct {
	logger.Options
}

type callerSkipKey struct{}

func WithCallerSkip(i int) logger.Option {
	return logger.SetOption(callerSkipKey{}, i)
}

type configKey struct{}

// WithConfig pass zap.Config to logger
func WithConfig(c zap.Config) logger.Option {
	return logger.SetOption(configKey{}, c)
}

type encoderConfigKey struct{}

// WithEncoderConfig pass zapcore.EncoderConfig to logger
func WithEncoderConfig(c zapcore.EncoderConfig) logger.Option {
	return logger.SetOption(encoderConfigKey{}, c)
}

type coreKey struct{}

// WithCore pass zap.Core to logger
func WithCore(c zapcore.Core) logger.Option {
	return logger.SetOption(coreKey{}, c)
}

// WithCores pass slice of zap.Core to logger
func WithCores(c ...zapcore.Core) logger.Option {
	return logger.SetOption(coreKey{}, zapcore.NewTee(c...))
}

type namespaceKey struct{}

func WithNamespace(namespace string) logger.Option {
	return logger.SetOption(namespaceKey{}, namespace)
}
