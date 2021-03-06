// Code generated by mockery v1.1.2. DO NOT EDIT.

package mocks

import (
	context "context"
	database "go-cicd/app/database"

	mock "github.com/stretchr/testify/mock"

	mongo "go.mongodb.org/mongo-driver/mongo"

	options "go.mongodb.org/mongo-driver/mongo/options"
)

// Collection is an autogenerated mock type for the Collection type
type Collection struct {
	mock.Mock
}

// CountDocuments provides a mock function with given fields: ctx, filter, opts
func (_m *Collection) CountDocuments(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (int64, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, filter)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, ...*options.CountOptions) int64); ok {
		r0 = rf(ctx, filter, opts...)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, interface{}, ...*options.CountOptions) error); ok {
		r1 = rf(ctx, filter, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteMany provides a mock function with given fields: ctx, filter, opts
func (_m *Collection) DeleteMany(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (database.DeleteResult, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, filter)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 database.DeleteResult
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, ...*options.DeleteOptions) database.DeleteResult); ok {
		r0 = rf(ctx, filter, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(database.DeleteResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, interface{}, ...*options.DeleteOptions) error); ok {
		r1 = rf(ctx, filter, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteOne provides a mock function with given fields: ctx, filter, opts
func (_m *Collection) DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (database.DeleteResult, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, filter)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 database.DeleteResult
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, ...*options.DeleteOptions) database.DeleteResult); ok {
		r0 = rf(ctx, filter, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(database.DeleteResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, interface{}, ...*options.DeleteOptions) error); ok {
		r1 = rf(ctx, filter, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Find provides a mock function with given fields: ctx, filter, opts
func (_m *Collection) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (database.Cursor, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, filter)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 database.Cursor
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, ...*options.FindOptions) database.Cursor); ok {
		r0 = rf(ctx, filter, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(database.Cursor)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, interface{}, ...*options.FindOptions) error); ok {
		r1 = rf(ctx, filter, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindOne provides a mock function with given fields: ctx, filter, opts
func (_m *Collection) FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) database.SingleResult {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, filter)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 database.SingleResult
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, ...*options.FindOneOptions) database.SingleResult); ok {
		r0 = rf(ctx, filter, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(database.SingleResult)
		}
	}

	return r0
}

// FindOneAndDelete provides a mock function with given fields: ctx, filter, opts
func (_m *Collection) FindOneAndDelete(ctx context.Context, filter interface{}, opts ...*options.FindOneAndDeleteOptions) database.SingleResult {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, filter)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 database.SingleResult
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, ...*options.FindOneAndDeleteOptions) database.SingleResult); ok {
		r0 = rf(ctx, filter, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(database.SingleResult)
		}
	}

	return r0
}

// FindOneAndReplace provides a mock function with given fields: ctx, filter, replacement, opts
func (_m *Collection) FindOneAndReplace(ctx context.Context, filter interface{}, replacement interface{}, opts ...*options.FindOneAndReplaceOptions) database.SingleResult {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, filter, replacement)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 database.SingleResult
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, interface{}, ...*options.FindOneAndReplaceOptions) database.SingleResult); ok {
		r0 = rf(ctx, filter, replacement, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(database.SingleResult)
		}
	}

	return r0
}

// FindOneAndUpdate provides a mock function with given fields: ctx, filter, update, opts
func (_m *Collection) FindOneAndUpdate(ctx context.Context, filter interface{}, update interface{}, opts ...*options.FindOneAndUpdateOptions) database.SingleResult {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, filter, update)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 database.SingleResult
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, interface{}, ...*options.FindOneAndUpdateOptions) database.SingleResult); ok {
		r0 = rf(ctx, filter, update, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(database.SingleResult)
		}
	}

	return r0
}

// Indexes provides a mock function with given fields:
func (_m *Collection) Indexes() mongo.IndexView {
	ret := _m.Called()

	var r0 mongo.IndexView
	if rf, ok := ret.Get(0).(func() mongo.IndexView); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(mongo.IndexView)
	}

	return r0
}

// InsertOne provides a mock function with given fields: ctx, document, opts
func (_m *Collection) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (database.InsertOneResult, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, document)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 database.InsertOneResult
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, ...*options.InsertOneOptions) database.InsertOneResult); ok {
		r0 = rf(ctx, document, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(database.InsertOneResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, interface{}, ...*options.InsertOneOptions) error); ok {
		r1 = rf(ctx, document, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReplaceOne provides a mock function with given fields: ctx, filter, replacement, opts
func (_m *Collection) ReplaceOne(ctx context.Context, filter interface{}, replacement interface{}, opts ...*options.ReplaceOptions) (database.UpdateResult, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, filter, replacement)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 database.UpdateResult
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, interface{}, ...*options.ReplaceOptions) database.UpdateResult); ok {
		r0 = rf(ctx, filter, replacement, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(database.UpdateResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, interface{}, interface{}, ...*options.ReplaceOptions) error); ok {
		r1 = rf(ctx, filter, replacement, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateMany provides a mock function with given fields: ctx, filter, update, opts
func (_m *Collection) UpdateMany(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (database.UpdateResult, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, filter, update)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 database.UpdateResult
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, interface{}, ...*options.UpdateOptions) database.UpdateResult); ok {
		r0 = rf(ctx, filter, update, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(database.UpdateResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, interface{}, interface{}, ...*options.UpdateOptions) error); ok {
		r1 = rf(ctx, filter, update, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateOne provides a mock function with given fields: ctx, filter, update, opts
func (_m *Collection) UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (database.UpdateResult, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, filter, update)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 database.UpdateResult
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, interface{}, ...*options.UpdateOptions) database.UpdateResult); ok {
		r0 = rf(ctx, filter, update, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(database.UpdateResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, interface{}, interface{}, ...*options.UpdateOptions) error); ok {
		r1 = rf(ctx, filter, update, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
