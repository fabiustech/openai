package openai

import (
	"context"
	"encoding/json"

	"github.com/fabiustech/openai/models"

	"github.com/fabiustech/openai/routes"
)

// ModerationRequest contains all relevant fields for requests to the moderations endpoint.
type ModerationRequest struct {
	// Input is the input text to classify.
	Input string `json:"input,omitempty"`
	// Model specifies the model to use for moderation.
	// Defaults to models.TextModerationLatest.
	Model models.Moderation `json:"model,omitempty"`
}

// Result represents one of possible moderation results.
type Result struct {
	Categories     *ResultCategories     `json:"categories"`
	CategoryScores *ResultCategoryScores `json:"category_scores"`
	Flagged        bool                  `json:"flagged"`
}

// ResultCategories represents Categories of Result.
type ResultCategories struct {
	Hate            bool `json:"hate"`
	HateThreatening bool `json:"hate/threatening"`
	SelfHarm        bool `json:"self-harm"`
	Sexual          bool `json:"sexual"`
	SexualMinors    bool `json:"sexual/minors"`
	Violence        bool `json:"violence"`
	ViolenceGraphic bool `json:"violence/graphic"`
}

// ResultCategoryScores represents CategoryScores of Result.
type ResultCategoryScores struct {
	Hate            float32 `json:"hate"`
	HateThreatening float32 `json:"hate/threatening"`
	SelfHarm        float32 `json:"self-harm"`
	Sexual          float32 `json:"sexual"`
	SexualMinors    float32 `json:"sexual/minors"`
	Violence        float32 `json:"violence"`
	ViolenceGraphic float32 `json:"violence/graphic"`
}

// ModerationResponse represents a response structure for moderation API.
type ModerationResponse struct {
	ID      string   `json:"id"`
	Model   string   `json:"model"`
	Results []Result `json:"results"`
}

// Moderations ...
func (c *Client) Moderations(ctx context.Context, mr *ModerationRequest) (*ModerationResponse, error) {
	var b, err = c.post(ctx, routes.Moderations, mr)
	if err != nil {
		return nil, err
	}

	var resp *ModerationResponse
	if err = json.Unmarshal(b, resp); err != nil {
		return nil, err
	}

	return resp, nil
}
