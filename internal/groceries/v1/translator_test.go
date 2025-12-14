package v1

import (
	"testing"
	"time"

	store "github.com/kiliandbigblue/octoback/gen/store/v1"
	"github.com/stretchr/testify/assert"
)

func TestTranslator(t *testing.T) {
	t.Run("can translate a grocery list", func(t *testing.T) {
		// Given
		createdAt := time.Now().Format("2006-01-02 15:04:05")
		li := store.GroceryList{
			ID:        1,
			Name:      "Grocery List",
			CreatedAt: createdAt,
			Version:   1,
		}

		// When
		translated := translateGroceryList(li)

		// Then
		assert.Equal(t, li.ID, translated.Id)
		assert.Equal(t, li.Name, translated.Name)
		assert.Equal(t, parseTime(li.CreatedAt), translated.CreateTime.AsTime())
		assert.Equal(t, int32(li.Version), translated.Version)
	})

	t.Run("can translate a grocery item", func(t *testing.T) {
		// Given
		createdAt := time.Now().Format("2006-01-02 15:04:05")
		it := store.GroceryItem{
			ID:            1,
			GroceryListID: 2,
			Name:          "Grocery Item",
			Quantity:      3,
			Checked:       1,
			CreatedAt:     createdAt,
			Version:       1,
		}

		// When
		translated := translateGroceryItem(it)

		// Then
		assert.Equal(t, it.ID, translated.Id)
		assert.Equal(t, it.GroceryListID, translated.GroceryList)
		assert.Equal(t, it.Name, translated.Name)
		assert.Equal(t, int32(it.Quantity), translated.Quantity)
		assert.Equal(t, it.Checked != 0, translated.Checked)
		assert.Equal(t, parseTime(it.CreatedAt), translated.CreateTime.AsTime())
		assert.Equal(t, int32(it.Version), translated.Version)
	})
}
