// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// TooManyRequestsProblemDetailProblemDetailStatus - The HTTP status code generated by the origin server for this occurrence of the problem.
type TooManyRequestsProblemDetailProblemDetailStatus int64

const (
	TooManyRequestsProblemDetailProblemDetailStatusFourHundred              TooManyRequestsProblemDetailProblemDetailStatus = 400
	TooManyRequestsProblemDetailProblemDetailStatusFourHundredAndOne        TooManyRequestsProblemDetailProblemDetailStatus = 401
	TooManyRequestsProblemDetailProblemDetailStatusFourHundredAndTwo        TooManyRequestsProblemDetailProblemDetailStatus = 402
	TooManyRequestsProblemDetailProblemDetailStatusFourHundredAndThree      TooManyRequestsProblemDetailProblemDetailStatus = 403
	TooManyRequestsProblemDetailProblemDetailStatusFourHundredAndFour       TooManyRequestsProblemDetailProblemDetailStatus = 404
	TooManyRequestsProblemDetailProblemDetailStatusFourHundredAndNine       TooManyRequestsProblemDetailProblemDetailStatus = 409
	TooManyRequestsProblemDetailProblemDetailStatusFourHundredAndTwentyNine TooManyRequestsProblemDetailProblemDetailStatus = 429
	TooManyRequestsProblemDetailProblemDetailStatusFiveHundred              TooManyRequestsProblemDetailProblemDetailStatus = 500
)

func (e TooManyRequestsProblemDetailProblemDetailStatus) ToPointer() *TooManyRequestsProblemDetailProblemDetailStatus {
	return &e
}

func (e *TooManyRequestsProblemDetailProblemDetailStatus) UnmarshalJSON(data []byte) error {
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
		*e = TooManyRequestsProblemDetailProblemDetailStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TooManyRequestsProblemDetailProblemDetailStatus: %v", v)
	}
}

// TooManyRequestsProblemDetail - Represents when the client has exhausted its rate limiting quota.
type TooManyRequestsProblemDetail struct {
	// A human-readable explanation specific to this occurrence of the problem.
	Detail *string `json:"detail,omitempty"`
	// A URI reference that identifies the specific occurrence of the problem.
	Instance *string `json:"instance,omitempty"`
	// The number of seconds to wait to retry the request again.
	RetryAfter *int                                            `json:"retry_after,omitempty"`
	Status     TooManyRequestsProblemDetailProblemDetailStatus `json:"status"`
	// A short, human-readable summary of the problem type.
	Title string `json:"title"`
	// A URI reference that identifies the problem type.
	Type string `json:"type"`
}

func (o *TooManyRequestsProblemDetail) GetDetail() *string {
	if o == nil {
		return nil
	}
	return o.Detail
}

func (o *TooManyRequestsProblemDetail) GetInstance() *string {
	if o == nil {
		return nil
	}
	return o.Instance
}

func (o *TooManyRequestsProblemDetail) GetRetryAfter() *int {
	if o == nil {
		return nil
	}
	return o.RetryAfter
}

func (o *TooManyRequestsProblemDetail) GetStatus() TooManyRequestsProblemDetailProblemDetailStatus {
	if o == nil {
		return TooManyRequestsProblemDetailProblemDetailStatus(0)
	}
	return o.Status
}

func (o *TooManyRequestsProblemDetail) GetTitle() string {
	if o == nil {
		return ""
	}
	return o.Title
}

func (o *TooManyRequestsProblemDetail) GetType() string {
	if o == nil {
		return ""
	}
	return o.Type
}
