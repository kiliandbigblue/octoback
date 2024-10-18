-- name: GetGroceryList :one
SELECT * FROM grocery_list WHERE id = $1 LIMIT 1;

-- name: CreateGroceryList :one
INSERT INTO grocery_list (name) VALUES ($1) RETURNING *;

-- name: UpdateGroceryList :one
UPDATE grocery_list 
SET name = $2, version = version + 1
WHERE id = $1
RETURNING *;

-- name: DeleteGroceryList :one
DELETE FROM grocery_list WHERE id = $1 RETURNING *;

-- name: ListGroceryLists :many
SELECT * FROM grocery_list;

