package main

import (
	"context"
	"database/sql"
	"flag"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"

	"connectrpc.com/connect"
	"connectrpc.com/validate"
	"github.com/kiliandbigblue/octoback/gen/proto/go/octoback/groceries/v1/groceriesv1connect"
	v1 "github.com/kiliandbigblue/octoback/internal/groceries/v1"
	"github.com/kiliandbigblue/octoback/internal/groceries/v1/store"
	"github.com/kiliandbigblue/octoback/internal/x/cloudzap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

// The config struct holds configuration for the server.
type config struct {
	port string
	db   struct {
		dsn          string
		maxOpenConns int
		maxIdleConns int
		maxIdleTime  string
	}
}

func main() {
	var cfg config
	flag.StringVar(&cfg.port, "port", os.Getenv("PORT"), "port to listen on")
	flag.StringVar(&cfg.db.dsn, "db-dsn", os.Getenv("OCTOMEAL_DB_DSN"), "PostgreSQL DSN")
	flag.IntVar(&cfg.db.maxOpenConns, "db-max-open-conns", 25, "PostgreSQL max open connections")
	flag.IntVar(&cfg.db.maxIdleConns, "db-max-idle-conns", 25, "PostgreSQL max idle connections")
	flag.StringVar(&cfg.db.maxIdleTime, "db-max-idle-time", "15m", "PostgreSQL max connection idle time")
	flag.Parse()

	log, err := cloudzap.NewLogger(cloudzap.LoggerConfig{
		Development: false,
		Level:       zapcore.DebugLevel,
	})
	if err != nil {
		panic(err)
	}

	db, err := openDB(cfg)
	if err != nil {
		log.Fatal("failed to open database", zap.Error(err))
	}
	defer func() {
		_ = db.Close()
	}()
	log.Info("database connection pool established")

	cs := v1.NewService(store.NewGroceryStore(db))

	vi, err := validate.NewInterceptor()
	if err != nil {
		log.Fatal("failed to create interceptor", zap.Error(err))
	}

	li := cloudzap.NewLoggerInterceptor(log)

	mux := http.NewServeMux()
	path, handler := groceriesv1connect.NewServiceHandler(cs, connect.WithInterceptors(vi, li))
	mux.Handle(path, handler)

	log.Info("starting server", zap.String("port", cfg.port))

	//nolint:gosec //No timeout.
	err = http.ListenAndServe(
		"0.0.0.0:"+cfg.port,
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
	if err != nil {
		log.Fatal("server crashed", zap.Error(err))
	}
	log.Info("server stopped")
}

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open database")
	}

	db.SetMaxOpenConns(cfg.db.maxOpenConns)
	db.SetMaxIdleConns(cfg.db.maxIdleConns)

	duration, err := time.ParseDuration(cfg.db.maxIdleTime)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse max idle time")
	}
	db.SetConnMaxIdleTime(duration)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, errors.Wrap(err, "failed to ping database")
	}

	return db, nil
}
