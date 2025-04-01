// Code generated by mockery v2.53.3. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// URLSaver is an autogenerated mock type for the URLSaver type
type URLSaver struct {
	mock.Mock
}

type URLSaver_Expecter struct {
	mock *mock.Mock
}

func (_m *URLSaver) EXPECT() *URLSaver_Expecter {
	return &URLSaver_Expecter{mock: &_m.Mock}
}

// SaveURL provides a mock function with given fields: urlToSave, alias
func (_m *URLSaver) SaveURL(urlToSave string, alias string) (int64, error) {
	ret := _m.Called(urlToSave, alias)

	if len(ret) == 0 {
		panic("no return value specified for SaveURL")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (int64, error)); ok {
		return rf(urlToSave, alias)
	}
	if rf, ok := ret.Get(0).(func(string, string) int64); ok {
		r0 = rf(urlToSave, alias)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(urlToSave, alias)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// URLSaver_SaveURL_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SaveURL'
type URLSaver_SaveURL_Call struct {
	*mock.Call
}

// SaveURL is a helper method to define mock.On call
//   - urlToSave string
//   - alias string
func (_e *URLSaver_Expecter) SaveURL(urlToSave interface{}, alias interface{}) *URLSaver_SaveURL_Call {
	return &URLSaver_SaveURL_Call{Call: _e.mock.On("SaveURL", urlToSave, alias)}
}

func (_c *URLSaver_SaveURL_Call) Run(run func(urlToSave string, alias string)) *URLSaver_SaveURL_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *URLSaver_SaveURL_Call) Return(_a0 int64, _a1 error) *URLSaver_SaveURL_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *URLSaver_SaveURL_Call) RunAndReturn(run func(string, string) (int64, error)) *URLSaver_SaveURL_Call {
	_c.Call.Return(run)
	return _c
}

// NewURLSaver creates a new instance of URLSaver. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewURLSaver(t interface {
	mock.TestingT
	Cleanup(func())
}) *URLSaver {
	mock := &URLSaver{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
