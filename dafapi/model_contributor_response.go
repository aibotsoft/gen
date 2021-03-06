/*
 * Sample API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package dafapi

import (
	"encoding/json"
)

// ContributorResponse struct for ContributorResponse
type ContributorResponse struct {
	ErrorCode int64 `json:"ErrorCode"`
	ErrorMsg string `json:"ErrorMsg"`
	Data []SportItem `json:"Data"`
}

// NewContributorResponse instantiates a new ContributorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContributorResponse(errorCode int64, errorMsg string, data []SportItem, ) *ContributorResponse {
	this := ContributorResponse{}
	this.ErrorCode = errorCode
	this.ErrorMsg = errorMsg
	this.Data = data
	return &this
}

// NewContributorResponseWithDefaults instantiates a new ContributorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContributorResponseWithDefaults() *ContributorResponse {
	this := ContributorResponse{}
	return &this
}

// GetErrorCode returns the ErrorCode field value
func (o *ContributorResponse) GetErrorCode() int64 {
	if o == nil  {
		var ret int64
		return ret
	}

	return o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value
// and a boolean to check if the value has been set.
func (o *ContributorResponse) GetErrorCodeOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ErrorCode, true
}

// SetErrorCode sets field value
func (o *ContributorResponse) SetErrorCode(v int64) {
	o.ErrorCode = v
}

// GetErrorMsg returns the ErrorMsg field value
func (o *ContributorResponse) GetErrorMsg() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ErrorMsg
}

// GetErrorMsgOk returns a tuple with the ErrorMsg field value
// and a boolean to check if the value has been set.
func (o *ContributorResponse) GetErrorMsgOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ErrorMsg, true
}

// SetErrorMsg sets field value
func (o *ContributorResponse) SetErrorMsg(v string) {
	o.ErrorMsg = v
}

// GetData returns the Data field value
func (o *ContributorResponse) GetData() []SportItem {
	if o == nil  {
		var ret []SportItem
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ContributorResponse) GetDataOk() (*[]SportItem, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ContributorResponse) SetData(v []SportItem) {
	o.Data = v
}

func (o ContributorResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ErrorCode"] = o.ErrorCode
	}
	if true {
		toSerialize["ErrorMsg"] = o.ErrorMsg
	}
	if true {
		toSerialize["Data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableContributorResponse struct {
	value *ContributorResponse
	isSet bool
}

func (v NullableContributorResponse) Get() *ContributorResponse {
	return v.value
}

func (v *NullableContributorResponse) Set(val *ContributorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableContributorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableContributorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContributorResponse(val *ContributorResponse) *NullableContributorResponse {
	return &NullableContributorResponse{value: val, isSet: true}
}

func (v NullableContributorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContributorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
