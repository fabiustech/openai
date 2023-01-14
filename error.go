package openai

import (
	"fmt"
	"net/http"
)

type ErrorResponse struct {
	Error *Error `json:"error,omitempty"`
}

type Error struct {
	Code    int     `json:"code,omitempty"`
	Message string  `json:"message"`
	Param   *string `json:"param,omitempty"`
	Type    string  `json:"type"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("Code: %v, Message: %s, Type: %s, Param: %v", e.Code, e.Message, e.Type, e.Param)
}

func (e *Error) Retryable() bool {
	if e.Code >= http.StatusInternalServerError {
		return true
	}
	return e.Code == http.StatusTooManyRequests
}
