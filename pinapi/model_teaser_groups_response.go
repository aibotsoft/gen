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

// TeaserGroupsResponse struct for TeaserGroupsResponse
type TeaserGroupsResponse struct {
	// A collection of TeaserGroups containing available teasers.
	TeaserGroups *[]TeaserGroups `json:"teaserGroups,omitempty"`
}

// NewTeaserGroupsResponse instantiates a new TeaserGroupsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTeaserGroupsResponse() *TeaserGroupsResponse {
	this := TeaserGroupsResponse{}
	return &this
}

// NewTeaserGroupsResponseWithDefaults instantiates a new TeaserGroupsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTeaserGroupsResponseWithDefaults() *TeaserGroupsResponse {
	this := TeaserGroupsResponse{}
	return &this
}

// GetTeaserGroups returns the TeaserGroups field value if set, zero value otherwise.
func (o *TeaserGroupsResponse) GetTeaserGroups() []TeaserGroups {
	if o == nil || o.TeaserGroups == nil {
		var ret []TeaserGroups
		return ret
	}
	return *o.TeaserGroups
}

// GetTeaserGroupsOk returns a tuple with the TeaserGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeaserGroupsResponse) GetTeaserGroupsOk() (*[]TeaserGroups, bool) {
	if o == nil || o.TeaserGroups == nil {
		return nil, false
	}
	return o.TeaserGroups, true
}

// HasTeaserGroups returns a boolean if a field has been set.
func (o *TeaserGroupsResponse) HasTeaserGroups() bool {
	if o != nil && o.TeaserGroups != nil {
		return true
	}

	return false
}

// SetTeaserGroups gets a reference to the given []TeaserGroups and assigns it to the TeaserGroups field.
func (o *TeaserGroupsResponse) SetTeaserGroups(v []TeaserGroups) {
	o.TeaserGroups = &v
}

func (o TeaserGroupsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TeaserGroups != nil {
		toSerialize["teaserGroups"] = o.TeaserGroups
	}
	return json.Marshal(toSerialize)
}

type NullableTeaserGroupsResponse struct {
	value *TeaserGroupsResponse
	isSet bool
}

func (v NullableTeaserGroupsResponse) Get() *TeaserGroupsResponse {
	return v.value
}

func (v *NullableTeaserGroupsResponse) Set(val *TeaserGroupsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTeaserGroupsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTeaserGroupsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTeaserGroupsResponse(val *TeaserGroupsResponse) *NullableTeaserGroupsResponse {
	return &NullableTeaserGroupsResponse{value: val, isSet: true}
}

func (v NullableTeaserGroupsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTeaserGroupsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
