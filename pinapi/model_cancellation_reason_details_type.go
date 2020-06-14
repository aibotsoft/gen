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

// CancellationReasonDetailsType struct for CancellationReasonDetailsType
type CancellationReasonDetailsType struct {
	Key *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewCancellationReasonDetailsType instantiates a new CancellationReasonDetailsType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCancellationReasonDetailsType() *CancellationReasonDetailsType {
	this := CancellationReasonDetailsType{}
	return &this
}

// NewCancellationReasonDetailsTypeWithDefaults instantiates a new CancellationReasonDetailsType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCancellationReasonDetailsTypeWithDefaults() *CancellationReasonDetailsType {
	this := CancellationReasonDetailsType{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *CancellationReasonDetailsType) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancellationReasonDetailsType) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *CancellationReasonDetailsType) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *CancellationReasonDetailsType) SetKey(v string) {
	o.Key = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CancellationReasonDetailsType) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancellationReasonDetailsType) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CancellationReasonDetailsType) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *CancellationReasonDetailsType) SetValue(v string) {
	o.Value = &v
}

func (o CancellationReasonDetailsType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableCancellationReasonDetailsType struct {
	value *CancellationReasonDetailsType
	isSet bool
}

func (v NullableCancellationReasonDetailsType) Get() *CancellationReasonDetailsType {
	return v.value
}

func (v *NullableCancellationReasonDetailsType) Set(val *CancellationReasonDetailsType) {
	v.value = val
	v.isSet = true
}

func (v NullableCancellationReasonDetailsType) IsSet() bool {
	return v.isSet
}

func (v *NullableCancellationReasonDetailsType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCancellationReasonDetailsType(val *CancellationReasonDetailsType) *NullableCancellationReasonDetailsType {
	return &NullableCancellationReasonDetailsType{value: val, isSet: true}
}

func (v NullableCancellationReasonDetailsType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCancellationReasonDetailsType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}