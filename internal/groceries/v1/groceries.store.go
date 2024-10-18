package v1

import (
	"context"

	models "github.com/kiliandbigblue/octoback/gen/proto/go/octoback/groceries/v1"
)

//go:generate mockery --name Store
type Store interface {
	GroceryLists(ctx context.Context) ([]*models.GroceryList, error)
	GroceryList(ctx context.Context, id int64) (*models.GroceryList, error)
	CreateGroceryList(ctx context.Context, r *models.GroceryList) (*models.GroceryList, error)
	UpdateGroceryList(ctx context.Context, r *models.GroceryList) (*models.GroceryList, error)
	DeleteGroceryList(ctx context.Context, id int64) error
}
