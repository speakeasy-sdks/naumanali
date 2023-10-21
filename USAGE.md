<!-- Start SDK Example Usage -->


```go
package main

import (
	"context"
	"github.com/speakeasy-sdks/naumanali"
	"github.com/speakeasy-sdks/naumanali/pkg/models/operations"
	"github.com/speakeasy-sdks/naumanali/pkg/models/shared"
	"log"
)

func main() {
	s := naumanali.New(
		naumanali.WithSecurity(""),
	)

	ctx := context.Background()
	res, err := s.Apitest.ExportFileByBranch(ctx, operations.ExportFileByBranchRequest{
		StoplightVersion: operations.ExportFileByBranchStoplightVersionStoplightAPIVersionStringTwoThousandAndTwentyTwo1205.ToPointer(),
		BranchName:       "string",
		FilePath:         "/etc/namedb/radian_southeast.csp",
		ProjectID:        "string",
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