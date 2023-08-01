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
    s := apitest.New()
    operationSecurity := operations.ExportFileByBranchSecurity{
            Authorization: "",
        }

    ctx := context.Background()
    res, err := s.ExportFileByBranch(ctx, operations.ExportFileByBranchRequest{
        StoplightVersion: operations.ExportFileByBranchStoplightVersionStoplightAPIVersionStringTwoThousandAndTwentyTwo1205.ToPointer(),
        BranchName: "corrupti",
        FilePath: "provident",
        IncludeInternal: apitest.Bool(false),
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