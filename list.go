package openai

import (
	"github.com/fabiustech/openai/objects"
)

type List[T any] struct {
	Object objects.Object `json:"object"`
	Data   []T            `json:"data"`
}
