// Code generated by mockery. DO NOT EDIT.

package crypt_mock

import mock "github.com/stretchr/testify/mock"

// VI is an autogenerated mock type for the IVI type
type VI struct {
	mock.Mock
}

type VI_Expecter struct {
	mock *mock.Mock
}

func (_m *VI) EXPECT() *VI_Expecter {
	return &VI_Expecter{mock: &_m.Mock}
}

// Invoke provides a mock function with given fields:
func (_m *VI) Invoke() {
	_m.Called()
}

// VI_Invoke_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Invoke'
type VI_Invoke_Call struct {
	*mock.Call
}

// Invoke is a helper method to define mock.On call
func (_e *VI_Expecter) Invoke() *VI_Invoke_Call {
	return &VI_Invoke_Call{Call: _e.mock.On("Invoke")}
}

func (_c *VI_Invoke_Call) Run(run func()) *VI_Invoke_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *VI_Invoke_Call) Return() *VI_Invoke_Call {
	_c.Call.Return()
	return _c
}

func (_c *VI_Invoke_Call) RunAndReturn(run func()) *VI_Invoke_Call {
	_c.Call.Return(run)
	return _c
}

// Raw provides a mock function with given fields:
func (_m *VI) Raw() []byte {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Raw")
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

// VI_Raw_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Raw'
type VI_Raw_Call struct {
	*mock.Call
}

// Raw is a helper method to define mock.On call
func (_e *VI_Expecter) Raw() *VI_Raw_Call {
	return &VI_Raw_Call{Call: _e.mock.On("Raw")}
}

func (_c *VI_Raw_Call) Run(run func()) *VI_Raw_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *VI_Raw_Call) Return(rawIV []byte) *VI_Raw_Call {
	_c.Call.Return(rawIV)
	return _c
}

func (_c *VI_Raw_Call) RunAndReturn(run func() []byte) *VI_Raw_Call {
	_c.Call.Return(run)
	return _c
}

// NewVI creates a new instance of VI. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewVI(t interface {
	mock.TestingT
	Cleanup(func())
}) *VI {
	mock := &VI{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
