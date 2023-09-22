// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/naumanali/pkg/models/shared"
	"github.com/speakeasy-sdks/naumanali/pkg/types"
	"github.com/speakeasy-sdks/naumanali/pkg/utils"
	"net/http"
)

type ExportFileByCommitRequest struct {
	// A string representing the Stoplight API version that is being requested. If not supplied: TODO document policy
	stoplightVersion *string `const:"2022-12-05" header:"style=simple,explode=false,name=Stoplight-Version"`
	// A reference to a commit or change in a Stoplight project.
	CommitHash string `pathParam:"style=simple,explode=false,name=commit_hash"`
	// A path to a file in a Stoplight project. Use forward slashes (`/`) to separate path directories.
	FilePath string `pathParam:"style=simple,explode=false,name=file_path"`
	// When `false`, indicates that any declaration (schema, operation, property, etc.) in the OpenAPI or JSON Schema marked with the `x-internal` extension property set to `true` will be omitted from the exported file. Setting this parameter to `true` with an anonymous or workspace guest user caller results in an error response.
	IncludeInternal *bool `default:"true" queryParam:"style=form,explode=true,name=include_internal"`
	// A string that uniquely identifies a Stoplight Project resource. This can be found in the [Project Settings](https://docs.stoplight.io/docs/platform/252039ebe8fb2-project-settings) page.
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

func (e ExportFileByCommitRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(e, "", false)
}

func (e *ExportFileByCommitRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &e, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ExportFileByCommitRequest) GetStoplightVersion() *string {
	return types.String("2022-12-05")
}

func (o *ExportFileByCommitRequest) GetCommitHash() string {
	if o == nil {
		return ""
	}
	return o.CommitHash
}

func (o *ExportFileByCommitRequest) GetFilePath() string {
	if o == nil {
		return ""
	}
	return o.FilePath
}

func (o *ExportFileByCommitRequest) GetIncludeInternal() *bool {
	if o == nil {
		return nil
	}
	return o.IncludeInternal
}

func (o *ExportFileByCommitRequest) GetProjectID() string {
	if o == nil {
		return ""
	}
	return o.ProjectID
}

type ExportFileByCommitResponse struct {
	// The server could not understand the request due to invalid syntax.
	BadRequestProblemDetail *shared.BadRequestProblemDetail
	Body                    []byte
	// This response is sent when a request conflicts with the current state of the server.
	ConflictProblemDetail *shared.ConflictProblemDetail
	ContentType           string
	// The server could not export the file contents requested.
	ExportFileFailedProblemDetail *shared.ExportFileFailedProblemDetail
	// The client does not have permissions to access the content.
	ForbiddenProblemDetail *shared.ForbiddenProblemDetail
	Headers                map[string][]string
	// The server has encountered a situation it doesn't know how to handle.
	InternalServerErrorProblemDetail *shared.InternalServerErrorProblemDetail
	// The server can not find the requested resource.
	NotFoundProblemDetail *shared.NotFoundProblemDetail
	// The client needs to provide payment to access the request content.
	PaymentRequiredProblemDetail *shared.PaymentRequiredProblemDetail
	StatusCode                   int
	RawResponse                  *http.Response
	// This response is sent when the client has exhausted the request quota.
	TooManyRequestsProblemDetail *shared.TooManyRequestsProblemDetail
	// The client must authenticate itself to get the requested response.
	UnauthorizedProblemDetail *shared.UnauthorizedProblemDetail
	// The exported OpenAPI or JSON Schema file.
	ExportFileByCommit200ApplicationJSONAny interface{}
}

func (o *ExportFileByCommitResponse) GetBadRequestProblemDetail() *shared.BadRequestProblemDetail {
	if o == nil {
		return nil
	}
	return o.BadRequestProblemDetail
}

func (o *ExportFileByCommitResponse) GetBody() []byte {
	if o == nil {
		return nil
	}
	return o.Body
}

func (o *ExportFileByCommitResponse) GetConflictProblemDetail() *shared.ConflictProblemDetail {
	if o == nil {
		return nil
	}
	return o.ConflictProblemDetail
}

func (o *ExportFileByCommitResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ExportFileByCommitResponse) GetExportFileFailedProblemDetail() *shared.ExportFileFailedProblemDetail {
	if o == nil {
		return nil
	}
	return o.ExportFileFailedProblemDetail
}

func (o *ExportFileByCommitResponse) GetForbiddenProblemDetail() *shared.ForbiddenProblemDetail {
	if o == nil {
		return nil
	}
	return o.ForbiddenProblemDetail
}

func (o *ExportFileByCommitResponse) GetHeaders() map[string][]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *ExportFileByCommitResponse) GetInternalServerErrorProblemDetail() *shared.InternalServerErrorProblemDetail {
	if o == nil {
		return nil
	}
	return o.InternalServerErrorProblemDetail
}

func (o *ExportFileByCommitResponse) GetNotFoundProblemDetail() *shared.NotFoundProblemDetail {
	if o == nil {
		return nil
	}
	return o.NotFoundProblemDetail
}

func (o *ExportFileByCommitResponse) GetPaymentRequiredProblemDetail() *shared.PaymentRequiredProblemDetail {
	if o == nil {
		return nil
	}
	return o.PaymentRequiredProblemDetail
}

func (o *ExportFileByCommitResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ExportFileByCommitResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ExportFileByCommitResponse) GetTooManyRequestsProblemDetail() *shared.TooManyRequestsProblemDetail {
	if o == nil {
		return nil
	}
	return o.TooManyRequestsProblemDetail
}

func (o *ExportFileByCommitResponse) GetUnauthorizedProblemDetail() *shared.UnauthorizedProblemDetail {
	if o == nil {
		return nil
	}
	return o.UnauthorizedProblemDetail
}

func (o *ExportFileByCommitResponse) GetExportFileByCommit200ApplicationJSONAny() interface{} {
	if o == nil {
		return nil
	}
	return o.ExportFileByCommit200ApplicationJSONAny
}
