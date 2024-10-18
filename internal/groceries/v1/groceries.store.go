package v1

import (
	"context"

	models "github.com/kiliandbigblue/octoback/gen/proto/go/octoback/groceries/v1"
)

//go:generate mockery --name Store
type Store interface {
	CreateGroceryItem(ctx context.Context, groceryList int64, r *models.GroceryItem) (*models.GroceryItem, error)
	CreateGroceryList(ctx context.Context, r *models.GroceryList) (*models.GroceryList, error)
	DeleteGroceryList(ctx context.Context, id int64) error
	GroceryList(ctx context.Context, id int64) (*models.GroceryList, error)
	GroceryLists(ctx context.Context) ([]*models.GroceryList, error)
	UpdateGroceryList(ctx context.Context, r *models.GroceryList) (*models.GroceryList, error)
}
