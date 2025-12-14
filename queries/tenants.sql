-- name: GetTenant :one
SELECT *
FROM tenant
WHERE id = ?
LIMIT 1;

-- name: CreateTenant :one
INSERT INTO tenant (name)
VALUES (?)
RETURNING *;

-- name: UpdateTenant :one
UPDATE tenant
SET name = ?, version = version + 1
WHERE id = ? AND version = ?
RETURNING *;

-- name: DeleteTenant :one
DELETE FROM tenant
WHERE id = ?
RETURNING *;

-- name: ListTenants :many
SELECT *
FROM tenant;

-- name: GetTenantMember :one
SELECT *
FROM tenant_member
WHERE id = ?
LIMIT 1;

-- name: GetTenantMemberByUserAndTenant :one
SELECT *
FROM tenant_member
WHERE user_id = ? AND tenant_id = ?
LIMIT 1;

-- name: GetTenantsByUserId :many
SELECT t.*
FROM tenant t
JOIN tenant_member tm ON t.id = tm.tenant_id
WHERE tm.user_id = ?;

-- name: GetDefaultTenantByUserId :one
SELECT t.*
FROM tenant t
JOIN tenant_member tm ON t.id = tm.tenant_id
WHERE tm.user_id = ?
ORDER BY tm.created_at ASC
LIMIT 1;

-- name: CreateTenantMember :one
INSERT INTO tenant_member (tenant_id, user_id, role)
VALUES (?, ?, ?)
RETURNING *;

-- name: UpdateTenantMemberRole :one
UPDATE tenant_member
SET role = ?
WHERE id = ?
RETURNING *;

-- name: DeleteTenantMember :one
DELETE FROM tenant_member
WHERE id = ?
RETURNING *;

-- name: ListTenantMembers :many
SELECT *
FROM tenant_member
WHERE tenant_id = ?;

-- name: IsTenantMember :one
SELECT COUNT(*) > 0 AS is_member
FROM tenant_member
WHERE tenant_id = ? AND user_id = ?;

-- name: CountUserTenants :one
SELECT COUNT(*) AS count
FROM tenant_member
WHERE user_id = ?;

-- ============ Invitation Queries ============

-- name: CreateTenantInvitation :one
INSERT INTO tenant_invitation (tenant_id, invite_code, created_by, expires_at)
VALUES (?, ?, ?, ?)
RETURNING *;

-- name: GetTenantInvitationByCode :one
SELECT *
FROM tenant_invitation
WHERE invite_code = ?
LIMIT 1;

-- name: GetTenantInvitation :one
SELECT *
FROM tenant_invitation
WHERE id = ?
LIMIT 1;

-- name: ListTenantInvitations :many
SELECT *
FROM tenant_invitation
WHERE tenant_id = ?
ORDER BY created_at DESC;

-- name: MarkInvitationUsed :one
UPDATE tenant_invitation
SET used_at = datetime('now'), used_by = ?
WHERE id = ? AND used_at IS NULL
RETURNING *;

-- name: DeleteTenantInvitation :one
DELETE FROM tenant_invitation
WHERE id = ?
RETURNING *;

-- name: DeleteExpiredInvitations :exec
DELETE FROM tenant_invitation
WHERE expires_at IS NOT NULL AND expires_at < datetime('now') AND used_at IS NULL;

