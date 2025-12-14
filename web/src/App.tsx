import '@mantine/core/styles.css';

import { MantineProvider } from '@mantine/core';
import { QueryClientProvider } from '@tanstack/react-query';
import { AuthProvider } from './lib/auth0';
import { AuthenticatedTransportProvider } from './lib/AuthenticatedTransportProvider';
import { queryClient } from './lib/query-client';
import { Router } from './Router';
import { theme } from './theme';

export default function App() {
  return (
    <AuthProvider>
      <QueryClientProvider client={queryClient}>
        <AuthenticatedTransportProvider>
          <MantineProvider theme={theme}>
            <Router />
          </MantineProvider>
        </AuthenticatedTransportProvider>
      </QueryClientProvider>
    </AuthProvider>
  );
}
