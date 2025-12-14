import { useEffect, useRef, type ReactNode } from 'react';
import { useEnsureUserTenant } from '../hooks';
import { useAuth0 } from './auth0';

interface TenantInitializerProps {
  children: ReactNode;
}

/**
 * Component that ensures the authenticated user has at least one tenant.
 * Automatically creates a personal tenant for new users on first login.
 */
export function TenantInitializer({ children }: TenantInitializerProps) {
  const { isAuthenticated, isLoading } = useAuth0();
  const ensureUserTenantMutation = useEnsureUserTenant();
  const hasInitialized = useRef(false);

  useEffect(() => {
    // Only run once per session when user is authenticated
    if (isAuthenticated && !isLoading && !hasInitialized.current) {
      hasInitialized.current = true;
      ensureUserTenantMutation.mutate({});
    }
  }, [isAuthenticated, isLoading, ensureUserTenantMutation]);

  return <>{children}</>;
}

