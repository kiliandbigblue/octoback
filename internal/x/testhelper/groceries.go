package testhelper

import (
	"crypto/rand"
	"math/big"

	"github.com/icrowley/fake"
	models "github.com/kiliandbigblue/octoback/gen/proto/go/octoback/groceries/v1"
	"github.com/segmentio/ksuid"
)

func FakeGroceryListID() string {
	return "gl_" + ksuid.New().String()[:8]
}

func FakeGroceryList() *models.GroceryList {
	nItems, _ := rand.Int(rand.Reader, big.NewInt(5))
	items := make([]*models.GroceryItem, 0, nItems.Int64()+1)

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
	qty, _ := rand.Int(rand.Reader, big.NewInt(5))
	return &models.GroceryItem{
		Name:     fake.Characters(),
		Quantity: int32(qty.Int64() + 1),
	}
}
