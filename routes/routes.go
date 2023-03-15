// Package routes contains constants for all OpenAI endpoint routes.
package routes

const (
	// Completions is the route for the completions endpoint.
	// https://platform.openai.com/docs/api-reference/completions
	Completions = "completions"
	// ChatCompletions  is the route for the chat completions endpoint.
	// https://platform.openai.com/docs/api-reference/chat
	ChatCompletions = "chat/completions"
	// Edits is the route for the edits endpoint.
	// https://platform.openai.com/docs/api-reference/edits
	Edits = "edits"
	// Embeddings is the route for the embeddings endpoint.
	// https://platform.openai.com/docs/api-reference/embeddings
	Embeddings = "embeddings"

	// Engines is the route for the engines endpoint.
	// https://platform.openai.com/docs/api-reference/engines
	// Deprecated: Use Models instead.
	Engines = "engines"

	// Files is the route for the files endpoint.
	// https://platform.openai.com/docs/api-reference/files
	Files = "files"

	// FineTunes is the route for the fine-tunes endpoint.
	// https://platform.openai.com/docs/api-reference/fine-tunes
	FineTunes = "fines-tunes"

	imagesBase = "images/"

	// ImageGenerations is the route for the create images endpoint.
	// https://platform.openai.com/docs/api-reference/images/create
	ImageGenerations = imagesBase + "generations"
	// ImageEdits is the route for the create image edits endpoint.
	// https://platform.openai.com/docs/api-reference/images/create-edit
	ImageEdits = imagesBase + "edits"
	// ImageVariations is the route for the create image variations endpoint.
	// https://platform.openai.com/docs/api-reference/images/create-variation
	ImageVariations = imagesBase + "variations"

	// Moderations is the route for the moderations endpoint.
	// https://platform.openai.com/docs/api-reference/moderations
	Moderations = "moderations"
)
