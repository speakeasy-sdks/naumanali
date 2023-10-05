<!-- Start SDK Example Usage -->


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
    res, err := s.ExportFileByBranch(ctx, operations.ExportFileByBranchRequest{
        StoplightVersion: operations.ExportFileByBranchStoplightVersionStoplightAPIVersionStringTwoThousandAndTwentyTwo1205.ToPointer(),
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
<!-- End SDK Example Usage -->