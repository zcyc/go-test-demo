// Code generated by MockGen. DO NOT EDIT.
// Source: struct.go
//
// Generated by this command:
//
//	mockgen -source=struct.go -destination=mocks/gomock_struct.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	dao "MockTest/internal/user/dao"
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockUserDao is a mock of UserDao interface.
type MockUserDao struct {
	ctrl     *gomock.Controller
	recorder *MockUserDaoMockRecorder
}

// MockUserDaoMockRecorder is the mock recorder for MockUserDao.
type MockUserDaoMockRecorder struct {
	mock *MockUserDao
}

// NewMockUserDao creates a new mock instance.
func NewMockUserDao(ctrl *gomock.Controller) *MockUserDao {
	mock := &MockUserDao{ctrl: ctrl}
	mock.recorder = &MockUserDaoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserDao) EXPECT() *MockUserDaoMockRecorder {
	return m.recorder
}

// GetUserByMobile mocks base method.
func (m *MockUserDao) GetUserByMobile(ctx context.Context, mobile string) (*dao.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByMobile", ctx, mobile)
	ret0, _ := ret[0].(*dao.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByMobile indicates an expected call of GetUserByMobile.
func (mr *MockUserDaoMockRecorder) GetUserByMobile(ctx, mobile any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByMobile", reflect.TypeOf((*MockUserDao)(nil).GetUserByMobile), ctx, mobile)
}