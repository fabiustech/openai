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
	"os"
	"path"
	"strings"
)

const (
	scheme   = "https"
	host     = "api.openai.com"
	basePath = "vi"
)

func reqURL(route string) string {
	var u = &url.URL{
		Scheme: scheme,
		Host:   host,
		Path:   path.Join(basePath, route),
	}
	return u.String()
}

// Client is OpenAI GPT-3 API client.
type Client struct {
	token string
	orgID *string
}

// NewClient creates new OpenAI API client.
func NewClient(token string) *Client {
	return &Client{
		token: token,
	}
}

// NewClientWithOrg creates new OpenAI API client for specified Organization ID.
func NewClientWithOrg(token, org string) *Client {
	return &Client{
		token: token,
		orgID: &org,
	}
}

func (c *Client) post(ctx context.Context, path string, payload any) ([]byte, error) {
	var b, err = json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	var req *http.Request
	req, err = http.NewRequestWithContext(ctx, "POST", reqURL(path), bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}

	switch payload.(type) {
	case FileRequest:
		req.Header.Set("Content-Type", "") // TODO
	default:
		req.Header.Set("Content-Type", "application/json; charset=utf-8")
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

// TODO: improve this.
func (c *Client) postFile(ctx context.Context, fr *FileRequest) ([]byte, error) {
	var b bytes.Buffer
	var w = multipart.NewWriter(&b)

	var pw, err = w.CreateFormField("purpose")
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(pw, strings.NewReader(fr.Purpose))
	if err != nil {
		return nil, err
	}

	var fw io.Writer
	fw, err = w.CreateFormFile("file", fr.FileName)
	if err != nil {
		return nil, err
	}

	var file io.ReadCloser
	file, err = readFile(fr.FilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	if _, err = io.Copy(fw, file); err != nil {
		return nil, err
	}

	w.Close()

	var req *http.Request
	req, err = http.NewRequestWithContext(ctx, "POST", reqURL(routeFiles), &b)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", w.FormDataContentType())

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

func readFile(path string) (io.ReadCloser, error) {
	if !isURL(path) {
		return os.Open(path)
	}

	var resp, err = http.Get(path)
	if err != nil {
		return nil, err
	}

	// Check server response.
	if resp.StatusCode != http.StatusOK {
		_ = resp.Body.Close()
		return nil, fmt.Errorf("error, status code: %d, message: failed to fetch file", resp.StatusCode)
	}

	return resp.Body, nil
}

func (c *Client) get(ctx context.Context, path string) ([]byte, error) {
	var req, err = http.NewRequestWithContext(ctx, "POST", reqURL(path), nil)
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

func (c *Client) delete(ctx context.Context, path string) error {
	var req, err = http.NewRequestWithContext(ctx, "DELETE", reqURL(path), nil)
	if err != nil {
		return err
	}

	var resp *http.Response
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return interpretResponse(resp)
}

// TODO: implement.
func interpretResponse(resp *http.Response) error {
	return nil
}
