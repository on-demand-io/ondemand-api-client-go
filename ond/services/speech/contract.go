package speech

type ToTextResponse struct {
	Data ToTextResponseData `json:"data"`
}

type ToTextResponseData struct {
	Text string `json:"text"`
}

type ToSpeechResponse struct {
	Data TextToSpeechData `json:"data"`
}

type TextToSpeechData struct {
	AudioURL string `json:"audioUrl"`
}
