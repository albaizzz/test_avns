package mock

import (
	"reflect"

	"github.com/golang/mock/gomock"
)

type AuthRepositoryMock struct {
	ctrl     *gomock.Controller
	recorder *AuthRepositoryMockRecorder
}

type AuthRepositoryMockRecorder struct {
	mock *AuthRepositoryMock
}

func NewAuthRepositoryMock(ctrl *gomock.Controller) *AuthRepositoryMock {
	mock := &AuthRepositoryMock{ctrl: ctrl}
	mock.recorder = &AuthRepositoryMockRecorder{mock: mock}
	return mock
}

func (a *AuthRepositoryMock) EXPECT() *AuthRepositoryMockRecorder {
	return a.recorder
}

func (a *AuthRepositoryMock) SaveSession(arg0 interface{}, arg1 interface{}) (err error) {
	ret := a.ctrl.Call(a, "SaveSession", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (r *AuthRepositoryMockRecorder) SaveSession(arg0 interface{}, arg1 interface{}) *gomock.Call {
	return r.mock.ctrl.RecordCallWithMethodType(r.mock, "SaveSession", reflect.TypeOf((*AuthRepositoryMock)(nil).SaveSession), arg0, arg1)
}

func (a *AuthRepositoryMock) Logout(arg0 interface{}, arg1 string) (err error) {
	ret := a.ctrl.Call(a, "Logout", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}
