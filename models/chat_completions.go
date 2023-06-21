package models

type ChatCompletion int

const (
	// UnknownChatCompletion represents an invalid ChatCompletion model.
	UnknownChatCompletion ChatCompletion = iota
	// GPT3Dot5Turbo is the most capable GPT-3.5 model and optimized for chat at 1/10th the cost of text-davinci-003.
	// Will be updated with our latest model iteration.
	GPT3Dot5Turbo
	// GPT3Dot5Turbo0301 is a snapshot of gpt-3.5-turbo from March 1st 2023. Unlike gpt-3.5-turbo, this model will not
	// receive updates, and will only be supported for a three month period ending on June 1st 2023.
	// Deprecated.
	GPT3Dot5Turbo0301

	// GPT3Dot5Turbo0613 includes the same function calling as GPT-4 as well as more reliable steerability via the
	// system message, two features that allow developers to guide the model's responses more effectively.
	GPT3Dot5Turbo0613

	// GPT3Dot5Turbo16K offers 4 times the context length of gpt-3.5-turbo at twice the price: $0.003 per 1K input
	// tokens and $0.004 per 1K output tokens. 16k context means the model can now support ~20 pages of text in a
	// single request.
	GPT3Dot5Turbo16K

	// GPT-4 models are currently in a limited beta and only accessible to those who have been granted access.
	// Please join the waitlist to get access when capacity is available.
	// https://openai.com/waitlist/gpt-4-api

	// GPT4 is more capable than any GPT-3.5 model, able to do more complex tasks, and optimized for chat.
	// Will be updated with our latest model iteration.
	GPT4
	// GPT4_0314 is a snapshot of gpt-4 from March 14th 2023. Unlike gpt-4, this model will not receive updates, and
	// will only be supported for a three month period ending on June 14th 2023.
	//nolint:revive,stylecheck // This would be unreadable otherwise.
	// Deprecated.
	GPT4_0314

	GPT4_0613
	// GPT4_32K has the same capabilities as the base gpt-4 mode but with 4x the context length.
	// Will be updated with our latest model iteration.
	//nolint:revive,stylecheck // This would be unreadable otherwise.
	GPT4_32K
	// GPT4_32K_0314 is a snapshot of gpt-4-32 from March 14th 2023. Unlike gpt-4-32k, this model will not receive
	// updates, and will only be supported for a three month period ending on June 14th 2023.
	// Deprecated.
	//nolint:revive,stylecheck // This would be unreadable otherwise.
	GPT4_32K_0314
)

// String implements the fmt.Stringer interface.
func (c ChatCompletion) String() string {
	return chatCompletionToString[c]
}

// MarshalText implements the encoding.TextMarshaler interface.
func (c ChatCompletion) MarshalText() ([]byte, error) {
	return []byte(c.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
// On unrecognized value, it sets |e| to Unknown.
func (c *ChatCompletion) UnmarshalText(b []byte) error {
	if val, ok := stringToChatCompletion[(string(b))]; ok {
		*c = val
		return nil
	}

	*c = UnknownChatCompletion

	return nil
}

var chatCompletionToString = map[ChatCompletion]string{
	GPT3Dot5Turbo:     "gpt-3.5-turbo",
	GPT3Dot5Turbo0301: "gpt-3.5-turbo-0301",
	GPT3Dot5Turbo0613: "gpt-3.5-turbo-0613",
	GPT3Dot5Turbo16K:  "gpt-3.5-turbo-16k",
	GPT4:              "gpt-4",
	GPT4_0613:         "gpt-4-0613",
	GPT4_0314:         "gpt-4-0314",
	GPT4_32K:          "gpt-4-32k",
	GPT4_32K_0314:     "gpt-4-32k-0314",
}

var stringToChatCompletion = map[string]ChatCompletion{
	"gpt-3.5-turbo":      GPT3Dot5Turbo,
	"gpt-3.5-turbo-0301": GPT3Dot5Turbo0301,
	"gpt-3.5-turbo-0613": GPT3Dot5Turbo0613,
	"gpt-3.5-turbo-16k":  GPT3Dot5Turbo16K,
	"gpt-4":              GPT4,
	"gpt-4-0314":         GPT4_0314,
	"gpt-4-0613":         GPT4_0613,
	"gpt-4-32k":          GPT4_32K,
	"gpt-4-32k-0314":     GPT4_32K_0314,
}
