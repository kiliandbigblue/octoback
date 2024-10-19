package v1

import (
	"context"
	"errors"

	"connectrpc.com/connect"
	models "github.com/kiliandbigblue/octoback/gen/proto/go/octoback/groceries/v1"
	"github.com/kiliandbigblue/octoback/internal/groceries/v1/store"
	"github.com/kiliandbigblue/octoback/internal/x/cloudzap"
)

type Service struct {
	s Store
}

func NewService(s Store) *Service {
	return &Service{
		s: s,
	}
}

// GetGroceryList returns a grocery list by its ID.
func (s *Service) GetGroceryList(ctx context.Context, request *connect.Request[models.GetGroceryListRequest]) (*connect.Response[models.GetGroceryListResponse], error) {
	log, _ := cloudzap.GetLogger(ctx)
	log.Info("GetGroceryList")

	gli, err := s.s.GroceryList(ctx, request.Msg.GetId())
	if err != nil {
		if errors.Is(err, store.ErrNoSuchEntity) {
			return nil, connect.NewError(connect.CodeNotFound, errors.New("grocery list not found"))
		}
		panic(err)
	}
	return connect.NewResponse(&models.GetGroceryListResponse{
		GroceryList: gli,
	}), nil
}

// Create a grocery list.
func (s *Service) CreateGroceryList(ctx context.Context, request *connect.Request[models.CreateGroceryListRequest]) (*connect.Response[models.CreateGroceryListResponse], error) {
	log, _ := cloudzap.GetLogger(ctx)
	log.Info("CreateGroceryList")

	gl := &models.GroceryList{
		Id:   1, // will be overwritten by the store
		Name: request.Msg.GetName(),
	}

	if _, err := s.s.CreateGroceryList(ctx, gl); err != nil {
		var sve *store.StoreValidationError
		if errors.As(err, &sve) {
			return nil, connect.NewError(connect.CodeInvalidArgument, err)
		}
		panic(err)
	}

	return connect.NewResponse(&models.CreateGroceryListResponse{
		GroceryList: gl,
	}), nil
}

// Update a grocery list.
func (s *Service) UpdateGroceryList(ctx context.Context, request *connect.Request[models.UpdateGroceryListRequest]) (*connect.Response[models.UpdateGroceryListResponse], error) {
	log, _ := cloudzap.GetLogger(ctx)
	log.Info("UpdateGroceryList")

	fields := map[string]struct{}{}
	for _, field := range request.Msg.GetUpdateMask().GetPaths() {
		switch field {
		case "name":
			fields[field] = struct{}{}
		}
	}
	if len(fields) == 0 {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("invalid mask"))
	}

	gl, err := s.s.GroceryList(ctx, request.Msg.GetGroceryList().GetId())
	if err != nil {
		if errors.Is(err, store.ErrNoSuchEntity) {
			return nil, connect.NewError(connect.CodeNotFound, errors.New("grocery list not found"))
		}
		panic(err)
	}

	if _, ok := fields["name"]; ok {
		gl.Name = request.Msg.GetGroceryList().GetName()
	}

	if _, err := s.s.UpdateGroceryList(ctx, gl); err != nil {
		var sve *store.StoreValidationError
		if errors.As(err, &sve) {
			return nil, connect.NewError(connect.CodeInvalidArgument, err)
		}
		panic(err)
	}

	return connect.NewResponse(&models.UpdateGroceryListResponse{
		GroceryList: gl,
	}), nil
}

// List grocery lists.
//
//nolint:unparam // error is never used.
func (s *Service) ListGroceryLists(ctx context.Context, _ *connect.Request[models.ListGroceryListsRequest]) (*connect.Response[models.ListGroceryListsResponse], error) {
	log, _ := cloudzap.GetLogger(ctx)
	log.Info("ListGroceryLists")

	gls, err := s.s.GroceryLists(ctx)
	if err != nil {
		panic(err)
	}

	return connect.NewResponse(&models.ListGroceryListsResponse{
		GroceryLists: gls,
	}), nil
}

// Delete a grocery list by its ID.
//
//nolint:unparam // Result is never used.
func (s *Service) DeleteGroceryList(ctx context.Context, request *connect.Request[models.DeleteGroceryListRequest]) (*connect.Response[models.DeleteGroceryListResponse], error) {
	log, _ := cloudzap.GetLogger(ctx)
	log.Info("DeleteGroceryList")

	if err := s.s.DeleteGroceryList(ctx, request.Msg.GetId()); err != nil {
		if errors.Is(err, store.ErrNoSuchEntity) {
			return nil, connect.NewError(connect.CodeNotFound, errors.New("grocery list not found"))
		}
		panic(err)
	}

	return connect.NewResponse(&models.DeleteGroceryListResponse{}), nil
}

// Create a grocery item
func (s *Service) CreateGroceryItem(ctx context.Context, request *connect.Request[models.CreateGroceryItemRequest]) (*connect.Response[models.CreateGroceryItemResponse], error) {
	log, _ := cloudzap.GetLogger(ctx)
	log.Info("CreateGroceryItem")

	gi := &models.GroceryItem{
		Id:       1, // will be overwritten by the store
		Name:     request.Msg.GetName(),
		Quantity: request.Msg.GetQuantity(),
		Checked:  request.Msg.GetChecked(),
	}

	gi, err := s.s.CreateGroceryItem(ctx, request.Msg.GetGroceryList(), gi)
	if err != nil {
		var sve *store.StoreValidationError
		if errors.As(err, &sve) {
			return nil, connect.NewError(connect.CodeInvalidArgument, err)
		}
		panic(err)
	}

	return connect.NewResponse(&models.CreateGroceryItemResponse{
		Item: gi,
	}), nil
}
