-- name: GetGroceryList :one
SELECT * FROM grocery_list WHERE id = $1 LIMIT 1;

-- name: CreateGroceryList :one
INSERT INTO grocery_list (name) VALUES ($1) RETURNING *;

-- name: UpdateGroceryList :one
UPDATE grocery_list 
SET name = $2, version = version + 1
WHERE id = $1 AND version = $3
RETURNING *;

-- name: DeleteGroceryList :one
DELETE FROM grocery_list WHERE id = $1 RETURNING *;

-- name: ListGroceryLists :many
SELECT * FROM grocery_list;

-- name: GetGroceryItem :one
SELECT * FROM grocery_item WHERE id = $1 LIMIT 1;

-- name: CreateGroceryItem :one
INSERT INTO grocery_item (grocery_list_id, name, quantity, checked) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: UpdateGroceryItem :one
UPDATE grocery_item
SET name = $2, quantity = $3, checked = $4, version = version + 1
WHERE id = $1 AND version = $5
RETURNING *;

-- name: DeleteGroceryItem :one
DELETE FROM grocery_item WHERE id = $1 RETURNING *;

-- name: ListGroceryItemsByGroceryList :many
SELECT * FROM grocery_item WHERE grocery_list_id = $1;


