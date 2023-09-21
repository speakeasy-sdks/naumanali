# github.com/speakeasy-sdks/naumanali

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/naumanali
```
<!-- End SDK Installation -->

## SDK Example Usage
<!-- Start SDK Example Usage -->


```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/naumanali"
	"github.com/speakeasy-sdks/naumanali/pkg/models/operations"
)

func main() {
    s := naumanali.New()
    operationSecurity := operations.ExportFileByBranchSecurity{
            Authorization: "",
        }

    ctx := context.Background()
    res, err := s.ExportFileByBranch(ctx, operations.ExportFileByBranchRequest{
        StoplightVersion: naumanali.String("2022-12-05"),
        BranchName: "corrupti",
        FilePath: "provident",
        IncludeInternal: naumanali.Bool(false),
        ProjectID: "distinctio",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.ExportFileByBranch200ApplicationJSONAny != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## Available Resources and Operations

### [Apitest SDK](docs/sdks/apitest/README.md)

* [ExportFileByBranch](docs/sdks/apitest/README.md#exportfilebybranch) - exportFileByBranch
* [ExportFileByCommit](docs/sdks/apitest/README.md#exportfilebycommit) - exportFileByCommit
<!-- End SDK Available Operations -->

### Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
