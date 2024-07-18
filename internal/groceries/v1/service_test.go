package v1

import (
	"context"
	"testing"

	"connectrpc.com/connect"
	models "github.com/kiliandbigblue/octoback/gen/proto/go/octoback/groceries/v1"
	"github.com/kiliandbigblue/octoback/internal/groceries/v1/mocks"
	"github.com/kiliandbigblue/octoback/internal/groceries/v1/store"
	"github.com/kiliandbigblue/octoback/internal/x/cloudzap"
	"github.com/kiliandbigblue/octoback/internal/x/testhelper"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type serviceTestSuite struct {
	suite.Suite
	log   *zap.Logger
	ctx   context.Context
	s     *Service
	store *mocks.Store
}

func TestServiceTestSuite(t *testing.T) {
	suite.Run(t, new(serviceTestSuite))
}

func (s *serviceTestSuite) SetupTest() {
	s.log = zap.NewNop()
	s.ctx = cloudzap.WithLogger(context.Background(), s.log)
	s.store = mocks.NewStore(s.T())
	s.s = NewService(s.store)
}

// Test that we can retrieve a grocery list by its ID.
func (s *serviceTestSuite) TestGetReceipe_Ok() {
	gl := testhelper.FakeGroceryList()

	s.store.EXPECT().GroceryList(gl.GetId()).Return(gl, nil)

	response, err := s.s.GetGroceryList(s.ctx, connect.NewRequest(&models.GetGroceryListRequest{
		Id: gl.GetId(),
	}))
	s.NoError(err)
	s.Equal(gl, response.Msg.GetGroceryList())
}

// Test that we get a NotFound error when the grocery list does not exist.
func (s *serviceTestSuite) TestGetReceipe_Err() {
	s.store.EXPECT().GroceryList(mock.Anything).Return(nil, store.ErrNoSuchEntity)

	response, err := s.s.GetGroceryList(s.ctx, connect.NewRequest(&models.GetGroceryListRequest{
		Id: testhelper.FakeGroceryListID(),
	}))
	s.Error(err)
	s.Equal(connect.CodeNotFound, connect.CodeOf(err))
	s.Nil(response)
}

// Test that we can create a grocery list.
func (s *serviceTestSuite) TestCreateGroceryList_Ok() {
	name := "My first grocery list"

	s.store.EXPECT().SetGroceryList(mock.MatchedBy(func(gl *models.GroceryList) bool {
		s.NotEmpty(gl.GetId())
		return s.True(proto.Equal(&models.GroceryList{
			Id:   gl.GetId(),
			Name: name,
		}, gl))
	})).Return(nil)

	response, err := s.s.CreateGroceryList(s.ctx, connect.NewRequest(&models.CreateGroceryListRequest{
		Name: name,
	}))
	s.NoError(err)
	s.NotEmpty(response.Msg.GetGroceryList().GetId())
	s.Equal("My first grocery list", response.Msg.GetGroceryList().GetName())
	s.Empty(response.Msg.GetGroceryList().GetItems())
}

// Test that we can't create a grocery list with an empty name.
func (s *serviceTestSuite) TestCreateGroceryList_Err_StoreValidation() {
	s.store.EXPECT().SetGroceryList(mock.Anything).Return(&store.StoreValidationError{}).Once()

	response, err := s.s.CreateGroceryList(s.ctx, connect.NewRequest(&models.CreateGroceryListRequest{
		Name: "",
	}))
	s.Error(err)
	s.Nil(response)
}

// Test that we can update a grocery list.
func (s *serviceTestSuite) TestUpdateGroceryList_Ok() {
	actual := testhelper.FakeGroceryList()

	input := testhelper.FakeGroceryList()
	input.Id = actual.GetId()

	s.store.EXPECT().GroceryList(actual.GetId()).Return(actual, nil)
	s.store.EXPECT().SetGroceryList(input).Return(nil)

	response, err := s.s.UpdateGroceryList(s.ctx, connect.NewRequest(&models.UpdateGroceryListRequest{
		GroceryList: input,
		UpdateMask:  &fieldmaskpb.FieldMask{Paths: []string{"name", "items"}},
	}))
	s.NoError(err)
	s.EqualExportedValues(*input, *response.Msg.GetGroceryList()) //nolint:govet //lock value
}

// Test that we can't update a grocery list with an invalid mask.
func (s *serviceTestSuite) TestUpdateGroceryList_Err_InvalidMask() {
	response, err := s.s.UpdateGroceryList(s.ctx, connect.NewRequest(&models.UpdateGroceryListRequest{
		GroceryList: testhelper.FakeGroceryList(),
		UpdateMask:  &fieldmaskpb.FieldMask{Paths: []string{"id"}},
	}))
	s.Error(err)
	s.Nil(response)
	s.Equal(connect.CodeInvalidArgument, connect.CodeOf(err))
	s.Contains(err.Error(), "invalid mask")
}

// Test that we get a NotFound error when the grocery list does not exist.
func (s *serviceTestSuite) TestUpdateGroceryList_Err_NotFound() {
	gl := testhelper.FakeGroceryList()

	s.store.EXPECT().GroceryList(gl.GetId()).Return(nil, store.ErrNoSuchEntity)

	response, err := s.s.UpdateGroceryList(s.ctx, connect.NewRequest(&models.UpdateGroceryListRequest{
		GroceryList: gl,
		UpdateMask:  &fieldmaskpb.FieldMask{Paths: []string{"name"}},
	}))
	s.Error(err)
	s.Nil(response)
	s.Equal(connect.CodeNotFound, connect.CodeOf(err))
}

// Test that we can't update a grocery list if the proto validation fails.
func (s *serviceTestSuite) TestUpdateGroceryList_Err_Validation() {
	gl := testhelper.FakeGroceryList()

	input := testhelper.FakeGroceryList()
	input.Id = gl.GetId()
	input.Name = ""

	s.store.EXPECT().GroceryList(gl.GetId()).Return(gl, nil)
	s.store.EXPECT().SetGroceryList(&models.GroceryList{
		Id:    gl.GetId(),
		Name:  "",
		Items: gl.GetItems(),
	}).Return(&store.StoreValidationError{})

	response, err := s.s.UpdateGroceryList(s.ctx, connect.NewRequest(&models.UpdateGroceryListRequest{
		GroceryList: input,
		UpdateMask:  &fieldmaskpb.FieldMask{Paths: []string{"name"}},
	}))
	s.Error(err)
	s.Nil(response)
	s.Equal(connect.CodeInvalidArgument, connect.CodeOf(err))
}

// Test that we can list all grocery lists.
func (s *serviceTestSuite) TestListGroceryLists_Ok() {
	gl1 := testhelper.FakeGroceryList()
	gl2 := testhelper.FakeGroceryList()

	s.store.EXPECT().GroceryLists().Return([]*models.GroceryList{gl1, gl2}, nil)

	response, err := s.s.ListGroceryLists(s.ctx, connect.NewRequest(&models.ListGroceryListsRequest{}))
	s.NoError(err)
	s.Equal([]*models.GroceryList{gl1, gl2}, response.Msg.GetGroceryLists())
}

// Test that we can delete a grocery list.
func (s *serviceTestSuite) TestDeleteGroceryList_Ok() {
	gl := testhelper.FakeGroceryList()

	s.store.EXPECT().DeleteGroceryList(gl.GetId()).Return(nil)

	_, err := s.s.DeleteGroceryList(s.ctx, connect.NewRequest(&models.DeleteGroceryListRequest{
		Id: gl.GetId(),
	}))
	s.NoError(err)
}

// Test that we get a NotFound error when the grocery list does not exist.
func (s *serviceTestSuite) TestDeleteGroceryList_Err_NotFound() {
	s.store.EXPECT().DeleteGroceryList(mock.Anything).Return(store.ErrNoSuchEntity)

	_, err := s.s.DeleteGroceryList(s.ctx, connect.NewRequest(&models.DeleteGroceryListRequest{
		Id: testhelper.FakeGroceryListID(),
	}))
	s.Error(err)
	s.Equal(connect.CodeNotFound, connect.CodeOf(err))
}
