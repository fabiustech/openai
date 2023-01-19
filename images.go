package openai

import (
	"context"
	"encoding/json"

	"github.com/fabiustech/openai/images"
	"github.com/fabiustech/openai/routes"
)

// CreateImageRequest contains all relevant fields for requests to the images/generations endpoint.
type CreateImageRequest struct {
	// Prompt is a text description of the desired image(s). The maximum length is 1000 characters.
	Prompt string `json:"prompt"`
	// N specifies the number of images to generate. Must be between 1 and 10.
	// Defaults to 1.
	N int `json:"n,omitempty"`
	// Size specifies the size of the generated images. Must be one of images.Size256x256, images.Size512x512, or
	// images.Size1024x1024.
	// Defaults to images.Size1024x1024.
	Size images.Size `json:"size,omitempty"`
	// ResponseFormat specifies the format in which the generated images are returned. Must be one of images.FormatURL
	// or images.FormatB64JSON.
	// Defaults to images.FormatURL.
	ResponseFormat images.Format `json:"response_format,omitempty"`
	// User specifies a unique identifier representing your end-user, which can help OpenAI to monitor and detect abuse:
	// https://beta.openai.com/docs/guides/safety-best-practices/end-user-ids.
	User string `json:"user,omitempty"`
}

// EditImageRequest contains all relevant fields for requests to the images/edits endpoint.
type EditImageRequest struct {
	// Image is the image to edit. Must be a valid PNG file, less than 4MB, and square. If Mask is not provided, image
	// must have transparency, which will be used as the mask.
	Image string `json:"image"`
	// Mask is an additional image whose fully transparent areas (e.g. where alpha is zero) indicate where image should
	// be edited. Must be a valid PNG file, less than 4MB, and have the same dimensions as Image.
	Mask string `json:"mask,omitempty"`
	// Prompt is a text description of the desired image(s). The maximum length is 1000 characters.
	Prompt string `json:"prompt"`
	// N specifies the number of images to generate. Must be between 1 and 10.
	// Defaults to 1.
	N int `json:"n,omitempty"`
	// Size specifies the size of the generated images. Must be one of images.Size256x256, images.Size512x512, or
	// images.Size1024x1024.
	// Defaults to images.Size1024x1024.
	Size images.Size `json:"size,omitempty"`
	// ResponseFormat specifies the format in which the generated images are returned. Must be one of images.FormatURL
	// or images.FormatB64JSON.
	// Defaults to images.FormatURL.
	ResponseFormat images.Format `json:"response_format,omitempty"`
	// User specifies a unique identifier representing your end-user, which can help OpenAI to monitor and detect abuse:
	// https://beta.openai.com/docs/guides/safety-best-practices/end-user-ids.
	User string `json:"user,omitempty"`
}

// VariationImageRequest contains all relevant fields for requests to the images/variations endpoint.
type VariationImageRequest struct {
	// Image is the image to use as the basis for the variation(s). Must be a valid PNG file, less than 4MB, and square.
	Image string `json:"image"`
	// N specifies the number of images to generate. Must be between 1 and 10.
	// Defaults to 1.
	N int `json:"n,omitempty"`
	// Size specifies the size of the generated images. Must be one of images.Size256x256, images.Size512x512, or
	// images.Size1024x1024.
	// Defaults to images.Size1024x1024.
	Size images.Size `json:"size,omitempty"`
	// ResponseFormat specifies the format in which the generated images are returned. Must be one of images.FormatURL
	// or images.FormatB64JSON.
	// Defaults to images.FormatURL.
	ResponseFormat images.Format `json:"response_format,omitempty"`
	// User specifies a unique identifier representing your end-user, which can help OpenAI to monitor and detect abuse:
	// https://beta.openai.com/docs/guides/safety-best-practices/end-user-ids.
	User string `json:"user,omitempty"`
}

// ImageResponse represents a response structure for image API.
type ImageResponse struct {
	Created uint64       `json:"created,omitempty"`
	Data    []*ImageData `json:"data,omitempty"`
}

// ImageData represents a response data structure for image API.
// Only one field will be non-nil.
type ImageData struct {
	URL     *string `json:"url,omitempty"`
	B64JSON *string `json:"b64_json,omitempty"`
}

// CreateImage creates an image (or images) given a prompt.
func (c *Client) CreateImage(ctx context.Context, ir *CreateImageRequest) (*ImageResponse, error) {
	var b, err = c.post(ctx, routes.ImageGenerations, ir)
	if err != nil {
		return nil, err
	}

	var resp = &ImageResponse{}
	if err = json.Unmarshal(b, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

// EditImage creates an edited or extended image (or images) given an original image and a prompt.
func (c *Client) EditImage(ctx context.Context, eir *EditImageRequest) (*ImageResponse, error) {
	var b, err = c.post(ctx, routes.ImageEdits, eir)
	if err != nil {
		return nil, err
	}

	var resp = &ImageResponse{}
	if err = json.Unmarshal(b, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

// ImageVariation creates a variation (or variations) of a given image.
func (c *Client) ImageVariation(ctx context.Context, vir *VariationImageRequest) (*ImageResponse, error) {
	var b, err = c.post(ctx, routes.ImageVariations, vir)
	if err != nil {
		return nil, err
	}

	var resp = &ImageResponse{}
	if err = json.Unmarshal(b, resp); err != nil {
		return nil, err
	}

	return resp, nil
}
