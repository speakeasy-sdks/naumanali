// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"fmt"
)

// SchemasNotFoundProblemDetailProblemDetailStatus - The HTTP status code generated by the origin server for this occurrence of the problem.
type SchemasNotFoundProblemDetailProblemDetailStatus int64

const (
	SchemasNotFoundProblemDetailProblemDetailStatusFourHundred              SchemasNotFoundProblemDetailProblemDetailStatus = 400
	SchemasNotFoundProblemDetailProblemDetailStatusFourHundredAndOne        SchemasNotFoundProblemDetailProblemDetailStatus = 401
	SchemasNotFoundProblemDetailProblemDetailStatusFourHundredAndTwo        SchemasNotFoundProblemDetailProblemDetailStatus = 402
	SchemasNotFoundProblemDetailProblemDetailStatusFourHundredAndThree      SchemasNotFoundProblemDetailProblemDetailStatus = 403
	SchemasNotFoundProblemDetailProblemDetailStatusFourHundredAndFour       SchemasNotFoundProblemDetailProblemDetailStatus = 404
	SchemasNotFoundProblemDetailProblemDetailStatusFourHundredAndNine       SchemasNotFoundProblemDetailProblemDetailStatus = 409
	SchemasNotFoundProblemDetailProblemDetailStatusFourHundredAndTwentyNine SchemasNotFoundProblemDetailProblemDetailStatus = 429
	SchemasNotFoundProblemDetailProblemDetailStatusFiveHundred              SchemasNotFoundProblemDetailProblemDetailStatus = 500
)

func (e SchemasNotFoundProblemDetailProblemDetailStatus) ToPointer() *SchemasNotFoundProblemDetailProblemDetailStatus {
	return &e
}

func (e *SchemasNotFoundProblemDetailProblemDetailStatus) UnmarshalJSON(data []byte) error {
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
		*e = SchemasNotFoundProblemDetailProblemDetailStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SchemasNotFoundProblemDetailProblemDetailStatus: %v", v)
	}
}

// NotFoundProblemDetail - Represents when the server can not find the requested resource.
type NotFoundProblemDetail struct {
	// A human-readable explanation specific to this occurrence of the problem.
	Detail *string `json:"detail,omitempty"`
	// A URI reference that identifies the specific occurrence of the problem.
	Instance *string                                         `json:"instance,omitempty"`
	Status   SchemasNotFoundProblemDetailProblemDetailStatus `json:"status"`
	// A short, human-readable summary of the problem type.
	Title string `json:"title"`
	// A URI reference that identifies the problem type.
	Type string `json:"type"`
}

var _ error = &NotFoundProblemDetail{}

func (e *NotFoundProblemDetail) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
