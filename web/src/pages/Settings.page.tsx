import {
  ActionIcon,
  Alert,
  Avatar,
  Badge,
  Box,
  Button,
  Card,
  CopyButton,
  Container,
  Divider,
  Group,
  Loader,
  Modal,
  Stack,
  Text,
  TextInput,
  ThemeIcon,
  Title,
  Tooltip,
  useMantineColorScheme,
} from '@mantine/core';
import { useDisclosure } from '@mantine/hooks';
import { useQueryClient } from '@tanstack/react-query';
import { useState } from 'react';
import { Link } from 'react-router-dom';
import { ColorSchemeToggle } from '../components/ColorSchemeToggle/ColorSchemeToggle';
import {
  useAcceptInvitation,
  useCreateInvitation,
  useLeaveTenant,
  useMyTenants,
} from '../hooks';
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

function ShareIcon() {
  return (
    <svg width="18" height="18" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
      <circle cx="18" cy="5" r="3" stroke="currentColor" strokeWidth="2" />
      <circle cx="6" cy="12" r="3" stroke="currentColor" strokeWidth="2" />
      <circle cx="18" cy="19" r="3" stroke="currentColor" strokeWidth="2" />
      <path d="M8.59 13.51L15.42 17.49M15.41 6.51L8.59 10.49" stroke="currentColor" strokeWidth="2" />
    </svg>
  );
}

function CopyIcon() {
  return (
    <svg width="16" height="16" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
      <rect x="9" y="9" width="13" height="13" rx="2" stroke="currentColor" strokeWidth="2" />
      <path
        d="M5 15H4C2.89543 15 2 14.1046 2 13V4C2 2.89543 2.89543 2 4 2H13C14.1046 2 15 2.89543 15 4V5"
        stroke="currentColor"
        strokeWidth="2"
      />
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

const roleLabels: Record<string, string> = {
  owner: 'Propriétaire',
  admin: 'Admin',
  member: 'Membre',
};

const roleColors: Record<string, string> = {
  owner: 'coral',
  admin: 'blue',
  member: 'gray',
};

function TenantCard({
  name,
  role,
  isLoading,
  onLeave,
  onInvite,
}: {
  name: string;
  role: string;
  isLoading: boolean;
  onLeave: () => void;
  onInvite: () => void;
}) {
  const canInvite = role === 'owner' || role === 'admin';
  const canLeave = role !== 'owner';

  return (
    <Card className={classes.tenantCard} data-member={true}>
      <Group justify="space-between" align="flex-start">
        <Group gap="md">
          <ThemeIcon size={48} radius="xl" variant="filled" color="coral">
            <HouseIcon />
          </ThemeIcon>
          <Stack gap={4}>
            <Text fw={600} size="lg">
              {name}
            </Text>
            <Badge variant="light" color={roleColors[role] || 'gray'} size="sm">
              {roleLabels[role] || role}
            </Badge>
          </Stack>
        </Group>

        <Group gap="xs">
          {canInvite && (
            <Tooltip label="Inviter quelqu'un">
              <ActionIcon variant="light" color="coral" size="lg" radius="xl" onClick={onInvite}>
                <ShareIcon />
              </ActionIcon>
            </Tooltip>
          )}
          {isLoading ? (
            <Loader size="sm" color="coral" />
          ) : canLeave ? (
            <Button
              variant="subtle"
              color="red"
              size="xs"
              onClick={onLeave}
              className={classes.leaveButton}
            >
              Quitter
            </Button>
          ) : null}
        </Group>
      </Group>
    </Card>
  );
}

function InviteModal({
  opened,
  onClose,
  tenantId,
  tenantName,
}: {
  opened: boolean;
  onClose: () => void;
  tenantId: bigint;
  tenantName: string;
}) {
  const createInvitationMutation = useCreateInvitation();
  const [inviteCode, setInviteCode] = useState<string | null>(null);

  const handleCreate = async () => {
    const result = await createInvitationMutation.mutateAsync({ tenantId });
    if (result.invitation?.inviteCode) {
      setInviteCode(result.invitation.inviteCode);
    }
  };

  const handleClose = () => {
    setInviteCode(null);
    onClose();
  };

  return (
    <Modal opened={opened} onClose={handleClose} title={`Inviter dans "${tenantName}"`} centered>
      <Stack gap="md">
        {!inviteCode ? (
          <>
            <Text size="sm" c="dimmed">
              Créez un code d'invitation à partager avec la personne que vous souhaitez inviter.
              Le code expire après 7 jours.
            </Text>
            <Button
              color="coral"
              onClick={handleCreate}
              loading={createInvitationMutation.isPending}
              fullWidth
            >
              Créer un code d'invitation
            </Button>
          </>
        ) : (
          <>
            <Text size="sm" c="dimmed">
              Partagez ce code avec la personne que vous souhaitez inviter :
            </Text>
            <Card withBorder p="md">
              <Group justify="space-between">
                <Text fw={600} size="lg" ff="monospace">
                  {inviteCode}
                </Text>
                <CopyButton value={inviteCode}>
                  {({ copied, copy }) => (
                    <Tooltip label={copied ? 'Copié !' : 'Copier'}>
                      <ActionIcon color={copied ? 'teal' : 'coral'} variant="light" onClick={copy}>
                        {copied ? <CheckIcon /> : <CopyIcon />}
                      </ActionIcon>
                    </Tooltip>
                  )}
                </CopyButton>
              </Group>
            </Card>
            <Button variant="light" color="coral" onClick={handleClose} fullWidth>
              Fermer
            </Button>
          </>
        )}
      </Stack>
    </Modal>
  );
}

export function SettingsPage() {
  const { colorScheme } = useMantineColorScheme();
  const { user, isAuthenticated, isLoading: authLoading, loginWithRedirect } = useAuth0();
  const queryClient = useQueryClient();

  const { data: myTenantsData, isLoading: myTenantsLoading, error: tenantsError } = useMyTenants();
  const leaveTenantMutation = useLeaveTenant();
  const acceptInvitationMutation = useAcceptInvitation();

  const [inviteModalOpened, { open: openInviteModal, close: closeInviteModal }] =
    useDisclosure(false);
  const [selectedTenant, setSelectedTenant] = useState<{ id: bigint; name: string } | null>(null);
  const [joinCode, setJoinCode] = useState('');
  const [joinError, setJoinError] = useState<string | null>(null);

  const handleLeave = async (tenantId: bigint) => {
    await leaveTenantMutation.mutateAsync({ tenantId });
    queryClient.invalidateQueries({ queryKey: ['octoback.tenants.v1.Service'] });
  };

  const handleInvite = (tenantId: bigint, tenantName: string) => {
    setSelectedTenant({ id: tenantId, name: tenantName });
    openInviteModal();
  };

  const handleJoin = async () => {
    if (!joinCode.trim()) return;
    setJoinError(null);
    try {
      await acceptInvitationMutation.mutateAsync({ inviteCode: joinCode.trim() });
      setJoinCode('');
      queryClient.invalidateQueries({ queryKey: ['octoback.tenants.v1.Service'] });
    } catch (error: unknown) {
      if (error instanceof Error) {
        if (error.message.includes('not found')) {
          setJoinError('Code d\'invitation invalide ou expiré');
        } else if (error.message.includes('already')) {
          setJoinError('Vous êtes déjà membre de ce foyer');
        } else if (error.message.includes('expired')) {
          setJoinError('Ce code d\'invitation a expiré');
        } else {
          setJoinError('Une erreur est survenue');
        }
      }
    }
  };

  // Create a map of tenant ID to membership for quick lookup
  const membershipMap = new Map(
    myTenantsData?.memberships?.map((m) => [m.tenantId, m]) ?? []
  );

  const isLoading = authLoading || myTenantsLoading;

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
                  Gérez vos foyers et invitez d'autres personnes à les rejoindre.
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
              <Stack gap="md">
                {myTenantsData?.tenants?.map((tenant) => {
                  const membership = membershipMap.get(tenant.id);
                  return (
                    <TenantCard
                      key={tenant.id.toString()}
                      name={tenant.name}
                      role={membership?.role || 'member'}
                      isLoading={
                        leaveTenantMutation.isPending &&
                        leaveTenantMutation.variables?.tenantId === tenant.id
                      }
                      onLeave={() => handleLeave(tenant.id)}
                      onInvite={() => handleInvite(tenant.id, tenant.name)}
                    />
                  );
                })}
              </Stack>
            )}

            {myTenantsData?.tenants?.length === 0 && !isLoading && (
              <Card className={classes.emptyCard}>
                <Stack align="center" gap="md" py="xl">
                  <ThemeIcon size={64} radius="xl" variant="light" color="gray">
                    <HouseIcon />
                  </ThemeIcon>
                  <Text c="dimmed" ta="center">
                    Vous n'êtes membre d'aucun foyer.
                  </Text>
                </Stack>
              </Card>
            )}
          </Stack>

          <Divider />

          {/* Join with invitation code */}
          <Stack gap="md">
            <Stack gap={4}>
              <Title order={4}>Rejoindre un foyer</Title>
              <Text c="dimmed" size="sm">
                Entrez un code d'invitation pour rejoindre un foyer existant.
              </Text>
            </Stack>

            <Group>
              <TextInput
                placeholder="Code d'invitation"
                value={joinCode}
                onChange={(e) => {
                  setJoinCode(e.currentTarget.value);
                  setJoinError(null);
                }}
                style={{ flex: 1 }}
                error={joinError}
              />
              <Button
                color="coral"
                onClick={handleJoin}
                loading={acceptInvitationMutation.isPending}
                disabled={!joinCode.trim()}
              >
                Rejoindre
              </Button>
            </Group>
          </Stack>
        </Stack>
      </Container>

      {/* Invite Modal */}
      {selectedTenant && (
        <InviteModal
          opened={inviteModalOpened}
          onClose={closeInviteModal}
          tenantId={selectedTenant.id}
          tenantName={selectedTenant.name}
        />
      )}
    </Box>
  );
}
