package v1

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"errors"
	"time"

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

// generateInviteCode generates a random invite code.
func generateInviteCode() (string, error) {
	b := make([]byte, 12)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b)[:16], nil
}

// EnsureUserTenant ensures the current user has at least one tenant.
// If the user has no tenant, a personal tenant is created.
func (s *Service) EnsureUserTenant(ctx context.Context, _ *connect.Request[models.EnsureUserTenantRequest]) (*connect.Response[models.EnsureUserTenantResponse], error) {
	log, _ := cloudzap.GetLogger(ctx)
	log.Info("EnsureUserTenant")

	userID, err := userIDFromContext(ctx)
	if err != nil {
		return nil, connect.NewError(connect.CodeUnauthenticated, err)
	}

	// Check if user already has tenants
	count, err := s.querier.CountUserTenants(ctx, userID)
	if err != nil {
		panic(err)
	}

	if count > 0 {
		// User already has at least one tenant, return the first one
		tenant, err := s.querier.GetDefaultTenantByUserId(ctx, userID)
		if err != nil {
			panic(err)
		}
		member, err := s.querier.GetTenantMemberByUserAndTenant(ctx, store.GetTenantMemberByUserAndTenantParams{
			UserID:   userID,
			TenantID: tenant.ID,
		})
		if err != nil {
			panic(err)
		}
		return connect.NewResponse(&models.EnsureUserTenantResponse{
			Tenant:     translateTenant(tenant),
			Membership: translateMembership(member),
			Created:    false,
		}), nil
	}

	// Create a new tenant for the user
	tenant, err := s.querier.CreateTenant(ctx, "Mon foyer")
	if err != nil {
		panic(err)
	}

	// Make the user the owner
	member, err := s.querier.CreateTenantMember(ctx, store.CreateTenantMemberParams{
		TenantID: tenant.ID,
		UserID:   userID,
		Role:     "owner",
	})
	if err != nil {
		panic(err)
	}

	return connect.NewResponse(&models.EnsureUserTenantResponse{
		Tenant:     translateTenant(tenant),
		Membership: translateMembership(member),
		Created:    true,
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

	// Prevent owner from leaving (they should transfer ownership first or delete the tenant)
	if member.Role == "owner" {
		return nil, connect.NewError(connect.CodeFailedPrecondition, errors.New("owner cannot leave tenant, transfer ownership first"))
	}

	// Delete the membership
	if _, err := s.querier.DeleteTenantMember(ctx, member.ID); err != nil {
		panic(err)
	}

	return connect.NewResponse(&models.LeaveTenantResponse{}), nil
}

// CreateInvitation creates an invitation code for a tenant.
func (s *Service) CreateInvitation(ctx context.Context, r *connect.Request[models.CreateInvitationRequest]) (*connect.Response[models.CreateInvitationResponse], error) {
	log, _ := cloudzap.GetLogger(ctx)
	log.Info("CreateInvitation")

	userID, err := userIDFromContext(ctx)
	if err != nil {
		return nil, connect.NewError(connect.CodeUnauthenticated, err)
	}

	// Check if user is a member with sufficient permissions (owner or admin)
	member, err := s.querier.GetTenantMemberByUserAndTenant(ctx, store.GetTenantMemberByUserAndTenantParams{
		UserID:   userID,
		TenantID: r.Msg.GetTenantId(),
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, connect.NewError(connect.CodePermissionDenied, errors.New("not a member of this tenant"))
		}
		panic(err)
	}

	if member.Role != "owner" && member.Role != "admin" {
		return nil, connect.NewError(connect.CodePermissionDenied, errors.New("only owner or admin can create invitations"))
	}

	// Generate invite code
	inviteCode, err := generateInviteCode()
	if err != nil {
		panic(err)
	}

	// Create invitation (expires in 7 days)
	expiresAt := time.Now().Add(7 * 24 * time.Hour).Format("2006-01-02 15:04:05")
	invitation, err := s.querier.CreateTenantInvitation(ctx, store.CreateTenantInvitationParams{
		TenantID:   r.Msg.GetTenantId(),
		InviteCode: inviteCode,
		CreatedBy:  userID,
		ExpiresAt:  sql.NullString{String: expiresAt, Valid: true},
	})
	if err != nil {
		panic(err)
	}

	return connect.NewResponse(&models.CreateInvitationResponse{
		Invitation: translateInvitation(invitation),
	}), nil
}

// AcceptInvitation accepts an invitation and joins the tenant.
func (s *Service) AcceptInvitation(ctx context.Context, r *connect.Request[models.AcceptInvitationRequest]) (*connect.Response[models.AcceptInvitationResponse], error) {
	log, _ := cloudzap.GetLogger(ctx)
	log.Info("AcceptInvitation")

	userID, err := userIDFromContext(ctx)
	if err != nil {
		return nil, connect.NewError(connect.CodeUnauthenticated, err)
	}

	// Get the invitation by code
	invitation, err := s.querier.GetTenantInvitationByCode(ctx, r.Msg.GetInviteCode())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, connect.NewError(connect.CodeNotFound, errors.New("invitation not found or invalid"))
		}
		panic(err)
	}

	// Check if already used
	if invitation.UsedAt.Valid {
		return nil, connect.NewError(connect.CodeFailedPrecondition, errors.New("invitation has already been used"))
	}

	// Check if expired
	if invitation.ExpiresAt.Valid {
		expiresAt, _ := time.Parse("2006-01-02 15:04:05", invitation.ExpiresAt.String)
		if time.Now().After(expiresAt) {
			return nil, connect.NewError(connect.CodeFailedPrecondition, errors.New("invitation has expired"))
		}
	}

	// Check if user is already a member
	isMember, err := s.querier.IsTenantMember(ctx, store.IsTenantMemberParams{
		TenantID: invitation.TenantID,
		UserID:   userID,
	})
	if err != nil {
		panic(err)
	}
	if isMember {
		return nil, connect.NewError(connect.CodeAlreadyExists, errors.New("already a member of this tenant"))
	}

	// Mark invitation as used
	_, err = s.querier.MarkInvitationUsed(ctx, store.MarkInvitationUsedParams{
		UsedBy: sql.NullString{String: userID, Valid: true},
		ID:     invitation.ID,
	})
	if err != nil {
		panic(err)
	}

	// Create membership
	member, err := s.querier.CreateTenantMember(ctx, store.CreateTenantMemberParams{
		TenantID: invitation.TenantID,
		UserID:   userID,
		Role:     "member",
	})
	if err != nil {
		panic(err)
	}

	// Get the tenant
	tenant, err := s.querier.GetTenant(ctx, invitation.TenantID)
	if err != nil {
		panic(err)
	}

	return connect.NewResponse(&models.AcceptInvitationResponse{
		Tenant:     translateTenant(tenant),
		Membership: translateMembership(member),
	}), nil
}

// ListInvitations lists invitations for a tenant (owner/admin only).
func (s *Service) ListInvitations(ctx context.Context, r *connect.Request[models.ListInvitationsRequest]) (*connect.Response[models.ListInvitationsResponse], error) {
	log, _ := cloudzap.GetLogger(ctx)
	log.Info("ListInvitations")

	userID, err := userIDFromContext(ctx)
	if err != nil {
		return nil, connect.NewError(connect.CodeUnauthenticated, err)
	}

	// Check if user is a member with sufficient permissions
	member, err := s.querier.GetTenantMemberByUserAndTenant(ctx, store.GetTenantMemberByUserAndTenantParams{
		UserID:   userID,
		TenantID: r.Msg.GetTenantId(),
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, connect.NewError(connect.CodePermissionDenied, errors.New("not a member of this tenant"))
		}
		panic(err)
	}

	if member.Role != "owner" && member.Role != "admin" {
		return nil, connect.NewError(connect.CodePermissionDenied, errors.New("only owner or admin can view invitations"))
	}

	invitations, err := s.querier.ListTenantInvitations(ctx, r.Msg.GetTenantId())
	if err != nil {
		panic(err)
	}

	result := make([]*models.TenantInvitation, len(invitations))
	for i, inv := range invitations {
		result[i] = translateInvitation(inv)
	}

	return connect.NewResponse(&models.ListInvitationsResponse{
		Invitations: result,
	}), nil
}

// RevokeInvitation revokes an invitation (owner/admin only).
func (s *Service) RevokeInvitation(ctx context.Context, r *connect.Request[models.RevokeInvitationRequest]) (*connect.Response[models.RevokeInvitationResponse], error) {
	log, _ := cloudzap.GetLogger(ctx)
	log.Info("RevokeInvitation")

	userID, err := userIDFromContext(ctx)
	if err != nil {
		return nil, connect.NewError(connect.CodeUnauthenticated, err)
	}

	// Get the invitation
	invitation, err := s.querier.GetTenantInvitation(ctx, r.Msg.GetInvitationId())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, connect.NewError(connect.CodeNotFound, errors.New("invitation not found"))
		}
		panic(err)
	}

	// Check if user is a member with sufficient permissions
	member, err := s.querier.GetTenantMemberByUserAndTenant(ctx, store.GetTenantMemberByUserAndTenantParams{
		UserID:   userID,
		TenantID: invitation.TenantID,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, connect.NewError(connect.CodePermissionDenied, errors.New("not a member of this tenant"))
		}
		panic(err)
	}

	if member.Role != "owner" && member.Role != "admin" {
		return nil, connect.NewError(connect.CodePermissionDenied, errors.New("only owner or admin can revoke invitations"))
	}

	// Delete the invitation
	if _, err := s.querier.DeleteTenantInvitation(ctx, invitation.ID); err != nil {
		panic(err)
	}

	return connect.NewResponse(&models.RevokeInvitationResponse{}), nil
}
