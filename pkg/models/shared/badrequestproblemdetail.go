// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// BadRequestProblemDetailProblemDetailStatus - The HTTP status code generated by the origin server for this occurrence of the problem.
type BadRequestProblemDetailProblemDetailStatus int64

const (
	BadRequestProblemDetailProblemDetailStatusFourHundred              BadRequestProblemDetailProblemDetailStatus = 400
	BadRequestProblemDetailProblemDetailStatusFourHundredAndOne        BadRequestProblemDetailProblemDetailStatus = 401
	BadRequestProblemDetailProblemDetailStatusFourHundredAndTwo        BadRequestProblemDetailProblemDetailStatus = 402
	BadRequestProblemDetailProblemDetailStatusFourHundredAndThree      BadRequestProblemDetailProblemDetailStatus = 403
	BadRequestProblemDetailProblemDetailStatusFourHundredAndFour       BadRequestProblemDetailProblemDetailStatus = 404
	BadRequestProblemDetailProblemDetailStatusFourHundredAndNine       BadRequestProblemDetailProblemDetailStatus = 409
	BadRequestProblemDetailProblemDetailStatusFourHundredAndTwentyNine BadRequestProblemDetailProblemDetailStatus = 429
	BadRequestProblemDetailProblemDetailStatusFiveHundred              BadRequestProblemDetailProblemDetailStatus = 500
)

func (e BadRequestProblemDetailProblemDetailStatus) ToPointer() *BadRequestProblemDetailProblemDetailStatus {
	return &e
}

func (e *BadRequestProblemDetailProblemDetailStatus) UnmarshalJSON(data []byte) error {
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
		*e = BadRequestProblemDetailProblemDetailStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for BadRequestProblemDetailProblemDetailStatus: %v", v)
	}
}

// BadRequestProblemDetail - Represents when the server could not understand the request due to invalid syntax.
type BadRequestProblemDetail struct {
	// A human-readable explanation specific to this occurrence of the problem.
	Detail *string `json:"detail,omitempty"`
	// A URI reference that identifies the specific occurrence of the problem.
	Instance          *string                                    `json:"instance,omitempty"`
	InvalidParameters []InvalidParameter                         `json:"invalid_parameters"`
	Status            BadRequestProblemDetailProblemDetailStatus `json:"status"`
	// A short, human-readable summary of the problem type.
	Title string `json:"title"`
	// A URI reference that identifies the problem type.
	Type string `json:"type"`
}

func (o *BadRequestProblemDetail) GetDetail() *string {
	if o == nil {
		return nil
	}
	return o.Detail
}

func (o *BadRequestProblemDetail) GetInstance() *string {
	if o == nil {
		return nil
	}
	return o.Instance
}

func (o *BadRequestProblemDetail) GetInvalidParameters() []InvalidParameter {
	if o == nil {
		return []InvalidParameter{}
	}
	return o.InvalidParameters
}

func (o *BadRequestProblemDetail) GetStatus() BadRequestProblemDetailProblemDetailStatus {
	if o == nil {
		return BadRequestProblemDetailProblemDetailStatus(0)
	}
	return o.Status
}

func (o *BadRequestProblemDetail) GetTitle() string {
	if o == nil {
		return ""
	}
	return o.Title
}

func (o *BadRequestProblemDetail) GetType() string {
	if o == nil {
		return ""
	}
	return o.Type
}
