package models

// Completion represents all models available for use with the Completions endpoint.
type Completion int

const (
	// UnknownCompletion represents and invalid Completion model.
	UnknownCompletion Completion = iota
	// TextDavinci003 is the most capable GPT-3 model. Can do any task the other models can do,
	// often with higher quality, longer output and better instruction-following. Also supports
	// inserting completions within text.
	//
	// Supports up to 4,000 tokens. Training data up to Jun 2021.
	TextDavinci003
	// TextDavinci002 is an older version of the most capable GPT-3 model. Can do any task the
	// other models can do, often with higher quality, longer output and better
	// instruction-following. Also supports inserting completions within text.
	//
	// Supports up to 4,000 tokens.
	//
	// Deprecated: Use TextDavinci003 instead.
	TextDavinci002
	// TextCurie001 is very capable, but faster and lower cost than Davinci.
	//
	// Supports up to 2,048 tokens. Training data up to Oct 2019.
	TextCurie001
	// TextBabbage001 is capable of straightforward tasks, very fast, and lower cost.
	//
	// Supports up to 2,048 tokens. Training data up to Oct 2019.
	TextBabbage001
	// TextAda001 is capable of very simple tasks, usually the fastest model in the
	// GPT-3 series, and lowest cost.
	//
	// Supports up to 2,048 tokens. Training data up to Oct 2019.
	TextAda001
	// TextDavinci001 ... (?).
	TextDavinci001

	// DavinciInstructBeta is the most capable model in the InstructGPT series.
	// It is much better at following user intentions than GPT-3 while also being
	// more truthful and less toxic. InstructGPT is better than GPT-3 at following
	// English instructions.
	DavinciInstructBeta
	// CurieInstructBeta is very capable, but faster and lower cost than Davinci.
	// It is much better at following user intentions than GPT-3 while also being
	// more truthful and less toxic. InstructGPT is better than GPT-3 at following
	// English instructions.
	CurieInstructBeta

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
	// CodeDavinci002 is the most capable Codex model. Particularly good at
	// translating natural language to code. In addition to completing code,
	// also supports inserting completions within code.
	//
	// Supports up to 8,000 tokens. Training data up to Jun 2021.
	CodeDavinci002
	// CodeCushman001 is almost as capable as Davinci Codex, but slightly faster.
	// This speed advantage may make it preferable for real-time applications.
	//
	// Supports up to 2,048 tokens.
	CodeCushman001
	// CodeDavinci001 is and older version of the most capable Codex model.
	// Particularly good at translating natural language to code. In addition
	// to completing code, also supports inserting completions within code.
	//
	// Deprecated: Use CodeDavinci002 instead.
	CodeDavinci001

	// TextDavinciInsert002 was a beta model released for insertion.
	//
	// Deprecated: Insertion should be done via the text models.
	TextDavinciInsert002
	// TextDavinciInsert001 was a beta model released for insertion.
	//
	// Deprecated: Insertion should be done via the text models.
	TextDavinciInsert001
)

// String implements the fmt.Stringer interface.
func (c Completion) String() string {
	return completionToString[c]
}

// MarshalText implements the encoding.TextMarshaler interface.
func (c Completion) MarshalText() ([]byte, error) {
	return []byte(c.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
// On unrecognized value, it sets |e| to Unknown.
func (c *Completion) UnmarshalText(b []byte) error {
	if val, ok := stringToCompletion[(string(b))]; ok {
		*c = val
		return nil
	}

	*c = UnknownCompletion

	return nil
}

var completionToString = map[Completion]string{
	TextDavinci003:       "text-davinci-003",
	TextDavinci002:       "text-davinci-002",
	TextCurie001:         "text-curie-001",
	TextBabbage001:       "text-babbage-001",
	TextAda001:           "text-ada-001",
	TextDavinci001:       "text-davinci-001",
	DavinciInstructBeta:  "davinci-instruct-beta",
	CurieInstructBeta:    "curie-instruct-beta",
	Davinci:              "davinci",
	Curie:                "curie",
	Ada:                  "ada",
	Babbage:              "babbage",
	CodeDavinci002:       "code-davinci-002",
	CodeCushman001:       "code-cushman-001",
	CodeDavinci001:       "code-davinci-001",
	TextDavinciInsert002: "text-davinci-insert-002",
	TextDavinciInsert001: "text-davinci-insert-001",
}

var stringToCompletion = map[string]Completion{
	"text-davinci-003":        TextDavinci003,
	"text-davinci-002":        TextDavinci002,
	"text-curie-001":          TextCurie001,
	"text-babbage-001":        TextBabbage001,
	"text-ada-001":            TextAda001,
	"text-davinci-001":        TextDavinci001,
	"davinci-instruct-beta":   DavinciInstructBeta,
	"curie-instruct-beta":     CurieInstructBeta,
	"davinci":                 Davinci,
	"curie":                   Curie,
	"ada":                     Ada,
	"babbage":                 Babbage,
	"code-davinci-002":        CodeDavinci002,
	"code-cushman-001":        CodeCushman001,
	"code-davinci-001":        CodeDavinci001,
	"text-davinci-insert-002": TextDavinciInsert002,
	"text-davinci-insert-001": TextDavinciInsert001,
}
