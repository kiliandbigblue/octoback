package main

import (
	"context"
	"database/sql"
	"flag"
	"net/http"
	"os"
	"time"

	"github.com/pkg/errors"
	_ "modernc.org/sqlite"

	"connectrpc.com/connect"
	"connectrpc.com/validate"
	"github.com/kiliandbigblue/octoback/gen/proto/go/octoback/groceries/v1/groceriesv1connect"
	"github.com/kiliandbigblue/octoback/gen/proto/go/octoback/recipes/v1/recipesv1connect"
	store "github.com/kiliandbigblue/octoback/gen/store/v1"
	groceriesv1 "github.com/kiliandbigblue/octoback/internal/groceries/v1"
	recipesv1 "github.com/kiliandbigblue/octoback/internal/recipes/v1"
	"github.com/kiliandbigblue/octoback/internal/x/cloudzap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

// The config struct holds configuration for the server.
type config struct {
	port   string
	dbPath string
}

func main() {
	var cfg config
	flag.StringVar(&cfg.port, "port", os.Getenv("PORT"), "port to listen on")
	flag.StringVar(&cfg.dbPath, "db-path", os.Getenv("OCTOMEAL_DB_PATH"), "SQLite database file path")
	flag.Parse()

	if cfg.dbPath == "" {
		cfg.dbPath = "octoback.db"
	}

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
	log.Info("database connection established", zap.String("path", cfg.dbPath))

	querier := store.New(db)
	groceriesSvc := groceriesv1.NewService(querier)
	recipesSvc := recipesv1.NewService(querier)

	vi, err := validate.NewInterceptor()
	if err != nil {
		log.Fatal("failed to create interceptor", zap.Error(err))
	}

	li := cloudzap.NewLoggerInterceptor(log)

	mux := http.NewServeMux()
	groceriesPath, groceriesHandler := groceriesv1connect.NewServiceHandler(groceriesSvc, connect.WithInterceptors(vi, li))
	mux.Handle(groceriesPath, groceriesHandler)
	recipesPath, recipesHandler := recipesv1connect.NewServiceHandler(recipesSvc, connect.WithInterceptors(vi, li))
	mux.Handle(recipesPath, recipesHandler)

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
	db, err := sql.Open("sqlite", cfg.dbPath)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open database")
	}

	// Enable foreign keys for SQLite
	if _, err := db.Exec("PRAGMA foreign_keys = ON"); err != nil {
		return nil, errors.Wrap(err, "failed to enable foreign keys")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, errors.Wrap(err, "failed to ping database")
	}

	return db, nil
}
