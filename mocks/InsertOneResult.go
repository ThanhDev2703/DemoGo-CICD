// Code generated by mockery v1.1.2. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// InsertOneResult is an autogenerated mock type for the InsertOneResult type
type InsertOneResult struct {
	mock.Mock
}

// InsertedID provides a mock function with given fields:
func (_m *InsertOneResult) InsertedID() interface{} {
	ret := _m.Called()

	var r0 interface{}
	if rf, ok := ret.Get(0).(func() interface{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}