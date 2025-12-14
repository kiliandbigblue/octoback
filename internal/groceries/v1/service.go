package v1

import (
	"context"
	"database/sql"
	"errors"

	"connectrpc.com/connect"
	models "github.com/kiliandbigblue/octoback/gen/proto/go/octoback/groceries/v1"
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

// GetGroceryList returns a grocery list by its ID.
func (s *Service) GetGroceryList(ctx context.Context, r *connect.Request[models.GetGroceryListRequest]) (*connect.Response[models.GetGroceryListResponse], error) {
	log, _ := cloudzap.GetLogger(ctx)
	log.Info("GetGroceryList")

	gli, err := s.querier.GetGroceryList(ctx, store.GetGroceryListParams{
		ID:       r.Msg.GetId(),
		TenantID: tenantIDFromContext(ctx),
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, connect.NewError(connect.CodeNotFound, errors.New("grocery list not found"))
		}
		panic(err)
	}
	return connect.NewResponse(&models.GetGroceryListResponse{
		GroceryList: translateGroceryList(gli),
	}), nil
}

// Create a grocery list.
func (s *Service) CreateGroceryList(ctx context.Context, r *connect.Request[models.CreateGroceryListRequest]) (*connect.Response[models.CreateGroceryListResponse], error) {
	log, _ := cloudzap.GetLogger(ctx)
	log.Info("CreateGroceryList")

	gli, err := s.querier.CreateGroceryList(ctx, store.CreateGroceryListParams{
		Name:     r.Msg.GetName(),
		TenantID: tenantIDFromContext(ctx),
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("invalid name"))
		}
		panic(err)
	}

	return connect.NewResponse(&models.CreateGroceryListResponse{
		GroceryList: translateGroceryList(gli),
	}), nil
}

// Update a grocery list.
func (s *Service) UpdateGroceryList(ctx context.Context, r *connect.Request[models.UpdateGroceryListRequest]) (*connect.Response[models.UpdateGroceryListResponse], error) {
	log, _ := cloudzap.GetLogger(ctx)
	log.Info("UpdateGroceryList")

	tenantID := tenantIDFromContext(ctx)

	fields := map[string]struct{}{}
	for _, field := range r.Msg.GetUpdateMask().GetPaths() {
		if field == "name" {
			fields[field] = struct{}{}
		}
	}
	if len(fields) == 0 {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("invalid mask"))
	}

	gl, err := s.querier.GetGroceryList(ctx, store.GetGroceryListParams{
		ID:       r.Msg.GetGroceryList().GetId(),
		TenantID: tenantID,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, connect.NewError(connect.CodeNotFound, errors.New("grocery list not found"))
		}
		panic(err)
	}

	if _, ok := fields["name"]; ok {
		gl.Name = r.Msg.GetGroceryList().GetName()
	}

	if _, err := s.querier.UpdateGroceryList(ctx, store.UpdateGroceryListParams{
		ID:       gl.ID,
		Name:     gl.Name,
		TenantID: tenantID,
		Version:  gl.Version,
	}); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("invalid name"))
		}
		panic(err)
	}

	return connect.NewResponse(&models.UpdateGroceryListResponse{
		GroceryList: translateGroceryList(gl),
	}), nil
}

// List grocery lists.
//
//nolint:unparam // error is never used.
func (s *Service) ListGroceryLists(ctx context.Context, _ *connect.Request[models.ListGroceryListsRequest]) (*connect.Response[models.ListGroceryListsResponse], error) {
	log, _ := cloudzap.GetLogger(ctx)
	log.Info("ListGroceryLists")

	gls, err := s.querier.ListGroceryLists(ctx, tenantIDFromContext(ctx))
	if err != nil {
		panic(err)
	}

	glis := make([]*models.GroceryList, len(gls))
	for i, gl := range gls {
		glis[i] = translateGroceryList(gl)
	}

	return connect.NewResponse(&models.ListGroceryListsResponse{
		GroceryLists: glis,
	}), nil
}

// Delete a grocery list by its ID.
//
//nolint:unparam // Result is never used.
func (s *Service) DeleteGroceryList(ctx context.Context, r *connect.Request[models.DeleteGroceryListRequest]) (*connect.Response[models.DeleteGroceryListResponse], error) {
	log, _ := cloudzap.GetLogger(ctx)
	log.Info("DeleteGroceryList")

	if _, err := s.querier.DeleteGroceryList(ctx, store.DeleteGroceryListParams{
		ID:       r.Msg.GetId(),
		TenantID: tenantIDFromContext(ctx),
	}); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, connect.NewError(connect.CodeNotFound, errors.New("grocery list not found"))
		}
		panic(err)
	}

	return connect.NewResponse(&models.DeleteGroceryListResponse{}), nil
}

// Get the grocery item.
func (s *Service) GetGroceryItem(ctx context.Context, r *connect.Request[models.GetGroceryItemRequest]) (*connect.Response[models.GetGroceryItemResponse], error) {
	log, _ := cloudzap.GetLogger(ctx)
	log.Info("GetGroceryItem")

	git, err := s.querier.GetGroceryItem(ctx, store.GetGroceryItemParams{
		ID:       r.Msg.GetId(),
		TenantID: tenantIDFromContext(ctx),
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, connect.NewError(connect.CodeNotFound, errors.New("grocery item not found"))
		}
		panic(err)
	}
	return connect.NewResponse(&models.GetGroceryItemResponse{
		GroceryItem: translateGroceryItem(git),
	}), nil
}

// Create a grocery item.
func (s *Service) CreateGroceryItem(ctx context.Context, r *connect.Request[models.CreateGroceryItemRequest]) (*connect.Response[models.CreateGroceryItemResponse], error) {
	log, _ := cloudzap.GetLogger(ctx)
	log.Info("CreateGroceryItem")

	// First verify the grocery list belongs to the tenant
	if _, err := s.querier.GetGroceryList(ctx, store.GetGroceryListParams{
		ID:       r.Msg.GetGroceryList(),
		TenantID: tenantIDFromContext(ctx),
	}); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, connect.NewError(connect.CodeNotFound, errors.New("grocery list not found"))
		}
		panic(err)
	}

	git, err := s.querier.CreateGroceryItem(ctx, store.CreateGroceryItemParams{
		GroceryListID: r.Msg.GetGroceryList(),
		Name:          r.Msg.GetName(),
		Quantity:      int64(r.Msg.GetQuantity()),
		Checked:       boolToInt64(r.Msg.GetChecked()),
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("invalid name"))
		}
		panic(err)
	}

	return connect.NewResponse(&models.CreateGroceryItemResponse{
		GroceryItem: translateGroceryItem(git),
	}), nil
}

// Update a grocery item.
func (s *Service) UpdateGroceryItem(ctx context.Context, r *connect.Request[models.UpdateGroceryItemRequest]) (*connect.Response[models.UpdateGroceryItemResponse], error) {
	log, _ := cloudzap.GetLogger(ctx)
	log.Info("UpdateGroceryItem")

	tenantID := tenantIDFromContext(ctx)

	fields := map[string]struct{}{}
	for _, field := range r.Msg.GetUpdateMask().GetPaths() {
		switch field {
		case "name", "quantity", "checked":
			fields[field] = struct{}{}
		}
	}
	if len(fields) == 0 {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("invalid mask"))
	}

	git, err := s.querier.GetGroceryItem(ctx, store.GetGroceryItemParams{
		ID:       r.Msg.GetGroceryItem().GetId(),
		TenantID: tenantID,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, connect.NewError(connect.CodeNotFound, errors.New("grocery item not found"))
		}
		panic(err)
	}

	if _, ok := fields["name"]; ok {
		git.Name = r.Msg.GetGroceryItem().GetName()
	}
	if _, ok := fields["quantity"]; ok {
		git.Quantity = int64(r.Msg.GetGroceryItem().GetQuantity())
	}
	if _, ok := fields["checked"]; ok {
		git.Checked = boolToInt64(r.Msg.GetGroceryItem().GetChecked())
	}

	if _, err := s.querier.UpdateGroceryItem(ctx, store.UpdateGroceryItemParams{
		ID:       git.ID,
		Name:     git.Name,
		Quantity: git.Quantity,
		Checked:  git.Checked,
		Version:  git.Version,
	}); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("invalid name"))
		}
		panic(err)
	}

	return connect.NewResponse(&models.UpdateGroceryItemResponse{
		GroceryItem: translateGroceryItem(git),
	}), nil
}

// List grocery items
//
//nolint:unparam // error is never used.
func (s *Service) ListGroceryItems(ctx context.Context, r *connect.Request[models.ListGroceryItemsRequest]) (*connect.Response[models.ListGroceryItemsResponse], error) {
	log, _ := cloudzap.GetLogger(ctx)
	log.Info("ListGroceryItems")

	gitsi, err := s.querier.ListGroceryItemsByGroceryList(ctx, store.ListGroceryItemsByGroceryListParams{
		GroceryListID: r.Msg.GetGroceryList(),
		TenantID:      tenantIDFromContext(ctx),
	})
	if err != nil {
		panic(err)
	}

	gits := make([]*models.GroceryItem, len(gitsi))
	for i, git := range gitsi {
		gits[i] = translateGroceryItem(git)
	}

	return connect.NewResponse(&models.ListGroceryItemsResponse{
		GroceryItems: gits,
	}), nil
}

// Delete a grocery item.
//
//nolint:unparam // Result is never used.
func (s *Service) DeleteGroceryItem(ctx context.Context, r *connect.Request[models.DeleteGroceryItemRequest]) (*connect.Response[models.DeleteGroceryItemResponse], error) {
	log, _ := cloudzap.GetLogger(ctx)
	log.Info("DeleteGroceryItem")

	// First verify the grocery item belongs to the tenant
	if _, err := s.querier.GetGroceryItem(ctx, store.GetGroceryItemParams{
		ID:       r.Msg.GetId(),
		TenantID: tenantIDFromContext(ctx),
	}); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, connect.NewError(connect.CodeNotFound, errors.New("grocery item not found"))
		}
		panic(err)
	}

	if _, err := s.querier.DeleteGroceryItem(ctx, r.Msg.GetId()); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, connect.NewError(connect.CodeNotFound, errors.New("grocery item not found"))
		}
		panic(err)
	}

	return connect.NewResponse(&models.DeleteGroceryItemResponse{}), nil
}
