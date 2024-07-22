package params

type ListPluginParams struct {
	PluginIDs []string `url:"pluginIds"`
	Page      int      `url:"page"`
	Limit     int      `url:"limit"`
}
