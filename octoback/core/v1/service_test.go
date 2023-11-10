package v1

import (
	"context"
	"testing"

	corev1 "github.com/kiliandbigblue/octoback/gen/proto/go/octoback/core/v1"
	"github.com/stretchr/testify/assert"
)

func TestService_GetReceipe(t *testing.T) {
	s := &Service{}
	response, err := s.GetGroceryList(context.Background(), &corev1.GetGroceryListRequest{
		Id: "ID",
	})
	assert.NoError(t, err)
	assert.Equal(t, &corev1.GroceryList{
		Id: "ID",
	}, response.GetGroceryList())
}
