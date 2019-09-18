// Copyright 2019 Honey Science Corporation
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, you can obtain one at http://mozilla.org/MPL/2.0/.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: internal/workflow/session.go

// Package mock_workflow is a generated GoMock package.
package mock_workflow

import (
	gomock "github.com/golang/mock/gomock"
	dipper "github.com/honeydipper/honeydipper/pkg/dipper"
	reflect "reflect"
)

// MockSessionHandler is a mock of SessionHandler interface
type MockSessionHandler struct {
	ctrl     *gomock.Controller
	recorder *MockSessionHandlerMockRecorder
}

// MockSessionHandlerMockRecorder is the mock recorder for MockSessionHandler
type MockSessionHandlerMockRecorder struct {
	mock *MockSessionHandler
}

// NewMockSessionHandler creates a new mock instance
func NewMockSessionHandler(ctrl *gomock.Controller) *MockSessionHandler {
	mock := &MockSessionHandler{ctrl: ctrl}
	mock.recorder = &MockSessionHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSessionHandler) EXPECT() *MockSessionHandlerMockRecorder {
	return m.recorder
}

// prepare mocks base method
func (m *MockSessionHandler) prepare(msg *dipper.Message, parent interface{}, ctx map[string]interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "prepare", msg, parent, ctx)
}

// prepare indicates an expected call of prepare
func (mr *MockSessionHandlerMockRecorder) prepare(msg, parent, ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "prepare", reflect.TypeOf((*MockSessionHandler)(nil).prepare), msg, parent, ctx)
}

// execute mocks base method
func (m *MockSessionHandler) execute(msg *dipper.Message) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "execute", msg)
}

// execute indicates an expected call of execute
func (mr *MockSessionHandlerMockRecorder) execute(msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "execute", reflect.TypeOf((*MockSessionHandler)(nil).execute), msg)
}

// continueExec mocks base method
func (m *MockSessionHandler) continueExec(msg *dipper.Message, export map[string]interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "continueExec", msg, export)
}

// continueExec indicates an expected call of continueExec
func (mr *MockSessionHandlerMockRecorder) continueExec(msg, export interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "continueExec", reflect.TypeOf((*MockSessionHandler)(nil).continueExec), msg, export)
}

// onError mocks base method
func (m *MockSessionHandler) onError() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "onError")
}

// onError indicates an expected call of onError
func (mr *MockSessionHandlerMockRecorder) onError() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "onError", reflect.TypeOf((*MockSessionHandler)(nil).onError))
}