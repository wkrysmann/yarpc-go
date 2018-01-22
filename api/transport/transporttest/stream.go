// Code generated by MockGen. DO NOT EDIT.
// Source: go.uber.org/yarpc/api/transport (interfaces: Stream,StreamCloser)

// Copyright (c) 2018 Uber Technologies, Inc.
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

package transporttest

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	transport "go.uber.org/yarpc/api/transport"
	reflect "reflect"
)

// MockStream is a mock of Stream interface
type MockStream struct {
	ctrl     *gomock.Controller
	recorder *MockStreamMockRecorder
}

// MockStreamMockRecorder is the mock recorder for MockStream
type MockStreamMockRecorder struct {
	mock *MockStream
}

// NewMockStream creates a new mock instance
func NewMockStream(ctrl *gomock.Controller) *MockStream {
	mock := &MockStream{ctrl: ctrl}
	mock.recorder = &MockStreamMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockStream) EXPECT() *MockStreamMockRecorder {
	return _m.recorder
}

// Context mocks base method
func (_m *MockStream) Context() context.Context {
	ret := _m.ctrl.Call(_m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (_mr *MockStreamMockRecorder) Context() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Context", reflect.TypeOf((*MockStream)(nil).Context))
}

// ReceiveMessage mocks base method
func (_m *MockStream) ReceiveMessage(_param0 context.Context) (*transport.StreamMessage, error) {
	ret := _m.ctrl.Call(_m, "ReceiveMessage", _param0)
	ret0, _ := ret[0].(*transport.StreamMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReceiveMessage indicates an expected call of ReceiveMessage
func (_mr *MockStreamMockRecorder) ReceiveMessage(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "ReceiveMessage", reflect.TypeOf((*MockStream)(nil).ReceiveMessage), arg0)
}

// Request mocks base method
func (_m *MockStream) Request() *transport.StreamRequest {
	ret := _m.ctrl.Call(_m, "Request")
	ret0, _ := ret[0].(*transport.StreamRequest)
	return ret0
}

// Request indicates an expected call of Request
func (_mr *MockStreamMockRecorder) Request() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Request", reflect.TypeOf((*MockStream)(nil).Request))
}

// SendMessage mocks base method
func (_m *MockStream) SendMessage(_param0 context.Context, _param1 *transport.StreamMessage) error {
	ret := _m.ctrl.Call(_m, "SendMessage", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMessage indicates an expected call of SendMessage
func (_mr *MockStreamMockRecorder) SendMessage(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SendMessage", reflect.TypeOf((*MockStream)(nil).SendMessage), arg0, arg1)
}

// MockStreamCloser is a mock of StreamCloser interface
type MockStreamCloser struct {
	ctrl     *gomock.Controller
	recorder *MockStreamCloserMockRecorder
}

// MockStreamCloserMockRecorder is the mock recorder for MockStreamCloser
type MockStreamCloserMockRecorder struct {
	mock *MockStreamCloser
}

// NewMockStreamCloser creates a new mock instance
func NewMockStreamCloser(ctrl *gomock.Controller) *MockStreamCloser {
	mock := &MockStreamCloser{ctrl: ctrl}
	mock.recorder = &MockStreamCloserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockStreamCloser) EXPECT() *MockStreamCloserMockRecorder {
	return _m.recorder
}

// Close mocks base method
func (_m *MockStreamCloser) Close(_param0 context.Context) error {
	ret := _m.ctrl.Call(_m, "Close", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (_mr *MockStreamCloserMockRecorder) Close(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Close", reflect.TypeOf((*MockStreamCloser)(nil).Close), arg0)
}

// Context mocks base method
func (_m *MockStreamCloser) Context() context.Context {
	ret := _m.ctrl.Call(_m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (_mr *MockStreamCloserMockRecorder) Context() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Context", reflect.TypeOf((*MockStreamCloser)(nil).Context))
}

// ReceiveMessage mocks base method
func (_m *MockStreamCloser) ReceiveMessage(_param0 context.Context) (*transport.StreamMessage, error) {
	ret := _m.ctrl.Call(_m, "ReceiveMessage", _param0)
	ret0, _ := ret[0].(*transport.StreamMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReceiveMessage indicates an expected call of ReceiveMessage
func (_mr *MockStreamCloserMockRecorder) ReceiveMessage(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "ReceiveMessage", reflect.TypeOf((*MockStreamCloser)(nil).ReceiveMessage), arg0)
}

// Request mocks base method
func (_m *MockStreamCloser) Request() *transport.StreamRequest {
	ret := _m.ctrl.Call(_m, "Request")
	ret0, _ := ret[0].(*transport.StreamRequest)
	return ret0
}

// Request indicates an expected call of Request
func (_mr *MockStreamCloserMockRecorder) Request() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Request", reflect.TypeOf((*MockStreamCloser)(nil).Request))
}

// SendMessage mocks base method
func (_m *MockStreamCloser) SendMessage(_param0 context.Context, _param1 *transport.StreamMessage) error {
	ret := _m.ctrl.Call(_m, "SendMessage", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMessage indicates an expected call of SendMessage
func (_mr *MockStreamCloserMockRecorder) SendMessage(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SendMessage", reflect.TypeOf((*MockStreamCloser)(nil).SendMessage), arg0, arg1)
}
