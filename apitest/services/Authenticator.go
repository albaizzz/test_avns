package services

import (
	"test_avns/apitest/constants"
	"test_avns/apitest/helpers"
	"test_avns/apitest/interfaces"
	"database/sql"
	"net/http"
	"strconv"
)

type Authenticator struct {
	AuthRepo interfaces.IAuthRepository
}

func NewAuthenticator(AuthRepository interfaces.IAuthRepository) *Authenticator {
	return &Authenticator{
		AuthRepo: AuthRepository,
	}
}

func (a *Authenticator) Authenticate(res http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	authSession := req.Header.Get("Authorization")

	if authSession == "" {
		resp := helpers.GetAPIResponse(constants.APIUserNotAuthenticated)
		helpers.APIResponse(res, resp)
		return
	}

	userAuth, err := a.AuthRepo.GetSession(req.Context(), authSession)
	if err != nil {
		if err == sql.ErrNoRows {
			resp := helpers.GetAPIResponse(constants.APIUserNotAuthenticated)
			helpers.APIResponse(res, resp)
			return
		}
		resp := helpers.GetAPIResponse(constants.APIInternalServerError)
		helpers.APIResponse(res, resp)
		return
	}

	req.Header.Set("username", userAuth.Username)
	req.Header.Set("user_id", strconv.Itoa(userAuth.UserId))

	next(res, req)
}
