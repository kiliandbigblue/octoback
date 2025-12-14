-- name: GetGroceryList :one
SELECT *
FROM grocery_list
WHERE id = ?
LIMIT 1;

-- name: CreateGroceryList :one
INSERT INTO grocery_list (name)
VALUES (?)
RETURNING *;

-- name: UpdateGroceryList :one
UPDATE grocery_list
SET name = ?, version = version + 1
WHERE id = ? AND version = ?
RETURNING *;

-- name: DeleteGroceryList :one
DELETE FROM grocery_list
WHERE id = ?
RETURNING *;

-- name: ListGroceryLists :many
SELECT *
FROM grocery_list;

-- name: GetGroceryItem :one
SELECT *
FROM grocery_item
WHERE id = ?
LIMIT 1;

-- name: CreateGroceryItem :one
INSERT INTO grocery_item (grocery_list_id, name, quantity, checked)
VALUES (?, ?, ?, ?)
RETURNING *;

-- name: UpdateGroceryItem :one
UPDATE grocery_item
SET name = ?, quantity = ?, checked = ?, version = version + 1
WHERE id = ? AND version = ?
RETURNING *;

-- name: DeleteGroceryItem :one
DELETE FROM grocery_item
WHERE id = ?
RETURNING *;

-- name: ListGroceryItemsByGroceryList :many
SELECT *
FROM grocery_item
WHERE grocery_list_id = ?;
