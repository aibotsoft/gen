/*
 * Pinnacle - Customer API Reference
 *
 *  # Authentication   API uses HTTP Basic access authentication. You need to send Authorization HTTP Request header:    `Authorization: Basic <Base64 value of UTF-8 encoded \"username:password\">`  Example:  `Authorization: Basic U03MyOT23YbzMDc6d3c3O1DQ1` 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"encoding/json"
)

// SpecialsFixturesContestant struct for SpecialsFixturesContestant
type SpecialsFixturesContestant struct {
	// Contestant Id.
	Id *int64 `json:"id,omitempty"`
	// Name of the contestant.
	Name *string `json:"name,omitempty"`
	// Rotation Number.
	RotNum *int `json:"rotNum,omitempty"`
}

// NewSpecialsFixturesContestant instantiates a new SpecialsFixturesContestant object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpecialsFixturesContestant() *SpecialsFixturesContestant {
	this := SpecialsFixturesContestant{}
	return &this
}

// NewSpecialsFixturesContestantWithDefaults instantiates a new SpecialsFixturesContestant object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpecialsFixturesContestantWithDefaults() *SpecialsFixturesContestant {
	this := SpecialsFixturesContestant{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SpecialsFixturesContestant) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpecialsFixturesContestant) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SpecialsFixturesContestant) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *SpecialsFixturesContestant) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SpecialsFixturesContestant) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpecialsFixturesContestant) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SpecialsFixturesContestant) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SpecialsFixturesContestant) SetName(v string) {
	o.Name = &v
}

// GetRotNum returns the RotNum field value if set, zero value otherwise.
func (o *SpecialsFixturesContestant) GetRotNum() int {
	if o == nil || o.RotNum == nil {
		var ret int
		return ret
	}
	return *o.RotNum
}

// GetRotNumOk returns a tuple with the RotNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpecialsFixturesContestant) GetRotNumOk() (*int, bool) {
	if o == nil || o.RotNum == nil {
		return nil, false
	}
	return o.RotNum, true
}

// HasRotNum returns a boolean if a field has been set.
func (o *SpecialsFixturesContestant) HasRotNum() bool {
	if o != nil && o.RotNum != nil {
		return true
	}

	return false
}

// SetRotNum gets a reference to the given int and assigns it to the RotNum field.
func (o *SpecialsFixturesContestant) SetRotNum(v int) {
	o.RotNum = &v
}

func (o SpecialsFixturesContestant) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.RotNum != nil {
		toSerialize["rotNum"] = o.RotNum
	}
	return json.Marshal(toSerialize)
}

type NullableSpecialsFixturesContestant struct {
	value *SpecialsFixturesContestant
	isSet bool
}

func (v NullableSpecialsFixturesContestant) Get() *SpecialsFixturesContestant {
	return v.value
}

func (v *NullableSpecialsFixturesContestant) Set(val *SpecialsFixturesContestant) {
	v.value = val
	v.isSet = true
}

func (v NullableSpecialsFixturesContestant) IsSet() bool {
	return v.isSet
}

func (v *NullableSpecialsFixturesContestant) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpecialsFixturesContestant(val *SpecialsFixturesContestant) *NullableSpecialsFixturesContestant {
	return &NullableSpecialsFixturesContestant{value: val, isSet: true}
}

func (v NullableSpecialsFixturesContestant) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpecialsFixturesContestant) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
