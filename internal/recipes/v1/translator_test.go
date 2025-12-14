package v1

import (
	"testing"
	"time"

	store "github.com/kiliandbigblue/octoback/gen/store/v1"
	"github.com/stretchr/testify/assert"
)

func TestTranslator(t *testing.T) {
	t.Run("can translate a recipe", func(t *testing.T) {
		// Given
		createdAt := time.Now().Format("2006-01-02 15:04:05")
		recipe := store.Recipe{
			ID:              1,
			Name:            "Pasta Carbonara",
			Description:     "Classic Italian pasta dish",
			Servings:        4,
			PrepTimeMinutes: 15,
			CookTimeMinutes: 20,
			CreatedAt:       createdAt,
			Version:         1,
		}

		// When
		translated := translateRecipe(recipe)

		// Then
		assert.Equal(t, recipe.ID, translated.Id)
		assert.Equal(t, recipe.Name, translated.Name)
		assert.Equal(t, recipe.Description, translated.Description)
		assert.Equal(t, int32(recipe.Servings), translated.Servings)
		assert.Equal(t, int32(recipe.PrepTimeMinutes), translated.PrepTimeMinutes)
		assert.Equal(t, int32(recipe.CookTimeMinutes), translated.CookTimeMinutes)
		assert.Equal(t, parseTime(recipe.CreatedAt), translated.CreateTime.AsTime())
		assert.Equal(t, int32(recipe.Version), translated.Version)
	})

	t.Run("can translate a recipe ingredient", func(t *testing.T) {
		// Given
		createdAt := time.Now().Format("2006-01-02 15:04:05")
		ingredient := store.RecipeIngredient{
			ID:        1,
			RecipeID:  2,
			Name:      "Spaghetti",
			Quantity:  "400",
			Unit:      "g",
			CreatedAt: createdAt,
			Version:   1,
		}

		// When
		translated := translateRecipeIngredient(ingredient)

		// Then
		assert.Equal(t, ingredient.ID, translated.Id)
		assert.Equal(t, ingredient.RecipeID, translated.Recipe)
		assert.Equal(t, ingredient.Name, translated.Name)
		assert.Equal(t, ingredient.Quantity, translated.Quantity)
		assert.Equal(t, ingredient.Unit, translated.Unit)
		assert.Equal(t, parseTime(ingredient.CreatedAt), translated.CreateTime.AsTime())
		assert.Equal(t, int32(ingredient.Version), translated.Version)
	})
}

