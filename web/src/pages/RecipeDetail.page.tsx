import { useState, useRef, useEffect } from 'react';
import {
  ActionIcon,
  Badge,
  Box,
  Button,
  Group,
  Loader,
  Modal,
  NumberInput,
  Paper,
  Stack,
  Text,
  Textarea,
  TextInput,
  ThemeIcon,
  Title,
  Transition,
  UnstyledButton,
  useMantineColorScheme,
} from '@mantine/core';
import { useDisclosure } from '@mantine/hooks';
import { useParams, useNavigate } from 'react-router-dom';
import { AppLayout } from '../components/AppLayout/AppLayout';
import {
  useRecipe,
  useRecipeIngredients,
  useCreateRecipeIngredient,
  useUpdateRecipe,
  useUpdateRecipeIngredient,
  useDeleteRecipeIngredient,
} from '../hooks/use-recipes';
import { RecipeIngredient } from '../gen/octoback/recipes/v1/recipes_pb';
import { useQueryClient } from '@tanstack/react-query';
import { FieldMask } from '@bufbuild/protobuf';
import classes from './RecipeDetail.module.css';

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
    <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round" strokeLinejoin="round">
      <polyline points="3 6 5 6 21 6" />
      <path d="M19 6v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V6m3 0V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2" />
    </svg>
  );
}

function EditIcon() {
  return (
    <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round">
      <path d="M11 4H4a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-7" />
      <path d="M18.5 2.5a2.121 2.121 0 0 1 3 3L12 15l-4 1 1-4 9.5-9.5z" />
    </svg>
  );
}

function ClockIcon() {
  return (
    <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round">
      <circle cx="12" cy="12" r="10" />
      <polyline points="12 6 12 12 16 14" />
    </svg>
  );
}

function ServingsIcon() {
  return (
    <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round">
      <path d="M17 21v-2a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v2" />
      <circle cx="9" cy="7" r="4" />
      <path d="M23 21v-2a4 4 0 0 0-3-3.87" />
      <path d="M16 3.13a4 4 0 0 1 0 7.75" />
    </svg>
  );
}

function IngredientIcon() {
  return (
    <svg width="80" height="80" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="1" strokeLinecap="round" strokeLinejoin="round">
      <path d="M12 2L2 7l10 5 10-5-10-5z" />
      <path d="M2 17l10 5 10-5" />
      <path d="M2 12l10 5 10-5" />
    </svg>
  );
}

function ChefHatIcon() {
  return (
    <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round" strokeLinejoin="round">
      <path d="M6 13.87A4 4 0 0 1 7.41 6.6a5.11 5.11 0 0 1 1.05-1.11 5 5 0 0 1 7.08 0A5.11 5.11 0 0 1 16.59 6.6 4 4 0 0 1 18 13.87V21H6Z" />
      <line x1="6" y1="17" x2="18" y2="17" />
    </svg>
  );
}

function EmptyState() {
  const { colorScheme } = useMantineColorScheme();

  return (
    <Box className={classes.emptyState} data-theme={colorScheme}>
      <Box className={classes.emptyIcon}>
        <IngredientIcon />
      </Box>
      <Text className={classes.emptyTitle}>Aucun ingrédient</Text>
      <Text c="dimmed" size="md" maw={320} ta="center" lh={1.6}>
        Ajoutez les ingrédients nécessaires à cette recette pour ne rien oublier lors de vos courses.
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

interface IngredientRowProps {
  ingredient: RecipeIngredient;
  onEdit: (ingredient: RecipeIngredient) => void;
  onDelete: (id: bigint) => void;
  delay: number;
}

function IngredientRow({ ingredient, onEdit, onDelete, delay }: IngredientRowProps) {
  const { colorScheme } = useMantineColorScheme();
  const [isHovered, setIsHovered] = useState(false);

  const quantityDisplay = [ingredient.quantity, ingredient.unit].filter(Boolean).join(' ');

  return (
    <UnstyledButton
      className={classes.ingredientRowButton}
      onClick={() => onEdit(ingredient)}
    >
      <Paper
        className={classes.ingredientRow}
        data-theme={colorScheme}
        style={{ animationDelay: `${delay}ms` }}
        onMouseEnter={() => setIsHovered(true)}
        onMouseLeave={() => setIsHovered(false)}
      >
        <Group gap="md" wrap="nowrap">
          <Box className={classes.ingredientBullet} />
          <Box className={classes.ingredientContent}>
            <Text className={classes.ingredientName}>{ingredient.name}</Text>
            {quantityDisplay && (
              <Text className={classes.ingredientQuantity}>{quantityDisplay}</Text>
            )}
          </Box>

          <Transition mounted={isHovered} transition="fade" duration={150}>
            {(styles) => (
              <ActionIcon
                variant="subtle"
                color="red"
                size="md"
                style={styles}
                onClick={(e) => {
                  e.stopPropagation();
                  onDelete(ingredient.id);
                }}
                className={classes.deleteButton}
              >
                <TrashIcon />
              </ActionIcon>
            )}
          </Transition>
        </Group>
      </Paper>
    </UnstyledButton>
  );
}

export function RecipeDetailPage() {
  const { colorScheme } = useMantineColorScheme();
  const { id } = useParams<{ id: string }>();
  const navigate = useNavigate();
  const queryClient = useQueryClient();
  const inputRef = useRef<HTMLInputElement>(null);

  const [newIngredientName, setNewIngredientName] = useState('');
  const [newIngredientQuantity, setNewIngredientQuantity] = useState('');
  const [newIngredientUnit, setNewIngredientUnit] = useState('');

  const [editOpened, { open: openEdit, close: closeEdit }] = useDisclosure(false);
  const [editName, setEditName] = useState('');
  const [editDescription, setEditDescription] = useState('');
  const [editServings, setEditServings] = useState<number | ''>(0);
  const [editPrepTime, setEditPrepTime] = useState<number | ''>(0);
  const [editCookTime, setEditCookTime] = useState<number | ''>(0);

  // Edit ingredient state
  const [editIngredientOpened, { open: openEditIngredient, close: closeEditIngredient }] = useDisclosure(false);
  const [editingIngredient, setEditingIngredient] = useState<RecipeIngredient | null>(null);
  const [editIngredientName, setEditIngredientName] = useState('');
  const [editIngredientQuantity, setEditIngredientQuantity] = useState('');
  const [editIngredientUnit, setEditIngredientUnit] = useState('');

  const recipeId = BigInt(id ?? '0');

  const { data: recipeData, isLoading: recipeLoading } = useRecipe(recipeId);
  const { data: ingredientsData, isLoading: ingredientsLoading } = useRecipeIngredients(recipeId);
  const createIngredientMutation = useCreateRecipeIngredient();
  const updateRecipeMutation = useUpdateRecipe();
  const updateIngredientMutation = useUpdateRecipeIngredient();
  const deleteIngredientMutation = useDeleteRecipeIngredient();

  const recipe = recipeData?.recipe;
  const ingredients = ingredientsData?.recipeIngredients ?? [];

  useEffect(() => {
    if (!recipeLoading && !recipe) {
      navigate('/recipes');
    }
  }, [recipeLoading, recipe, navigate]);

  useEffect(() => {
    if (recipe) {
      setEditName(recipe.name);
      setEditDescription(recipe.description);
      setEditServings(recipe.servings || '');
      setEditPrepTime(recipe.prepTimeMinutes || '');
      setEditCookTime(recipe.cookTimeMinutes || '');
    }
  }, [recipe]);

  const handleAddIngredient = async () => {
    if (!newIngredientName.trim()) return;

    await createIngredientMutation.mutateAsync({
      recipe: recipeId,
      name: newIngredientName.trim(),
      quantity: newIngredientQuantity.trim(),
      unit: newIngredientUnit.trim(),
    });

    queryClient.invalidateQueries({
      queryKey: ['octoback.recipes.v1.Service', 'ListRecipeIngredients'],
    });

    setNewIngredientName('');
    setNewIngredientQuantity('');
    setNewIngredientUnit('');
    inputRef.current?.focus();
  };

  const handleDeleteIngredient = async (ingredientId: bigint) => {
    await deleteIngredientMutation.mutateAsync({ id: ingredientId });
    queryClient.invalidateQueries({
      queryKey: ['octoback.recipes.v1.Service', 'ListRecipeIngredients'],
    });
  };

  const handleOpenEditIngredient = (ingredient: RecipeIngredient) => {
    setEditingIngredient(ingredient);
    setEditIngredientName(ingredient.name);
    setEditIngredientQuantity(ingredient.quantity);
    setEditIngredientUnit(ingredient.unit);
    openEditIngredient();
  };

  const handleUpdateIngredient = async () => {
    if (!editIngredientName.trim() || !editingIngredient) {
      return;
    }

    await updateIngredientMutation.mutateAsync({
      recipeIngredient: {
        id: editingIngredient.id,
        recipe: editingIngredient.recipe,
        name: editIngredientName.trim(),
        quantity: editIngredientQuantity.trim(),
        unit: editIngredientUnit.trim(),
        version: editingIngredient.version,
      },
      updateMask: new FieldMask({
        paths: ['name', 'quantity', 'unit'],
      }),
    });

    queryClient.invalidateQueries({
      queryKey: ['octoback.recipes.v1.Service', 'ListRecipeIngredients'],
    });

    closeEditIngredient();
    setEditingIngredient(null);
  };

  const handleUpdateRecipe = async () => {
    if (!editName.trim() || !recipe) return;

    await updateRecipeMutation.mutateAsync({
      recipe: {
        id: recipeId,
        name: editName.trim(),
        description: editDescription.trim(),
        servings: Number(editServings) || 0,
        prepTimeMinutes: Number(editPrepTime) || 0,
        cookTimeMinutes: Number(editCookTime) || 0,
        version: recipe.version,
      },
      updateMask: new FieldMask({
        paths: ['name', 'description', 'servings', 'prep_time_minutes', 'cook_time_minutes'],
      }),
    });

    queryClient.invalidateQueries({
      queryKey: ['octoback.recipes.v1.Service', 'GetRecipe'],
    });

    closeEdit();
  };

  const isLoading = recipeLoading || ingredientsLoading;
  const totalTime = (recipe?.prepTimeMinutes ?? 0) + (recipe?.cookTimeMinutes ?? 0);

  return (
    <AppLayout
      title={recipe?.name ?? 'Chargement...'}
      backTo="/recipes"
      action={
        recipe && (
          <ActionIcon
            variant="light"
            color="sage"
            size="lg"
            radius="xl"
            onClick={openEdit}
          >
            <EditIcon />
          </ActionIcon>
        )
      }
    >
      <Box className={classes.wrapper} data-theme={colorScheme}>
        {/* Recipe info header */}
        {recipe && (
          <Paper className={classes.recipeHeader} data-theme={colorScheme}>
            <Group gap="xl" align="flex-start">
              <ThemeIcon size={56} radius="xl" variant="light" color="sage">
                <ChefHatIcon />
              </ThemeIcon>
              <Box style={{ flex: 1 }}>
                <Title order={2} className={classes.recipeTitle}>
                  {recipe.name}
                </Title>
                {recipe.description && (
                  <Text c="dimmed" mt="xs" lh={1.6}>
                    {recipe.description}
                  </Text>
                )}
                <Group gap="lg" mt="md">
                  {totalTime > 0 && (
                    <Group gap="xs">
                      <ClockIcon />
                      <Text size="sm" fw={500}>
                        {formatTime(totalTime)}
                      </Text>
                      <Text size="xs" c="dimmed">
                        ({formatTime(recipe.prepTimeMinutes)} prépa + {formatTime(recipe.cookTimeMinutes)} cuisson)
                      </Text>
                    </Group>
                  )}
                  {recipe.servings > 0 && (
                    <Group gap="xs">
                      <ServingsIcon />
                      <Text size="sm" fw={500}>
                        {recipe.servings} portions
                      </Text>
                    </Group>
                  )}
                </Group>
              </Box>
            </Group>
          </Paper>
        )}

        {/* Ingredients section */}
        <Box mt="xl">
          <Group justify="space-between" mb="md">
            <Title order={3} className={classes.sectionTitle}>
              Ingrédients
            </Title>
            <Badge variant="light" color="sage" size="lg">
              {ingredients.length} article{ingredients.length !== 1 ? 's' : ''}
            </Badge>
          </Group>

          {/* Add ingredient form */}
          <Paper className={classes.addForm} data-theme={colorScheme}>
            <Group gap="sm" wrap="nowrap">
              <TextInput
                ref={inputRef}
                placeholder="Nom de l'ingrédient..."
                value={newIngredientName}
                onChange={(e) => setNewIngredientName(e.currentTarget.value)}
                onKeyDown={(e) => {
                  if (e.key === 'Enter') handleAddIngredient();
                }}
                className={classes.addInputMain}
                size="md"
                radius="md"
                styles={{
                  input: {
                    border: 'none',
                    background: 'transparent',
                  },
                }}
              />
              <TextInput
                placeholder="Qté"
                value={newIngredientQuantity}
                onChange={(e) => setNewIngredientQuantity(e.currentTarget.value)}
                className={classes.addInputSmall}
                size="md"
                radius="md"
                w={80}
                styles={{
                  input: {
                    textAlign: 'center',
                  },
                }}
              />
              <TextInput
                placeholder="Unité"
                value={newIngredientUnit}
                onChange={(e) => setNewIngredientUnit(e.currentTarget.value)}
                className={classes.addInputSmall}
                size="md"
                radius="md"
                w={80}
                styles={{
                  input: {
                    textAlign: 'center',
                  },
                }}
              />
              <Button
                color="sage"
                radius="md"
                onClick={handleAddIngredient}
                loading={createIngredientMutation.isPending}
                className={classes.addButton}
              >
                <PlusIcon />
              </Button>
            </Group>
          </Paper>

          {/* Ingredients list */}
          {isLoading ? (
            <Box className={classes.loadingState}>
              <Loader color="sage" size="lg" />
            </Box>
          ) : ingredients.length === 0 ? (
            <EmptyState />
          ) : (
            <Stack gap="sm" className={classes.ingredientsList}>
              {ingredients.map((ingredient, index) => (
                <IngredientRow
                  key={ingredient.id.toString()}
                  ingredient={ingredient}
                  onEdit={handleOpenEditIngredient}
                  onDelete={handleDeleteIngredient}
                  delay={index * 30}
                />
              ))}
            </Stack>
          )}
        </Box>
      </Box>

      {/* Edit Recipe Modal */}
      <Modal
        opened={editOpened}
        onClose={closeEdit}
        title="Modifier la recette"
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
            value={editName}
            onChange={(e) => setEditName(e.currentTarget.value)}
            size="md"
            radius="md"
            required
          />
          <Textarea
            label="Description"
            placeholder="Une brève description de votre recette..."
            value={editDescription}
            onChange={(e) => setEditDescription(e.currentTarget.value)}
            size="md"
            radius="md"
            rows={3}
          />
          <Group grow>
            <NumberInput
              label="Portions"
              placeholder="4"
              value={editServings}
              onChange={(value) => setEditServings(value === '' ? '' : Number(value))}
              min={1}
              max={50}
              size="md"
              radius="md"
            />
            <NumberInput
              label="Préparation (min)"
              placeholder="15"
              value={editPrepTime}
              onChange={(value) => setEditPrepTime(value === '' ? '' : Number(value))}
              min={0}
              max={480}
              size="md"
              radius="md"
            />
            <NumberInput
              label="Cuisson (min)"
              placeholder="30"
              value={editCookTime}
              onChange={(value) => setEditCookTime(value === '' ? '' : Number(value))}
              min={0}
              max={480}
              size="md"
              radius="md"
            />
          </Group>
          <Group justify="flex-end" gap="sm" mt="md">
            <Button variant="subtle" color="gray" onClick={closeEdit} radius="md">
              Annuler
            </Button>
            <Button
              color="sage"
              onClick={handleUpdateRecipe}
              loading={updateRecipeMutation.isPending}
              radius="md"
            >
              Enregistrer
            </Button>
          </Group>
        </Stack>
      </Modal>

      {/* Edit Ingredient Modal */}
      <Modal
        opened={editIngredientOpened}
        onClose={closeEditIngredient}
        title="Modifier l'ingrédient"
        centered
        radius="lg"
        size="sm"
        classNames={{
          title: classes.modalTitle,
          content: classes.modalContent,
        }}
      >
        <Stack gap="md">
          <TextInput
            label="Nom de l'ingrédient"
            placeholder="Ex: Poulet"
            value={editIngredientName}
            onChange={(e) => setEditIngredientName(e.currentTarget.value)}
            size="md"
            radius="md"
            required
          />
          <Group grow>
            <TextInput
              label="Quantité"
              placeholder="Ex: 500"
              value={editIngredientQuantity}
              onChange={(e) => setEditIngredientQuantity(e.currentTarget.value)}
              size="md"
              radius="md"
            />
            <TextInput
              label="Unité"
              placeholder="Ex: g"
              value={editIngredientUnit}
              onChange={(e) => setEditIngredientUnit(e.currentTarget.value)}
              size="md"
              radius="md"
            />
          </Group>
          <Group justify="flex-end" gap="sm" mt="md">
            <Button variant="subtle" color="gray" onClick={closeEditIngredient} radius="md">
              Annuler
            </Button>
            <Button
              color="sage"
              onClick={handleUpdateIngredient}
              loading={updateIngredientMutation.isPending}
              radius="md"
            >
              Enregistrer
            </Button>
          </Group>
        </Stack>
      </Modal>
    </AppLayout>
  );
}

