package openai

import (
	"context"
	"encoding/json"
)

// Image sizes defined by the OpenAI API.
// TODO: make enum.
const (
	CreateImageSize256x256   = "256x256"
	CreateImageSize512x512   = "512x512"
	CreateImageSize1024x1024 = "1024x1024"
)

const (
	CreateImageResponseFormatURL     = "url"
	CreateImageResponseFormatB64JSON = "b64_json"
)

// ImageRequest represents the request structure for the image API.
type ImageRequest struct {
	Prompt         string `json:"prompt,omitempty"`
	N              int    `json:"n,omitempty"`
	Size           string `json:"size,omitempty"`
	ResponseFormat string `json:"response_format,omitempty"`
	User           string `json:"user,omitempty"`
}

// ImageResponse represents a response structure for image API.
type ImageResponse struct {
	Created uint64                   `json:"created,omitempty"`
	Data    []ImageResponseDataInner `json:"data,omitempty"`
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

// CreateImage - API call to create an image. This is the main endpoint of the DALL-E API.
func (c *Client) CreateImage(ctx context.Context, ir *ImageRequest) (*ImageResponse, error) {
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

// TODO: ImageEdit
// TODO: ImageVariation
