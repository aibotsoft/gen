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
	"time"
)

// ParlayLeg struct for ParlayLeg
type ParlayLeg struct {
	SportId *int32 `json:"sportId,omitempty" xml:"sportId"`
	// Parlay leg type.
	LegBetType *string `json:"legBetType,omitempty" xml:"legBetType"`
	// Parlay Leg status. CANCELLED = The leg is canceled- the stake on this leg will be transferred to the next one. In this case the leg will be ignored when calculating the winLoss,   LOSE = The leg is a loss or a push-lose. When Push-lose happens, the half of the stake on the leg will be pushed to the next leg, and the other half will be a lose. This can happen only when the leg is placed on a quarter points handicap,   PUSH = The leg is a push - the stake on this leg will be transferred to the next one. In this case the leg will be ignored when calculating the winLoss,   REFUNDED = The leg is refunded - the stake on this leg will be transferred to the next one. In this case the leg will be ignored when calculating the winLoss,   WON = The leg is a won or a push-won. When Push-won happens, the half of the stake on the leg will be pushed to the next leg, and the other half is won. This can happen only when the leg is placed on a quarter points handicap
	LegBetStatus *string `json:"legBetStatus,omitempty" xml:"legBetStatus"`
	LeagueId     *int32  `json:"leagueId,omitempty" xml:"leagueId"`
	EventId      *int64  `json:"eventId,omitempty" xml:"eventId"`
	// Date time when the event starts.
	EventStartTime *time.Time       `json:"eventStartTime,omitempty" xml:"eventStartTime"`
	Handicap       *NullableFloat64 `json:"handicap,omitempty" xml:"handicap"`
	Price          *float64         `json:"price,omitempty" xml:"price"`
	TeamName       *string          `json:"teamName,omitempty" xml:"teamName"`
	// Side type.
	Side              *NullableString `json:"side,omitempty" xml:"side"`
	Pitcher1          *NullableString `json:"pitcher1,omitempty" xml:"pitcher1"`
	Pitcher2          *NullableString `json:"pitcher2,omitempty" xml:"pitcher2"`
	Pitcher1MustStart *bool           `json:"pitcher1MustStart,omitempty" xml:"pitcher1MustStart"`
	Pitcher2MustStart *bool           `json:"pitcher2MustStart,omitempty" xml:"pitcher2MustStart"`
	// Wellington Phoenix
	Team1 *string `json:"team1,omitempty" xml:"team1"`
	// Adelaide United
	Team2        *string `json:"team2,omitempty" xml:"team2"`
	PeriodNumber *int32  `json:"periodNumber,omitempty" xml:"periodNumber"`
	// Full time team 1 score
	FtTeam1Score *NullableFloat64 `json:"ftTeam1Score,omitempty" xml:"ftTeam1Score"`
	// Full time team 2 score
	FtTeam2Score *NullableFloat64 `json:"ftTeam2Score,omitempty" xml:"ftTeam2Score"`
	// End of period team 1 score. If the bet was placed on Game period (periodNumber =0) , this will be null
	PTeam1Score *NullableFloat64 `json:"pTeam1Score,omitempty" xml:"pTeam1Score"`
	// End of period team 2 score. If the bet was placed on Game period (periodNumber =0) , this will be null
	PTeam2Score        *NullableFloat64    `json:"pTeam2Score,omitempty" xml:"pTeam2Score"`
	CancellationReason *CancellationReason `json:"cancellationReason,omitempty" xml:"cancellationReason"`
}

// GetSportId returns the SportId field value if set, zero value otherwise.
func (o *ParlayLeg) GetSportId() int32 {
	if o == nil || o.SportId == nil {
		var ret int32
		return ret
	}
	return *o.SportId
}

// GetSportIdOk returns a tuple with the SportId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ParlayLeg) GetSportIdOk() (int32, bool) {
	if o == nil || o.SportId == nil {
		var ret int32
		return ret, false
	}
	return *o.SportId, true
}

// HasSportId returns a boolean if a field has been set.
func (o *ParlayLeg) HasSportId() bool {
	if o != nil && o.SportId != nil {
		return true
	}

	return false
}

// SetSportId gets a reference to the given int32 and assigns it to the SportId field.
func (o *ParlayLeg) SetSportId(v int32) {
	o.SportId = &v
}

// GetLegBetType returns the LegBetType field value if set, zero value otherwise.
func (o *ParlayLeg) GetLegBetType() string {
	if o == nil || o.LegBetType == nil {
		var ret string
		return ret
	}
	return *o.LegBetType
}

// GetLegBetTypeOk returns a tuple with the LegBetType field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ParlayLeg) GetLegBetTypeOk() (string, bool) {
	if o == nil || o.LegBetType == nil {
		var ret string
		return ret, false
	}
	return *o.LegBetType, true
}

// HasLegBetType returns a boolean if a field has been set.
func (o *ParlayLeg) HasLegBetType() bool {
	if o != nil && o.LegBetType != nil {
		return true
	}

	return false
}

// SetLegBetType gets a reference to the given string and assigns it to the LegBetType field.
func (o *ParlayLeg) SetLegBetType(v string) {
	o.LegBetType = &v
}

// GetLegBetStatus returns the LegBetStatus field value if set, zero value otherwise.
func (o *ParlayLeg) GetLegBetStatus() string {
	if o == nil || o.LegBetStatus == nil {
		var ret string
		return ret
	}
	return *o.LegBetStatus
}

// GetLegBetStatusOk returns a tuple with the LegBetStatus field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ParlayLeg) GetLegBetStatusOk() (string, bool) {
	if o == nil || o.LegBetStatus == nil {
		var ret string
		return ret, false
	}
	return *o.LegBetStatus, true
}

// HasLegBetStatus returns a boolean if a field has been set.
func (o *ParlayLeg) HasLegBetStatus() bool {
	if o != nil && o.LegBetStatus != nil {
		return true
	}

	return false
}

// SetLegBetStatus gets a reference to the given string and assigns it to the LegBetStatus field.
func (o *ParlayLeg) SetLegBetStatus(v string) {
	o.LegBetStatus = &v
}

// GetLeagueId returns the LeagueId field value if set, zero value otherwise.
func (o *ParlayLeg) GetLeagueId() int32 {
	if o == nil || o.LeagueId == nil {
		var ret int32
		return ret
	}
	return *o.LeagueId
}

// GetLeagueIdOk returns a tuple with the LeagueId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ParlayLeg) GetLeagueIdOk() (int32, bool) {
	if o == nil || o.LeagueId == nil {
		var ret int32
		return ret, false
	}
	return *o.LeagueId, true
}

// HasLeagueId returns a boolean if a field has been set.
func (o *ParlayLeg) HasLeagueId() bool {
	if o != nil && o.LeagueId != nil {
		return true
	}

	return false
}

// SetLeagueId gets a reference to the given int32 and assigns it to the LeagueId field.
func (o *ParlayLeg) SetLeagueId(v int32) {
	o.LeagueId = &v
}

// GetEventId returns the EventId field value if set, zero value otherwise.
func (o *ParlayLeg) GetEventId() int64 {
	if o == nil || o.EventId == nil {
		var ret int64
		return ret
	}
	return *o.EventId
}

// GetEventIdOk returns a tuple with the EventId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ParlayLeg) GetEventIdOk() (int64, bool) {
	if o == nil || o.EventId == nil {
		var ret int64
		return ret, false
	}
	return *o.EventId, true
}

// HasEventId returns a boolean if a field has been set.
func (o *ParlayLeg) HasEventId() bool {
	if o != nil && o.EventId != nil {
		return true
	}

	return false
}

// SetEventId gets a reference to the given int64 and assigns it to the EventId field.
func (o *ParlayLeg) SetEventId(v int64) {
	o.EventId = &v
}

// GetEventStartTime returns the EventStartTime field value if set, zero value otherwise.
func (o *ParlayLeg) GetEventStartTime() time.Time {
	if o == nil || o.EventStartTime == nil {
		var ret time.Time
		return ret
	}
	return *o.EventStartTime
}

// GetEventStartTimeOk returns a tuple with the EventStartTime field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ParlayLeg) GetEventStartTimeOk() (time.Time, bool) {
	if o == nil || o.EventStartTime == nil {
		var ret time.Time
		return ret, false
	}
	return *o.EventStartTime, true
}

// HasEventStartTime returns a boolean if a field has been set.
func (o *ParlayLeg) HasEventStartTime() bool {
	if o != nil && o.EventStartTime != nil {
		return true
	}

	return false
}

// SetEventStartTime gets a reference to the given time.Time and assigns it to the EventStartTime field.
func (o *ParlayLeg) SetEventStartTime(v time.Time) {
	o.EventStartTime = &v
}

// GetHandicap returns the Handicap field value if set, zero value otherwise.
func (o *ParlayLeg) GetHandicap() NullableFloat64 {
	if o == nil || o.Handicap == nil {
		var ret NullableFloat64
		return ret
	}
	return *o.Handicap
}

// GetHandicapOk returns a tuple with the Handicap field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ParlayLeg) GetHandicapOk() (NullableFloat64, bool) {
	if o == nil || o.Handicap == nil {
		var ret NullableFloat64
		return ret, false
	}
	return *o.Handicap, true
}

// HasHandicap returns a boolean if a field has been set.
func (o *ParlayLeg) HasHandicap() bool {
	if o != nil && o.Handicap != nil {
		return true
	}

	return false
}

// SetHandicap gets a reference to the given NullableFloat64 and assigns it to the Handicap field.
func (o *ParlayLeg) SetHandicap(v NullableFloat64) {
	o.Handicap = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *ParlayLeg) GetPrice() float64 {
	if o == nil || o.Price == nil {
		var ret float64
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ParlayLeg) GetPriceOk() (float64, bool) {
	if o == nil || o.Price == nil {
		var ret float64
		return ret, false
	}
	return *o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *ParlayLeg) HasPrice() bool {
	if o != nil && o.Price != nil {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float64 and assigns it to the Price field.
func (o *ParlayLeg) SetPrice(v float64) {
	o.Price = &v
}

// GetTeamName returns the TeamName field value if set, zero value otherwise.
func (o *ParlayLeg) GetTeamName() string {
	if o == nil || o.TeamName == nil {
		var ret string
		return ret
	}
	return *o.TeamName
}

// GetTeamNameOk returns a tuple with the TeamName field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ParlayLeg) GetTeamNameOk() (string, bool) {
	if o == nil || o.TeamName == nil {
		var ret string
		return ret, false
	}
	return *o.TeamName, true
}

// HasTeamName returns a boolean if a field has been set.
func (o *ParlayLeg) HasTeamName() bool {
	if o != nil && o.TeamName != nil {
		return true
	}

	return false
}

// SetTeamName gets a reference to the given string and assigns it to the TeamName field.
func (o *ParlayLeg) SetTeamName(v string) {
	o.TeamName = &v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *ParlayLeg) GetSide() NullableString {
	if o == nil || o.Side == nil {
		var ret NullableString
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ParlayLeg) GetSideOk() (NullableString, bool) {
	if o == nil || o.Side == nil {
		var ret NullableString
		return ret, false
	}
	return *o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *ParlayLeg) HasSide() bool {
	if o != nil && o.Side != nil {
		return true
	}

	return false
}

// SetSide gets a reference to the given NullableString and assigns it to the Side field.
func (o *ParlayLeg) SetSide(v NullableString) {
	o.Side = &v
}

// GetPitcher1 returns the Pitcher1 field value if set, zero value otherwise.
func (o *ParlayLeg) GetPitcher1() NullableString {
	if o == nil || o.Pitcher1 == nil {
		var ret NullableString
		return ret
	}
	return *o.Pitcher1
}

// GetPitcher1Ok returns a tuple with the Pitcher1 field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ParlayLeg) GetPitcher1Ok() (NullableString, bool) {
	if o == nil || o.Pitcher1 == nil {
		var ret NullableString
		return ret, false
	}
	return *o.Pitcher1, true
}

// HasPitcher1 returns a boolean if a field has been set.
func (o *ParlayLeg) HasPitcher1() bool {
	if o != nil && o.Pitcher1 != nil {
		return true
	}

	return false
}

// SetPitcher1 gets a reference to the given NullableString and assigns it to the Pitcher1 field.
func (o *ParlayLeg) SetPitcher1(v NullableString) {
	o.Pitcher1 = &v
}

// GetPitcher2 returns the Pitcher2 field value if set, zero value otherwise.
func (o *ParlayLeg) GetPitcher2() NullableString {
	if o == nil || o.Pitcher2 == nil {
		var ret NullableString
		return ret
	}
	return *o.Pitcher2
}

// GetPitcher2Ok returns a tuple with the Pitcher2 field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ParlayLeg) GetPitcher2Ok() (NullableString, bool) {
	if o == nil || o.Pitcher2 == nil {
		var ret NullableString
		return ret, false
	}
	return *o.Pitcher2, true
}

// HasPitcher2 returns a boolean if a field has been set.
func (o *ParlayLeg) HasPitcher2() bool {
	if o != nil && o.Pitcher2 != nil {
		return true
	}

	return false
}

// SetPitcher2 gets a reference to the given NullableString and assigns it to the Pitcher2 field.
func (o *ParlayLeg) SetPitcher2(v NullableString) {
	o.Pitcher2 = &v
}

// GetPitcher1MustStart returns the Pitcher1MustStart field value if set, zero value otherwise.
func (o *ParlayLeg) GetPitcher1MustStart() bool {
	if o == nil || o.Pitcher1MustStart == nil {
		var ret bool
		return ret
	}
	return *o.Pitcher1MustStart
}

// GetPitcher1MustStartOk returns a tuple with the Pitcher1MustStart field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ParlayLeg) GetPitcher1MustStartOk() (bool, bool) {
	if o == nil || o.Pitcher1MustStart == nil {
		var ret bool
		return ret, false
	}
	return *o.Pitcher1MustStart, true
}

// HasPitcher1MustStart returns a boolean if a field has been set.
func (o *ParlayLeg) HasPitcher1MustStart() bool {
	if o != nil && o.Pitcher1MustStart != nil {
		return true
	}

	return false
}

// SetPitcher1MustStart gets a reference to the given bool and assigns it to the Pitcher1MustStart field.
func (o *ParlayLeg) SetPitcher1MustStart(v bool) {
	o.Pitcher1MustStart = &v
}

// GetPitcher2MustStart returns the Pitcher2MustStart field value if set, zero value otherwise.
func (o *ParlayLeg) GetPitcher2MustStart() bool {
	if o == nil || o.Pitcher2MustStart == nil {
		var ret bool
		return ret
	}
	return *o.Pitcher2MustStart
}

// GetPitcher2MustStartOk returns a tuple with the Pitcher2MustStart field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ParlayLeg) GetPitcher2MustStartOk() (bool, bool) {
	if o == nil || o.Pitcher2MustStart == nil {
		var ret bool
		return ret, false
	}
	return *o.Pitcher2MustStart, true
}

// HasPitcher2MustStart returns a boolean if a field has been set.
func (o *ParlayLeg) HasPitcher2MustStart() bool {
	if o != nil && o.Pitcher2MustStart != nil {
		return true
	}

	return false
}

// SetPitcher2MustStart gets a reference to the given bool and assigns it to the Pitcher2MustStart field.
func (o *ParlayLeg) SetPitcher2MustStart(v bool) {
	o.Pitcher2MustStart = &v
}

// GetTeam1 returns the Team1 field value if set, zero value otherwise.
func (o *ParlayLeg) GetTeam1() string {
	if o == nil || o.Team1 == nil {
		var ret string
		return ret
	}
	return *o.Team1
}

// GetTeam1Ok returns a tuple with the Team1 field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ParlayLeg) GetTeam1Ok() (string, bool) {
	if o == nil || o.Team1 == nil {
		var ret string
		return ret, false
	}
	return *o.Team1, true
}

// HasTeam1 returns a boolean if a field has been set.
func (o *ParlayLeg) HasTeam1() bool {
	if o != nil && o.Team1 != nil {
		return true
	}

	return false
}

// SetTeam1 gets a reference to the given string and assigns it to the Team1 field.
func (o *ParlayLeg) SetTeam1(v string) {
	o.Team1 = &v
}

// GetTeam2 returns the Team2 field value if set, zero value otherwise.
func (o *ParlayLeg) GetTeam2() string {
	if o == nil || o.Team2 == nil {
		var ret string
		return ret
	}
	return *o.Team2
}

// GetTeam2Ok returns a tuple with the Team2 field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ParlayLeg) GetTeam2Ok() (string, bool) {
	if o == nil || o.Team2 == nil {
		var ret string
		return ret, false
	}
	return *o.Team2, true
}

// HasTeam2 returns a boolean if a field has been set.
func (o *ParlayLeg) HasTeam2() bool {
	if o != nil && o.Team2 != nil {
		return true
	}

	return false
}

// SetTeam2 gets a reference to the given string and assigns it to the Team2 field.
func (o *ParlayLeg) SetTeam2(v string) {
	o.Team2 = &v
}

// GetPeriodNumber returns the PeriodNumber field value if set, zero value otherwise.
func (o *ParlayLeg) GetPeriodNumber() int32 {
	if o == nil || o.PeriodNumber == nil {
		var ret int32
		return ret
	}
	return *o.PeriodNumber
}

// GetPeriodNumberOk returns a tuple with the PeriodNumber field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ParlayLeg) GetPeriodNumberOk() (int32, bool) {
	if o == nil || o.PeriodNumber == nil {
		var ret int32
		return ret, false
	}
	return *o.PeriodNumber, true
}

// HasPeriodNumber returns a boolean if a field has been set.
func (o *ParlayLeg) HasPeriodNumber() bool {
	if o != nil && o.PeriodNumber != nil {
		return true
	}

	return false
}

// SetPeriodNumber gets a reference to the given int32 and assigns it to the PeriodNumber field.
func (o *ParlayLeg) SetPeriodNumber(v int32) {
	o.PeriodNumber = &v
}

// GetFtTeam1Score returns the FtTeam1Score field value if set, zero value otherwise.
func (o *ParlayLeg) GetFtTeam1Score() NullableFloat64 {
	if o == nil || o.FtTeam1Score == nil {
		var ret NullableFloat64
		return ret
	}
	return *o.FtTeam1Score
}

// GetFtTeam1ScoreOk returns a tuple with the FtTeam1Score field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ParlayLeg) GetFtTeam1ScoreOk() (NullableFloat64, bool) {
	if o == nil || o.FtTeam1Score == nil {
		var ret NullableFloat64
		return ret, false
	}
	return *o.FtTeam1Score, true
}

// HasFtTeam1Score returns a boolean if a field has been set.
func (o *ParlayLeg) HasFtTeam1Score() bool {
	if o != nil && o.FtTeam1Score != nil {
		return true
	}

	return false
}

// SetFtTeam1Score gets a reference to the given NullableFloat64 and assigns it to the FtTeam1Score field.
func (o *ParlayLeg) SetFtTeam1Score(v NullableFloat64) {
	o.FtTeam1Score = &v
}

// GetFtTeam2Score returns the FtTeam2Score field value if set, zero value otherwise.
func (o *ParlayLeg) GetFtTeam2Score() NullableFloat64 {
	if o == nil || o.FtTeam2Score == nil {
		var ret NullableFloat64
		return ret
	}
	return *o.FtTeam2Score
}

// GetFtTeam2ScoreOk returns a tuple with the FtTeam2Score field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ParlayLeg) GetFtTeam2ScoreOk() (NullableFloat64, bool) {
	if o == nil || o.FtTeam2Score == nil {
		var ret NullableFloat64
		return ret, false
	}
	return *o.FtTeam2Score, true
}

// HasFtTeam2Score returns a boolean if a field has been set.
func (o *ParlayLeg) HasFtTeam2Score() bool {
	if o != nil && o.FtTeam2Score != nil {
		return true
	}

	return false
}

// SetFtTeam2Score gets a reference to the given NullableFloat64 and assigns it to the FtTeam2Score field.
func (o *ParlayLeg) SetFtTeam2Score(v NullableFloat64) {
	o.FtTeam2Score = &v
}

// GetPTeam1Score returns the PTeam1Score field value if set, zero value otherwise.
func (o *ParlayLeg) GetPTeam1Score() NullableFloat64 {
	if o == nil || o.PTeam1Score == nil {
		var ret NullableFloat64
		return ret
	}
	return *o.PTeam1Score
}

// GetPTeam1ScoreOk returns a tuple with the PTeam1Score field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ParlayLeg) GetPTeam1ScoreOk() (NullableFloat64, bool) {
	if o == nil || o.PTeam1Score == nil {
		var ret NullableFloat64
		return ret, false
	}
	return *o.PTeam1Score, true
}

// HasPTeam1Score returns a boolean if a field has been set.
func (o *ParlayLeg) HasPTeam1Score() bool {
	if o != nil && o.PTeam1Score != nil {
		return true
	}

	return false
}

// SetPTeam1Score gets a reference to the given NullableFloat64 and assigns it to the PTeam1Score field.
func (o *ParlayLeg) SetPTeam1Score(v NullableFloat64) {
	o.PTeam1Score = &v
}

// GetPTeam2Score returns the PTeam2Score field value if set, zero value otherwise.
func (o *ParlayLeg) GetPTeam2Score() NullableFloat64 {
	if o == nil || o.PTeam2Score == nil {
		var ret NullableFloat64
		return ret
	}
	return *o.PTeam2Score
}

// GetPTeam2ScoreOk returns a tuple with the PTeam2Score field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ParlayLeg) GetPTeam2ScoreOk() (NullableFloat64, bool) {
	if o == nil || o.PTeam2Score == nil {
		var ret NullableFloat64
		return ret, false
	}
	return *o.PTeam2Score, true
}

// HasPTeam2Score returns a boolean if a field has been set.
func (o *ParlayLeg) HasPTeam2Score() bool {
	if o != nil && o.PTeam2Score != nil {
		return true
	}

	return false
}

// SetPTeam2Score gets a reference to the given NullableFloat64 and assigns it to the PTeam2Score field.
func (o *ParlayLeg) SetPTeam2Score(v NullableFloat64) {
	o.PTeam2Score = &v
}

// GetCancellationReason returns the CancellationReason field value if set, zero value otherwise.
func (o *ParlayLeg) GetCancellationReason() CancellationReason {
	if o == nil || o.CancellationReason == nil {
		var ret CancellationReason
		return ret
	}
	return *o.CancellationReason
}

// GetCancellationReasonOk returns a tuple with the CancellationReason field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ParlayLeg) GetCancellationReasonOk() (CancellationReason, bool) {
	if o == nil || o.CancellationReason == nil {
		var ret CancellationReason
		return ret, false
	}
	return *o.CancellationReason, true
}

// HasCancellationReason returns a boolean if a field has been set.
func (o *ParlayLeg) HasCancellationReason() bool {
	if o != nil && o.CancellationReason != nil {
		return true
	}

	return false
}

// SetCancellationReason gets a reference to the given CancellationReason and assigns it to the CancellationReason field.
func (o *ParlayLeg) SetCancellationReason(v CancellationReason) {
	o.CancellationReason = &v
}

type NullableParlayLeg struct {
	Value        ParlayLeg
	ExplicitNull bool
}

func (v NullableParlayLeg) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableParlayLeg) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
