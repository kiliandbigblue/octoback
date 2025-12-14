package tenant

import (
	"context"
	"database/sql"
	"errors"

	store "github.com/kiliandbigblue/octoback/gen/store/v1"
)

// DBResolver implements TenantResolver using the database.
type DBResolver struct {
	querier store.Querier
}

// NewDBResolver creates a new DBResolver.
func NewDBResolver(querier store.Querier) *DBResolver {
	return &DBResolver{
		querier: querier,
	}
}

// GetDefaultTenantID returns the default tenant ID for a user.
func (r *DBResolver) GetDefaultTenantID(ctx context.Context, userID string) (int64, bool, error) {
	tenant, err := r.querier.GetDefaultTenantByUserId(ctx, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, false, nil
		}
		return 0, false, err
	}
	return tenant.ID, true, nil
}

// IsTenantMember checks if a user is a member of a specific tenant.
func (r *DBResolver) IsTenantMember(ctx context.Context, tenantID int64, userID string) (bool, error) {
	return r.querier.IsTenantMember(ctx, store.IsTenantMemberParams{
		TenantID: tenantID,
		UserID:   userID,
	})
}
