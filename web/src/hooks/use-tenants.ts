import { useMutation, useQuery } from '@connectrpc/connect-query';
import {
  getMyTenants,
  joinTenant,
  leaveTenant,
  listTenants,
} from '../gen/octoback/tenants/v1/service-Service_connectquery';

// ============ Query Hooks ============

/**
 * Fetch all available tenants
 */
export function useTenants() {
  return useQuery(listTenants, {});
}

/**
 * Fetch tenants the current user is a member of
 */
export function useMyTenants() {
  return useQuery(getMyTenants, {});
}

// ============ Mutation Hooks ============

/**
 * Join a tenant
 */
export function useJoinTenant() {
  return useMutation(joinTenant);
}

/**
 * Leave a tenant
 */
export function useLeaveTenant() {
  return useMutation(leaveTenant);
}

