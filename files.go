package openai

import (
	"context"
	"encoding/json"
	"github.com/fabiustech/openai/objects"
	"github.com/fabiustech/openai/routes"
	"net/url"
	"path"
)

// FileRequest ...
type FileRequest struct {
	FileName string `json:"file"`
	FilePath string `json:"-"`
	Purpose  string `json:"purpose"`
}

// File struct represents an OpenAPI file.
type File struct {
	Bytes     int            `json:"bytes"`
	CreatedAt int            `json:"created_at"`
	ID        string         `json:"id"`
	FileName  string         `json:"filename"`
	Object    objects.Object `json:"object"`
	Owner     string         `json:"owner"`
	Purpose   string         `json:"purpose"`
}

// FilesList is a list of files that belong to the user or organization.
// TODO: wrap.
type FilesList struct {
	Files []File `json:"data"`
}

// isUrl is a helper function that determines whether the given FilePath
// is a remote URL or a local file path.
func isURL(path string) bool {
	if _, err := url.ParseRequestURI(path); err != nil {
		return false
	}

	if u, err := url.Parse(path); err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	return true
}

// CreateFile uploads a jsonl file to GPT3
// FilePath can be either a local file path or a URL.
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

// DeleteFile deletes an existing file.
func (c *Client) DeleteFile(ctx context.Context, id string) error {
	return c.delete(ctx, path.Join(routes.Files, id))
}

// ListFiles Lists the currently available files,
// and provides basic information about each file such as the file name and purpose.
func (c *Client) ListFiles(ctx context.Context) (*FilesList, error) {
	var b, err = c.get(ctx, routes.Files)
	if err != nil {
		return nil, err
	}

	var fl *FilesList
	if err = json.Unmarshal(b, fl); err != nil {
		return nil, err
	}

	return fl, nil
}

// GetFile Retrieves a file instance, providing basic information about the file
// such as the file name and purpose.
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
