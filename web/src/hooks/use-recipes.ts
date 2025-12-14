import { useMutation, useQuery } from '@connectrpc/connect-query';
import {
  createRecipe,
  createRecipeIngredient,
  deleteRecipe,
  deleteRecipeIngredient,
  getRecipe,
  getRecipeIngredient,
  listRecipeIngredients,
  listRecipes,
  updateRecipe,
  updateRecipeIngredient,
} from '../gen/octoback/recipes/v1/service-Service_connectquery';

// ============ Query Hooks ============

/**
 * Fetch a single recipe by ID
 */
export function useRecipe(id: bigint) {
  return useQuery(getRecipe, { id });
}

/**
 * Fetch all recipes
 */
export function useRecipes() {
  return useQuery(listRecipes, {});
}

/**
 * Fetch a single recipe ingredient by ID
 */
export function useRecipeIngredient(id: bigint) {
  return useQuery(getRecipeIngredient, { id });
}

/**
 * Fetch all ingredients for a specific recipe
 */
export function useRecipeIngredients(recipeId: bigint) {
  return useQuery(listRecipeIngredients, { recipe: recipeId });
}

// ============ Mutation Hooks ============

/**
 * Create a new recipe
 */
export function useCreateRecipe() {
  return useMutation(createRecipe);
}

/**
 * Update an existing recipe
 */
export function useUpdateRecipe() {
  return useMutation(updateRecipe);
}

/**
 * Delete a recipe
 */
export function useDeleteRecipe() {
  return useMutation(deleteRecipe);
}

/**
 * Create a new recipe ingredient
 */
export function useCreateRecipeIngredient() {
  return useMutation(createRecipeIngredient);
}

/**
 * Update an existing recipe ingredient
 */
export function useUpdateRecipeIngredient() {
  return useMutation(updateRecipeIngredient);
}

/**
 * Delete a recipe ingredient
 */
export function useDeleteRecipeIngredient() {
  return useMutation(deleteRecipeIngredient);
}

