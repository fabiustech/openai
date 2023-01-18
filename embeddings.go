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
type EmbeddingResponse struct {
	*List[*Embedding]
	Model models.Embedding
	Usage *Usage
}

// EmbeddingRequest contains all relevant fields for requests to the embeddings endpoint.
type EmbeddingRequest struct {
	// Input represents input text to get embeddings for, encoded as a strings. To get embeddings for multiple inputs in
	//a single request, pass a slice of length > 1. Each input string must not exceed 8192 tokens in length.
	Input []string `json:"input"`
	// Model is the ID of the model to use.
	Model models.Embedding `json:"model"`
	// User is a unique identifier representing your end-user, which will help OpenAI to monitor and detect abuse.
	User string `json:"user"`
}

// CreateEmbeddings creates an embedding vector representing the input text.
func (c *Client) CreateEmbeddings(ctx context.Context, request *EmbeddingRequest) (*EmbeddingResponse, error) {
	var b, err = c.post(ctx, routes.Embeddings, request)
	if err != nil {
		return nil, err
	}

	var resp = &EmbeddingResponse{}
	if err = json.Unmarshal(b, resp); err != nil {
		return nil, err
	}

	return resp, nil
}
