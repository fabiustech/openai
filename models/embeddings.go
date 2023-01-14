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
	Unknown Embedding = iota
	AdaSimilarity
	BabbageSimilarity
	CurieSimilarity
	DavinciSimilarity
	AdaSearchDocument
	AdaSearchQuery
	BabbageSearchDocument
	BabbageSearchQuery
	CurieSearchDocument
	CurieSearchQuery
	DavinciSearchDocument
	DavinciSearchQuery
	AdaCodeSearchCode
	AdaCodeSearchText
	BabbageCodeSearchCode
	BabbageCodeSearchText
	AdaEmbeddingV2
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
