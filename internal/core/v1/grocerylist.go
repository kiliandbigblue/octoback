package v1

import (
	"encoding/json"
	"io"

	models "github.com/kiliandbigblue/octoback/gen/proto/go/octoback/core/v1"
	"github.com/pkg/errors"
)

func NewGroceryLists(r io.Reader) ([]*models.GroceryList, error) {
	var groceryLists []*models.GroceryList
	err := json.NewDecoder(r).Decode(&groceryLists)
	if err != nil {
		err = errors.Wrapf(err, "failed to decode grocery lists")
	}

	return groceryLists, err
}
