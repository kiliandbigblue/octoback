package testhelper

import (
	"crypto/rand"
	"math/big"
	"time"

	"github.com/icrowley/fake"
	models "github.com/kiliandbigblue/octoback/gen/proto/go/octoback/recipes/v1"
	store "github.com/kiliandbigblue/octoback/gen/store/v1"
	"github.com/segmentio/ksuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func FakeRecipeID() int64 {
	return ksuid.New().Time().Unix()
}

func FakeRecipe() *models.Recipe {
	servings, _ := rand.Int(rand.Reader, big.NewInt(8))
	prepTime, _ := rand.Int(rand.Reader, big.NewInt(60))
	cookTime, _ := rand.Int(rand.Reader, big.NewInt(120))
	return &models.Recipe{
		Id:              FakeRecipeID(),
		Name:            fake.Characters(),
		Description:     fake.Sentences(),
		Servings:        int32(servings.Int64() + 1),
		PrepTimeMinutes: int32(prepTime.Int64()),
		CookTimeMinutes: int32(cookTime.Int64()),
		CreateTime:      timestamppb.Now(),
		Version:         1,
	}
}

func FakeStoreRecipe() store.Recipe {
	servings, _ := rand.Int(rand.Reader, big.NewInt(8))
	prepTime, _ := rand.Int(rand.Reader, big.NewInt(60))
	cookTime, _ := rand.Int(rand.Reader, big.NewInt(120))
	return store.Recipe{
		ID:              FakeRecipeID(),
		Name:            fake.Characters(),
		Description:     fake.Sentences(),
		Servings:        servings.Int64() + 1,
		PrepTimeMinutes: prepTime.Int64(),
		CookTimeMinutes: cookTime.Int64(),
		CreatedAt:       time.Now().Format("2006-01-02 15:04:05"),
		Version:         1,
	}
}

func FakeRecipeIngredientID() int64 {
	return ksuid.New().Time().Unix()
}

func FakeRecipeIngredient(recipe int64) *models.RecipeIngredient {
	return &models.RecipeIngredient{
		Id:         FakeRecipeIngredientID(),
		Recipe:     recipe,
		Name:       fake.Characters(),
		Quantity:   fake.DigitsN(2),
		Unit:       fake.Word(),
		CreateTime: timestamppb.Now(),
		Version:    1,
	}
}

func FakeStoreRecipeIngredient(recipe int64) store.RecipeIngredient {
	return store.RecipeIngredient{
		ID:        FakeRecipeIngredientID(),
		RecipeID:  recipe,
		Name:      fake.Characters(),
		Quantity:  fake.DigitsN(2),
		Unit:      fake.Word(),
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		Version:   1,
	}
}

