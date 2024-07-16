// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/web-platform-tests/wpt.fyi/api/manifest (interfaces: API)
//
// Generated by this command:
//
//	mockgen -destination mock_manifest/api_mock.go github.com/web-platform-tests/wpt.fyi/api/manifest API
//

// Package mock_manifest is a generated GoMock package.
package mock_manifest

import (
	reflect "reflect"
	time "time"

	shared "github.com/web-platform-tests/wpt.fyi/shared"
	gomock "go.uber.org/mock/gomock"
)

// MockAPI is a mock of API interface.
type MockAPI struct {
	ctrl     *gomock.Controller
	recorder *MockAPIMockRecorder
}

// MockAPIMockRecorder is the mock recorder for MockAPI.
type MockAPIMockRecorder struct {
	mock *MockAPI
}

// NewMockAPI creates a new mock instance.
func NewMockAPI(ctrl *gomock.Controller) *MockAPI {
	mock := &MockAPI{ctrl: ctrl}
	mock.recorder = &MockAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAPI) EXPECT() *MockAPIMockRecorder {
	return m.recorder
}

// GetManifestForSHA mocks base method.
func (m *MockAPI) GetManifestForSHA(arg0 string) (string, []byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetManifestForSHA", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].([]byte)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetManifestForSHA indicates an expected call of GetManifestForSHA.
func (mr *MockAPIMockRecorder) GetManifestForSHA(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetManifestForSHA", reflect.TypeOf((*MockAPI)(nil).GetManifestForSHA), arg0)
}

// NewRedis mocks base method.
func (m *MockAPI) NewRedis(arg0 time.Duration) shared.ReadWritable {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewRedis", arg0)
	ret0, _ := ret[0].(shared.ReadWritable)
	return ret0
}

// NewRedis indicates an expected call of NewRedis.
func (mr *MockAPIMockRecorder) NewRedis(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewRedis", reflect.TypeOf((*MockAPI)(nil).NewRedis), arg0)
}
