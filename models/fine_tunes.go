package models

type FineTune int

const (
	// Davinci most capable of the older versions of the GPT-3 models
	// and is intended to be used with the fine-tuning endpoints.
	Davinci FineTune = iota
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
