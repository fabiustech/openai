package openai

import (
	"context"
	"encoding/json"
	"github.com/fabiustech/openai/images"
)

// CreateImageRequest represents the request structure for the image API.
type CreateImageRequest struct {
	Prompt string `json:"prompt"`
	*ImageRequestFields
}

type EditImageRequest struct {
	Image  string  `json:"image"`
	Mask   *string `json:"mask,omitempty"`
	Prompt string  `json:"prompt"`
	*ImageRequestFields
}

type VariationImageRequest struct {
	Image string `json:"image"`
	*ImageRequestFields
}

type ImageRequestFields struct {
	N              *int           `json:"n,omitempty"`
	Size           *images.Size   `json:"size,omitempty"`
	ResponseFormat *images.Format `json:"response_format,omitempty"`
	User           *string        `json:"user,omitempty"`
}

// ImageResponse represents a response structure for image API.
type ImageResponse struct {
	Created uint64                    `json:"created,omitempty"`
	Data    []*ImageResponseDataInner `json:"data,omitempty"`
}

// ImageResponseDataInner represents a response data structure for image API.
type ImageResponseDataInner struct {
	URL     string `json:"url,omitempty"`
	B64JSON string `json:"b64_json,omitempty"`
}

const (
	routeGenerations = "images/generations"
	routeEdits       = "images/edits"
	routeVariations  = "images/variations"
)

// CreateImage ...
func (c *Client) CreateImage(ctx context.Context, ir *CreateImageRequest) (*ImageResponse, error) {
	var b, err = c.post(ctx, routeGenerations, ir)
	if err != nil {
		return nil, err
	}

	var resp *ImageResponse
	if err = json.Unmarshal(b, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

// EditImage ...
func (c *Client) EditImage(ctx context.Context, eir *EditImageRequest) (*ImageResponse, error) {
	var b, err = c.post(ctx, routeEdits, eir)
	if err != nil {
		return nil, err
	}

	var resp *ImageResponse
	if err = json.Unmarshal(b, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) ImageVariation(ctx context.Context, vir *VariationImageRequest) (*ImageResponse, error) {
	var b, err = c.post(ctx, routeVariations, vir)
	if err != nil {
		return nil, err
	}

	var resp *ImageResponse
	if err = json.Unmarshal(b, resp); err != nil {
		return nil, err
	}

	return resp, nil
}
