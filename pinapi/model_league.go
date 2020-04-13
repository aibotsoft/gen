/*
 * Pinnacle - Customer API Reference
 *
 *  # Authentication   API uses HTTP Basic access authentication. You need to send Authorization HTTP Request header:    `Authorization: Basic <Base64 value of UTF-8 encoded \"username:password\">`  Example:  `Authorization: Basic U03MyOT23YbzMDc6d3c3O1DQ1`
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pinapi

import (
	"bytes"
	"encoding/json"
)

// League struct for League
type League struct {
	// League Id.
	Id *int32 `json:"id,omitempty" xml:"id"`
	// Name of the league.
	Name *string `json:"name,omitempty" xml:"name"`
	// Specifies whether the home team is team1 or team2. You need this information to place a bet.
	HomeTeamType *string `json:"homeTeamType,omitempty" xml:"homeTeamType"`
	// Whether the league currently has events or specials.
	HasOfferings *bool `json:"hasOfferings,omitempty" xml:"hasOfferings"`
	// Represents grouping for the league, usually a region/country
	Container *string `json:"container,omitempty" xml:"container"`
	// Specifies whether you can place parlay round robins on events in this league.
	AllowRoundRobins *bool `json:"allowRoundRobins,omitempty" xml:"allowRoundRobins"`
	// Indicates how many specials are in the given league.
	LeagueSpecialsCount *int32 `json:"leagueSpecialsCount,omitempty" xml:"leagueSpecialsCount"`
	// Indicates how many game specials are in the given league.
	EventSpecialsCount *int32 `json:"eventSpecialsCount,omitempty" xml:"eventSpecialsCount"`
	// Indicates how many events are in the given league.
	EventCount *int32 `json:"eventCount,omitempty" xml:"eventCount"`
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *League) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *League) GetIdOk() (int32, bool) {
	if o == nil || o.Id == nil {
		var ret int32
		return ret, false
	}
	return *o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *League) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *League) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *League) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *League) GetNameOk() (string, bool) {
	if o == nil || o.Name == nil {
		var ret string
		return ret, false
	}
	return *o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *League) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *League) SetName(v string) {
	o.Name = &v
}

// GetHomeTeamType returns the HomeTeamType field value if set, zero value otherwise.
func (o *League) GetHomeTeamType() string {
	if o == nil || o.HomeTeamType == nil {
		var ret string
		return ret
	}
	return *o.HomeTeamType
}

// GetHomeTeamTypeOk returns a tuple with the HomeTeamType field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *League) GetHomeTeamTypeOk() (string, bool) {
	if o == nil || o.HomeTeamType == nil {
		var ret string
		return ret, false
	}
	return *o.HomeTeamType, true
}

// HasHomeTeamType returns a boolean if a field has been set.
func (o *League) HasHomeTeamType() bool {
	if o != nil && o.HomeTeamType != nil {
		return true
	}

	return false
}

// SetHomeTeamType gets a reference to the given string and assigns it to the HomeTeamType field.
func (o *League) SetHomeTeamType(v string) {
	o.HomeTeamType = &v
}

// GetHasOfferings returns the HasOfferings field value if set, zero value otherwise.
func (o *League) GetHasOfferings() bool {
	if o == nil || o.HasOfferings == nil {
		var ret bool
		return ret
	}
	return *o.HasOfferings
}

// GetHasOfferingsOk returns a tuple with the HasOfferings field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *League) GetHasOfferingsOk() (bool, bool) {
	if o == nil || o.HasOfferings == nil {
		var ret bool
		return ret, false
	}
	return *o.HasOfferings, true
}

// HasHasOfferings returns a boolean if a field has been set.
func (o *League) HasHasOfferings() bool {
	if o != nil && o.HasOfferings != nil {
		return true
	}

	return false
}

// SetHasOfferings gets a reference to the given bool and assigns it to the HasOfferings field.
func (o *League) SetHasOfferings(v bool) {
	o.HasOfferings = &v
}

// GetContainer returns the Container field value if set, zero value otherwise.
func (o *League) GetContainer() string {
	if o == nil || o.Container == nil {
		var ret string
		return ret
	}
	return *o.Container
}

// GetContainerOk returns a tuple with the Container field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *League) GetContainerOk() (string, bool) {
	if o == nil || o.Container == nil {
		var ret string
		return ret, false
	}
	return *o.Container, true
}

// HasContainer returns a boolean if a field has been set.
func (o *League) HasContainer() bool {
	if o != nil && o.Container != nil {
		return true
	}

	return false
}

// SetContainer gets a reference to the given string and assigns it to the Container field.
func (o *League) SetContainer(v string) {
	o.Container = &v
}

// GetAllowRoundRobins returns the AllowRoundRobins field value if set, zero value otherwise.
func (o *League) GetAllowRoundRobins() bool {
	if o == nil || o.AllowRoundRobins == nil {
		var ret bool
		return ret
	}
	return *o.AllowRoundRobins
}

// GetAllowRoundRobinsOk returns a tuple with the AllowRoundRobins field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *League) GetAllowRoundRobinsOk() (bool, bool) {
	if o == nil || o.AllowRoundRobins == nil {
		var ret bool
		return ret, false
	}
	return *o.AllowRoundRobins, true
}

// HasAllowRoundRobins returns a boolean if a field has been set.
func (o *League) HasAllowRoundRobins() bool {
	if o != nil && o.AllowRoundRobins != nil {
		return true
	}

	return false
}

// SetAllowRoundRobins gets a reference to the given bool and assigns it to the AllowRoundRobins field.
func (o *League) SetAllowRoundRobins(v bool) {
	o.AllowRoundRobins = &v
}

// GetLeagueSpecialsCount returns the LeagueSpecialsCount field value if set, zero value otherwise.
func (o *League) GetLeagueSpecialsCount() int32 {
	if o == nil || o.LeagueSpecialsCount == nil {
		var ret int32
		return ret
	}
	return *o.LeagueSpecialsCount
}

// GetLeagueSpecialsCountOk returns a tuple with the LeagueSpecialsCount field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *League) GetLeagueSpecialsCountOk() (int32, bool) {
	if o == nil || o.LeagueSpecialsCount == nil {
		var ret int32
		return ret, false
	}
	return *o.LeagueSpecialsCount, true
}

// HasLeagueSpecialsCount returns a boolean if a field has been set.
func (o *League) HasLeagueSpecialsCount() bool {
	if o != nil && o.LeagueSpecialsCount != nil {
		return true
	}

	return false
}

// SetLeagueSpecialsCount gets a reference to the given int32 and assigns it to the LeagueSpecialsCount field.
func (o *League) SetLeagueSpecialsCount(v int32) {
	o.LeagueSpecialsCount = &v
}

// GetEventSpecialsCount returns the EventSpecialsCount field value if set, zero value otherwise.
func (o *League) GetEventSpecialsCount() int32 {
	if o == nil || o.EventSpecialsCount == nil {
		var ret int32
		return ret
	}
	return *o.EventSpecialsCount
}

// GetEventSpecialsCountOk returns a tuple with the EventSpecialsCount field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *League) GetEventSpecialsCountOk() (int32, bool) {
	if o == nil || o.EventSpecialsCount == nil {
		var ret int32
		return ret, false
	}
	return *o.EventSpecialsCount, true
}

// HasEventSpecialsCount returns a boolean if a field has been set.
func (o *League) HasEventSpecialsCount() bool {
	if o != nil && o.EventSpecialsCount != nil {
		return true
	}

	return false
}

// SetEventSpecialsCount gets a reference to the given int32 and assigns it to the EventSpecialsCount field.
func (o *League) SetEventSpecialsCount(v int32) {
	o.EventSpecialsCount = &v
}

// GetEventCount returns the EventCount field value if set, zero value otherwise.
func (o *League) GetEventCount() int32 {
	if o == nil || o.EventCount == nil {
		var ret int32
		return ret
	}
	return *o.EventCount
}

// GetEventCountOk returns a tuple with the EventCount field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *League) GetEventCountOk() (int32, bool) {
	if o == nil || o.EventCount == nil {
		var ret int32
		return ret, false
	}
	return *o.EventCount, true
}

// HasEventCount returns a boolean if a field has been set.
func (o *League) HasEventCount() bool {
	if o != nil && o.EventCount != nil {
		return true
	}

	return false
}

// SetEventCount gets a reference to the given int32 and assigns it to the EventCount field.
func (o *League) SetEventCount(v int32) {
	o.EventCount = &v
}

type NullableLeague struct {
	Value        League
	ExplicitNull bool
}

func (v NullableLeague) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableLeague) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
