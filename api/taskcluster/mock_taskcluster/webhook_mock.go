// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/web-platform-tests/wpt.fyi/api/taskcluster (interfaces: API)
//
// Generated by this command:
//
//	mockgen -build_flags=--mod=mod -destination mock_taskcluster/webhook_mock.go github.com/web-platform-tests/wpt.fyi/api/taskcluster API
//

// Package mock_taskcluster is a generated GoMock package.
package mock_taskcluster

import (
	reflect "reflect"

	github "github.com/google/go-github/v65/github"
	taskcluster "github.com/web-platform-tests/wpt.fyi/api/taskcluster"
	gomock "go.uber.org/mock/gomock"
)

// MockAPI is a mock of API interface.
type MockAPI struct {
	ctrl     *gomock.Controller
	recorder *MockAPIMockRecorder
	isgomock struct{}
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

// GetTaskGroupInfo mocks base method.
func (m *MockAPI) GetTaskGroupInfo(arg0, arg1 string) (*taskcluster.TaskGroupInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTaskGroupInfo", arg0, arg1)
	ret0, _ := ret[0].(*taskcluster.TaskGroupInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTaskGroupInfo indicates an expected call of GetTaskGroupInfo.
func (mr *MockAPIMockRecorder) GetTaskGroupInfo(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTaskGroupInfo", reflect.TypeOf((*MockAPI)(nil).GetTaskGroupInfo), arg0, arg1)
}

// ListCheckRuns mocks base method.
func (m *MockAPI) ListCheckRuns(owner, repo string, checkSuiteID int64) ([]*github.CheckRun, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListCheckRuns", owner, repo, checkSuiteID)
	ret0, _ := ret[0].([]*github.CheckRun)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCheckRuns indicates an expected call of ListCheckRuns.
func (mr *MockAPIMockRecorder) ListCheckRuns(owner, repo, checkSuiteID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCheckRuns", reflect.TypeOf((*MockAPI)(nil).ListCheckRuns), owner, repo, checkSuiteID)
}
