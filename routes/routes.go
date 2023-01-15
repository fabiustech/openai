// Package routes contains constants for all OpenAI endpoint routes.
package routes

const (
	// Completions is the route for the completions endpoint.
	// https://beta.openai.com/docs/api-reference/completions
	Completions = "completions"
	// Edits is the route for the edits endpoint.
	// https://beta.openai.com/docs/api-reference/edits
	Edits = "edits"
	// Embeddings is the route for the embeddings endpoint.
	// https://beta.openai.com/docs/api-reference/embeddings
	Embeddings = "embeddings"

	// Engines is the route for the engines endpoint.
	// https://beta.openai.com/docs/api-reference/engines
	// Deprecated: Use Models instead.
	Engines = "engines"

	// Files is the route for the files endpoint.
	// https://beta.openai.com/docs/api-reference/files
	Files = "files"

	//
	imagesBase = "images/"

	// ImageGenerations is the route for the create images endpoint.
	// https://beta.openai.com/docs/api-reference/images/create
	ImageGenerations = imagesBase + "generations"
	// ImageEdits is the route for the create image edits endpoint.
	// https://beta.openai.com/docs/api-reference/images/create-edit
	ImageEdits = imagesBase + "edits"
	// ImageVariations is the route for the create image variations endpoint.
	// https://beta.openai.com/docs/api-reference/images/create-variation
	ImageVariations = imagesBase + "variations"

	// Moderations is the route for the moderations endpoint.
	// https://beta.openai.com/docs/api-reference/moderations
	Moderations = "moderations"
)
