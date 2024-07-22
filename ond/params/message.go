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

type ListMessageParams struct {
	// SessionID - Required
	// id of the chat session
	SessionID string

	// ExternalUserID - Optional
	// An identifier of the external user who created this message query
	ExternalUserID string `url:"externalUserId"`

	// Sort order of the messages based on its creation time
	// Default order is descending
	Sort Sort `url:"sort"`

	// Limit
	// Specifies the total number of results to retrieve. If the provided value is less than the minimum allowed limit,
	// it will be reset to the default value. Conversely, if the provided value exceeds the maximum allowed limit,
	// it will be reset to the maximum value.
	// Default values is 10
	Limit int32 `url:"limit"`

	// Cursor
	// It acts as a pagination key and used to retrieve next results.
	// For first iteration, this parameter must not be set.
	// For subsequent iterations, it must be set to "pagination.next" value from previous response.
	Cursor string `url:"cursor"`
}
