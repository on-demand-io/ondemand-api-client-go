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

type GetMessageResponse struct{}

type ListMessagesResponse struct{}

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
