package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type LLM_Response struct {
	Model   string   `json:"model"`
	Choices []Choice `json:"choices"`
}

type Choice struct {
	Message Message `json:"message"`
	Reason  string  `json:"finish_reason"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type CompletionBody struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
	Stream   bool      `json:"stream"`
}

func getResponse(question string) string {

	var messages []Message
	var response LLM_Response

	m := Message{
		Role:    "user",
		Content: question,
	}
	messages = append(messages, m)

	cb := CompletionBody{
		Model:    "gpt-4o",
		Messages: messages,
		Stream:   false,
	}

	raw_body, err := json.Marshal(cb)
	check_err(err)

	body := new(bytes.Buffer)
	body.Write(raw_body)

	request, _ := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", body)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("OPENAI_KEY")))

	resp, _ := Client.Do(request)

	resp_body := resp.Body
	defer resp.Body.Close()

	content, _ := io.ReadAll(resp_body)
	json.Unmarshal(content, &response)

	return response.Choices[0].Message.Content
}
