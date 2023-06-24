// Code generated by mockery v2.26.1. DO NOT EDIT.

package mocks

import (
	context "context"

	users "github.com/ergomake/ergomake/internal/users"
	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

type Service_Expecter struct {
	mock *mock.Mock
}

func (_m *Service) EXPECT() *Service_Expecter {
	return &Service_Expecter{mock: &_m.Mock}
}

// Save provides a mock function with given fields: ctx, user
func (_m *Service) Save(ctx context.Context, user users.User) error {
	ret := _m.Called(ctx, user)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, users.User) error); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Service_Save_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Save'
type Service_Save_Call struct {
	*mock.Call
}

// Save is a helper method to define mock.On call
//   - ctx context.Context
//   - user users.User
func (_e *Service_Expecter) Save(ctx interface{}, user interface{}) *Service_Save_Call {
	return &Service_Save_Call{Call: _e.mock.On("Save", ctx, user)}
}

func (_c *Service_Save_Call) Run(run func(ctx context.Context, user users.User)) *Service_Save_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(users.User))
	})
	return _c
}

func (_c *Service_Save_Call) Return(_a0 error) *Service_Save_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Service_Save_Call) RunAndReturn(run func(context.Context, users.User) error) *Service_Save_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewService interface {
	mock.TestingT
	Cleanup(func())
}

// NewService creates a new instance of Service. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewService(t mockConstructorTestingTNewService) *Service {
	mock := &Service{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
