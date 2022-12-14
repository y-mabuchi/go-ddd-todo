// Code generated by MockGen. DO NOT EDIT.
// Source: task.go

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/y-mabuchi/go-ddd-todo/domain"
)

// MockTaskUseCaseInterface is a mock of TaskUseCaseInterface interface.
type MockTaskUseCaseInterface struct {
	ctrl     *gomock.Controller
	recorder *MockTaskUseCaseInterfaceMockRecorder
}

// MockTaskUseCaseInterfaceMockRecorder is the mock recorder for MockTaskUseCaseInterface.
type MockTaskUseCaseInterfaceMockRecorder struct {
	mock *MockTaskUseCaseInterface
}

// NewMockTaskUseCaseInterface creates a new mock instance.
func NewMockTaskUseCaseInterface(ctrl *gomock.Controller) *MockTaskUseCaseInterface {
	mock := &MockTaskUseCaseInterface{ctrl: ctrl}
	mock.recorder = &MockTaskUseCaseInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTaskUseCaseInterface) EXPECT() *MockTaskUseCaseInterfaceMockRecorder {
	return m.recorder
}

// CreateTask mocks base method.
func (m *MockTaskUseCaseInterface) CreateTask(name string, dueDate time.Time) (*domain.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTask", name, dueDate)
	ret0, _ := ret[0].(*domain.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTask indicates an expected call of CreateTask.
func (mr *MockTaskUseCaseInterfaceMockRecorder) CreateTask(name, dueDate interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTask", reflect.TypeOf((*MockTaskUseCaseInterface)(nil).CreateTask), name, dueDate)
}

// PostponeTask mocks base method.
func (m *MockTaskUseCaseInterface) PostponeTask(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostponeTask", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// PostponeTask indicates an expected call of PostponeTask.
func (mr *MockTaskUseCaseInterfaceMockRecorder) PostponeTask(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostponeTask", reflect.TypeOf((*MockTaskUseCaseInterface)(nil).PostponeTask), id)
}
