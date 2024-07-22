package chat

import (
	"github.com/dinson/ond-api-client-go/ond/params"
	"time"
)

type SubmitQueryResponse struct {
	Data QueryResponse `json:"data"`
}

type QueryResponse struct {
	// SessionID
	// id of chat session
	SessionID string `json:"sessionId"`

	// MessageID
	// id of the chat message
	MessageID string `json:"messageId"`

	// Answer to the query
	Answer string `json:"answer"`

	// Status
	// Current status of the chat message
	// Values can be "processing", "completed" or "failed"
	Status params.ChatStatus `json:"status"`
}

type GetSessionResponse struct {
	Data SessionData `json:"data"`
}

type GetMessageResponse struct {
	Data Message `json:"data"`
}

type Message struct {
	// ID of chat message
	ID string `json:"id"`

	// SessionID
	// id of chat session
	SessionID string `json:"sessionId"`

	// CompanyID
	// id of the company linked to chat session
	CompanyID string `json:"companyId"`

	// ExternalUserID
	// id of the external user
	ExternalUserID string `json:"externalUserId"`

	// PluginIDs
	// List of plugins associated with the chat message
	PluginIDs []string `json:"pluginIds"`

	// ResponseMode
	// Possible values - "sync", "stream" or "webhook"
	ResponseMode params.ResponseMode `json:"responseMode"`

	// EndpointID
	// id of the endpoint used to fulfill the query
	EndpointID string `json:"endpointId"`

	// Status
	// Current status of the chat message
	// Possible values - "processing", "completed", "failed"
	Status params.ChatStatus `json:"status"`

	// Type of the chat message
	// Possible values - "text", "media"
	Type params.MessageType `json:"type"`

	// Media  content of the chat message. Only set when message type is media
	Media *MessageMedia `json:"media,omitempty"`

	// Query
	// Only set when message type is text
	Query *string `json:"query,omitempty"`

	// Answer to the query. Only set when message type is text
	Answer *string `json:"answer,omitempty"`

	// CreatedBy
	// id of the company user created the chat message
	CreatedBy string `json:"createdBy"`

	// CreatedAt
	// Timestamp when the chat message was created
	CreatedAt time.Time `json:"createdAt"`

	// UpdatedAt
	// Timestamp when the chat message was last updated
	UpdatedAt time.Time `json:"updatedAt"`
}

type MessageMedia struct {
	ID string `json:"id"`

	// Name of the media
	Name string `json:"name"`

	// Source of the media
	// Possible values - "document", "video, "audio",  "youtube",  "image"
	Source params.MediaSource `json:"source"`

	// URL of the media content
	URL string `json:"url"`

	// Context of the media
	Context string `json:"context"`
}

type ListMessagesResponse struct {
	Data       []Message  `json:"data"`
	Pagination Pagination `json:"pagination"`
}

type ListSessionsResponse struct {
	Data       []SessionData `json:"data"`
	Pagination Pagination    `json:"pagination"`
}

type CreateSessionResponse struct {
	Data SessionData `json:"data"`
}

type SessionData struct {
	ID             string    `json:"id"`
	CompanyID      string    `json:"companyId"`
	ExternalUserID string    `json:"externalUserId"`
	PluginIDs      []string  `json:"pluginIds"`
	CreatedBy      string    `json:"createdBy"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

type Pagination struct {
	Next  string `json:"next"`
	Limit int    `json:"limit"`
}
