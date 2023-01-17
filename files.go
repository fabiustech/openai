package openai

import (
	"context"
	"encoding/json"
	"path"

	"github.com/fabiustech/openai/objects"
	"github.com/fabiustech/openai/routes"
)

// FileRequest ...
type FileRequest struct {
	// File is the JSON Lines file to be uploaded. If the purpose is set to "fine-tune", each line is a JSON record
	// with "prompt" and "completion" fields representing your training examples:
	// https://beta.openai.com/docs/guides/fine-tuning/prepare-training-data.
	File     string `json:"file"`
	FilePath string `json:"-"`
	// The intended purpose of the uploaded documents. Use "fine-tune" for Fine-tuning.
	// This allows OpenAI to validate the format of the uploaded file.
	Purpose string `json:"purpose"`
}

// File struct represents an OpenAPI file.
type File struct {
	ID        string         `json:"id"`
	Object    objects.Object `json:"object"`
	Bytes     int            `json:"bytes"`
	CreatedAt int            `json:"created_at"`
	Filename  string         `json:"filename"`
	Purpose   string         `json:"purpose"`
}

// TODO: FileRequest should take a file.File.
// CreateFile ...
func (c *Client) CreateFile(ctx context.Context, fr *FileRequest) (*File, error) {
	var b, err = c.postFile(ctx, fr)
	if err != nil {
		return nil, err
	}

	var f *File
	if err = json.Unmarshal(b, f); err != nil {
		return nil, err
	}

	return f, nil
}

// DeleteFile ...
func (c *Client) DeleteFile(ctx context.Context, id string) error {
	var _, err = c.delete(ctx, path.Join(routes.Files, id))

	return err
}

// ListFiles ...
func (c *Client) ListFiles(ctx context.Context) (*List[*File], error) {
	var b, err = c.get(ctx, routes.Files)
	if err != nil {
		return nil, err
	}

	var fl *List[*File]
	if err = json.Unmarshal(b, fl); err != nil {
		return nil, err
	}

	return fl, nil
}

// GetFile ...
func (c *Client) GetFile(ctx context.Context, id string) (*File, error) {
	var b, err = c.get(ctx, path.Join(routes.Files, id))
	if err != nil {
		return nil, err
	}

	var f *File
	if err = json.Unmarshal(b, f); err != nil {
		return nil, err
	}

	return f, nil
}
