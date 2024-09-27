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

	// GPT3Dot5Turbo1106 is a GPT-3.5 Turbo model with improved instruction following, JSON mode, reproducible outputs,
	// parallel function calling, and more. Returns a maximum of 4,096 output tokens.
	GPT3Dot5Turbo1106

	// GPT3Dot5Turbo0125 is the latest GPT-3.5 Turbo model with higher accuracy at responding in requested formats and a
	// fix for a bug which caused a text encoding issue for non-English language function calls. Returns a maximum of
	// 4,096 output tokens.
	GPT3Dot5Turbo0125

	// GPT4 is more capable than any GPT-3.5 model, able to do more complex tasks, and optimized for chat.
	// Will be updated with our latest model iteration.
	GPT4
	// GPT4_0314 is a snapshot of gpt-4 from March 14th 2023. Unlike gpt-4, this model will not receive updates, and
	// will only be supported for a three month period ending on June 14th 2023.
	//nolint:revive,stylecheck // This would be unreadable otherwise.
	// Deprecated.
	GPT4_0314

	// GPT4_0613 includes an updated and improved model with function calling.
	//nolint:revive,stylecheck // This would be unreadable otherwise.
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
	// GPT4_32K_0613 includes the same improvements as gpt-4-0613, along with an extended context length for better
	// comprehension of larger texts.
	//nolint:revive,stylecheck // This would be unreadable otherwise.
	GPT4_32K_0613

	// GPT4Turbo1106Preview (GPT-4 Turbo) is more capable and has knowledge of world events up to April 2023. It has a
	// 128k context window so it can fit the equivalent of more than 300 pages of text in a single prompt.
	// It also optimized its performance so we are able to offer GPT-4 Turbo at a 3x cheaper price for input tokens and
	// a 2x cheaper price for output tokens compared to GPT-4.
	GPT4Turbo1106Preview

	// GPT4Turbo0125Preview (GPT-4 Turbo) is the latest GPT-4 model intended to reduce cases of "laziness" where the model
	// doesn’t complete a task. Returns a maximum of 4,096 output tokens.
	GPT4Turbo0125Preview

	// GPT4TurboPreview (GPT-4 Turbo) is the latest GPT-4. It currently points to gpt-4-0125-preview, however, may
	// receive updates in the future.
	GPT4TurboPreview

	// GPT4o (GPT-4o) is the most advanced model. It is multimodal (accepting text or image inputs and outputting text),
	// and it has the same high intelligence as GPT-4 Turbo but is much more efficient—it generates text 2x faster and
	// is 50% cheaper. Additionally, GPT-4o has the best vision and performance across non-English languages of any model.
	GPT4o

	// GPT4o20240503 is a snapshot of gpt-4o from May 13th, 2024.
	GPT4o20240503

	// O1Preview is a part of the o1 series of large language models, which are trained with reinforcement learning to
	// perform complex reasoning. o1 models think before they answer, producing a long internal chain of thought before
	// responding to the user.
	O1Preview

	// O1Mini is a part of the o1 series of large language models, which are trained with reinforcement learning to
	// perform complex reasoning. o1 models think before they answer, producing a long internal chain of thought before
	// responding to the user.
	O1Mini
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
	GPT3Dot5Turbo:        "gpt-3.5-turbo",
	GPT3Dot5Turbo0301:    "gpt-3.5-turbo-0301",
	GPT3Dot5Turbo0613:    "gpt-3.5-turbo-0613",
	GPT3Dot5Turbo16K:     "gpt-3.5-turbo-16k",
	GPT3Dot5Turbo1106:    "gpt-3.5-turbo-1106",
	GPT3Dot5Turbo0125:    "gpt-3.5-turbo-0125",
	GPT4:                 "gpt-4",
	GPT4_0613:            "gpt-4-0613",
	GPT4_0314:            "gpt-4-0314",
	GPT4_32K:             "gpt-4-32k",
	GPT4_32K_0314:        "gpt-4-32k-0314",
	GPT4_32K_0613:        "gpt-4-32k-0613",
	GPT4Turbo1106Preview: "gpt-4-1106-preview",
	GPT4Turbo0125Preview: "gpt-4-0125-preview",
	GPT4TurboPreview:     "gpt-4-turbo-preview",
	GPT4o:                "gpt-4o",
	GPT4o20240503:        "gpt-4o-2024-05-13",
	O1Preview:            "o1-preview",
	O1Mini:               "o1-mini",
}

var stringToChatCompletion = map[string]ChatCompletion{
	"gpt-3.5-turbo":       GPT3Dot5Turbo,
	"gpt-3.5-turbo-0301":  GPT3Dot5Turbo0301,
	"gpt-3.5-turbo-0613":  GPT3Dot5Turbo0613,
	"gpt-3.5-turbo-16k":   GPT3Dot5Turbo16K,
	"gpt-3.5-turbo-1106":  GPT3Dot5Turbo1106,
	"gpt-3.5-turbo-0125":  GPT3Dot5Turbo0125,
	"gpt-4":               GPT4,
	"gpt-4-0314":          GPT4_0314,
	"gpt-4-0613":          GPT4_0613,
	"gpt-4-32k":           GPT4_32K,
	"gpt-4-32k-0314":      GPT4_32K_0314,
	"gpt-4-32k-0613":      GPT4_32K_0613,
	"gpt-4-1106-preview":  GPT4Turbo1106Preview,
	"gpt-4-0125-preview":  GPT4Turbo0125Preview,
	"gpt-4-turbo-preview": GPT4TurboPreview,
	"gpt-4o":              GPT4o,
	"gpt-4o-2024-05-13":   GPT4o20240503,
	"o1-preview":          O1Preview,
	"o1-mini":             O1Mini,
}
