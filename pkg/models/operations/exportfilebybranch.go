// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/naumanali/v2/pkg/utils"
	"net/http"
)

// StoplightAPIVersionString - A string representing the Stoplight API version that is being requested. If not supplied: TODO document policy
type StoplightAPIVersionString string

const (
	StoplightAPIVersionStringTwoThousandAndTwentyTwo1205 StoplightAPIVersionString = "2022-12-05"
)

func (e StoplightAPIVersionString) ToPointer() *StoplightAPIVersionString {
	return &e
}

func (e *StoplightAPIVersionString) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "2022-12-05":
		*e = StoplightAPIVersionString(v)
		return nil
	default:
		return fmt.Errorf("invalid value for StoplightAPIVersionString: %v", v)
	}
}

type ExportFileByBranchRequest struct {
	// A string representing the Stoplight API version that is being requested. If not supplied: TODO document policy
	StoplightVersion *StoplightAPIVersionString `header:"style=simple,explode=false,name=Stoplight-Version"`
	// A reference to a branch tracked by a Stoplight project. Must be percent encoded.
	BranchName string `pathParam:"style=simple,explode=false,name=branch_name"`
	// A path to a file in a Stoplight project. Use forward slashes (`/`) to separate path directories.
	FilePath string `pathParam:"style=simple,explode=false,name=file_path"`
	// When `false`, indicates that any declaration (schema, operation, property, etc.) in the OpenAPI or JSON Schema marked with the `x-internal` extension property set to `true` will be omitted from the exported file. Setting this parameter to `true` with an anonymous or workspace guest user caller results in an error response.
	IncludeInternal *bool `default:"true" queryParam:"style=form,explode=true,name=include_internal"`
	// A string that uniquely identifies a Stoplight Project resource. This can be found in the [Project Settings](https://docs.stoplight.io/docs/platform/252039ebe8fb2-project-settings) page.
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

func (e ExportFileByBranchRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(e, "", false)
}

func (e *ExportFileByBranchRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &e, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ExportFileByBranchRequest) GetStoplightVersion() *StoplightAPIVersionString {
	if o == nil {
		return nil
	}
	return o.StoplightVersion
}

func (o *ExportFileByBranchRequest) GetBranchName() string {
	if o == nil {
		return ""
	}
	return o.BranchName
}

func (o *ExportFileByBranchRequest) GetFilePath() string {
	if o == nil {
		return ""
	}
	return o.FilePath
}

func (o *ExportFileByBranchRequest) GetIncludeInternal() *bool {
	if o == nil {
		return nil
	}
	return o.IncludeInternal
}

func (o *ExportFileByBranchRequest) GetProjectID() string {
	if o == nil {
		return ""
	}
	return o.ProjectID
}

type ExportFileByBranchResponse struct {
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

func (o *ExportFileByBranchResponse) GetTwoHundredApplicationJSONAny() interface{} {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationJSONAny
}

func (o *ExportFileByBranchResponse) GetBody() []byte {
	if o == nil {
		return nil
	}
	return o.Body
}

func (o *ExportFileByBranchResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ExportFileByBranchResponse) GetHeaders() map[string][]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *ExportFileByBranchResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ExportFileByBranchResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
