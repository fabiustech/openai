package images

// Size represents the enum values for the image sizes that
// you can generate. Smaller sizes are faster to generate.
type Size int

const (
	// SizeInvalid represents and invalid Size option.
	SizeInvalid Size = iota
	// Size256x256 specifies that the API will return an image that is
	// 256x256 pixels.
	Size256x256
	// Size512x512 specifies that the API will return an image that is
	// 512x512 pixels.
	Size512x512
	// Size1024x1024 specifies that the API will return an image that is
	// 1024x1024 pixels.
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
