// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/naumanali/v2/pkg/utils"
	"net/http"
)

// HeaderStoplightAPIVersionString - A string representing the Stoplight API version that is being requested. If not supplied: TODO document policy
type HeaderStoplightAPIVersionString string

const (
	HeaderStoplightAPIVersionStringTwoThousandAndTwentyTwo1205 HeaderStoplightAPIVersionString = "2022-12-05"
)

func (e HeaderStoplightAPIVersionString) ToPointer() *HeaderStoplightAPIVersionString {
	return &e
}

func (e *HeaderStoplightAPIVersionString) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "2022-12-05":
		*e = HeaderStoplightAPIVersionString(v)
		return nil
	default:
		return fmt.Errorf("invalid value for HeaderStoplightAPIVersionString: %v", v)
	}
}

type ExportFileByCommitRequest struct {
	// A string representing the Stoplight API version that is being requested. If not supplied: TODO document policy
	StoplightVersion *HeaderStoplightAPIVersionString `header:"style=simple,explode=false,name=Stoplight-Version"`
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

func (o *ExportFileByCommitRequest) GetStoplightVersion() *HeaderStoplightAPIVersionString {
	if o == nil {
		return nil
	}
	return o.StoplightVersion
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
	// The exported OpenAPI or JSON Schema file.
	TwoHundredApplicationJSONAny interface{}
	Body                         []byte
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ExportFileByCommitResponse) GetTwoHundredApplicationJSONAny() interface{} {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationJSONAny
}

func (o *ExportFileByCommitResponse) GetBody() []byte {
	if o == nil {
		return nil
	}
	return o.Body
}

func (o *ExportFileByCommitResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ExportFileByCommitResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
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
