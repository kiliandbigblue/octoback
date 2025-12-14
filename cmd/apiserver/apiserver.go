package main

import (
	"context"
	"database/sql"
	"flag"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/pkg/errors"
	_ "modernc.org/sqlite"

	"connectrpc.com/connect"
	"connectrpc.com/validate"
	"github.com/kiliandbigblue/octoback/gen/proto/go/octoback/groceries/v1/groceriesv1connect"
	"github.com/kiliandbigblue/octoback/gen/proto/go/octoback/recipes/v1/recipesv1connect"
	"github.com/kiliandbigblue/octoback/gen/proto/go/octoback/tenants/v1/tenantsv1connect"
	store "github.com/kiliandbigblue/octoback/gen/store/v1"
	groceriesv1 "github.com/kiliandbigblue/octoback/internal/groceries/v1"
	recipesv1 "github.com/kiliandbigblue/octoback/internal/recipes/v1"
	tenantsv1 "github.com/kiliandbigblue/octoback/internal/tenants/v1"
	"github.com/kiliandbigblue/octoback/internal/x/auth0"
	"github.com/kiliandbigblue/octoback/internal/x/cloudzap"
	"github.com/kiliandbigblue/octoback/internal/x/tenant"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

// The config struct holds configuration for the server.
type config struct {
	port          string
	dbPath        string
	staticDir     string
	auth0Domain   string
	auth0Audience string
}

func main() {
	var cfg config
	flag.StringVar(&cfg.port, "port", os.Getenv("PORT"), "port to listen on")
	flag.StringVar(&cfg.dbPath, "db-path", os.Getenv("OCTOMEAL_DB_PATH"), "SQLite database file path")
	flag.StringVar(&cfg.staticDir, "static-dir", os.Getenv("STATIC_DIR"), "directory containing static files to serve")
	flag.StringVar(&cfg.auth0Domain, "auth0-domain", os.Getenv("AUTH0_DOMAIN"), "Auth0 domain (e.g., your-tenant.auth0.com)")
	flag.StringVar(&cfg.auth0Audience, "auth0-audience", os.Getenv("AUTH0_AUDIENCE"), "Auth0 API audience identifier")
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

	// Validate Auth0 configuration
	if cfg.auth0Domain == "" || cfg.auth0Audience == "" {
		log.Fatal("AUTH0_DOMAIN and AUTH0_AUDIENCE environment variables are required")
	}

	db, err := openDB(cfg)
	if err != nil {
		log.Fatal("failed to open database", zap.Error(err))
	}
	defer func() {
		_ = db.Close()
	}()
	log.Info("database connection established", zap.String("path", cfg.dbPath))

	// Initialize Auth0 validator
	auth0Validator, err := auth0.NewValidator(auth0.Config{
		Domain:   cfg.auth0Domain,
		Audience: cfg.auth0Audience,
	}, log)
	if err != nil {
		log.Fatal("failed to create Auth0 validator", zap.Error(err))
	}
	log.Info("Auth0 validator initialized",
		zap.String("domain", cfg.auth0Domain),
		zap.String("audience", cfg.auth0Audience),
	)

	querier := store.New(db)
	groceriesSvc := groceriesv1.NewService(querier)
	recipesSvc := recipesv1.NewService(querier)
	tenantsSvc := tenantsv1.NewService(querier)

	vi, err := validate.NewInterceptor()
	if err != nil {
		log.Fatal("failed to create interceptor", zap.Error(err))
	}

	li := cloudzap.NewLoggerInterceptor(log)
	ai := auth0.NewAuthInterceptor(auth0Validator)

	// Initialize tenant middleware
	tenantResolver := tenant.NewDBResolver(querier)
	tenantMiddleware := tenant.NewMiddleware(tenantResolver, tenant.Config{}, log)
	ti := tenantMiddleware.NewTenantInterceptor()

	// Interceptors are executed in order: auth first, then tenant resolution, then logging, then validation
	interceptors := connect.WithInterceptors(ai, ti, li, vi)

	mux := http.NewServeMux()
	groceriesPath, groceriesHandler := groceriesv1connect.NewServiceHandler(groceriesSvc, interceptors)
	mux.Handle(groceriesPath, groceriesHandler)
	recipesPath, recipesHandler := recipesv1connect.NewServiceHandler(recipesSvc, interceptors)
	mux.Handle(recipesPath, recipesHandler)

	// Tenant service uses only auth interceptor (no tenant context needed)
	tenantInterceptors := connect.WithInterceptors(ai, li, vi)
	tenantsPath, tenantsHandler := tenantsv1connect.NewServiceHandler(tenantsSvc, tenantInterceptors)
	mux.Handle(tenantsPath, tenantsHandler)

	// Serve static files if configured (for production deployment)
	if cfg.staticDir != "" {
		log.Info("serving static files", zap.String("dir", cfg.staticDir))
		mux.Handle("/", spaHandler(cfg.staticDir))
	}

	// Wrap with CORS middleware to allow Authorization header
	handler := auth0.CORSMiddleware(h2c.NewHandler(mux, &http2.Server{}))

	log.Info("starting server", zap.String("port", cfg.port))

	//nolint:gosec //No timeout.
	err = http.ListenAndServe("0.0.0.0:"+cfg.port, handler)
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

// spaHandler returns a handler that serves static files with SPA fallback.
// For any request that doesn't match a static file, it serves index.html.
func spaHandler(staticDir string) http.Handler {
	fileServer := http.FileServer(http.Dir(staticDir))

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Clean the path to prevent directory traversal
		path := filepath.Clean(r.URL.Path)
		if path == "/" {
			path = "/index.html"
		}

		// Check if the file exists
		fullPath := filepath.Join(staticDir, path)
		info, err := os.Stat(fullPath)
		if err != nil {
			if errors.Is(err, fs.ErrNotExist) {
				// File doesn't exist, serve index.html for SPA routing
				http.ServeFile(w, r, filepath.Join(staticDir, "index.html"))
				return
			}
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Don't serve directories directly
		if info.IsDir() {
			// Try to serve index.html from the directory
			indexPath := filepath.Join(fullPath, "index.html")
			if _, err := os.Stat(indexPath); err == nil {
				http.ServeFile(w, r, indexPath)
				return
			}
			// Fall back to SPA index.html
			http.ServeFile(w, r, filepath.Join(staticDir, "index.html"))
			return
		}

		// Set cache headers for static assets
		if strings.HasPrefix(path, "/assets/") {
			w.Header().Set("Cache-Control", "public, max-age=31536000, immutable")
		}

		fileServer.ServeHTTP(w, r)
	})
}
