package openai

import (
	"github.com/fabiustech/openai/objects"
)

// List represents a generic form of list of objects returned from many get endpoints.
type List[T any] struct {
	// Object specifies the object type (e.g. Model).
	Object objects.Object `json:"object"`
	// Data contains the list of objects.
	Data []T `json:"data"`
}
