package rest_errors

import "net/http"

type RestError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewBadRequestRestError(message string) *RestError {
	return &RestError{
		Code:    http.StatusBadRequest,
		Message: message,
	}
}

func NewInternalServerRestError(message string) *RestError {
	return &RestError{
		Code:    http.StatusInternalServerError,
		Message: message,
	}
}
