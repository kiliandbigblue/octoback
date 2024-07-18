package cloudzap

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// LoggerConfig configures the application logger.
type LoggerConfig struct {
	// Development indicates if the logger should output human-readable output for development.
	Development bool
	// Level indicates which log level the logger should output at.
	Level zapcore.Level
}

// NewLogger creates a new Logger.
func NewLogger(config LoggerConfig) (*zap.Logger, error) {
	if config.Development {
		zapConfig := zap.NewDevelopmentConfig()
		zapConfig.EncoderConfig.EncodeLevel = zapcore.LowercaseColorLevelEncoder
		zapConfig.Level = zap.NewAtomicLevelAt(config.Level)
		logger, err := zapConfig.Build(
			zap.AddCaller(),
			zap.AddStacktrace(zap.FatalLevel), // add stacktraces manually where needed
		)
		if err != nil {
			return nil, fmt.Errorf("init logger: %w", err)
		}
		return logger, nil
	}
	zapConfig := zap.NewProductionConfig()
	zapConfig.EncoderConfig = NewEncoderConfig()
	zapConfig.Level = zap.NewAtomicLevelAt(config.Level)
	zapOptions := []zap.Option{
		zap.AddCaller(),
		zap.AddStacktrace(zap.FatalLevel), // add stacktraces manually where needed
	}
	logger, err := zapConfig.Build(zapOptions...)
	if err != nil {
		return nil, fmt.Errorf("init logger: %w", err)
	}
	return logger, nil
}
