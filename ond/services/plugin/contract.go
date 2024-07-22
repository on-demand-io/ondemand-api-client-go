package plugin

import "github.com/dinson/ond-api-client-go/ond/model"

type ListResponse struct {
	Data Data `json:"data"`
}

type Data struct {
	Plugins []*model.Plugin `json:"plugins"`
}
