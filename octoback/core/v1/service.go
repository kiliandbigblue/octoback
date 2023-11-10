package v1

import (
	"context"

	corev1 "github.com/kiliandbigblue/octoback/gen/proto/go/octoback/core/v1"
)

type Service struct{}

// GetGroceryList returns a grocery list by its ID.
//
//nolint:unparam // The context is not used in this example.
func (s *Service) GetGroceryList(_ context.Context, request *corev1.GetGroceryListRequest) (*corev1.GetGroceryListResponse, error) {
	return &corev1.GetGroceryListResponse{
		GroceryList: &corev1.GroceryList{
			Id: request.Id,
		},
	}, nil
}
