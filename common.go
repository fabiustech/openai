package openai

// Usage Represents the total token usage per request to OpenAI.
type Usage struct {
	// PromptTokens is the number of tokens in the passed prompt.
	PromptTokens int `json:"prompt_tokens"`
	// CompletionTokens is the number of tokens in the completion response.
	CompletionTokens int `json:"completion_tokens"`
	// Total tokens is the sum of PromptTokens and CompletionTokens.
	TotalTokens int `json:"total_tokens"`
}
