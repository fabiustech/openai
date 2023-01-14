package openai

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

type AnswerRequest struct {
	Documents       []string   `json:"documents,omitempty"`
	File            string     `json:"file,omitempty"`
	Question        string     `json:"question"`
	SearchModel     string     `json:"search_model,omitempty"`
	Model           string     `json:"model"`
	ExamplesContext string     `json:"examples_context"`
	Examples        [][]string `json:"examples"`
	MaxTokens       int        `json:"max_tokens,omitempty"`
	Stop            []string   `json:"stop,omitempty"`
	Temperature     *float64   `json:"temperature,omitempty"`
}

type AnswerResponse struct {
	Answers           []string `json:"answers"`
	Completion        string   `json:"completion"`
	Model             string   `json:"model"`
	Object            string   `json:"object"`
	SearchModel       string   `json:"search_model"`
	SelectedDocuments []struct {
		Document int    `json:"document"`
		Text     string `json:"text"`
	} `json:"selected_documents"`
}

// Answers ...
func (c *Client) Answers(ctx context.Context, ar AnswerRequest) (*AnswerResponse, error) {
	var b, err = json.Marshal(ar)
	if err != nil {
		return nil, err
	}

	var req *http.Request
	req, err = http.NewRequest("POST", c.fullURL("/answers"), bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)

	var resp *AnswerResponse
	if err = c.sendRequest(req, resp); err != nil {
		return nil, err
	}
	
	return resp, err
}
