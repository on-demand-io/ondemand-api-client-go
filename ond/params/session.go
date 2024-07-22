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

type ListSessionParams struct {
	// ExternalUserID - Optional
	// An identifier of the external user creating this chat session.
	// This user is external to OnDemand but internal to your own system which can be used for filtering sessions and auditing.
	// If not managing chat users internally, use any unique string.
	ExternalUserID string `url:"externalUserId"`

	// Sort order of the sessions based on its creation time
	Sort Sort `url:"sort"`

	// Limit
	// Specifies the total number of results to retrieve. If the provided value is less than the minimum allowed limit,
	// it will be reset to the default value. Conversely, if the provided value exceeds the maximum allowed limit,
	// it will be reset to the maximum value.
	Limit int32 `url:"limit"`

	// Cursor
	// It acts as a pagination key and used to retrieve next results.
	//For first iteration, this parameter must not be set.
	//For subsequent iterations, it must be set to "pagination.next" value from previous response.
	Cursor string `url:"cursor"`
}

type Sort string

var (
	// SortAsc sorts according to ascending order of values
	SortAsc Sort = "asc"
	// SortDesc sorts according to descending order of values
	SortDesc Sort = "desc"
)

func (s Sort) String() string { return string(s) }
