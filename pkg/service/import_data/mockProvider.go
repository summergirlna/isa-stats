// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/service/import_data/service.go
//
// Generated by this command:
//
//	mockgen -source ./pkg/service/import_data/service.go -destination ./pkg//service/import_data/mockProvider.go -package import_data
//

// Package import_data is a generated GoMock package.
package import_data

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
	isgomock struct{}
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// Import mocks base method.
func (m *MockService) Import(ctx context.Context, in *Input) (*Output, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Import", ctx, in)
	ret0, _ := ret[0].(*Output)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Import indicates an expected call of Import.
func (mr *MockServiceMockRecorder) Import(ctx, in any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Import", reflect.TypeOf((*MockService)(nil).Import), ctx, in)
}
