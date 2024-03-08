<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	naumanali "github.com/speakeasy-sdks/naumanali/v3"
	"github.com/speakeasy-sdks/naumanali/v3/pkg/models/operations"
	"github.com/speakeasy-sdks/naumanali/v3/pkg/models/shared"
	"log"
)

func main() {
	s := naumanali.New(
		naumanali.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)

	ctx := context.Background()
	res, err := s.ExportFileByBranch(ctx, operations.ExportFileByBranchRequest{
		StoplightVersion: operations.StoplightAPIVersionStringTwoThousandAndTwentyTwo1205.ToPointer(),
		BranchName:       "<value>",
		FilePath:         "/opt/lib/bronze_override_yuck.rmvb",
		ProjectID:        "<value>",
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.TwoHundredApplicationJSONAny != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->