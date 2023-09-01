// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// ExportFileFailedProblemDetailProblemDetailStatus - The HTTP status code generated by the origin server for this occurrence of the problem.
type ExportFileFailedProblemDetailProblemDetailStatus int64

const (
	ExportFileFailedProblemDetailProblemDetailStatusFourHundred              ExportFileFailedProblemDetailProblemDetailStatus = 400
	ExportFileFailedProblemDetailProblemDetailStatusFourHundredAndOne        ExportFileFailedProblemDetailProblemDetailStatus = 401
	ExportFileFailedProblemDetailProblemDetailStatusFourHundredAndTwo        ExportFileFailedProblemDetailProblemDetailStatus = 402
	ExportFileFailedProblemDetailProblemDetailStatusFourHundredAndThree      ExportFileFailedProblemDetailProblemDetailStatus = 403
	ExportFileFailedProblemDetailProblemDetailStatusFourHundredAndFour       ExportFileFailedProblemDetailProblemDetailStatus = 404
	ExportFileFailedProblemDetailProblemDetailStatusFourHundredAndNine       ExportFileFailedProblemDetailProblemDetailStatus = 409
	ExportFileFailedProblemDetailProblemDetailStatusFourHundredAndTwentyNine ExportFileFailedProblemDetailProblemDetailStatus = 429
	ExportFileFailedProblemDetailProblemDetailStatusFiveHundred              ExportFileFailedProblemDetailProblemDetailStatus = 500
)

func (e ExportFileFailedProblemDetailProblemDetailStatus) ToPointer() *ExportFileFailedProblemDetailProblemDetailStatus {
	return &e
}

func (e *ExportFileFailedProblemDetailProblemDetailStatus) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 400:
		fallthrough
	case 401:
		fallthrough
	case 402:
		fallthrough
	case 403:
		fallthrough
	case 404:
		fallthrough
	case 409:
		fallthrough
	case 429:
		fallthrough
	case 500:
		*e = ExportFileFailedProblemDetailProblemDetailStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ExportFileFailedProblemDetailProblemDetailStatus: %v", v)
	}
}

// ExportFileFailedProblemDetail - Represents when the export request is well-formed, but the server was unable to successfully export the requested file.
type ExportFileFailedProblemDetail struct {
	// A human-readable explanation specific to this occurrence of the problem.
	Detail         *string         `json:"detail,omitempty"`
	ExportProblems []ExportProblem `json:"export_problems"`
	// A URI reference that identifies the specific occurrence of the problem.
	Instance *string                                          `json:"instance,omitempty"`
	Status   ExportFileFailedProblemDetailProblemDetailStatus `json:"status"`
	// A short, human-readable summary of the problem type.
	Title string `json:"title"`
	// A URI reference that identifies the problem type.
	Type string `json:"type"`
}

func (o *ExportFileFailedProblemDetail) GetDetail() *string {
	if o == nil {
		return nil
	}
	return o.Detail
}

func (o *ExportFileFailedProblemDetail) GetExportProblems() []ExportProblem {
	if o == nil {
		return []ExportProblem{}
	}
	return o.ExportProblems
}

func (o *ExportFileFailedProblemDetail) GetInstance() *string {
	if o == nil {
		return nil
	}
	return o.Instance
}

func (o *ExportFileFailedProblemDetail) GetStatus() ExportFileFailedProblemDetailProblemDetailStatus {
	if o == nil {
		return ExportFileFailedProblemDetailProblemDetailStatus(0)
	}
	return o.Status
}

func (o *ExportFileFailedProblemDetail) GetTitle() string {
	if o == nil {
		return ""
	}
	return o.Title
}

func (o *ExportFileFailedProblemDetail) GetType() string {
	if o == nil {
		return ""
	}
	return o.Type
}