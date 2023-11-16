<!-- Start SDK Example Usage -->
```go
package main

import (
	"context"
	naumanali "github.com/speakeasy-sdks/naumanali/v2"
	"github.com/speakeasy-sdks/naumanali/v2/pkg/models/operations"
	"github.com/speakeasy-sdks/naumanali/v2/pkg/models/shared"
	"log"
)

func main() {
	s := naumanali.New(
		naumanali.WithSecurity(""),
	)

	ctx := context.Background()
	res, err := s.ExportFileByBranch(ctx, operations.ExportFileByBranchRequest{
		StoplightVersion: operations.StoplightAPIVersionStringTwoThousandAndTwentyTwo1205.ToPointer(),
		BranchName:       "string",
		FilePath:         "/etc/namedb/radian_southeast.csp",
		ProjectID:        "string",
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.TwoHundredApplicationJSONAny != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage -->