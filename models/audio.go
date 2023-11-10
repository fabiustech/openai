package models

type Audio int

const (
	UnknownAudioModel Audio = iota
	Whisper1
	Whisper2
)

// String implements the fmt.Stringer interface.
func (a Audio) String() string {
	return audioToString[a]
}

// MarshalText implements the encoding.TextMarshaler interface.
func (a Audio) MarshalText() ([]byte, error) {
	return []byte(a.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
// On unrecognized value, it sets |e| to Unknown.
func (a *Audio) UnmarshalText(b []byte) error {
	if val, ok := stringToAudio[(string(b))]; ok {
		*a = val
		return nil
	}

	*a = UnknownAudioModel

	return nil
}

var audioToString = map[Audio]string{
	Whisper1: "whisper-1",
	Whisper2: "whisper-2",
}

var stringToAudio = map[string]Audio{
	"whisper-1": Whisper1,
	"whisper-2": Whisper2,
}
