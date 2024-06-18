// Code generated by MockGen. DO NOT EDIT.
// Source: /home/vinicius/Documents/go-projects/go-web-app/src/model/repository/user_repository.go
//
// Generated by this command:
//
//	mockgen -source=/home/vinicius/Documents/go-projects/go-web-app/src/model/repository/user_repository.go -destination=/home/vinicius/Documents/go-projects/go-web-app/src/tests/mocks/user_repository_mock.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	rest_err "github.com/Vinicius-Madeira/go-web-app/src/configuration/rest_err"
	model "github.com/Vinicius-Madeira/go-web-app/src/model"
	gomock "go.uber.org/mock/gomock"
)

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockUserRepository) CreateUser(UserDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", UserDomain)
	ret0, _ := ret[0].(model.UserDomainInterface)
	ret1, _ := ret[1].(*rest_err.RestError)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserRepositoryMockRecorder) CreateUser(UserDomain any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserRepository)(nil).CreateUser), UserDomain)
}

// DeleteUser mocks base method.
func (m *MockUserRepository) DeleteUser(userId string) *rest_err.RestError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", userId)
	ret0, _ := ret[0].(*rest_err.RestError)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockUserRepositoryMockRecorder) DeleteUser(userId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockUserRepository)(nil).DeleteUser), userId)
}

// FindUserByEmail mocks base method.
func (m *MockUserRepository) FindUserByEmail(email string) (model.UserDomainInterface, *rest_err.RestError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByEmail", email)
	ret0, _ := ret[0].(model.UserDomainInterface)
	ret1, _ := ret[1].(*rest_err.RestError)
	return ret0, ret1
}

// FindUserByEmail indicates an expected call of FindUserByEmail.
func (mr *MockUserRepositoryMockRecorder) FindUserByEmail(email any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByEmail", reflect.TypeOf((*MockUserRepository)(nil).FindUserByEmail), email)
}

// FindUserByEmailAndPassword mocks base method.
func (m *MockUserRepository) FindUserByEmailAndPassword(email, password string) (model.UserDomainInterface, *rest_err.RestError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByEmailAndPassword", email, password)
	ret0, _ := ret[0].(model.UserDomainInterface)
	ret1, _ := ret[1].(*rest_err.RestError)
	return ret0, ret1
}

// FindUserByEmailAndPassword indicates an expected call of FindUserByEmailAndPassword.
func (mr *MockUserRepositoryMockRecorder) FindUserByEmailAndPassword(email, password any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByEmailAndPassword", reflect.TypeOf((*MockUserRepository)(nil).FindUserByEmailAndPassword), email, password)
}

// FindUserByID mocks base method.
func (m *MockUserRepository) FindUserByID(id string) (model.UserDomainInterface, *rest_err.RestError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByID", id)
	ret0, _ := ret[0].(model.UserDomainInterface)
	ret1, _ := ret[1].(*rest_err.RestError)
	return ret0, ret1
}

// FindUserByID indicates an expected call of FindUserByID.
func (mr *MockUserRepositoryMockRecorder) FindUserByID(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByID", reflect.TypeOf((*MockUserRepository)(nil).FindUserByID), id)
}

// UpdateUser mocks base method.
func (m *MockUserRepository) UpdateUser(userId string, userDomain model.UserDomainInterface) *rest_err.RestError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", userId, userDomain)
	ret0, _ := ret[0].(*rest_err.RestError)
	return ret0
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockUserRepositoryMockRecorder) UpdateUser(userId, userDomain any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockUserRepository)(nil).UpdateUser), userId, userDomain)
}