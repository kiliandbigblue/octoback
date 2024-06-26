package v1

import (
	"testing"

	"github.com/kiliandbigblue/octoback/internal/core/v1/mocks"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type serviceTestSuite struct {
	suite.Suite
	s     *Service
	log   *zap.Logger
	store *mocks.Store
}

func TestServiceTestSuite(t *testing.T) {
	suite.Run(t, new(serviceTestSuite))
}

func (s *serviceTestSuite) SetupTest() {
	s.log = zap.NewNop()
	s.store = mocks.NewStore(s.T())
	s.s = NewService(s.log, s.store)
}

// // Test that we can retrieve a grocery list by its ID.
// func (s *serviceTestSuite) TestGetReceipe_Ok() {
// 	gl := FakeGroceryList()
//
// 	s.store.EXPECT().GroceryList(gl.GetId()).Return(gl, nil)
//
// 	response, err := s.s.GetGroceryList(context.Background(), &models.GetGroceryListRequest{
// 		Id: gl.GetId(),
// 	})
// 	s.NoError(err)
// 	s.Equal(gl, response.GetGroceryList())
// }
//
// // Test that we get a NotFound error when the grocery list does not exist.
// func (s *serviceTestSuite) TestGetReceipe_Err() {
// 	s.store.EXPECT().GroceryList(mock.Anything).Return(nil, ErrNoSuchEntity)
//
// 	response, err := s.s.GetGroceryList(context.Background(), &models.GetGroceryListRequest{
// 		Id: FakeGroceryListID(),
// 	})
// 	s.Error(err)
// 	s.Equal(codes.NotFound, status.Code(err))
// 	s.Nil(response)
// }
//
// // Test that we can create a grocery list.
// func (s *serviceTestSuite) TestCreateGroceryList_Ok() {
// 	name := "My first grocery list"
//
// 	s.store.EXPECT().SetGroceryList(mock.MatchedBy(func(gl *models.GroceryList) bool {
// 		s.NotEmpty(gl.GetId())
// 		return s.True(proto.Equal(&models.GroceryList{
// 			Id:   gl.GetId(),
// 			Name: name,
// 		}, gl))
// 	})).Return(nil)
//
// 	response, err := s.s.CreateGroceryList(context.Background(), &models.CreateGroceryListRequest{
// 		Name: name,
// 	})
// 	s.NoError(err)
// 	s.NotEmpty(response.GetGroceryList().GetId())
// 	s.Equal("My first grocery list", response.GetGroceryList().GetName())
// 	s.Empty(response.GetGroceryList().GetItems())
// }
//
// // Test that we can't create a grocery list with an empty name.
// func (s *serviceTestSuite) TestCreateGroceryList_Err_StoreValidation() {
// 	s.store.EXPECT().SetGroceryList(mock.Anything).Return(&StoreValidationError{}).Once()
//
// 	response, err := s.s.CreateGroceryList(context.Background(), &models.CreateGroceryListRequest{
// 		Name: "",
// 	})
// 	s.Error(err)
// 	s.Nil(response)
// }
//
// // Test that we can update a grocery list.
// func (s *serviceTestSuite) TestUpdateGroceryList_Ok() {
// 	actual := FakeGroceryList()
//
// 	input := FakeGroceryList()
// 	input.Id = actual.GetId()
//
// 	s.store.EXPECT().GroceryList(actual.GetId()).Return(actual, nil)
// 	s.store.EXPECT().SetGroceryList(input).Return(nil)
//
// 	response, err := s.s.UpdateGroceryList(context.Background(), &models.UpdateGroceryListRequest{
// 		GroceryList: input,
// 		UpdateMask:  &fieldmaskpb.FieldMask{Paths: []string{"name", "items"}},
// 	})
// 	s.NoError(err)
// 	s.EqualExportedValues(*input, *response.GetGroceryList()) //nolint:govet //lock value
// }
//
// // Test that we can't update a grocery list with an invalid mask.
// func (s *serviceTestSuite) TestUpdateGroceryList_Err_InvalidMask() {
// 	response, err := s.s.UpdateGroceryList(context.Background(), &models.UpdateGroceryListRequest{
// 		GroceryList: FakeGroceryList(),
// 		UpdateMask:  &fieldmaskpb.FieldMask{Paths: []string{"id"}},
// 	})
// 	s.Error(err)
// 	s.Nil(response)
// 	s.Equal(codes.InvalidArgument, status.Code(err))
// 	s.Contains(err.Error(), "invalid mask")
// }
//
// // Test that we get a NotFound error when the grocery list does not exist.
// func (s *serviceTestSuite) TestUpdateGroceryList_Err_NotFound() {
// 	gl := FakeGroceryList()
//
// 	s.store.EXPECT().GroceryList(gl.GetId()).Return(nil, ErrNoSuchEntity)
//
// 	response, err := s.s.UpdateGroceryList(context.Background(), &models.UpdateGroceryListRequest{
// 		GroceryList: gl,
// 		UpdateMask:  &fieldmaskpb.FieldMask{Paths: []string{"name"}},
// 	})
// 	s.Error(err)
// 	s.Nil(response)
// 	s.Equal(codes.NotFound, status.Code(err))
// }
//
// // Test that we can't update a grocery list if the proto validation fails.
// func (s *serviceTestSuite) TestUpdateGroceryList_Err_Validation() {
// 	gl := FakeGroceryList()
//
// 	input := FakeGroceryList()
// 	input.Id = gl.GetId()
// 	input.Name = ""
//
// 	s.store.EXPECT().GroceryList(gl.GetId()).Return(gl, nil)
// 	s.store.EXPECT().SetGroceryList(&models.GroceryList{
// 		Id:    gl.GetId(),
// 		Name:  "",
// 		Items: gl.GetItems(),
// 	}).Return(&StoreValidationError{})
//
// 	response, err := s.s.UpdateGroceryList(context.Background(), &models.UpdateGroceryListRequest{
// 		GroceryList: input,
// 		UpdateMask:  &fieldmaskpb.FieldMask{Paths: []string{"name"}},
// 	})
// 	s.Error(err)
// 	s.Nil(response)
// 	s.Equal(codes.InvalidArgument, status.Code(err))
// }
//
// // Test that we can list all grocery lists.
// func (s *serviceTestSuite) TestListGroceryLists_Ok() {
// 	gl1 := FakeGroceryList()
// 	gl2 := FakeGroceryList()
//
// 	s.store.EXPECT().GroceryLists().Return([]*models.GroceryList{gl1, gl2}, nil)
//
// 	response, err := s.s.ListGroceryLists(context.Background(), &models.ListGroceryListsRequest{})
// 	s.NoError(err)
// 	s.Equal([]*models.GroceryList{gl1, gl2}, response.GetGroceryLists())
// }
//
// // Test that we can delete a grocery list.
// func (s *serviceTestSuite) TestDeleteGroceryList_Ok() {
// 	gl := FakeGroceryList()
//
// 	s.store.EXPECT().DeleteGroceryList(gl.GetId()).Return(nil)
//
// 	_, err := s.s.DeleteGroceryList(context.Background(), &models.DeleteGroceryListRequest{
// 		Id: gl.GetId(),
// 	})
// 	s.NoError(err)
// }
//
// // Test that we get a NotFound error when the grocery list does not exist.
// func (s *serviceTestSuite) TestDeleteGroceryList_Err_NotFound() {
// 	s.store.EXPECT().DeleteGroceryList(mock.Anything).Return(ErrNoSuchEntity)
//
// 	_, err := s.s.DeleteGroceryList(context.Background(), &models.DeleteGroceryListRequest{
// 		Id: FakeGroceryListID(),
// 	})
// 	s.Error(err)
// 	s.Equal(codes.NotFound, status.Code(err))
// }
