import {
  ActionIcon,
  Alert,
  Avatar,
  Badge,
  Box,
  Button,
  Card,
  Container,
  Divider,
  Group,
  Loader,
  SimpleGrid,
  Stack,
  Text,
  ThemeIcon,
  Title,
  useMantineColorScheme,
} from '@mantine/core';
import { useQueryClient } from '@tanstack/react-query';
import { Link } from 'react-router-dom';
import { ColorSchemeToggle } from '../components/ColorSchemeToggle/ColorSchemeToggle';
import { useJoinTenant, useLeaveTenant, useMyTenants, useTenants } from '../hooks';
import { useAuth0 } from '../lib/auth0';
import classes from './Settings.module.css';

function HouseIcon() {
  return (
    <svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
      <path
        d="M3 10.5L12 3L21 10.5V21H15V15H9V21H3V10.5Z"
        stroke="currentColor"
        strokeWidth="2"
        strokeLinecap="round"
        strokeLinejoin="round"
      />
    </svg>
  );
}

function CheckIcon() {
  return (
    <svg width="16" height="16" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
      <path
        d="M5 12L10 17L20 7"
        stroke="currentColor"
        strokeWidth="2"
        strokeLinecap="round"
        strokeLinejoin="round"
      />
    </svg>
  );
}

function PlusIcon() {
  return (
    <svg width="16" height="16" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
      <path d="M12 5V19M5 12H19" stroke="currentColor" strokeWidth="2" strokeLinecap="round" />
    </svg>
  );
}

function BackIcon() {
  return (
    <svg width="20" height="20" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
      <path
        d="M19 12H5M5 12L12 19M5 12L12 5"
        stroke="currentColor"
        strokeWidth="2"
        strokeLinecap="round"
        strokeLinejoin="round"
      />
    </svg>
  );
}

function OctopusLogo({ size = 48 }: { size?: number }) {
  return (
    <svg
      width={size}
      height={size}
      viewBox="0 0 64 64"
      fill="none"
      xmlns="http://www.w3.org/2000/svg"
    >
      <circle cx="32" cy="24" r="16" fill="currentColor" />
      <circle cx="26" cy="22" r="3" fill="white" />
      <circle cx="38" cy="22" r="3" fill="white" />
      <circle cx="26" cy="22" r="1.5" fill="#2d3436" />
      <circle cx="38" cy="22" r="1.5" fill="#2d3436" />
      <path
        d="M12 40C12 40 8 52 12 56C14 58 18 56 18 52C18 48 16 44 20 40"
        stroke="currentColor"
        strokeWidth="4"
        strokeLinecap="round"
      />
      <path
        d="M20 42C20 42 18 54 22 58C24 60 28 58 28 54C28 50 26 46 28 42"
        stroke="currentColor"
        strokeWidth="4"
        strokeLinecap="round"
      />
      <path
        d="M36 42C36 42 38 54 42 58C44 60 40 60 38 56C36 52 38 48 36 42"
        stroke="currentColor"
        strokeWidth="4"
        strokeLinecap="round"
      />
      <path
        d="M44 40C44 40 48 52 52 56C54 58 50 58 48 54C46 50 48 44 44 40"
        stroke="currentColor"
        strokeWidth="4"
        strokeLinecap="round"
      />
      <ellipse cx="32" cy="28" rx="4" ry="2" fill="#f68f71" opacity="0.6" />
    </svg>
  );
}

function TenantCard({
  name,
  isMember,
  isLoading,
  onJoin,
  onLeave,
}: {
  name: string;
  isMember: boolean;
  isLoading: boolean;
  onJoin: () => void;
  onLeave: () => void;
}) {
  return (
    <Card className={classes.tenantCard} data-member={isMember}>
      <Group justify="space-between" align="flex-start">
        <Group gap="md">
          <ThemeIcon
            size={48}
            radius="xl"
            variant={isMember ? 'filled' : 'light'}
            color={isMember ? 'coral' : 'gray'}
          >
            <HouseIcon />
          </ThemeIcon>
          <Stack gap={4}>
            <Text fw={600} size="lg">
              {name}
            </Text>
            {isMember && (
              <Badge variant="light" color="coral" size="sm">
                Membre
              </Badge>
            )}
          </Stack>
        </Group>

        {isLoading ? (
          <Loader size="sm" color="coral" />
        ) : isMember ? (
          <Button
            variant="subtle"
            color="red"
            size="xs"
            onClick={onLeave}
            className={classes.leaveButton}
          >
            Quitter
          </Button>
        ) : (
          <ActionIcon
            variant="light"
            color="coral"
            size="lg"
            radius="xl"
            onClick={onJoin}
            className={classes.joinButton}
          >
            <PlusIcon />
          </ActionIcon>
        )}
      </Group>
    </Card>
  );
}

export function SettingsPage() {
  const { colorScheme } = useMantineColorScheme();
  const { user, isAuthenticated, isLoading: authLoading, loginWithRedirect } = useAuth0();
  const queryClient = useQueryClient();

  const { data: tenantsData, isLoading: tenantsLoading, error: tenantsError } = useTenants();
  const { data: myTenantsData, isLoading: myTenantsLoading } = useMyTenants();

  const joinTenantMutation = useJoinTenant();
  const leaveTenantMutation = useLeaveTenant();

  const myTenantIds = new Set(myTenantsData?.tenants?.map((t) => t.id) ?? []);

  const handleJoin = async (tenantId: bigint) => {
    await joinTenantMutation.mutateAsync({ tenantId });
    queryClient.invalidateQueries({ queryKey: ['octoback.tenants.v1.Service'] });
  };

  const handleLeave = async (tenantId: bigint) => {
    await leaveTenantMutation.mutateAsync({ tenantId });
    queryClient.invalidateQueries({ queryKey: ['octoback.tenants.v1.Service'] });
  };

  const isLoading = authLoading || tenantsLoading || myTenantsLoading;

  if (authLoading) {
    return (
      <Box className={classes.wrapper} data-theme={colorScheme}>
        <Container size="md" py={100}>
          <Stack align="center" gap="xl">
            <Loader size="lg" color="coral" />
            <Text c="dimmed">Chargement...</Text>
          </Stack>
        </Container>
      </Box>
    );
  }

  if (!isAuthenticated) {
    return (
      <Box className={classes.wrapper} data-theme={colorScheme}>
        <Container size="md" py={100}>
          <Stack align="center" gap="xl">
            <Box c="coral.5">
              <OctopusLogo size={80} />
            </Box>
            <Title order={2} ta="center">
              Connexion requise
            </Title>
            <Text c="dimmed" ta="center" maw={400}>
              Connectez-vous pour accéder à vos paramètres et gérer vos foyers.
            </Text>
            <Button size="lg" radius="xl" color="coral" onClick={() => loginWithRedirect()}>
              Se connecter
            </Button>
          </Stack>
        </Container>
      </Box>
    );
  }

  return (
    <Box className={classes.wrapper} data-theme={colorScheme}>
      {/* Background decoration */}
      <Box className={classes.bgDecoration} />

      {/* Header */}
      <Box component="header" className={classes.header}>
        <Container size="xl">
          <Group justify="space-between" h={70}>
            <Group gap="md">
              <ActionIcon
                component={Link}
                to="/"
                variant="subtle"
                color={colorScheme === 'dark' ? 'gray' : 'charcoal'}
                size="lg"
              >
                <BackIcon />
              </ActionIcon>
              <Group gap="xs">
                <Box c="coral.5">
                  <OctopusLogo size={36} />
                </Box>
                <Text fw={700} size="lg" style={{ letterSpacing: '-0.5px' }}>
                  Paramètres
                </Text>
              </Group>
            </Group>
            <ColorSchemeToggle />
          </Group>
        </Container>
      </Box>

      <Container size="md" py="xl">
        <Stack gap="xl">
          {/* User Profile Section */}
          <Card className={classes.profileCard}>
            <Group gap="lg">
              <Avatar src={user?.picture} alt={user?.name} size={80} radius="xl" />
              <Stack gap={4}>
                <Title order={3}>{user?.name}</Title>
                <Text c="dimmed" size="sm">
                  {user?.email}
                </Text>
              </Stack>
            </Group>
          </Card>

          <Divider />

          {/* Tenants Section */}
          <Stack gap="md">
            <Group justify="space-between" align="center">
              <Stack gap={4}>
                <Title order={4}>Mes foyers</Title>
                <Text c="dimmed" size="sm">
                  Rejoignez un foyer pour partager vos recettes et listes de courses.
                </Text>
              </Stack>
              {myTenantsData?.tenants && myTenantsData.tenants.length > 0 && (
                <Badge size="lg" variant="light" color="coral" leftSection={<CheckIcon />}>
                  {myTenantsData.tenants.length} foyer
                  {myTenantsData.tenants.length > 1 ? 's' : ''}
                </Badge>
              )}
            </Group>

            {tenantsError && (
              <Alert variant="light" color="red" title="Erreur">
                Impossible de charger les foyers. Veuillez réessayer.
              </Alert>
            )}

            {isLoading ? (
              <Stack align="center" py="xl">
                <Loader color="coral" />
              </Stack>
            ) : (
              <SimpleGrid cols={{ base: 1, sm: 2 }} spacing="md">
                {tenantsData?.tenants?.map((tenant) => (
                  <TenantCard
                    key={tenant.id.toString()}
                    name={tenant.name}
                    isMember={myTenantIds.has(tenant.id)}
                    isLoading={
                      (joinTenantMutation.isPending &&
                        joinTenantMutation.variables?.tenantId === tenant.id) ||
                      (leaveTenantMutation.isPending &&
                        leaveTenantMutation.variables?.tenantId === tenant.id)
                    }
                    onJoin={() => handleJoin(tenant.id)}
                    onLeave={() => handleLeave(tenant.id)}
                  />
                ))}
              </SimpleGrid>
            )}

            {tenantsData?.tenants?.length === 0 && !isLoading && (
              <Card className={classes.emptyCard}>
                <Stack align="center" gap="md" py="xl">
                  <ThemeIcon size={64} radius="xl" variant="light" color="gray">
                    <HouseIcon />
                  </ThemeIcon>
                  <Text c="dimmed" ta="center">
                    Aucun foyer disponible pour le moment.
                  </Text>
                </Stack>
              </Card>
            )}
          </Stack>
        </Stack>
      </Container>
    </Box>
  );
}

