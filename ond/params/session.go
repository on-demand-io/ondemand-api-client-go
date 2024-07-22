package params

type CreateChatSessionParams struct {
	// ExternalUserID - Required
	// An identifier of the external user creating this chat session.
	// This user is external to OnDemand but internal to your own system which can be used for filtering sessions and auditing.
	// If not managing chat users internally, use any unique string.
	ExternalUserID string `json:"externalUserId"`

	// PluginIDs - Optional
	// A list of plugin IDs to be used in the chat session.
	// A maximum of 10 plugins are allowed. This list can be empty.
	// if set here and not overwritten through /query endpoint, then these plugins will be used for all queries in this session.
	PluginIDs []string `json:"pluginIds"`
}

// ?externalUserId=ddk&sort=asc&cursor=cu&limit=10

type ListSessionsParams struct {
	ExternalUserID string `url:"externalUserId"`
	Sort           Sort   `url:"sort"`
	Limit          int32  `url:"limit"`
	Cursor         string `url:"cursor"`
}

type Sort string

var (
	SortAsc  Sort = "asc"
	SortDesc Sort = "desc"
)

func (s Sort) String() string { return string(s) }
