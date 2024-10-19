package v1

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"connectrpc.com/connect"
	"github.com/icrowley/fake"
	models "github.com/kiliandbigblue/octoback/gen/proto/go/octoback/groceries/v1"
	store "github.com/kiliandbigblue/octoback/gen/store/v1"
	"github.com/kiliandbigblue/octoback/internal/groceries/v1/mocks"
	"github.com/kiliandbigblue/octoback/internal/x/cloudzap"
	"github.com/kiliandbigblue/octoback/internal/x/testhelper"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type serviceTestSuite struct {
	suite.Suite
	log     *zap.Logger
	ctx     context.Context
	s       *Service
	querier *mocks.Querier
}

func TestServiceTestSuite(t *testing.T) {
	suite.Run(t, new(serviceTestSuite))
}

func (s *serviceTestSuite) SetupTest() {
	s.log = zap.NewNop()
	s.ctx = cloudzap.WithLogger(context.Background(), s.log)
	s.querier = mocks.NewQuerier(s.T())
	s.s = NewService(s.querier)
}

// Test that we can retrieve a grocery list by its ID.
func (s *serviceTestSuite) TestGetGroceryList_OK() {
	gl := testhelper.FakeStoreGroceryList()

	s.querier.EXPECT().GetGroceryList(s.ctx, gl.ID).Return(gl, nil)

	response, err := s.s.GetGroceryList(s.ctx, connect.NewRequest(&models.GetGroceryListRequest{
		Id: gl.ID,
	}))
	s.NoError(err)
	s.Equal(&models.GroceryList{
		Id:         gl.ID,
		Name:       gl.Name,
		CreateTime: timestamppb.New(gl.CreatedAt),
		Version:    gl.Version,
	}, response.Msg.GetGroceryList())
}

// Test that we get a NotFound error when the grocery list does not exist.
func (s *serviceTestSuite) TestGetGroceryList_Err_NotFound() {
	s.querier.EXPECT().GetGroceryList(s.ctx, mock.Anything).Return(store.GroceryList{}, sql.ErrNoRows)

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

	s.querier.EXPECT().CreateGroceryList(s.ctx, name).Return(store.GroceryList{
		ID:        testhelper.FakeGroceryListID(),
		Name:      name,
		CreatedAt: time.Now(),
		Version:   1,
	}, nil)

	response, err := s.s.CreateGroceryList(s.ctx, connect.NewRequest(&models.CreateGroceryListRequest{
		Name: name,
	}))
	s.NoError(err)
	s.NotEmpty(response.Msg.GetGroceryList().GetId())
	s.Equal("My first grocery list", response.Msg.GetGroceryList().GetName())
}

// Test that we can't create a grocery list with an empty name.
func (s *serviceTestSuite) TestCreateGroceryList_Err() {
	s.querier.EXPECT().CreateGroceryList(s.ctx, mock.Anything).Return(store.GroceryList{}, sql.ErrNoRows)

	response, err := s.s.CreateGroceryList(s.ctx, connect.NewRequest(&models.CreateGroceryListRequest{
		Name: "",
	}))
	s.Error(err)
	s.Nil(response)
}

// Test that we can update a grocery list.
func (s *serviceTestSuite) TestUpdateGroceryList_Ok() {
	actual := testhelper.FakeStoreGroceryList()

	input := actual
	input.Name = fake.Characters()

	s.querier.EXPECT().GetGroceryList(s.ctx, actual.ID).Return(actual, nil)

	s.querier.EXPECT().UpdateGroceryList(s.ctx, store.UpdateGroceryListParams{
		ID:      input.ID,
		Name:    input.Name,
		Version: input.Version,
	}).Return(input, nil)

	response, err := s.s.UpdateGroceryList(s.ctx, connect.NewRequest(&models.UpdateGroceryListRequest{
		GroceryList: translateGroceryList(input),
		UpdateMask:  &fieldmaskpb.FieldMask{Paths: []string{"name"}},
	}))
	s.NoError(err)
	s.EqualExportedValues(*translateGroceryList(input), *response.Msg.GetGroceryList())
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

	s.querier.EXPECT().GetGroceryList(s.ctx, gl.GetId()).Return(store.GroceryList{}, sql.ErrNoRows)

	response, err := s.s.UpdateGroceryList(s.ctx, connect.NewRequest(&models.UpdateGroceryListRequest{
		GroceryList: gl,
		UpdateMask:  &fieldmaskpb.FieldMask{Paths: []string{"name"}},
	}))
	s.Error(err)
	s.Nil(response)
	s.Equal(connect.CodeNotFound, connect.CodeOf(err))
}

// Test that we can list all grocery lists.
func (s *serviceTestSuite) TestListGroceryLists_Ok() {
	gl1 := testhelper.FakeStoreGroceryList()
	gl2 := testhelper.FakeStoreGroceryList()

	s.querier.EXPECT().ListGroceryLists(s.ctx).Return([]store.GroceryList{gl1, gl2}, nil)

	response, err := s.s.ListGroceryLists(s.ctx, connect.NewRequest(&models.ListGroceryListsRequest{}))
	s.NoError(err)
	s.Equal([]*models.GroceryList{translateGroceryList(gl1), translateGroceryList(gl2)}, response.Msg.GetGroceryLists())
}

// Test that we can delete a grocery list.
func (s *serviceTestSuite) TestDeleteGroceryList_Ok() {
	id := testhelper.FakeGroceryListID()

	s.querier.EXPECT().DeleteGroceryList(s.ctx, id).Return(store.GroceryList{}, nil)

	_, err := s.s.DeleteGroceryList(s.ctx, connect.NewRequest(&models.DeleteGroceryListRequest{
		Id: id,
	}))
	s.NoError(err)
}

// Test that we get a NotFound error when the grocery list does not exist.
func (s *serviceTestSuite) TestDeleteGroceryList_Err_NotFound() {
	s.querier.EXPECT().DeleteGroceryList(s.ctx, mock.Anything).Return(store.GroceryList{}, sql.ErrNoRows)

	_, err := s.s.DeleteGroceryList(s.ctx, connect.NewRequest(&models.DeleteGroceryListRequest{
		Id: testhelper.FakeGroceryListID(),
	}))
	s.Error(err)
	s.Equal(connect.CodeNotFound, connect.CodeOf(err))
}

// Test that we can retrieve a grocery item by its ID.
func (s *serviceTestSuite) TestGetGroceryItem_OK() {
	gl := testhelper.FakeStoreGroceryItem(testhelper.FakeGroceryListID())

	s.querier.EXPECT().GetGroceryItem(s.ctx, gl.ID).Return(gl, nil)

	response, err := s.s.GetGroceryItem(s.ctx, connect.NewRequest(&models.GetGroceryItemRequest{
		Id: gl.ID,
	}))
	s.NoError(err)
	s.Equal(&models.GroceryItem{
		Id:          gl.ID,
		GroceryList: gl.GroceryListID,
		Name:        gl.Name,
		Quantity:    gl.Quantity,
		Checked:     gl.Checked,
		CreateTime:  timestamppb.New(gl.CreatedAt),
		Version:     gl.Version,
	}, response.Msg.GetGroceryItem())
}

// Test that we get a NotFound error when the grocery item does not exist.
func (s *serviceTestSuite) TestGetGroceryItem_Err_NotFound() {
	s.querier.EXPECT().GetGroceryItem(s.ctx, mock.Anything).Return(store.GroceryItem{}, sql.ErrNoRows)

	response, err := s.s.GetGroceryItem(s.ctx, connect.NewRequest(&models.GetGroceryItemRequest{
		Id: testhelper.FakeGroceryItemID(),
	}))
	s.Error(err)
	s.Equal(connect.CodeNotFound, connect.CodeOf(err))
	s.Nil(response)
}

// Test that we can create a grocery item.
func (s *serviceTestSuite) TestCreateGroceryItem_Ok() {
	gil := testhelper.FakeGroceryListID()
	git := testhelper.FakeStoreGroceryItem(gil)

	s.querier.EXPECT().CreateGroceryItem(s.ctx, store.CreateGroceryItemParams{
		GroceryListID: gil,
		Name:          git.Name,
		Quantity:      git.Quantity,
		Checked:       git.Checked,
	}).Return(git, nil)

	response, err := s.s.CreateGroceryItem(s.ctx, connect.NewRequest(&models.CreateGroceryItemRequest{
		GroceryList: gil,
		Name:        git.Name,
		Quantity:    git.Quantity,
		Checked:     git.Checked,
	}))
	s.NoError(err)
	s.NotEmpty(response.Msg.GetGroceryItem().GetId())
}

// Test that we can't create a grocery item with an empty name.
func (s *serviceTestSuite) TestCreateGroceryItem_Err() {
	s.querier.EXPECT().CreateGroceryItem(s.ctx, mock.Anything).Return(store.GroceryItem{}, sql.ErrNoRows)

	response, err := s.s.CreateGroceryItem(s.ctx, connect.NewRequest(&models.CreateGroceryItemRequest{
		Name: "",
	}))
	s.Error(err)
	s.Nil(response)
}

// Test that we can update a grocery item.
func (s *serviceTestSuite) TestUpdateGroceryItem_Ok() {
	gil := testhelper.FakeGroceryListID()

	actual := testhelper.FakeStoreGroceryItem(gil)
	input := actual
	input.Name = actual.Name + " updated"
	input.Quantity = actual.Quantity + 1
	input.Checked = !actual.Checked

	s.querier.EXPECT().GetGroceryItem(s.ctx, actual.ID).Return(actual, nil)

	s.querier.EXPECT().UpdateGroceryItem(s.ctx, store.UpdateGroceryItemParams{
		ID:       actual.ID,
		Name:     input.Name,
		Quantity: input.Quantity,
		Checked:  input.Checked,
		Version:  actual.Version,
	}).Return(input, nil)

	response, err := s.s.UpdateGroceryItem(s.ctx, connect.NewRequest(&models.UpdateGroceryItemRequest{
		GroceryItem: translateGroceryItem(input),
		UpdateMask:  &fieldmaskpb.FieldMask{Paths: []string{"name", "quantity", "checked"}},
	}))
	s.NoError(err)
	s.EqualExportedValues(*translateGroceryItem(input), *response.Msg.GetGroceryItem())
}

// Test that we can't update a grocery item with an invalid mask.
func (s *serviceTestSuite) TestUpdateGroceryItem_Err_InvalidMask() {
	response, err := s.s.UpdateGroceryItem(s.ctx, connect.NewRequest(&models.UpdateGroceryItemRequest{
		GroceryItem: testhelper.FakeGroceryItem(testhelper.FakeGroceryListID()),
		UpdateMask:  &fieldmaskpb.FieldMask{Paths: []string{"id"}},
	}))
	s.Error(err)
	s.Nil(response)
	s.Equal(connect.CodeInvalidArgument, connect.CodeOf(err))
	s.Contains(err.Error(), "invalid mask")
}

// Test that we get a NotFound error when the grocery item does not exist.
func (s *serviceTestSuite) TestUpdateGroceryItem_Err_NotFound() {
	gl := testhelper.FakeGroceryItem(testhelper.FakeGroceryListID())

	s.querier.EXPECT().GetGroceryItem(s.ctx, gl.GetId()).Return(store.GroceryItem{}, sql.ErrNoRows)

	response, err := s.s.UpdateGroceryItem(s.ctx, connect.NewRequest(&models.UpdateGroceryItemRequest{
		GroceryItem: gl,
		UpdateMask:  &fieldmaskpb.FieldMask{Paths: []string{"name"}},
	}))
	s.Error(err)
	s.Nil(response)
	s.Equal(connect.CodeNotFound, connect.CodeOf(err))
}

// Test that we can list all grocery items of a grocery list.
func (s *serviceTestSuite) TestListGroceryItems_Ok() {
	gli := testhelper.FakeGroceryListID()

	gl1 := testhelper.FakeStoreGroceryItem(gli)
	gl2 := testhelper.FakeStoreGroceryItem(gli)

	s.querier.EXPECT().ListGroceryItemsByGroceryList(s.ctx, gli).Return([]store.GroceryItem{gl1, gl2}, nil)

	response, err := s.s.ListGroceryItems(s.ctx, connect.NewRequest(&models.ListGroceryItemsRequest{
		GroceryList: gli,
	}))
	s.NoError(err)
	s.Equal([]*models.GroceryItem{translateGroceryItem(gl1), translateGroceryItem(gl2)}, response.Msg.GetGroceryItems())
}

// Test that we can delete a grocery item.
func (s *serviceTestSuite) TestDeleteGroceryItem_Ok() {
	id := testhelper.FakeGroceryItemID()

	s.querier.EXPECT().DeleteGroceryItem(s.ctx, id).Return(store.GroceryItem{}, nil)

	_, err := s.s.DeleteGroceryItem(s.ctx, connect.NewRequest(&models.DeleteGroceryItemRequest{
		Id: id,
	}))
	s.NoError(err)
}

// Test that we get a NotFound error when the grocery item does not exist.
func (s *serviceTestSuite) TestDeleteGroceryItem_Err_NotFound() {
	s.querier.EXPECT().DeleteGroceryItem(s.ctx, mock.Anything).Return(store.GroceryItem{}, sql.ErrNoRows)

	_, err := s.s.DeleteGroceryItem(s.ctx, connect.NewRequest(&models.DeleteGroceryItemRequest{
		Id: testhelper.FakeGroceryItemID(),
	}))
	s.Error(err)
	s.Equal(connect.CodeNotFound, connect.CodeOf(err))
}
