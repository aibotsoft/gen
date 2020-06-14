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
	"time"
)

// TeaserLeg struct for TeaserLeg
type TeaserLeg struct {
	SportId *int `json:"sportId,omitempty"`
	// Teaser leg type.
	LegBetType *string `json:"legBetType,omitempty"`
	// CANCELLED = The leg is canceled- the stake on this leg will be transferred to the next one. In this case the leg will be ignored when calculating the winLoss,   LOSE = The leg is a loss or a push-lose. When Push-lose happens, the half of the stake on the leg will be pushed to the next leg, and the other half will be a lose. This can happen only when the leg is placed on a quarter points handicap,   PUSH = The leg is a push - the stake on this leg will be transferred to the next one. In this case the leg will be ignored when calculating the winLoss,   REFUNDED = The leg is refunded - the stake on this leg will be transferred to the next one. In this case the leg will be ignored when calculating the winLoss,   WON = The leg is a won or a push-won. When Push-won happens, the half of the stake on the leg will be pushed to the next leg, and the other half is won. This can happen only when the leg is placed on a quarter points handicap   
	LegBetStatus *string `json:"legBetStatus,omitempty"`
	LeagueId *int `json:"leagueId,omitempty"`
	EventId *int64 `json:"eventId,omitempty"`
	// Date time when the event starts.
	EventStartTime *time.Time `json:"eventStartTime,omitempty"`
	Handicap *float64 `json:"handicap,omitempty"`
	TeamName *string `json:"teamName,omitempty"`
	// Side type.
	Side *string `json:"side,omitempty"`
	Team1 *string `json:"team1,omitempty"`
	Team2 *string `json:"team2,omitempty"`
	PeriodNumber *int `json:"periodNumber,omitempty"`
}

// NewTeaserLeg instantiates a new TeaserLeg object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTeaserLeg() *TeaserLeg {
	this := TeaserLeg{}
	return &this
}

// NewTeaserLegWithDefaults instantiates a new TeaserLeg object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTeaserLegWithDefaults() *TeaserLeg {
	this := TeaserLeg{}
	return &this
}

// GetSportId returns the SportId field value if set, zero value otherwise.
func (o *TeaserLeg) GetSportId() int {
	if o == nil || o.SportId == nil {
		var ret int
		return ret
	}
	return *o.SportId
}

// GetSportIdOk returns a tuple with the SportId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeaserLeg) GetSportIdOk() (*int, bool) {
	if o == nil || o.SportId == nil {
		return nil, false
	}
	return o.SportId, true
}

// HasSportId returns a boolean if a field has been set.
func (o *TeaserLeg) HasSportId() bool {
	if o != nil && o.SportId != nil {
		return true
	}

	return false
}

// SetSportId gets a reference to the given int and assigns it to the SportId field.
func (o *TeaserLeg) SetSportId(v int) {
	o.SportId = &v
}

// GetLegBetType returns the LegBetType field value if set, zero value otherwise.
func (o *TeaserLeg) GetLegBetType() string {
	if o == nil || o.LegBetType == nil {
		var ret string
		return ret
	}
	return *o.LegBetType
}

// GetLegBetTypeOk returns a tuple with the LegBetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeaserLeg) GetLegBetTypeOk() (*string, bool) {
	if o == nil || o.LegBetType == nil {
		return nil, false
	}
	return o.LegBetType, true
}

// HasLegBetType returns a boolean if a field has been set.
func (o *TeaserLeg) HasLegBetType() bool {
	if o != nil && o.LegBetType != nil {
		return true
	}

	return false
}

// SetLegBetType gets a reference to the given string and assigns it to the LegBetType field.
func (o *TeaserLeg) SetLegBetType(v string) {
	o.LegBetType = &v
}

// GetLegBetStatus returns the LegBetStatus field value if set, zero value otherwise.
func (o *TeaserLeg) GetLegBetStatus() string {
	if o == nil || o.LegBetStatus == nil {
		var ret string
		return ret
	}
	return *o.LegBetStatus
}

// GetLegBetStatusOk returns a tuple with the LegBetStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeaserLeg) GetLegBetStatusOk() (*string, bool) {
	if o == nil || o.LegBetStatus == nil {
		return nil, false
	}
	return o.LegBetStatus, true
}

// HasLegBetStatus returns a boolean if a field has been set.
func (o *TeaserLeg) HasLegBetStatus() bool {
	if o != nil && o.LegBetStatus != nil {
		return true
	}

	return false
}

// SetLegBetStatus gets a reference to the given string and assigns it to the LegBetStatus field.
func (o *TeaserLeg) SetLegBetStatus(v string) {
	o.LegBetStatus = &v
}

// GetLeagueId returns the LeagueId field value if set, zero value otherwise.
func (o *TeaserLeg) GetLeagueId() int {
	if o == nil || o.LeagueId == nil {
		var ret int
		return ret
	}
	return *o.LeagueId
}

// GetLeagueIdOk returns a tuple with the LeagueId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeaserLeg) GetLeagueIdOk() (*int, bool) {
	if o == nil || o.LeagueId == nil {
		return nil, false
	}
	return o.LeagueId, true
}

// HasLeagueId returns a boolean if a field has been set.
func (o *TeaserLeg) HasLeagueId() bool {
	if o != nil && o.LeagueId != nil {
		return true
	}

	return false
}

// SetLeagueId gets a reference to the given int and assigns it to the LeagueId field.
func (o *TeaserLeg) SetLeagueId(v int) {
	o.LeagueId = &v
}

// GetEventId returns the EventId field value if set, zero value otherwise.
func (o *TeaserLeg) GetEventId() int64 {
	if o == nil || o.EventId == nil {
		var ret int64
		return ret
	}
	return *o.EventId
}

// GetEventIdOk returns a tuple with the EventId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeaserLeg) GetEventIdOk() (*int64, bool) {
	if o == nil || o.EventId == nil {
		return nil, false
	}
	return o.EventId, true
}

// HasEventId returns a boolean if a field has been set.
func (o *TeaserLeg) HasEventId() bool {
	if o != nil && o.EventId != nil {
		return true
	}

	return false
}

// SetEventId gets a reference to the given int64 and assigns it to the EventId field.
func (o *TeaserLeg) SetEventId(v int64) {
	o.EventId = &v
}

// GetEventStartTime returns the EventStartTime field value if set, zero value otherwise.
func (o *TeaserLeg) GetEventStartTime() time.Time {
	if o == nil || o.EventStartTime == nil {
		var ret time.Time
		return ret
	}
	return *o.EventStartTime
}

// GetEventStartTimeOk returns a tuple with the EventStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeaserLeg) GetEventStartTimeOk() (*time.Time, bool) {
	if o == nil || o.EventStartTime == nil {
		return nil, false
	}
	return o.EventStartTime, true
}

// HasEventStartTime returns a boolean if a field has been set.
func (o *TeaserLeg) HasEventStartTime() bool {
	if o != nil && o.EventStartTime != nil {
		return true
	}

	return false
}

// SetEventStartTime gets a reference to the given time.Time and assigns it to the EventStartTime field.
func (o *TeaserLeg) SetEventStartTime(v time.Time) {
	o.EventStartTime = &v
}

// GetHandicap returns the Handicap field value if set, zero value otherwise.
func (o *TeaserLeg) GetHandicap() float64 {
	if o == nil || o.Handicap == nil {
		var ret float64
		return ret
	}
	return *o.Handicap
}

// GetHandicapOk returns a tuple with the Handicap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeaserLeg) GetHandicapOk() (*float64, bool) {
	if o == nil || o.Handicap == nil {
		return nil, false
	}
	return o.Handicap, true
}

// HasHandicap returns a boolean if a field has been set.
func (o *TeaserLeg) HasHandicap() bool {
	if o != nil && o.Handicap != nil {
		return true
	}

	return false
}

// SetHandicap gets a reference to the given float64 and assigns it to the Handicap field.
func (o *TeaserLeg) SetHandicap(v float64) {
	o.Handicap = &v
}

// GetTeamName returns the TeamName field value if set, zero value otherwise.
func (o *TeaserLeg) GetTeamName() string {
	if o == nil || o.TeamName == nil {
		var ret string
		return ret
	}
	return *o.TeamName
}

// GetTeamNameOk returns a tuple with the TeamName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeaserLeg) GetTeamNameOk() (*string, bool) {
	if o == nil || o.TeamName == nil {
		return nil, false
	}
	return o.TeamName, true
}

// HasTeamName returns a boolean if a field has been set.
func (o *TeaserLeg) HasTeamName() bool {
	if o != nil && o.TeamName != nil {
		return true
	}

	return false
}

// SetTeamName gets a reference to the given string and assigns it to the TeamName field.
func (o *TeaserLeg) SetTeamName(v string) {
	o.TeamName = &v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *TeaserLeg) GetSide() string {
	if o == nil || o.Side == nil {
		var ret string
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeaserLeg) GetSideOk() (*string, bool) {
	if o == nil || o.Side == nil {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *TeaserLeg) HasSide() bool {
	if o != nil && o.Side != nil {
		return true
	}

	return false
}

// SetSide gets a reference to the given string and assigns it to the Side field.
func (o *TeaserLeg) SetSide(v string) {
	o.Side = &v
}

// GetTeam1 returns the Team1 field value if set, zero value otherwise.
func (o *TeaserLeg) GetTeam1() string {
	if o == nil || o.Team1 == nil {
		var ret string
		return ret
	}
	return *o.Team1
}

// GetTeam1Ok returns a tuple with the Team1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeaserLeg) GetTeam1Ok() (*string, bool) {
	if o == nil || o.Team1 == nil {
		return nil, false
	}
	return o.Team1, true
}

// HasTeam1 returns a boolean if a field has been set.
func (o *TeaserLeg) HasTeam1() bool {
	if o != nil && o.Team1 != nil {
		return true
	}

	return false
}

// SetTeam1 gets a reference to the given string and assigns it to the Team1 field.
func (o *TeaserLeg) SetTeam1(v string) {
	o.Team1 = &v
}

// GetTeam2 returns the Team2 field value if set, zero value otherwise.
func (o *TeaserLeg) GetTeam2() string {
	if o == nil || o.Team2 == nil {
		var ret string
		return ret
	}
	return *o.Team2
}

// GetTeam2Ok returns a tuple with the Team2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeaserLeg) GetTeam2Ok() (*string, bool) {
	if o == nil || o.Team2 == nil {
		return nil, false
	}
	return o.Team2, true
}

// HasTeam2 returns a boolean if a field has been set.
func (o *TeaserLeg) HasTeam2() bool {
	if o != nil && o.Team2 != nil {
		return true
	}

	return false
}

// SetTeam2 gets a reference to the given string and assigns it to the Team2 field.
func (o *TeaserLeg) SetTeam2(v string) {
	o.Team2 = &v
}

// GetPeriodNumber returns the PeriodNumber field value if set, zero value otherwise.
func (o *TeaserLeg) GetPeriodNumber() int {
	if o == nil || o.PeriodNumber == nil {
		var ret int
		return ret
	}
	return *o.PeriodNumber
}

// GetPeriodNumberOk returns a tuple with the PeriodNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeaserLeg) GetPeriodNumberOk() (*int, bool) {
	if o == nil || o.PeriodNumber == nil {
		return nil, false
	}
	return o.PeriodNumber, true
}

// HasPeriodNumber returns a boolean if a field has been set.
func (o *TeaserLeg) HasPeriodNumber() bool {
	if o != nil && o.PeriodNumber != nil {
		return true
	}

	return false
}

// SetPeriodNumber gets a reference to the given int and assigns it to the PeriodNumber field.
func (o *TeaserLeg) SetPeriodNumber(v int) {
	o.PeriodNumber = &v
}

func (o TeaserLeg) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SportId != nil {
		toSerialize["sportId"] = o.SportId
	}
	if o.LegBetType != nil {
		toSerialize["legBetType"] = o.LegBetType
	}
	if o.LegBetStatus != nil {
		toSerialize["legBetStatus"] = o.LegBetStatus
	}
	if o.LeagueId != nil {
		toSerialize["leagueId"] = o.LeagueId
	}
	if o.EventId != nil {
		toSerialize["eventId"] = o.EventId
	}
	if o.EventStartTime != nil {
		toSerialize["eventStartTime"] = o.EventStartTime
	}
	if o.Handicap != nil {
		toSerialize["handicap"] = o.Handicap
	}
	if o.TeamName != nil {
		toSerialize["teamName"] = o.TeamName
	}
	if o.Side != nil {
		toSerialize["side"] = o.Side
	}
	if o.Team1 != nil {
		toSerialize["team1"] = o.Team1
	}
	if o.Team2 != nil {
		toSerialize["team2"] = o.Team2
	}
	if o.PeriodNumber != nil {
		toSerialize["periodNumber"] = o.PeriodNumber
	}
	return json.Marshal(toSerialize)
}

type NullableTeaserLeg struct {
	value *TeaserLeg
	isSet bool
}

func (v NullableTeaserLeg) Get() *TeaserLeg {
	return v.value
}

func (v *NullableTeaserLeg) Set(val *TeaserLeg) {
	v.value = val
	v.isSet = true
}

func (v NullableTeaserLeg) IsSet() bool {
	return v.isSet
}

func (v *NullableTeaserLeg) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTeaserLeg(val *TeaserLeg) *NullableTeaserLeg {
	return &NullableTeaserLeg{value: val, isSet: true}
}

func (v NullableTeaserLeg) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTeaserLeg) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}