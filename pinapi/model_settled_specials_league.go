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

// SettledSpecialsLeague League Dto to hold all settled specials for the league
type SettledSpecialsLeague struct {
	// League Id.
	Id *int `json:"id,omitempty"`
	// A collection of Settled Specials
	Specials *[]SettledSpecial `json:"specials,omitempty"`
}

// NewSettledSpecialsLeague instantiates a new SettledSpecialsLeague object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSettledSpecialsLeague() *SettledSpecialsLeague {
	this := SettledSpecialsLeague{}
	return &this
}

// NewSettledSpecialsLeagueWithDefaults instantiates a new SettledSpecialsLeague object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSettledSpecialsLeagueWithDefaults() *SettledSpecialsLeague {
	this := SettledSpecialsLeague{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SettledSpecialsLeague) GetId() int {
	if o == nil || o.Id == nil {
		var ret int
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettledSpecialsLeague) GetIdOk() (*int, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SettledSpecialsLeague) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int and assigns it to the Id field.
func (o *SettledSpecialsLeague) SetId(v int) {
	o.Id = &v
}

// GetSpecials returns the Specials field value if set, zero value otherwise.
func (o *SettledSpecialsLeague) GetSpecials() []SettledSpecial {
	if o == nil || o.Specials == nil {
		var ret []SettledSpecial
		return ret
	}
	return *o.Specials
}

// GetSpecialsOk returns a tuple with the Specials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettledSpecialsLeague) GetSpecialsOk() (*[]SettledSpecial, bool) {
	if o == nil || o.Specials == nil {
		return nil, false
	}
	return o.Specials, true
}

// HasSpecials returns a boolean if a field has been set.
func (o *SettledSpecialsLeague) HasSpecials() bool {
	if o != nil && o.Specials != nil {
		return true
	}

	return false
}

// SetSpecials gets a reference to the given []SettledSpecial and assigns it to the Specials field.
func (o *SettledSpecialsLeague) SetSpecials(v []SettledSpecial) {
	o.Specials = &v
}

func (o SettledSpecialsLeague) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Specials != nil {
		toSerialize["specials"] = o.Specials
	}
	return json.Marshal(toSerialize)
}

type NullableSettledSpecialsLeague struct {
	value *SettledSpecialsLeague
	isSet bool
}

func (v NullableSettledSpecialsLeague) Get() *SettledSpecialsLeague {
	return v.value
}

func (v *NullableSettledSpecialsLeague) Set(val *SettledSpecialsLeague) {
	v.value = val
	v.isSet = true
}

func (v NullableSettledSpecialsLeague) IsSet() bool {
	return v.isSet
}

func (v *NullableSettledSpecialsLeague) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSettledSpecialsLeague(val *SettledSpecialsLeague) *NullableSettledSpecialsLeague {
	return &NullableSettledSpecialsLeague{value: val, isSet: true}
}

func (v NullableSettledSpecialsLeague) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSettledSpecialsLeague) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
