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

// SetOddsTypeResponseData struct for SetOddsTypeResponseData
type SetOddsTypeResponseData struct {
	OddsType *int64 `json:"OddsType,omitempty"`
}

// NewSetOddsTypeResponseData instantiates a new SetOddsTypeResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetOddsTypeResponseData() *SetOddsTypeResponseData {
	this := SetOddsTypeResponseData{}
	return &this
}

// NewSetOddsTypeResponseDataWithDefaults instantiates a new SetOddsTypeResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetOddsTypeResponseDataWithDefaults() *SetOddsTypeResponseData {
	this := SetOddsTypeResponseData{}
	return &this
}

// GetOddsType returns the OddsType field value if set, zero value otherwise.
func (o *SetOddsTypeResponseData) GetOddsType() int64 {
	if o == nil || o.OddsType == nil {
		var ret int64
		return ret
	}
	return *o.OddsType
}

// GetOddsTypeOk returns a tuple with the OddsType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetOddsTypeResponseData) GetOddsTypeOk() (*int64, bool) {
	if o == nil || o.OddsType == nil {
		return nil, false
	}
	return o.OddsType, true
}

// HasOddsType returns a boolean if a field has been set.
func (o *SetOddsTypeResponseData) HasOddsType() bool {
	if o != nil && o.OddsType != nil {
		return true
	}

	return false
}

// SetOddsType gets a reference to the given int64 and assigns it to the OddsType field.
func (o *SetOddsTypeResponseData) SetOddsType(v int64) {
	o.OddsType = &v
}

func (o SetOddsTypeResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.OddsType != nil {
		toSerialize["OddsType"] = o.OddsType
	}
	return json.Marshal(toSerialize)
}

type NullableSetOddsTypeResponseData struct {
	value *SetOddsTypeResponseData
	isSet bool
}

func (v NullableSetOddsTypeResponseData) Get() *SetOddsTypeResponseData {
	return v.value
}

func (v *NullableSetOddsTypeResponseData) Set(val *SetOddsTypeResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableSetOddsTypeResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableSetOddsTypeResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetOddsTypeResponseData(val *SetOddsTypeResponseData) *NullableSetOddsTypeResponseData {
	return &NullableSetOddsTypeResponseData{value: val, isSet: true}
}

func (v NullableSetOddsTypeResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetOddsTypeResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
