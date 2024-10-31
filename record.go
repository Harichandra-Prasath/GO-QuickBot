package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-audio/audio"
	"github.com/go-audio/wav"
	"github.com/gordonklaus/portaudio"
)

func Listen() {
	fmt.Printf("Listening... Press ctrl + c to stop:")

	sig_chan := make(chan os.Signal, 1)
	signal.Notify(sig_chan, os.Interrupt, syscall.SIGTERM)

	f, err := os.Create("temp.wav")
	check_err(err)
	defer f.Close()

	buffer := make([]int16, 64)

	recorder, err := portaudio.OpenDefaultStream(1, 0, 44100, len(buffer), buffer)
	check_err(err)

	encoder := wav.NewEncoder(f, 44100, 16, 1, 1)
	defer encoder.Close()

	check_err(recorder.Start())
	defer recorder.Close()

loop:
	for {
		check_err(recorder.Read())
		audioData := make([]int, 64)
		for i, v := range buffer {
			audioData[i] = int(v)
		}
		audio_buff := &audio.IntBuffer{
			Data: audioData,
			Format: &audio.Format{
				SampleRate:  44100,
				NumChannels: 1,
			},
			SourceBitDepth: 16,
		}
		check_err(encoder.Write(audio_buff))
		select {
		case <-sig_chan:
			break loop
		default:
		}
	}
	check_err(recorder.Stop())

}
