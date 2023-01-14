package openai

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/fabiustech/openai/models"
	"net/http"
)

// Embedding is a special format of data representation that can be easily utilized by machine
// learning models and algorithms. The embedding is an information dense representation of the
// semantic meaning of a piece of text. Each embedding is a vector of floating point numbers,
// such that the distance between two embeddings in the vector space is correlated with semantic similarity
// between two inputs in the original format. For example, if two texts are similar,
// then their vector representations should also be similar.
type Embedding struct {
	Object    string    `json:"object"`
	Embedding []float64 `json:"embedding"`
	Index     int       `json:"index"`
}

// EmbeddingResponse is the response from a Create embeddings request.
type EmbeddingResponse struct {
	Object string           `json:"object"`
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
func (c *Client) CreateEmbeddings(ctx context.Context, request EmbeddingRequest) (resp EmbeddingResponse, err error) {
	var reqBytes []byte
	reqBytes, err = json.Marshal(request)
	if err != nil {
		return
	}

	urlSuffix := "/embeddings"
	req, err := http.NewRequest(http.MethodPost, c.fullURL(urlSuffix), bytes.NewBuffer(reqBytes))
	if err != nil {
		return
	}

	req = req.WithContext(ctx)
	err = c.sendRequest(req, &resp)

	return
}
