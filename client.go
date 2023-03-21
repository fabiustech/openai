package openai

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"path"

	"github.com/fabiustech/openai/routes"
)

const (
	// Below are the defaults for calling OpenAI.
	scheme   = "https"
	host     = "api.openai.com"
	basePath = "v1"
)

// Client is OpenAI API client.
type Client struct {
	token string
	orgID *string

	scheme, host, base, params string
}

// SetBaseURL configures the client to make requests to a different base URL.
// If configuring a call to an Azure hosted endpoint, include the `api-version` parameter
// in the passed URL string. E.g.
//
//	https://{your-resource-name}.openai.azure.com/openai/deployments/{deployment-id}/?api-version=2022-12-01
func (c *Client) SetBaseURL(u string) error {
	var parsed, err = url.Parse(u)
	if err != nil {
		return err
	}

	c.scheme = parsed.Scheme
	c.host = parsed.Host
	c.base = parsed.Path
	c.params = parsed.RawQuery

	return nil
}

// NewClient creates new OpenAI API client.
func NewClient(token string) *Client {
	return &Client{
		token:  token,
		scheme: scheme,
		host:   host,
		base:   basePath,
	}
}

// NewClientWithOrg creates new OpenAI API client for specified Organization ID.
func NewClientWithOrg(token, org string) *Client {
	return &Client{
		token:  token,
		orgID:  &org,
		scheme: scheme,
		host:   host,
		base:   basePath,
	}
}

func (c *Client) newRequest(ctx context.Context, method string, url string, body io.Reader) (*http.Request, error) {
	var req, err = http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json; charset=utf-8")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))

	if c.orgID != nil {
		req.Header.Set("OpenAI-Organization", *c.orgID)
	}

	return req, nil
}

func (c *Client) post(ctx context.Context, path string, payload any) ([]byte, error) {
	var b, err = json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	var req *http.Request
	req, err = c.newRequest(ctx, "POST", c.reqURL(path), bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

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

const bufferSize = 4096

func (c *Client) postStream(ctx context.Context, path string, payload any) (<-chan []byte, <-chan error, error) {
	var b, err = json.Marshal(payload)
	if err != nil {
		return nil, nil, err
	}

	var req *http.Request
	req, err = c.newRequest(ctx, "POST", c.reqURL(path), bytes.NewBuffer(b))
	if err != nil {
		return nil, nil, err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "text/event-stream; charset=utf-8")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cache-Control", "no-cache")

	var resp *http.Response
	resp, err = http.DefaultClient.Do(req) //nolint:bodyclose // The body is closed in the error check or the go routine.
	if err != nil {
		return nil, nil, err
	}
	if err = interpretResponse(resp); err != nil {
		_ = resp.Body.Close()
		return nil, nil, err
	}

	var events = make(chan []byte)
	var errCh = make(chan error)

	go func() {
		defer resp.Body.Close()
		defer close(events)
		defer close(errCh)

		for {
			var msg = make([]byte, bufferSize)
			_, err = resp.Body.Read(msg)

			switch {
			case errors.Is(err, io.EOF):
				return
			case err != nil:
				errCh <- err
				return
			case ctx.Err() != nil:
				errCh <- ctx.Err()
				return
			default:
				// No-op.
			}

			events <- msg
		}
	}()

	return events, errCh, nil
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
	req, err = c.newRequest(ctx, "POST", c.reqURL(routes.Files), &b)
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

func (c *Client) get(ctx context.Context, path string) ([]byte, error) {
	var req, err = c.newRequest(ctx, "GET", c.reqURL(path), nil)
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

func (c *Client) delete(ctx context.Context, path string) ([]byte, error) {
	var req, err = c.newRequest(ctx, "DELETE", c.reqURL(path), nil)
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
		Scheme:   c.scheme,
		Host:     c.host,
		Path:     path.Join(c.base, route),
		RawQuery: c.params,
	}

	return u.String()
}

func interpretResponse(resp *http.Response) error {
	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusBadRequest {
		var b, err = io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("error, HTTP status code: %d", resp.StatusCode)
		}
		var er = &errorResponse{}
		if err = json.Unmarshal(b, er); err != nil || er.Err == nil {
			return fmt.Errorf("error, HTTP status code: %d, msg: %s", resp.StatusCode, string(b))
		}

		return er
	}

	return nil
}
