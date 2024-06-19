package main

type PipeLine struct {
	LLM_History []Message
}

func getNewPipeLine() *PipeLine {
	return &PipeLine{
		LLM_History: make([]Message, 0),
	}
}

func (P *PipeLine) Process() {

	// Get the transcript
	transcript := P.getTranscript()

	if transcript == "Kill." {
		close(Stop_Chan)
		return
	}

	// Get the response
	generated_response := P.getResponse(transcript)

	// Play the generated response
	P.Speak(generated_response)
}
