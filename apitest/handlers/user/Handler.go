package user

import (
	"test_avns/apitest/helpers"
	"test_avns/apitest/interfaces"
	"test_avns/apitest/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	UserService interfaces.IUserService
}

func (u *UserHandler) AuthenticatUser(res http.ResponseWriter, req *http.Request) {
	var user models.User
	decoder := json.NewDecoder(req.Body)
	decoder.Decode(&user)
	respReg := u.UserService.AuthUser(req.Context(), user)
	helpers.APIResponse(res, respReg)
	return
}

func (u *UserHandler) Logout(res http.ResponseWriter, req *http.Request) {
	authSession := req.Header.Get("Authorization")
	respLogout := u.UserService.Logout(req.Context(), authSession)
	helpers.APIResponse(res, respLogout)
	return
}

func (u *UserHandler) Register(res http.ResponseWriter, req *http.Request) {
	var user models.User
	decoder := json.NewDecoder(req.Body)
	decoder.Decode(&user)
	respReg := u.UserService.Register(req.Context(), user)
	helpers.APIResponse(res, respReg)
	return
}

func (u *UserHandler) EditUser(res http.ResponseWriter, req *http.Request) {
	var user models.User
	decoder := json.NewDecoder(req.Body)
	decoder.Decode(&user)
	param := mux.Vars(req)
	user.ID = helpers.ToInt(param["id"])
	respEdit := u.UserService.EditUser(req.Context(), user)
	helpers.APIResponse(res, respEdit)
	return
}

func (u *UserHandler) DeleteUser(res http.ResponseWriter, req *http.Request) {
	var user models.User
	decoder := json.NewDecoder(req.Body)
	decoder.Decode(&user)
	param := mux.Vars(req)
	user.ID = helpers.ToInt(param["id"])
	respDel := u.UserService.DeleteUser(req.Context(), user)
	helpers.APIResponse(res, respDel)
	return
}

func (u *UserHandler) GetUserByUsername(res http.ResponseWriter, req *http.Request) {
	param := mux.Vars(req)
	var username = param["username"]
	respData := u.UserService.GetUserByUsername(req.Context(), username)
	helpers.APIResponse(res, respData)
	return
}
