//nolint:wrapcheck // Internal stuff.
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
		zap.WrapCore(func(core zapcore.Core) zapcore.Core {
			return sourceLocationCore{nextCore: core}
		}),
	}
	logger, err := zapConfig.Build(zapOptions...)
	if err != nil {
		return nil, fmt.Errorf("init logger: %w", err)
	}
	return logger, nil
}

type sourceLocationCore struct {
	nextCore zapcore.Core
}

func (c sourceLocationCore) Enabled(level zapcore.Level) bool {
	return c.nextCore.Enabled(level)
}

func (c sourceLocationCore) With(fields []zapcore.Field) zapcore.Core {
	return sourceLocationCore{
		nextCore: c.nextCore.With(fields),
	}
}

func (c sourceLocationCore) Sync() error {
	return c.nextCore.Sync()
}

// Check implements zapcore.Core.
func (c sourceLocationCore) Check(entry zapcore.Entry, checked *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	if !c.nextCore.Enabled(entry.Level) {
		return checked
	}
	return checked.AddCore(entry, c)
}

// Write implements zapcore.Core.
func (c sourceLocationCore) Write(entry zapcore.Entry, fields []zapcore.Field) error {
	if entry.Caller.Defined {
		fields = appendIfNotExists(fields, SourceLocationForCaller(entry.Caller))
	}
	return c.nextCore.Write(entry, fields)
}

func appendIfNotExists(fields []zapcore.Field, field zap.Field) []zapcore.Field {
	for _, existing := range fields {
		if existing.Key == field.Key {
			return fields
		}
	}
	return append(fields, field)
}
