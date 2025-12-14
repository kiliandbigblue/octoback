import { useMutation, useQuery } from '@connectrpc/connect-query';
import {
  acceptInvitation,
  createInvitation,
  ensureUserTenant,
  getMyTenants,
  leaveTenant,
  listInvitations,
  revokeInvitation,
} from '../gen/octoback/tenants/v1/service-Service_connectquery';

// ============ Query Hooks ============

/**
 * Fetch tenants the current user is a member of
 */
export function useMyTenants() {
  return useQuery(getMyTenants, {});
}

/**
 * List invitations for a specific tenant (owner/admin only)
 */
export function useListInvitations(tenantId: bigint) {
  return useQuery(listInvitations, { tenantId });
}

// ============ Mutation Hooks ============

/**
 * Ensure the user has a tenant (creates one if needed)
 */
export function useEnsureUserTenant() {
  return useMutation(ensureUserTenant);
}

/**
 * Leave a tenant
 */
export function useLeaveTenant() {
  return useMutation(leaveTenant);
}

/**
 * Create an invitation for a tenant
 */
export function useCreateInvitation() {
  return useMutation(createInvitation);
}

/**
 * Accept an invitation to join a tenant
 */
export function useAcceptInvitation() {
  return useMutation(acceptInvitation);
}

/**
 * Revoke an invitation
 */
export function useRevokeInvitation() {
  return useMutation(revokeInvitation);
}
