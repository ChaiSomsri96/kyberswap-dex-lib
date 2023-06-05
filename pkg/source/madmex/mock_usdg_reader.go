// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/madmex (interfaces: IUSDGReader)

// Package madmex is a generated GoMock package.
package madmex

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIUSDGReader is a mock of IUSDGReader interface.
type MockIUSDGReader struct {
	ctrl     *gomock.Controller
	recorder *MockIUSDGReaderMockRecorder
}

// MockIUSDGReaderMockRecorder is the mock recorder for MockIUSDGReader.
type MockIUSDGReaderMockRecorder struct {
	mock *MockIUSDGReader
}

// NewMockIUSDGReader creates a new mock instance.
func NewMockIUSDGReader(ctrl *gomock.Controller) *MockIUSDGReader {
	mock := &MockIUSDGReader{ctrl: ctrl}
	mock.recorder = &MockIUSDGReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUSDGReader) EXPECT() *MockIUSDGReaderMockRecorder {
	return m.recorder
}

// Read mocks base method.
func (m *MockIUSDGReader) Read(arg0 context.Context, arg1 string) (*USDG, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", arg0, arg1)
	ret0, _ := ret[0].(*USDG)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockIUSDGReaderMockRecorder) Read(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockIUSDGReader)(nil).Read), arg0, arg1)
}