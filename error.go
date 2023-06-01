package openai

import (
	"fmt"
	"net/http"
)

type retryableError struct {
	Err *Error `json:"error"`
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
	if e.Code >= http.StatusInternalServerError || e.Code == 0 {
		return true
	}
	return e.Code == http.StatusTooManyRequests
}
