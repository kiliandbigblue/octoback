-- name: GetGroceryList :one
select *
from GROCERY_LIST
where ID = $1
limit 1;

-- name: CreateGroceryList :one
insert into GROCERY_LIST (NAME)
values ($1)
returning *;

-- name: UpdateGroceryList :one
update GROCERY_LIST
set NAME = $2, VERSION = VERSION + 1
where ID = $1 and VERSION = $3
returning *;

-- name: DeleteGroceryList :one
delete
from GROCERY_LIST
where ID = $1
returning *;

-- name: ListGroceryLists :many
select *
from GROCERY_LIST;

-- name: GetGroceryItem :one
select *
from GROCERY_ITEM
where ID = $1
limit 1;

-- name: CreateGroceryItem :one
insert into GROCERY_ITEM (GROCERY_LIST_ID, NAME, QUANTITY, CHECKED)
values ($1, $2, $3, $4)
returning *;

-- name: UpdateGroceryItem :one
update GROCERY_ITEM
set NAME = $2, QUANTITY = $3, CHECKED = $4, VERSION = VERSION + 1
where ID = $1 and VERSION = $5
returning *;

-- name: DeleteGroceryItem :one
delete
from GROCERY_ITEM
where ID = $1
returning *;

-- name: ListGroceryItemsByGroceryList :many
select *
from GROCERY_ITEM
where GROCERY_LIST_ID = $1;


