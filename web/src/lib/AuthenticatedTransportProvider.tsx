import { TransportProvider } from '@connectrpc/connect-query';
import { useMemo, type ReactNode } from 'react';
import { useAccessToken } from './auth0';
import { createAuthenticatedTransport } from './connect';

interface AuthenticatedTransportProviderProps {
  children: ReactNode;
}

/**
 * Transport provider that automatically includes the Auth0 access token
 * in all Connect-RPC requests.
 */
export function AuthenticatedTransportProvider({ children }: AuthenticatedTransportProviderProps) {
  const { getToken } = useAccessToken();

  // Memoize the transport to avoid recreating it on every render
  const transport = useMemo(() => createAuthenticatedTransport(getToken), [getToken]);

  return <TransportProvider transport={transport}>{children}</TransportProvider>;
}

