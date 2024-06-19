package main

import (
	"net/http"

	"github.com/gordonklaus/portaudio"
	"github.com/joho/godotenv"
)

var Client *http.Client
var Stop_Chan chan struct{}
var Pipeline *PipeLine

func check_err(err error) {
	if err != nil {
		panic(err)
	}
}

func Initialize_system() {
	check_err(godotenv.Load())

	Stop_Chan = make(chan struct{})
	Client = &http.Client{}
	Pipeline = getNewPipeLine()
	portaudio.Initialize()

}

func main() {

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
