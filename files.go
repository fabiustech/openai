package openai

import (
	"context"
	"encoding/json"
	"os"
	"path"

	"github.com/fabiustech/openai/objects"
	"github.com/fabiustech/openai/routes"
)

// FileRequest contains all relevant data for upload requests to the files endpoint.
type FileRequest struct {
	// File is the JSON Lines file to be uploaded. If the purpose is set to "fine-tune", each line is a JSON record
	// with "prompt" and "completion" fields representing your training examples:
	// https://beta.openai.com/docs/guides/fine-tuning/prepare-training-data.
	File *os.File
	// Purpose is the intended purpose of the uploaded documents. Use "fine-tune" for Fine-tuning.
	// This allows OpenAI to validate the format of the uploaded file.
	Purpose string
}

// NewFineTuneFileRequest returns a |*FileRequest| with File opened from |path| and Purpose set to "fine-tuned".
func NewFineTuneFileRequest(path string) (*FileRequest, error) {
	var f, err = os.Open(path)
	if err != nil {
		return nil, err
	}

	return &FileRequest{
		File:    f,
		Purpose: "fine-tune",
	}, nil
}

// File represents an OpenAPI file.
type File struct {
	ID        string         `json:"id"`
	Object    objects.Object `json:"object"`
	Bytes     int            `json:"bytes"`
	CreatedAt int            `json:"created_at"`
	Filename  string         `json:"filename"`
	Purpose   string         `json:"purpose"`
}

// ListFiles returns a list of files that belong to the user's organization.
func (c *Client) ListFiles(ctx context.Context) (*List[*File], error) {
	var b, err = c.get(ctx, routes.Files)
	if err != nil {
		return nil, err
	}

	var fl = &List[*File]{}
	if err = json.Unmarshal(b, fl); err != nil {
		return nil, err
	}

	return fl, nil
}

// UploadFile uploads a file that contains document(s) to be used across various endpoints/features. Currently, the size
// of all the files uploaded by one organization can be up to 1 GB.
func (c *Client) UploadFile(ctx context.Context, fr *FileRequest) (*File, error) {
	var b, err = c.postFile(ctx, fr)
	if err != nil {
		return nil, err
	}

	var f = &File{}
	if err = json.Unmarshal(b, f); err != nil {
		return nil, err
	}

	return f, nil
}

// DeleteFile deletes a file.
func (c *Client) DeleteFile(ctx context.Context, id string) error {
	var _, err = c.delete(ctx, path.Join(routes.Files, id))

	return err
}

// RetrieveFile returns information about a specific file.
func (c *Client) RetrieveFile(ctx context.Context, id string) (*File, error) {
	var b, err = c.get(ctx, path.Join(routes.Files, id))
	if err != nil {
		return nil, err
	}

	var f = &File{}
	if err = json.Unmarshal(b, f); err != nil {
		return nil, err
	}

	return f, nil
}
