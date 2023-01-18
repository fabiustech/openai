package openai

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"path"

	"github.com/fabiustech/openai/routes"
)

const (
	scheme   = "https"
	host     = "api.openai.com"
	basePath = "v1"
)

// Client is OpenAI GPT-3 API client.
type Client struct {
	token string
	orgID *string

	// scheme and host are only used for testing.
	scheme, host string
}

// NewClient creates new OpenAI API client.
func NewClient(token string) *Client {
	return &Client{
		token:  token,
		scheme: scheme,
		host:   host,
	}
}

// NewClientWithOrg creates new OpenAI API client for specified Organization ID.
func NewClientWithOrg(token, org string) *Client {
	return &Client{
		token:  token,
		orgID:  &org,
		scheme: scheme,
		host:   host,
	}
}

func (c *Client) post(ctx context.Context, path string, payload any) ([]byte, error) {
	var b, err = json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	var req *http.Request
	req, err = http.NewRequestWithContext(ctx, "POST", c.reqURL(path), bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	if c.orgID != nil {
		req.Header.Set("OpenAI-Organization", *c.orgID)
	}

	var resp *http.Response
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err = interpretResponse(resp); err != nil {
		return nil, err
	}

	return io.ReadAll(resp.Body)
}

func (c *Client) postFile(ctx context.Context, fr *FileRequest) ([]byte, error) {
	var b bytes.Buffer
	var w = multipart.NewWriter(&b)

	if err := w.WriteField("purposes", fr.Purpose); err != nil {
		return nil, err
	}

	var fw, err = w.CreateFormFile("file", fr.File.Name())
	if err != nil {
		return nil, err
	}

	if _, err = io.Copy(fw, fr.File); err != nil {
		return nil, err
	}

	if err = w.Close(); err != nil {
		return nil, err
	}

	var req *http.Request
	req, err = http.NewRequestWithContext(ctx, "POST", c.reqURL(routes.Files), &b)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", w.FormDataContentType())

	if c.orgID != nil {
		req.Header.Set("OpenAI-Organization", *c.orgID)
	}

	var resp *http.Response
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err = interpretResponse(resp); err != nil {
		return nil, err
	}

	return io.ReadAll(resp.Body)
}

func (c *Client) get(ctx context.Context, path string) ([]byte, error) {
	var req, err = http.NewRequestWithContext(ctx, "POST", c.reqURL(path), nil)
	if err != nil {
		return nil, err
	}

	if c.orgID != nil {
		req.Header.Set("OpenAI-Organization", *c.orgID)
	}

	var resp *http.Response
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err = interpretResponse(resp); err != nil {
		return nil, err
	}

	return io.ReadAll(resp.Body)
}

func (c *Client) delete(ctx context.Context, path string) ([]byte, error) {
	var req, err = http.NewRequestWithContext(ctx, "DELETE", c.reqURL(path), nil)
	if err != nil {
		return nil, err
	}

	var resp *http.Response
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err = interpretResponse(resp); err != nil {
		return nil, err
	}

	return io.ReadAll(resp.Body)
}

func (c *Client) reqURL(route string) string {
	var u = &url.URL{
		Scheme: c.scheme,
		Host:   c.host,
		Path:   path.Join(basePath, route),
	}
	return u.String()
}

func interpretResponse(resp *http.Response) error {
	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusBadRequest {
		var b, err = io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("error, status code: %d", resp.StatusCode)
		}
		var er *errorResponse
		if err = json.Unmarshal(b, er); err != nil || er.Error == nil {
			return fmt.Errorf("error, status code: %d, msg: %s", resp.StatusCode, string(b))
		}

		return er.Error
	}

	return nil
}
