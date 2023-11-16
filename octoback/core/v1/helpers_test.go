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
	items := make([]*models.GroceryItem, rand.Intn(5)+1)
	for i := range items {
		items[i] = FakeGroceryItem()
	}

	return &models.GroceryList{
		Id:    FakeGroceryListID(),
		Name:  fake.Characters(),
		Items: items,
	}
}

func FakeGroceryItemID() string {
	return "gi_" + ksuid.New().String()[:8]
}

func FakeGroceryItem() *models.GroceryItem {
	return &models.GroceryItem{
		Id:       FakeGroceryItemID(),
		Name:     fake.Characters(),
		Quantity: int32(rand.Intn(5) + 1),
	}
}
