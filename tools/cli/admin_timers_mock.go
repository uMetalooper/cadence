// The MIT License (MIT)

// Copyright (c) 2017-2020 Uber Technologies Inc.

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/uber/cadence/tools/cli (interfaces: LoadCloser,Printer)
//
// Generated by this command:
//
//	mockgen -package cli -destination admin_timers_mock.go github.com/uber/cadence/tools/cli LoadCloser,Printer
//

// Package cli is a generated GoMock package.
package cli

import (
	io "io"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"

	persistence "github.com/uber/cadence/common/persistence"
)

// MockLoadCloser is a mock of LoadCloser interface.
type MockLoadCloser struct {
	ctrl     *gomock.Controller
	recorder *MockLoadCloserMockRecorder
	isgomock struct{}
}

// MockLoadCloserMockRecorder is the mock recorder for MockLoadCloser.
type MockLoadCloserMockRecorder struct {
	mock *MockLoadCloser
}

// NewMockLoadCloser creates a new mock instance.
func NewMockLoadCloser(ctrl *gomock.Controller) *MockLoadCloser {
	mock := &MockLoadCloser{ctrl: ctrl}
	mock.recorder = &MockLoadCloserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLoadCloser) EXPECT() *MockLoadCloserMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockLoadCloser) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockLoadCloserMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockLoadCloser)(nil).Close))
}

// Load mocks base method.
func (m *MockLoadCloser) Load() ([]*persistence.TimerTaskInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Load")
	ret0, _ := ret[0].([]*persistence.TimerTaskInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Load indicates an expected call of Load.
func (mr *MockLoadCloserMockRecorder) Load() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Load", reflect.TypeOf((*MockLoadCloser)(nil).Load))
}

// MockPrinter is a mock of Printer interface.
type MockPrinter struct {
	ctrl     *gomock.Controller
	recorder *MockPrinterMockRecorder
	isgomock struct{}
}

// MockPrinterMockRecorder is the mock recorder for MockPrinter.
type MockPrinterMockRecorder struct {
	mock *MockPrinter
}

// NewMockPrinter creates a new mock instance.
func NewMockPrinter(ctrl *gomock.Controller) *MockPrinter {
	mock := &MockPrinter{ctrl: ctrl}
	mock.recorder = &MockPrinterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPrinter) EXPECT() *MockPrinterMockRecorder {
	return m.recorder
}

// Print mocks base method.
func (m *MockPrinter) Print(output io.Writer, timers []*persistence.TimerTaskInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Print", output, timers)
	ret0, _ := ret[0].(error)
	return ret0
}

// Print indicates an expected call of Print.
func (mr *MockPrinterMockRecorder) Print(output, timers any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Print", reflect.TypeOf((*MockPrinter)(nil).Print), output, timers)
}
