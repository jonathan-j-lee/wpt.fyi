// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/web-platform-tests/wpt.fyi/shared (interfaces: CachedStore,ObjectCache,ObjectStore,ReadWritable,Readable,RedisSet)
//
// Generated by this command:
//
//	mockgen -destination sharedtest/cache_mock.go -package sharedtest github.com/web-platform-tests/wpt.fyi/shared CachedStore,ObjectCache,ObjectStore,ReadWritable,Readable,RedisSet
//

// Package sharedtest is a generated GoMock package.
package sharedtest

import (
	io "io"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockCachedStore is a mock of CachedStore interface.
type MockCachedStore struct {
	ctrl     *gomock.Controller
	recorder *MockCachedStoreMockRecorder
}

// MockCachedStoreMockRecorder is the mock recorder for MockCachedStore.
type MockCachedStoreMockRecorder struct {
	mock *MockCachedStore
}

// NewMockCachedStore creates a new mock instance.
func NewMockCachedStore(ctrl *gomock.Controller) *MockCachedStore {
	mock := &MockCachedStore{ctrl: ctrl}
	mock.recorder = &MockCachedStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCachedStore) EXPECT() *MockCachedStoreMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockCachedStore) Get(arg0, arg1, arg2 any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockCachedStoreMockRecorder) Get(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockCachedStore)(nil).Get), arg0, arg1, arg2)
}

// MockObjectCache is a mock of ObjectCache interface.
type MockObjectCache struct {
	ctrl     *gomock.Controller
	recorder *MockObjectCacheMockRecorder
}

// MockObjectCacheMockRecorder is the mock recorder for MockObjectCache.
type MockObjectCacheMockRecorder struct {
	mock *MockObjectCache
}

// NewMockObjectCache creates a new mock instance.
func NewMockObjectCache(ctrl *gomock.Controller) *MockObjectCache {
	mock := &MockObjectCache{ctrl: ctrl}
	mock.recorder = &MockObjectCacheMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockObjectCache) EXPECT() *MockObjectCacheMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockObjectCache) Get(arg0, arg1 any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockObjectCacheMockRecorder) Get(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockObjectCache)(nil).Get), arg0, arg1)
}

// Put mocks base method.
func (m *MockObjectCache) Put(arg0, arg1 any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Put", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Put indicates an expected call of Put.
func (mr *MockObjectCacheMockRecorder) Put(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockObjectCache)(nil).Put), arg0, arg1)
}

// MockObjectStore is a mock of ObjectStore interface.
type MockObjectStore struct {
	ctrl     *gomock.Controller
	recorder *MockObjectStoreMockRecorder
}

// MockObjectStoreMockRecorder is the mock recorder for MockObjectStore.
type MockObjectStoreMockRecorder struct {
	mock *MockObjectStore
}

// NewMockObjectStore creates a new mock instance.
func NewMockObjectStore(ctrl *gomock.Controller) *MockObjectStore {
	mock := &MockObjectStore{ctrl: ctrl}
	mock.recorder = &MockObjectStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockObjectStore) EXPECT() *MockObjectStoreMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockObjectStore) Get(arg0, arg1 any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockObjectStoreMockRecorder) Get(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockObjectStore)(nil).Get), arg0, arg1)
}

// MockReadWritable is a mock of ReadWritable interface.
type MockReadWritable struct {
	ctrl     *gomock.Controller
	recorder *MockReadWritableMockRecorder
}

// MockReadWritableMockRecorder is the mock recorder for MockReadWritable.
type MockReadWritableMockRecorder struct {
	mock *MockReadWritable
}

// NewMockReadWritable creates a new mock instance.
func NewMockReadWritable(ctrl *gomock.Controller) *MockReadWritable {
	mock := &MockReadWritable{ctrl: ctrl}
	mock.recorder = &MockReadWritableMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReadWritable) EXPECT() *MockReadWritableMockRecorder {
	return m.recorder
}

// NewReadCloser mocks base method.
func (m *MockReadWritable) NewReadCloser(arg0 any) (io.ReadCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewReadCloser", arg0)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewReadCloser indicates an expected call of NewReadCloser.
func (mr *MockReadWritableMockRecorder) NewReadCloser(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewReadCloser", reflect.TypeOf((*MockReadWritable)(nil).NewReadCloser), arg0)
}

// NewWriteCloser mocks base method.
func (m *MockReadWritable) NewWriteCloser(arg0 any) (io.WriteCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewWriteCloser", arg0)
	ret0, _ := ret[0].(io.WriteCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewWriteCloser indicates an expected call of NewWriteCloser.
func (mr *MockReadWritableMockRecorder) NewWriteCloser(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewWriteCloser", reflect.TypeOf((*MockReadWritable)(nil).NewWriteCloser), arg0)
}

// MockReadable is a mock of Readable interface.
type MockReadable struct {
	ctrl     *gomock.Controller
	recorder *MockReadableMockRecorder
}

// MockReadableMockRecorder is the mock recorder for MockReadable.
type MockReadableMockRecorder struct {
	mock *MockReadable
}

// NewMockReadable creates a new mock instance.
func NewMockReadable(ctrl *gomock.Controller) *MockReadable {
	mock := &MockReadable{ctrl: ctrl}
	mock.recorder = &MockReadableMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReadable) EXPECT() *MockReadableMockRecorder {
	return m.recorder
}

// NewReadCloser mocks base method.
func (m *MockReadable) NewReadCloser(arg0 any) (io.ReadCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewReadCloser", arg0)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewReadCloser indicates an expected call of NewReadCloser.
func (mr *MockReadableMockRecorder) NewReadCloser(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewReadCloser", reflect.TypeOf((*MockReadable)(nil).NewReadCloser), arg0)
}

// MockRedisSet is a mock of RedisSet interface.
type MockRedisSet struct {
	ctrl     *gomock.Controller
	recorder *MockRedisSetMockRecorder
}

// MockRedisSetMockRecorder is the mock recorder for MockRedisSet.
type MockRedisSetMockRecorder struct {
	mock *MockRedisSet
}

// NewMockRedisSet creates a new mock instance.
func NewMockRedisSet(ctrl *gomock.Controller) *MockRedisSet {
	mock := &MockRedisSet{ctrl: ctrl}
	mock.recorder = &MockRedisSetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRedisSet) EXPECT() *MockRedisSetMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockRedisSet) Add(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockRedisSetMockRecorder) Add(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockRedisSet)(nil).Add), arg0, arg1)
}

// GetAll mocks base method.
func (m *MockRedisSet) GetAll(arg0 string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockRedisSetMockRecorder) GetAll(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockRedisSet)(nil).GetAll), arg0)
}

// Remove mocks base method.
func (m *MockRedisSet) Remove(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remove", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove.
func (mr *MockRedisSetMockRecorder) Remove(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockRedisSet)(nil).Remove), arg0, arg1)
}
