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
	"context"
	"log"
	"github.com/speakeasy-sdks/naumanali"
	"github.com/speakeasy-sdks/naumanali/pkg/models/shared"
	"github.com/speakeasy-sdks/naumanali/pkg/models/operations"
)

func main() {
    s := naumanali.New(
        naumanali.WithSecurity(shared.Security{
            Authorization: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Apitest.ExportFileByBranch(ctx, operations.ExportFileByBranchRequest{
        StoplightVersion: naumanali.String("2022-12-05"),
        BranchName: "Auto",
        FilePath: "/etc/defaults/sausages_east.c4p",
        IncludeInternal: naumanali.Bool(false),
        ProjectID: "sievert Applications Bike",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ExportFileByBranch200ApplicationJSONAny != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ExportFileByBranchRequest](../../models/operations/exportfilebybranchrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.ExportFileByBranchResponse](../../models/operations/exportfilebybranchresponse.md), error**


## ExportFileByCommit

Exports an OpenAPI or JSON Schema file from a Stoplight project that exists on the specified commit.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/naumanali"
	"github.com/speakeasy-sdks/naumanali/pkg/models/shared"
	"github.com/speakeasy-sdks/naumanali/pkg/models/operations"
)

func main() {
    s := naumanali.New(
        naumanali.WithSecurity(shared.Security{
            Authorization: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Apitest.ExportFileByCommit(ctx, operations.ExportFileByCommitRequest{
        StoplightVersion: naumanali.String("2022-12-05"),
        CommitHash: "Frozen",
        FilePath: "/boot/sore_account_east.oxt",
        IncludeInternal: naumanali.Bool(false),
        ProjectID: "Cambridgeshire Classical",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ExportFileByCommit200ApplicationJSONAny != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ExportFileByCommitRequest](../../models/operations/exportfilebycommitrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.ExportFileByCommitResponse](../../models/operations/exportfilebycommitresponse.md), error**

