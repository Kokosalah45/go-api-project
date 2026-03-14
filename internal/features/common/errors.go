package common

import "net/http"

type GenericError struct {
	ErrorCode string
	Message   string
}

const (
	GENERIC_ERROR_CODE    = "GENERIC_ERROR"
	GENERIC_ERROR_MESSAGE = "AN ERROR OCCURRED"
)

func NewGenericError(message string, errorCode string) *GenericError {
	currentMessage := GENERIC_ERROR_MESSAGE
	currentErrorCode := GENERIC_ERROR_CODE

	if message != "" {
		currentMessage = message
	}

	if errorCode != "" {
		currentErrorCode = errorCode
	}

	return &GenericError{
		Message:   currentMessage,
		ErrorCode: currentErrorCode,
	}
}

func (e *GenericError) Error() string {
	return e.Message
}

type BadRequestError struct {
	Message    string
	StatusCode int
}

func NewBadRequestError() *BadRequestError {
	currentMessage := "BAD_REQUEST"
	statusCode := http.StatusBadRequest

	return &BadRequestError{
		Message:    currentMessage,
		StatusCode: statusCode,
	}
}
