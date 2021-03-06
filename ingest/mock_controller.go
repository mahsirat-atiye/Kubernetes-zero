// Code generated by MockGen. DO NOT EDIT.
// Source: controller.go

// Package mock_ingest is a generated GoMock package.
package ingest

import (
	reflect "reflect"

	gin "github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
)

// MockController is a mock of Controller interface.
type MockController struct {
	ctrl     *gomock.Controller
	recorder *MockControllerMockRecorder
}

// MockControllerMockRecorder is the mock recorder for MockController.
type MockControllerMockRecorder struct {
	mock *MockController
}

// NewMockController creates a new mock instance.
func NewMockController(ctrl *gomock.Controller) *MockController {
	mock := &MockController{ctrl: ctrl}
	mock.recorder = &MockControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockController) EXPECT() *MockControllerMockRecorder {
	return m.recorder
}

// CreateBook mocks base method.
func (m *MockController) CreateBook(ctx *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CreateBook", ctx)
}

// CreateBook indicates an expected call of CreateBook.
func (mr *MockControllerMockRecorder) CreateBook(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBook", reflect.TypeOf((*MockController)(nil).CreateBook), ctx)
}
