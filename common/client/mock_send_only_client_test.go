// Code generated by mockery v2.43.2. DO NOT EDIT.

package client

import (
	context "context"

	types "github.com/goplugin/pluginv3.0/v2/common/types"
	mock "github.com/stretchr/testify/mock"
)

// mockSendOnlyClient is an autogenerated mock type for the sendOnlyClient type
type mockSendOnlyClient[CHAIN_ID types.ID] struct {
	mock.Mock
}

type mockSendOnlyClient_Expecter[CHAIN_ID types.ID] struct {
	mock *mock.Mock
}

func (_m *mockSendOnlyClient[CHAIN_ID]) EXPECT() *mockSendOnlyClient_Expecter[CHAIN_ID] {
	return &mockSendOnlyClient_Expecter[CHAIN_ID]{mock: &_m.Mock}
}

// ChainID provides a mock function with given fields: _a0
func (_m *mockSendOnlyClient[CHAIN_ID]) ChainID(_a0 context.Context) (CHAIN_ID, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for ChainID")
	}

	var r0 CHAIN_ID
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (CHAIN_ID, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(context.Context) CHAIN_ID); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(CHAIN_ID)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockSendOnlyClient_ChainID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ChainID'
type mockSendOnlyClient_ChainID_Call[CHAIN_ID types.ID] struct {
	*mock.Call
}

// ChainID is a helper method to define mock.On call
//   - _a0 context.Context
func (_e *mockSendOnlyClient_Expecter[CHAIN_ID]) ChainID(_a0 interface{}) *mockSendOnlyClient_ChainID_Call[CHAIN_ID] {
	return &mockSendOnlyClient_ChainID_Call[CHAIN_ID]{Call: _e.mock.On("ChainID", _a0)}
}

func (_c *mockSendOnlyClient_ChainID_Call[CHAIN_ID]) Run(run func(_a0 context.Context)) *mockSendOnlyClient_ChainID_Call[CHAIN_ID] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *mockSendOnlyClient_ChainID_Call[CHAIN_ID]) Return(_a0 CHAIN_ID, _a1 error) *mockSendOnlyClient_ChainID_Call[CHAIN_ID] {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *mockSendOnlyClient_ChainID_Call[CHAIN_ID]) RunAndReturn(run func(context.Context) (CHAIN_ID, error)) *mockSendOnlyClient_ChainID_Call[CHAIN_ID] {
	_c.Call.Return(run)
	return _c
}

// Close provides a mock function with given fields:
func (_m *mockSendOnlyClient[CHAIN_ID]) Close() {
	_m.Called()
}

// mockSendOnlyClient_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type mockSendOnlyClient_Close_Call[CHAIN_ID types.ID] struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *mockSendOnlyClient_Expecter[CHAIN_ID]) Close() *mockSendOnlyClient_Close_Call[CHAIN_ID] {
	return &mockSendOnlyClient_Close_Call[CHAIN_ID]{Call: _e.mock.On("Close")}
}

func (_c *mockSendOnlyClient_Close_Call[CHAIN_ID]) Run(run func()) *mockSendOnlyClient_Close_Call[CHAIN_ID] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *mockSendOnlyClient_Close_Call[CHAIN_ID]) Return() *mockSendOnlyClient_Close_Call[CHAIN_ID] {
	_c.Call.Return()
	return _c
}

func (_c *mockSendOnlyClient_Close_Call[CHAIN_ID]) RunAndReturn(run func()) *mockSendOnlyClient_Close_Call[CHAIN_ID] {
	_c.Call.Return(run)
	return _c
}

// Dial provides a mock function with given fields: ctx
func (_m *mockSendOnlyClient[CHAIN_ID]) Dial(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Dial")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// mockSendOnlyClient_Dial_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Dial'
type mockSendOnlyClient_Dial_Call[CHAIN_ID types.ID] struct {
	*mock.Call
}

// Dial is a helper method to define mock.On call
//   - ctx context.Context
func (_e *mockSendOnlyClient_Expecter[CHAIN_ID]) Dial(ctx interface{}) *mockSendOnlyClient_Dial_Call[CHAIN_ID] {
	return &mockSendOnlyClient_Dial_Call[CHAIN_ID]{Call: _e.mock.On("Dial", ctx)}
}

func (_c *mockSendOnlyClient_Dial_Call[CHAIN_ID]) Run(run func(ctx context.Context)) *mockSendOnlyClient_Dial_Call[CHAIN_ID] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *mockSendOnlyClient_Dial_Call[CHAIN_ID]) Return(_a0 error) *mockSendOnlyClient_Dial_Call[CHAIN_ID] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mockSendOnlyClient_Dial_Call[CHAIN_ID]) RunAndReturn(run func(context.Context) error) *mockSendOnlyClient_Dial_Call[CHAIN_ID] {
	_c.Call.Return(run)
	return _c
}

// newMockSendOnlyClient creates a new instance of mockSendOnlyClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockSendOnlyClient[CHAIN_ID types.ID](t interface {
	mock.TestingT
	Cleanup(func())
}) *mockSendOnlyClient[CHAIN_ID] {
	mock := &mockSendOnlyClient[CHAIN_ID]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
