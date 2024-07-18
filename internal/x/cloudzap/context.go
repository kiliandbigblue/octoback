package cloudzap

import (
	"context"

	"go.uber.org/zap"
)

type loggerContextKey struct{}

// WithLogger adds a cloudzap to the current context.
func WithLogger(ctx context.Context, cloudzap *zap.Logger) context.Context {
	return context.WithValue(ctx, loggerContextKey{}, cloudzap)
}

// GetLogger returns the cloudzap for the current context.
func GetLogger(ctx context.Context) (*zap.Logger, bool) {
	cloudzap, ok := ctx.Value(loggerContextKey{}).(*zap.Logger)
	return cloudzap, ok
}
