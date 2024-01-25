package v1

import (
	"math/rand"

	"github.com/icrowley/fake"
	models "github.com/kiliandbigblue/octoback/gen/proto/go/octoback/core/v1"
	"github.com/segmentio/ksuid"
)

func FakeGroceryListID() string {
	return "gl_" + ksuid.New().String()[:8]
}

func FakeGroceryList() *models.GroceryList {
	items := make([]*models.GroceryItem, 0, rand.Intn(5)+1)
	for i := range items {
		item := FakeGroceryItem()
		item.Id = int64(i)
		items = append(items, item)
	}

	return &models.GroceryList{
		Id:    FakeGroceryListID(),
		Name:  fake.Characters(),
		Items: items,
	}
}

func FakeGroceryItem() *models.GroceryItem {
	return &models.GroceryItem{
		Name:     fake.Characters(),
		Quantity: int32(rand.Intn(5) + 1),
	}
}
