-- name: GetGroceryList :one
SELECT *
FROM grocery_list
WHERE id = ? AND tenant_id = ?
LIMIT 1;

-- name: CreateGroceryList :one
INSERT INTO grocery_list (name, tenant_id)
VALUES (?, ?)
RETURNING *;

-- name: UpdateGroceryList :one
UPDATE grocery_list
SET name = ?, version = version + 1
WHERE id = ? AND tenant_id = ? AND version = ?
RETURNING *;

-- name: DeleteGroceryList :one
DELETE FROM grocery_list
WHERE id = ? AND tenant_id = ?
RETURNING *;

-- name: ListGroceryLists :many
SELECT *
FROM grocery_list
WHERE tenant_id = ?;

-- name: GetGroceryItem :one
SELECT gi.*
FROM grocery_item gi
JOIN grocery_list gl ON gi.grocery_list_id = gl.id
WHERE gi.id = ? AND gl.tenant_id = ?
LIMIT 1;

-- name: CreateGroceryItem :one
INSERT INTO grocery_item (grocery_list_id, name, quantity, checked)
VALUES (?, ?, ?, ?)
RETURNING *;

-- name: UpdateGroceryItem :one
UPDATE grocery_item
SET name = ?, quantity = ?, checked = ?, version = version + 1
WHERE grocery_item.id = ? AND grocery_item.version = ?
RETURNING *;

-- name: DeleteGroceryItem :one
DELETE FROM grocery_item
WHERE grocery_item.id = ?
RETURNING *;

-- name: ListGroceryItemsByGroceryList :many
SELECT gi.*
FROM grocery_item gi
JOIN grocery_list gl ON gi.grocery_list_id = gl.id
WHERE gi.grocery_list_id = ? AND gl.tenant_id = ?;
