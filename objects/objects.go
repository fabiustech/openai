// Package objects contains the enum values which represent the various
// objects returned by all OpenAI endpoints.
package objects

// Object enumerates the various object types returned by OpenAI endpoints.
type Object int

const (
	// Unknown is an invalid object.
	Unknown Object = iota
	// Model is a model (can be either a base model or fine-tuned).
	Model
	// List is a list of other objects.
	List
	// TextCompletion is a text completion.
	TextCompletion
	// CodeCompletion is a code completion.
	CodeCompletion
	// Edit is an edit.
	Edit
	// Embedding is an embedding.
	Embedding
	// File is a file.
	File
	// FineTune is a fine-tuned model.
	FineTune
	FineTimeEvent
	// Engine represents an engine.
	// Deprecated: use Model instead.
	Engine
)

// String implements the fmt.Stringer interface.
func (o Object) String() string {
	return objectToString[o]
}

// MarshalText implements the encoding.TextMarshaler interface.
func (o Object) MarshalText() ([]byte, error) {
	return []byte(o.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
// On unrecognized value, it sets |e| to Unknown.
func (o *Object) UnmarshalText(b []byte) error {
	if val, ok := stringToObject[(string(b))]; ok {
		*o = val
		return nil
	}

	*o = Unknown

	return nil
}

var objectToString = map[Object]string{
	Model:          "model",
	List:           "list",
	TextCompletion: "text_completion",
	CodeCompletion: "code_completion",
	Edit:           "edit",
	Embedding:      "embedding",
	File:           "file",
	FineTune:       "fine-tune",
	FineTimeEvent:  "fine-tune-event",
	Engine:         "engine",
}

var stringToObject = map[string]Object{
	"model":           Model,
	"list":            List,
	"text_completion": TextCompletion,
	"code_completion": CodeCompletion,
	"edit":            Edit,
	"embedding":       Embedding,
	"file":            File,
	"fine-tune":       FineTune,
	"fine-tune-event": FineTimeEvent,
	"engine":          Engine,
}
