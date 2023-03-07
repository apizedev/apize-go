# apize-go

The official [Apize](https://apize.dev) SDK for the Go programming language, for adding AI-powered intelligence to your apps.

## Overview

This module provides access to the [Apize APIs](https://docs.apize.dev) from your Go code, handling authentication and the processing of API requests.

## Summarization example

```go
import "github.com/apizedev/apize-go"

// Create an Apize client with your API token
client := apize.New("token")

// Summarize some text
ctx := context.Background()
res, _ := client.Summarize(ctx, &apize.SummarizeRequest{
    Text: "The quick brown fox jumped over the lazy dog",
})
fmt.Println("Summary: ", res.Summary)
```

## Resources

-   Read our [Quickstart guide here](https://docs.apize.dev/quickstart)
-   Read the [API docs here](https://docs.apize.dev/text-intelligence)
-   Setup your [API tokens here](https://apize.dev/tokens)
