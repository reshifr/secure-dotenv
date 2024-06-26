// Code generated by mockery. DO NOT EDIT.

package crypto_mock

import mock "github.com/stretchr/testify/mock"

// IV is an autogenerated mock type for the IV type
type IV struct {
	mock.Mock
}

type IV_Expecter struct {
	mock *mock.Mock
}

func (_m *IV) EXPECT() *IV_Expecter {
	return &IV_Expecter{mock: &_m.Mock}
}

// Invoke provides a mock function with given fields:
func (_m *IV) Invoke() []byte {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Invoke")
	}

	var r0 []byte
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	return r0
}

// IV_Invoke_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Invoke'
type IV_Invoke_Call struct {
	*mock.Call
}

// Invoke is a helper method to define mock.On call
func (_e *IV_Expecter) Invoke() *IV_Invoke_Call {
	return &IV_Invoke_Call{Call: _e.mock.On("Invoke")}
}

func (_c *IV_Invoke_Call) Run(run func()) *IV_Invoke_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *IV_Invoke_Call) Return(invokedRawIV []byte) *IV_Invoke_Call {
	_c.Call.Return(invokedRawIV)
	return _c
}

func (_c *IV_Invoke_Call) RunAndReturn(run func() []byte) *IV_Invoke_Call {
	_c.Call.Return(run)
	return _c
}

// Len provides a mock function with given fields:
func (_m *IV) Len() uint32 {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Len")
	}

	var r0 uint32
	if rf, ok := ret.Get(0).(func() uint32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint32)
	}

	return r0
}

// IV_Len_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Len'
type IV_Len_Call struct {
	*mock.Call
}

// Len is a helper method to define mock.On call
func (_e *IV_Expecter) Len() *IV_Len_Call {
	return &IV_Len_Call{Call: _e.mock.On("Len")}
}

func (_c *IV_Len_Call) Run(run func()) *IV_Len_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *IV_Len_Call) Return(ivLen uint32) *IV_Len_Call {
	_c.Call.Return(ivLen)
	return _c
}

func (_c *IV_Len_Call) RunAndReturn(run func() uint32) *IV_Len_Call {
	_c.Call.Return(run)
	return _c
}

// NewIV creates a new instance of IV. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIV(t interface {
	mock.TestingT
	Cleanup(func())
}) *IV {
	mock := &IV{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
