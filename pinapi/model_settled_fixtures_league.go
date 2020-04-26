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

// SettledFixturesLeague struct for SettledFixturesLeague
type SettledFixturesLeague struct {
	// League Id.
	Id *int `json:"id,omitempty"`
	// Contains a list of events.
	Events *[]SettledFixturesEvent `json:"events,omitempty"`
}

// NewSettledFixturesLeague instantiates a new SettledFixturesLeague object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSettledFixturesLeague() *SettledFixturesLeague {
	this := SettledFixturesLeague{}
	return &this
}

// NewSettledFixturesLeagueWithDefaults instantiates a new SettledFixturesLeague object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSettledFixturesLeagueWithDefaults() *SettledFixturesLeague {
	this := SettledFixturesLeague{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SettledFixturesLeague) GetId() int {
	if o == nil || o.Id == nil {
		var ret int
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettledFixturesLeague) GetIdOk() (*int, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SettledFixturesLeague) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int and assigns it to the Id field.
func (o *SettledFixturesLeague) SetId(v int) {
	o.Id = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *SettledFixturesLeague) GetEvents() []SettledFixturesEvent {
	if o == nil || o.Events == nil {
		var ret []SettledFixturesEvent
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettledFixturesLeague) GetEventsOk() (*[]SettledFixturesEvent, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *SettledFixturesLeague) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []SettledFixturesEvent and assigns it to the Events field.
func (o *SettledFixturesLeague) SetEvents(v []SettledFixturesEvent) {
	o.Events = &v
}

func (o SettledFixturesLeague) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	return json.Marshal(toSerialize)
}

type NullableSettledFixturesLeague struct {
	value *SettledFixturesLeague
	isSet bool
}

func (v NullableSettledFixturesLeague) Get() *SettledFixturesLeague {
	return v.value
}

func (v *NullableSettledFixturesLeague) Set(val *SettledFixturesLeague) {
	v.value = val
	v.isSet = true
}

func (v NullableSettledFixturesLeague) IsSet() bool {
	return v.isSet
}

func (v *NullableSettledFixturesLeague) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSettledFixturesLeague(val *SettledFixturesLeague) *NullableSettledFixturesLeague {
	return &NullableSettledFixturesLeague{value: val, isSet: true}
}

func (v NullableSettledFixturesLeague) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSettledFixturesLeague) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
