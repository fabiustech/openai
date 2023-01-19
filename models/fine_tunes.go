package models

type FineTune int

const (
	UnknownFineTune FineTune = iota
	// Davinci most capable of the older versions of the GPT-3 models
	// and is intended to be used with the fine-tuning endpoints.
	Davinci
	// Curie is very capable, but faster and lower cost than Davinci. It is
	// an older version of the GPT-3 models and is intended to be used with
	// the fine-tuning endpoints.
	Curie
	// Babbage is capable of straightforward tasks, very fast, and lower cost.
	// It is an older version of the GPT-3 models and is intended to be used
	// with the fine-tuning endpoints.
	Babbage
	// Ada is capable of very simple tasks, usually the fastest model in the
	//  GPT-3 series, and lowest cost. It is an older version of the GPT-3
	// models and is intended to be used with the fine-tuning endpoints.
	Ada
)

// String implements the fmt.Stringer interface.
func (f FineTune) String() string {
	return fineTuneToString[f]
}

// MarshalText implements the encoding.TextMarshaler interface.
func (f FineTune) MarshalText() ([]byte, error) {
	return []byte(f.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
// On unrecognized value, it sets |e| to Unknown.
func (f *FineTune) UnmarshalText(b []byte) error {
	if val, ok := stringToFineTune[(string(b))]; ok {
		*f = val
		return nil
	}

	*f = UnknownFineTune

	return nil
}

var fineTuneToString = map[FineTune]string{
	Davinci: "davinci",
	Curie:   "curie",
	Ada:     "ada",
	Babbage: "babbage",
}

var stringToFineTune = map[string]FineTune{
	"davinci": Davinci,
	"curie":   Curie,
	"ada":     Ada,
	"babbage": Babbage,
}

// FineTunedModel represents the name of a fine-tuned model which was
// previously generated.
type FineTunedModel string

// NewFineTunedModel converts a string to FineTunedModel.
func NewFineTunedModel(name string) FineTunedModel {
	return FineTunedModel(name)
}
