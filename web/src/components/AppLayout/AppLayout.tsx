import {
  ActionIcon,
  Box,
  Container,
  Drawer,
  Group,
  Stack,
  Text,
  Tooltip,
  UnstyledButton,
  useMantineColorScheme,
} from '@mantine/core';
import { useDisclosure } from '@mantine/hooks';
import { Link, useLocation } from 'react-router-dom';
import { ColorSchemeToggle } from '../ColorSchemeToggle/ColorSchemeToggle';
import { useAuth0 } from '../../lib/auth0';
import classes from './AppLayout.module.css';

function OctopusLogo({ size = 32 }: { size?: number }) {
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

function ListIcon() {
  return (
    <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round" strokeLinejoin="round">
      <path d="M8 6h13M8 12h13M8 18h13M3 6h.01M3 12h.01M3 18h.01" />
    </svg>
  );
}

function ChefHatIcon() {
  return (
    <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round" strokeLinejoin="round">
      <path d="M6 13.87A4 4 0 0 1 7.41 6a5.11 5.11 0 0 1 1.05-1.54 5 5 0 0 1 7.08 0A5.11 5.11 0 0 1 16.59 6 4 4 0 0 1 18 13.87V21H6V13.87Z" />
      <line x1="6" y1="17" x2="18" y2="17" />
    </svg>
  );
}

function CalendarIcon() {
  return (
    <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round" strokeLinejoin="round">
      <rect x="3" y="4" width="18" height="18" rx="2" />
      <line x1="3" y1="10" x2="21" y2="10" />
      <line x1="8" y1="2" x2="8" y2="6" />
      <line x1="16" y1="2" x2="16" y2="6" />
    </svg>
  );
}

function SettingsIcon() {
  return (
    <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round" strokeLinejoin="round">
      <circle cx="12" cy="12" r="3" />
      <path d="M19.4 15a1.65 1.65 0 0 0 .33 1.82l.06.06a2 2 0 0 1 0 2.83 2 2 0 0 1-2.83 0l-.06-.06a1.65 1.65 0 0 0-1.82-.33 1.65 1.65 0 0 0-1 1.51V21a2 2 0 0 1-2 2 2 2 0 0 1-2-2v-.09A1.65 1.65 0 0 0 9 19.4a1.65 1.65 0 0 0-1.82.33l-.06.06a2 2 0 0 1-2.83 0 2 2 0 0 1 0-2.83l.06-.06a1.65 1.65 0 0 0 .33-1.82 1.65 1.65 0 0 0-1.51-1H3a2 2 0 0 1-2-2 2 2 0 0 1 2-2h.09A1.65 1.65 0 0 0 4.6 9a1.65 1.65 0 0 0-.33-1.82l-.06-.06a2 2 0 0 1 0-2.83 2 2 0 0 1 2.83 0l.06.06a1.65 1.65 0 0 0 1.82.33H9a1.65 1.65 0 0 0 1-1.51V3a2 2 0 0 1 2-2 2 2 0 0 1 2 2v.09a1.65 1.65 0 0 0 1 1.51 1.65 1.65 0 0 0 1.82-.33l.06-.06a2 2 0 0 1 2.83 0 2 2 0 0 1 0 2.83l-.06.06a1.65 1.65 0 0 0-.33 1.82V9a1.65 1.65 0 0 0 1.51 1H21a2 2 0 0 1 2 2 2 2 0 0 1-2 2h-.09a1.65 1.65 0 0 0-1.51 1z" />
    </svg>
  );
}

function LogoutIcon() {
  return (
    <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round" strokeLinejoin="round">
      <path d="M9 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h4" />
      <polyline points="16 17 21 12 16 7" />
      <line x1="21" y1="12" x2="9" y2="12" />
    </svg>
  );
}

function MoreIcon() {
  return (
    <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round" strokeLinejoin="round">
      <circle cx="12" cy="12" r="1" />
      <circle cx="12" cy="5" r="1" />
      <circle cx="12" cy="19" r="1" />
    </svg>
  );
}

const navItems = [
  { icon: ListIcon, label: 'Listes', path: '/groceries' },
  { icon: ChefHatIcon, label: 'Recettes', path: '/recipes' },
  { icon: CalendarIcon, label: 'Planning', path: '/planning' },
];

interface AppLayoutProps {
  children: React.ReactNode;
  title?: string;
  backTo?: string;
  action?: React.ReactNode;
}

export function AppLayout({ children, title, backTo, action }: AppLayoutProps) {
  const { colorScheme } = useMantineColorScheme();
  const { logout } = useAuth0();
  const location = useLocation();
  const [mobileMenuOpened, { open: openMobileMenu, close: closeMobileMenu }] = useDisclosure(false);

  return (
    <Box className={classes.wrapper} data-theme={colorScheme}>
      {/* Sidebar for desktop */}
      <Box component="aside" className={classes.sidebar}>
        <Box className={classes.sidebarHeader}>
          <Link to="/" className={classes.logoLink}>
            <Box c="coral.5">
              <OctopusLogo size={36} />
            </Box>
          </Link>
        </Box>

        <Box className={classes.sidebarNav}>
          {navItems.map((item) => {
            const isActive = location.pathname.startsWith(item.path);
            return (
              <Tooltip key={item.path} label={item.label} position="right" offset={12}>
                <UnstyledButton
                  component={Link}
                  to={item.path}
                  className={classes.navItem}
                  data-active={isActive || undefined}
                >
                  <item.icon />
                </UnstyledButton>
              </Tooltip>
            );
          })}
        </Box>

        <Box className={classes.sidebarFooter}>
          <Tooltip label="Paramètres" position="right" offset={12}>
            <UnstyledButton
              component={Link}
              to="/settings"
              className={classes.navItem}
              data-active={location.pathname === '/settings' || undefined}
            >
              <SettingsIcon />
            </UnstyledButton>
          </Tooltip>
          <ColorSchemeToggle />
          <Tooltip label="Déconnexion" position="right" offset={12}>
            <ActionIcon
              variant="subtle"
              color="gray"
              size="lg"
              onClick={() => logout({ logoutParams: { returnTo: window.location.origin } })}
            >
              <LogoutIcon />
            </ActionIcon>
          </Tooltip>
        </Box>
      </Box>

      {/* Main content */}
      <Box className={classes.main}>
        {title && (
          <Box className={classes.pageHeader}>
            <Container size="lg">
              <Group justify="space-between" align="center">
                <Group gap="md">
                  {backTo && (
                    <UnstyledButton component={Link} to={backTo} className={classes.backButton}>
                      ←
                    </UnstyledButton>
                  )}
                  <Text className={classes.pageTitle}>{title}</Text>
                </Group>
                {action}
              </Group>
            </Container>
          </Box>
        )}
        <Box className={classes.content}>
          <Container size="lg" py="xl">
            {children}
          </Container>
        </Box>
      </Box>

      {/* Bottom navigation for mobile */}
      <Box component="nav" className={classes.bottomNav}>
        {navItems.map((item) => {
          const isActive = location.pathname.startsWith(item.path);
          return (
            <UnstyledButton
              key={item.path}
              component={Link}
              to={item.path}
              className={classes.bottomNavItem}
              data-active={isActive || undefined}
            >
              <item.icon />
              <Text size="xs" mt={4}>
                {item.label}
              </Text>
            </UnstyledButton>
          );
        })}
        <UnstyledButton
          className={classes.bottomNavItem}
          onClick={openMobileMenu}
        >
          <MoreIcon />
          <Text size="xs" mt={4}>
            Plus
          </Text>
        </UnstyledButton>
      </Box>

      {/* Mobile menu drawer */}
      <Drawer
        opened={mobileMenuOpened}
        onClose={closeMobileMenu}
        position="bottom"
        size="auto"
        title="Menu"
        classNames={{
          content: classes.mobileDrawer,
          header: classes.mobileDrawerHeader,
          title: classes.mobileDrawerTitle,
        }}
      >
        <Stack gap="xs" pb="md">
          <UnstyledButton
            component={Link}
            to="/settings"
            className={classes.mobileMenuItem}
            onClick={closeMobileMenu}
          >
            <SettingsIcon />
            <Text ml="md">Paramètres</Text>
          </UnstyledButton>
          
          <Box className={classes.mobileMenuItem}>
            <ColorSchemeToggle />
            <Text ml="md">Thème</Text>
          </Box>
          
          <UnstyledButton
            className={classes.mobileMenuItem}
            onClick={() => {
              closeMobileMenu();
              logout({ logoutParams: { returnTo: window.location.origin } });
            }}
          >
            <LogoutIcon />
            <Text ml="md">Déconnexion</Text>
          </UnstyledButton>
        </Stack>
      </Drawer>
    </Box>
  );
}
