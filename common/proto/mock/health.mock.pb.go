// Code generated by MockGen. DO NOT EDIT.
// Source: health.pb.go

// Package mock_proto is a generated GoMock package.
package mock_proto

import (
	context "context"
	reflect "reflect"

	proto "github.com/crafted-systems/smartcollect-pro/common/proto"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// MockHealthClient is a mock of HealthClient interface.
type MockHealthClient struct {
	ctrl     *gomock.Controller
	recorder *MockHealthClientMockRecorder
}

// MockHealthClientMockRecorder is the mock recorder for MockHealthClient.
type MockHealthClientMockRecorder struct {
	mock *MockHealthClient
}

// NewMockHealthClient creates a new mock instance.
func NewMockHealthClient(ctrl *gomock.Controller) *MockHealthClient {
	mock := &MockHealthClient{ctrl: ctrl}
	mock.recorder = &MockHealthClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHealthClient) EXPECT() *MockHealthClientMockRecorder {
	return m.recorder
}

// Health mocks base method.
func (m *MockHealthClient) Health(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*proto.HealthResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Health", varargs...)
	ret0, _ := ret[0].(*proto.HealthResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Health indicates an expected call of Health.
func (mr *MockHealthClientMockRecorder) Health(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Health", reflect.TypeOf((*MockHealthClient)(nil).Health), varargs...)
}

// MockHealthServer is a mock of HealthServer interface.
type MockHealthServer struct {
	ctrl     *gomock.Controller
	recorder *MockHealthServerMockRecorder
}

// MockHealthServerMockRecorder is the mock recorder for MockHealthServer.
type MockHealthServerMockRecorder struct {
	mock *MockHealthServer
}

// NewMockHealthServer creates a new mock instance.
func NewMockHealthServer(ctrl *gomock.Controller) *MockHealthServer {
	mock := &MockHealthServer{ctrl: ctrl}
	mock.recorder = &MockHealthServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHealthServer) EXPECT() *MockHealthServerMockRecorder {
	return m.recorder
}

// Health mocks base method.
func (m *MockHealthServer) Health(arg0 context.Context, arg1 *emptypb.Empty) (*proto.HealthResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Health", arg0, arg1)
	ret0, _ := ret[0].(*proto.HealthResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Health indicates an expected call of Health.
func (mr *MockHealthServerMockRecorder) Health(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Health", reflect.TypeOf((*MockHealthServer)(nil).Health), arg0, arg1)
}
