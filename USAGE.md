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
		BranchName:       "Auto",
		FilePath:         "/etc/defaults/sausages_east.c4p",
		ProjectID:        "sievert Applications Bike",
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