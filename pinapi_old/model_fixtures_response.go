/*
 * Pinnacle - Customer API Reference
 *
 *  # Authentication   API uses HTTP Basic access authentication. You need to send Authorization HTTP Request header:    `Authorization: Basic <Base64 value of UTF-8 encoded \"username:password\">`  Example:  `Authorization: Basic U03MyOT23YbzMDc6d3c3O1DQ1`
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pinapi_old

import (
	"bytes"
	"encoding/json"
)

// FixturesResponse struct for FixturesResponse
type FixturesResponse struct {
	// Same as requested sport Id.
	SportId *int32 `json:"sportId,omitempty" xml:"sportId"`
	// Use this value for the subsequent requests for since query parameter to get just the changes since previous response.
	Last *int64 `json:"last,omitempty" xml:"last"`
	// Contains a list of Leagues.
	League *[]FixturesLeague `json:"league,omitempty" xml:"league"`
}

// GetSportId returns the SportId field value if set, zero value otherwise.
func (o *FixturesResponse) GetSportId() int32 {
	if o == nil || o.SportId == nil {
		var ret int32
		return ret
	}
	return *o.SportId
}

// GetSportIdOk returns a tuple with the SportId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FixturesResponse) GetSportIdOk() (int32, bool) {
	if o == nil || o.SportId == nil {
		var ret int32
		return ret, false
	}
	return *o.SportId, true
}

// HasSportId returns a boolean if a field has been set.
func (o *FixturesResponse) HasSportId() bool {
	if o != nil && o.SportId != nil {
		return true
	}

	return false
}

// SetSportId gets a reference to the given int32 and assigns it to the SportId field.
func (o *FixturesResponse) SetSportId(v int32) {
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

// GetLastOk returns a tuple with the Last field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FixturesResponse) GetLastOk() (int64, bool) {
	if o == nil || o.Last == nil {
		var ret int64
		return ret, false
	}
	return *o.Last, true
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

// GetLeagueOk returns a tuple with the League field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FixturesResponse) GetLeagueOk() ([]FixturesLeague, bool) {
	if o == nil || o.League == nil {
		var ret []FixturesLeague
		return ret, false
	}
	return *o.League, true
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

type NullableFixturesResponse struct {
	Value        FixturesResponse
	ExplicitNull bool
}

func (v NullableFixturesResponse) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableFixturesResponse) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
