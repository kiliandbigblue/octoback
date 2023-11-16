package v1

import (
	"context"
	"errors"

	models "github.com/kiliandbigblue/octoback/gen/proto/go/octoback/core/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var ErrNoSuchEntity = errors.New("no such entity")

type Store interface {
	GetGroceryList(id string) (*models.GroceryList, error)
}

type Service struct {
	log   *zap.Logger
	store Store

	models.UnimplementedServiceServer
}

func NewService(log *zap.Logger, store Store) *Service {
	return &Service{
		store: store,
		log:   log,
	}
}

// GetGroceryList returns a grocery list by its ID.
//
//nolint:unparam // The context is not used in this example.
func (s *Service) GetGroceryList(_ context.Context, request *models.GetGroceryListRequest) (*models.GetGroceryListResponse, error) {
	s.log.Info("GetGroceryList")

	gli, err := s.store.GetGroceryList(request.GetId())
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
