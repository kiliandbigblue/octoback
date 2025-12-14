package v1

import (
	"time"

	models "github.com/kiliandbigblue/octoback/gen/proto/go/octoback/groceries/v1"
	store "github.com/kiliandbigblue/octoback/gen/store/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// parseTime parses a SQLite datetime string to time.Time.
func parseTime(s string) time.Time {
	t, _ := time.Parse("2006-01-02 15:04:05", s)
	return t
}

// translateGroceryList translates a store.GroceryList to a models.GroceryList.
func translateGroceryList(groceryList store.GroceryList) *models.GroceryList {
	return &models.GroceryList{
		Id:         groceryList.ID,
		Name:       groceryList.Name,
		CreateTime: timestamppb.New(parseTime(groceryList.CreatedAt)),
		Version:    int32(groceryList.Version),
	}
}

// translateGroceryItem translates a store.GroceryItem to a models.GroceryItem.
func translateGroceryItem(groceryItem store.GroceryItem) *models.GroceryItem {
	return &models.GroceryItem{
		Id:          groceryItem.ID,
		GroceryList: groceryItem.GroceryListID,
		Name:        groceryItem.Name,
		Quantity:    int32(groceryItem.Quantity),
		Checked:     groceryItem.Checked != 0,
		CreateTime:  timestamppb.New(parseTime(groceryItem.CreatedAt)),
		Version:     int32(groceryItem.Version),
	}
}

// boolToInt64 converts a bool to int64 for SQLite.
func boolToInt64(b bool) int64 {
	if b {
		return 1
	}
	return 0
}
