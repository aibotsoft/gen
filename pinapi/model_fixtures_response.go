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

// FixturesResponse struct for FixturesResponse
type FixturesResponse struct {
	// Same as requested sport Id.
	SportId *int `json:"sportId,omitempty"`
	// Use this value for the subsequent requests for since query parameter to get just the changes since previous response.
	Last *int64 `json:"last,omitempty"`
	// Contains a list of Leagues.
	League *[]FixturesLeague `json:"league,omitempty"`
}

// NewFixturesResponse instantiates a new FixturesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFixturesResponse() *FixturesResponse {
	this := FixturesResponse{}
	return &this
}

// NewFixturesResponseWithDefaults instantiates a new FixturesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFixturesResponseWithDefaults() *FixturesResponse {
	this := FixturesResponse{}
	return &this
}

// GetSportId returns the SportId field value if set, zero value otherwise.
func (o *FixturesResponse) GetSportId() int {
	if o == nil || o.SportId == nil {
		var ret int
		return ret
	}
	return *o.SportId
}

// GetSportIdOk returns a tuple with the SportId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FixturesResponse) GetSportIdOk() (*int, bool) {
	if o == nil || o.SportId == nil {
		return nil, false
	}
	return o.SportId, true
}

// HasSportId returns a boolean if a field has been set.
func (o *FixturesResponse) HasSportId() bool {
	if o != nil && o.SportId != nil {
		return true
	}

	return false
}

// SetSportId gets a reference to the given int and assigns it to the SportId field.
func (o *FixturesResponse) SetSportId(v int) {
	o.SportId = &v
}

// GetLast returns the Last field value if set, zero value otherwise.
func (o *FixturesResponse) GetLast() int64 {
	if o == nil || o.Last == nil {
		var ret int64
		return ret
	}
	return *o.Last
}

// GetLastOk returns a tuple with the Last field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FixturesResponse) GetLastOk() (*int64, bool) {
	if o == nil || o.Last == nil {
		return nil, false
	}
	return o.Last, true
}

// HasLast returns a boolean if a field has been set.
func (o *FixturesResponse) HasLast() bool {
	if o != nil && o.Last != nil {
		return true
	}

	return false
}

// SetLast gets a reference to the given int64 and assigns it to the Last field.
func (o *FixturesResponse) SetLast(v int64) {
	o.Last = &v
}

// GetLeague returns the League field value if set, zero value otherwise.
func (o *FixturesResponse) GetLeague() []FixturesLeague {
	if o == nil || o.League == nil {
		var ret []FixturesLeague
		return ret
	}
	return *o.League
}

// GetLeagueOk returns a tuple with the League field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FixturesResponse) GetLeagueOk() (*[]FixturesLeague, bool) {
	if o == nil || o.League == nil {
		return nil, false
	}
	return o.League, true
}

// HasLeague returns a boolean if a field has been set.
func (o *FixturesResponse) HasLeague() bool {
	if o != nil && o.League != nil {
		return true
	}

	return false
}

// SetLeague gets a reference to the given []FixturesLeague and assigns it to the League field.
func (o *FixturesResponse) SetLeague(v []FixturesLeague) {
	o.League = &v
}

func (o FixturesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SportId != nil {
		toSerialize["sportId"] = o.SportId
	}
	if o.Last != nil {
		toSerialize["last"] = o.Last
	}
	if o.League != nil {
		toSerialize["league"] = o.League
	}
	return json.Marshal(toSerialize)
}

type NullableFixturesResponse struct {
	value *FixturesResponse
	isSet bool
}

func (v NullableFixturesResponse) Get() *FixturesResponse {
	return v.value
}

func (v *NullableFixturesResponse) Set(val *FixturesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFixturesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFixturesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFixturesResponse(val *FixturesResponse) *NullableFixturesResponse {
	return &NullableFixturesResponse{value: val, isSet: true}
}

func (v NullableFixturesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFixturesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
