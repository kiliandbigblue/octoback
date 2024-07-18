package cloudzap

import (
	"context"

	"connectrpc.com/connect"
	"github.com/segmentio/ksuid"
	"go.uber.org/zap"
)

func NewLoggerInterceptor(log *zap.Logger) connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {
			log := log.With(zap.String("corID", ksuid.New().String()))
			ctx = WithLogger(ctx, log)
			return next(ctx, req)
		})
	}
	return connect.UnaryInterceptorFunc(interceptor)
}
