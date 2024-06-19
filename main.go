package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/joho/godotenv"
)

var Client *http.Client

func check_err(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	check_err(godotenv.Load())
	Client = &http.Client{}
	Listen()
	t1 := time.Now()
	transcript := getTranscript()
	fmt.Println(time.Since(t1).Seconds())
	response := getResponse(transcript)
	fmt.Println(time.Since(t1).Seconds())
	Speak(response)
}
