// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import conc "github.com/applike/gosoline/pkg/conc"
import context "context"
import mock "github.com/stretchr/testify/mock"

// DistributedLockProvider is an autogenerated mock type for the DistributedLockProvider type
type DistributedLockProvider struct {
	mock.Mock
}

// Acquire provides a mock function with given fields: ctx, resource
func (_m *DistributedLockProvider) Acquire(ctx context.Context, resource string) (conc.DistributedLock, error) {
	ret := _m.Called(ctx, resource)

	var r0 conc.DistributedLock
	if rf, ok := ret.Get(0).(func(context.Context, string) conc.DistributedLock); ok {
		r0 = rf(ctx, resource)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(conc.DistributedLock)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, resource)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
