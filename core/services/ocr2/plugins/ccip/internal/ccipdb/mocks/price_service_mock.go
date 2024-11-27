// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	big "math/big"

	ccip "github.com/goplugin/plugin-common/pkg/types/ccip"
	ccipdata "github.com/goplugin/pluginv3.0/v2/core/services/ocr2/plugins/ccip/internal/ccipdata"

	context "context"

	mock "github.com/stretchr/testify/mock"

	prices "github.com/goplugin/pluginv3.0/v2/core/services/ocr2/plugins/ccip/prices"
)

// PriceService is an autogenerated mock type for the PriceService type
type PriceService struct {
	mock.Mock
}

type PriceService_Expecter struct {
	mock *mock.Mock
}

func (_m *PriceService) EXPECT() *PriceService_Expecter {
	return &PriceService_Expecter{mock: &_m.Mock}
}

// Close provides a mock function with given fields:
func (_m *PriceService) Close() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Close")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PriceService_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type PriceService_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *PriceService_Expecter) Close() *PriceService_Close_Call {
	return &PriceService_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *PriceService_Close_Call) Run(run func()) *PriceService_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *PriceService_Close_Call) Return(_a0 error) *PriceService_Close_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PriceService_Close_Call) RunAndReturn(run func() error) *PriceService_Close_Call {
	_c.Call.Return(run)
	return _c
}

// GetGasAndTokenPrices provides a mock function with given fields: ctx, destChainSelector
func (_m *PriceService) GetGasAndTokenPrices(ctx context.Context, destChainSelector uint64) (map[uint64]*big.Int, map[ccip.Address]*big.Int, error) {
	ret := _m.Called(ctx, destChainSelector)

	if len(ret) == 0 {
		panic("no return value specified for GetGasAndTokenPrices")
	}

	var r0 map[uint64]*big.Int
	var r1 map[ccip.Address]*big.Int
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) (map[uint64]*big.Int, map[ccip.Address]*big.Int, error)); ok {
		return rf(ctx, destChainSelector)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64) map[uint64]*big.Int); ok {
		r0 = rf(ctx, destChainSelector)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[uint64]*big.Int)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64) map[ccip.Address]*big.Int); ok {
		r1 = rf(ctx, destChainSelector)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(map[ccip.Address]*big.Int)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, uint64) error); ok {
		r2 = rf(ctx, destChainSelector)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// PriceService_GetGasAndTokenPrices_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetGasAndTokenPrices'
type PriceService_GetGasAndTokenPrices_Call struct {
	*mock.Call
}

// GetGasAndTokenPrices is a helper method to define mock.On call
//   - ctx context.Context
//   - destChainSelector uint64
func (_e *PriceService_Expecter) GetGasAndTokenPrices(ctx interface{}, destChainSelector interface{}) *PriceService_GetGasAndTokenPrices_Call {
	return &PriceService_GetGasAndTokenPrices_Call{Call: _e.mock.On("GetGasAndTokenPrices", ctx, destChainSelector)}
}

func (_c *PriceService_GetGasAndTokenPrices_Call) Run(run func(ctx context.Context, destChainSelector uint64)) *PriceService_GetGasAndTokenPrices_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64))
	})
	return _c
}

func (_c *PriceService_GetGasAndTokenPrices_Call) Return(_a0 map[uint64]*big.Int, _a1 map[ccip.Address]*big.Int, _a2 error) *PriceService_GetGasAndTokenPrices_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *PriceService_GetGasAndTokenPrices_Call) RunAndReturn(run func(context.Context, uint64) (map[uint64]*big.Int, map[ccip.Address]*big.Int, error)) *PriceService_GetGasAndTokenPrices_Call {
	_c.Call.Return(run)
	return _c
}

// Start provides a mock function with given fields: _a0
func (_m *PriceService) Start(_a0 context.Context) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Start")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PriceService_Start_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Start'
type PriceService_Start_Call struct {
	*mock.Call
}

// Start is a helper method to define mock.On call
//   - _a0 context.Context
func (_e *PriceService_Expecter) Start(_a0 interface{}) *PriceService_Start_Call {
	return &PriceService_Start_Call{Call: _e.mock.On("Start", _a0)}
}

func (_c *PriceService_Start_Call) Run(run func(_a0 context.Context)) *PriceService_Start_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *PriceService_Start_Call) Return(_a0 error) *PriceService_Start_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PriceService_Start_Call) RunAndReturn(run func(context.Context) error) *PriceService_Start_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateDynamicConfig provides a mock function with given fields: ctx, gasPriceEstimator, destPriceRegistryReader
func (_m *PriceService) UpdateDynamicConfig(ctx context.Context, gasPriceEstimator prices.GasPriceEstimatorCommit, destPriceRegistryReader ccipdata.PriceRegistryReader) error {
	ret := _m.Called(ctx, gasPriceEstimator, destPriceRegistryReader)

	if len(ret) == 0 {
		panic("no return value specified for UpdateDynamicConfig")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, prices.GasPriceEstimatorCommit, ccipdata.PriceRegistryReader) error); ok {
		r0 = rf(ctx, gasPriceEstimator, destPriceRegistryReader)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PriceService_UpdateDynamicConfig_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateDynamicConfig'
type PriceService_UpdateDynamicConfig_Call struct {
	*mock.Call
}

// UpdateDynamicConfig is a helper method to define mock.On call
//   - ctx context.Context
//   - gasPriceEstimator prices.GasPriceEstimatorCommit
//   - destPriceRegistryReader ccipdata.PriceRegistryReader
func (_e *PriceService_Expecter) UpdateDynamicConfig(ctx interface{}, gasPriceEstimator interface{}, destPriceRegistryReader interface{}) *PriceService_UpdateDynamicConfig_Call {
	return &PriceService_UpdateDynamicConfig_Call{Call: _e.mock.On("UpdateDynamicConfig", ctx, gasPriceEstimator, destPriceRegistryReader)}
}

func (_c *PriceService_UpdateDynamicConfig_Call) Run(run func(ctx context.Context, gasPriceEstimator prices.GasPriceEstimatorCommit, destPriceRegistryReader ccipdata.PriceRegistryReader)) *PriceService_UpdateDynamicConfig_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(prices.GasPriceEstimatorCommit), args[2].(ccipdata.PriceRegistryReader))
	})
	return _c
}

func (_c *PriceService_UpdateDynamicConfig_Call) Return(_a0 error) *PriceService_UpdateDynamicConfig_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PriceService_UpdateDynamicConfig_Call) RunAndReturn(run func(context.Context, prices.GasPriceEstimatorCommit, ccipdata.PriceRegistryReader) error) *PriceService_UpdateDynamicConfig_Call {
	_c.Call.Return(run)
	return _c
}

// NewPriceService creates a new instance of PriceService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPriceService(t interface {
	mock.TestingT
	Cleanup(func())
}) *PriceService {
	mock := &PriceService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
