import { Avatar, Button, Group, Loader, Menu, Text, UnstyledButton } from '@mantine/core';
import { Link } from 'react-router-dom';
import { useAuth0 } from '../../lib/auth0';

/**
 * Authentication button that shows login/logout based on auth state.
 * Shows user avatar and menu when logged in.
 */
export function AuthButton() {
  const { isAuthenticated, isLoading, loginWithRedirect, logout, user } = useAuth0();

  if (isLoading) {
    return <Loader size="sm" color="coral" />;
  }

  if (isAuthenticated && user) {
    return (
      <Menu shadow="md" width={200} position="bottom-end">
        <Menu.Target>
          <UnstyledButton>
            <Group gap="xs">
              <Avatar src={user.picture} alt={user.name} radius="xl" size="sm" />
              <Text size="sm" fw={500} visibleFrom="sm">
                {user.name}
              </Text>
            </Group>
          </UnstyledButton>
        </Menu.Target>

        <Menu.Dropdown>
          <Menu.Label>Compte</Menu.Label>
          <Menu.Item>
            <Text size="sm" c="dimmed">
              {user.email}
            </Text>
          </Menu.Item>
          <Menu.Divider />
          <Menu.Item component={Link} to="/settings">
            Paramètres
          </Menu.Item>
          <Menu.Item
            color="red"
            onClick={() => logout({ logoutParams: { returnTo: window.location.origin } })}
          >
            Déconnexion
          </Menu.Item>
        </Menu.Dropdown>
      </Menu>
    );
  }

  return (
    <Group gap="sm">
      <Button variant="light" color="coral" radius="xl" onClick={() => loginWithRedirect()}>
        Connexion
      </Button>
      <Button
        color="coral"
        radius="xl"
        onClick={() =>
          loginWithRedirect({
            authorizationParams: {
              screen_hint: 'signup',
            },
          })
        }
      >
        Commencer
      </Button>
    </Group>
  );
}

