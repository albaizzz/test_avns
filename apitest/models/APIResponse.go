package models

// APIHeader defines attributes for api header
type APIHeader struct {
	Authorization string `valid:"required"`
	Channel       string `valid:"required"`
	ClientVersion int
}

// APIResponse defines attributes for api Response
type APIResponse struct {
	Code     interface{} `json:"code"`
	Data     interface{} `json:"data"`
	HttpCode int
}
