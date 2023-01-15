package images

// Format represents the enum values for the formats in which
// generated images are returned.
type Format int

const (
	// FormatInvalid represents and invalid Format option.
	FormatInvalid Format = iota
	// FormatURL specifies that the API will return a url to the generated image.
	// URLs will expire after an hour.
	FormatURL
	// FormatB64JSON specifies that the API will return the image as Base64 data.
	FormatB64JSON
)

// String implements the fmt.Stringer interface.
func (f Format) String() string {
	return formatToString[f]
}

// MarshalText implements the encoding.TextMarshaler interface.
func (f Format) MarshalText() ([]byte, error) {
	return []byte(f.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
// On unrecognized value, it sets |e| to Unknown.
func (f *Format) UnmarshalText(b []byte) error {
	if val, ok := stringToFormat[(string(b))]; ok {
		*f = val
		return nil
	}

	*f = FormatInvalid

	return nil
}

var formatToString = map[Format]string{
	FormatURL:     "url",
	FormatB64JSON: "b64_json",
}

var stringToFormat = map[string]Format{
	"url":      FormatURL,
	"b64_json": FormatB64JSON,
}
