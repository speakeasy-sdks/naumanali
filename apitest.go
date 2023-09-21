// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package naumanali

import (
	"bytes"
	"context"
	"fmt"
	"github.com/speakeasy-sdks/naumanali/pkg/models/operations"
	"github.com/speakeasy-sdks/naumanali/pkg/models/sdkerrors"
	"github.com/speakeasy-sdks/naumanali/pkg/models/shared"
	"github.com/speakeasy-sdks/naumanali/pkg/utils"
	"io"
	"net/http"
	"time"
)

// ServerList contains the list of servers available to the SDK
var ServerList = []string{
	"https://api.stoplight.io",
}

// HTTPClient provides an interface for suplying the SDK with a custom HTTP client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// String provides a helper function to return a pointer to a string
func String(s string) *string { return &s }

// Bool provides a helper function to return a pointer to a bool
func Bool(b bool) *bool { return &b }

// Int provides a helper function to return a pointer to an int
func Int(i int) *int { return &i }

// Int64 provides a helper function to return a pointer to an int64
func Int64(i int64) *int64 { return &i }

// Float32 provides a helper function to return a pointer to a float32
func Float32(f float32) *float32 { return &f }

// Float64 provides a helper function to return a pointer to a float64
func Float64(f float64) *float64 { return &f }

type sdkConfiguration struct {
	DefaultClient  HTTPClient
	SecurityClient HTTPClient

	ServerURL         string
	ServerIndex       int
	Language          string
	OpenAPIDocVersion string
	SDKVersion        string
	GenVersion        string
	RetryConfig       *utils.RetryConfig
}

func (c *sdkConfiguration) GetServerDetails() (string, map[string]string) {
	if c.ServerURL != "" {
		return c.ServerURL, nil
	}

	return ServerList[c.ServerIndex], nil
}

// Apitest - Stoplight Catalog API: The Catalog API provides [bundled exports](https://docs.stoplight.io/docs/platform/37d160068e33c-export-api-files#ref-options) of your OpenAPI and JSON Schema files. A bundled export resolves remote `$refs` once, and then reuses references to the same objects to avoid repetition and produce a cleaner file.
//
// Use Catalog API exports to:
//
// - Facilitate automated testing
// - Update internal API Directories
// - Update a locally hosted [Prism  mock server](https://github.com/stoplightio/prism)
// - Avoid direct Git access
type Apitest struct {
	sdkConfiguration sdkConfiguration
}

type SDKOption func(*Apitest)

// WithServerURL allows the overriding of the default server URL
func WithServerURL(serverURL string) SDKOption {
	return func(sdk *Apitest) {
		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithTemplatedServerURL allows the overriding of the default server URL with a templated URL populated with the provided parameters
func WithTemplatedServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *Apitest) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithServerIndex allows the overriding of the default server by index
func WithServerIndex(serverIndex int) SDKOption {
	return func(sdk *Apitest) {
		if serverIndex < 0 || serverIndex >= len(ServerList) {
			panic(fmt.Errorf("server index %d out of range", serverIndex))
		}

		sdk.sdkConfiguration.ServerIndex = serverIndex
	}
}

// WithClient allows the overriding of the default HTTP client used by the SDK
func WithClient(client HTTPClient) SDKOption {
	return func(sdk *Apitest) {
		sdk.sdkConfiguration.DefaultClient = client
	}
}

func WithRetryConfig(retryConfig utils.RetryConfig) SDKOption {
	return func(sdk *Apitest) {
		sdk.sdkConfiguration.RetryConfig = &retryConfig
	}
}

// New creates a new instance of the SDK with the provided options
func New(opts ...SDKOption) *Apitest {
	sdk := &Apitest{
		sdkConfiguration: sdkConfiguration{
			Language:          "go",
			OpenAPIDocVersion: "2022-12-05",
			SDKVersion:        "1.2.1",
			GenVersion:        "2.122.1",
		},
	}
	for _, opt := range opts {
		opt(sdk)
	}

	// Use WithClient to override the default client if you would like to customize the timeout
	if sdk.sdkConfiguration.DefaultClient == nil {
		sdk.sdkConfiguration.DefaultClient = &http.Client{Timeout: 60 * time.Second}
	}
	if sdk.sdkConfiguration.SecurityClient == nil {
		sdk.sdkConfiguration.SecurityClient = sdk.sdkConfiguration.DefaultClient
	}

	return sdk
}

// ExportFileByBranch - exportFileByBranch
// Exports an OpenAPI or JSON Schema file from a Stoplight project that exists on the specified branch.
func (s *Apitest) ExportFileByBranch(ctx context.Context, request operations.ExportFileByBranchRequest, security operations.ExportFileByBranchSecurity, opts ...operations.Option) (*operations.ExportFileByBranchResponse, error) {
	o := operations.Options{}
	supportedOptions := []string{
		operations.SupportedOptionAcceptHeaderOverride,
	}

	for _, opt := range opts {
		if err := opt(&o, supportedOptions...); err != nil {
			return nil, fmt.Errorf("error applying option: %w", err)
		}
	}
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/projects/{project_id}/branches/{branch_name}/export/{file_path}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	if o.AcceptHeaderOverride != nil {
		req.Header.Set("Accept", string(*o.AcceptHeaderOverride))
	} else {
		req.Header.Set("Accept", "application/json;q=1, application/problem+json;q=0.7, application/yaml;q=0")
	}

	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s %s", s.sdkConfiguration.Language, s.sdkConfiguration.SDKVersion, s.sdkConfiguration.GenVersion, s.sdkConfiguration.OpenAPIDocVersion))

	utils.PopulateHeaders(ctx, req, request)

	if err := utils.PopulateQueryParams(ctx, req, request, nil); err != nil {
		return nil, fmt.Errorf("error populating query params: %w", err)
	}

	client := utils.ConfigureSecurityClient(s.sdkConfiguration.DefaultClient, security)

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.ExportFileByBranchResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out interface{}
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.ExportFileByBranch200ApplicationJSONAny = out
		case utils.MatchContentType(contentType, `application/yaml`):
			res.Body = rawBody
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 400:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(contentType, `application/problem+json`):
			var out shared.BadRequestProblemDetail
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.BadRequestProblemDetail = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 401:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(contentType, `application/problem+json`):
			var out shared.UnauthorizedProblemDetail
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.UnauthorizedProblemDetail = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 402:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(contentType, `application/problem+json`):
			var out shared.PaymentRequiredProblemDetail
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.PaymentRequiredProblemDetail = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 403:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(contentType, `application/problem+json`):
			var out shared.ForbiddenProblemDetail
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.ForbiddenProblemDetail = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 404:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(contentType, `application/problem+json`):
			var out shared.NotFoundProblemDetail
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.NotFoundProblemDetail = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 409:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(contentType, `application/problem+json`):
			var out shared.ConflictProblemDetail
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.ConflictProblemDetail = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 422:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(contentType, `application/problem+json`):
			var out shared.ExportFileFailedProblemDetail
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.ExportFileFailedProblemDetail = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 429:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(contentType, `application/problem+json`):
			var out shared.TooManyRequestsProblemDetail
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.TooManyRequestsProblemDetail = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 500:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(contentType, `application/problem+json`):
			var out shared.InternalServerErrorProblemDetail
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.InternalServerErrorProblemDetail = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	}

	return res, nil
}

// ExportFileByCommit - exportFileByCommit
// Exports an OpenAPI or JSON Schema file from a Stoplight project that exists on the specified commit.
func (s *Apitest) ExportFileByCommit(ctx context.Context, request operations.ExportFileByCommitRequest, security operations.ExportFileByCommitSecurity, opts ...operations.Option) (*operations.ExportFileByCommitResponse, error) {
	o := operations.Options{}
	supportedOptions := []string{
		operations.SupportedOptionAcceptHeaderOverride,
	}

	for _, opt := range opts {
		if err := opt(&o, supportedOptions...); err != nil {
			return nil, fmt.Errorf("error applying option: %w", err)
		}
	}
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/projects/{project_id}/commits/{commit_hash}/export/{file_path}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	if o.AcceptHeaderOverride != nil {
		req.Header.Set("Accept", string(*o.AcceptHeaderOverride))
	} else {
		req.Header.Set("Accept", "application/json;q=1, application/problem+json;q=0.7, application/yaml;q=0")
	}

	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s %s", s.sdkConfiguration.Language, s.sdkConfiguration.SDKVersion, s.sdkConfiguration.GenVersion, s.sdkConfiguration.OpenAPIDocVersion))

	utils.PopulateHeaders(ctx, req, request)

	if err := utils.PopulateQueryParams(ctx, req, request, nil); err != nil {
		return nil, fmt.Errorf("error populating query params: %w", err)
	}

	client := utils.ConfigureSecurityClient(s.sdkConfiguration.DefaultClient, security)

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.ExportFileByCommitResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out interface{}
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.ExportFileByCommit200ApplicationJSONAny = out
		case utils.MatchContentType(contentType, `application/yaml`):
			res.Body = rawBody
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 400:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(contentType, `application/problem+json`):
			var out shared.BadRequestProblemDetail
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.BadRequestProblemDetail = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 401:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(contentType, `application/problem+json`):
			var out shared.UnauthorizedProblemDetail
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.UnauthorizedProblemDetail = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 402:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(contentType, `application/problem+json`):
			var out shared.PaymentRequiredProblemDetail
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.PaymentRequiredProblemDetail = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 403:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(contentType, `application/problem+json`):
			var out shared.ForbiddenProblemDetail
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.ForbiddenProblemDetail = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 404:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(contentType, `application/problem+json`):
			var out shared.NotFoundProblemDetail
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.NotFoundProblemDetail = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 409:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(contentType, `application/problem+json`):
			var out shared.ConflictProblemDetail
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.ConflictProblemDetail = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 422:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(contentType, `application/problem+json`):
			var out shared.ExportFileFailedProblemDetail
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.ExportFileFailedProblemDetail = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 429:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(contentType, `application/problem+json`):
			var out shared.TooManyRequestsProblemDetail
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.TooManyRequestsProblemDetail = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 500:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(contentType, `application/problem+json`):
			var out shared.InternalServerErrorProblemDetail
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.InternalServerErrorProblemDetail = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	}

	return res, nil
}
