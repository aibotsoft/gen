/*
 * EpinApi
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package epinapi

import (
	"encoding/json"
)

// UnauthorizedError Unauthorized Response
type UnauthorizedError struct {
	// detail
	Detail string `json:"detail"`
	// status
	Status int64 `json:"status"`
	// title
	Title string `json:"title"`
	// type
	Type string `json:"type"`
}

// NewUnauthorizedError instantiates a new UnauthorizedError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnauthorizedError(detail string, status int64, title string, type_ string, ) *UnauthorizedError {
	this := UnauthorizedError{}
	this.Detail = detail
	this.Status = status
	this.Title = title
	this.Type = type_
	return &this
}

// NewUnauthorizedErrorWithDefaults instantiates a new UnauthorizedError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnauthorizedErrorWithDefaults() *UnauthorizedError {
	this := UnauthorizedError{}
	return &this
}

// GetDetail returns the Detail field value
func (o *UnauthorizedError) GetDetail() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Detail
}

// GetDetailOk returns a tuple with the Detail field value
// and a boolean to check if the value has been set.
func (o *UnauthorizedError) GetDetailOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Detail, true
}

// SetDetail sets field value
func (o *UnauthorizedError) SetDetail(v string) {
	o.Detail = v
}

// GetStatus returns the Status field value
func (o *UnauthorizedError) GetStatus() int64 {
	if o == nil  {
		var ret int64
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *UnauthorizedError) GetStatusOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *UnauthorizedError) SetStatus(v int64) {
	o.Status = v
}

// GetTitle returns the Title field value
func (o *UnauthorizedError) GetTitle() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *UnauthorizedError) GetTitleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *UnauthorizedError) SetTitle(v string) {
	o.Title = v
}

// GetType returns the Type field value
func (o *UnauthorizedError) GetType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *UnauthorizedError) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *UnauthorizedError) SetType(v string) {
	o.Type = v
}

func (o UnauthorizedError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["detail"] = o.Detail
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["title"] = o.Title
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableUnauthorizedError struct {
	value *UnauthorizedError
	isSet bool
}

func (v NullableUnauthorizedError) Get() *UnauthorizedError {
	return v.value
}

func (v *NullableUnauthorizedError) Set(val *UnauthorizedError) {
	v.value = val
	v.isSet = true
}

func (v NullableUnauthorizedError) IsSet() bool {
	return v.isSet
}

func (v *NullableUnauthorizedError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnauthorizedError(val *UnauthorizedError) *NullableUnauthorizedError {
	return &NullableUnauthorizedError{value: val, isSet: true}
}

func (v NullableUnauthorizedError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnauthorizedError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
