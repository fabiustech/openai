package models

// Embedding enumerates the models which can be used
// to generate Embedding vectors.
type Embedding int

// String implements the fmt.Stringer interface.
func (e Embedding) String() string {
	return enumToString[e]
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e Embedding) MarshalText() ([]byte, error) {
	return []byte(e.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
// On unrecognized value, it sets |e| to Unknown.
func (e *Embedding) UnmarshalText(b []byte) error {
	if val, ok := stringToEnum[(string(b))]; ok {
		*e = val
		return nil
	}

	*e = Unknown

	return nil
}

const (
	// Unknown represents an invalid Embedding model.
	Unknown Embedding = iota

	// AdaEmbeddingV2 is the second-generation embedding model. OpenAI recommends using
	// text-embedding-ada-002 for nearly all use cases. It’s better, cheaper, and simpler to use.
	//
	// Supports up to 8191. Knowledge cutoff Sep 2021.
	AdaEmbeddingV2

	// The below models are first-generation models (those ending in -001) use the GPT-3
	// tokenizer and have a max input of 2046 tokens. First-generation embeddings are generated
	// by five different model families tuned for three different tasks: text search, text similarity
	// and code search. The search models come in pairs: one for short queries and one for long documents.
	// Each family includes up to four models on a spectrum of quality and speed.

	// Deprecated: OpenAI recommends using text-embedding-ada-002 for nearly all use cases.
	AdaSimilarity
	// Deprecated: OpenAI recommends using text-embedding-ada-002 for nearly all use cases.
	BabbageSimilarity
	// Deprecated: OpenAI recommends using text-embedding-ada-002 for nearly all use cases.
	CurieSimilarity
	// Deprecated: OpenAI recommends using text-embedding-ada-002 for nearly all use cases.
	DavinciSimilarity
	// Deprecated: OpenAI recommends using text-embedding-ada-002 for nearly all use cases.
	AdaSearchDocument
	// Deprecated: OpenAI recommends using text-embedding-ada-002 for nearly all use cases.
	AdaSearchQuery
	// Deprecated: OpenAI recommends using text-embedding-ada-002 for nearly all use cases.
	BabbageSearchDocument
	// Deprecated: OpenAI recommends using text-embedding-ada-002 for nearly all use cases.
	BabbageSearchQuery
	// Deprecated: OpenAI recommends using text-embedding-ada-002 for nearly all use cases.
	CurieSearchDocument
	// Deprecated: OpenAI recommends using text-embedding-ada-002 for nearly all use cases.
	CurieSearchQuery
	// Deprecated: OpenAI recommends using text-embedding-ada-002 for nearly all use cases.
	DavinciSearchDocument
	// Deprecated: OpenAI recommends using text-embedding-ada-002 for nearly all use cases.
	DavinciSearchQuery
	// Deprecated: OpenAI recommends using text-embedding-ada-002 for nearly all use cases.
	AdaCodeSearchCode
	// Deprecated: OpenAI recommends using text-embedding-ada-002 for nearly all use cases.
	AdaCodeSearchText
	// Deprecated: OpenAI recommends using text-embedding-ada-002 for nearly all use cases.
	BabbageCodeSearchCode
	// Deprecated: OpenAI recommends using text-embedding-ada-002 for nearly all use cases.
	BabbageCodeSearchText
)

var enumToString = map[Embedding]string{
	AdaSimilarity:         "text-similarity-ada-001",
	BabbageSimilarity:     "text-similarity-babbage-001",
	CurieSimilarity:       "text-similarity-curie-001",
	DavinciSimilarity:     "text-similarity-davinci-001",
	AdaSearchDocument:     "text-search-ada-doc-001",
	AdaSearchQuery:        "text-search-ada-query-001",
	BabbageSearchDocument: "text-search-babbage-doc-001",
	BabbageSearchQuery:    "text-search-babbage-query-001",
	CurieSearchDocument:   "text-search-curie-doc-001",
	CurieSearchQuery:      "text-search-curie-query-001",
	DavinciSearchDocument: "text-search-davinci-doc-001",
	DavinciSearchQuery:    "text-search-davinci-query-001",
	AdaCodeSearchCode:     "code-search-ada-code-001",
	AdaCodeSearchText:     "code-search-ada-text-001",
	BabbageCodeSearchCode: "code-search-babbage-code-001",
	BabbageCodeSearchText: "code-search-babbage-text-001",
	AdaEmbeddingV2:        "text-embedding-ada-002",
}

var stringToEnum = map[string]Embedding{
	"text-similarity-ada-001":       AdaSimilarity,
	"text-similarity-babbage-001":   BabbageSimilarity,
	"text-similarity-curie-001":     CurieSimilarity,
	"text-similarity-davinci-001":   DavinciSimilarity,
	"text-search-ada-doc-001":       AdaSearchDocument,
	"text-search-ada-query-001":     AdaSearchQuery,
	"text-search-babbage-doc-001":   BabbageSearchDocument,
	"text-search-babbage-query-001": BabbageSearchQuery,
	"text-search-curie-doc-001":     CurieSearchDocument,
	"text-search-curie-query-001":   CurieSearchQuery,
	"text-search-davinci-doc-001":   DavinciSearchDocument,
	"text-search-davinci-query-001": DavinciSearchQuery,
	"code-search-ada-code-001":      AdaCodeSearchCode,
	"code-search-ada-text-001":      AdaCodeSearchText,
	"code-search-babbage-code-001":  BabbageCodeSearchCode,
	"code-search-babbage-text-001":  BabbageCodeSearchText,
	"text-embedding-ada-002":        AdaEmbeddingV2,
}
