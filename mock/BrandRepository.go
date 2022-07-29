// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	context "context"
	domain "jamtangan/domain"

	mock "github.com/stretchr/testify/mock"
)

// BrandRepository is an autogenerated mock type for the BrandRepository type
type BrandRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: _a0, _a1
func (_m *BrandRepository) Create(_a0 context.Context, _a1 *domain.Brand) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Brand) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *BrandRepository) GetByID(ctx context.Context, id int64) (domain.Brand, error) {
	ret := _m.Called(ctx, id)

	var r0 domain.Brand
	if rf, ok := ret.Get(0).(func(context.Context, int64) domain.Brand); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(domain.Brand)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
