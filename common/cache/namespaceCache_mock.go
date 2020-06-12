// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by MockGen. DO NOT EDIT.
// Source: namespaceCache.go

// Package cache is a generated GoMock package.
package cache

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockNamespaceCache is a mock of NamespaceCache interface
type MockNamespaceCache struct {
	ctrl     *gomock.Controller
	recorder *MockNamespaceCacheMockRecorder
}

// MockNamespaceCacheMockRecorder is the mock recorder for MockNamespaceCache
type MockNamespaceCacheMockRecorder struct {
	mock *MockNamespaceCache
}

// NewMockNamespaceCache creates a new mock instance
func NewMockNamespaceCache(ctrl *gomock.Controller) *MockNamespaceCache {
	mock := &MockNamespaceCache{ctrl: ctrl}
	mock.recorder = &MockNamespaceCacheMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNamespaceCache) EXPECT() *MockNamespaceCacheMockRecorder {
	return m.recorder
}

// Start mocks base method
func (m *MockNamespaceCache) Start() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Start")
}

// Start indicates an expected call of Start
func (mr *MockNamespaceCacheMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockNamespaceCache)(nil).Start))
}

// Stop mocks base method
func (m *MockNamespaceCache) Stop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop
func (mr *MockNamespaceCacheMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockNamespaceCache)(nil).Stop))
}

// RegisterNamespaceChangeCallback mocks base method
func (m *MockNamespaceCache) RegisterNamespaceChangeCallback(shard int, initialNotificationVersion int64, prepareCallback PrepareCallbackFn, callback CallbackFn) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterNamespaceChangeCallback", shard, initialNotificationVersion, prepareCallback, callback)
}

// RegisterNamespaceChangeCallback indicates an expected call of RegisterNamespaceChangeCallback
func (mr *MockNamespaceCacheMockRecorder) RegisterNamespaceChangeCallback(shard, initialNotificationVersion, prepareCallback, callback interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterNamespaceChangeCallback", reflect.TypeOf((*MockNamespaceCache)(nil).RegisterNamespaceChangeCallback), shard, initialNotificationVersion, prepareCallback, callback)
}

// UnregisterNamespaceChangeCallback mocks base method
func (m *MockNamespaceCache) UnregisterNamespaceChangeCallback(shard int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UnregisterNamespaceChangeCallback", shard)
}

// UnregisterNamespaceChangeCallback indicates an expected call of UnregisterNamespaceChangeCallback
func (mr *MockNamespaceCacheMockRecorder) UnregisterNamespaceChangeCallback(shard interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnregisterNamespaceChangeCallback", reflect.TypeOf((*MockNamespaceCache)(nil).UnregisterNamespaceChangeCallback), shard)
}

// GetNamespace mocks base method
func (m *MockNamespaceCache) GetNamespace(name string) (*NamespaceCacheEntry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNamespace", name)
	ret0, _ := ret[0].(*NamespaceCacheEntry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNamespace indicates an expected call of GetNamespace
func (mr *MockNamespaceCacheMockRecorder) GetNamespace(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNamespace", reflect.TypeOf((*MockNamespaceCache)(nil).GetNamespace), name)
}

// GetNamespaceByID mocks base method
func (m *MockNamespaceCache) GetNamespaceByID(id string) (*NamespaceCacheEntry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNamespaceByID", id)
	ret0, _ := ret[0].(*NamespaceCacheEntry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNamespaceByID indicates an expected call of GetNamespaceByID
func (mr *MockNamespaceCacheMockRecorder) GetNamespaceByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNamespaceByID", reflect.TypeOf((*MockNamespaceCache)(nil).GetNamespaceByID), id)
}

// GetNamespaceID mocks base method
func (m *MockNamespaceCache) GetNamespaceID(name string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNamespaceID", name)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNamespaceID indicates an expected call of GetNamespaceID
func (mr *MockNamespaceCacheMockRecorder) GetNamespaceID(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNamespaceID", reflect.TypeOf((*MockNamespaceCache)(nil).GetNamespaceID), name)
}

// GetNamespaceName mocks base method
func (m *MockNamespaceCache) GetNamespaceName(id string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNamespaceName", id)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNamespaceName indicates an expected call of GetNamespaceName
func (mr *MockNamespaceCacheMockRecorder) GetNamespaceName(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNamespaceName", reflect.TypeOf((*MockNamespaceCache)(nil).GetNamespaceName), id)
}

// GetAllNamespace mocks base method
func (m *MockNamespaceCache) GetAllNamespace() map[string]*NamespaceCacheEntry {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllNamespace")
	ret0, _ := ret[0].(map[string]*NamespaceCacheEntry)
	return ret0
}

// GetAllNamespace indicates an expected call of GetAllNamespace
func (mr *MockNamespaceCacheMockRecorder) GetAllNamespace() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllNamespace", reflect.TypeOf((*MockNamespaceCache)(nil).GetAllNamespace))
}

// GetCacheSize mocks base method
func (m *MockNamespaceCache) GetCacheSize() (int64, int64) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCacheSize")
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(int64)
	return ret0, ret1
}

// GetCacheSize indicates an expected call of GetCacheSize
func (mr *MockNamespaceCacheMockRecorder) GetCacheSize() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCacheSize", reflect.TypeOf((*MockNamespaceCache)(nil).GetCacheSize))
}
