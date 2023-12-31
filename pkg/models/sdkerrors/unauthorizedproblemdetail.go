// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"fmt"
)

// SchemasUnauthorizedProblemDetailProblemDetailStatus - The HTTP status code generated by the origin server for this occurrence of the problem.
type SchemasUnauthorizedProblemDetailProblemDetailStatus int64

const (
	SchemasUnauthorizedProblemDetailProblemDetailStatusFourHundred              SchemasUnauthorizedProblemDetailProblemDetailStatus = 400
	SchemasUnauthorizedProblemDetailProblemDetailStatusFourHundredAndOne        SchemasUnauthorizedProblemDetailProblemDetailStatus = 401
	SchemasUnauthorizedProblemDetailProblemDetailStatusFourHundredAndTwo        SchemasUnauthorizedProblemDetailProblemDetailStatus = 402
	SchemasUnauthorizedProblemDetailProblemDetailStatusFourHundredAndThree      SchemasUnauthorizedProblemDetailProblemDetailStatus = 403
	SchemasUnauthorizedProblemDetailProblemDetailStatusFourHundredAndFour       SchemasUnauthorizedProblemDetailProblemDetailStatus = 404
	SchemasUnauthorizedProblemDetailProblemDetailStatusFourHundredAndNine       SchemasUnauthorizedProblemDetailProblemDetailStatus = 409
	SchemasUnauthorizedProblemDetailProblemDetailStatusFourHundredAndTwentyNine SchemasUnauthorizedProblemDetailProblemDetailStatus = 429
	SchemasUnauthorizedProblemDetailProblemDetailStatusFiveHundred              SchemasUnauthorizedProblemDetailProblemDetailStatus = 500
)

func (e SchemasUnauthorizedProblemDetailProblemDetailStatus) ToPointer() *SchemasUnauthorizedProblemDetailProblemDetailStatus {
	return &e
}

func (e *SchemasUnauthorizedProblemDetailProblemDetailStatus) UnmarshalJSON(data []byte) error {
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
		*e = SchemasUnauthorizedProblemDetailProblemDetailStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SchemasUnauthorizedProblemDetailProblemDetailStatus: %v", v)
	}
}

// UnauthorizedProblemDetail - Represents when the client must authenticate itself to get the requested response.
type UnauthorizedProblemDetail struct {
	// A human-readable explanation specific to this occurrence of the problem.
	Detail *string `json:"detail,omitempty"`
	// A URI reference that identifies the specific occurrence of the problem.
	Instance *string                                             `json:"instance,omitempty"`
	Status   SchemasUnauthorizedProblemDetailProblemDetailStatus `json:"status"`
	// A short, human-readable summary of the problem type.
	Title string `json:"title"`
	// A URI reference that identifies the problem type.
	Type string `json:"type"`
}

var _ error = &UnauthorizedProblemDetail{}

func (e *UnauthorizedProblemDetail) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
