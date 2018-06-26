// Code generated by mockery v1.0.0
package cache

import mock "github.com/stretchr/testify/mock"

// MockRequest is an autogenerated mock type for the Request type
type MockRequest struct {
	mock.Mock
}

// CacheInfo provides a mock function with given fields:
func (_m *MockRequest) CacheInfo() RequestInfo {
	ret := _m.Called()

	var r0 RequestInfo
	if rf, ok := ret.Get(0).(func() RequestInfo); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(RequestInfo)
	}

	return r0
}
