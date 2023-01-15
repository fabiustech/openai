package openai

import (
	"context"
	"encoding/json"
	"github.com/fabiustech/openai/models"
	"github.com/fabiustech/openai/objects"
	"github.com/fabiustech/openai/routes"
)

// Embedding is a special format of data representation that can be easily utilized by machine
// learning models and algorithms. The embedding is an information dense representation of the
// semantic meaning of a piece of text. Each embedding is a vector of floating point numbers,
// such that the distance between two embeddings in the vector space is correlated with semantic similarity
// between two inputs in the original format. For example, if two texts are similar,
// then their vector representations should also be similar.
type Embedding struct {
	Object    objects.Object `json:"object"`
	Embedding []float64      `json:"embedding"`
	Index     int            `json:"index"`
}

// EmbeddingResponse is the response from a Create embeddings request.
// Todo: Wrap
type EmbeddingResponse struct {
	Object objects.Object   `json:"object"`
	Data   []Embedding      `json:"data"`
	Model  models.Embedding `json:"model"`
	Usage  Usage            `json:"usage"`
}

// EmbeddingRequest is the input to a Create embeddings request.
type EmbeddingRequest struct {
	// Input is a slice of strings for which you want to generate an Embedding vector.
	// Each input must not exceed 2048 tokens in length.
	// OpenAPI suggests replacing newlines (\n) in your input with a single space, as they
	// have observed inferior results when newlines are present.
	// E.g.
	//	"The food was delicious and the waiter..."
	Input []string `json:"input"`
	// ID of the model to use. You can use the List models API to see all of your available models,
	// or see our Model overview for descriptions of them.
	Model models.Embedding `json:"model"`
	// A unique identifier representing your end-user, which will help OpenAI to monitor and detect abuse.
	User string `json:"user"`
}

// CreateEmbeddings returns an EmbeddingResponse which will contain an Embedding for every item in |request.Input|.
// https://beta.openai.com/docs/api-reference/embeddings/create
func (c *Client) CreateEmbeddings(ctx context.Context, request *EmbeddingRequest) (*EmbeddingResponse, error) {
	var b, err = c.post(ctx, routes.Embeddings, request)
	if err != nil {
		return nil, err
	}

	var resp *EmbeddingResponse
	if err = json.Unmarshal(b, resp); err != nil {
		return nil, err
	}

	return resp, nil
}
