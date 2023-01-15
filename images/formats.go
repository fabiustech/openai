package images

type Format int

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

	*f = FormatUnkown

	return nil
}

const (
	FormatUnkown Format = iota
	FormatURL
	FormatB64JSON
)

var formatToString = map[Format]string{
	FormatURL:     "url",
	FormatB64JSON: "b64_json",
}

var stringToFormat = map[string]Format{
	"url":      FormatURL,
	"b64_json": FormatB64JSON,
}
