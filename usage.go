package openai

// Usage Represents the total token usage per request to OpenAI.
type Usage struct {
	// PromptTokens is the number of tokens in the request's prompt.
	PromptTokens int `json:"prompt_tokens"`
	// CompletionTokens is the number of tokens in the completion response.
	// Will not be set for requests to the embeddings endpoint.
	CompletionTokens int `json:"completion_tokens,omitempty"`
	// Total tokens is the sum of PromptTokens and CompletionTokens.
	TotalTokens int `json:"total_tokens"`
}
