package models

// Moderation represents all models available for use with the Moderations endpoint.
type Moderation int

const (
	// UnknownModeration represents and invalid Moderation model.
	UnknownModeration Moderation = iota
	// TextModerationStable ...
	TextModerationStable
	// TextModerationLatest ...
	TextModerationLatest
)

// String implements the fmt.Stringer interface.
func (m Moderation) String() string {
	return moderationToString[m]
}

// MarshalText implements the encoding.TextMarshaler interface.
func (m Moderation) MarshalText() ([]byte, error) {
	return []byte(m.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
// On unrecognized value, it sets |e| to Unknown.
func (m *Moderation) UnmarshalText(b []byte) error {
	if val, ok := stringToModeration[(string(b))]; ok {
		*m = val
		return nil
	}

	*m = UnknownModeration

	return nil
}

var moderationToString = map[Moderation]string{
	// TextDavinciEdit001 can be used to edit text, rather than just completing it.
	TextModerationStable: "text-moderation-stable",
	// CodeDavinciEdit001 can be used to edit code, rather than just completing it.
	TextModerationLatest: "text-moderation-latest",
}

var stringToModeration = map[string]Moderation{
	"text-moderation-stable": TextModerationStable,
	"text-moderation-latest": TextModerationLatest,
}
