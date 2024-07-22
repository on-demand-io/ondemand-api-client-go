package chat

import "time"

type SubmitQueryResponse struct{}

type GetSessionResponse struct{}

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
