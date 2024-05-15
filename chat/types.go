package chat

type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
	Name    string `json:"name,omitempty"`
}

type ChatRequest struct {
	Model            string        `json:"model"`
	Messages         []ChatMessage `json:"messages"`
	MaxTokens        int           `json:"max_tokens,omitempty"`
	Temperature      float32       `json:"temperature,omitempty"`
	TopP             float32       `json:"top_p,omitempty"`
	FrequencyPenalty int           `json:"frequecy_penalty,omitempty"`
	PresencePenalty  int           `json:"presence_penatly,omitempty"`
}

type ChatResponse struct {
	Id                string       `json:"id"`
	Choices           []ChatChoice `json:"choices"`
	Created           int          `json:"created"`
	Model             string       `json:"model"`
	SystemFingerprint string       `json:"system_fingerprint"`
	Object            string       `json:"object"`
	Usage             ChatUsage    `json:"usage"`
}

type ChatChoice struct {
	Index        int         `json:"index"`
	Message      ChatMessage `json:"message"`
	FinishReason string      `json:"finish_reson"`
}

type ChatUsage struct {
	PromptTokens     int `json:"prompt_token"`
	CompletionTokens int `json:"completion_token"`
	TotalTokens      int `json:"total_tokens"`
}
