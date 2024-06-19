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

func getTranscript() string {

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
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("OPENAI_KEY")))
	req.Header.Set("Content-Type", writer.FormDataContentType())

	openai_res, err := Client.Do(req)
	check_err(err)

	defer openai_res.Body.Close()
	bodyText, err := io.ReadAll(openai_res.Body)
	check_err(err)

	check_err(json.Unmarshal(bodyText, &response))
	return response.Text
}
