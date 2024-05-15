package chat

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/charmbracelet/log"
)

var openAiUrl = "https://api.openai.com/v1/chat/completions"

func GetChatResponse(input string) string {
	systemPrompt := ChatMessage{
		Role:    "system",
		Content: "You are a helpful assistant. Respond in markdown only.",
	}

	userPrompt := ChatMessage{
		Role:    "user",
		Content: input,
	}
	chatBody := ChatRequest{
		Model: "gpt-4o",
		Messages: []ChatMessage{
			userPrompt,
			systemPrompt,
		},
		MaxTokens:        3500,
		Temperature:      1,
		TopP:             1,
		FrequencyPenalty: 0,
		PresencePenalty:  0,
	}

	body, err := json.Marshal(chatBody)
	if err != nil {
		log.Fatal("Error converting body to JSON")
	}

	bodyReader := bytes.NewReader(body)

	req, err := http.NewRequest("POST", openAiUrl, bodyReader)
	if err != nil {
		log.Fatal("Error creating Open AI API request")
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+os.Getenv("openAiKey"))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("Error with openAI API", "Error", err)
	}

	var chatResp ChatResponse
	respBodyBuffer, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Error reading response body", "error", err)
	}

	err = json.Unmarshal(respBodyBuffer, &chatResp)
	if err != nil {
		log.Fatal("Error trying to parse response JSON", "Error", err)
	}

	return chatResp.Choices[0].Message.Content
}
