package openai

import (
	"context"
	"encoding/json"
	"github.com/fabiustech/openai/routes"
	"path"
)

// Engine struct represents engine from OpenAPI API.
type Engine struct {
	ID     string `json:"id"`
	Object string `json:"object"`
	Owner  string `json:"owner"`
	Ready  bool   `json:"ready"`
}

// EnginesList is a list of engines.
type EnginesList struct {
	Engines []*Engine `json:"data"`
}

// ListEngines Lists the currently available engines, and provides basic
// information about each option such as the owner and availability.
//
// Deprecated: Please use their replacement, Models, instead.
// https://beta.openai.com/docs/api-reference/models
func (c *Client) ListEngines(ctx context.Context) (*EnginesList, error) {
	var b, err = c.get(ctx, routes.Engines)
	if err != nil {
		return nil, err
	}

	var el *EnginesList
	if err = json.Unmarshal(b, el); err != nil {
		return nil, err
	}

	return el, nil
}

// GetEngine Retrieves an engine instance, providing basic information about
// the engine such as the owner and availability.
//
// Deprecated: Please use their replacement, Models, instead.
// https://beta.openai.com/docs/api-reference/models
func (c *Client) GetEngine(ctx context.Context, id string) (*Engine, error) {
	var b, err = c.get(ctx, path.Join(routes.Engines, id))
	if err != nil {
		return nil, err
	}

	var e *Engine
	if err = json.Unmarshal(b, e); err != nil {
		return nil, err
	}

	return e, nil
}
