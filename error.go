package openai

import (
	"fmt"
	"net/http"
)

// responseError wraps the returned error.
type responseError struct {
	ID  string `json:"id"`
	Err *Error `json:"error,omitempty"`
}

// Error implements the error interface.
func (e *responseError) Error() string {
	return fmt.Sprintf("Request ID: %s, Code: %v, Message: %s, Type: %s, Param: %v", e.ID, e.Err.Code, e.Err.Message, e.Err.Type, e.Err.Param)
}

// Retryable returns true if the error is retryable.
func (e *responseError) Retryable() bool {
	if e.Err.Code >= http.StatusInternalServerError {
		return true
	}
	return e.Err.Code == http.StatusTooManyRequests
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
