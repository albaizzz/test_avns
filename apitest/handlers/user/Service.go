package user

import (
	"test_avns/apitest/constants"
	"test_avns/apitest/helpers"
	"test_avns/apitest/interfaces"
	"test_avns/apitest/models"
	"context"
	"database/sql"
	"net/http"
	"time"

	"github.com/go-sql-driver/mysql"
)

type UserService struct {
	UserRepository interfaces.IUserRepository
	AuthRepository interfaces.IAuthRepository
}

func NewUserService(userRepository interfaces.IUserRepository, authRepo interfaces.IAuthRepository) *UserService {
	return &UserService{
		UserRepository: userRepository,
		AuthRepository: authRepo,
	}
}

func (u *UserService) AuthUser(ctx context.Context, user models.User) (response models.APIResponse) {
	userData, err := u.UserRepository.GetUserByUsername(ctx, user.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			response = helpers.GetAPIResponse(constants.APIUserNotAuthenticated)
			return
		}
		response = helpers.GetAPIResponse(constants.APIInternalServerError)
		return
	}

	validPass := helpers.ComparePasswords(userData.Password, []byte(user.Password))
	if validPass {
		//save to redis
		userData.Password, userData.Address = "", ""
		authCode := helpers.GenerateAuthCode()
		userSession := models.UserSession{
			UserId:      userData.ID,
			Session:     authCode,
			User:        userData,
			ExpiredDate: time.Now().AddDate(0, 0, 3),
		}
		err := u.AuthRepository.SaveSession(ctx, userSession)
		if err != nil {
			response = helpers.GetAPIResponse(constants.APIInternalServerError)
			return
		}
		response = helpers.GetAPIResponse(constants.APIGeneralSuccess, authCode)
	} else {
		response = helpers.GetAPIResponse(constants.APIUserNotAuthenticated)
	}
	return
}

func (u *UserService) Register(ctx context.Context, user models.User) (response models.APIResponse) {
	//encrypted password
	password := helpers.HashAndSalt([]byte(user.Password))
	user.Password = password
	err := u.UserRepository.Add(ctx, user)
	if err != nil {
		me, ok := err.(*mysql.MySQLError)
		if !ok {
			response = helpers.GetAPIResponse(constants.APIInternalServerError)
		}
		if me.Number == 1062 {
			response = helpers.GetAPIResponse(constants.APIUserAlreadyExist)
		}
		return
	}
	response.Data = "OK"
	response.HttpCode = http.StatusCreated
	return
}

func (u *UserService) GetUserByUsername(ctx context.Context, username string) (response models.APIResponse) {

	user, err := u.UserRepository.GetUserByUsername(ctx, username)

	if err != nil {
		if err == sql.ErrNoRows {
			response = helpers.GetAPIResponse(constants.APIDataNotFound)
			return
		}
		response = helpers.GetAPIResponse(constants.APIInternalServerError)
		return
	}

	response = helpers.GetAPIResponse(constants.APIGeneralSuccess, user)
	return
}

func (u *UserService) EditUser(ctx context.Context, user models.User) (resp models.APIResponse) {
	if user.Password != "" {
		user.Password = helpers.HashAndSalt([]byte(user.Password))
	}
	err := u.UserRepository.Edit(ctx, user)
	if err != nil {
		resp = helpers.GetAPIResponse(constants.APIInternalServerError)
		return
	}
	resp = helpers.GetAPIResponse(constants.APIAccepted)
	return
}

func (u *UserService) DeleteUser(ctx context.Context, user models.User) (resp models.APIResponse) {
	err := u.UserRepository.Delete(ctx, user.ID)
	if err != nil {
		resp = helpers.GetAPIResponse(constants.APIInternalServerError)
		return
	}
	resp = helpers.GetAPIResponse(constants.APIAccepted)
	return
}

func (u *UserService) Logout(ctx context.Context, session string) (resp models.APIResponse) {
	err := u.AuthRepository.Logout(ctx, session)
	if err != nil {
		resp = helpers.GetAPIResponse(constants.APIInternalServerError)
		return
	}
	resp = helpers.GetAPIResponse(constants.APIGeneralSuccess)
	return
}
