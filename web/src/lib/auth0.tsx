import { Auth0Provider, useAuth0 } from '@auth0/auth0-react';
import type { ReactNode } from 'react';

// Auth0 configuration from environment variables
const auth0Config = {
  domain: import.meta.env.VITE_AUTH0_DOMAIN || '',
  clientId: import.meta.env.VITE_AUTH0_CLIENT_ID || '',
  audience: import.meta.env.VITE_AUTH0_AUDIENCE || '',
};

interface AuthProviderProps {
  children: ReactNode;
}

/**
 * Auth0 provider wrapper that configures authentication for the app.
 */
export function AuthProvider({ children }: AuthProviderProps) {
  const redirectUri = window.location.origin;

  if (!auth0Config.domain || !auth0Config.clientId) {
    console.warn('Auth0 configuration missing. Set VITE_AUTH0_DOMAIN and VITE_AUTH0_CLIENT_ID environment variables.');
    return <>{children}</>;
  }

  return (
    <Auth0Provider
      domain={auth0Config.domain}
      clientId={auth0Config.clientId}
      authorizationParams={{
        redirect_uri: redirectUri,
        audience: auth0Config.audience,
      }}
      cacheLocation="localstorage"
    >
      {children}
    </Auth0Provider>
  );
}

/**
 * Hook to get the access token for API calls.
 * Returns a function that retrieves the token asynchronously.
 */
export function useAccessToken() {
  const { getAccessTokenSilently, isAuthenticated } = useAuth0();

  const getToken = async (): Promise<string | null> => {
    if (!isAuthenticated) {
      return null;
    }

    try {
      const token = await getAccessTokenSilently({
        authorizationParams: {
          audience: auth0Config.audience,
        },
      });
      return token;
    } catch (error) {
      console.error('Failed to get access token:', error);
      return null;
    }
  };

  return { getToken, isAuthenticated };
}

// Re-export useAuth0 for convenience
export { useAuth0 };

