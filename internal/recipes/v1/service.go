package v1

import (
	"context"
	"database/sql"
	"errors"

	"connectrpc.com/connect"
	models "github.com/kiliandbigblue/octoback/gen/proto/go/octoback/recipes/v1"
	store "github.com/kiliandbigblue/octoback/gen/store/v1"
	"github.com/kiliandbigblue/octoback/internal/x/cloudzap"
	"github.com/kiliandbigblue/octoback/internal/x/tenant"
)

type Service struct {
	querier store.Querier
}

func NewService(querier store.Querier) *Service {
	return &Service{
		querier: querier,
	}
}

// tenantIDFromContext extracts tenant_id from context and converts it to sql.NullInt64.
func tenantIDFromContext(ctx context.Context) sql.NullInt64 {
	id := tenant.MustTenantIDFromContext(ctx)
	return sql.NullInt64{Int64: id, Valid: true}
}

// GetRecipe returns a recipe by its ID.
func (s *Service) GetRecipe(ctx context.Context, r *connect.Request[models.GetRecipeRequest]) (*connect.Response[models.GetRecipeResponse], error) {
	log, _ := cloudzap.GetLogger(ctx)
	log.Info("GetRecipe")

	recipe, err := s.querier.GetRecipe(ctx, store.GetRecipeParams{
		ID:       r.Msg.GetId(),
		TenantID: tenantIDFromContext(ctx),
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, connect.NewError(connect.CodeNotFound, errors.New("recipe not found"))
		}
		panic(err)
	}
	return connect.NewResponse(&models.GetRecipeResponse{
		Recipe: translateRecipe(recipe),
	}), nil
}

// CreateRecipe creates a new recipe.
func (s *Service) CreateRecipe(ctx context.Context, r *connect.Request[models.CreateRecipeRequest]) (*connect.Response[models.CreateRecipeResponse], error) {
	log, _ := cloudzap.GetLogger(ctx)
	log.Info("CreateRecipe")

	recipe, err := s.querier.CreateRecipe(ctx, store.CreateRecipeParams{
		Name:            r.Msg.GetName(),
		Description:     r.Msg.GetDescription(),
		Servings:        int64(r.Msg.GetServings()),
		PrepTimeMinutes: int64(r.Msg.GetPrepTimeMinutes()),
		CookTimeMinutes: int64(r.Msg.GetCookTimeMinutes()),
		TenantID:        tenantIDFromContext(ctx),
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("invalid recipe data"))
		}
		panic(err)
	}

	return connect.NewResponse(&models.CreateRecipeResponse{
		Recipe: translateRecipe(recipe),
	}), nil
}

// UpdateRecipe updates a recipe.
func (s *Service) UpdateRecipe(ctx context.Context, r *connect.Request[models.UpdateRecipeRequest]) (*connect.Response[models.UpdateRecipeResponse], error) {
	log, _ := cloudzap.GetLogger(ctx)
	log.Info("UpdateRecipe")

	tenantID := tenantIDFromContext(ctx)

	fields := map[string]struct{}{}
	for _, field := range r.Msg.GetUpdateMask().GetPaths() {
		switch field {
		case "name", "description", "servings", "prep_time_minutes", "cook_time_minutes":
			fields[field] = struct{}{}
		}
	}
	if len(fields) == 0 {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("invalid mask"))
	}

	recipe, err := s.querier.GetRecipe(ctx, store.GetRecipeParams{
		ID:       r.Msg.GetRecipe().GetId(),
		TenantID: tenantID,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, connect.NewError(connect.CodeNotFound, errors.New("recipe not found"))
		}
		panic(err)
	}

	if _, ok := fields["name"]; ok {
		recipe.Name = r.Msg.GetRecipe().GetName()
	}
	if _, ok := fields["description"]; ok {
		recipe.Description = r.Msg.GetRecipe().GetDescription()
	}
	if _, ok := fields["servings"]; ok {
		recipe.Servings = int64(r.Msg.GetRecipe().GetServings())
	}
	if _, ok := fields["prep_time_minutes"]; ok {
		recipe.PrepTimeMinutes = int64(r.Msg.GetRecipe().GetPrepTimeMinutes())
	}
	if _, ok := fields["cook_time_minutes"]; ok {
		recipe.CookTimeMinutes = int64(r.Msg.GetRecipe().GetCookTimeMinutes())
	}

	if _, err := s.querier.UpdateRecipe(ctx, store.UpdateRecipeParams{
		ID:              recipe.ID,
		Name:            recipe.Name,
		Description:     recipe.Description,
		Servings:        recipe.Servings,
		PrepTimeMinutes: recipe.PrepTimeMinutes,
		CookTimeMinutes: recipe.CookTimeMinutes,
		TenantID:        tenantID,
		Version:         recipe.Version,
	}); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("invalid recipe data"))
		}
		panic(err)
	}

	return connect.NewResponse(&models.UpdateRecipeResponse{
		Recipe: translateRecipe(recipe),
	}), nil
}

// ListRecipes lists all recipes.
//
//nolint:unparam // error is never used.
func (s *Service) ListRecipes(ctx context.Context, _ *connect.Request[models.ListRecipesRequest]) (*connect.Response[models.ListRecipesResponse], error) {
	log, _ := cloudzap.GetLogger(ctx)
	log.Info("ListRecipes")

	recipes, err := s.querier.ListRecipes(ctx, tenantIDFromContext(ctx))
	if err != nil {
		panic(err)
	}

	result := make([]*models.Recipe, len(recipes))
	for i, recipe := range recipes {
		result[i] = translateRecipe(recipe)
	}

	return connect.NewResponse(&models.ListRecipesResponse{
		Recipes: result,
	}), nil
}

// DeleteRecipe deletes a recipe by its ID.
//
//nolint:unparam // Result is never used.
func (s *Service) DeleteRecipe(ctx context.Context, r *connect.Request[models.DeleteRecipeRequest]) (*connect.Response[models.DeleteRecipeResponse], error) {
	log, _ := cloudzap.GetLogger(ctx)
	log.Info("DeleteRecipe")

	if _, err := s.querier.DeleteRecipe(ctx, store.DeleteRecipeParams{
		ID:       r.Msg.GetId(),
		TenantID: tenantIDFromContext(ctx),
	}); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, connect.NewError(connect.CodeNotFound, errors.New("recipe not found"))
		}
		panic(err)
	}

	return connect.NewResponse(&models.DeleteRecipeResponse{}), nil
}

// GetRecipeIngredient returns a recipe ingredient by its ID.
func (s *Service) GetRecipeIngredient(ctx context.Context, r *connect.Request[models.GetRecipeIngredientRequest]) (*connect.Response[models.GetRecipeIngredientResponse], error) {
	log, _ := cloudzap.GetLogger(ctx)
	log.Info("GetRecipeIngredient")

	ingredient, err := s.querier.GetRecipeIngredient(ctx, store.GetRecipeIngredientParams{
		ID:       r.Msg.GetId(),
		TenantID: tenantIDFromContext(ctx),
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, connect.NewError(connect.CodeNotFound, errors.New("recipe ingredient not found"))
		}
		panic(err)
	}
	return connect.NewResponse(&models.GetRecipeIngredientResponse{
		RecipeIngredient: translateRecipeIngredient(ingredient),
	}), nil
}

// CreateRecipeIngredient creates a new recipe ingredient.
func (s *Service) CreateRecipeIngredient(ctx context.Context, r *connect.Request[models.CreateRecipeIngredientRequest]) (*connect.Response[models.CreateRecipeIngredientResponse], error) {
	log, _ := cloudzap.GetLogger(ctx)
	log.Info("CreateRecipeIngredient")

	// First verify the recipe belongs to the tenant
	if _, err := s.querier.GetRecipe(ctx, store.GetRecipeParams{
		ID:       r.Msg.GetRecipe(),
		TenantID: tenantIDFromContext(ctx),
	}); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, connect.NewError(connect.CodeNotFound, errors.New("recipe not found"))
		}
		panic(err)
	}

	ingredient, err := s.querier.CreateRecipeIngredient(ctx, store.CreateRecipeIngredientParams{
		RecipeID: r.Msg.GetRecipe(),
		Name:     r.Msg.GetName(),
		Quantity: r.Msg.GetQuantity(),
		Unit:     r.Msg.GetUnit(),
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("invalid ingredient data"))
		}
		panic(err)
	}

	return connect.NewResponse(&models.CreateRecipeIngredientResponse{
		RecipeIngredient: translateRecipeIngredient(ingredient),
	}), nil
}

// UpdateRecipeIngredient updates a recipe ingredient.
func (s *Service) UpdateRecipeIngredient(ctx context.Context, r *connect.Request[models.UpdateRecipeIngredientRequest]) (*connect.Response[models.UpdateRecipeIngredientResponse], error) {
	log, _ := cloudzap.GetLogger(ctx)
	log.Info("UpdateRecipeIngredient")

	tenantID := tenantIDFromContext(ctx)

	fields := map[string]struct{}{}
	for _, field := range r.Msg.GetUpdateMask().GetPaths() {
		switch field {
		case "name", "quantity", "unit":
			fields[field] = struct{}{}
		}
	}
	if len(fields) == 0 {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("invalid mask"))
	}

	ingredient, err := s.querier.GetRecipeIngredient(ctx, store.GetRecipeIngredientParams{
		ID:       r.Msg.GetRecipeIngredient().GetId(),
		TenantID: tenantID,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, connect.NewError(connect.CodeNotFound, errors.New("recipe ingredient not found"))
		}
		panic(err)
	}

	if _, ok := fields["name"]; ok {
		ingredient.Name = r.Msg.GetRecipeIngredient().GetName()
	}
	if _, ok := fields["quantity"]; ok {
		ingredient.Quantity = r.Msg.GetRecipeIngredient().GetQuantity()
	}
	if _, ok := fields["unit"]; ok {
		ingredient.Unit = r.Msg.GetRecipeIngredient().GetUnit()
	}

	if _, err := s.querier.UpdateRecipeIngredient(ctx, store.UpdateRecipeIngredientParams{
		ID:       ingredient.ID,
		Name:     ingredient.Name,
		Quantity: ingredient.Quantity,
		Unit:     ingredient.Unit,
		Version:  ingredient.Version,
	}); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("invalid ingredient data"))
		}
		panic(err)
	}

	return connect.NewResponse(&models.UpdateRecipeIngredientResponse{
		RecipeIngredient: translateRecipeIngredient(ingredient),
	}), nil
}

// ListRecipeIngredients lists all ingredients for a recipe.
//
//nolint:unparam // error is never used.
func (s *Service) ListRecipeIngredients(ctx context.Context, r *connect.Request[models.ListRecipeIngredientsRequest]) (*connect.Response[models.ListRecipeIngredientsResponse], error) {
	log, _ := cloudzap.GetLogger(ctx)
	log.Info("ListRecipeIngredients")

	ingredients, err := s.querier.ListRecipeIngredientsByRecipe(ctx, store.ListRecipeIngredientsByRecipeParams{
		RecipeID: r.Msg.GetRecipe(),
		TenantID: tenantIDFromContext(ctx),
	})
	if err != nil {
		panic(err)
	}

	result := make([]*models.RecipeIngredient, len(ingredients))
	for i, ingredient := range ingredients {
		result[i] = translateRecipeIngredient(ingredient)
	}

	return connect.NewResponse(&models.ListRecipeIngredientsResponse{
		RecipeIngredients: result,
	}), nil
}

// DeleteRecipeIngredient deletes a recipe ingredient.
//
//nolint:unparam // Result is never used.
func (s *Service) DeleteRecipeIngredient(ctx context.Context, r *connect.Request[models.DeleteRecipeIngredientRequest]) (*connect.Response[models.DeleteRecipeIngredientResponse], error) {
	log, _ := cloudzap.GetLogger(ctx)
	log.Info("DeleteRecipeIngredient")

	// First verify the ingredient belongs to the tenant
	if _, err := s.querier.GetRecipeIngredient(ctx, store.GetRecipeIngredientParams{
		ID:       r.Msg.GetId(),
		TenantID: tenantIDFromContext(ctx),
	}); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, connect.NewError(connect.CodeNotFound, errors.New("recipe ingredient not found"))
		}
		panic(err)
	}

	if _, err := s.querier.DeleteRecipeIngredient(ctx, r.Msg.GetId()); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, connect.NewError(connect.CodeNotFound, errors.New("recipe ingredient not found"))
		}
		panic(err)
	}

	return connect.NewResponse(&models.DeleteRecipeIngredientResponse{}), nil
}
