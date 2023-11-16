# github.com/speakeasy-sdks/naumanali

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/naumanali
```
<!-- End SDK Installation -->

## SDK Example Usage
<!-- Start SDK Example Usage -->
### Example

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

<!-- Start SDK Available Operations -->
## Available Resources and Operations

### [Apitest SDK](docs/sdks/apitest/README.md)

* [ExportFileByBranch](docs/sdks/apitest/README.md#exportfilebybranch) - exportFileByBranch
* [ExportFileByCommit](docs/sdks/apitest/README.md#exportfilebycommit) - exportFileByCommit
<!-- End SDK Available Operations -->



<!-- Start Dev Containers -->

<!-- End Dev Containers -->



<!-- Start Pagination -->
# Pagination

Some of the endpoints in this SDK support pagination. To use pagination, you make your SDK calls as usual, but the
returned response object will have a `Next` method that can be called to pull down the next group of results. If the
return value of `Next` is `nil`, then there are no more pages to be fetched.

Here's an example of one such pagination call:
<!-- End Pagination -->



<!-- Start Go Types -->

<!-- End Go Types -->



<!-- Start Error Handling -->
## Error Handling

Handling errors in this SDK should largely match your expectations.  All operations return a response object or an error, they will never return both.  When specified by the OpenAPI spec document, the SDK will return the appropriate subclass.

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
| sdkerrors.SDKError                         | 400-600                                    | */*                                        |

### Example

```go
package main

import (
	"context"
	"errors"
	naumanali "github.com/speakeasy-sdks/naumanali/v2"
	"github.com/speakeasy-sdks/naumanali/v2/pkg/models/operations"
	"github.com/speakeasy-sdks/naumanali/v2/pkg/models/sdkerrors"
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

		var e *sdkerrors.BadRequestProblemDetail
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.UnauthorizedProblemDetail
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.PaymentRequiredProblemDetail
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.ForbiddenProblemDetail
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.NotFoundProblemDetail
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.ConflictProblemDetail
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.ExportFileFailedProblemDetail
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.TooManyRequestsProblemDetail
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.InternalServerErrorProblemDetail
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.SDKError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling -->



<!-- Start Server Selection -->
## Server Selection

### Select Server by Index

You can override the default server globally using the `WithServerIndex` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the indexes associated with the available servers:

| # | Server | Variables |
| - | ------ | --------- |
| 0 | `https://api.stoplight.io` | None |

#### Example

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
		naumanali.WithServerIndex(0),
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


### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL` option when initializing the SDK client instance. For example:
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
		naumanali.WithServerURL("https://api.stoplight.io"),
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
<!-- End Server Selection -->



<!-- Start Custom HTTP Client -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client -->



<!-- Start Authentication -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security scheme globally:

| Name            | Type            | Scheme          |
| --------------- | --------------- | --------------- |
| `Authorization` | http            | HTTP Bearer     |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	naumanali "github.com/speakeasy-sdks/naumanali/v2"
	"github.com/speakeasy-sdks/naumanali/v2/pkg/models/operations"
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
<!-- End Authentication -->

<!-- Placeholder for Future Speakeasy SDK Sections -->



### Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
