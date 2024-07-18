package model

import "time"

type Plugin struct {
	ID                   string    `json:"id"`
	Name                 string    `json:"name"`
	Identifier           string    `json:"identifier"`
	Description          string    `json:"description"`
	Category             string    `json:"category"`
	ConversationStarters []string  `json:"conversationStarters"`
	LogoURL              string    `json:"logoUrl"`
	Type                 string    `json:"type"`
	Source               string    `json:"source"`
	Status               string    `json:"status"`
	PluginID             string    `json:"pluginId"`
	CompanyID            string    `json:"companyId"`
	FileSubType          string    `json:"fileSubType"`
	ChatSubType          string    `json:"chatSubType"`
	PrivacyPolicy        string    `json:"privacyPolicy"`
	CreatedAt            time.Time `json:"createdAt"`
	UpdatedAt            time.Time `json:"updatedAt"`
}
