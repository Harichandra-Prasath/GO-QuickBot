package main

import (
	"fmt"
	"time"

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
	t1 := time.Now()
	transcript := getTranscript()
	fmt.Println(time.Since(t1).Seconds())
	fmt.Println(getResponse(transcript))
	fmt.Println(time.Since(t1).Seconds())
}
