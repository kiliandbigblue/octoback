package testhelper

import (
	"crypto/rand"
	"math/big"
	"time"

	"github.com/icrowley/fake"
	models "github.com/kiliandbigblue/octoback/gen/proto/go/octoback/groceries/v1"
	store "github.com/kiliandbigblue/octoback/gen/store/v1"
	"github.com/segmentio/ksuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func FakeGroceryListID() int64 {
	return ksuid.New().Time().Unix()
}

func FakeGroceryList() *models.GroceryList {
	return &models.GroceryList{
		Id:         FakeGroceryListID(),
		Name:       fake.Characters(),
		CreateTime: timestamppb.Now(),
		Version:    1,
	}
}

func FakeStoreGroceryList() store.GroceryList {
	return store.GroceryList{
		ID:        FakeGroceryListID(),
		Name:      fake.Characters(),
		CreatedAt: time.Now(),
		Version:   1,
	}
}

func FakeGroceryItemID() int64 {
	return ksuid.New().Time().Unix()
}

func FakeGroceryItem(groceryList int64) *models.GroceryItem {
	qty, _ := rand.Int(rand.Reader, big.NewInt(5))
	checked, _ := rand.Int(rand.Reader, big.NewInt(2))
	return &models.GroceryItem{
		Id:          FakeGroceryItemID(),
		GroceryList: groceryList,
		Name:        fake.Characters(),
		Quantity:    int32(qty.Int64() + 1),
		Checked:     bool(checked.Int64()%2 == 0),
		CreateTime:  timestamppb.Now(),
		Version:     1,
	}
}

func FakeStoreGroceryItem(groceryList int64) store.GroceryItem {
	qty, _ := rand.Int(rand.Reader, big.NewInt(5))
	checked, _ := rand.Int(rand.Reader, big.NewInt(2))
	return store.GroceryItem{
		ID:            FakeGroceryItemID(),
		GroceryListID: groceryList,
		Name:          fake.Characters(),
		Quantity:      int32(qty.Int64() + 1),
		Checked:       bool(checked.Int64()%2 == 0),
		CreatedAt:     time.Now(),
		Version:       1,
	}
}
