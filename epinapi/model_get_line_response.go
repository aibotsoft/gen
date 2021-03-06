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

// GetLineResponse struct for GetLineResponse
type GetLineResponse struct {
	Classes *[]ClassItem `json:"classes,omitempty"`
	Limits *[]LimitItem `json:"limits,omitempty"`
	Selections *[]SelectionItem `json:"selections,omitempty"`
}

// NewGetLineResponse instantiates a new GetLineResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetLineResponse() *GetLineResponse {
	this := GetLineResponse{}
	return &this
}

// NewGetLineResponseWithDefaults instantiates a new GetLineResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetLineResponseWithDefaults() *GetLineResponse {
	this := GetLineResponse{}
	return &this
}

// GetClasses returns the Classes field value if set, zero value otherwise.
func (o *GetLineResponse) GetClasses() []ClassItem {
	if o == nil || o.Classes == nil {
		var ret []ClassItem
		return ret
	}
	return *o.Classes
}

// GetClassesOk returns a tuple with the Classes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLineResponse) GetClassesOk() (*[]ClassItem, bool) {
	if o == nil || o.Classes == nil {
		return nil, false
	}
	return o.Classes, true
}

// HasClasses returns a boolean if a field has been set.
func (o *GetLineResponse) HasClasses() bool {
	if o != nil && o.Classes != nil {
		return true
	}

	return false
}

// SetClasses gets a reference to the given []ClassItem and assigns it to the Classes field.
func (o *GetLineResponse) SetClasses(v []ClassItem) {
	o.Classes = &v
}

// GetLimits returns the Limits field value if set, zero value otherwise.
func (o *GetLineResponse) GetLimits() []LimitItem {
	if o == nil || o.Limits == nil {
		var ret []LimitItem
		return ret
	}
	return *o.Limits
}

// GetLimitsOk returns a tuple with the Limits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLineResponse) GetLimitsOk() (*[]LimitItem, bool) {
	if o == nil || o.Limits == nil {
		return nil, false
	}
	return o.Limits, true
}

// HasLimits returns a boolean if a field has been set.
func (o *GetLineResponse) HasLimits() bool {
	if o != nil && o.Limits != nil {
		return true
	}

	return false
}

// SetLimits gets a reference to the given []LimitItem and assigns it to the Limits field.
func (o *GetLineResponse) SetLimits(v []LimitItem) {
	o.Limits = &v
}

// GetSelections returns the Selections field value if set, zero value otherwise.
func (o *GetLineResponse) GetSelections() []SelectionItem {
	if o == nil || o.Selections == nil {
		var ret []SelectionItem
		return ret
	}
	return *o.Selections
}

// GetSelectionsOk returns a tuple with the Selections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLineResponse) GetSelectionsOk() (*[]SelectionItem, bool) {
	if o == nil || o.Selections == nil {
		return nil, false
	}
	return o.Selections, true
}

// HasSelections returns a boolean if a field has been set.
func (o *GetLineResponse) HasSelections() bool {
	if o != nil && o.Selections != nil {
		return true
	}

	return false
}

// SetSelections gets a reference to the given []SelectionItem and assigns it to the Selections field.
func (o *GetLineResponse) SetSelections(v []SelectionItem) {
	o.Selections = &v
}

func (o GetLineResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Classes != nil {
		toSerialize["classes"] = o.Classes
	}
	if o.Limits != nil {
		toSerialize["limits"] = o.Limits
	}
	if o.Selections != nil {
		toSerialize["selections"] = o.Selections
	}
	return json.Marshal(toSerialize)
}

type NullableGetLineResponse struct {
	value *GetLineResponse
	isSet bool
}

func (v NullableGetLineResponse) Get() *GetLineResponse {
	return v.value
}

func (v *NullableGetLineResponse) Set(val *GetLineResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetLineResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetLineResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetLineResponse(val *GetLineResponse) *NullableGetLineResponse {
	return &NullableGetLineResponse{value: val, isSet: true}
}

func (v NullableGetLineResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetLineResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
