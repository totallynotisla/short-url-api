package pkg

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApiError struct {
	Code    int
	Message string
	Errors  []string
}

func NewApiErrorNotFound(message string) *ApiError {
	if message == "" {
		message = "Not Found"
	}

	return &ApiError{
		Code:    http.StatusNotFound,
		Message: message,
	}
}

func NewApiErrorBadRequest(message string, errors []string) *ApiError {
	if message == "" {
		message = "Bad Request"
	}

	return &ApiError{
		Code:    http.StatusBadRequest,
		Message: message,
		Errors:  errors,
	}
}

func NewApiErrorUnauthorized(message string) *ApiError {
	if message == "" {
		message = "Unauthorized"
	}

	return &ApiError{
		Code:    http.StatusUnauthorized,
		Message: message,
	}
}

func (ae ApiError) JSON(c *gin.Context) {
	c.JSON(ae.Code, ApiResponse{
		Message: ae.Message,
		Errors:  ae.Errors,
	})
}

func NewApiErrorInternal() *ApiError {
	return &ApiError{
		Code:    http.StatusInternalServerError,
		Message: "Internal Server Error",
	}
}

func NewApiErrorForbidden(message string) *ApiError {
	if message == "" {
		message = "Forbidden"
	}

	return &ApiError{
		Code:    http.StatusForbidden,
		Message: message,
	}
}
