// Package tenant provides multi-tenancy support for the application.
package tenant

import (
	"context"
	"strconv"

	"connectrpc.com/connect"
	"github.com/kiliandbigblue/octoback/internal/x/auth0"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// contextKey is used to store tenant ID in context.
type contextKey string

const tenantIDKey contextKey = "tenant_id"

// TenantResolver is an interface for resolving tenant information for a user.
type TenantResolver interface {
	// GetDefaultTenantID returns the default tenant ID for a user.
	// If the user is not a member of any tenant, it returns 0 and false.
	GetDefaultTenantID(ctx context.Context, userID string) (int64, bool, error)
	// IsTenantMember checks if a user is a member of a specific tenant.
	IsTenantMember(ctx context.Context, tenantID int64, userID string) (bool, error)
}

// Config holds the tenant middleware configuration.
type Config struct {
	// HeaderName is the name of the header that can override the default tenant.
	// If empty, defaults to "X-Tenant-ID".
	HeaderName string
}

// Middleware provides tenant-aware middleware for Connect-RPC services.
type Middleware struct {
	resolver TenantResolver
	config   Config
	log      *zap.Logger
}

// NewMiddleware creates a new tenant middleware.
func NewMiddleware(resolver TenantResolver, config Config, log *zap.Logger) *Middleware {
	if config.HeaderName == "" {
		config.HeaderName = "X-Tenant-ID"
	}
	return &Middleware{
		resolver: resolver,
		config:   config,
		log:      log,
	}
}

// WithTenantID adds a tenant ID to the context.
func WithTenantID(ctx context.Context, tenantID int64) context.Context {
	return context.WithValue(ctx, tenantIDKey, tenantID)
}

// TenantIDFromContext retrieves the tenant ID from the context.
// Returns the tenant ID and true if found, or 0 and false if not.
func TenantIDFromContext(ctx context.Context) (int64, bool) {
	id, ok := ctx.Value(tenantIDKey).(int64)
	return id, ok
}

// MustTenantIDFromContext retrieves the tenant ID from the context or panics.
func MustTenantIDFromContext(ctx context.Context) int64 {
	id, ok := TenantIDFromContext(ctx)
	if !ok {
		panic("tenant ID not found in context")
	}
	return id
}

// NewTenantInterceptor creates a Connect interceptor that resolves and validates tenant access.
func (m *Middleware) NewTenantInterceptor() connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			// Get the authenticated user from context
			claims, ok := auth0.ClaimsFromContext(ctx)
			if !ok {
				return nil, connect.NewError(
					connect.CodeUnauthenticated,
					errors.New("no authenticated user found"),
				)
			}

			userID := claims.Subject
			var tenantID int64

			// Check if a specific tenant is requested via header
			tenantHeader := req.Header().Get(m.config.HeaderName)
			if tenantHeader != "" {
				// Parse the tenant ID from the header
				requestedTenantID, err := strconv.ParseInt(tenantHeader, 10, 64)
				if err != nil {
					return nil, connect.NewError(
						connect.CodeInvalidArgument,
						errors.New("invalid tenant ID format"),
					)
				}

				// Verify the user has access to this tenant
				isMember, err := m.resolver.IsTenantMember(ctx, requestedTenantID, userID)
				if err != nil {
					m.log.Error("failed to check tenant membership",
						zap.String("user_id", userID),
						zap.Int64("tenant_id", requestedTenantID),
						zap.Error(err),
					)
					return nil, connect.NewError(
						connect.CodeInternal,
						errors.New("failed to verify tenant access"),
					)
				}
				if !isMember {
					return nil, connect.NewError(
						connect.CodePermissionDenied,
						errors.New("user is not a member of the requested tenant"),
					)
				}
				tenantID = requestedTenantID
			} else {
				// Use the user's default tenant
				defaultTenantID, found, err := m.resolver.GetDefaultTenantID(ctx, userID)
				if err != nil {
					m.log.Error("failed to get default tenant",
						zap.String("user_id", userID),
						zap.Error(err),
					)
					return nil, connect.NewError(
						connect.CodeInternal,
						errors.New("failed to resolve tenant"),
					)
				}
				if !found {
					return nil, connect.NewError(
						connect.CodeFailedPrecondition,
						errors.New("user is not a member of any tenant"),
					)
				}
				tenantID = defaultTenantID
			}

			m.log.Debug("tenant resolved",
				zap.String("user_id", userID),
				zap.Int64("tenant_id", tenantID),
			)

			ctx = WithTenantID(ctx, tenantID)
			return next(ctx, req)
		}
	}
}
