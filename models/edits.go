package models

// Edit represents all models available for use with the Edits endpoint.
type Edit int

const (
	TextDavinciEdit001 Edit = iota
	CodeDavinciEdit001
)

var editToString = map[Edit]string{
	// TextDavinciEdit001 can be used to edit text, rather than just completing it.
	TextDavinciEdit001: "text-davinci-edit-001",
	// CodeDavinciEdit001 can be used to edit code, rather than just completing it.
	CodeDavinciEdit001: "code-davinci-edit-001",
}

var stringToEdit = map[string]Edit{
	"text-davinci-edit-001": TextDavinciEdit001,
	"code-davinci-edit-001": CodeDavinciEdit001,
}
