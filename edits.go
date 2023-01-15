package openai

import (
	"context"
	"encoding/json"
	"github.com/fabiustech/openai/models"
	"github.com/fabiustech/openai/objects"
	"github.com/fabiustech/openai/routes"
)

// EditsRequest represents a request structure for Edits API.
type EditsRequest struct {
	Model models.Edit `json:"model"`
	// Input is the input text to use as a starting point for the edit.
	// Defaults to "".
	Input string `json:"input,omitempty"`
	// Instruction is the instruction that tells the model how to edit the prompt.
	Instruction string `json:"instruction,omitempty"`
	// N specifies how many edits to generate for the input and instruction.
	// Defaults to 1.
	N *int `json:"n,omitempty"`
	// Temperature specifies what sampling temperature to use. Higher values means the model will take more risks. Try 0.9 for more creative
	// applications, and 0 (argmax sampling) for ones with a well-defined answer. OpenAI generally recommends altering
	// this or top_p but not both.
	// More on sampling temperature: https://towardsdatascience.com/how-to-sample-from-language-models-682bceb97277
	// Defaults to 1.
	Temperature *float32 `json:"temperature,omitempty"`
	// TopP specifies an alternative to sampling with temperature, called nucleus sampling, where the model considers
	// the results of the tokens with top_p probability mass. So 0.1 means only the tokens comprising the top 10%
	// probability mass are considered. OpenAI generally recommends altering this or temperature but not both.
	// Defaults to 1.
	TopP *float32 `json:"top_p,omitempty"`
}

// EditsChoice represents one of possible edits.
type EditsChoice struct {
	Text  string `json:"text"`
	Index int    `json:"index"`
}

// EditsResponse represents a response structure for Edits API.
type EditsResponse struct {
	Object  objects.Object `json:"object"` // "edit"
	Created uint64         `json:"created"`
	Usage   *Usage         `json:"usage"`
	Choices []*EditsChoice `json:"choices"`
}

// Edits ...
func (c *Client) Edits(ctx context.Context, er *EditsRequest) (*EditsResponse, error) {
	var b, err = c.post(ctx, routes.Edits, er)

	var resp *EditsResponse
	if err = json.Unmarshal(b, resp); err != nil {
		return nil, err
	}

	return resp, nil
}
