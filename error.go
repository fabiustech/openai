package openai

import (
	"fmt"
	"net/http"
)

// errorResponse wraps the returned error.
type errorResponse struct {
	Error *Error `json:"error,omitempty"`
}

// Error represents an error response from the API.
type Error struct {
	Code    int     `json:"code"`
	Message string  `json:"message"`
	Param   *string `json:"param,omitempty"`
	Type    string  `json:"type"`
}

// Error implements the error interface.
func (e *Error) Error() string {
	return fmt.Sprintf("Code: %v, Message: %s, Type: %s, Param: %v", e.Code, e.Message, e.Type, e.Param)
}

// Retryable returns true if the error is retryable.
func (e *Error) Retryable() bool {
	if e.Code >= http.StatusInternalServerError {
		return true
	}
	return e.Code == http.StatusTooManyRequests
}
