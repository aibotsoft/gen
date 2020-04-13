/*
 * Pinnacle - Customer API Reference
 *
 *  # Authentication   API uses HTTP Basic access authentication. You need to send Authorization HTTP Request header:    `Authorization: Basic <Base64 value of UTF-8 encoded \"username:password\">`  Example:  `Authorization: Basic U03MyOT23YbzMDc6d3c3O1DQ1`
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pinapi

import (
	"bytes"
	"encoding/json"
)

// CancellationReasonType struct for CancellationReasonType
type CancellationReasonType struct {
	// Cancellation Reason Code
	Code    *string                        `json:"code,omitempty" xml:"code"`
	Details *CancellationReasonDetailsType `json:"details,omitempty" xml:"details"`
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *CancellationReasonType) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CancellationReasonType) GetCodeOk() (string, bool) {
	if o == nil || o.Code == nil {
		var ret string
		return ret, false
	}
	return *o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *CancellationReasonType) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *CancellationReasonType) SetCode(v string) {
	o.Code = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *CancellationReasonType) GetDetails() CancellationReasonDetailsType {
	if o == nil || o.Details == nil {
		var ret CancellationReasonDetailsType
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CancellationReasonType) GetDetailsOk() (CancellationReasonDetailsType, bool) {
	if o == nil || o.Details == nil {
		var ret CancellationReasonDetailsType
		return ret, false
	}
	return *o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *CancellationReasonType) HasDetails() bool {
	if o != nil && o.Details != nil {
		return true
	}

	return false
}

// SetDetails gets a reference to the given CancellationReasonDetailsType and assigns it to the Details field.
func (o *CancellationReasonType) SetDetails(v CancellationReasonDetailsType) {
	o.Details = &v
}

type NullableCancellationReasonType struct {
	Value        CancellationReasonType
	ExplicitNull bool
}

func (v NullableCancellationReasonType) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableCancellationReasonType) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}