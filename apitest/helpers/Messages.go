package helpers

import (
	"test_avns/apitest/constants"
	"test_avns/apitest/models"
	"net/http"
)

var Message = map[int]models.APIResponse{
	constants.APIGeneralSuccess:       models.APIResponse{Data: "OK", HttpCode: http.StatusOK},
	constants.APIAccepted:             models.APIResponse{Data: "OK", HttpCode: http.StatusAccepted},
	constants.APIInternalServerError:  models.APIResponse{Data: "Internal Server Error.", HttpCode: http.StatusInternalServerError},
	constants.APIUserAlreadyExist:     models.APIResponse{Data: "User or email already exist.", HttpCode: http.StatusConflict},
	constants.APIUserNotExist:         models.APIResponse{Data: "username not exist", HttpCode: http.StatusNotFound},
	constants.APIUserNotAuthenticated: models.APIResponse{Data: "User not authorized", HttpCode: http.StatusForbidden},
	constants.APIDataNotFound:         models.APIResponse{Data: "Data not exist", HttpCode: http.StatusNotFound},
}

func GetAPIResponse(apiStatusCode int, data ...interface{}) (resp models.APIResponse) {
	resp.Code = apiStatusCode
	if data == nil {
		resp = Message[apiStatusCode]
		resp.Code = apiStatusCode
	} else {
		resp.Data = data
		resp.HttpCode = http.StatusOK
	}
	return
}
