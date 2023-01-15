package objects

type Object int

const (
	Model Object = iota
	List
	TextCompletion
	CodeCompletion
	Edit
	Embedding
	File
	FineTune
	Engine
)

var objectToString = map[Object]string{
	Model:          "model",
	List:           "list",
	TextCompletion: "text_completion",
	CodeCompletion: "code_completion",
	Edit:           "edit",
	Embedding:      "embedding",
	File:           "file",
	FineTune:       "fine-tune",
	Engine:         "engine",
}

var stringToModel = map[string]Object{
	"model":           Model,
	"list":            List,
	"text_completion": TextCompletion,
	"code_completion": CodeCompletion,
	"edit":            Edit,
	"embedding":       Embedding,
	"file":            File,
	"fine-tune":       FineTune,
	"engine":          Engine,
}
