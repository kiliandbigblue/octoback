package v1

import (
	"time"

	models "github.com/kiliandbigblue/octoback/gen/proto/go/octoback/tenants/v1"
	store "github.com/kiliandbigblue/octoback/gen/store/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func translateTenant(t store.Tenant) *models.Tenant {
	createdAt, _ := time.Parse("2006-01-02 15:04:05", t.CreatedAt)
	return &models.Tenant{
		Id:         t.ID,
		Name:       t.Name,
		CreateTime: timestamppb.New(createdAt),
		Version:    int32(t.Version),
	}
}

func translateMembership(m store.TenantMember) *models.TenantMembership {
	createdAt, _ := time.Parse("2006-01-02 15:04:05", m.CreatedAt)
	return &models.TenantMembership{
		Id:         m.ID,
		TenantId:   m.TenantID,
		UserId:     m.UserID,
		Role:       m.Role,
		CreateTime: timestamppb.New(createdAt),
	}
}

func translateInvitation(i store.TenantInvitation) *models.TenantInvitation {
	createdAt, _ := time.Parse("2006-01-02 15:04:05", i.CreatedAt)
	invitation := &models.TenantInvitation{
		Id:         i.ID,
		TenantId:   i.TenantID,
		InviteCode: i.InviteCode,
		CreatedBy:  i.CreatedBy,
		CreateTime: timestamppb.New(createdAt),
	}
	if i.ExpiresAt.Valid {
		expiresAt, _ := time.Parse("2006-01-02 15:04:05", i.ExpiresAt.String)
		invitation.ExpiresAt = timestamppb.New(expiresAt)
	}
	if i.UsedAt.Valid {
		usedAt, _ := time.Parse("2006-01-02 15:04:05", i.UsedAt.String)
		invitation.UsedAt = timestamppb.New(usedAt)
	}
	if i.UsedBy.Valid {
		invitation.UsedBy = i.UsedBy.String
	}
	return invitation
}
