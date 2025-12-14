import {
  Box,
  Button,
  Container,
  Group,
  SimpleGrid,
  Stack,
  Text,
  ThemeIcon,
  Title,
  useMantineColorScheme,
} from '@mantine/core';
import { ColorSchemeToggle } from '../components/ColorSchemeToggle/ColorSchemeToggle';
import classes from './Home.module.css';

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

function RecipeIcon() {
  return (
    <svg width="32" height="32" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
      <path
        d="M4 6H20M4 10H20M4 14H14M4 18H10"
        stroke="currentColor"
        strokeWidth="2"
        strokeLinecap="round"
      />
      <path
        d="M18 14L20 16L18 18"
        stroke="currentColor"
        strokeWidth="2"
        strokeLinecap="round"
        strokeLinejoin="round"
      />
    </svg>
  );
}

function CalendarIcon() {
  return (
    <svg width="32" height="32" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
      <rect
        x="3"
        y="4"
        width="18"
        height="18"
        rx="2"
        stroke="currentColor"
        strokeWidth="2"
      />
      <path d="M3 10H21" stroke="currentColor" strokeWidth="2" />
      <path d="M8 2V6" stroke="currentColor" strokeWidth="2" strokeLinecap="round" />
      <path d="M16 2V6" stroke="currentColor" strokeWidth="2" strokeLinecap="round" />
      <circle cx="8" cy="15" r="1.5" fill="currentColor" />
      <circle cx="12" cy="15" r="1.5" fill="currentColor" />
      <circle cx="16" cy="15" r="1.5" fill="currentColor" />
    </svg>
  );
}

function CartIcon() {
  return (
    <svg width="32" height="32" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
      <path
        d="M6 6H22L20 16H8L6 6Z"
        stroke="currentColor"
        strokeWidth="2"
        strokeLinejoin="round"
      />
      <path d="M6 6L5 3H2" stroke="currentColor" strokeWidth="2" strokeLinecap="round" />
      <circle cx="10" cy="20" r="2" fill="currentColor" />
      <circle cx="18" cy="20" r="2" fill="currentColor" />
    </svg>
  );
}

const features = [
  {
    icon: RecipeIcon,
    title: 'Stockage de recettes',
    description:
      'Stockez et organisez facilement vos recettes préférées en un seul endroit. Ajoutez des recettes du web ou saisissez-les manuellement.',
    color: 'coral',
  },
  {
    icon: CalendarIcon,
    title: 'Planification de repas',
    description:
      'Planifiez vos repas pour la semaine avec la possibilité d\'attribuer des recettes à des jours et repas spécifiques.',
    color: 'sage',
  },
  {
    icon: CartIcon,
    title: 'Liste d\'épicerie',
    description:
      'OctoPlanner génère automatiquement votre liste d\'épicerie avec les ingrédients nécessaires, en tenant compte de ce que vous avez déjà.',
    color: 'coral',
  },
];

function FeatureCard({
  icon: Icon,
  title,
  description,
  color,
  delay,
}: {
  icon: () => React.ReactElement;
  title: string;
  description: string;
  color: string;
  delay: number;
}) {
  return (
    <Box className={classes.featureCard} style={{ animationDelay: `${delay}ms` }}>
      <ThemeIcon size={64} radius="xl" variant="light" color={color} mb="lg">
        <Icon />
      </ThemeIcon>
      <Title order={3} mb="sm" className={classes.featureTitle}>
        {title}
      </Title>
      <Text c="dimmed" size="md" lh={1.6}>
        {description}
      </Text>
    </Box>
  );
}

export function HomePage() {
  const { colorScheme } = useMantineColorScheme();
  const isDark = colorScheme === 'dark';

  return (
    <Box className={classes.wrapper} data-theme={colorScheme}>
      {/* Background decoration */}
      <Box className={classes.bgDecoration} />
      <Box className={classes.bgDecoration2} />

      {/* Header */}
      <Box component="header" className={classes.header}>
        <Container size="xl">
          <Group justify="space-between" h={70}>
            <Group gap="xs">
              <Box c="coral.5">
                <OctopusLogo size={40} />
              </Box>
              <Text
                fw={700}
                size="xl"
                style={{ letterSpacing: '-0.5px' }}
              >
                OctoPlanner
              </Text>
            </Group>
            <Group>
              <ColorSchemeToggle />
              <Button variant="light" color="coral" radius="xl">
                Connexion
              </Button>
              <Button color="coral" radius="xl">
                Commencer
              </Button>
            </Group>
          </Group>
        </Container>
      </Box>

      {/* Hero Section */}
      <Container size="lg" className={classes.heroContainer}>
        <Stack align="center" gap="xl" className={classes.hero}>
          <Box className={classes.badge}>
            <Text size="sm" fw={600}>
              🧑🏼‍🍳 Cuisiner n'a jamais été aussi simple
            </Text>
          </Box>

          <Title className={classes.title} ta="center">
            Planifiez vos repas,
            <br />
            <Text
              component="span"
              inherit
              variant="gradient"
              gradient={{ from: 'coral.5', to: 'coral.7', deg: 135 }}
            >
              simplifiez votre vie
            </Text>
          </Title>

          <Text
            c="dimmed"
            size="xl"
            maw={600}
            ta="center"
            lh={1.7}
            className={classes.subtitle}
          >
            OctoPlanner est votre assistant culinaire. Stockez vos recettes, planifiez vos repas 
            et générez votre liste d'épicerie en quelques clics.
          </Text>

          <Group mt="md">
            <Button
              size="xl"
              radius="xl"
              color="coral"
              className={classes.ctaButton}
            >
              Commencer gratuitement
            </Button>
            <Button
              size="xl"
              radius="xl"
              variant="outline"
              color={isDark ? 'gray.4' : 'charcoal.9'}
            >
              Découvrir
            </Button>
          </Group>
        </Stack>

        {/* Hero Visual */}
        <Box className={classes.heroVisual}>
          <Box className={classes.floatingCard} style={{ '--delay': '0s' } as React.CSSProperties}>
            <Text size="xl" mb="xs">📅</Text>
            <Text size="sm" fw={600}>Lundi</Text>
            <Text size="xs" c="dimmed">Poulet rôti</Text>
          </Box>
          <Box className={classes.floatingCard} style={{ '--delay': '0.2s', '--x': '120px', '--y': '60px' } as React.CSSProperties}>
            <Text size="xl" mb="xs">🥗</Text>
            <Text size="sm" fw={600}>Mardi</Text>
            <Text size="xs" c="dimmed">Salade César</Text>
          </Box>
          <Box className={classes.floatingCard} style={{ '--delay': '0.4s', '--x': '-100px', '--y': '120px' } as React.CSSProperties}>
            <Text size="xl" mb="xs">🍝</Text>
            <Text size="sm" fw={600}>Mercredi</Text>
            <Text size="xs" c="dimmed">Spaghetti</Text>
          </Box>
          <Box className={classes.floatingCard} style={{ '--delay': '0.6s', '--x': '80px', '--y': '180px' } as React.CSSProperties}>
            <Text size="xl" mb="xs">🛒</Text>
            <Text size="sm" fw={600}>Liste</Text>
            <Text size="xs" c="dimmed">8 articles</Text>
          </Box>
        </Box>
      </Container>

      {/* Features Section */}
      <Box className={classes.featuresSection}>
        <Container size="lg">
          <Stack align="center" mb={60}>
            <Text c="coral.5" fw={700} size="sm" tt="uppercase" style={{ letterSpacing: '2px' }}>
              Fonctionnalités
            </Text>
            <Title order={2} ta="center" className={classes.sectionTitle}>
              Tout ce dont vous avez besoin
              <br />
              pour cuisiner sereinement
            </Title>
          </Stack>

          <SimpleGrid cols={{ base: 1, sm: 2, md: 3 }} spacing={40}>
            {features.map((feature, index) => (
              <FeatureCard
                key={feature.title}
                {...feature}
                delay={index * 150}
              />
            ))}
          </SimpleGrid>
        </Container>
      </Box>

      {/* CTA Section */}
      <Box className={classes.ctaSection}>
        <Container size="md">
          <Box className={classes.ctaCard}>
            <Box c="coral.5" mb="lg">
              <OctopusLogo size={64} />
            </Box>
            <Title order={2} ta="center" mb="md" c="white">
              Prêt à simplifier vos repas ?
            </Title>
            <Text c="gray.4" ta="center" size="lg" maw={500} mx="auto" mb="xl">
              Rejoignez des milliers de cuisiniers qui ont déjà adopté OctoPlanner
              pour organiser leur semaine culinaire.
            </Text>
            <Button size="xl" radius="xl" color="white" c="coral.6">
              Créer un compte gratuit
            </Button>
          </Box>
        </Container>
      </Box>

      {/* Footer */}
      <Box component="footer" className={classes.footer}>
        <Container size="xl">
          <Group justify="space-between">
            <Group gap="xs">
              <Box c="coral.5">
                <OctopusLogo size={32} />
              </Box>
              <Text fw={600}>OctoPlanner</Text>
            </Group>
            <Text size="sm" c="dimmed">
              © 2024 OctoPlanner. Fait avec ❤️ pour les amoureux de la cuisine.
            </Text>
          </Group>
        </Container>
      </Box>
    </Box>
  );
}
