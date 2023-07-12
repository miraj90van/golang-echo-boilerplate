// Code generated by mockery v2.23.0. DO NOT EDIT.

package mock_db

import (
	context "context"

	ent "myapp/ent"

	mock "github.com/stretchr/testify/mock"
)

// TemplateRepository is an autogenerated mock type for the TemplateRepository type
type TemplateRepository struct {
	mock.Mock
}

// DatabaseAction provides a mock function with given fields: ctx, requestDomain
func (_m *TemplateRepository) DatabaseAction(ctx context.Context, requestDomain *ent.Pet) (*ent.Pet, error) {
	ret := _m.Called(ctx, requestDomain)

	var r0 *ent.Pet
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ent.Pet) (*ent.Pet, error)); ok {
		return rf(ctx, requestDomain)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ent.Pet) *ent.Pet); ok {
		r0 = rf(ctx, requestDomain)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ent.Pet)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ent.Pet) error); ok {
		r1 = rf(ctx, requestDomain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewTemplateRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewTemplateRepository creates a new instance of TemplateRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTemplateRepository(t mockConstructorTestingTNewTemplateRepository) *TemplateRepository {
	mock := &TemplateRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
