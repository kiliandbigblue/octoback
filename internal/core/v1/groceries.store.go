package v1

import (
	"errors"

	models "github.com/kiliandbigblue/octoback/gen/proto/go/octoback/core/v1"
)

var ErrNoSuchEntity = errors.New("no such entity")

type StoreValidationError struct {
	Err error
}

func (e *StoreValidationError) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	}
	return "store validation error"
}

//go:generate mockery --name Store
type Store interface {
	GroceryLists() ([]*models.GroceryList, error)
	GroceryList(id string) (*models.GroceryList, error)
	SetGroceryList(gl *models.GroceryList) error
	DeleteGroceryList(id string) error
}
