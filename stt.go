package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

type STT_Response struct {
	Text string
}

func (P *PipeLine) getTranscript() string {

	var response STT_Response

	form := new(bytes.Buffer)
	writer := multipart.NewWriter(form)
	fw, _ := writer.CreateFormFile("file", filepath.Base("temp.wav"))
	fd, _ := os.Open("temp.wav")

	defer fd.Close()
	defer os.Remove(fd.Name())

	io.Copy(fw, fd)

	writer.WriteField("language", "en")
	writer.WriteField("model", "whisper-1")
	writer.Close()

	req, _ := http.NewRequest("POST", "https://api.openai.com/v1/audio/transcriptions", form)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", OPENAI_KEY))
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := Client.Do(req)
	check_err(err)

	defer resp.Body.Close()
	content, err := io.ReadAll(resp.Body)
	check_err(err)

	if resp.StatusCode != 200 {
		fmt.Println(string(content))
		return ""
	}

	check_err(json.Unmarshal(content, &response))

	fmt.Printf("\n\nUser: %s\n", response.Text)
	return response.Text
}
