/*
 * Sportmarket API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package sportmarketapi

import (
	"encoding/json"
)

// ErrorData struct for ErrorData
type ErrorData struct {
	ValidationErrors *map[string][]string `json:"validation_errors,omitempty"`
}

// NewErrorData instantiates a new ErrorData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorData() *ErrorData {
	this := ErrorData{}
	return &this
}

// NewErrorDataWithDefaults instantiates a new ErrorData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorDataWithDefaults() *ErrorData {
	this := ErrorData{}
	return &this
}

// GetValidationErrors returns the ValidationErrors field value if set, zero value otherwise.
func (o *ErrorData) GetValidationErrors() map[string][]string {
	if o == nil || o.ValidationErrors == nil {
		var ret map[string][]string
		return ret
	}
	return *o.ValidationErrors
}

// GetValidationErrorsOk returns a tuple with the ValidationErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorData) GetValidationErrorsOk() (*map[string][]string, bool) {
	if o == nil || o.ValidationErrors == nil {
		return nil, false
	}
	return o.ValidationErrors, true
}

// HasValidationErrors returns a boolean if a field has been set.
func (o *ErrorData) HasValidationErrors() bool {
	if o != nil && o.ValidationErrors != nil {
		return true
	}

	return false
}

// SetValidationErrors gets a reference to the given map[string][]string and assigns it to the ValidationErrors field.
func (o *ErrorData) SetValidationErrors(v map[string][]string) {
	o.ValidationErrors = &v
}

func (o ErrorData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ValidationErrors != nil {
		toSerialize["validation_errors"] = o.ValidationErrors
	}
	return json.Marshal(toSerialize)
}

type NullableErrorData struct {
	value *ErrorData
	isSet bool
}

func (v NullableErrorData) Get() *ErrorData {
	return v.value
}

func (v *NullableErrorData) Set(val *ErrorData) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorData) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorData(val *ErrorData) *NullableErrorData {
	return &NullableErrorData{value: val, isSet: true}
}

func (v NullableErrorData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
