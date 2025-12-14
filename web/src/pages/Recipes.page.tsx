import { useState } from 'react';
import {
  ActionIcon,
  Badge,
  Box,
  Button,
  Group,
  Loader,
  Modal,
  SimpleGrid,
  Stack,
  Text,
  Textarea,
  TextInput,
  ThemeIcon,
  UnstyledButton,
  useMantineColorScheme,
} from '@mantine/core';
import { useDisclosure } from '@mantine/hooks';
import { Link, useNavigate } from 'react-router-dom';
import { AppLayout } from '../components/AppLayout/AppLayout';
import { useRecipes, useCreateRecipe, useDeleteRecipe } from '../hooks/use-recipes';
import { Recipe } from '../gen/octoback/recipes/v1/recipes_pb';
import { useQueryClient } from '@tanstack/react-query';
import classes from './Recipes.module.css';

function ChefHatIcon({ size = 48 }: { size?: number }) {
  return (
    <svg width={size} height={size} viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round" strokeLinejoin="round">
      <path d="M6 13.87A4 4 0 0 1 7.41 6.6a5.11 5.11 0 0 1 1.05-1.11 5 5 0 0 1 7.08 0A5.11 5.11 0 0 1 16.59 6.6 4 4 0 0 1 18 13.87V21H6Z" />
      <line x1="6" y1="17" x2="18" y2="17" />
    </svg>
  );
}

function PlusIcon() {
  return (
    <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round">
      <line x1="12" y1="5" x2="12" y2="19" />
      <line x1="5" y1="12" x2="19" y2="12" />
    </svg>
  );
}

function TrashIcon() {
  return (
    <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round" strokeLinejoin="round">
      <polyline points="3 6 5 6 21 6" />
      <path d="M19 6v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V6m3 0V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2" />
    </svg>
  );
}

function ClockIcon() {
  return (
    <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round">
      <circle cx="12" cy="12" r="10" />
      <polyline points="12 6 12 12 16 14" />
    </svg>
  );
}

function ServingsIcon() {
  return (
    <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round">
      <path d="M17 21v-2a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v2" />
      <circle cx="9" cy="7" r="4" />
      <path d="M23 21v-2a4 4 0 0 0-3-3.87" />
      <path d="M16 3.13a4 4 0 0 1 0 7.75" />
    </svg>
  );
}

function EmptyState() {
  const { colorScheme } = useMantineColorScheme();

  return (
    <Box className={classes.emptyState} data-theme={colorScheme}>
      <Box className={classes.emptyIcon}>
        <ChefHatIcon size={64} />
      </Box>
      <Text className={classes.emptyTitle}>Aucune recette</Text>
      <Text c="dimmed" size="lg" maw={400} ta="center" lh={1.6}>
        Ajoutez votre première recette pour commencer à construire votre livre de cuisine personnel.
      </Text>
    </Box>
  );
}

function formatTime(minutes: number): string {
  if (minutes < 60) {
    return `${minutes} min`;
  }
  const hours = Math.floor(minutes / 60);
  const mins = minutes % 60;
  return mins > 0 ? `${hours}h ${mins}min` : `${hours}h`;
}

interface RecipeCardProps {
  recipe: Recipe;
  onDelete: (id: bigint) => void;
  delay: number;
}

function RecipeCard({ recipe, onDelete, delay }: RecipeCardProps) {
  const { colorScheme } = useMantineColorScheme();
  const [isHovered, setIsHovered] = useState(false);

  const totalTime = recipe.prepTimeMinutes + recipe.cookTimeMinutes;

  return (
    <UnstyledButton
      component={Link}
      to={`/recipes/${recipe.id}`}
      className={classes.recipeCard}
      data-theme={colorScheme}
      style={{ animationDelay: `${delay}ms` }}
      onMouseEnter={() => setIsHovered(true)}
      onMouseLeave={() => setIsHovered(false)}
    >
      <Box className={classes.cardPattern} />

      <Group justify="space-between" align="flex-start" mb="lg" style={{ position: 'relative', zIndex: 1 }}>
        <ThemeIcon
          size={48}
          radius="xl"
          variant="light"
          color="sage"
          className={classes.cardIcon}
        >
          <ChefHatIcon size={24} />
        </ThemeIcon>

        {isHovered && (
          <ActionIcon
            variant="subtle"
            color="red"
            size="sm"
            onClick={(e) => {
              e.preventDefault();
              e.stopPropagation();
              onDelete(recipe.id);
            }}
            className={classes.deleteButton}
          >
            <TrashIcon />
          </ActionIcon>
        )}
      </Group>

      <Stack gap="xs" style={{ position: 'relative', zIndex: 1 }}>
        <Text className={classes.cardTitle}>{recipe.name}</Text>
        {recipe.description && (
          <Text className={classes.cardDescription} lineClamp={2}>
            {recipe.description}
          </Text>
        )}
        <Group gap="md" mt="xs">
          {totalTime > 0 && (
            <Badge
              variant="light"
              color="sage"
              size="sm"
              leftSection={<ClockIcon />}
              className={classes.metaBadge}
            >
              {formatTime(totalTime)}
            </Badge>
          )}
          {recipe.servings > 0 && (
            <Badge
              variant="light"
              color="coral"
              size="sm"
              leftSection={<ServingsIcon />}
              className={classes.metaBadge}
            >
              {recipe.servings} pers.
            </Badge>
          )}
        </Group>
      </Stack>

      <Box className={classes.cardArrow}>→</Box>
    </UnstyledButton>
  );
}

export function RecipesPage() {
  const { colorScheme } = useMantineColorScheme();
  const [opened, { open, close }] = useDisclosure(false);
  const [newRecipeName, setNewRecipeName] = useState('');
  const [newRecipeDescription, setNewRecipeDescription] = useState('');
  const navigate = useNavigate();

  const { data, isLoading } = useRecipes();
  const createMutation = useCreateRecipe();
  const deleteMutation = useDeleteRecipe();
  const queryClient = useQueryClient();

  const recipes = data?.recipes ?? [];

  const handleCreate = async () => {
    if (!newRecipeName.trim()) return;

    const result = await createMutation.mutateAsync({
      name: newRecipeName.trim(),
      description: newRecipeDescription.trim(),
      servings: 0,
      prepTimeMinutes: 0,
      cookTimeMinutes: 0,
    });
    queryClient.invalidateQueries({ queryKey: ['octoback.recipes.v1.Service', 'ListRecipes'] });
    resetForm();
    close();
    
    if (result.recipe) {
      navigate(`/recipes/${result.recipe.id}`);
    }
  };

  const handleDelete = async (id: bigint) => {
    await deleteMutation.mutateAsync({ id });
    queryClient.invalidateQueries({ queryKey: ['octoback.recipes.v1.Service', 'ListRecipes'] });
  };

  const resetForm = () => {
    setNewRecipeName('');
    setNewRecipeDescription('');
  };

  return (
    <AppLayout
      title="Mes recettes"
      action={
        <Button
          leftSection={<PlusIcon />}
          color="sage"
          radius="xl"
          onClick={open}
          className={classes.createButton}
        >
          Nouvelle recette
        </Button>
      }
    >
      {isLoading ? (
        <Box className={classes.loadingState}>
          <Loader color="sage" size="lg" />
        </Box>
      ) : recipes.length === 0 ? (
        <EmptyState />
      ) : (
        <SimpleGrid
          cols={{ base: 1, xs: 2, md: 3 }}
          spacing={{ base: 'md', md: 'lg' }}
          className={classes.grid}
        >
          {recipes.map((recipe, index) => (
            <RecipeCard
              key={recipe.id.toString()}
              recipe={recipe}
              onDelete={handleDelete}
              delay={index * 50}
            />
          ))}

          {/* Add new recipe card */}
          <UnstyledButton
            className={classes.addCard}
            data-theme={colorScheme}
            onClick={open}
          >
            <Box className={classes.addIcon}>
              <PlusIcon />
            </Box>
            <Text size="sm" c="dimmed" fw={500}>
              Ajouter une recette
            </Text>
          </UnstyledButton>
        </SimpleGrid>
      )}

      {/* Create Modal */}
      <Modal
        opened={opened}
        onClose={() => {
          close();
          resetForm();
        }}
        title="Nouvelle recette"
        centered
        radius="lg"
        size="md"
        classNames={{
          title: classes.modalTitle,
          content: classes.modalContent,
        }}
      >
        <Stack gap="md">
          <TextInput
            label="Nom de la recette"
            placeholder="Ex: Poulet rôti aux herbes"
            value={newRecipeName}
            onChange={(e) => setNewRecipeName(e.currentTarget.value)}
            size="md"
            radius="md"
            autoFocus
            required
          />
          <Textarea
            label="Description"
            placeholder="Une brève description de votre recette..."
            value={newRecipeDescription}
            onChange={(e) => setNewRecipeDescription(e.currentTarget.value)}
            size="md"
            radius="md"
            rows={3}
          />
          <Group justify="flex-end" gap="sm" mt="md">
            <Button
              variant="subtle"
              color="gray"
              onClick={() => {
                close();
                resetForm();
              }}
              radius="md"
            >
              Annuler
            </Button>
            <Button
              color="sage"
              onClick={handleCreate}
              loading={createMutation.isPending}
              radius="md"
            >
              Créer
            </Button>
          </Group>
        </Stack>
      </Modal>
    </AppLayout>
  );
}

