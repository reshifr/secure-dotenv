// Code generated by mockery. DO NOT EDIT.

package crypt_mock

import mock "github.com/stretchr/testify/mock"

// KDF is an autogenerated mock type for the IKDF type
type KDF struct {
	mock.Mock
}

type KDF_Expecter struct {
	mock *mock.Mock
}

func (_m *KDF) EXPECT() *KDF_Expecter {
	return &KDF_Expecter{mock: &_m.Mock}
}

// PassphraseKey provides a mock function with given fields: passphrase, salt, keyLen
func (_m *KDF) PassphraseKey(passphrase string, salt []byte, keyLen uint32) []byte {
	ret := _m.Called(passphrase, salt, keyLen)

	if len(ret) == 0 {
		panic("no return value specified for PassphraseKey")
	}

	var r0 []byte
	if rf, ok := ret.Get(0).(func(string, []byte, uint32) []byte); ok {
		r0 = rf(passphrase, salt, keyLen)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	return r0
}

// KDF_PassphraseKey_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PassphraseKey'
type KDF_PassphraseKey_Call struct {
	*mock.Call
}

// PassphraseKey is a helper method to define mock.On call
//   - passphrase string
//   - salt []byte
//   - keyLen uint32
func (_e *KDF_Expecter) PassphraseKey(passphrase interface{}, salt interface{}, keyLen interface{}) *KDF_PassphraseKey_Call {
	return &KDF_PassphraseKey_Call{Call: _e.mock.On("PassphraseKey", passphrase, salt, keyLen)}
}

func (_c *KDF_PassphraseKey_Call) Run(run func(passphrase string, salt []byte, keyLen uint32)) *KDF_PassphraseKey_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].([]byte), args[2].(uint32))
	})
	return _c
}

func (_c *KDF_PassphraseKey_Call) Return(key []byte) *KDF_PassphraseKey_Call {
	_c.Call.Return(key)
	return _c
}

func (_c *KDF_PassphraseKey_Call) RunAndReturn(run func(string, []byte, uint32) []byte) *KDF_PassphraseKey_Call {
	_c.Call.Return(run)
	return _c
}

// NewKDF creates a new instance of KDF. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewKDF(t interface {
	mock.TestingT
	Cleanup(func())
}) *KDF {
	mock := &KDF{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}