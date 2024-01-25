package v1

import (
	"context"
	"errors"

	"github.com/bufbuild/protovalidate-go"
	models "github.com/kiliandbigblue/octoback/gen/proto/go/octoback/core/v1"
	"github.com/segmentio/ksuid"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	log   *zap.Logger
	store Store
	v     *protovalidate.Validator

	models.UnimplementedServiceServer
}

func NewService(log *zap.Logger, store Store) *Service {
	v, _ := protovalidate.New()
	return &Service{
		store: store,
		log:   log,
		v:     v,
	}
}

// GetGroceryList returns a grocery list by its ID.
func (s *Service) GetGroceryList(_ context.Context, request *models.GetGroceryListRequest) (*models.GetGroceryListResponse, error) {
	s.log.Info("GetGroceryList")

	gli, err := s.store.GroceryList(request.GetId())
	if err != nil {
		if errors.Is(err, ErrNoSuchEntity) {
			return nil, status.Error(codes.NotFound, "grocery list not found")
		}
		panic(err)
	}
	return &models.GetGroceryListResponse{
		GroceryList: gli,
	}, nil
}

// Create a grocery list.
func (s *Service) CreateGroceryList(_ context.Context, request *models.CreateGroceryListRequest) (*models.CreateGroceryListResponse, error) {
	s.log.Info("CreateGroceryList")

	gl := &models.GroceryList{
		Id:   "gl_" + ksuid.New().String()[:8],
		Name: request.GetName(),
	}

	if err := s.store.SetGroceryList(gl); err != nil {
		var sve *StoreValidationError
		if errors.As(err, &sve) {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		panic(err)
	}

	return &models.CreateGroceryListResponse{
		GroceryList: gl,
	}, nil
}

// Update a grocery list.
func (s *Service) UpdateGroceryList(_ context.Context, request *models.UpdateGroceryListRequest) (*models.UpdateGroceryListResponse, error) {
	s.log.Info("UpdateGroceryList")

	fields := map[string]struct{}{}
	for _, field := range request.GetUpdateMask().GetPaths() {
		switch field {
		case "name", "items":
			fields[field] = struct{}{}
		}
	}
	if len(fields) == 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid mask")
	}

	gl, err := s.store.GroceryList(request.GetGroceryList().GetId())
	if err != nil {
		if errors.Is(err, ErrNoSuchEntity) {
			return nil, status.Error(codes.NotFound, "grocery list not found")
		}
		panic(err)
	}

	if _, ok := fields["name"]; ok {
		gl.Name = request.GetGroceryList().GetName()
	}
	if _, ok := fields["items"]; ok {
		gl.Items = request.GetGroceryList().GetItems()
	}

	if err := s.store.SetGroceryList(gl); err != nil {
		var sve *StoreValidationError
		if errors.As(err, &sve) {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		panic(err)
	}

	return &models.UpdateGroceryListResponse{
		GroceryList: gl,
	}, nil
}

// List grocery lists.
//
//nolint:unparam // Error is always nil.
func (s *Service) ListGroceryLists(_ context.Context, _ *models.ListGroceryListsRequest) (*models.ListGroceryListsResponse, error) {
	s.log.Info("ListGroceryLists")

	gls, err := s.store.GroceryLists()
	if err != nil {
		panic(err)
	}

	return &models.ListGroceryListsResponse{
		GroceryLists: gls,
	}, nil
}

// Delete a grocery list by its ID.
//
//nolint:unparam // Result is never used.
func (s *Service) DeleteGroceryList(_ context.Context, request *models.DeleteGroceryListRequest) (*models.DeleteGroceryListResponse, error) {
	s.log.Info("DeleteGroceryList")

	if err := s.store.DeleteGroceryList(request.GetId()); err != nil {
		if errors.Is(err, ErrNoSuchEntity) {
			return nil, status.Error(codes.NotFound, "grocery list not found")
		}
		panic(err)
	}

	return &models.DeleteGroceryListResponse{}, nil
}
