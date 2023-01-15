package models

// Edit represents all models available for use with the Edits endpoint.
type Edit int

const (
	// UnknownEdit represents and invalid Edit model.
	UnknownEdit Edit = iota
	// TextDavinciEdit001 ...
	TextDavinciEdit001
	// CodeDavinciEdit001 ...
	CodeDavinciEdit001
)

// String implements the fmt.Stringer interface.
func (e Edit) String() string {
	return editToString[e]
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e Edit) MarshalText() ([]byte, error) {
	return []byte(e.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
// On unrecognized value, it sets |e| to Unknown.
func (e *Edit) UnmarshalText(b []byte) error {
	if val, ok := stringToEdit[(string(b))]; ok {
		*e = val
		return nil
	}

	*e = UnknownEdit

	return nil
}

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
