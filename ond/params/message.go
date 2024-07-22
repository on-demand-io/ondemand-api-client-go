package params

type MediaSource string

var (
	MediaDocument MediaSource = "document"
	MediaVideo    MediaSource = "video"
	MediaAudio    MediaSource = "audio"
	MediaYoutube  MediaSource = "youtube"
	MediaImage    MediaSource = "image"
)

func (m MediaSource) String() string { return string(m) }

type MessageType string

var (
	MessageTypeText  MessageType = "text"
	MessageTypeMedia MessageType = "media"
)

func (m MessageType) String() string { return string(m) }
