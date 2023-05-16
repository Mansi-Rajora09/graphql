package response

import "net/http"

const (
	SUCCESS = "success"
	FAILURE = "failure"
)

// HTTPResponse format
type HTTPResponse struct {
	Success string      `json:"success"`
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data,omitempty"`
}

// NewHTTPResponse for create common response
func NewHTTPResponse(code int, message string, data interface{}) *HTTPResponse {
	commonResponse := new(HTTPResponse)

	commonResponse.Data = data

	if code < http.StatusBadRequest {
		commonResponse.Success = SUCCESS
	}

	commonResponse.Code = code
	commonResponse.Message = message
	return commonResponse
}
