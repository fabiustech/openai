package openai

import (
	"fmt"
	"net/http"
)

type wrappedError struct {
	Err *Error `json:"error"`
}

// Error represents an error response from the API.
type Error struct {
	StatusCode int     `json:"statusCode"`
	Code       string  `json:"code"`
	Message    string  `json:"message"`
	Param      *string `json:"param,omitempty"`
	Type       string  `json:"type"`
}

// Error implements the error interface.
func (e *Error) Error() string {
	return fmt.Sprintf("Code: %v, Message: %s, Type: %s, Param: %v", e.Code, e.Message, e.Type, e.Param)
}

// Retryable returns true if the error is retryable.
func (e *Error) Retryable() bool {
	if e.StatusCode >= http.StatusInternalServerError {
		return true
	}
	return e.StatusCode == http.StatusTooManyRequests
}
