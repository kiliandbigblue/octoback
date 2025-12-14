import '@mantine/core/styles.css';

import { TransportProvider } from '@connectrpc/connect-query';
import { MantineProvider } from '@mantine/core';
import { QueryClientProvider } from '@tanstack/react-query';
import { transport } from './lib/connect';
import { queryClient } from './lib/query-client';
import { Router } from './Router';
import { theme } from './theme';

export default function App() {
  return (
    <QueryClientProvider client={queryClient}>
      <TransportProvider transport={transport}>
        <MantineProvider theme={theme}>
          <Router />
        </MantineProvider>
      </TransportProvider>
    </QueryClientProvider>
  );
}
