package v1

import (
	"context"
	"database/sql"
	"errors"

	"connectrpc.com/connect"
	models "github.com/kiliandbigblue/octoback/gen/proto/go/octoback/tenants/v1"
	store "github.com/kiliandbigblue/octoback/gen/store/v1"
	"github.com/kiliandbigblue/octoback/internal/x/auth0"
	"github.com/kiliandbigblue/octoback/internal/x/cloudzap"
)

type Service struct {
	querier store.Querier
}

func NewService(querier store.Querier) *Service {
	return &Service{
		querier: querier,
	}
}

// userIDFromContext extracts the user ID from the auth0 claims in context.
func userIDFromContext(ctx context.Context) (string, error) {
	claims, ok := auth0.ClaimsFromContext(ctx)
	if !ok {
		return "", errors.New("no authenticated user found")
	}
	return claims.Subject, nil
}

// ListTenants lists all available tenants.
func (s *Service) ListTenants(ctx context.Context, _ *connect.Request[models.ListTenantsRequest]) (*connect.Response[models.ListTenantsResponse], error) {
	log, _ := cloudzap.GetLogger(ctx)
	log.Info("ListTenants")

	tenants, err := s.querier.ListTenants(ctx)
	if err != nil {
		panic(err)
	}

	result := make([]*models.Tenant, len(tenants))
	for i, t := range tenants {
		result[i] = translateTenant(t)
	}

	return connect.NewResponse(&models.ListTenantsResponse{
		Tenants: result,
	}), nil
}

// GetMyTenants lists tenants the current user is a member of.
func (s *Service) GetMyTenants(ctx context.Context, _ *connect.Request[models.GetMyTenantsRequest]) (*connect.Response[models.GetMyTenantsResponse], error) {
	log, _ := cloudzap.GetLogger(ctx)
	log.Info("GetMyTenants")

	userID, err := userIDFromContext(ctx)
	if err != nil {
		return nil, connect.NewError(connect.CodeUnauthenticated, err)
	}

	tenants, err := s.querier.GetTenantsByUserId(ctx, userID)
	if err != nil {
		panic(err)
	}

	tenantResults := make([]*models.Tenant, len(tenants))
	memberships := make([]*models.TenantMembership, len(tenants))

	for i, t := range tenants {
		tenantResults[i] = translateTenant(t)

		// Get the membership for this tenant
		member, err := s.querier.GetTenantMemberByUserAndTenant(ctx, store.GetTenantMemberByUserAndTenantParams{
			UserID:   userID,
			TenantID: t.ID,
		})
		if err != nil {
			panic(err)
		}
		memberships[i] = translateMembership(member)
	}

	return connect.NewResponse(&models.GetMyTenantsResponse{
		Tenants:     tenantResults,
		Memberships: memberships,
	}), nil
}

// JoinTenant adds the current user to a tenant.
func (s *Service) JoinTenant(ctx context.Context, r *connect.Request[models.JoinTenantRequest]) (*connect.Response[models.JoinTenantResponse], error) {
	log, _ := cloudzap.GetLogger(ctx)
	log.Info("JoinTenant")

	userID, err := userIDFromContext(ctx)
	if err != nil {
		return nil, connect.NewError(connect.CodeUnauthenticated, err)
	}

	// Verify the tenant exists
	if _, err := s.querier.GetTenant(ctx, r.Msg.GetTenantId()); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, connect.NewError(connect.CodeNotFound, errors.New("tenant not found"))
		}
		panic(err)
	}

	// Check if already a member
	isMember, err := s.querier.IsTenantMember(ctx, store.IsTenantMemberParams{
		TenantID: r.Msg.GetTenantId(),
		UserID:   userID,
	})
	if err != nil {
		panic(err)
	}
	if isMember {
		return nil, connect.NewError(connect.CodeAlreadyExists, errors.New("already a member of this tenant"))
	}

	// Create the membership
	member, err := s.querier.CreateTenantMember(ctx, store.CreateTenantMemberParams{
		TenantID: r.Msg.GetTenantId(),
		UserID:   userID,
		Role:     "member",
	})
	if err != nil {
		panic(err)
	}

	return connect.NewResponse(&models.JoinTenantResponse{
		Membership: translateMembership(member),
	}), nil
}

// LeaveTenant removes the current user from a tenant.
func (s *Service) LeaveTenant(ctx context.Context, r *connect.Request[models.LeaveTenantRequest]) (*connect.Response[models.LeaveTenantResponse], error) {
	log, _ := cloudzap.GetLogger(ctx)
	log.Info("LeaveTenant")

	userID, err := userIDFromContext(ctx)
	if err != nil {
		return nil, connect.NewError(connect.CodeUnauthenticated, err)
	}

	// Get the membership
	member, err := s.querier.GetTenantMemberByUserAndTenant(ctx, store.GetTenantMemberByUserAndTenantParams{
		UserID:   userID,
		TenantID: r.Msg.GetTenantId(),
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, connect.NewError(connect.CodeNotFound, errors.New("not a member of this tenant"))
		}
		panic(err)
	}

	// Delete the membership
	if _, err := s.querier.DeleteTenantMember(ctx, member.ID); err != nil {
		panic(err)
	}

	return connect.NewResponse(&models.LeaveTenantResponse{}), nil
}
