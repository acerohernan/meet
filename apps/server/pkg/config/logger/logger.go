package logger

import (
	"github.com/acerohernan/meet/pkg/config"
	"go.uber.org/zap"
)

var defaultLogger *zap.SugaredLogger

func InitLogger(conf *config.LoggerConfig) error {
	level, err := zap.ParseAtomicLevel(conf.Level)

	if err != nil {
		return err
	}

	zapConfig := zap.Config{
		Level:            level,
		Development:      false,
		Encoding:         "console",
		EncoderConfig:    zap.NewDevelopmentEncoderConfig(),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}

	logger, err := zapConfig.Build()

	if err != nil {
		return err
	}

	defaultLogger = logger.Sugar().WithOptions(zap.AddCallerSkip(1))

	return nil
}

func Debugw(msg string, keysAndValues ...interface{}) {
	if defaultLogger == nil {
		return
	}

	defaultLogger.Debugw(msg, keysAndValues...)
}

func Infow(msg string, keysAndValues ...interface{}) {
	if defaultLogger == nil {
		return
	}

	defaultLogger.Infow(msg, keysAndValues...)
}

func Warnw(msg string, keysAndValues ...interface{}) {
	if defaultLogger == nil {
		return
	}

	defaultLogger.Warnw(msg, keysAndValues...)
}

func Errorw(msg string, err error, keysAndValues ...interface{}) {
	if defaultLogger == nil {
		return
	}

	if err != nil {
		keysAndValues = append(keysAndValues, "error", err)
	}

	defaultLogger.Errorw(msg, keysAndValues...)
}
