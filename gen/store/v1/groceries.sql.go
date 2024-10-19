// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: groceries.sql

package v1

import (
	"context"
)

const createGroceryItem = `-- name: CreateGroceryItem :one
insert into GROCERY_ITEM (GROCERY_LIST_ID, NAME, QUANTITY, CHECKED)
values ($1, $2, $3, $4)
returning id, grocery_list_id, name, quantity, checked, created_at, version
`

type CreateGroceryItemParams struct {
	GroceryListID int64
	Name          string
	Quantity      int32
	Checked       bool
}

func (q *Queries) CreateGroceryItem(ctx context.Context, arg CreateGroceryItemParams) (GroceryItem, error) {
	row := q.db.QueryRowContext(ctx, createGroceryItem,
		arg.GroceryListID,
		arg.Name,
		arg.Quantity,
		arg.Checked,
	)
	var i GroceryItem
	err := row.Scan(
		&i.ID,
		&i.GroceryListID,
		&i.Name,
		&i.Quantity,
		&i.Checked,
		&i.CreatedAt,
		&i.Version,
	)
	return i, err
}

const createGroceryList = `-- name: CreateGroceryList :one
insert into GROCERY_LIST (NAME)
values ($1)
returning id, name, created_at, version
`

func (q *Queries) CreateGroceryList(ctx context.Context, name string) (GroceryList, error) {
	row := q.db.QueryRowContext(ctx, createGroceryList, name)
	var i GroceryList
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.Version,
	)
	return i, err
}

const deleteGroceryItem = `-- name: DeleteGroceryItem :one
delete
from GROCERY_ITEM
where ID = $1
returning id, grocery_list_id, name, quantity, checked, created_at, version
`

func (q *Queries) DeleteGroceryItem(ctx context.Context, id int64) (GroceryItem, error) {
	row := q.db.QueryRowContext(ctx, deleteGroceryItem, id)
	var i GroceryItem
	err := row.Scan(
		&i.ID,
		&i.GroceryListID,
		&i.Name,
		&i.Quantity,
		&i.Checked,
		&i.CreatedAt,
		&i.Version,
	)
	return i, err
}

const deleteGroceryList = `-- name: DeleteGroceryList :one
delete
from GROCERY_LIST
where ID = $1
returning id, name, created_at, version
`

func (q *Queries) DeleteGroceryList(ctx context.Context, id int64) (GroceryList, error) {
	row := q.db.QueryRowContext(ctx, deleteGroceryList, id)
	var i GroceryList
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.Version,
	)
	return i, err
}

const getGroceryItem = `-- name: GetGroceryItem :one
select id, grocery_list_id, name, quantity, checked, created_at, version
from GROCERY_ITEM
where ID = $1
limit 1
`

func (q *Queries) GetGroceryItem(ctx context.Context, id int64) (GroceryItem, error) {
	row := q.db.QueryRowContext(ctx, getGroceryItem, id)
	var i GroceryItem
	err := row.Scan(
		&i.ID,
		&i.GroceryListID,
		&i.Name,
		&i.Quantity,
		&i.Checked,
		&i.CreatedAt,
		&i.Version,
	)
	return i, err
}

const getGroceryList = `-- name: GetGroceryList :one
select id, name, created_at, version
from GROCERY_LIST
where ID = $1
limit 1
`

func (q *Queries) GetGroceryList(ctx context.Context, id int64) (GroceryList, error) {
	row := q.db.QueryRowContext(ctx, getGroceryList, id)
	var i GroceryList
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.Version,
	)
	return i, err
}

const listGroceryItemsByGroceryList = `-- name: ListGroceryItemsByGroceryList :many
select id, grocery_list_id, name, quantity, checked, created_at, version
from GROCERY_ITEM
where GROCERY_LIST_ID = $1
`

func (q *Queries) ListGroceryItemsByGroceryList(ctx context.Context, groceryListID int64) ([]GroceryItem, error) {
	rows, err := q.db.QueryContext(ctx, listGroceryItemsByGroceryList, groceryListID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GroceryItem
	for rows.Next() {
		var i GroceryItem
		if err := rows.Scan(
			&i.ID,
			&i.GroceryListID,
			&i.Name,
			&i.Quantity,
			&i.Checked,
			&i.CreatedAt,
			&i.Version,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listGroceryLists = `-- name: ListGroceryLists :many
select id, name, created_at, version
from GROCERY_LIST
`

func (q *Queries) ListGroceryLists(ctx context.Context) ([]GroceryList, error) {
	rows, err := q.db.QueryContext(ctx, listGroceryLists)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GroceryList
	for rows.Next() {
		var i GroceryList
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.CreatedAt,
			&i.Version,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateGroceryItem = `-- name: UpdateGroceryItem :one
update GROCERY_ITEM
set NAME = $2, QUANTITY = $3, CHECKED = $4, VERSION = VERSION + 1
where ID = $1 and VERSION = $5
returning id, grocery_list_id, name, quantity, checked, created_at, version
`

type UpdateGroceryItemParams struct {
	ID       int64
	Name     string
	Quantity int32
	Checked  bool
	Version  int32
}

func (q *Queries) UpdateGroceryItem(ctx context.Context, arg UpdateGroceryItemParams) (GroceryItem, error) {
	row := q.db.QueryRowContext(ctx, updateGroceryItem,
		arg.ID,
		arg.Name,
		arg.Quantity,
		arg.Checked,
		arg.Version,
	)
	var i GroceryItem
	err := row.Scan(
		&i.ID,
		&i.GroceryListID,
		&i.Name,
		&i.Quantity,
		&i.Checked,
		&i.CreatedAt,
		&i.Version,
	)
	return i, err
}

const updateGroceryList = `-- name: UpdateGroceryList :one
update GROCERY_LIST
set NAME = $2, VERSION = VERSION + 1
where ID = $1 and VERSION = $3
returning id, name, created_at, version
`

type UpdateGroceryListParams struct {
	ID      int64
	Name    string
	Version int32
}

func (q *Queries) UpdateGroceryList(ctx context.Context, arg UpdateGroceryListParams) (GroceryList, error) {
	row := q.db.QueryRowContext(ctx, updateGroceryList, arg.ID, arg.Name, arg.Version)
	var i GroceryList
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.Version,
	)
	return i, err
}
