package images

type Size int

// Image sizes defined by the OpenAI API.
// TODO: make enum.
const (
	SizeInvalid Size = iota
	Size256x256
	Size512x512
	Size1024x1024
)

// String implements the fmt.Stringer interface.
func (s Size) String() string {
	return imageToString[s]
}

// MarshalText implements the encoding.TextMarshaler interface.
func (s Size) MarshalText() ([]byte, error) {
	return []byte(s.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
// On unrecognized value, it sets |e| to Unknown.
func (s *Size) UnmarshalText(b []byte) error {
	if val, ok := stringToImage[(string(b))]; ok {
		*s = val
		return nil
	}

	*s = SizeInvalid

	return nil
}

var imageToString = map[Size]string{
	Size256x256:   "256x256",
	Size512x512:   "512x512",
	Size1024x1024: "1024x1024",
}

var stringToImage = map[string]Size{
	"256x256":   Size256x256,
	"512x512":   Size512x512,
	"1024x1024": Size1024x1024,
}
