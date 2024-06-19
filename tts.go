package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
)

type SpeechBody struct {
	Model string `json:"model"`
	Voice string `json:"voice"`
	Text  string `json:"input"`
}

func (P *PipeLine) Speak(text string) {

	payload := SpeechBody{
		Model: "tts-1",
		Voice: "fable",
		Text:  text,
	}

	raw_body, _ := json.Marshal(payload)

	body := new(bytes.Buffer)
	body.Write(raw_body)

	request, _ := http.NewRequest("POST", "https://api.openai.com/v1/audio/speech", body)
	request.Header.Add("Content-type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("OPENAI_KEY")))

	chunk := make([]byte, 512)

	resp, _ := Client.Do(request)

	resp_body := resp.Body
	defer resp_body.Close()

	process := exec.Command("ffplay", "-autoexit", "-", "-nodisp")
	P_in, _ := process.StdinPipe()
	process.Start()

	for {
		n, _ := resp_body.Read(chunk)
		if n == 0 {
			break
		}
		P_in.Write(chunk)
	}

	P_in.Close()
	process.Wait()
}
