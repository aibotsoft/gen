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

// TeaserGroupsBetType struct for TeaserGroupsBetType
type TeaserGroupsBetType struct {
	// Number of points the line will be teased for the given league.
	Points *float64 `json:"points,omitempty"`
}

// NewTeaserGroupsBetType instantiates a new TeaserGroupsBetType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTeaserGroupsBetType() *TeaserGroupsBetType {
	this := TeaserGroupsBetType{}
	return &this
}

// NewTeaserGroupsBetTypeWithDefaults instantiates a new TeaserGroupsBetType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTeaserGroupsBetTypeWithDefaults() *TeaserGroupsBetType {
	this := TeaserGroupsBetType{}
	return &this
}

// GetPoints returns the Points field value if set, zero value otherwise.
func (o *TeaserGroupsBetType) GetPoints() float64 {
	if o == nil || o.Points == nil {
		var ret float64
		return ret
	}
	return *o.Points
}

// GetPointsOk returns a tuple with the Points field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeaserGroupsBetType) GetPointsOk() (*float64, bool) {
	if o == nil || o.Points == nil {
		return nil, false
	}
	return o.Points, true
}

// HasPoints returns a boolean if a field has been set.
func (o *TeaserGroupsBetType) HasPoints() bool {
	if o != nil && o.Points != nil {
		return true
	}

	return false
}

// SetPoints gets a reference to the given float64 and assigns it to the Points field.
func (o *TeaserGroupsBetType) SetPoints(v float64) {
	o.Points = &v
}

func (o TeaserGroupsBetType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Points != nil {
		toSerialize["points"] = o.Points
	}
	return json.Marshal(toSerialize)
}

type NullableTeaserGroupsBetType struct {
	value *TeaserGroupsBetType
	isSet bool
}

func (v NullableTeaserGroupsBetType) Get() *TeaserGroupsBetType {
	return v.value
}

func (v *NullableTeaserGroupsBetType) Set(val *TeaserGroupsBetType) {
	v.value = val
	v.isSet = true
}

func (v NullableTeaserGroupsBetType) IsSet() bool {
	return v.isSet
}

func (v *NullableTeaserGroupsBetType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTeaserGroupsBetType(val *TeaserGroupsBetType) *NullableTeaserGroupsBetType {
	return &NullableTeaserGroupsBetType{value: val, isSet: true}
}

func (v NullableTeaserGroupsBetType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTeaserGroupsBetType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
