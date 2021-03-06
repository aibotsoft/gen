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

// FixturesLeague struct for FixturesLeague
type FixturesLeague struct {
	// League ID.
	Id *int `json:"id,omitempty"`
	// League Name.
	Name *string `json:"name,omitempty"`
	// Contains a list of events.
	Events *[]Fixture `json:"events,omitempty"`
}

// NewFixturesLeague instantiates a new FixturesLeague object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFixturesLeague() *FixturesLeague {
	this := FixturesLeague{}
	return &this
}

// NewFixturesLeagueWithDefaults instantiates a new FixturesLeague object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFixturesLeagueWithDefaults() *FixturesLeague {
	this := FixturesLeague{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FixturesLeague) GetId() int {
	if o == nil || o.Id == nil {
		var ret int
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FixturesLeague) GetIdOk() (*int, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FixturesLeague) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int and assigns it to the Id field.
func (o *FixturesLeague) SetId(v int) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FixturesLeague) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FixturesLeague) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FixturesLeague) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FixturesLeague) SetName(v string) {
	o.Name = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *FixturesLeague) GetEvents() []Fixture {
	if o == nil || o.Events == nil {
		var ret []Fixture
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FixturesLeague) GetEventsOk() (*[]Fixture, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *FixturesLeague) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []Fixture and assigns it to the Events field.
func (o *FixturesLeague) SetEvents(v []Fixture) {
	o.Events = &v
}

func (o FixturesLeague) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	return json.Marshal(toSerialize)
}

type NullableFixturesLeague struct {
	value *FixturesLeague
	isSet bool
}

func (v NullableFixturesLeague) Get() *FixturesLeague {
	return v.value
}

func (v *NullableFixturesLeague) Set(val *FixturesLeague) {
	v.value = val
	v.isSet = true
}

func (v NullableFixturesLeague) IsSet() bool {
	return v.isSet
}

func (v *NullableFixturesLeague) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFixturesLeague(val *FixturesLeague) *NullableFixturesLeague {
	return &NullableFixturesLeague{value: val, isSet: true}
}

func (v NullableFixturesLeague) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFixturesLeague) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
