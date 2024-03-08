# Apitest SDK


## Overview

Stoplight Catalog API: The Catalog API provides [bundled exports](https://docs.stoplight.io/docs/platform/37d160068e33c-export-api-files#ref-options) of your OpenAPI and JSON Schema files. A bundled export resolves remote `$refs` once, and then reuses references to the same objects to avoid repetition and produce a cleaner file.

Use Catalog API exports to:

- Facilitate automated testing
- Update internal API Directories
- Update a locally hosted [Prism  mock server](https://github.com/stoplightio/prism)
- Avoid direct Git access

### Available Operations

* [ExportFileByBranch](#exportfilebybranch) - exportFileByBranch
* [ExportFileByCommit](#exportfilebycommit) - exportFileByCommit

## ExportFileByBranch

Exports an OpenAPI or JSON Schema file from a Stoplight project that exists on the specified branch.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/naumanali/v3/pkg/models/shared"
	naumanali "github.com/speakeasy-sdks/naumanali/v3"
	"context"
	"github.com/speakeasy-sdks/naumanali/v3/pkg/models/operations"
	"log"
)

func main() {
    s := naumanali.New(
        naumanali.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.ExportFileByBranch(ctx, operations.ExportFileByBranchRequest{
        StoplightVersion: operations.StoplightAPIVersionStringTwoThousandAndTwentyTwo1205.ToPointer(),
        BranchName: "<value>",
        FilePath: "/opt/lib/bronze_override_yuck.rmvb",
        ProjectID: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TwoHundredApplicationJSONAny != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.ExportFileByBranchRequest](../../pkg/models/operations/exportfilebybranchrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.ExportFileByBranchResponse](../../pkg/models/operations/exportfilebybranchresponse.md), error**
| Error Object                               | Status Code                                | Content Type                               |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| sdkerrors.BadRequestProblemDetail          | 400                                        | application/problem+json                   |
| sdkerrors.UnauthorizedProblemDetail        | 401                                        | application/problem+json                   |
| sdkerrors.PaymentRequiredProblemDetail     | 402                                        | application/problem+json                   |
| sdkerrors.ForbiddenProblemDetail           | 403                                        | application/problem+json                   |
| sdkerrors.NotFoundProblemDetail            | 404                                        | application/problem+json                   |
| sdkerrors.ConflictProblemDetail            | 409                                        | application/problem+json                   |
| sdkerrors.ExportFileFailedProblemDetail    | 422                                        | application/problem+json                   |
| sdkerrors.TooManyRequestsProblemDetail     | 429                                        | application/problem+json                   |
| sdkerrors.InternalServerErrorProblemDetail | 500                                        | application/problem+json                   |
| sdkerrors.SDKError                         | 4xx-5xx                                    | */*                                        |

## ExportFileByCommit

Exports an OpenAPI or JSON Schema file from a Stoplight project that exists on the specified commit.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/naumanali/v3/pkg/models/shared"
	naumanali "github.com/speakeasy-sdks/naumanali/v3"
	"context"
	"github.com/speakeasy-sdks/naumanali/v3/pkg/models/operations"
	"log"
)

func main() {
    s := naumanali.New(
        naumanali.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.ExportFileByCommit(ctx, operations.ExportFileByCommitRequest{
        StoplightVersion: operations.HeaderStoplightAPIVersionStringTwoThousandAndTwentyTwo1205.ToPointer(),
        CommitHash: "<value>",
        FilePath: "/lib/jose_jewelery_east.oxt",
        ProjectID: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TwoHundredApplicationJSONAny != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.ExportFileByCommitRequest](../../pkg/models/operations/exportfilebycommitrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.ExportFileByCommitResponse](../../pkg/models/operations/exportfilebycommitresponse.md), error**
| Error Object                               | Status Code                                | Content Type                               |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| sdkerrors.BadRequestProblemDetail          | 400                                        | application/problem+json                   |
| sdkerrors.UnauthorizedProblemDetail        | 401                                        | application/problem+json                   |
| sdkerrors.PaymentRequiredProblemDetail     | 402                                        | application/problem+json                   |
| sdkerrors.ForbiddenProblemDetail           | 403                                        | application/problem+json                   |
| sdkerrors.NotFoundProblemDetail            | 404                                        | application/problem+json                   |
| sdkerrors.ConflictProblemDetail            | 409                                        | application/problem+json                   |
| sdkerrors.ExportFileFailedProblemDetail    | 422                                        | application/problem+json                   |
| sdkerrors.TooManyRequestsProblemDetail     | 429                                        | application/problem+json                   |
| sdkerrors.InternalServerErrorProblemDetail | 500                                        | application/problem+json                   |
| sdkerrors.SDKError                         | 4xx-5xx                                    | */*                                        |
