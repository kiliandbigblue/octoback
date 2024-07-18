package main

import (
	"net/http"
	"os"

	"connectrpc.com/connect"
	"connectrpc.com/validate"
	"github.com/kiliandbigblue/octoback/gen/proto/go/octoback/core/v1/corev1connect"
	corev1 "github.com/kiliandbigblue/octoback/internal/core/v1"
	"github.com/kiliandbigblue/octoback/internal/x/cloudzap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	log, err := cloudzap.NewLogger(cloudzap.LoggerConfig{
		Development: false,
		Level:       zapcore.DebugLevel,
	})
	if err != nil {
		panic(err)
	}

	log.Info("starting server")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Info("defaulting to port", zap.String("port", port))
	}

	db, _ := os.OpenFile("database.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0o666) //nolint:gosec //Permissive permissions.

	fs, err := corev1.NewFileSystemGroceryStore(db)
	if err != nil {
		log.Fatal("failed to create file system grocery store", zap.Error(err))
	}

	cs := corev1.NewService(fs)

	vi, err := validate.NewInterceptor()
	if err != nil {
		log.Fatal("failed to create interceptor", zap.Error(err))
	}

	li := cloudzap.NewLoggerInterceptor(log)

	mux := http.NewServeMux()
	path, handler := corev1connect.NewServiceHandler(cs, connect.WithInterceptors(vi, li))
	mux.Handle(path, handler)

	//nolint:gosec //No timeout.
	err = http.ListenAndServe(
		"0.0.0.0:"+port,
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
	if err != nil {
		log.Fatal("server crashed", zap.Error(err))
	}
	log.Info("server stopped")
}
