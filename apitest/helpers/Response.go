package helpers

import (
	"encoding/json"
	"net/http"

	"test_avns/apitest/models"
)

// APIResponse defines attributes of API Response
type APIResponsemodel struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

func APIResponse(w http.ResponseWriter, resp models.APIResponse) {
	apiResponse := new(APIResponsemodel)
	apiResponse.Data = resp.Data
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.HttpCode)
	json.NewEncoder(w).Encode(apiResponse)
}

// Response sets the json response output to client
func Response(w http.ResponseWriter, httpStatus int, data interface{}, code int) {
	apiResponse := new(APIResponsemodel)
	apiResponse.Data = data
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	json.NewEncoder(w).Encode(apiResponse)
}

// ResponseResult sets the json response output to client
func ResponseResult(w http.ResponseWriter, data interface{}, httpStatus int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	json.NewEncoder(w).Encode(data)
}
