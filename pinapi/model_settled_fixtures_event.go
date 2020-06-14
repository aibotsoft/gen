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

// SettledFixturesEvent struct for SettledFixturesEvent
type SettledFixturesEvent struct {
	// Event Id.
	Id *int64 `json:"id,omitempty"`
	// Contains a list of periods.
	Periods *[]SettledFixturesPeriod `json:"periods,omitempty"`
}

// NewSettledFixturesEvent instantiates a new SettledFixturesEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSettledFixturesEvent() *SettledFixturesEvent {
	this := SettledFixturesEvent{}
	return &this
}

// NewSettledFixturesEventWithDefaults instantiates a new SettledFixturesEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSettledFixturesEventWithDefaults() *SettledFixturesEvent {
	this := SettledFixturesEvent{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SettledFixturesEvent) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettledFixturesEvent) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SettledFixturesEvent) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *SettledFixturesEvent) SetId(v int64) {
	o.Id = &v
}

// GetPeriods returns the Periods field value if set, zero value otherwise.
func (o *SettledFixturesEvent) GetPeriods() []SettledFixturesPeriod {
	if o == nil || o.Periods == nil {
		var ret []SettledFixturesPeriod
		return ret
	}
	return *o.Periods
}

// GetPeriodsOk returns a tuple with the Periods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettledFixturesEvent) GetPeriodsOk() (*[]SettledFixturesPeriod, bool) {
	if o == nil || o.Periods == nil {
		return nil, false
	}
	return o.Periods, true
}

// HasPeriods returns a boolean if a field has been set.
func (o *SettledFixturesEvent) HasPeriods() bool {
	if o != nil && o.Periods != nil {
		return true
	}

	return false
}

// SetPeriods gets a reference to the given []SettledFixturesPeriod and assigns it to the Periods field.
func (o *SettledFixturesEvent) SetPeriods(v []SettledFixturesPeriod) {
	o.Periods = &v
}

func (o SettledFixturesEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Periods != nil {
		toSerialize["periods"] = o.Periods
	}
	return json.Marshal(toSerialize)
}

type NullableSettledFixturesEvent struct {
	value *SettledFixturesEvent
	isSet bool
}

func (v NullableSettledFixturesEvent) Get() *SettledFixturesEvent {
	return v.value
}

func (v *NullableSettledFixturesEvent) Set(val *SettledFixturesEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableSettledFixturesEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableSettledFixturesEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSettledFixturesEvent(val *SettledFixturesEvent) *NullableSettledFixturesEvent {
	return &NullableSettledFixturesEvent{value: val, isSet: true}
}

func (v NullableSettledFixturesEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSettledFixturesEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}