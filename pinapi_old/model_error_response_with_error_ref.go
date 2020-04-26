/*
 * Pinnacle - Customer API Reference
 *
 *  # Authentication   API uses HTTP Basic access authentication. You need to send Authorization HTTP Request header:    `Authorization: Basic <Base64 value of UTF-8 encoded \"username:password\">`  Example:  `Authorization: Basic U03MyOT23YbzMDc6d3c3O1DQ1`
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pinapi_old

import (
	"bytes"
	"encoding/json"
)

// ErrorResponseWithErrorRef struct for ErrorResponseWithErrorRef
type ErrorResponseWithErrorRef struct {
	Ref     *string `json:"ref,omitempty" xml:"ref"`
	Code    *string `json:"code,omitempty" xml:"code"`
	Message *string `json:"message,omitempty" xml:"message"`
}

// GetRef returns the Ref field value if set, zero value otherwise.
func (o *ErrorResponseWithErrorRef) GetRef() string {
	if o == nil || o.Ref == nil {
		var ret string
		return ret
	}
	return *o.Ref
}

// GetRefOk returns a tuple with the Ref field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ErrorResponseWithErrorRef) GetRefOk() (string, bool) {
	if o == nil || o.Ref == nil {
		var ret string
		return ret, false
	}
	return *o.Ref, true
}

// HasRef returns a boolean if a field has been set.
func (o *ErrorResponseWithErrorRef) HasRef() bool {
	if o != nil && o.Ref != nil {
		return true
	}

	return false
}

// SetRef gets a reference to the given string and assigns it to the Ref field.
func (o *ErrorResponseWithErrorRef) SetRef(v string) {
	o.Ref = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ErrorResponseWithErrorRef) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ErrorResponseWithErrorRef) GetCodeOk() (string, bool) {
	if o == nil || o.Code == nil {
		var ret string
		return ret, false
	}
	return *o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ErrorResponseWithErrorRef) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ErrorResponseWithErrorRef) SetCode(v string) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ErrorResponseWithErrorRef) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ErrorResponseWithErrorRef) GetMessageOk() (string, bool) {
	if o == nil || o.Message == nil {
		var ret string
		return ret, false
	}
	return *o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ErrorResponseWithErrorRef) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ErrorResponseWithErrorRef) SetMessage(v string) {
	o.Message = &v
}

type NullableErrorResponseWithErrorRef struct {
	Value        ErrorResponseWithErrorRef
	ExplicitNull bool
}

func (v NullableErrorResponseWithErrorRef) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableErrorResponseWithErrorRef) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
