package params

type QueryParams struct {
	// SessionID - Required
	// ID of the session linked to this query
	SessionID string

	// EndpointID - Required
	// Endpoint ID of the fulfillment model selected to fulfill the query.
	// This can be a predefined, BYOI or BYOM model endpoint.
	// You can get a list of all predefined models at Fulfilllment Models.
	// To get a list of BYOI and BYOM models, please refer to BYOI and BYOM respectively.
	// Refer https://docs.on-demand.io/docs/fulfillment-models#predefined-models to get the IDs
	EndpointID string `json:"endpointId"`

	// ResponseMode - Required
	// Response mode to get the query answer.
	// Value can be "sync", "stream" or "webhook"
	// Refer https://docs.on-demand.io/docs/chat-queries-and-responses-modes
	ResponseMode ResponseMode `json:"responseMode"`

	// ModelConfig
	// Sets fulfillment model configuration.
	// If not passed, default configuration will be used to configure the fulfillment model referenced by endpointId parameter
	ModelConfig ModelConfig `json:"modelConfigs"`

	// Query - Required
	// Question to ask the AI
	Query string `json:"query"`

	// PluginIDs
	// A list of plugin IDs to be made accessible to RAG to answer the query. A maximum of 10 plugins are allowed.
	// This list can be empty. If specified, it will replace the plugin IDs set during session creation.
	// If not specified or left empty, the plugin IDs set during session creation will be utilized.
	// If not specified at any level, the system will bypass the RAG and proceed directly to execute the fulfillment.
	PluginIDs []string `json:"pluginIds"`

	// OnlyFulfillment
	// If set to true, skips the RAG and only executes the fulfillment even if pluginIds parameter is set at any level.
	// Skipping RAG will also skip the plugins execution so the RAG dependent queries will not be answered correctly.
	OnlyFulfillment bool `json:"onlyFulfillment"`
}

type ModelConfig struct {
	// FulfillmentPrompt
	// Prompt to set for fulfillment.
	// It must include two template variables in the format: Context: {context} and Question: {question}.
	// These template variables will be replaced automatically by OnDemand.
	// If not passed or empty, default prompt will be used.
	FulfillmentPrompt string `json:"fulfillmentPrompt"`

	// StopSequences
	// Up to 4 sequences where the API will stop generating further tokens.
	StopSequences []string `json:"stopSequences"`

	// Temperature
	// What sampling temperature to use, between 0 and 2. Higher values like 0.8 will make the output more random,
	// while lower values like 0.2 will make it more focused and deterministic.
	// It is recommended to alter either this or topP but not both.
	// Defaults to 0.7
	Temperature int `json:"temperature"`

	// TopP
	// An alternative to sampling with temperature, called nucleus sampling,
	// where the model considers the results of the tokens with topP probability mass.
	// So 0.1 means only the tokens comprising the top 10% probability mass are considered.
	// It is recommended to alter either this or Temperature but not both.
	// Defaults to 1
	TopP int `json:"topP"`

	// PresencePenalty
	// Number between 0.0 and 2.0. Higher values penalize new tokens based on whether they appear in the text so far,
	// increasing the model's likelihood to talk about new topics.
	// Defaults to 0
	PresencePenalty int `json:"presencePenalty"`

	// FrequencyPenalty
	// Number between 0.0 and 2.0. Higher values penalize new tokens based on their existing frequency in the text so far,
	// decreasing the model's likelihood to repeat the same line verbatim.
	// Defaults to 0
	FrequencyPenalty int `json:"frequencyPenalty"`
}

type ResponseMode string

var (
	// ResponseModeSync
	// In sync mode, the query answer is provided synchronously.
	// This means all answer tokens are combined on the server and then sent back in the response once the process is complete.
	ResponseModeSync ResponseMode = "sync"

	// ResponseModeStream
	// when responseMode is set to stream. you get the answer through server-side events (SSE).
	ResponseModeStream ResponseMode = "stream"

	// ResponseModeWebhook
	// In webhook mode, the response is sent to a webhook URL that you can configure on the Webhook Settings page of the OnDemand dashboard.
	// Go to https://app.on-demand.io/settings/webhooks to manage your webhook configurations.
	ResponseModeWebhook ResponseMode = "webhook"
)

func (rm ResponseMode) String() string { return string(rm) }

// ChatStatus ... values can be "processing", "completed" or "failed"
type ChatStatus string

var (
	StatusProcession ChatStatus = "processing"
	StatusCompleted  ChatStatus = "completed"
	StatusFailed     ChatStatus = "failed"
)

func (cs ChatStatus) String() string { return string(cs) }
