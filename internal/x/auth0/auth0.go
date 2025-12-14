// Package auth0 provides Auth0 JWT validation for Connect-RPC services.
package auth0

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"connectrpc.com/connect"
	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// Config holds the Auth0 configuration.
type Config struct {
	// Domain is the Auth0 domain (e.g., "your-tenant.auth0.com")
	Domain string
	// Audience is the API identifier configured in Auth0
	Audience string
}

// CustomClaims contains custom data we want from the token.
type CustomClaims struct {
	Scope string `json:"scope"`
}

// Validate does nothing for this example, but we need it to satisfy the validator.CustomClaims interface.
func (c CustomClaims) Validate(_ context.Context) error {
	return nil
}

// Validator wraps the Auth0 JWT validator.
type Validator struct {
	validator *validator.Validator
	log       *zap.Logger
}

// NewValidator creates a new Auth0 JWT validator.
func NewValidator(cfg Config, log *zap.Logger) (*Validator, error) {
	issuerURL, err := url.Parse("https://" + cfg.Domain + "/")
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse issuer URL")
	}

	provider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)

	jwtValidator, err := validator.New(
		provider.KeyFunc,
		validator.RS256,
		issuerURL.String(),
		[]string{cfg.Audience},
		validator.WithCustomClaims(
			func() validator.CustomClaims {
				return &CustomClaims{}
			},
		),
		validator.WithAllowedClockSkew(time.Minute),
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create JWT validator")
	}

	return &Validator{
		validator: jwtValidator,
		log:       log,
	}, nil
}

// contextKey is used to store validated claims in context.
type contextKey string

const claimsKey contextKey = "auth0_claims"

// ValidatedClaims contains the validated token claims.
type ValidatedClaims struct {
	Subject          string
	CustomClaims     *CustomClaims
	RegisteredClaims *validator.RegisteredClaims
}

// ClaimsFromContext retrieves the validated claims from the context.
func ClaimsFromContext(ctx context.Context) (*ValidatedClaims, bool) {
	claims, ok := ctx.Value(claimsKey).(*ValidatedClaims)
	return claims, ok
}

// withClaims adds validated claims to the context.
func withClaims(ctx context.Context, claims *ValidatedClaims) context.Context {
	return context.WithValue(ctx, claimsKey, claims)
}

// extractBearerToken extracts the token from the Authorization header.
func extractBearerToken(authHeader string) (string, error) {
	if authHeader == "" {
		return "", errors.New("authorization header is required")
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || !strings.EqualFold(parts[0], "bearer") {
		return "", errors.New("authorization header format must be 'Bearer {token}'")
	}

	return parts[1], nil
}

// NewAuthInterceptor creates a Connect interceptor that validates Auth0 JWT tokens.
func NewAuthInterceptor(v *Validator) connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			authHeader := req.Header().Get("Authorization")

			token, err := extractBearerToken(authHeader)
			if err != nil {
				v.log.Debug("failed to extract bearer token", zap.Error(err))
				return nil, connect.NewError(connect.CodeUnauthenticated, err)
			}

			claims, err := v.validator.ValidateToken(ctx, token)
			if err != nil {
				v.log.Debug("failed to validate token", zap.Error(err))
				return nil, connect.NewError(
					connect.CodeUnauthenticated,
					errors.New("invalid or expired token"),
				)
			}

			validatedClaims, ok := claims.(*validator.ValidatedClaims)
			if !ok {
				return nil, connect.NewError(
					connect.CodeInternal,
					errors.New("failed to cast claims"),
				)
			}

			customClaims, _ := validatedClaims.CustomClaims.(*CustomClaims)

			ctx = withClaims(ctx, &ValidatedClaims{
				Subject:          validatedClaims.RegisteredClaims.Subject,
				CustomClaims:     customClaims,
				RegisteredClaims: &validatedClaims.RegisteredClaims,
			})

			v.log.Debug("token validated successfully",
				zap.String("subject", validatedClaims.RegisteredClaims.Subject),
			)

			return next(ctx, req)
		}
	}
}

// RequireScope creates an interceptor that checks for a specific scope in the token.
func RequireScope(requiredScope string) connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			claims, ok := ClaimsFromContext(ctx)
			if !ok {
				return nil, connect.NewError(
					connect.CodeUnauthenticated,
					errors.New("no claims found in context"),
				)
			}

			if claims.CustomClaims == nil {
				return nil, connect.NewError(
					connect.CodePermissionDenied,
					errors.New("no custom claims found"),
				)
			}

			scopes := strings.Split(claims.CustomClaims.Scope, " ")
			for _, scope := range scopes {
				if scope == requiredScope {
					return next(ctx, req)
				}
			}

			return nil, connect.NewError(
				connect.CodePermissionDenied,
				fmt.Errorf("missing required scope: %s", requiredScope),
			)
		}
	}
}

// CORSMiddleware adds CORS headers for Auth0 authentication.
// This should wrap your HTTP handler to allow the Authorization header.
func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, X-CSRF-Token, Connect-Protocol-Version, X-Tenant-ID")
		w.Header().Set("Access-Control-Expose-Headers", "Link")
		w.Header().Set("Access-Control-Max-Age", "300")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
