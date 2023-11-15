// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"fmt"
)

// SchemasProblemDetailStatus - The HTTP status code generated by the origin server for this occurrence of the problem.
type SchemasProblemDetailStatus int64

const (
	SchemasProblemDetailStatusFourHundred              SchemasProblemDetailStatus = 400
	SchemasProblemDetailStatusFourHundredAndOne        SchemasProblemDetailStatus = 401
	SchemasProblemDetailStatusFourHundredAndTwo        SchemasProblemDetailStatus = 402
	SchemasProblemDetailStatusFourHundredAndThree      SchemasProblemDetailStatus = 403
	SchemasProblemDetailStatusFourHundredAndFour       SchemasProblemDetailStatus = 404
	SchemasProblemDetailStatusFourHundredAndNine       SchemasProblemDetailStatus = 409
	SchemasProblemDetailStatusFourHundredAndTwentyNine SchemasProblemDetailStatus = 429
	SchemasProblemDetailStatusFiveHundred              SchemasProblemDetailStatus = 500
)

func (e SchemasProblemDetailStatus) ToPointer() *SchemasProblemDetailStatus {
	return &e
}

func (e *SchemasProblemDetailStatus) UnmarshalJSON(data []byte) error {
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
		*e = SchemasProblemDetailStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SchemasProblemDetailStatus: %v", v)
	}
}

// ConflictProblemDetail - Represents when a request conflicts with the current state of the server.
type ConflictProblemDetail struct {
	// A human-readable explanation specific to this occurrence of the problem.
	Detail *string `json:"detail,omitempty"`
	// A URI reference that identifies the specific occurrence of the problem.
	Instance *string                    `json:"instance,omitempty"`
	Status   SchemasProblemDetailStatus `json:"status"`
	// A short, human-readable summary of the problem type.
	Title string `json:"title"`
	// A URI reference that identifies the problem type.
	Type string `json:"type"`
}

var _ error = &ConflictProblemDetail{}

func (e *ConflictProblemDetail) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
