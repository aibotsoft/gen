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

// SettledFixturesSport struct for SettledFixturesSport
type SettledFixturesSport struct {
	// Same as requested sport Id.
	SportId *int `json:"sportId,omitempty"`
	// Use this value for the subsequent requests for since query parameter to get just the changes since previous response.
	Last *int64 `json:"last,omitempty"`
	// Contains a list of Leagues.
	Leagues *[]SettledFixturesLeague `json:"leagues,omitempty"`
}

// NewSettledFixturesSport instantiates a new SettledFixturesSport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSettledFixturesSport() *SettledFixturesSport {
	this := SettledFixturesSport{}
	return &this
}

// NewSettledFixturesSportWithDefaults instantiates a new SettledFixturesSport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSettledFixturesSportWithDefaults() *SettledFixturesSport {
	this := SettledFixturesSport{}
	return &this
}

// GetSportId returns the SportId field value if set, zero value otherwise.
func (o *SettledFixturesSport) GetSportId() int {
	if o == nil || o.SportId == nil {
		var ret int
		return ret
	}
	return *o.SportId
}

// GetSportIdOk returns a tuple with the SportId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettledFixturesSport) GetSportIdOk() (*int, bool) {
	if o == nil || o.SportId == nil {
		return nil, false
	}
	return o.SportId, true
}

// HasSportId returns a boolean if a field has been set.
func (o *SettledFixturesSport) HasSportId() bool {
	if o != nil && o.SportId != nil {
		return true
	}

	return false
}

// SetSportId gets a reference to the given int and assigns it to the SportId field.
func (o *SettledFixturesSport) SetSportId(v int) {
	o.SportId = &v
}

// GetLast returns the Last field value if set, zero value otherwise.
func (o *SettledFixturesSport) GetLast() int64 {
	if o == nil || o.Last == nil {
		var ret int64
		return ret
	}
	return *o.Last
}

// GetLastOk returns a tuple with the Last field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettledFixturesSport) GetLastOk() (*int64, bool) {
	if o == nil || o.Last == nil {
		return nil, false
	}
	return o.Last, true
}

// HasLast returns a boolean if a field has been set.
func (o *SettledFixturesSport) HasLast() bool {
	if o != nil && o.Last != nil {
		return true
	}

	return false
}

// SetLast gets a reference to the given int64 and assigns it to the Last field.
func (o *SettledFixturesSport) SetLast(v int64) {
	o.Last = &v
}

// GetLeagues returns the Leagues field value if set, zero value otherwise.
func (o *SettledFixturesSport) GetLeagues() []SettledFixturesLeague {
	if o == nil || o.Leagues == nil {
		var ret []SettledFixturesLeague
		return ret
	}
	return *o.Leagues
}

// GetLeaguesOk returns a tuple with the Leagues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettledFixturesSport) GetLeaguesOk() (*[]SettledFixturesLeague, bool) {
	if o == nil || o.Leagues == nil {
		return nil, false
	}
	return o.Leagues, true
}

// HasLeagues returns a boolean if a field has been set.
func (o *SettledFixturesSport) HasLeagues() bool {
	if o != nil && o.Leagues != nil {
		return true
	}

	return false
}

// SetLeagues gets a reference to the given []SettledFixturesLeague and assigns it to the Leagues field.
func (o *SettledFixturesSport) SetLeagues(v []SettledFixturesLeague) {
	o.Leagues = &v
}

func (o SettledFixturesSport) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SportId != nil {
		toSerialize["sportId"] = o.SportId
	}
	if o.Last != nil {
		toSerialize["last"] = o.Last
	}
	if o.Leagues != nil {
		toSerialize["leagues"] = o.Leagues
	}
	return json.Marshal(toSerialize)
}

type NullableSettledFixturesSport struct {
	value *SettledFixturesSport
	isSet bool
}

func (v NullableSettledFixturesSport) Get() *SettledFixturesSport {
	return v.value
}

func (v *NullableSettledFixturesSport) Set(val *SettledFixturesSport) {
	v.value = val
	v.isSet = true
}

func (v NullableSettledFixturesSport) IsSet() bool {
	return v.isSet
}

func (v *NullableSettledFixturesSport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSettledFixturesSport(val *SettledFixturesSport) *NullableSettledFixturesSport {
	return &NullableSettledFixturesSport{value: val, isSet: true}
}

func (v NullableSettledFixturesSport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSettledFixturesSport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
