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

// BetStatusResponse struct for BetStatusResponse
type BetStatusResponse struct {
	ErrorCode int64 `json:"ErrorCode"`
	ErrorMsg string `json:"ErrorMsg"`
	Data BetStatusResponseData `json:"Data"`
}

// NewBetStatusResponse instantiates a new BetStatusResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBetStatusResponse(errorCode int64, errorMsg string, data BetStatusResponseData, ) *BetStatusResponse {
	this := BetStatusResponse{}
	this.ErrorCode = errorCode
	this.ErrorMsg = errorMsg
	this.Data = data
	return &this
}

// NewBetStatusResponseWithDefaults instantiates a new BetStatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBetStatusResponseWithDefaults() *BetStatusResponse {
	this := BetStatusResponse{}
	return &this
}

// GetErrorCode returns the ErrorCode field value
func (o *BetStatusResponse) GetErrorCode() int64 {
	if o == nil  {
		var ret int64
		return ret
	}

	return o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value
// and a boolean to check if the value has been set.
func (o *BetStatusResponse) GetErrorCodeOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ErrorCode, true
}

// SetErrorCode sets field value
func (o *BetStatusResponse) SetErrorCode(v int64) {
	o.ErrorCode = v
}

// GetErrorMsg returns the ErrorMsg field value
func (o *BetStatusResponse) GetErrorMsg() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ErrorMsg
}

// GetErrorMsgOk returns a tuple with the ErrorMsg field value
// and a boolean to check if the value has been set.
func (o *BetStatusResponse) GetErrorMsgOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ErrorMsg, true
}

// SetErrorMsg sets field value
func (o *BetStatusResponse) SetErrorMsg(v string) {
	o.ErrorMsg = v
}

// GetData returns the Data field value
func (o *BetStatusResponse) GetData() BetStatusResponseData {
	if o == nil  {
		var ret BetStatusResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BetStatusResponse) GetDataOk() (*BetStatusResponseData, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *BetStatusResponse) SetData(v BetStatusResponseData) {
	o.Data = v
}

func (o BetStatusResponse) MarshalJSON() ([]byte, error) {
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

type NullableBetStatusResponse struct {
	value *BetStatusResponse
	isSet bool
}

func (v NullableBetStatusResponse) Get() *BetStatusResponse {
	return v.value
}

func (v *NullableBetStatusResponse) Set(val *BetStatusResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBetStatusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBetStatusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetStatusResponse(val *BetStatusResponse) *NullableBetStatusResponse {
	return &NullableBetStatusResponse{value: val, isSet: true}
}

func (v NullableBetStatusResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetStatusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
