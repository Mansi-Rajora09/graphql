package api

import (
	"example/model/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthCheckHandler struct {
}

func NewHealthCheckHandler() HealthCheckHandlerInterface {
	return &HealthCheckHandler{}
}

type HealthCheckHandlerInterface interface {
	HealthCheck(context *gin.Context)
}

const (
	HealthCheckString = "/health-check"
)

// Healthcheck API Handler
func (hc HealthCheckHandler) HealthCheck(context *gin.Context) {
	context.JSON(http.StatusOK,
		response.NewHTTPResponse(200, "OK", "Healthcheck working"))

}
