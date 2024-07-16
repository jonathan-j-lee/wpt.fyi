// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/web-platform-tests/wpt.fyi/shared (interfaces: TestRunQuery)
//
// Generated by this command:
//
//	mockgen -destination sharedtest/test_run_query_mock.go -package sharedtest github.com/web-platform-tests/wpt.fyi/shared TestRunQuery
//

// Package sharedtest is a generated GoMock package.
package sharedtest

import (
	reflect "reflect"
	time "time"

	mapset "github.com/deckarep/golang-set"
	shared "github.com/web-platform-tests/wpt.fyi/shared"
	gomock "go.uber.org/mock/gomock"
)

// MockTestRunQuery is a mock of TestRunQuery interface.
type MockTestRunQuery struct {
	ctrl     *gomock.Controller
	recorder *MockTestRunQueryMockRecorder
}

// MockTestRunQueryMockRecorder is the mock recorder for MockTestRunQuery.
type MockTestRunQueryMockRecorder struct {
	mock *MockTestRunQuery
}

// NewMockTestRunQuery creates a new mock instance.
func NewMockTestRunQuery(ctrl *gomock.Controller) *MockTestRunQuery {
	mock := &MockTestRunQuery{ctrl: ctrl}
	mock.recorder = &MockTestRunQueryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTestRunQuery) EXPECT() *MockTestRunQueryMockRecorder {
	return m.recorder
}

// GetAlignedRunSHAs mocks base method.
func (m *MockTestRunQuery) GetAlignedRunSHAs(arg0 shared.ProductSpecs, arg1 mapset.Set, arg2, arg3 *time.Time, arg4, arg5 *int) ([]string, map[string]shared.KeysByProduct, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAlignedRunSHAs", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(map[string]shared.KeysByProduct)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetAlignedRunSHAs indicates an expected call of GetAlignedRunSHAs.
func (mr *MockTestRunQueryMockRecorder) GetAlignedRunSHAs(arg0, arg1, arg2, arg3, arg4, arg5 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAlignedRunSHAs", reflect.TypeOf((*MockTestRunQuery)(nil).GetAlignedRunSHAs), arg0, arg1, arg2, arg3, arg4, arg5)
}

// LoadTestRunKeys mocks base method.
func (m *MockTestRunQuery) LoadTestRunKeys(arg0 []shared.ProductSpec, arg1 mapset.Set, arg2 []string, arg3, arg4 *time.Time, arg5, arg6 *int) (shared.KeysByProduct, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadTestRunKeys", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	ret0, _ := ret[0].(shared.KeysByProduct)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadTestRunKeys indicates an expected call of LoadTestRunKeys.
func (mr *MockTestRunQueryMockRecorder) LoadTestRunKeys(arg0, arg1, arg2, arg3, arg4, arg5, arg6 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadTestRunKeys", reflect.TypeOf((*MockTestRunQuery)(nil).LoadTestRunKeys), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// LoadTestRuns mocks base method.
func (m *MockTestRunQuery) LoadTestRuns(arg0 []shared.ProductSpec, arg1 mapset.Set, arg2 []string, arg3, arg4 *time.Time, arg5, arg6 *int) (shared.TestRunsByProduct, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadTestRuns", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	ret0, _ := ret[0].(shared.TestRunsByProduct)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadTestRuns indicates an expected call of LoadTestRuns.
func (mr *MockTestRunQueryMockRecorder) LoadTestRuns(arg0, arg1, arg2, arg3, arg4, arg5, arg6 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadTestRuns", reflect.TypeOf((*MockTestRunQuery)(nil).LoadTestRuns), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// LoadTestRunsByKeys mocks base method.
func (m *MockTestRunQuery) LoadTestRunsByKeys(arg0 shared.KeysByProduct) (shared.TestRunsByProduct, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadTestRunsByKeys", arg0)
	ret0, _ := ret[0].(shared.TestRunsByProduct)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadTestRunsByKeys indicates an expected call of LoadTestRunsByKeys.
func (mr *MockTestRunQueryMockRecorder) LoadTestRunsByKeys(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadTestRunsByKeys", reflect.TypeOf((*MockTestRunQuery)(nil).LoadTestRunsByKeys), arg0)
}
