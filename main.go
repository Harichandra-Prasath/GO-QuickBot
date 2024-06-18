package main

import (
	"fmt"

	"github.com/joho/godotenv"
)

func check_err(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	check_err(godotenv.Load())
	Listen()
	transcript := getTranscript()
	fmt.Println(transcript)
}
