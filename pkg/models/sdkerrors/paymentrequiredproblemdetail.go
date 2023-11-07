// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"fmt"
)

// SchemasPaymentRequiredProblemDetailProblemDetailStatus - The HTTP status code generated by the origin server for this occurrence of the problem.
type SchemasPaymentRequiredProblemDetailProblemDetailStatus int64

const (
	SchemasPaymentRequiredProblemDetailProblemDetailStatusFourHundred              SchemasPaymentRequiredProblemDetailProblemDetailStatus = 400
	SchemasPaymentRequiredProblemDetailProblemDetailStatusFourHundredAndOne        SchemasPaymentRequiredProblemDetailProblemDetailStatus = 401
	SchemasPaymentRequiredProblemDetailProblemDetailStatusFourHundredAndTwo        SchemasPaymentRequiredProblemDetailProblemDetailStatus = 402
	SchemasPaymentRequiredProblemDetailProblemDetailStatusFourHundredAndThree      SchemasPaymentRequiredProblemDetailProblemDetailStatus = 403
	SchemasPaymentRequiredProblemDetailProblemDetailStatusFourHundredAndFour       SchemasPaymentRequiredProblemDetailProblemDetailStatus = 404
	SchemasPaymentRequiredProblemDetailProblemDetailStatusFourHundredAndNine       SchemasPaymentRequiredProblemDetailProblemDetailStatus = 409
	SchemasPaymentRequiredProblemDetailProblemDetailStatusFourHundredAndTwentyNine SchemasPaymentRequiredProblemDetailProblemDetailStatus = 429
	SchemasPaymentRequiredProblemDetailProblemDetailStatusFiveHundred              SchemasPaymentRequiredProblemDetailProblemDetailStatus = 500
)

func (e SchemasPaymentRequiredProblemDetailProblemDetailStatus) ToPointer() *SchemasPaymentRequiredProblemDetailProblemDetailStatus {
	return &e
}

func (e *SchemasPaymentRequiredProblemDetailProblemDetailStatus) UnmarshalJSON(data []byte) error {
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
		*e = SchemasPaymentRequiredProblemDetailProblemDetailStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SchemasPaymentRequiredProblemDetailProblemDetailStatus: %v", v)
	}
}

// PaymentRequiredProblemDetail - Represents when the server is refusing access to the resource until payment has been made.
type PaymentRequiredProblemDetail struct {
	// A human-readable explanation specific to this occurrence of the problem.
	Detail *string `json:"detail,omitempty"`
	// A URI reference that identifies the specific occurrence of the problem.
	Instance *string                                                `json:"instance,omitempty"`
	Status   SchemasPaymentRequiredProblemDetailProblemDetailStatus `json:"status"`
	// A short, human-readable summary of the problem type.
	Title string `json:"title"`
	// A URI reference that identifies the problem type.
	Type string `json:"type"`
}

var _ error = &PaymentRequiredProblemDetail{}

func (e *PaymentRequiredProblemDetail) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}