package openai

import (
	"context"
	"encoding/json"
	"path"

	"github.com/fabiustech/openai/routes"
)

// Engine contains all relevant fields for requests to the engines endpoint.
type Engine struct {
	ID     string `json:"id"`
	Object string `json:"object"`
	Owner  string `json:"owner"`
	Ready  bool   `json:"ready"`
}

// ListEngines lists the currently available engines, and provides basic
// information about each option such as the owner and availability.
//
// Deprecated: Please use their replacement, Models, instead.
// https://beta.openai.com/docs/api-reference/models
func (c *Client) ListEngines(ctx context.Context) (*List[*Engine], error) {
	var b, err = c.get(ctx, routes.Engines)
	if err != nil {
		return nil, err
	}

	var el = &List[*Engine]{}
	if err = json.Unmarshal(b, el); err != nil {
		return nil, err
	}

	return el, nil
}

// GetEngine retrieves a model instance, providing basic information about it such as the owner and availability.
//
// Deprecated: Please use their replacement, Models, instead.
// https://beta.openai.com/docs/api-reference/models
func (c *Client) GetEngine(ctx context.Context, id string) (*Engine, error) {
	var b, err = c.get(ctx, path.Join(routes.Engines, id))
	if err != nil {
		return nil, err
	}

	var e = &Engine{}
	if err = json.Unmarshal(b, e); err != nil {
		return nil, err
	}

	return e, nil
}
