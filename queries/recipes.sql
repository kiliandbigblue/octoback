-- name: GetRecipe :one
SELECT *
FROM recipe
WHERE id = ?
LIMIT 1;

-- name: CreateRecipe :one
INSERT INTO recipe (name, description, servings, prep_time_minutes, cook_time_minutes)
VALUES (?, ?, ?, ?, ?)
RETURNING *;

-- name: UpdateRecipe :one
UPDATE recipe
SET name = ?, description = ?, servings = ?, prep_time_minutes = ?, cook_time_minutes = ?, version = version + 1
WHERE id = ? AND version = ?
RETURNING *;

-- name: DeleteRecipe :one
DELETE FROM recipe
WHERE id = ?
RETURNING *;

-- name: ListRecipes :many
SELECT *
FROM recipe;

-- name: GetRecipeIngredient :one
SELECT *
FROM recipe_ingredient
WHERE id = ?
LIMIT 1;

-- name: CreateRecipeIngredient :one
INSERT INTO recipe_ingredient (recipe_id, name, quantity, unit)
VALUES (?, ?, ?, ?)
RETURNING *;

-- name: UpdateRecipeIngredient :one
UPDATE recipe_ingredient
SET name = ?, quantity = ?, unit = ?, version = version + 1
WHERE id = ? AND version = ?
RETURNING *;

-- name: DeleteRecipeIngredient :one
DELETE FROM recipe_ingredient
WHERE id = ?
RETURNING *;

-- name: ListRecipeIngredientsByRecipe :many
SELECT *
FROM recipe_ingredient
WHERE recipe_id = ?;

