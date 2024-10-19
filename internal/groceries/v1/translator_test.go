package v1

import (
	"testing"
	"time"

	store "github.com/kiliandbigblue/octoback/internal/store/v1"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestTranslator(t *testing.T) {
	t.Run("can translate a grocery list", func(t *testing.T) {
		// Given
		li := store.GroceryList{
			ID:        1,
			Name:      "Grocery List",
			CreatedAt: time.Now(),
			Version:   1,
		}

		// When
		translated := translateGroceryList(li)

		// Then
		assert.Equal(t, li.ID, translated.Id)
		assert.Equal(t, li.Name, translated.Name)
		assert.Equal(t, timestamppb.New(li.CreatedAt), translated.CreateTime)
		assert.Equal(t, li.Version, translated.Version)
	})

	t.Run("can translate a grocery item", func(t *testing.T) {
		// Given
		it := store.GroceryItem{
			ID:        1,
			Name:      "Grocery Item",
			CreatedAt: time.Now(),
			Version:   1,
		}

		// When
		translated := translateGroceryItem(it)

		// Then
		assert.Equal(t, it.ID, translated.Id)
		assert.Equal(t, it.Name, translated.Name)
		assert.Equal(t, timestamppb.New(it.CreatedAt), translated.CreateTime)
		assert.Equal(t, it.Version, translated.Version)
	})
}
