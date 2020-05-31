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

// SportItemM0 struct for SportItemM0
type SportItemM0 struct {
	E int64 `json:"E"`
	T int64 `json:"T"`
	L int64 `json:"L"`
	Days *map[string]int64 `json:"Days,omitempty"`
}

// NewSportItemM0 instantiates a new SportItemM0 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSportItemM0(e int64, t int64, l int64, ) *SportItemM0 {
	this := SportItemM0{}
	this.E = e
	this.T = t
	this.L = l
	return &this
}

// NewSportItemM0WithDefaults instantiates a new SportItemM0 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSportItemM0WithDefaults() *SportItemM0 {
	this := SportItemM0{}
	return &this
}

// GetE returns the E field value
func (o *SportItemM0) GetE() int64 {
	if o == nil  {
		var ret int64
		return ret
	}

	return o.E
}

// GetEOk returns a tuple with the E field value
// and a boolean to check if the value has been set.
func (o *SportItemM0) GetEOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.E, true
}

// SetE sets field value
func (o *SportItemM0) SetE(v int64) {
	o.E = v
}

// GetT returns the T field value
func (o *SportItemM0) GetT() int64 {
	if o == nil  {
		var ret int64
		return ret
	}

	return o.T
}

// GetTOk returns a tuple with the T field value
// and a boolean to check if the value has been set.
func (o *SportItemM0) GetTOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.T, true
}

// SetT sets field value
func (o *SportItemM0) SetT(v int64) {
	o.T = v
}

// GetL returns the L field value
func (o *SportItemM0) GetL() int64 {
	if o == nil  {
		var ret int64
		return ret
	}

	return o.L
}

// GetLOk returns a tuple with the L field value
// and a boolean to check if the value has been set.
func (o *SportItemM0) GetLOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.L, true
}

// SetL sets field value
func (o *SportItemM0) SetL(v int64) {
	o.L = v
}

// GetDays returns the Days field value if set, zero value otherwise.
func (o *SportItemM0) GetDays() map[string]int64 {
	if o == nil || o.Days == nil {
		var ret map[string]int64
		return ret
	}
	return *o.Days
}

// GetDaysOk returns a tuple with the Days field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SportItemM0) GetDaysOk() (*map[string]int64, bool) {
	if o == nil || o.Days == nil {
		return nil, false
	}
	return o.Days, true
}

// HasDays returns a boolean if a field has been set.
func (o *SportItemM0) HasDays() bool {
	if o != nil && o.Days != nil {
		return true
	}

	return false
}

// SetDays gets a reference to the given map[string]int64 and assigns it to the Days field.
func (o *SportItemM0) SetDays(v map[string]int64) {
	o.Days = &v
}

func (o SportItemM0) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["E"] = o.E
	}
	if true {
		toSerialize["T"] = o.T
	}
	if true {
		toSerialize["L"] = o.L
	}
	if o.Days != nil {
		toSerialize["Days"] = o.Days
	}
	return json.Marshal(toSerialize)
}

type NullableSportItemM0 struct {
	value *SportItemM0
	isSet bool
}

func (v NullableSportItemM0) Get() *SportItemM0 {
	return v.value
}

func (v *NullableSportItemM0) Set(val *SportItemM0) {
	v.value = val
	v.isSet = true
}

func (v NullableSportItemM0) IsSet() bool {
	return v.isSet
}

func (v *NullableSportItemM0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSportItemM0(val *SportItemM0) *NullableSportItemM0 {
	return &NullableSportItemM0{value: val, isSet: true}
}

func (v NullableSportItemM0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSportItemM0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}