# OnDemand API Client - Golang
Official golang client package for on-demand API.

This is the source code for the Go client library for <a href="https://docs.on-demand.io/reference/intro-to-ondemand-api" target="_blank">OnDemand API</a>.

The versioning scheme of this library is inspired by <a href="https://semver.org/" target="_blank">SemVer</a> and the format is v{MAJOR}.{MINOR}.{PATCH}. For example, v3.0.0 and v2.5.1 are valid library versions.

### Table of contents

1. Installation.
2. How to use the library ?
3. Using the chat API.
4. Streaming response from chat API.
5. Error Handling.

## Install the library
Install the latest version of the library with the following commands:

`go get github.com/dinson/ond-api-client-go`

## Use the library

Some examples on how to use the client library is listed below:

### Initialise the API client

```go
package main

import (
    "github.com/dinson/ond-api-client-go/ond"
)

func main() {
	ondAPIClient := ond.Init("your-ondemand-api-key-here")
}
```

### Use Chat API

Chat API supports sync, webhook and streaming response modes. You can choose between sync and webhook modes by passing the appropriate values in the `params.QueryParams.ResponseMode` field when calling the `chat.Query()` method.

Sync mode 
 - Return the query response synchronously.
 - Use `params.ResponseModeSync`

Webhook mode 
- Return the query response to your configured webhook.
- Use `params.ResponseModeWebhook`
- Configure webhook via the OnDemand dashboard - https://app.on-demand.io/settings/webhooks

#### Example usage with sync mode
```go
package main

import (
	"context"
	"fmt"
	"github.com/dinson/ond-api-client-go/ond"
	"github.com/dinson/ond-api-client-go/ond/params"
)

func main() {
	ctx := context.Background()
	
	ondAPIClient := ond.Init("your-ondemand-api-key-here")
    
	// create chat session
	session, _ := ondAPIClient.Chat.CreateSession(ctx, &params.CreateChatSessionParams{...})
	
        // call query endpoint to receive answer to a question
	q, _ := ondAPIClient.Chat.Query(ctx, &params.QueryParams{
		SessionID:  session.Data.ID,
		EndpointID: "predefined-gemini-1.5-pro", // refer https://docs.on-demand.io/docs/fulfillment-models#predefined-models
		Query:      "How are you today",
		ResponseMode: params.ResponseModeSync,  // alternatively, use params.ResponseModeWebhook
	})
    
	fmt.Println(q.Data.Answer)  // answer to your query from the LLM model
	
}
```

### Streaming mode for chat API
- You can also consume chat API response in streaming mode via SSE events.
- Use `chat.OpenStream()` to create an SSE connection and then use `Consume()` method to start consuming the streaming response.
- Refer to the example shown below on how to implement streaming response mode.

#### Example usage with streaming mode
```go
package main

import (
	"context"
	"fmt"
	"github.com/dinson/ond-api-client-go/ond"
	"github.com/dinson/ond-api-client-go/ond/params"
)

func main() {
	ctx := context.Background()
	
	ondAPIClient := ond.Init("your-ondemand-api-key-here")
    
	// create chat session
	session, _ := ondAPIClient.Chat.CreateSession(ctx, &params.CreateChatSessionParams{...})
	
        // call query endpoint to receive answer to a question
	stream, _ := ondAPIClient.Chat.OpenStream(ctx, &params.QueryParams{...})

	go stream.Consume() // start the `Consume()` in a go routine to prevent blocking

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()

		// responses are made available through the `stream.EventChan` channel
		for e := range stream.EventChan {
			if e.Done {
				// do operations to be conducted after consuming the whole response
				break
			}
			if e.Error != nil {
				// handle error
				break
			}

			fmt.Print(e.Data.Answer)
		}
	}()

	wg.Wait()
	
}
```

### Error Handling
OnDemand API return error status, error code and an optional message when encountering errors.

Every error returned from the API client library will be an object with the following fields:
```go
type ErrResponse struct {
    Message   string 
    ErrorCode string 
    Status    int
}
```

You can use the error object just like you would use the `error` package in go.
Example:
```go
func main() {
	...
	_, err := ondAPIClient.Chat.Query(...)
	if err != nil {
	    fmt.Println(err.Error())	
    }
}
```
The `err.Error()` method will return a native error package `err`.

For detailed documentation on error handling, please visit: https://docs.on-demand.io/reference/errors
