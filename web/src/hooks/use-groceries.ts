import { useMutation, useQuery } from '@connectrpc/connect-query';
import {
  createGroceryItem,
  createGroceryList,
  deleteGroceryItem,
  deleteGroceryList,
  getGroceryItem,
  getGroceryList,
  listGroceryItems,
  listGroceryLists,
  updateGroceryItem,
  updateGroceryList,
} from '../gen/octoback/groceries/v1/service-Service_connectquery';

// ============ Query Hooks ============

/**
 * Fetch a single grocery list by ID
 */
export function useGroceryList(id: bigint) {
  return useQuery(getGroceryList, { id });
}

/**
 * Fetch all grocery lists
 */
export function useGroceryLists() {
  return useQuery(listGroceryLists, {});
}

/**
 * Fetch a single grocery item by ID
 */
export function useGroceryItem(id: bigint) {
  return useQuery(getGroceryItem, { id });
}

/**
 * Fetch all grocery items for a specific list
 */
export function useGroceryItems(groceryListId: bigint) {
  return useQuery(listGroceryItems, { groceryList: groceryListId });
}

// ============ Mutation Hooks ============

/**
 * Create a new grocery list
 */
export function useCreateGroceryList() {
  return useMutation(createGroceryList);
}

/**
 * Update an existing grocery list
 */
export function useUpdateGroceryList() {
  return useMutation(updateGroceryList);
}

/**
 * Delete a grocery list
 */
export function useDeleteGroceryList() {
  return useMutation(deleteGroceryList);
}

/**
 * Create a new grocery item
 */
export function useCreateGroceryItem() {
  return useMutation(createGroceryItem);
}

/**
 * Update an existing grocery item
 */
export function useUpdateGroceryItem() {
  return useMutation(updateGroceryItem);
}

/**
 * Delete a grocery item
 */
export function useDeleteGroceryItem() {
  return useMutation(deleteGroceryItem);
}

