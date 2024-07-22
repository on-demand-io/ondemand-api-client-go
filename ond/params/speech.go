package params

type SpeechToTextParams struct {
	AudioURL string `json:"audioUrl"`
}

type TextToSpeechParams struct {
	Model     ServiceModel `json:"model"`
	TextInput string       `json:"input"`
	Voice     ServiceVoice `json:"voice"`
}

type ServiceModel string

var (
	ModelTTS1 ServiceModel = "tts-1"
)

func (sm ServiceModel) String() string { return string(sm) }

type ServiceVoice string

var (
	VoiceAlloy ServiceVoice = "alloy"
)

func (v ServiceVoice) String() string { return string(v) }
