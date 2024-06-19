package main

import (
	"flag"
	"net/http"

	"github.com/gordonklaus/portaudio"
)

var Client *http.Client
var Stop_Chan chan struct{}
var Pipeline *PipeLine

var OPENAI_KEY string

func check_err(err error) {
	if err != nil {
		panic(err)
	}
}

func Initialize_system() {
	Stop_Chan = make(chan struct{})
	Client = &http.Client{}
	Pipeline = getNewPipeLine()
	portaudio.Initialize()
}

func main() {

	flag.StringVar(&OPENAI_KEY, "OPENAI_KEY", "", "Provide Open AI Api key")
	flag.Parse()

	if OPENAI_KEY == "" {
		panic("No API Key Provided")
	}
	Initialize_system()
	defer portaudio.Terminate()

main_loop:
	for {
		select {
		case <-Stop_Chan:
			break main_loop
		default:
			Listen()
			Pipeline.Process()
		}
	}

}
