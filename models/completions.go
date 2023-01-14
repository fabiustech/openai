package models

type Completion int

const (
	TextDavinci003 Completion = iota
	TextDavinci002
	TextCurie001
	TextBabbage001
	TextAda001
	TextDavinci001
	DavinciInstructBeta
	Davinci
	CurieInstructBeta
	Curie
	Ada
	Babbage
)

var completionToString = map[Completion]string{
	TextDavinci003:      "text-davinci-003",
	TextDavinci002:      "text-davinci-002",
	TextCurie001:        "text-curie-001",
	TextBabbage001:      "text-babbage-001",
	TextAda001:          "text-ada-001",
	TextDavinci001:      "text-davinci-001",
	DavinciInstructBeta: "davinci-instruct-beta",
	Davinci:             "davinci",
	CurieInstructBeta:   "curie-instruct-beta",
	Curie:               "curie",
	Ada:                 "ada",
	Babbage:             "babbage",
}

// Codex Defines the models provided by OpenAI.
// These models are designed for code-specific tasks, and use
// a different tokenizer which optimizes for whitespace.
const (
	CodexCodeDavinci002 = "code-davinci-002"
	CodexCodeCushman001 = "code-cushman-001"
	CodexCodeDavinci001 = "code-davinci-001"
)
