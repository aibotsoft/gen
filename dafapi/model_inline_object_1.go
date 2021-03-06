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

// InlineObject1 struct for InlineObject1
type InlineObject1 struct {
	Fdate *string `json:"fdate,omitempty"`
	Datatype *int64 `json:"datatype,omitempty"`
}

// NewInlineObject1 instantiates a new InlineObject1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1() *InlineObject1 {
	this := InlineObject1{}
	return &this
}

// NewInlineObject1WithDefaults instantiates a new InlineObject1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1WithDefaults() *InlineObject1 {
	this := InlineObject1{}
	return &this
}

// GetFdate returns the Fdate field value if set, zero value otherwise.
func (o *InlineObject1) GetFdate() string {
	if o == nil || o.Fdate == nil {
		var ret string
		return ret
	}
	return *o.Fdate
}

// GetFdateOk returns a tuple with the Fdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1) GetFdateOk() (*string, bool) {
	if o == nil || o.Fdate == nil {
		return nil, false
	}
	return o.Fdate, true
}

// HasFdate returns a boolean if a field has been set.
func (o *InlineObject1) HasFdate() bool {
	if o != nil && o.Fdate != nil {
		return true
	}

	return false
}

// SetFdate gets a reference to the given string and assigns it to the Fdate field.
func (o *InlineObject1) SetFdate(v string) {
	o.Fdate = &v
}

// GetDatatype returns the Datatype field value if set, zero value otherwise.
func (o *InlineObject1) GetDatatype() int64 {
	if o == nil || o.Datatype == nil {
		var ret int64
		return ret
	}
	return *o.Datatype
}

// GetDatatypeOk returns a tuple with the Datatype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1) GetDatatypeOk() (*int64, bool) {
	if o == nil || o.Datatype == nil {
		return nil, false
	}
	return o.Datatype, true
}

// HasDatatype returns a boolean if a field has been set.
func (o *InlineObject1) HasDatatype() bool {
	if o != nil && o.Datatype != nil {
		return true
	}

	return false
}

// SetDatatype gets a reference to the given int64 and assigns it to the Datatype field.
func (o *InlineObject1) SetDatatype(v int64) {
	o.Datatype = &v
}

func (o InlineObject1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Fdate != nil {
		toSerialize["fdate"] = o.Fdate
	}
	if o.Datatype != nil {
		toSerialize["datatype"] = o.Datatype
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1 struct {
	value *InlineObject1
	isSet bool
}

func (v NullableInlineObject1) Get() *InlineObject1 {
	return v.value
}

func (v *NullableInlineObject1) Set(val *InlineObject1) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1(val *InlineObject1) *NullableInlineObject1 {
	return &NullableInlineObject1{value: val, isSet: true}
}

func (v NullableInlineObject1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
