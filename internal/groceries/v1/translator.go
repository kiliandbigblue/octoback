package v1

import (
	models "github.com/kiliandbigblue/octoback/gen/proto/go/octoback/groceries/v1"
	store "github.com/kiliandbigblue/octoback/internal/store/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// translateGroceryList translates a store.GroceryList to a models.GroceryList.
func translateGroceryList(groceryList store.GroceryList) *models.GroceryList {
	return &models.GroceryList{
		Id:         groceryList.ID,
		Name:       groceryList.Name,
		CreateTime: timestamppb.New(groceryList.CreatedAt),
		Version:    groceryList.Version,
	}
}

// translateGroceryItem translates a store.GroceryItem to a models.GroceryItem.
func translateGroceryItem(groceryItem store.GroceryItem) *models.GroceryItem {
	return &models.GroceryItem{
		Id:          groceryItem.ID,
		GroceryList: groceryItem.GroceryListID,
		Name:        groceryItem.Name,
		Quantity:    groceryItem.Quantity,
		Checked:     groceryItem.Checked,
		CreateTime:  timestamppb.New(groceryItem.CreatedAt),
		Version:     groceryItem.Version,
	}
}
