package errorhandler

import (
	response "Pokemon-API/Response"
	"net/http"
)

var (
	getStatusNotFound = &response.BodyError{
		Code:   http.StatusNotFound,
		Status: http.StatusText(http.StatusNotFound),
	}

	getStatusBadRequest = &response.BodyError{
		Code:   http.StatusBadRequest,
		Status: http.StatusText(http.StatusBadRequest),
	}

	getStatusInternalServerError = &response.BodyError{
		Code:   http.StatusInternalServerError,
		Status: http.StatusText(http.StatusInternalServerError),
	}
)

// GetErrorCode will return status body with message based on HTTP code
func GetErrorCode(code int, err error) *response.BodyError {
	switch {
	case code == http.StatusNotFound:
		getStatusNotFound.Error = err.Error()
		return getStatusNotFound
	case code == http.StatusBadRequest:
		getStatusBadRequest.Error = err.Error()
		return getStatusBadRequest
	default:
		getStatusInternalServerError.Error = err.Error()
		return getStatusInternalServerError
	}
}
