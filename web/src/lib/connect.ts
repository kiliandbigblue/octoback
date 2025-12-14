import { createConnectTransport } from '@connectrpc/connect-web';

// Create a transport for Connect-RPC
// In development, we use localhost:8080 where the Go server runs
// In production, this should point to your backend URL
export const transport = createConnectTransport({
  baseUrl: import.meta.env.VITE_API_URL || 'http://localhost:8080',
});

