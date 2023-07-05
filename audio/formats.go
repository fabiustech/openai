// Package audio contains the enum values which represent the output formats returned by the
// OpenAI transcription endpoint.
package audio

// Format represents the enum values for the formats in which
// generated transcript are returned.
type Format int

const (
	// FormatInvalid represents and invalid Format option.
	FormatInvalid Format = iota
	// FormatJSON specifies that the API will return the transcript as JSON.
	FormatJSON
	// FormatText specifies that the API will return the transcript as plain text.
	FormatText
	// FormatSRT specifies that the API will return the transcript as SRT.
	FormatSRT
	// FormatVerboseJSON specifies that the API will return the transcript as verbose JSON.
	FormatVerboseJSON
	// FormatVTT specifies that the API will return the transcript as VTT.
	FormatVTT
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
	FormatJSON:        "json",
	FormatText:        "text",
	FormatSRT:         "srt",
	FormatVerboseJSON: "verbose_json",
	FormatVTT:         "vtt",
}

var stringToFormat = map[string]Format{
	"json":         FormatJSON,
	"text":         FormatText,
	"srt":          FormatSRT,
	"verbose_json": FormatVerboseJSON,
	"vtt":          FormatVTT,
}
