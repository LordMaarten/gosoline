// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"

import mock "github.com/stretchr/testify/mock"

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// HasCurrency provides a mock function with given fields: _a0, _a1
func (_m *Service) HasCurrency(_a0 context.Context, _a1 string) (bool, error) {
	ret := _m.Called(_a0, _a1)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ToCurrency provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *Service) ToCurrency(_a0 context.Context, _a1 string, _a2 float64, _a3 string) (float64, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 float64
	if rf, ok := ret.Get(0).(func(context.Context, string, float64, string) float64); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		r0 = ret.Get(0).(float64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, float64, string) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ToEur provides a mock function with given fields: _a0, _a1, _a2
func (_m *Service) ToEur(_a0 context.Context, _a1 float64, _a2 string) (float64, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 float64
	if rf, ok := ret.Get(0).(func(context.Context, float64, string) float64); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).(float64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, float64, string) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ToUsd provides a mock function with given fields: _a0, _a1, _a2
func (_m *Service) ToUsd(_a0 context.Context, _a1 float64, _a2 string) (float64, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 float64
	if rf, ok := ret.Get(0).(func(context.Context, float64, string) float64); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).(float64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, float64, string) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
