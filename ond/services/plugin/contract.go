package plugin

import "github.com/dinson/ond-api-client-go/ond/model"

type ListRequest struct {
	PluginIDs []string `url:"pluginIds"`
	Page      int      `url:"page"`
	Limit     int      `url:"limit"`
}

type ListResponse struct {
	Data Data `json:"data"`
}

type Data struct {
	Plugins []*model.Plugin `json:"plugins"`
}
