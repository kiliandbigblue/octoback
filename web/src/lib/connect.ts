import type { Interceptor } from '@connectrpc/connect';
import { createConnectTransport } from '@connectrpc/connect-web';

// Create a transport for Connect-RPC
// In development, we use localhost:8080 where the Go server runs
// In production, this should point to your backend URL
const baseUrl = import.meta.env.VITE_API_URL || 'http://localhost:8080';

/**
 * Creates an auth interceptor that adds the Authorization header to requests.
 */
export function createAuthInterceptor(getToken: () => Promise<string | null>): Interceptor {
  return (next) => async (req) => {
    const token = await getToken();
    if (token) {
      req.header.set('Authorization', `Bearer ${token}`);
    }
    return next(req);
  };
}

/**
 * Creates a Connect transport with authentication.
 */
export function createAuthenticatedTransport(getToken: () => Promise<string | null>) {
  return createConnectTransport({
    baseUrl,
    interceptors: [createAuthInterceptor(getToken)],
  });
}

/**
 * Default transport without authentication (for unauthenticated routes).
 * @deprecated Use createAuthenticatedTransport with AuthenticatedTransportProvider instead.
 */
export const transport = createConnectTransport({
  baseUrl,
});
