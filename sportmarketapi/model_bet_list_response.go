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

// BetListResponse struct for BetListResponse
type BetListResponse struct {
	Status string `json:"status"`
	Data []OrderItem `json:"data"`
}

// NewBetListResponse instantiates a new BetListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBetListResponse(status string, data []OrderItem, ) *BetListResponse {
	this := BetListResponse{}
	this.Status = status
	this.Data = data
	return &this
}

// NewBetListResponseWithDefaults instantiates a new BetListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBetListResponseWithDefaults() *BetListResponse {
	this := BetListResponse{}
	return &this
}

// GetStatus returns the Status field value
func (o *BetListResponse) GetStatus() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *BetListResponse) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *BetListResponse) SetStatus(v string) {
	o.Status = v
}

// GetData returns the Data field value
func (o *BetListResponse) GetData() []OrderItem {
	if o == nil  {
		var ret []OrderItem
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BetListResponse) GetDataOk() (*[]OrderItem, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *BetListResponse) SetData(v []OrderItem) {
	o.Data = v
}

func (o BetListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableBetListResponse struct {
	value *BetListResponse
	isSet bool
}

func (v NullableBetListResponse) Get() *BetListResponse {
	return v.value
}

func (v *NullableBetListResponse) Set(val *BetListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBetListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBetListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetListResponse(val *BetListResponse) *NullableBetListResponse {
	return &NullableBetListResponse{value: val, isSet: true}
}

func (v NullableBetListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
