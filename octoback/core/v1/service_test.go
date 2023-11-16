package v1

import (
	"context"
	"testing"

	models "github.com/kiliandbigblue/octoback/gen/proto/go/octoback/core/v1"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type StubStore struct {
	groceryLists map[string]*models.GroceryList
}

func (s *StubStore) GetGroceryList(id string) (*models.GroceryList, error) {
	gli, ok := s.groceryLists[id]
	if !ok {
		return nil, ErrNoSuchEntity
	}
	return gli, nil
}

func TestService_OK_GetReceipe(t *testing.T) {
	gl := FakeGroceryList()

	store := &StubStore{
		groceryLists: map[string]*models.GroceryList{
			gl.GetId(): gl,
		},
	}

	s := NewService(zap.NewNop(), store)

	response, err := s.GetGroceryList(context.Background(), &models.GetGroceryListRequest{
		Id: gl.GetId(),
	})
	assert.NoError(t, err)
	assert.Equal(t, gl, response.GetGroceryList())
}

func TestService_Err_GetReceipe(t *testing.T) {
	store := &StubStore{}

	s := NewService(zap.NewNop(), store)

	response, err := s.GetGroceryList(context.Background(), &models.GetGroceryListRequest{
		Id: FakeGroceryListID(),
	})
	assert.Error(t, err)
	assert.Equal(t, codes.NotFound, status.Code(err))
	assert.Nil(t, response)
}
