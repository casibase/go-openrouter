package openrouter

const (
	GooglePalm2CodeChatBison = "google/palm-2-codechat-bison"
	GooglePalm2ChatBison     = "google/palm-2-chat-bison"
	OpenaiGpt35Turbo         = "openai/gpt-3.5-turbo"
	OpenaiGpt35Turbo16k      = "openai/gpt-3.5-turbo-16k"
	OpenaiGpt4               = "openai/gpt-4"
	OpenaiGpt432K            = "openai/gpt-4-32k"
	AnthropicClaude2         = "anthropic/claude-2"
	AnthropicClaudeInstantV1 = "anthropic/claude-instant-v1"
	MetaLlamaLlama213bChat   = "meta-llama/llama-2-13b-chat"
	MetaLlamaLlama270bChat   = "meta-llama/llama-2-70b-chat"
	Palm2CodeChatBison       = "palm-2-codechat-bison"
	Palm2ChatBison           = "palm-2-chat-bison"
	Gpt35Turbo               = "gpt-3.5-turbo"
	Gpt35Turbo16k            = "gpt-3.5-turbo-16k"
	Gpt4                     = "gpt-4"
	G432K                    = "gpt-4-32k"
	Claude2                  = "claude-2"
	ClaudeInstantV1          = "claude-instant-v1"
	Llama213bChat            = "llama-2-13b-chat"
	Llama270bChat            = "llama-2-70b-chat"
)

var (
	enableModels = map[string]bool{
		GooglePalm2CodeChatBison: true,
		GooglePalm2ChatBison:     true,
		OpenaiGpt35Turbo:         true,
		OpenaiGpt35Turbo16k:      true,
		OpenaiGpt4:               true,
		OpenaiGpt432K:            true,
		AnthropicClaude2:         true,
		AnthropicClaudeInstantV1: true,
		MetaLlamaLlama213bChat:   true,
		MetaLlamaLlama270bChat:   true,
	}
	wrapperModels = map[string]string{
		Palm2CodeChatBison: GooglePalm2CodeChatBison,
		Palm2ChatBison:     GooglePalm2ChatBison,
		Gpt35Turbo:         OpenaiGpt35Turbo,
		Gpt35Turbo16k:      OpenaiGpt35Turbo16k,
		Gpt4:               OpenaiGpt4,
		G432K:              OpenaiGpt432K,
		Claude2:            AnthropicClaude2,
		ClaudeInstantV1:    AnthropicClaudeInstantV1,
		Llama213bChat:      MetaLlamaLlama213bChat,
		Llama270bChat:      MetaLlamaLlama270bChat,
	}
)

func checkSupportsModel(model string) bool {
	return enableModels[model]
}

type ChatCompletionMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// ChatCompletionRequest represents a request structure for chat completion API.
type ChatCompletionRequest struct {
	Model       string                  `json:"model"`
	Messages    []ChatCompletionMessage `json:"messages"`
	MaxTokens   int                     `json:"max_tokens,omitempty"`
	Stream      bool                    `json:"stream,omitempty"`
	Temperature *float32                `json:"temperature,omitempty"`
	TopP        *float32                `json:"top_p,omitempty"`
	TopK        *uint                   `json:"top_k,omitempty"`
}

type Index struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatCompletionChoice struct {
	Message      Index  `json:"message,omitempty"`
	FinishReason string `json:"finish_reason,omitempty"`
	Delta        Index  `json:"delta,omitempty"`
	Index        uint   `json:"index,omitempty"`
}

// ChatCompletionResponse represents a response structure for chat completion API.
type ChatCompletionResponse struct {
	ID      string                 `json:"id,omitempty"`
	Object  string                 `json:"object,omitempty"`
	Created int64                  `json:"created,omitempty"`
	Model   string                 `json:"model"`
	Choices []ChatCompletionChoice `json:"choices"`
	//Usage   Usage                  `json:"usage,omitempty"`
}

type Usage struct {
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}
