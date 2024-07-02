package rest_err

import "net/http"

// RestError represents the error object.
// @Summary Error information
// @Description Structure for describing why the error occurred
type RestError struct {
	// Error message.
	Message string `json:"message" example:"error trying to process request"`

	// Error description.
	Err string `json:"error" example:"internal_server_error"`

	// Error code.
	Code int `json:"code" example:"500"`

	// Error causes.
	Causes []Cause `json:"causes"`
}

// Cause represents the error cause.
// @Summary Error Causes
// @Description Structure representing the causes of an error.
type Cause struct {
	// Field associated with the error cause.
	// @json
	// @jsonTag field
	Field string `json:"field" example:"name"`

	// Error message describing the cause.
	// @json
	// @jsonTag message
	Message string `json:"message" example:"name is required"`
}

func (r *RestError) Error() string {
	return r.Message
}

func NewBadRequestValidationError(message string, causes []Cause) *RestError {
	return &RestError{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

func NewBadRequestError(message string) *RestError {
	return &RestError{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
	}
}

func NewUnauthorizedRequestError(message string) *RestError {
	return &RestError{
		Message: message,
		Err:     "unauthorized",
		Code:    http.StatusUnauthorized,
	}
}

func NewForbiddenError(message string) *RestError {
	return &RestError{
		Message: message,
		Err:     "forbidden",
		Code:    http.StatusForbidden,
	}
}

func NewNotFoundError(message string) *RestError {
	return &RestError{
		Message: message,
		Err:     "not_found",
		Code:    http.StatusNotFound,
	}
}

func NewInternalServerError(message string) *RestError {
	return &RestError{
		Message: message,
		Err:     "internal_server_error",
		Code:    http.StatusInternalServerError,
	}
}
