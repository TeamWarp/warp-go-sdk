---
name: warp-go-sdk
description: "Go SDK for warp API. Use when writing Go code that calls warp API with the github.com/TeamWarp/warp-go-sdk package: installing it, constructing and authenticating the client, and calling API operations."
---

# warp Go SDK

Generated Go client for warp API, published as `github.com/TeamWarp/warp-go-sdk`. Use the generated client instead of hand-writing HTTP requests.

## Install

```sh
go get github.com/TeamWarp/warp-go-sdk
```

## Client setup and authentication

```go
import (
	"context"
	"fmt"

	sdk "github.com/TeamWarp/warp-go-sdk"
)

client := sdk.NewClient()
```

Provide credentials using the options below. Environment variables are read automatically when the target runtime supports them:

- `option.WithAPIKey` (env: `WARP_API_KEY`) — The API key for header authorization.

## Calling operations

```go
package main

import (
	"context"
	"fmt"
	"os"

	sdk "github.com/TeamWarp/warp-go-sdk"
	"github.com/TeamWarp/warp-go-sdk/option"
)

func main() {
	client := sdk.NewClient(
		option.WithAPIKey(os.Getenv("WARP_API_KEY")),
	)

	healthPlan, err := client.Benefits.HealthPlans.List(context.Background(), sdk.BenefitHealthPlanListParams{
		Limit:    sdk.F[string]("limit"),
		Statuses: sdk.F[[]sdk.BenefitHealthPlanListParamsStatus]([]sdk.BenefitHealthPlanListParamsStatus{"active"}),
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(healthPlan)
}
```

Method names, parameter shapes, and response types are generated from the API description — do not guess them. Look up the exact call signature in [api.md](./api.md) before writing a call.

## Error handling

Non-success responses return generated API errors. Error objects expose status, headers, response body, and request metadata where the target runtime supports it.

```go
healthPlan, err := client.Benefits.HealthPlans.List(context.Background(), sdk.BenefitHealthPlanListParams{
	Limit:    sdk.F[string]("limit"),
	Statuses: sdk.F[[]sdk.BenefitHealthPlanListParamsStatus]([]sdk.BenefitHealthPlanListParamsStatus{"active"}),
})
if err != nil {
	var apiErr *sdk.Error
	if errors.As(err, &apiErr) {
		fmt.Println(apiErr.StatusCode, apiErr.RawJSON())
	}
	panic(err)
}

// imports: "context", "errors", "fmt", sdk "github.com/TeamWarp/warp-go-sdk"
```

## Requirements

- Go 1.22 or newer

## Reference files

- [README.md](./README.md) — full feature tour: client options, request options, retries and timeouts, logging.
- [api.md](./api.md) — complete catalogue of every operation with request and response types.
