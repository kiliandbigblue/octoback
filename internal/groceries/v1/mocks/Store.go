// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	context "context"

	groceriesv1 "github.com/kiliandbigblue/octoback/gen/proto/go/octoback/groceries/v1"
	mock "github.com/stretchr/testify/mock"
)

// Store is an autogenerated mock type for the Store type
type Store struct {
	mock.Mock
}

type Store_Expecter struct {
	mock *mock.Mock
}

func (_m *Store) EXPECT() *Store_Expecter {
	return &Store_Expecter{mock: &_m.Mock}
}

// CreateGroceryList provides a mock function with given fields: ctx, r
func (_m *Store) CreateGroceryList(ctx context.Context, r *groceriesv1.GroceryList) (*groceriesv1.GroceryList, error) {
	ret := _m.Called(ctx, r)

	if len(ret) == 0 {
		panic("no return value specified for CreateGroceryList")
	}

	var r0 *groceriesv1.GroceryList
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *groceriesv1.GroceryList) (*groceriesv1.GroceryList, error)); ok {
		return rf(ctx, r)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *groceriesv1.GroceryList) *groceriesv1.GroceryList); ok {
		r0 = rf(ctx, r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*groceriesv1.GroceryList)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *groceriesv1.GroceryList) error); ok {
		r1 = rf(ctx, r)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store_CreateGroceryList_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateGroceryList'
type Store_CreateGroceryList_Call struct {
	*mock.Call
}

// CreateGroceryList is a helper method to define mock.On call
//   - ctx context.Context
//   - r *groceriesv1.GroceryList
func (_e *Store_Expecter) CreateGroceryList(ctx interface{}, r interface{}) *Store_CreateGroceryList_Call {
	return &Store_CreateGroceryList_Call{Call: _e.mock.On("CreateGroceryList", ctx, r)}
}

func (_c *Store_CreateGroceryList_Call) Run(run func(ctx context.Context, r *groceriesv1.GroceryList)) *Store_CreateGroceryList_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*groceriesv1.GroceryList))
	})
	return _c
}

func (_c *Store_CreateGroceryList_Call) Return(_a0 *groceriesv1.GroceryList, _a1 error) *Store_CreateGroceryList_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Store_CreateGroceryList_Call) RunAndReturn(run func(context.Context, *groceriesv1.GroceryList) (*groceriesv1.GroceryList, error)) *Store_CreateGroceryList_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteGroceryList provides a mock function with given fields: ctx, id
func (_m *Store) DeleteGroceryList(ctx context.Context, id int64) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteGroceryList")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Store_DeleteGroceryList_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteGroceryList'
type Store_DeleteGroceryList_Call struct {
	*mock.Call
}

// DeleteGroceryList is a helper method to define mock.On call
//   - ctx context.Context
//   - id int64
func (_e *Store_Expecter) DeleteGroceryList(ctx interface{}, id interface{}) *Store_DeleteGroceryList_Call {
	return &Store_DeleteGroceryList_Call{Call: _e.mock.On("DeleteGroceryList", ctx, id)}
}

func (_c *Store_DeleteGroceryList_Call) Run(run func(ctx context.Context, id int64)) *Store_DeleteGroceryList_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *Store_DeleteGroceryList_Call) Return(_a0 error) *Store_DeleteGroceryList_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Store_DeleteGroceryList_Call) RunAndReturn(run func(context.Context, int64) error) *Store_DeleteGroceryList_Call {
	_c.Call.Return(run)
	return _c
}

// GroceryList provides a mock function with given fields: ctx, id
func (_m *Store) GroceryList(ctx context.Context, id int64) (*groceriesv1.GroceryList, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GroceryList")
	}

	var r0 *groceriesv1.GroceryList
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (*groceriesv1.GroceryList, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) *groceriesv1.GroceryList); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*groceriesv1.GroceryList)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store_GroceryList_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GroceryList'
type Store_GroceryList_Call struct {
	*mock.Call
}

// GroceryList is a helper method to define mock.On call
//   - ctx context.Context
//   - id int64
func (_e *Store_Expecter) GroceryList(ctx interface{}, id interface{}) *Store_GroceryList_Call {
	return &Store_GroceryList_Call{Call: _e.mock.On("GroceryList", ctx, id)}
}

func (_c *Store_GroceryList_Call) Run(run func(ctx context.Context, id int64)) *Store_GroceryList_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *Store_GroceryList_Call) Return(_a0 *groceriesv1.GroceryList, _a1 error) *Store_GroceryList_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Store_GroceryList_Call) RunAndReturn(run func(context.Context, int64) (*groceriesv1.GroceryList, error)) *Store_GroceryList_Call {
	_c.Call.Return(run)
	return _c
}

// GroceryLists provides a mock function with given fields: ctx
func (_m *Store) GroceryLists(ctx context.Context) ([]*groceriesv1.GroceryList, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GroceryLists")
	}

	var r0 []*groceriesv1.GroceryList
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]*groceriesv1.GroceryList, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []*groceriesv1.GroceryList); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*groceriesv1.GroceryList)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store_GroceryLists_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GroceryLists'
type Store_GroceryLists_Call struct {
	*mock.Call
}

// GroceryLists is a helper method to define mock.On call
//   - ctx context.Context
func (_e *Store_Expecter) GroceryLists(ctx interface{}) *Store_GroceryLists_Call {
	return &Store_GroceryLists_Call{Call: _e.mock.On("GroceryLists", ctx)}
}

func (_c *Store_GroceryLists_Call) Run(run func(ctx context.Context)) *Store_GroceryLists_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Store_GroceryLists_Call) Return(_a0 []*groceriesv1.GroceryList, _a1 error) *Store_GroceryLists_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Store_GroceryLists_Call) RunAndReturn(run func(context.Context) ([]*groceriesv1.GroceryList, error)) *Store_GroceryLists_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateGroceryList provides a mock function with given fields: ctx, r
func (_m *Store) UpdateGroceryList(ctx context.Context, r *groceriesv1.GroceryList) (*groceriesv1.GroceryList, error) {
	ret := _m.Called(ctx, r)

	if len(ret) == 0 {
		panic("no return value specified for UpdateGroceryList")
	}

	var r0 *groceriesv1.GroceryList
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *groceriesv1.GroceryList) (*groceriesv1.GroceryList, error)); ok {
		return rf(ctx, r)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *groceriesv1.GroceryList) *groceriesv1.GroceryList); ok {
		r0 = rf(ctx, r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*groceriesv1.GroceryList)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *groceriesv1.GroceryList) error); ok {
		r1 = rf(ctx, r)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store_UpdateGroceryList_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateGroceryList'
type Store_UpdateGroceryList_Call struct {
	*mock.Call
}

// UpdateGroceryList is a helper method to define mock.On call
//   - ctx context.Context
//   - r *groceriesv1.GroceryList
func (_e *Store_Expecter) UpdateGroceryList(ctx interface{}, r interface{}) *Store_UpdateGroceryList_Call {
	return &Store_UpdateGroceryList_Call{Call: _e.mock.On("UpdateGroceryList", ctx, r)}
}

func (_c *Store_UpdateGroceryList_Call) Run(run func(ctx context.Context, r *groceriesv1.GroceryList)) *Store_UpdateGroceryList_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*groceriesv1.GroceryList))
	})
	return _c
}

func (_c *Store_UpdateGroceryList_Call) Return(_a0 *groceriesv1.GroceryList, _a1 error) *Store_UpdateGroceryList_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Store_UpdateGroceryList_Call) RunAndReturn(run func(context.Context, *groceriesv1.GroceryList) (*groceriesv1.GroceryList, error)) *Store_UpdateGroceryList_Call {
	_c.Call.Return(run)
	return _c
}

// NewStore creates a new instance of Store. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewStore(t interface {
	mock.TestingT
	Cleanup(func())
}) *Store {
	mock := &Store{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
