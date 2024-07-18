package plugin

import "airevai/ondemand-api/ond/model"

type ListRequest struct {
	PluginIDs []string `form:"pluginId"`
	Page      int      `form:"page"`
	Limit     int      `form:"limit"`
}

type ListResponse struct {
	Data Data `json:"data"`
}

type Data struct {
	Plugins []*model.Plugin `json:"plugins"`
}
