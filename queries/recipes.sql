-- name: GetRecipe :one
SELECT *
FROM recipe
WHERE id = ? AND tenant_id = ?
LIMIT 1;

-- name: CreateRecipe :one
INSERT INTO recipe (name, description, servings, prep_time_minutes, cook_time_minutes, tenant_id)
VALUES (?, ?, ?, ?, ?, ?)
RETURNING *;

-- name: UpdateRecipe :one
UPDATE recipe
SET name = ?, description = ?, servings = ?, prep_time_minutes = ?, cook_time_minutes = ?, version = version + 1
WHERE id = ? AND tenant_id = ? AND version = ?
RETURNING *;

-- name: DeleteRecipe :one
DELETE FROM recipe
WHERE id = ? AND tenant_id = ?
RETURNING *;

-- name: ListRecipes :many
SELECT *
FROM recipe
WHERE tenant_id = ?;

-- name: GetRecipeIngredient :one
SELECT ri.*
FROM recipe_ingredient ri
JOIN recipe r ON ri.recipe_id = r.id
WHERE ri.id = ? AND r.tenant_id = ?
LIMIT 1;

-- name: CreateRecipeIngredient :one
INSERT INTO recipe_ingredient (recipe_id, name, quantity, unit)
VALUES (?, ?, ?, ?)
RETURNING *;

-- name: UpdateRecipeIngredient :one
UPDATE recipe_ingredient
SET name = ?, quantity = ?, unit = ?, version = version + 1
WHERE recipe_ingredient.id = ? AND recipe_ingredient.version = ?
RETURNING *;

-- name: DeleteRecipeIngredient :one
DELETE FROM recipe_ingredient
WHERE recipe_ingredient.id = ?
RETURNING *;

-- name: ListRecipeIngredientsByRecipe :many
SELECT ri.*
FROM recipe_ingredient ri
JOIN recipe r ON ri.recipe_id = r.id
WHERE ri.recipe_id = ? AND r.tenant_id = ?;
