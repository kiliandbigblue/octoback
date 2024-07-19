package v1

import (
	models "github.com/kiliandbigblue/octoback/gen/proto/go/octoback/groceries/v1"
)

//go:generate mockery --name Store
type Store interface {
	GroceryLists() ([]*models.GroceryList, error)
	GroceryList(id string) (*models.GroceryList, error)
	SetGroceryList(gl *models.GroceryList) error
	DeleteGroceryList(id string) error
}
