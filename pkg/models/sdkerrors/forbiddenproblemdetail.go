// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"fmt"
)

// SchemasForbiddenProblemDetailProblemDetailStatus - The HTTP status code generated by the origin server for this occurrence of the problem.
type SchemasForbiddenProblemDetailProblemDetailStatus int64

const (
	SchemasForbiddenProblemDetailProblemDetailStatusFourHundred              SchemasForbiddenProblemDetailProblemDetailStatus = 400
	SchemasForbiddenProblemDetailProblemDetailStatusFourHundredAndOne        SchemasForbiddenProblemDetailProblemDetailStatus = 401
	SchemasForbiddenProblemDetailProblemDetailStatusFourHundredAndTwo        SchemasForbiddenProblemDetailProblemDetailStatus = 402
	SchemasForbiddenProblemDetailProblemDetailStatusFourHundredAndThree      SchemasForbiddenProblemDetailProblemDetailStatus = 403
	SchemasForbiddenProblemDetailProblemDetailStatusFourHundredAndFour       SchemasForbiddenProblemDetailProblemDetailStatus = 404
	SchemasForbiddenProblemDetailProblemDetailStatusFourHundredAndNine       SchemasForbiddenProblemDetailProblemDetailStatus = 409
	SchemasForbiddenProblemDetailProblemDetailStatusFourHundredAndTwentyNine SchemasForbiddenProblemDetailProblemDetailStatus = 429
	SchemasForbiddenProblemDetailProblemDetailStatusFiveHundred              SchemasForbiddenProblemDetailProblemDetailStatus = 500
)

func (e SchemasForbiddenProblemDetailProblemDetailStatus) ToPointer() *SchemasForbiddenProblemDetailProblemDetailStatus {
	return &e
}

func (e *SchemasForbiddenProblemDetailProblemDetailStatus) UnmarshalJSON(data []byte) error {
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
		*e = SchemasForbiddenProblemDetailProblemDetailStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SchemasForbiddenProblemDetailProblemDetailStatus: %v", v)
	}
}

// ForbiddenProblemDetail - Represents when the client does not have permissions to access the resource.
type ForbiddenProblemDetail struct {
	// A human-readable explanation specific to this occurrence of the problem.
	Detail *string `json:"detail,omitempty"`
	// A URI reference that identifies the specific occurrence of the problem.
	Instance *string                                          `json:"instance,omitempty"`
	Status   SchemasForbiddenProblemDetailProblemDetailStatus `json:"status"`
	// A short, human-readable summary of the problem type.
	Title string `json:"title"`
	// A URI reference that identifies the problem type.
	Type string `json:"type"`
}

var _ error = &ForbiddenProblemDetail{}

func (e *ForbiddenProblemDetail) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}