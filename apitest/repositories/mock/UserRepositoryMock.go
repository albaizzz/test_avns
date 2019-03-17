package mock

import (
	"test_avns/apitest/models"
	"reflect"

	"github.com/golang/mock/gomock"
)

type UserRepositoryMock struct {
	ctrl     *gomock.Controller
	recorder *UserRepositoryMockRecorder
}

type UserRepositoryMockRecorder struct {
	mock *UserRepositoryMock
}

func NewUserRepositoryMock(ctrl *gomock.Controller) *UserRepositoryMock {
	mock := &UserRepositoryMock{ctrl: ctrl}
	mock.recorder = &UserRepositoryMockRecorder{mock: mock}
	return mock
}

func (a *UserRepositoryMock) EXPECT() *UserRepositoryMockRecorder {
	return a.recorder
}

func (u *UserRepositoryMock) GetUserByUsername(arg0 interface{}, arg1 interface{}) (user models.User, err error) {
	ret := u.ctrl.Call(u, "GetUserByUsername", arg0, arg1)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (ur *UserRepositoryMockRecorder) GetUserByUsername(arg0 interface{}, arg1 interface{}) *gomock.Call {
	return ur.mock.ctrl.RecordCallWithMethodType(ur.mock, "GetUserByUsername", reflect.TypeOf((*UserRepositoryMock)(nil).GetUserByUsername), arg0, arg1)
}

func (u *UserRepositoryMock) GetUsers(arg0 interface{}) (users []models.User, err error) {
	ret := u.ctrl.Call(u, "GetUsers", arg0)
	ret0, _ := ret[0].([]models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (u *UserRepositoryMock) Add(arg0 interface{}, arg1 interface{}) error {
	ret := u.ctrl.Call(u, "Add", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (u *UserRepositoryMock) Edit(arg0 interface{}, arg1 interface{}) error {
	ret := u.ctrl.Call(u, "Edit", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (u *UserRepositoryMock) Delete(arg0 interface{}, arg1 interface{}) error {
	ret := u.ctrl.Call(u, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}
