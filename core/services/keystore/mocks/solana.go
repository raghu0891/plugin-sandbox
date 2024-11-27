// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	solkey "github.com/goplugin/pluginv3.0/v2/core/services/keystore/keys/solkey"
)

// Solana is an autogenerated mock type for the Solana type
type Solana struct {
	mock.Mock
}

type Solana_Expecter struct {
	mock *mock.Mock
}

func (_m *Solana) EXPECT() *Solana_Expecter {
	return &Solana_Expecter{mock: &_m.Mock}
}

// Add provides a mock function with given fields: ctx, key
func (_m *Solana) Add(ctx context.Context, key solkey.Key) error {
	ret := _m.Called(ctx, key)

	if len(ret) == 0 {
		panic("no return value specified for Add")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, solkey.Key) error); ok {
		r0 = rf(ctx, key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Solana_Add_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Add'
type Solana_Add_Call struct {
	*mock.Call
}

// Add is a helper method to define mock.On call
//   - ctx context.Context
//   - key solkey.Key
func (_e *Solana_Expecter) Add(ctx interface{}, key interface{}) *Solana_Add_Call {
	return &Solana_Add_Call{Call: _e.mock.On("Add", ctx, key)}
}

func (_c *Solana_Add_Call) Run(run func(ctx context.Context, key solkey.Key)) *Solana_Add_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(solkey.Key))
	})
	return _c
}

func (_c *Solana_Add_Call) Return(_a0 error) *Solana_Add_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Solana_Add_Call) RunAndReturn(run func(context.Context, solkey.Key) error) *Solana_Add_Call {
	_c.Call.Return(run)
	return _c
}

// Create provides a mock function with given fields: ctx
func (_m *Solana) Create(ctx context.Context) (solkey.Key, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 solkey.Key
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (solkey.Key, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) solkey.Key); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(solkey.Key)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Solana_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type Solana_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
func (_e *Solana_Expecter) Create(ctx interface{}) *Solana_Create_Call {
	return &Solana_Create_Call{Call: _e.mock.On("Create", ctx)}
}

func (_c *Solana_Create_Call) Run(run func(ctx context.Context)) *Solana_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Solana_Create_Call) Return(_a0 solkey.Key, _a1 error) *Solana_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Solana_Create_Call) RunAndReturn(run func(context.Context) (solkey.Key, error)) *Solana_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, id
func (_m *Solana) Delete(ctx context.Context, id string) (solkey.Key, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 solkey.Key
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (solkey.Key, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) solkey.Key); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(solkey.Key)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Solana_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type Solana_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *Solana_Expecter) Delete(ctx interface{}, id interface{}) *Solana_Delete_Call {
	return &Solana_Delete_Call{Call: _e.mock.On("Delete", ctx, id)}
}

func (_c *Solana_Delete_Call) Run(run func(ctx context.Context, id string)) *Solana_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *Solana_Delete_Call) Return(_a0 solkey.Key, _a1 error) *Solana_Delete_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Solana_Delete_Call) RunAndReturn(run func(context.Context, string) (solkey.Key, error)) *Solana_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// EnsureKey provides a mock function with given fields: ctx
func (_m *Solana) EnsureKey(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for EnsureKey")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Solana_EnsureKey_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'EnsureKey'
type Solana_EnsureKey_Call struct {
	*mock.Call
}

// EnsureKey is a helper method to define mock.On call
//   - ctx context.Context
func (_e *Solana_Expecter) EnsureKey(ctx interface{}) *Solana_EnsureKey_Call {
	return &Solana_EnsureKey_Call{Call: _e.mock.On("EnsureKey", ctx)}
}

func (_c *Solana_EnsureKey_Call) Run(run func(ctx context.Context)) *Solana_EnsureKey_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Solana_EnsureKey_Call) Return(_a0 error) *Solana_EnsureKey_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Solana_EnsureKey_Call) RunAndReturn(run func(context.Context) error) *Solana_EnsureKey_Call {
	_c.Call.Return(run)
	return _c
}

// Export provides a mock function with given fields: id, password
func (_m *Solana) Export(id string, password string) ([]byte, error) {
	ret := _m.Called(id, password)

	if len(ret) == 0 {
		panic("no return value specified for Export")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) ([]byte, error)); ok {
		return rf(id, password)
	}
	if rf, ok := ret.Get(0).(func(string, string) []byte); ok {
		r0 = rf(id, password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(id, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Solana_Export_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Export'
type Solana_Export_Call struct {
	*mock.Call
}

// Export is a helper method to define mock.On call
//   - id string
//   - password string
func (_e *Solana_Expecter) Export(id interface{}, password interface{}) *Solana_Export_Call {
	return &Solana_Export_Call{Call: _e.mock.On("Export", id, password)}
}

func (_c *Solana_Export_Call) Run(run func(id string, password string)) *Solana_Export_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *Solana_Export_Call) Return(_a0 []byte, _a1 error) *Solana_Export_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Solana_Export_Call) RunAndReturn(run func(string, string) ([]byte, error)) *Solana_Export_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: id
func (_m *Solana) Get(id string) (solkey.Key, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 solkey.Key
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (solkey.Key, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) solkey.Key); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(solkey.Key)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Solana_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type Solana_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - id string
func (_e *Solana_Expecter) Get(id interface{}) *Solana_Get_Call {
	return &Solana_Get_Call{Call: _e.mock.On("Get", id)}
}

func (_c *Solana_Get_Call) Run(run func(id string)) *Solana_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Solana_Get_Call) Return(_a0 solkey.Key, _a1 error) *Solana_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Solana_Get_Call) RunAndReturn(run func(string) (solkey.Key, error)) *Solana_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetAll provides a mock function with given fields:
func (_m *Solana) GetAll() ([]solkey.Key, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAll")
	}

	var r0 []solkey.Key
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]solkey.Key, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []solkey.Key); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]solkey.Key)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Solana_GetAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAll'
type Solana_GetAll_Call struct {
	*mock.Call
}

// GetAll is a helper method to define mock.On call
func (_e *Solana_Expecter) GetAll() *Solana_GetAll_Call {
	return &Solana_GetAll_Call{Call: _e.mock.On("GetAll")}
}

func (_c *Solana_GetAll_Call) Run(run func()) *Solana_GetAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Solana_GetAll_Call) Return(_a0 []solkey.Key, _a1 error) *Solana_GetAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Solana_GetAll_Call) RunAndReturn(run func() ([]solkey.Key, error)) *Solana_GetAll_Call {
	_c.Call.Return(run)
	return _c
}

// Import provides a mock function with given fields: ctx, keyJSON, password
func (_m *Solana) Import(ctx context.Context, keyJSON []byte, password string) (solkey.Key, error) {
	ret := _m.Called(ctx, keyJSON, password)

	if len(ret) == 0 {
		panic("no return value specified for Import")
	}

	var r0 solkey.Key
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []byte, string) (solkey.Key, error)); ok {
		return rf(ctx, keyJSON, password)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []byte, string) solkey.Key); ok {
		r0 = rf(ctx, keyJSON, password)
	} else {
		r0 = ret.Get(0).(solkey.Key)
	}

	if rf, ok := ret.Get(1).(func(context.Context, []byte, string) error); ok {
		r1 = rf(ctx, keyJSON, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Solana_Import_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Import'
type Solana_Import_Call struct {
	*mock.Call
}

// Import is a helper method to define mock.On call
//   - ctx context.Context
//   - keyJSON []byte
//   - password string
func (_e *Solana_Expecter) Import(ctx interface{}, keyJSON interface{}, password interface{}) *Solana_Import_Call {
	return &Solana_Import_Call{Call: _e.mock.On("Import", ctx, keyJSON, password)}
}

func (_c *Solana_Import_Call) Run(run func(ctx context.Context, keyJSON []byte, password string)) *Solana_Import_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]byte), args[2].(string))
	})
	return _c
}

func (_c *Solana_Import_Call) Return(_a0 solkey.Key, _a1 error) *Solana_Import_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Solana_Import_Call) RunAndReturn(run func(context.Context, []byte, string) (solkey.Key, error)) *Solana_Import_Call {
	_c.Call.Return(run)
	return _c
}

// Sign provides a mock function with given fields: ctx, id, msg
func (_m *Solana) Sign(ctx context.Context, id string, msg []byte) ([]byte, error) {
	ret := _m.Called(ctx, id, msg)

	if len(ret) == 0 {
		panic("no return value specified for Sign")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, []byte) ([]byte, error)); ok {
		return rf(ctx, id, msg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, []byte) []byte); ok {
		r0 = rf(ctx, id, msg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, []byte) error); ok {
		r1 = rf(ctx, id, msg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Solana_Sign_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Sign'
type Solana_Sign_Call struct {
	*mock.Call
}

// Sign is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
//   - msg []byte
func (_e *Solana_Expecter) Sign(ctx interface{}, id interface{}, msg interface{}) *Solana_Sign_Call {
	return &Solana_Sign_Call{Call: _e.mock.On("Sign", ctx, id, msg)}
}

func (_c *Solana_Sign_Call) Run(run func(ctx context.Context, id string, msg []byte)) *Solana_Sign_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].([]byte))
	})
	return _c
}

func (_c *Solana_Sign_Call) Return(signature []byte, err error) *Solana_Sign_Call {
	_c.Call.Return(signature, err)
	return _c
}

func (_c *Solana_Sign_Call) RunAndReturn(run func(context.Context, string, []byte) ([]byte, error)) *Solana_Sign_Call {
	_c.Call.Return(run)
	return _c
}

// NewSolana creates a new instance of Solana. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSolana(t interface {
	mock.TestingT
	Cleanup(func())
}) *Solana {
	mock := &Solana{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
