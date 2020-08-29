// Code generated by mockery v1.1.2. DO NOT EDIT.

package mocks

import (
	context "context"
	database "go-cicd/app/database"

	mock "github.com/stretchr/testify/mock"

	options "go.mongodb.org/mongo-driver/mongo/options"
)

// SessionContext is an autogenerated mock type for the SessionContext type
type SessionContext struct {
	mock.Mock
}

// AbortTransaction provides a mock function with given fields: ctx
func (_m *SessionContext) AbortTransaction(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Client provides a mock function with given fields:
func (_m *SessionContext) Client() database.Client {
	ret := _m.Called()

	var r0 database.Client
	if rf, ok := ret.Get(0).(func() database.Client); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(database.Client)
		}
	}

	return r0
}

// CommitTransaction provides a mock function with given fields: ctx
func (_m *SessionContext) CommitTransaction(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Context provides a mock function with given fields:
func (_m *SessionContext) Context() context.Context {
	ret := _m.Called()

	var r0 context.Context
	if rf, ok := ret.Get(0).(func() context.Context); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(context.Context)
		}
	}

	return r0
}

// StartTransaction provides a mock function with given fields: opts
func (_m *SessionContext) StartTransaction(opts ...*options.TransactionOptions) error {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(...*options.TransactionOptions) error); ok {
		r0 = rf(opts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
