package healthcheck

import (
	"encoding/json"
	"net/http"

	"test_avns/apitest/constants"
	"test_avns/apitest/helpers"
	"test_avns/apitest/interfaces"
)

type HealthCheckHandler struct {
	HealthService interfaces.HealthServiceContract
}

func (h *HealthCheckHandler) GetStatus(res http.ResponseWriter, req *http.Request) {
	// resp := helpers.GetAPIResponse(constants.APIGeneralSuccess, nil)
	helpers.Response(res, http.StatusOK, nil, constants.APIGeneralSuccess)
	// res.Write(resp)
	// return
}

func (h *HealthCheckHandler) Readiness(res http.ResponseWriter, req *http.Request) {
	status, svc := h.HealthService.HealthStatus()
	res.WriteHeader(status)
	b, _ := json.Marshal(svc)
	res.Write(b)
	return
}
