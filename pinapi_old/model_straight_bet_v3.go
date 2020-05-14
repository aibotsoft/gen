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

// StraightBetV3 struct for StraightBetV3
type StraightBetV3 struct {
	// Bet identification
	BetId int64 `json:"betId" xml:"betId"`
	// Wager identification. All bets placed thru the API will have value 1. Website Classic view supports multiple contest(special) bets placement in the same bet slip in that case the bet would have appropriate wager number, as well as all round robin parlay bets.
	WagerNumber int32 `json:"wagerNumber" xml:"wagerNumber"`
	// Date time when the bet was placed.
	PlacedAt time.Time `json:"placedAt" xml:"placedAt"`
	// Bet Status.    ACCEPTED = Bet was accepted,   CANCELLED = Bet is cancelled as per Pinnacle betting rules,   LOSE = The bet is settled as lose,   PENDING_ACCEPTANCE = This status is reserved only for live bets. If a live bet is placed during danger zone or live delay is applied, it will be in PENDING_ACCEPTANCE , otherwise in ACCEPTED status. From this status bet can go to ACCEPTED or NOT_ACCEPTED status,   REFUNDED = When an event is cancelled or when the bet is settled as push, the bet will have REFUNDED status,   NOT_ACCEPTED = Bet was not accepted. Bet can be in this status only if it was previously in PENDING_ACCEPTANCE status,   WON = The bet is settled as won
	BetStatus string `json:"betStatus" xml:"betStatus"`
	// Bet type.
	BetType string `json:"betType" xml:"betType"`
	// Win amount.
	Win float64 `json:"win" xml:"win"`
	// Risk amount.
	Risk float64 `json:"risk" xml:"risk"`
	// Win-Loss for settled bets.
	WinLoss    *NullableFloat64 `json:"winLoss,omitempty" xml:"winLoss"`
	OddsFormat OddsFormat       `json:"oddsFormat" xml:"oddsFormat"`
	// Client’s commission on the bet.
	CustomerCommission *NullableFloat64    `json:"customerCommission,omitempty" xml:"customerCommission"`
	CancellationReason *CancellationReason `json:"cancellationReason,omitempty" xml:"cancellationReason"`
	// Update Sequence
	UpdateSequence int64            `json:"updateSequence" xml:"updateSequence"`
	SportId        *int32           `json:"sportId,omitempty" xml:"sportId"`
	LeagueId       *int32           `json:"leagueId,omitempty" xml:"leagueId"`
	EventId        *int64           `json:"eventId,omitempty" xml:"eventId"`
	Handicap       *NullableFloat64 `json:"handicap,omitempty" xml:"handicap"`
	Price          *float64         `json:"price,omitempty" xml:"price"`
	TeamName       *string          `json:"teamName,omitempty" xml:"teamName"`
	// Side type.
	Side *NullableString `json:"side,omitempty" xml:"side"`
	// Pitcher name of team1. Only for bets on baseball.
	Pitcher1 *NullableString `json:"pitcher1,omitempty" xml:"pitcher1"`
	// Pitcher name of team2. Only for bets on baseball.
	Pitcher2 *NullableString `json:"pitcher2,omitempty" xml:"pitcher2"`
	// Baseball only. Refers to the pitcher for Team1.  This applicable only for MONEYLINE bet type, for all other bet types this has to be TRUE.
	Pitcher1MustStart *NullableBool `json:"pitcher1MustStart,omitempty" xml:"pitcher1MustStart"`
	// Baseball only. Refers to the pitcher for Team2.  This applicable only for MONEYLINE bet type, for all other bet types this has to be TRUE.
	Pitcher2MustStart *NullableBool `json:"pitcher2MustStart,omitempty" xml:"pitcher2MustStart"`
	Team1             *string       `json:"team1,omitempty" xml:"team1"`
	Team2             *string       `json:"team2,omitempty" xml:"team2"`
	PeriodNumber      *int32        `json:"periodNumber,omitempty" xml:"periodNumber"`
	// Team 1 score that the bet was placed on, only for live bets.
	Team1Score *NullableFloat64 `json:"team1Score,omitempty" xml:"team1Score"`
	// Team 2 score that the bet was placed, only for live bets.
	Team2Score *NullableFloat64 `json:"team2Score,omitempty" xml:"team2Score"`
	// Full time team 1 score, only for settled bets.
	FtTeam1Score *NullableFloat64 `json:"ftTeam1Score,omitempty" xml:"ftTeam1Score"`
	// Full time team 2 score, only for settled bets.
	FtTeam2Score *NullableFloat64 `json:"ftTeam2Score,omitempty" xml:"ftTeam2Score"`
	// .End of period team 1 score, only for settled bets. If the bet was placed on Game period (periodNumber =0), this will be null .
	PTeam1Score *NullableFloat64 `json:"pTeam1Score,omitempty" xml:"pTeam1Score"`
	// End of period team 2 score, only for settled bets. If the bet was placed on Game period (periodNumber =0), this will be null
	PTeam2Score *NullableFloat64 `json:"pTeam2Score,omitempty" xml:"pTeam2Score"`
	// Whether the bet is on live event
	IsLive *bool `json:"isLive,omitempty" xml:"isLive"`
	// Date time when the event starts.
	EventStartTime *time.Time `json:"eventStartTime,omitempty" xml:"eventStartTime"`
}

// GetBetId returns the BetId field value
func (o *StraightBetV3) GetBetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.BetId
}

// SetBetId sets field value
func (o *StraightBetV3) SetBetId(v int64) {
	o.BetId = v
}

// GetWagerNumber returns the WagerNumber field value
func (o *StraightBetV3) GetWagerNumber() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.WagerNumber
}

// SetWagerNumber sets field value
func (o *StraightBetV3) SetWagerNumber(v int32) {
	o.WagerNumber = v
}

// GetPlacedAt returns the PlacedAt field value
func (o *StraightBetV3) GetPlacedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.PlacedAt
}

// SetPlacedAt sets field value
func (o *StraightBetV3) SetPlacedAt(v time.Time) {
	o.PlacedAt = v
}

// GetBetStatus returns the BetStatus field value
func (o *StraightBetV3) GetBetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BetStatus
}

// SetBetStatus sets field value
func (o *StraightBetV3) SetBetStatus(v string) {
	o.BetStatus = v
}

// GetBetType returns the BetTypeId field value
func (o *StraightBetV3) GetBetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BetType
}

// SetBetType sets field value
func (o *StraightBetV3) SetBetType(v string) {
	o.BetType = v
}

// GetWin returns the Win field value
func (o *StraightBetV3) GetWin() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Win
}

// SetWin sets field value
func (o *StraightBetV3) SetWin(v float64) {
	o.Win = v
}

// GetRisk returns the Risk field value
func (o *StraightBetV3) GetRisk() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Risk
}

// SetRisk sets field value
func (o *StraightBetV3) SetRisk(v float64) {
	o.Risk = v
}

// GetWinLoss returns the WinLoss field value if set, zero value otherwise.
func (o *StraightBetV3) GetWinLoss() NullableFloat64 {
	if o == nil || o.WinLoss == nil {
		var ret NullableFloat64
		return ret
	}
	return *o.WinLoss
}

// GetWinLossOk returns a tuple with the WinLoss field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *StraightBetV3) GetWinLossOk() (NullableFloat64, bool) {
	if o == nil || o.WinLoss == nil {
		var ret NullableFloat64
		return ret, false
	}
	return *o.WinLoss, true
}

// HasWinLoss returns a boolean if a field has been set.
func (o *StraightBetV3) HasWinLoss() bool {
	if o != nil && o.WinLoss != nil {
		return true
	}

	return false
}

// SetWinLoss gets a reference to the given NullableFloat64 and assigns it to the WinLoss field.
func (o *StraightBetV3) SetWinLoss(v NullableFloat64) {
	o.WinLoss = &v
}

// GetOddsFormat returns the OddsFormat field value
func (o *StraightBetV3) GetOddsFormat() OddsFormat {
	if o == nil {
		var ret OddsFormat
		return ret
	}

	return o.OddsFormat
}

// SetOddsFormat sets field value
func (o *StraightBetV3) SetOddsFormat(v OddsFormat) {
	o.OddsFormat = v
}

// GetCustomerCommission returns the CustomerCommission field value if set, zero value otherwise.
func (o *StraightBetV3) GetCustomerCommission() NullableFloat64 {
	if o == nil || o.CustomerCommission == nil {
		var ret NullableFloat64
		return ret
	}
	return *o.CustomerCommission
}

// GetCustomerCommissionOk returns a tuple with the CustomerCommission field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *StraightBetV3) GetCustomerCommissionOk() (NullableFloat64, bool) {
	if o == nil || o.CustomerCommission == nil {
		var ret NullableFloat64
		return ret, false
	}
	return *o.CustomerCommission, true
}

// HasCustomerCommission returns a boolean if a field has been set.
func (o *StraightBetV3) HasCustomerCommission() bool {
	if o != nil && o.CustomerCommission != nil {
		return true
	}

	return false
}

// SetCustomerCommission gets a reference to the given NullableFloat64 and assigns it to the CustomerCommission field.
func (o *StraightBetV3) SetCustomerCommission(v NullableFloat64) {
	o.CustomerCommission = &v
}

// GetCancellationReason returns the CancellationReason field value if set, zero value otherwise.
func (o *StraightBetV3) GetCancellationReason() CancellationReason {
	if o == nil || o.CancellationReason == nil {
		var ret CancellationReason
		return ret
	}
	return *o.CancellationReason
}

// GetCancellationReasonOk returns a tuple with the CancellationReason field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *StraightBetV3) GetCancellationReasonOk() (CancellationReason, bool) {
	if o == nil || o.CancellationReason == nil {
		var ret CancellationReason
		return ret, false
	}
	return *o.CancellationReason, true
}

// HasCancellationReason returns a boolean if a field has been set.
func (o *StraightBetV3) HasCancellationReason() bool {
	if o != nil && o.CancellationReason != nil {
		return true
	}

	return false
}

// SetCancellationReason gets a reference to the given CancellationReason and assigns it to the CancellationReason field.
func (o *StraightBetV3) SetCancellationReason(v CancellationReason) {
	o.CancellationReason = &v
}

// GetUpdateSequence returns the UpdateSequence field value
func (o *StraightBetV3) GetUpdateSequence() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UpdateSequence
}

// SetUpdateSequence sets field value
func (o *StraightBetV3) SetUpdateSequence(v int64) {
	o.UpdateSequence = v
}

// GetSportId returns the SportId field value if set, zero value otherwise.
func (o *StraightBetV3) GetSportId() int32 {
	if o == nil || o.SportId == nil {
		var ret int32
		return ret
	}
	return *o.SportId
}

// GetSportIdOk returns a tuple with the SportId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *StraightBetV3) GetSportIdOk() (int32, bool) {
	if o == nil || o.SportId == nil {
		var ret int32
		return ret, false
	}
	return *o.SportId, true
}

// HasSportId returns a boolean if a field has been set.
func (o *StraightBetV3) HasSportId() bool {
	if o != nil && o.SportId != nil {
		return true
	}

	return false
}

// SetSportId gets a reference to the given int32 and assigns it to the SportId field.
func (o *StraightBetV3) SetSportId(v int32) {
	o.SportId = &v
}

// GetLeagueId returns the LeagueId field value if set, zero value otherwise.
func (o *StraightBetV3) GetLeagueId() int32 {
	if o == nil || o.LeagueId == nil {
		var ret int32
		return ret
	}
	return *o.LeagueId
}

// GetLeagueIdOk returns a tuple with the LeagueId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *StraightBetV3) GetLeagueIdOk() (int32, bool) {
	if o == nil || o.LeagueId == nil {
		var ret int32
		return ret, false
	}
	return *o.LeagueId, true
}

// HasLeagueId returns a boolean if a field has been set.
func (o *StraightBetV3) HasLeagueId() bool {
	if o != nil && o.LeagueId != nil {
		return true
	}

	return false
}

// SetLeagueId gets a reference to the given int32 and assigns it to the LeagueId field.
func (o *StraightBetV3) SetLeagueId(v int32) {
	o.LeagueId = &v
}

// GetEventId returns the EventId field value if set, zero value otherwise.
func (o *StraightBetV3) GetEventId() int64 {
	if o == nil || o.EventId == nil {
		var ret int64
		return ret
	}
	return *o.EventId
}

// GetEventIdOk returns a tuple with the EventId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *StraightBetV3) GetEventIdOk() (int64, bool) {
	if o == nil || o.EventId == nil {
		var ret int64
		return ret, false
	}
	return *o.EventId, true
}

// HasEventId returns a boolean if a field has been set.
func (o *StraightBetV3) HasEventId() bool {
	if o != nil && o.EventId != nil {
		return true
	}

	return false
}

// SetEventId gets a reference to the given int64 and assigns it to the EventId field.
func (o *StraightBetV3) SetEventId(v int64) {
	o.EventId = &v
}

// GetHandicap returns the Handicap field value if set, zero value otherwise.
func (o *StraightBetV3) GetHandicap() NullableFloat64 {
	if o == nil || o.Handicap == nil {
		var ret NullableFloat64
		return ret
	}
	return *o.Handicap
}

// GetHandicapOk returns a tuple with the Handicap field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *StraightBetV3) GetHandicapOk() (NullableFloat64, bool) {
	if o == nil || o.Handicap == nil {
		var ret NullableFloat64
		return ret, false
	}
	return *o.Handicap, true
}

// HasHandicap returns a boolean if a field has been set.
func (o *StraightBetV3) HasHandicap() bool {
	if o != nil && o.Handicap != nil {
		return true
	}

	return false
}

// SetHandicap gets a reference to the given NullableFloat64 and assigns it to the Handicap field.
func (o *StraightBetV3) SetHandicap(v NullableFloat64) {
	o.Handicap = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *StraightBetV3) GetPrice() float64 {
	if o == nil || o.Price == nil {
		var ret float64
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *StraightBetV3) GetPriceOk() (float64, bool) {
	if o == nil || o.Price == nil {
		var ret float64
		return ret, false
	}
	return *o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *StraightBetV3) HasPrice() bool {
	if o != nil && o.Price != nil {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float64 and assigns it to the Price field.
func (o *StraightBetV3) SetPrice(v float64) {
	o.Price = &v
}

// GetTeamName returns the TeamName field value if set, zero value otherwise.
func (o *StraightBetV3) GetTeamName() string {
	if o == nil || o.TeamName == nil {
		var ret string
		return ret
	}
	return *o.TeamName
}

// GetTeamNameOk returns a tuple with the TeamName field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *StraightBetV3) GetTeamNameOk() (string, bool) {
	if o == nil || o.TeamName == nil {
		var ret string
		return ret, false
	}
	return *o.TeamName, true
}

// HasTeamName returns a boolean if a field has been set.
func (o *StraightBetV3) HasTeamName() bool {
	if o != nil && o.TeamName != nil {
		return true
	}

	return false
}

// SetTeamName gets a reference to the given string and assigns it to the TeamName field.
func (o *StraightBetV3) SetTeamName(v string) {
	o.TeamName = &v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *StraightBetV3) GetSide() NullableString {
	if o == nil || o.Side == nil {
		var ret NullableString
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *StraightBetV3) GetSideOk() (NullableString, bool) {
	if o == nil || o.Side == nil {
		var ret NullableString
		return ret, false
	}
	return *o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *StraightBetV3) HasSide() bool {
	if o != nil && o.Side != nil {
		return true
	}

	return false
}

// SetSide gets a reference to the given NullableString and assigns it to the Side field.
func (o *StraightBetV3) SetSide(v NullableString) {
	o.Side = &v
}

// GetPitcher1 returns the Pitcher1 field value if set, zero value otherwise.
func (o *StraightBetV3) GetPitcher1() NullableString {
	if o == nil || o.Pitcher1 == nil {
		var ret NullableString
		return ret
	}
	return *o.Pitcher1
}

// GetPitcher1Ok returns a tuple with the Pitcher1 field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *StraightBetV3) GetPitcher1Ok() (NullableString, bool) {
	if o == nil || o.Pitcher1 == nil {
		var ret NullableString
		return ret, false
	}
	return *o.Pitcher1, true
}

// HasPitcher1 returns a boolean if a field has been set.
func (o *StraightBetV3) HasPitcher1() bool {
	if o != nil && o.Pitcher1 != nil {
		return true
	}

	return false
}

// SetPitcher1 gets a reference to the given NullableString and assigns it to the Pitcher1 field.
func (o *StraightBetV3) SetPitcher1(v NullableString) {
	o.Pitcher1 = &v
}

// GetPitcher2 returns the Pitcher2 field value if set, zero value otherwise.
func (o *StraightBetV3) GetPitcher2() NullableString {
	if o == nil || o.Pitcher2 == nil {
		var ret NullableString
		return ret
	}
	return *o.Pitcher2
}

// GetPitcher2Ok returns a tuple with the Pitcher2 field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *StraightBetV3) GetPitcher2Ok() (NullableString, bool) {
	if o == nil || o.Pitcher2 == nil {
		var ret NullableString
		return ret, false
	}
	return *o.Pitcher2, true
}

// HasPitcher2 returns a boolean if a field has been set.
func (o *StraightBetV3) HasPitcher2() bool {
	if o != nil && o.Pitcher2 != nil {
		return true
	}

	return false
}

// SetPitcher2 gets a reference to the given NullableString and assigns it to the Pitcher2 field.
func (o *StraightBetV3) SetPitcher2(v NullableString) {
	o.Pitcher2 = &v
}

// GetPitcher1MustStart returns the Pitcher1MustStart field value if set, zero value otherwise.
func (o *StraightBetV3) GetPitcher1MustStart() NullableBool {
	if o == nil || o.Pitcher1MustStart == nil {
		var ret NullableBool
		return ret
	}
	return *o.Pitcher1MustStart
}

// GetPitcher1MustStartOk returns a tuple with the Pitcher1MustStart field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *StraightBetV3) GetPitcher1MustStartOk() (NullableBool, bool) {
	if o == nil || o.Pitcher1MustStart == nil {
		var ret NullableBool
		return ret, false
	}
	return *o.Pitcher1MustStart, true
}

// HasPitcher1MustStart returns a boolean if a field has been set.
func (o *StraightBetV3) HasPitcher1MustStart() bool {
	if o != nil && o.Pitcher1MustStart != nil {
		return true
	}

	return false
}

// SetPitcher1MustStart gets a reference to the given NullableBool and assigns it to the Pitcher1MustStart field.
func (o *StraightBetV3) SetPitcher1MustStart(v NullableBool) {
	o.Pitcher1MustStart = &v
}

// GetPitcher2MustStart returns the Pitcher2MustStart field value if set, zero value otherwise.
func (o *StraightBetV3) GetPitcher2MustStart() NullableBool {
	if o == nil || o.Pitcher2MustStart == nil {
		var ret NullableBool
		return ret
	}
	return *o.Pitcher2MustStart
}

// GetPitcher2MustStartOk returns a tuple with the Pitcher2MustStart field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *StraightBetV3) GetPitcher2MustStartOk() (NullableBool, bool) {
	if o == nil || o.Pitcher2MustStart == nil {
		var ret NullableBool
		return ret, false
	}
	return *o.Pitcher2MustStart, true
}

// HasPitcher2MustStart returns a boolean if a field has been set.
func (o *StraightBetV3) HasPitcher2MustStart() bool {
	if o != nil && o.Pitcher2MustStart != nil {
		return true
	}

	return false
}

// SetPitcher2MustStart gets a reference to the given NullableBool and assigns it to the Pitcher2MustStart field.
func (o *StraightBetV3) SetPitcher2MustStart(v NullableBool) {
	o.Pitcher2MustStart = &v
}

// GetTeam1 returns the Team1 field value if set, zero value otherwise.
func (o *StraightBetV3) GetTeam1() string {
	if o == nil || o.Team1 == nil {
		var ret string
		return ret
	}
	return *o.Team1
}

// GetTeam1Ok returns a tuple with the Team1 field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *StraightBetV3) GetTeam1Ok() (string, bool) {
	if o == nil || o.Team1 == nil {
		var ret string
		return ret, false
	}
	return *o.Team1, true
}

// HasTeam1 returns a boolean if a field has been set.
func (o *StraightBetV3) HasTeam1() bool {
	if o != nil && o.Team1 != nil {
		return true
	}

	return false
}

// SetTeam1 gets a reference to the given string and assigns it to the Team1 field.
func (o *StraightBetV3) SetTeam1(v string) {
	o.Team1 = &v
}

// GetTeam2 returns the Team2 field value if set, zero value otherwise.
func (o *StraightBetV3) GetTeam2() string {
	if o == nil || o.Team2 == nil {
		var ret string
		return ret
	}
	return *o.Team2
}

// GetTeam2Ok returns a tuple with the Team2 field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *StraightBetV3) GetTeam2Ok() (string, bool) {
	if o == nil || o.Team2 == nil {
		var ret string
		return ret, false
	}
	return *o.Team2, true
}

// HasTeam2 returns a boolean if a field has been set.
func (o *StraightBetV3) HasTeam2() bool {
	if o != nil && o.Team2 != nil {
		return true
	}

	return false
}

// SetTeam2 gets a reference to the given string and assigns it to the Team2 field.
func (o *StraightBetV3) SetTeam2(v string) {
	o.Team2 = &v
}

// GetPeriodNumber returns the PeriodNumber field value if set, zero value otherwise.
func (o *StraightBetV3) GetPeriodNumber() int32 {
	if o == nil || o.PeriodNumber == nil {
		var ret int32
		return ret
	}
	return *o.PeriodNumber
}

// GetPeriodNumberOk returns a tuple with the PeriodNumber field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *StraightBetV3) GetPeriodNumberOk() (int32, bool) {
	if o == nil || o.PeriodNumber == nil {
		var ret int32
		return ret, false
	}
	return *o.PeriodNumber, true
}

// HasPeriodNumber returns a boolean if a field has been set.
func (o *StraightBetV3) HasPeriodNumber() bool {
	if o != nil && o.PeriodNumber != nil {
		return true
	}

	return false
}

// SetPeriodNumber gets a reference to the given int32 and assigns it to the PeriodNumber field.
func (o *StraightBetV3) SetPeriodNumber(v int32) {
	o.PeriodNumber = &v
}

// GetTeam1Score returns the Team1Score field value if set, zero value otherwise.
func (o *StraightBetV3) GetTeam1Score() NullableFloat64 {
	if o == nil || o.Team1Score == nil {
		var ret NullableFloat64
		return ret
	}
	return *o.Team1Score
}

// GetTeam1ScoreOk returns a tuple with the Team1Score field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *StraightBetV3) GetTeam1ScoreOk() (NullableFloat64, bool) {
	if o == nil || o.Team1Score == nil {
		var ret NullableFloat64
		return ret, false
	}
	return *o.Team1Score, true
}

// HasTeam1Score returns a boolean if a field has been set.
func (o *StraightBetV3) HasTeam1Score() bool {
	if o != nil && o.Team1Score != nil {
		return true
	}

	return false
}

// SetTeam1Score gets a reference to the given NullableFloat64 and assigns it to the Team1Score field.
func (o *StraightBetV3) SetTeam1Score(v NullableFloat64) {
	o.Team1Score = &v
}

// GetTeam2Score returns the Team2Score field value if set, zero value otherwise.
func (o *StraightBetV3) GetTeam2Score() NullableFloat64 {
	if o == nil || o.Team2Score == nil {
		var ret NullableFloat64
		return ret
	}
	return *o.Team2Score
}

// GetTeam2ScoreOk returns a tuple with the Team2Score field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *StraightBetV3) GetTeam2ScoreOk() (NullableFloat64, bool) {
	if o == nil || o.Team2Score == nil {
		var ret NullableFloat64
		return ret, false
	}
	return *o.Team2Score, true
}

// HasTeam2Score returns a boolean if a field has been set.
func (o *StraightBetV3) HasTeam2Score() bool {
	if o != nil && o.Team2Score != nil {
		return true
	}

	return false
}

// SetTeam2Score gets a reference to the given NullableFloat64 and assigns it to the Team2Score field.
func (o *StraightBetV3) SetTeam2Score(v NullableFloat64) {
	o.Team2Score = &v
}

// GetFtTeam1Score returns the FtTeam1Score field value if set, zero value otherwise.
func (o *StraightBetV3) GetFtTeam1Score() NullableFloat64 {
	if o == nil || o.FtTeam1Score == nil {
		var ret NullableFloat64
		return ret
	}
	return *o.FtTeam1Score
}

// GetFtTeam1ScoreOk returns a tuple with the FtTeam1Score field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *StraightBetV3) GetFtTeam1ScoreOk() (NullableFloat64, bool) {
	if o == nil || o.FtTeam1Score == nil {
		var ret NullableFloat64
		return ret, false
	}
	return *o.FtTeam1Score, true
}

// HasFtTeam1Score returns a boolean if a field has been set.
func (o *StraightBetV3) HasFtTeam1Score() bool {
	if o != nil && o.FtTeam1Score != nil {
		return true
	}

	return false
}

// SetFtTeam1Score gets a reference to the given NullableFloat64 and assigns it to the FtTeam1Score field.
func (o *StraightBetV3) SetFtTeam1Score(v NullableFloat64) {
	o.FtTeam1Score = &v
}

// GetFtTeam2Score returns the FtTeam2Score field value if set, zero value otherwise.
func (o *StraightBetV3) GetFtTeam2Score() NullableFloat64 {
	if o == nil || o.FtTeam2Score == nil {
		var ret NullableFloat64
		return ret
	}
	return *o.FtTeam2Score
}

// GetFtTeam2ScoreOk returns a tuple with the FtTeam2Score field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *StraightBetV3) GetFtTeam2ScoreOk() (NullableFloat64, bool) {
	if o == nil || o.FtTeam2Score == nil {
		var ret NullableFloat64
		return ret, false
	}
	return *o.FtTeam2Score, true
}

// HasFtTeam2Score returns a boolean if a field has been set.
func (o *StraightBetV3) HasFtTeam2Score() bool {
	if o != nil && o.FtTeam2Score != nil {
		return true
	}

	return false
}

// SetFtTeam2Score gets a reference to the given NullableFloat64 and assigns it to the FtTeam2Score field.
func (o *StraightBetV3) SetFtTeam2Score(v NullableFloat64) {
	o.FtTeam2Score = &v
}

// GetPTeam1Score returns the PTeam1Score field value if set, zero value otherwise.
func (o *StraightBetV3) GetPTeam1Score() NullableFloat64 {
	if o == nil || o.PTeam1Score == nil {
		var ret NullableFloat64
		return ret
	}
	return *o.PTeam1Score
}

// GetPTeam1ScoreOk returns a tuple with the PTeam1Score field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *StraightBetV3) GetPTeam1ScoreOk() (NullableFloat64, bool) {
	if o == nil || o.PTeam1Score == nil {
		var ret NullableFloat64
		return ret, false
	}
	return *o.PTeam1Score, true
}

// HasPTeam1Score returns a boolean if a field has been set.
func (o *StraightBetV3) HasPTeam1Score() bool {
	if o != nil && o.PTeam1Score != nil {
		return true
	}

	return false
}

// SetPTeam1Score gets a reference to the given NullableFloat64 and assigns it to the PTeam1Score field.
func (o *StraightBetV3) SetPTeam1Score(v NullableFloat64) {
	o.PTeam1Score = &v
}

// GetPTeam2Score returns the PTeam2Score field value if set, zero value otherwise.
func (o *StraightBetV3) GetPTeam2Score() NullableFloat64 {
	if o == nil || o.PTeam2Score == nil {
		var ret NullableFloat64
		return ret
	}
	return *o.PTeam2Score
}

// GetPTeam2ScoreOk returns a tuple with the PTeam2Score field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *StraightBetV3) GetPTeam2ScoreOk() (NullableFloat64, bool) {
	if o == nil || o.PTeam2Score == nil {
		var ret NullableFloat64
		return ret, false
	}
	return *o.PTeam2Score, true
}

// HasPTeam2Score returns a boolean if a field has been set.
func (o *StraightBetV3) HasPTeam2Score() bool {
	if o != nil && o.PTeam2Score != nil {
		return true
	}

	return false
}

// SetPTeam2Score gets a reference to the given NullableFloat64 and assigns it to the PTeam2Score field.
func (o *StraightBetV3) SetPTeam2Score(v NullableFloat64) {
	o.PTeam2Score = &v
}

// GetIsLive returns the IsLive field value if set, zero value otherwise.
func (o *StraightBetV3) GetIsLive() bool {
	if o == nil || o.IsLive == nil {
		var ret bool
		return ret
	}
	return *o.IsLive
}

// GetIsLiveOk returns a tuple with the IsLive field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *StraightBetV3) GetIsLiveOk() (bool, bool) {
	if o == nil || o.IsLive == nil {
		var ret bool
		return ret, false
	}
	return *o.IsLive, true
}

// HasIsLive returns a boolean if a field has been set.
func (o *StraightBetV3) HasIsLive() bool {
	if o != nil && o.IsLive != nil {
		return true
	}

	return false
}

// SetIsLive gets a reference to the given bool and assigns it to the IsLive field.
func (o *StraightBetV3) SetIsLive(v bool) {
	o.IsLive = &v
}

// GetEventStartTime returns the EventStartTime field value if set, zero value otherwise.
func (o *StraightBetV3) GetEventStartTime() time.Time {
	if o == nil || o.EventStartTime == nil {
		var ret time.Time
		return ret
	}
	return *o.EventStartTime
}

// GetEventStartTimeOk returns a tuple with the EventStartTime field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *StraightBetV3) GetEventStartTimeOk() (time.Time, bool) {
	if o == nil || o.EventStartTime == nil {
		var ret time.Time
		return ret, false
	}
	return *o.EventStartTime, true
}

// HasEventStartTime returns a boolean if a field has been set.
func (o *StraightBetV3) HasEventStartTime() bool {
	if o != nil && o.EventStartTime != nil {
		return true
	}

	return false
}

// SetEventStartTime gets a reference to the given time.Time and assigns it to the EventStartTime field.
func (o *StraightBetV3) SetEventStartTime(v time.Time) {
	o.EventStartTime = &v
}

type NullableStraightBetV3 struct {
	Value        StraightBetV3
	ExplicitNull bool
}

func (v NullableStraightBetV3) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableStraightBetV3) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
