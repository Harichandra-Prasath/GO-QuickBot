package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

func (P *PipeLine) getResponse(question string) string {

	var response LLM_Response

	user_message := Message{
		Role:    "user",
		Content: question,
	}

	P.LLM_History = append(P.LLM_History, user_message)

	cb := CompletionBody{
		Model:    "gpt-4o",
		Messages: P.LLM_History,
		Stream:   false,
	}

	raw_body, err := json.Marshal(cb)
	check_err(err)

	body := new(bytes.Buffer)
	body.Write(raw_body)

	request, _ := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", body)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", OPENAI_KEY))

	resp, _ := Client.Do(request)
	resp_body := resp.Body
	defer resp.Body.Close()

	content, _ := io.ReadAll(resp_body)

	if resp.StatusCode != 200 {
		fmt.Println(string(content))
		return ""
	}

	json.Unmarshal(content, &response)

	assistant_message := response.Choices[0].Message

	P.LLM_History = append(P.LLM_History, assistant_message)
	fmt.Printf("Assistant: %s\n\n", assistant_message.Content)

	return assistant_message.Content
}
