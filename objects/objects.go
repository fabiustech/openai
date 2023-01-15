package objects

type Object int

const (
	Model = "model"
	List  = "list"

	TextCompletion = "text_completion"
	CodeCompletion = "code_completion"
	Edit           = "edit"
	Embedding      = "embedding"
	File           = "file"
	FineTune       = "fine-tune"
	Engine         = "engine"
)
