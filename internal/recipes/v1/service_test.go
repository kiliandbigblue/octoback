package v1

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"connectrpc.com/connect"
	"github.com/icrowley/fake"
	models "github.com/kiliandbigblue/octoback/gen/proto/go/octoback/recipes/v1"
	store "github.com/kiliandbigblue/octoback/gen/store/v1"
	"github.com/kiliandbigblue/octoback/internal/recipes/v1/mocks"
	"github.com/kiliandbigblue/octoback/internal/x/cloudzap"
	"github.com/kiliandbigblue/octoback/internal/x/tenant"
	"github.com/kiliandbigblue/octoback/internal/x/testhelper"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type serviceTestSuite struct {
	suite.Suite
	log      *zap.Logger
	ctx      context.Context
	tenantID int64
	s        *Service
	querier  *mocks.Querier
}

func TestServiceTestSuite(t *testing.T) {
	suite.Run(t, new(serviceTestSuite))
}

func (s *serviceTestSuite) SetupTest() {
	s.log = zap.NewNop()
	s.tenantID = testhelper.FakeTenantID()
	s.ctx = cloudzap.WithLogger(context.Background(), s.log)
	s.ctx = tenant.WithTenantID(s.ctx, s.tenantID)
	s.querier = mocks.NewQuerier(s.T())
	s.s = NewService(s.querier)
}

// Helper to create a NullInt64 from the test suite's tenant ID
func (s *serviceTestSuite) tenantIDParam() sql.NullInt64 {
	return sql.NullInt64{Int64: s.tenantID, Valid: true}
}

// Test that we can retrieve a recipe by its ID.
func (s *serviceTestSuite) TestGetRecipe_OK() {
	recipe := testhelper.FakeStoreRecipeWithTenant(s.tenantID)

	s.querier.EXPECT().GetRecipe(s.ctx, store.GetRecipeParams{
		ID:       recipe.ID,
		TenantID: s.tenantIDParam(),
	}).Return(recipe, nil)

	response, err := s.s.GetRecipe(s.ctx, connect.NewRequest(&models.GetRecipeRequest{
		Id: recipe.ID,
	}))
	s.NoError(err)
	s.Equal(translateRecipe(recipe), response.Msg.GetRecipe())
}

// Test that we get a NotFound error when the recipe does not exist.
func (s *serviceTestSuite) TestGetRecipe_Err_NotFound() {
	s.querier.EXPECT().GetRecipe(s.ctx, mock.Anything).Return(store.Recipe{}, sql.ErrNoRows)

	response, err := s.s.GetRecipe(s.ctx, connect.NewRequest(&models.GetRecipeRequest{
		Id: testhelper.FakeRecipeID(),
	}))
	s.Error(err)
	s.Equal(connect.CodeNotFound, connect.CodeOf(err))
	s.Nil(response)
}

// Test that we can create a recipe.
func (s *serviceTestSuite) TestCreateRecipe_Ok() {
	name := "Pasta Carbonara"
	description := "Classic Italian pasta dish"
	servings := int32(4)
	prepTime := int32(15)
	cookTime := int32(20)

	s.querier.EXPECT().CreateRecipe(s.ctx, store.CreateRecipeParams{
		Name:            name,
		Description:     description,
		Servings:        int64(servings),
		PrepTimeMinutes: int64(prepTime),
		CookTimeMinutes: int64(cookTime),
		TenantID:        s.tenantIDParam(),
	}).Return(store.Recipe{
		ID:              testhelper.FakeRecipeID(),
		Name:            name,
		Description:     description,
		Servings:        int64(servings),
		PrepTimeMinutes: int64(prepTime),
		CookTimeMinutes: int64(cookTime),
		CreatedAt:       time.Now().Format("2006-01-02 15:04:05"),
		Version:         1,
		TenantID:        s.tenantIDParam(),
	}, nil)

	response, err := s.s.CreateRecipe(s.ctx, connect.NewRequest(&models.CreateRecipeRequest{
		Name:            name,
		Description:     description,
		Servings:        servings,
		PrepTimeMinutes: prepTime,
		CookTimeMinutes: cookTime,
	}))
	s.NoError(err)
	s.NotEmpty(response.Msg.GetRecipe().GetId())
	s.Equal(name, response.Msg.GetRecipe().GetName())
	s.Equal(description, response.Msg.GetRecipe().GetDescription())
	s.Equal(servings, response.Msg.GetRecipe().GetServings())
}

// Test that we can't create a recipe with invalid data.
func (s *serviceTestSuite) TestCreateRecipe_Err() {
	s.querier.EXPECT().CreateRecipe(s.ctx, mock.Anything).Return(store.Recipe{}, sql.ErrNoRows)

	response, err := s.s.CreateRecipe(s.ctx, connect.NewRequest(&models.CreateRecipeRequest{
		Name: "",
	}))
	s.Error(err)
	s.Nil(response)
}

// Test that we can update a recipe.
func (s *serviceTestSuite) TestUpdateRecipe_Ok() {
	actual := testhelper.FakeStoreRecipeWithTenant(s.tenantID)

	input := actual
	input.Name = fake.Characters()
	input.Description = fake.Sentences()
	input.Servings = actual.Servings + 2

	s.querier.EXPECT().GetRecipe(s.ctx, store.GetRecipeParams{
		ID:       actual.ID,
		TenantID: s.tenantIDParam(),
	}).Return(actual, nil)

	s.querier.EXPECT().UpdateRecipe(s.ctx, store.UpdateRecipeParams{
		ID:              input.ID,
		Name:            input.Name,
		Description:     input.Description,
		Servings:        input.Servings,
		PrepTimeMinutes: input.PrepTimeMinutes,
		CookTimeMinutes: input.CookTimeMinutes,
		TenantID:        s.tenantIDParam(),
		Version:         input.Version,
	}).Return(input, nil)

	response, err := s.s.UpdateRecipe(s.ctx, connect.NewRequest(&models.UpdateRecipeRequest{
		Recipe:     translateRecipe(input),
		UpdateMask: &fieldmaskpb.FieldMask{Paths: []string{"name", "description", "servings"}},
	}))
	s.NoError(err)
	s.EqualExportedValues(*translateRecipe(input), *response.Msg.GetRecipe())
}

// Test that we can't update a recipe with an invalid mask.
func (s *serviceTestSuite) TestUpdateRecipe_Err_InvalidMask() {
	response, err := s.s.UpdateRecipe(s.ctx, connect.NewRequest(&models.UpdateRecipeRequest{
		Recipe:     testhelper.FakeRecipe(),
		UpdateMask: &fieldmaskpb.FieldMask{Paths: []string{"id"}},
	}))
	s.Error(err)
	s.Nil(response)
	s.Equal(connect.CodeInvalidArgument, connect.CodeOf(err))
	s.Contains(err.Error(), "invalid mask")
}

// Test that we get a NotFound error when the recipe does not exist.
func (s *serviceTestSuite) TestUpdateRecipe_Err_NotFound() {
	recipe := testhelper.FakeRecipe()

	s.querier.EXPECT().GetRecipe(s.ctx, store.GetRecipeParams{
		ID:       recipe.GetId(),
		TenantID: s.tenantIDParam(),
	}).Return(store.Recipe{}, sql.ErrNoRows)

	response, err := s.s.UpdateRecipe(s.ctx, connect.NewRequest(&models.UpdateRecipeRequest{
		Recipe:     recipe,
		UpdateMask: &fieldmaskpb.FieldMask{Paths: []string{"name"}},
	}))
	s.Error(err)
	s.Nil(response)
	s.Equal(connect.CodeNotFound, connect.CodeOf(err))
}

// Test that we can list all recipes.
func (s *serviceTestSuite) TestListRecipes_Ok() {
	r1 := testhelper.FakeStoreRecipeWithTenant(s.tenantID)
	r2 := testhelper.FakeStoreRecipeWithTenant(s.tenantID)

	s.querier.EXPECT().ListRecipes(s.ctx, s.tenantIDParam()).Return([]store.Recipe{r1, r2}, nil)

	response, err := s.s.ListRecipes(s.ctx, connect.NewRequest(&models.ListRecipesRequest{}))
	s.NoError(err)
	s.Equal([]*models.Recipe{translateRecipe(r1), translateRecipe(r2)}, response.Msg.GetRecipes())
}

// Test that we can delete a recipe.
func (s *serviceTestSuite) TestDeleteRecipe_Ok() {
	id := testhelper.FakeRecipeID()

	s.querier.EXPECT().DeleteRecipe(s.ctx, store.DeleteRecipeParams{
		ID:       id,
		TenantID: s.tenantIDParam(),
	}).Return(store.Recipe{}, nil)

	_, err := s.s.DeleteRecipe(s.ctx, connect.NewRequest(&models.DeleteRecipeRequest{
		Id: id,
	}))
	s.NoError(err)
}

// Test that we get a NotFound error when the recipe does not exist.
func (s *serviceTestSuite) TestDeleteRecipe_Err_NotFound() {
	s.querier.EXPECT().DeleteRecipe(s.ctx, mock.Anything).Return(store.Recipe{}, sql.ErrNoRows)

	_, err := s.s.DeleteRecipe(s.ctx, connect.NewRequest(&models.DeleteRecipeRequest{
		Id: testhelper.FakeRecipeID(),
	}))
	s.Error(err)
	s.Equal(connect.CodeNotFound, connect.CodeOf(err))
}

// Test that we can retrieve a recipe ingredient by its ID.
func (s *serviceTestSuite) TestGetRecipeIngredient_OK() {
	ingredient := testhelper.FakeStoreRecipeIngredient(testhelper.FakeRecipeID())

	s.querier.EXPECT().GetRecipeIngredient(s.ctx, store.GetRecipeIngredientParams{
		ID:       ingredient.ID,
		TenantID: s.tenantIDParam(),
	}).Return(ingredient, nil)

	response, err := s.s.GetRecipeIngredient(s.ctx, connect.NewRequest(&models.GetRecipeIngredientRequest{
		Id: ingredient.ID,
	}))
	s.NoError(err)
	s.Equal(translateRecipeIngredient(ingredient), response.Msg.GetRecipeIngredient())
}

// Test that we get a NotFound error when the recipe ingredient does not exist.
func (s *serviceTestSuite) TestGetRecipeIngredient_Err_NotFound() {
	s.querier.EXPECT().GetRecipeIngredient(s.ctx, mock.Anything).Return(store.RecipeIngredient{}, sql.ErrNoRows)

	response, err := s.s.GetRecipeIngredient(s.ctx, connect.NewRequest(&models.GetRecipeIngredientRequest{
		Id: testhelper.FakeRecipeIngredientID(),
	}))
	s.Error(err)
	s.Equal(connect.CodeNotFound, connect.CodeOf(err))
	s.Nil(response)
}

// Test that we can create a recipe ingredient.
func (s *serviceTestSuite) TestCreateRecipeIngredient_Ok() {
	recipeID := testhelper.FakeRecipeID()
	ingredient := testhelper.FakeStoreRecipeIngredient(recipeID)

	// First, verify the recipe exists
	s.querier.EXPECT().GetRecipe(s.ctx, store.GetRecipeParams{
		ID:       recipeID,
		TenantID: s.tenantIDParam(),
	}).Return(testhelper.FakeStoreRecipeWithTenant(s.tenantID), nil)

	s.querier.EXPECT().CreateRecipeIngredient(s.ctx, store.CreateRecipeIngredientParams{
		RecipeID: recipeID,
		Name:     ingredient.Name,
		Quantity: ingredient.Quantity,
		Unit:     ingredient.Unit,
	}).Return(ingredient, nil)

	response, err := s.s.CreateRecipeIngredient(s.ctx, connect.NewRequest(&models.CreateRecipeIngredientRequest{
		Recipe:   recipeID,
		Name:     ingredient.Name,
		Quantity: ingredient.Quantity,
		Unit:     ingredient.Unit,
	}))
	s.NoError(err)
	s.NotEmpty(response.Msg.GetRecipeIngredient().GetId())
}

// Test that we can't create a recipe ingredient when the recipe doesn't exist.
func (s *serviceTestSuite) TestCreateRecipeIngredient_Err_RecipeNotFound() {
	s.querier.EXPECT().GetRecipe(s.ctx, mock.Anything).Return(store.Recipe{}, sql.ErrNoRows)

	response, err := s.s.CreateRecipeIngredient(s.ctx, connect.NewRequest(&models.CreateRecipeIngredientRequest{
		Recipe: testhelper.FakeRecipeID(),
		Name:   "Test Ingredient",
	}))
	s.Error(err)
	s.Nil(response)
	s.Equal(connect.CodeNotFound, connect.CodeOf(err))
}

// Test that we can update a recipe ingredient.
func (s *serviceTestSuite) TestUpdateRecipeIngredient_Ok() {
	recipeID := testhelper.FakeRecipeID()

	actual := testhelper.FakeStoreRecipeIngredient(recipeID)
	input := actual
	input.Name = actual.Name + " updated"
	input.Quantity = "500"
	input.Unit = "ml"

	s.querier.EXPECT().GetRecipeIngredient(s.ctx, store.GetRecipeIngredientParams{
		ID:       actual.ID,
		TenantID: s.tenantIDParam(),
	}).Return(actual, nil)

	s.querier.EXPECT().UpdateRecipeIngredient(s.ctx, store.UpdateRecipeIngredientParams{
		ID:       actual.ID,
		Name:     input.Name,
		Quantity: input.Quantity,
		Unit:     input.Unit,
		Version:  actual.Version,
	}).Return(input, nil)

	response, err := s.s.UpdateRecipeIngredient(s.ctx, connect.NewRequest(&models.UpdateRecipeIngredientRequest{
		RecipeIngredient: translateRecipeIngredient(input),
		UpdateMask:       &fieldmaskpb.FieldMask{Paths: []string{"name", "quantity", "unit"}},
	}))
	s.NoError(err)
	s.EqualExportedValues(*translateRecipeIngredient(input), *response.Msg.GetRecipeIngredient())
}

// Test that we can't update a recipe ingredient with an invalid mask.
func (s *serviceTestSuite) TestUpdateRecipeIngredient_Err_InvalidMask() {
	response, err := s.s.UpdateRecipeIngredient(s.ctx, connect.NewRequest(&models.UpdateRecipeIngredientRequest{
		RecipeIngredient: testhelper.FakeRecipeIngredient(testhelper.FakeRecipeID()),
		UpdateMask:       &fieldmaskpb.FieldMask{Paths: []string{"id"}},
	}))
	s.Error(err)
	s.Nil(response)
	s.Equal(connect.CodeInvalidArgument, connect.CodeOf(err))
	s.Contains(err.Error(), "invalid mask")
}

// Test that we get a NotFound error when the recipe ingredient does not exist.
func (s *serviceTestSuite) TestUpdateRecipeIngredient_Err_NotFound() {
	ingredient := testhelper.FakeRecipeIngredient(testhelper.FakeRecipeID())

	s.querier.EXPECT().GetRecipeIngredient(s.ctx, store.GetRecipeIngredientParams{
		ID:       ingredient.GetId(),
		TenantID: s.tenantIDParam(),
	}).Return(store.RecipeIngredient{}, sql.ErrNoRows)

	response, err := s.s.UpdateRecipeIngredient(s.ctx, connect.NewRequest(&models.UpdateRecipeIngredientRequest{
		RecipeIngredient: ingredient,
		UpdateMask:       &fieldmaskpb.FieldMask{Paths: []string{"name"}},
	}))
	s.Error(err)
	s.Nil(response)
	s.Equal(connect.CodeNotFound, connect.CodeOf(err))
}

// Test that we can list all recipe ingredients of a recipe.
func (s *serviceTestSuite) TestListRecipeIngredients_Ok() {
	recipeID := testhelper.FakeRecipeID()

	i1 := testhelper.FakeStoreRecipeIngredient(recipeID)
	i2 := testhelper.FakeStoreRecipeIngredient(recipeID)

	s.querier.EXPECT().ListRecipeIngredientsByRecipe(s.ctx, store.ListRecipeIngredientsByRecipeParams{
		RecipeID: recipeID,
		TenantID: s.tenantIDParam(),
	}).Return([]store.RecipeIngredient{i1, i2}, nil)

	response, err := s.s.ListRecipeIngredients(s.ctx, connect.NewRequest(&models.ListRecipeIngredientsRequest{
		Recipe: recipeID,
	}))
	s.NoError(err)
	s.Equal([]*models.RecipeIngredient{translateRecipeIngredient(i1), translateRecipeIngredient(i2)}, response.Msg.GetRecipeIngredients())
}

// Test that we can delete a recipe ingredient.
func (s *serviceTestSuite) TestDeleteRecipeIngredient_Ok() {
	id := testhelper.FakeRecipeIngredientID()

	// First, verify the ingredient exists
	s.querier.EXPECT().GetRecipeIngredient(s.ctx, store.GetRecipeIngredientParams{
		ID:       id,
		TenantID: s.tenantIDParam(),
	}).Return(testhelper.FakeStoreRecipeIngredient(testhelper.FakeRecipeID()), nil)

	s.querier.EXPECT().DeleteRecipeIngredient(s.ctx, id).Return(store.RecipeIngredient{}, nil)

	_, err := s.s.DeleteRecipeIngredient(s.ctx, connect.NewRequest(&models.DeleteRecipeIngredientRequest{
		Id: id,
	}))
	s.NoError(err)
}

// Test that we get a NotFound error when the recipe ingredient does not exist.
func (s *serviceTestSuite) TestDeleteRecipeIngredient_Err_NotFound() {
	s.querier.EXPECT().GetRecipeIngredient(s.ctx, mock.Anything).Return(store.RecipeIngredient{}, sql.ErrNoRows)

	_, err := s.s.DeleteRecipeIngredient(s.ctx, connect.NewRequest(&models.DeleteRecipeIngredientRequest{
		Id: testhelper.FakeRecipeIngredientID(),
	}))
	s.Error(err)
	s.Equal(connect.CodeNotFound, connect.CodeOf(err))
}
