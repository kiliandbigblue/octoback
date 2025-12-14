package v1

import (
	"time"

	models "github.com/kiliandbigblue/octoback/gen/proto/go/octoback/recipes/v1"
	store "github.com/kiliandbigblue/octoback/gen/store/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// parseTime parses a SQLite datetime string to time.Time.
func parseTime(s string) time.Time {
	t, _ := time.Parse("2006-01-02 15:04:05", s)
	return t
}

// translateRecipe translates a store.Recipe to a models.Recipe.
func translateRecipe(recipe store.Recipe) *models.Recipe {
	return &models.Recipe{
		Id:              recipe.ID,
		Name:            recipe.Name,
		Description:     recipe.Description,
		Servings:        int32(recipe.Servings),
		PrepTimeMinutes: int32(recipe.PrepTimeMinutes),
		CookTimeMinutes: int32(recipe.CookTimeMinutes),
		CreateTime:      timestamppb.New(parseTime(recipe.CreatedAt)),
		Version:         int32(recipe.Version),
	}
}

// translateRecipeIngredient translates a store.RecipeIngredient to a models.RecipeIngredient.
func translateRecipeIngredient(ingredient store.RecipeIngredient) *models.RecipeIngredient {
	return &models.RecipeIngredient{
		Id:         ingredient.ID,
		Recipe:     ingredient.RecipeID,
		Name:       ingredient.Name,
		Quantity:   ingredient.Quantity,
		Unit:       ingredient.Unit,
		CreateTime: timestamppb.New(parseTime(ingredient.CreatedAt)),
		Version:    int32(ingredient.Version),
	}
}

