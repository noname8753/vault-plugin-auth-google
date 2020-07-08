// Code generated by MockGen. DO NOT EDIT.
// Source: provider.go

// Package google is a generated GoMock package.
package google

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	oauth2 "golang.org/x/oauth2"
	oauth20 "google.golang.org/api/oauth2/v2"
	reflect "reflect"
)

// MockUserProvider is a mock of UserProvider interface
type MockUserProvider struct {
	ctrl     *gomock.Controller
	recorder *MockUserProviderMockRecorder
}

// MockUserProviderMockRecorder is the mock recorder for MockUserProvider
type MockUserProviderMockRecorder struct {
	mock *MockUserProvider
}

// NewMockUserProvider creates a new mock instance
func NewMockUserProvider(ctrl *gomock.Controller) *MockUserProvider {
	mock := &MockUserProvider{ctrl: ctrl}
	mock.recorder = &MockUserProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserProvider) EXPECT() *MockUserProviderMockRecorder {
	return m.recorder
}

// authUser mocks base method
func (m *MockUserProvider) authUser(ctx context.Context, config *oauth2.Config, token *oauth2.Token) (*oauth20.Userinfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "authUser", ctx, config, token)
	ret0, _ := ret[0].(*oauth20.Userinfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// authUser indicates an expected call of authUser
func (mr *MockUserProviderMockRecorder) authUser(ctx, config, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "authUser", reflect.TypeOf((*MockUserProvider)(nil).authUser), ctx, config, token)
}

// oauth2Exchange mocks base method
func (m *MockUserProvider) oauth2Exchange(ctx context.Context, code string, config *oauth2.Config) (*oauth2.Token, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "oauth2Exchange", ctx, code, config)
	ret0, _ := ret[0].(*oauth2.Token)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// oauth2Exchange indicates an expected call of oauth2Exchange
func (mr *MockUserProviderMockRecorder) oauth2Exchange(ctx, code, config interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "oauth2Exchange", reflect.TypeOf((*MockUserProvider)(nil).oauth2Exchange), ctx, code, config)
}