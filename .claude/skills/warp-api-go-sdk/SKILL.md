---
name: warp-api-go-sdk
description: "Go SDK for Warp API. Use when writing Go code that calls Warp API with the warp-hr package: installing it, constructing and authenticating the client, and calling API operations."
---

# Warp API Go SDK

Generated Go client for Warp API, published as `warp-hr`. Use the generated client instead of hand-writing HTTP requests.

## Install

```sh
go get warp-hr
```

## Client setup and authentication

```go
import (
	"context"
	"fmt"

	sdk "warp-hr"
)

client := sdk.NewClient()
```

Provide credentials using the options below. Environment variables are read automatically when the target runtime supports them:

- `option.WithAPIKey` (env: `API_KEY`) — The API key for header authorization.

## Calling operations

```go
package main

import (
	"context"
	"fmt"
	"os"

	sdk "warp-hr"
	"warp-hr/option"
)

func main() {
	client := sdk.NewClient(
		option.WithAPIKey(os.Getenv("API_KEY")),
	)

	customWorkerField, err := client.CustomWorkerFields.List(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(customWorkerField)
}
```

Method names, parameter shapes, and response types are generated from the API description — do not guess them. Look up the exact call signature in [api.md](../../../api.md) before writing a call.

## Error handling

Non-success responses return generated API errors. Error objects expose status, headers, response body, and request metadata where the target runtime supports it.

```go
customWorkerField, err := client.CustomWorkerFields.List(context.Background())
if err != nil {
	var apiErr *sdk.Error
	if errors.As(err, &apiErr) {
		fmt.Println(apiErr.StatusCode, apiErr.RawJSON())
	}
	panic(err)
}

// imports: sdk "warp-hr", "errors", "fmt"
```

## Requirements

- Go 1.22 or newer

## Reference files

- [README.md](../../../README.md) — full feature tour: client options, request options, retries and timeouts, logging.
- [api.md](../../../api.md) — complete catalogue of every operation with request and response types.
