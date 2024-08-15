// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	domain "cleantaskmanager/domain"

	mock "github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UserUsecase is an autogenerated mock type for the UserUsecase type
type UserUsecase struct {
	mock.Mock
}

// GetUserByID provides a mock function with given fields: id
func (_m *UserUsecase) GetUserByID(id primitive.ObjectID) (*domain.User, error) {
	ret := _m.Called(id)

	var r0 *domain.User
	if rf, ok := ret.Get(0).(func(primitive.ObjectID) *domain.User); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(primitive.ObjectID) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LoginUser provides a mock function with given fields: user
func (_m *UserUsecase) LoginUser(user *domain.Login) (string, error) {
	ret := _m.Called(user)

	var r0 string
	if rf, ok := ret.Get(0).(func(*domain.Login) string); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*domain.Login) error); ok {
		r1 = rf(user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterUser provides a mock function with given fields: user
func (_m *UserUsecase) RegisterUser(user *domain.User) (primitive.ObjectID, error) {
	ret := _m.Called(user)

	var r0 primitive.ObjectID
	if rf, ok := ret.Get(0).(func(*domain.User) (primitive.ObjectID, error)); ok {
		r0, _ = rf(user)
	} else {
		r0 = ret.Get(0).(primitive.ObjectID)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*domain.User) error); ok {
		r1 = rf(user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUserUsecase interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserUsecase creates a new instance of UserUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserUsecase(t mockConstructorTestingTNewUserUsecase) *UserUsecase {
	mock := &UserUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
