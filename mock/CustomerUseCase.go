// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	context "context"
	domain "jamtangan/domain"

	mock "github.com/stretchr/testify/mock"
)

// CustomerUseCase is an autogenerated mock type for the CustomerUseCase type
type CustomerUseCase struct {
	mock.Mock
}

// CreateTransaction provides a mock function with given fields: ctx, request
func (_m *CustomerUseCase) CreateTransaction(ctx context.Context, request *domain.TransactionDetail) error {
	ret := _m.Called(ctx, request)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.TransactionDetail) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FetchProductByBrandID provides a mock function with given fields: ctx, brandID
func (_m *CustomerUseCase) FetchProductByBrandID(ctx context.Context, brandID int64) ([]domain.Product, error) {
	ret := _m.Called(ctx, brandID)

	var r0 []domain.Product
	if rf, ok := ret.Get(0).(func(context.Context, int64) []domain.Product); ok {
		r0 = rf(ctx, brandID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, brandID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTransactionByID provides a mock function with given fields: ctx, id
func (_m *CustomerUseCase) GetTransactionByID(ctx context.Context, id int64) (domain.TransactionDetail, error) {
	ret := _m.Called(ctx, id)

	var r0 domain.TransactionDetail
	if rf, ok := ret.Get(0).(func(context.Context, int64) domain.TransactionDetail); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(domain.TransactionDetail)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
