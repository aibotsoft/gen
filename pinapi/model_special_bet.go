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

// SpecialBet struct for SpecialBet
type SpecialBet struct {
	// Bet identification
	BetId int64 `json:"betId"`
	// Unique Request Id
	UniqueRequestId *string `json:"uniqueRequestId,omitempty"`
	// Wager identification. All bets placed thru the API will have value 1. Website Classic view supports multiple contest(special) bets placement in the same bet slip in that case the bet would have appropriate wager number, as well as all round robin parlay bets.
	WagerNumber int `json:"wagerNumber"`
	// Date time when the bet was placed.
	PlacedAt time.Time `json:"placedAt"`
	// Bet Status.  ACCEPTED = Bet was accepted,  CANCELLED = Bet is cancelled as per Pinnacle betting rules,  LOSE = The bet is settled as lose, REFUNDED = When an event is cancelled or when the bet is settled as push, the bet will have REFUNDED status,  WON = The bet is settled as won  
	BetStatus string `json:"betStatus"`
	BetType string `json:"betType"`
	// Win amount.
	Win float64 `json:"win"`
	// Risk amount.
	Risk float64 `json:"risk"`
	// Win-Loss for settled bets.
	WinLoss NullableFloat64 `json:"winLoss,omitempty"`
	OddsFormat OddsFormat `json:"oddsFormat"`
	// Client’s commission on the bet.
	CustomerCommission NullableFloat64 `json:"customerCommission,omitempty"`
	CancellationReason *CancellationReason `json:"cancellationReason,omitempty"`
	// Update Sequence. It gets updated when the bet status change.
	UpdateSequence int64 `json:"updateSequence"`
	SpecialId int64 `json:"specialId"`
	SpecialName string `json:"specialName"`
	ContestantId int64 `json:"contestantId"`
	ContestantName string `json:"contestantName"`
	Price float64 `json:"price"`
	Handicap *float64 `json:"handicap,omitempty"`
	Units *string `json:"units,omitempty"`
	SportId int `json:"sportId"`
	LeagueId int `json:"leagueId"`
	// Populated if bet was placed on a special linked to the event.
	EventId NullableInt64 `json:"eventId,omitempty"`
	// Populated if bet was placed on a special linked to the event.
	PeriodNumber NullableInt `json:"periodNumber,omitempty"`
	// Populated if bet was placed on a special linked to the event.
	Team1 NullableString `json:"team1,omitempty"`
	// Populated if bet was placed on a special linked to the event.
	Team2 NullableString `json:"team2,omitempty"`
	// Date time when the event starts
	EventStartTime *time.Time `json:"eventStartTime,omitempty"`
}

// NewSpecialBet instantiates a new SpecialBet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpecialBet(betId int64, wagerNumber int, placedAt time.Time, betStatus string, betType string, win float64, risk float64, oddsFormat OddsFormat, updateSequence int64, specialId int64, specialName string, contestantId int64, contestantName string, price float64, sportId int, leagueId int, ) *SpecialBet {
	this := SpecialBet{}
	this.BetId = betId
	this.WagerNumber = wagerNumber
	this.PlacedAt = placedAt
	this.BetStatus = betStatus
	this.BetType = betType
	this.Win = win
	this.Risk = risk
	this.OddsFormat = oddsFormat
	this.UpdateSequence = updateSequence
	this.SpecialId = specialId
	this.SpecialName = specialName
	this.ContestantId = contestantId
	this.ContestantName = contestantName
	this.Price = price
	this.SportId = sportId
	this.LeagueId = leagueId
	return &this
}

// NewSpecialBetWithDefaults instantiates a new SpecialBet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpecialBetWithDefaults() *SpecialBet {
	this := SpecialBet{}
	var betType string = "SPECIAL"
	this.BetType = betType
	return &this
}

// GetBetId returns the BetId field value
func (o *SpecialBet) GetBetId() int64 {
	if o == nil  {
		var ret int64
		return ret
	}

	return o.BetId
}

// GetBetIdOk returns a tuple with the BetId field value
// and a boolean to check if the value has been set.
func (o *SpecialBet) GetBetIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BetId, true
}

// SetBetId sets field value
func (o *SpecialBet) SetBetId(v int64) {
	o.BetId = v
}

// GetUniqueRequestId returns the UniqueRequestId field value if set, zero value otherwise.
func (o *SpecialBet) GetUniqueRequestId() string {
	if o == nil || o.UniqueRequestId == nil {
		var ret string
		return ret
	}
	return *o.UniqueRequestId
}

// GetUniqueRequestIdOk returns a tuple with the UniqueRequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpecialBet) GetUniqueRequestIdOk() (*string, bool) {
	if o == nil || o.UniqueRequestId == nil {
		return nil, false
	}
	return o.UniqueRequestId, true
}

// HasUniqueRequestId returns a boolean if a field has been set.
func (o *SpecialBet) HasUniqueRequestId() bool {
	if o != nil && o.UniqueRequestId != nil {
		return true
	}

	return false
}

// SetUniqueRequestId gets a reference to the given string and assigns it to the UniqueRequestId field.
func (o *SpecialBet) SetUniqueRequestId(v string) {
	o.UniqueRequestId = &v
}

// GetWagerNumber returns the WagerNumber field value
func (o *SpecialBet) GetWagerNumber() int {
	if o == nil  {
		var ret int
		return ret
	}

	return o.WagerNumber
}

// GetWagerNumberOk returns a tuple with the WagerNumber field value
// and a boolean to check if the value has been set.
func (o *SpecialBet) GetWagerNumberOk() (*int, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WagerNumber, true
}

// SetWagerNumber sets field value
func (o *SpecialBet) SetWagerNumber(v int) {
	o.WagerNumber = v
}

// GetPlacedAt returns the PlacedAt field value
func (o *SpecialBet) GetPlacedAt() time.Time {
	if o == nil  {
		var ret time.Time
		return ret
	}

	return o.PlacedAt
}

// GetPlacedAtOk returns a tuple with the PlacedAt field value
// and a boolean to check if the value has been set.
func (o *SpecialBet) GetPlacedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PlacedAt, true
}

// SetPlacedAt sets field value
func (o *SpecialBet) SetPlacedAt(v time.Time) {
	o.PlacedAt = v
}

// GetBetStatus returns the BetStatus field value
func (o *SpecialBet) GetBetStatus() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.BetStatus
}

// GetBetStatusOk returns a tuple with the BetStatus field value
// and a boolean to check if the value has been set.
func (o *SpecialBet) GetBetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BetStatus, true
}

// SetBetStatus sets field value
func (o *SpecialBet) SetBetStatus(v string) {
	o.BetStatus = v
}

// GetBetType returns the BetType field value
func (o *SpecialBet) GetBetType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.BetType
}

// GetBetTypeOk returns a tuple with the BetType field value
// and a boolean to check if the value has been set.
func (o *SpecialBet) GetBetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BetType, true
}

// SetBetType sets field value
func (o *SpecialBet) SetBetType(v string) {
	o.BetType = v
}

// GetWin returns the Win field value
func (o *SpecialBet) GetWin() float64 {
	if o == nil  {
		var ret float64
		return ret
	}

	return o.Win
}

// GetWinOk returns a tuple with the Win field value
// and a boolean to check if the value has been set.
func (o *SpecialBet) GetWinOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Win, true
}

// SetWin sets field value
func (o *SpecialBet) SetWin(v float64) {
	o.Win = v
}

// GetRisk returns the Risk field value
func (o *SpecialBet) GetRisk() float64 {
	if o == nil  {
		var ret float64
		return ret
	}

	return o.Risk
}

// GetRiskOk returns a tuple with the Risk field value
// and a boolean to check if the value has been set.
func (o *SpecialBet) GetRiskOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Risk, true
}

// SetRisk sets field value
func (o *SpecialBet) SetRisk(v float64) {
	o.Risk = v
}

// GetWinLoss returns the WinLoss field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SpecialBet) GetWinLoss() float64 {
	if o == nil || o.WinLoss.Get() == nil {
		var ret float64
		return ret
	}
	return *o.WinLoss.Get()
}

// GetWinLossOk returns a tuple with the WinLoss field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SpecialBet) GetWinLossOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.WinLoss.Get(), o.WinLoss.IsSet()
}

// HasWinLoss returns a boolean if a field has been set.
func (o *SpecialBet) HasWinLoss() bool {
	if o != nil && o.WinLoss.IsSet() {
		return true
	}

	return false
}

// SetWinLoss gets a reference to the given NullableFloat64 and assigns it to the WinLoss field.
func (o *SpecialBet) SetWinLoss(v float64) {
	o.WinLoss.Set(&v)
}
// SetWinLossNil sets the value for WinLoss to be an explicit nil
func (o *SpecialBet) SetWinLossNil() {
	o.WinLoss.Set(nil)
}

// UnsetWinLoss ensures that no value is present for WinLoss, not even an explicit nil
func (o *SpecialBet) UnsetWinLoss() {
	o.WinLoss.Unset()
}

// GetOddsFormat returns the OddsFormat field value
func (o *SpecialBet) GetOddsFormat() OddsFormat {
	if o == nil  {
		var ret OddsFormat
		return ret
	}

	return o.OddsFormat
}

// GetOddsFormatOk returns a tuple with the OddsFormat field value
// and a boolean to check if the value has been set.
func (o *SpecialBet) GetOddsFormatOk() (*OddsFormat, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OddsFormat, true
}

// SetOddsFormat sets field value
func (o *SpecialBet) SetOddsFormat(v OddsFormat) {
	o.OddsFormat = v
}

// GetCustomerCommission returns the CustomerCommission field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SpecialBet) GetCustomerCommission() float64 {
	if o == nil || o.CustomerCommission.Get() == nil {
		var ret float64
		return ret
	}
	return *o.CustomerCommission.Get()
}

// GetCustomerCommissionOk returns a tuple with the CustomerCommission field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SpecialBet) GetCustomerCommissionOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CustomerCommission.Get(), o.CustomerCommission.IsSet()
}

// HasCustomerCommission returns a boolean if a field has been set.
func (o *SpecialBet) HasCustomerCommission() bool {
	if o != nil && o.CustomerCommission.IsSet() {
		return true
	}

	return false
}

// SetCustomerCommission gets a reference to the given NullableFloat64 and assigns it to the CustomerCommission field.
func (o *SpecialBet) SetCustomerCommission(v float64) {
	o.CustomerCommission.Set(&v)
}
// SetCustomerCommissionNil sets the value for CustomerCommission to be an explicit nil
func (o *SpecialBet) SetCustomerCommissionNil() {
	o.CustomerCommission.Set(nil)
}

// UnsetCustomerCommission ensures that no value is present for CustomerCommission, not even an explicit nil
func (o *SpecialBet) UnsetCustomerCommission() {
	o.CustomerCommission.Unset()
}

// GetCancellationReason returns the CancellationReason field value if set, zero value otherwise.
func (o *SpecialBet) GetCancellationReason() CancellationReason {
	if o == nil || o.CancellationReason == nil {
		var ret CancellationReason
		return ret
	}
	return *o.CancellationReason
}

// GetCancellationReasonOk returns a tuple with the CancellationReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpecialBet) GetCancellationReasonOk() (*CancellationReason, bool) {
	if o == nil || o.CancellationReason == nil {
		return nil, false
	}
	return o.CancellationReason, true
}

// HasCancellationReason returns a boolean if a field has been set.
func (o *SpecialBet) HasCancellationReason() bool {
	if o != nil && o.CancellationReason != nil {
		return true
	}

	return false
}

// SetCancellationReason gets a reference to the given CancellationReason and assigns it to the CancellationReason field.
func (o *SpecialBet) SetCancellationReason(v CancellationReason) {
	o.CancellationReason = &v
}

// GetUpdateSequence returns the UpdateSequence field value
func (o *SpecialBet) GetUpdateSequence() int64 {
	if o == nil  {
		var ret int64
		return ret
	}

	return o.UpdateSequence
}

// GetUpdateSequenceOk returns a tuple with the UpdateSequence field value
// and a boolean to check if the value has been set.
func (o *SpecialBet) GetUpdateSequenceOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UpdateSequence, true
}

// SetUpdateSequence sets field value
func (o *SpecialBet) SetUpdateSequence(v int64) {
	o.UpdateSequence = v
}

// GetSpecialId returns the SpecialId field value
func (o *SpecialBet) GetSpecialId() int64 {
	if o == nil  {
		var ret int64
		return ret
	}

	return o.SpecialId
}

// GetSpecialIdOk returns a tuple with the SpecialId field value
// and a boolean to check if the value has been set.
func (o *SpecialBet) GetSpecialIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SpecialId, true
}

// SetSpecialId sets field value
func (o *SpecialBet) SetSpecialId(v int64) {
	o.SpecialId = v
}

// GetSpecialName returns the SpecialName field value
func (o *SpecialBet) GetSpecialName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.SpecialName
}

// GetSpecialNameOk returns a tuple with the SpecialName field value
// and a boolean to check if the value has been set.
func (o *SpecialBet) GetSpecialNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SpecialName, true
}

// SetSpecialName sets field value
func (o *SpecialBet) SetSpecialName(v string) {
	o.SpecialName = v
}

// GetContestantId returns the ContestantId field value
func (o *SpecialBet) GetContestantId() int64 {
	if o == nil  {
		var ret int64
		return ret
	}

	return o.ContestantId
}

// GetContestantIdOk returns a tuple with the ContestantId field value
// and a boolean to check if the value has been set.
func (o *SpecialBet) GetContestantIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ContestantId, true
}

// SetContestantId sets field value
func (o *SpecialBet) SetContestantId(v int64) {
	o.ContestantId = v
}

// GetContestantName returns the ContestantName field value
func (o *SpecialBet) GetContestantName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ContestantName
}

// GetContestantNameOk returns a tuple with the ContestantName field value
// and a boolean to check if the value has been set.
func (o *SpecialBet) GetContestantNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ContestantName, true
}

// SetContestantName sets field value
func (o *SpecialBet) SetContestantName(v string) {
	o.ContestantName = v
}

// GetPrice returns the Price field value
func (o *SpecialBet) GetPrice() float64 {
	if o == nil  {
		var ret float64
		return ret
	}

	return o.Price
}

// GetPriceOk returns a tuple with the Price field value
// and a boolean to check if the value has been set.
func (o *SpecialBet) GetPriceOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Price, true
}

// SetPrice sets field value
func (o *SpecialBet) SetPrice(v float64) {
	o.Price = v
}

// GetHandicap returns the Handicap field value if set, zero value otherwise.
func (o *SpecialBet) GetHandicap() float64 {
	if o == nil || o.Handicap == nil {
		var ret float64
		return ret
	}
	return *o.Handicap
}

// GetHandicapOk returns a tuple with the Handicap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpecialBet) GetHandicapOk() (*float64, bool) {
	if o == nil || o.Handicap == nil {
		return nil, false
	}
	return o.Handicap, true
}

// HasHandicap returns a boolean if a field has been set.
func (o *SpecialBet) HasHandicap() bool {
	if o != nil && o.Handicap != nil {
		return true
	}

	return false
}

// SetHandicap gets a reference to the given float64 and assigns it to the Handicap field.
func (o *SpecialBet) SetHandicap(v float64) {
	o.Handicap = &v
}

// GetUnits returns the Units field value if set, zero value otherwise.
func (o *SpecialBet) GetUnits() string {
	if o == nil || o.Units == nil {
		var ret string
		return ret
	}
	return *o.Units
}

// GetUnitsOk returns a tuple with the Units field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpecialBet) GetUnitsOk() (*string, bool) {
	if o == nil || o.Units == nil {
		return nil, false
	}
	return o.Units, true
}

// HasUnits returns a boolean if a field has been set.
func (o *SpecialBet) HasUnits() bool {
	if o != nil && o.Units != nil {
		return true
	}

	return false
}

// SetUnits gets a reference to the given string and assigns it to the Units field.
func (o *SpecialBet) SetUnits(v string) {
	o.Units = &v
}

// GetSportId returns the SportId field value
func (o *SpecialBet) GetSportId() int {
	if o == nil  {
		var ret int
		return ret
	}

	return o.SportId
}

// GetSportIdOk returns a tuple with the SportId field value
// and a boolean to check if the value has been set.
func (o *SpecialBet) GetSportIdOk() (*int, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SportId, true
}

// SetSportId sets field value
func (o *SpecialBet) SetSportId(v int) {
	o.SportId = v
}

// GetLeagueId returns the LeagueId field value
func (o *SpecialBet) GetLeagueId() int {
	if o == nil  {
		var ret int
		return ret
	}

	return o.LeagueId
}

// GetLeagueIdOk returns a tuple with the LeagueId field value
// and a boolean to check if the value has been set.
func (o *SpecialBet) GetLeagueIdOk() (*int, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LeagueId, true
}

// SetLeagueId sets field value
func (o *SpecialBet) SetLeagueId(v int) {
	o.LeagueId = v
}

// GetEventId returns the EventId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SpecialBet) GetEventId() int64 {
	if o == nil || o.EventId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.EventId.Get()
}

// GetEventIdOk returns a tuple with the EventId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SpecialBet) GetEventIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EventId.Get(), o.EventId.IsSet()
}

// HasEventId returns a boolean if a field has been set.
func (o *SpecialBet) HasEventId() bool {
	if o != nil && o.EventId.IsSet() {
		return true
	}

	return false
}

// SetEventId gets a reference to the given NullableInt64 and assigns it to the EventId field.
func (o *SpecialBet) SetEventId(v int64) {
	o.EventId.Set(&v)
}
// SetEventIdNil sets the value for EventId to be an explicit nil
func (o *SpecialBet) SetEventIdNil() {
	o.EventId.Set(nil)
}

// UnsetEventId ensures that no value is present for EventId, not even an explicit nil
func (o *SpecialBet) UnsetEventId() {
	o.EventId.Unset()
}

// GetPeriodNumber returns the PeriodNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SpecialBet) GetPeriodNumber() int {
	if o == nil || o.PeriodNumber.Get() == nil {
		var ret int
		return ret
	}
	return *o.PeriodNumber.Get()
}

// GetPeriodNumberOk returns a tuple with the PeriodNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SpecialBet) GetPeriodNumberOk() (*int, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PeriodNumber.Get(), o.PeriodNumber.IsSet()
}

// HasPeriodNumber returns a boolean if a field has been set.
func (o *SpecialBet) HasPeriodNumber() bool {
	if o != nil && o.PeriodNumber.IsSet() {
		return true
	}

	return false
}

// SetPeriodNumber gets a reference to the given NullableInt and assigns it to the PeriodNumber field.
func (o *SpecialBet) SetPeriodNumber(v int) {
	o.PeriodNumber.Set(&v)
}
// SetPeriodNumberNil sets the value for PeriodNumber to be an explicit nil
func (o *SpecialBet) SetPeriodNumberNil() {
	o.PeriodNumber.Set(nil)
}

// UnsetPeriodNumber ensures that no value is present for PeriodNumber, not even an explicit nil
func (o *SpecialBet) UnsetPeriodNumber() {
	o.PeriodNumber.Unset()
}

// GetTeam1 returns the Team1 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SpecialBet) GetTeam1() string {
	if o == nil || o.Team1.Get() == nil {
		var ret string
		return ret
	}
	return *o.Team1.Get()
}

// GetTeam1Ok returns a tuple with the Team1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SpecialBet) GetTeam1Ok() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Team1.Get(), o.Team1.IsSet()
}

// HasTeam1 returns a boolean if a field has been set.
func (o *SpecialBet) HasTeam1() bool {
	if o != nil && o.Team1.IsSet() {
		return true
	}

	return false
}

// SetTeam1 gets a reference to the given NullableString and assigns it to the Team1 field.
func (o *SpecialBet) SetTeam1(v string) {
	o.Team1.Set(&v)
}
// SetTeam1Nil sets the value for Team1 to be an explicit nil
func (o *SpecialBet) SetTeam1Nil() {
	o.Team1.Set(nil)
}

// UnsetTeam1 ensures that no value is present for Team1, not even an explicit nil
func (o *SpecialBet) UnsetTeam1() {
	o.Team1.Unset()
}

// GetTeam2 returns the Team2 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SpecialBet) GetTeam2() string {
	if o == nil || o.Team2.Get() == nil {
		var ret string
		return ret
	}
	return *o.Team2.Get()
}

// GetTeam2Ok returns a tuple with the Team2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SpecialBet) GetTeam2Ok() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Team2.Get(), o.Team2.IsSet()
}

// HasTeam2 returns a boolean if a field has been set.
func (o *SpecialBet) HasTeam2() bool {
	if o != nil && o.Team2.IsSet() {
		return true
	}

	return false
}

// SetTeam2 gets a reference to the given NullableString and assigns it to the Team2 field.
func (o *SpecialBet) SetTeam2(v string) {
	o.Team2.Set(&v)
}
// SetTeam2Nil sets the value for Team2 to be an explicit nil
func (o *SpecialBet) SetTeam2Nil() {
	o.Team2.Set(nil)
}

// UnsetTeam2 ensures that no value is present for Team2, not even an explicit nil
func (o *SpecialBet) UnsetTeam2() {
	o.Team2.Unset()
}

// GetEventStartTime returns the EventStartTime field value if set, zero value otherwise.
func (o *SpecialBet) GetEventStartTime() time.Time {
	if o == nil || o.EventStartTime == nil {
		var ret time.Time
		return ret
	}
	return *o.EventStartTime
}

// GetEventStartTimeOk returns a tuple with the EventStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpecialBet) GetEventStartTimeOk() (*time.Time, bool) {
	if o == nil || o.EventStartTime == nil {
		return nil, false
	}
	return o.EventStartTime, true
}

// HasEventStartTime returns a boolean if a field has been set.
func (o *SpecialBet) HasEventStartTime() bool {
	if o != nil && o.EventStartTime != nil {
		return true
	}

	return false
}

// SetEventStartTime gets a reference to the given time.Time and assigns it to the EventStartTime field.
func (o *SpecialBet) SetEventStartTime(v time.Time) {
	o.EventStartTime = &v
}

func (o SpecialBet) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["betId"] = o.BetId
	}
	if o.UniqueRequestId != nil {
		toSerialize["uniqueRequestId"] = o.UniqueRequestId
	}
	if true {
		toSerialize["wagerNumber"] = o.WagerNumber
	}
	if true {
		toSerialize["placedAt"] = o.PlacedAt
	}
	if true {
		toSerialize["betStatus"] = o.BetStatus
	}
	if true {
		toSerialize["betType"] = o.BetType
	}
	if true {
		toSerialize["win"] = o.Win
	}
	if true {
		toSerialize["risk"] = o.Risk
	}
	if o.WinLoss.IsSet() {
		toSerialize["winLoss"] = o.WinLoss.Get()
	}
	if true {
		toSerialize["oddsFormat"] = o.OddsFormat
	}
	if o.CustomerCommission.IsSet() {
		toSerialize["customerCommission"] = o.CustomerCommission.Get()
	}
	if o.CancellationReason != nil {
		toSerialize["cancellationReason"] = o.CancellationReason
	}
	if true {
		toSerialize["updateSequence"] = o.UpdateSequence
	}
	if true {
		toSerialize["specialId"] = o.SpecialId
	}
	if true {
		toSerialize["specialName"] = o.SpecialName
	}
	if true {
		toSerialize["contestantId"] = o.ContestantId
	}
	if true {
		toSerialize["contestantName"] = o.ContestantName
	}
	if true {
		toSerialize["price"] = o.Price
	}
	if o.Handicap != nil {
		toSerialize["handicap"] = o.Handicap
	}
	if o.Units != nil {
		toSerialize["units"] = o.Units
	}
	if true {
		toSerialize["sportId"] = o.SportId
	}
	if true {
		toSerialize["leagueId"] = o.LeagueId
	}
	if o.EventId.IsSet() {
		toSerialize["eventId"] = o.EventId.Get()
	}
	if o.PeriodNumber.IsSet() {
		toSerialize["periodNumber"] = o.PeriodNumber.Get()
	}
	if o.Team1.IsSet() {
		toSerialize["team1"] = o.Team1.Get()
	}
	if o.Team2.IsSet() {
		toSerialize["team2"] = o.Team2.Get()
	}
	if o.EventStartTime != nil {
		toSerialize["eventStartTime"] = o.EventStartTime
	}
	return json.Marshal(toSerialize)
}

type NullableSpecialBet struct {
	value *SpecialBet
	isSet bool
}

func (v NullableSpecialBet) Get() *SpecialBet {
	return v.value
}

func (v *NullableSpecialBet) Set(val *SpecialBet) {
	v.value = val
	v.isSet = true
}

func (v NullableSpecialBet) IsSet() bool {
	return v.isSet
}

func (v *NullableSpecialBet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpecialBet(val *SpecialBet) *NullableSpecialBet {
	return &NullableSpecialBet{value: val, isSet: true}
}

func (v NullableSpecialBet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpecialBet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
