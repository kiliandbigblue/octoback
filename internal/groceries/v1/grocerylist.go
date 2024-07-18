package v1

import (
	"encoding/json"
	"io"

	models "github.com/kiliandbigblue/octoback/gen/proto/go/octoback/groceries/v1"
	"github.com/pkg/errors"
)

type GroceryLists []*models.GroceryList

func NewGroceryLists(r io.Reader) (GroceryLists, error) {
	var groceryLists GroceryLists
	err := json.NewDecoder(r).Decode(&groceryLists)
	if err != nil {
		err = errors.Wrapf(err, "failed to decode grocery lists")
	}

	return groceryLists, err
}

// Find returns the grocery list with the given id, or nil if not found.
func (g GroceryLists) Find(id string) *models.GroceryList {
	for _, gl := range g {
		if gl.Id == id {
			return gl
		}
	}

	return nil
}
