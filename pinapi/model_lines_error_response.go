/*
 * Pinnacle - Customer API Reference
 *
 *  # Authentication   API uses HTTP Basic access authentication. You need to send Authorization HTTP Request header:    `Authorization: Basic <Base64 value of UTF-8 encoded \"username:password\">`  Example:  `Authorization: Basic U03MyOT23YbzMDc6d3c3O1DQ1` 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package pinapi

import (
	"encoding/json"
)

// LinesErrorResponse struct for LinesErrorResponse
type LinesErrorResponse struct {
	Status *string `json:"status,omitempty"`
	Error *ErrorResponse `json:"error,omitempty"`
	// Code identifying an error that occurred.
	Code *int `json:"code,omitempty"`
}

// NewLinesErrorResponse instantiates a new LinesErrorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinesErrorResponse() *LinesErrorResponse {
	this := LinesErrorResponse{}
	return &this
}

// NewLinesErrorResponseWithDefaults instantiates a new LinesErrorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinesErrorResponseWithDefaults() *LinesErrorResponse {
	this := LinesErrorResponse{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *LinesErrorResponse) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinesErrorResponse) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *LinesErrorResponse) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *LinesErrorResponse) SetStatus(v string) {
	o.Status = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *LinesErrorResponse) GetError() ErrorResponse {
	if o == nil || o.Error == nil {
		var ret ErrorResponse
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinesErrorResponse) GetErrorOk() (*ErrorResponse, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *LinesErrorResponse) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given ErrorResponse and assigns it to the Error field.
func (o *LinesErrorResponse) SetError(v ErrorResponse) {
	o.Error = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *LinesErrorResponse) GetCode() int {
	if o == nil || o.Code == nil {
		var ret int
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinesErrorResponse) GetCodeOk() (*int, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *LinesErrorResponse) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given int and assigns it to the Code field.
func (o *LinesErrorResponse) SetCode(v int) {
	o.Code = &v
}

func (o LinesErrorResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	return json.Marshal(toSerialize)
}

type NullableLinesErrorResponse struct {
	value *LinesErrorResponse
	isSet bool
}

func (v NullableLinesErrorResponse) Get() *LinesErrorResponse {
	return v.value
}

func (v *NullableLinesErrorResponse) Set(val *LinesErrorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLinesErrorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLinesErrorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinesErrorResponse(val *LinesErrorResponse) *NullableLinesErrorResponse {
	return &NullableLinesErrorResponse{value: val, isSet: true}
}

func (v NullableLinesErrorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinesErrorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
