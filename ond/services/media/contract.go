package media

import "time"

type CreateMediaResponse struct {
	Data Data `json:"data"`
}

type Data struct {
	ID               string        `json:"_id"`
	CompanyID        string        `json:"companyId"`
	SessionID        string        `json:"sessionId"`
	URL              string        `json:"url"`
	SourceURL        string        `json:"sourceUrl"`
	ExtractedTextURL string        `json:"extractedTextUrl"`
	Name             string        `json:"name"`
	SizeInBytes      int           `json:"sizeBytes"`
	Source           string        `json:"source"`
	MimeType         string        `json:"mimeType"`
	Extension        string        `json:"extension"`
	Plugins          []string      `json:"plugins"`
	Context          string        `json:"context"`
	ExtractedText    interface{}   `json:"extractedText"`
	ActionStatus     string        `json:"actionStatus"`
	FailedReason     interface{}   `json:"failedReason"`
	IsDeleted        bool          `json:"isDeleted"`
	ResponseMode     string        `json:"responseMode"`
	CreatedBy        string        `json:"createdBy"`
	UpdatedBy        string        `json:"updatedBy"`
	PluginInputs     []PluginInput `json:"pluginInputs"`
	MediaStatsID     string        `json:"mediaStatsId"`
	CreatedAt        time.Time     `json:"createdAt"`
	UpdatedAt        time.Time     `json:"updatedAt"`
}

type PluginInput struct {
	AdditionalProperties AdditionalProperties `json:"additionalProp"`
}

type AdditionalProperties struct {
	PostProcess PostProcess `json:"postProcess"`
}

type PostProcess struct {
	ChatPluginID string `json:"chatPluginId"`
}

type FetchMediaResponse struct {
	Data []Data `json:"data"`
}
