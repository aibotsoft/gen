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

// StraightBetV3 struct for StraightBetV3
type StraightBetV3 struct {
	// Bet identification
	BetId int64 `json:"betId"`
	// Wager identification. All bets placed thru the API will have value 1. Website Classic view supports multiple contest(special) bets placement in the same bet slip in that case the bet would have appropriate wager number, as well as all round robin parlay bets.
	WagerNumber int `json:"wagerNumber"`
	// Date time when the bet was placed.
	PlacedAt time.Time `json:"placedAt"`
	// Bet Status.    ACCEPTED = Bet was accepted,   CANCELLED = Bet is cancelled as per Pinnacle betting rules,   LOSE = The bet is settled as lose,   PENDING_ACCEPTANCE = This status is reserved only for live bets. If a live bet is placed during danger zone or live delay is applied, it will be in PENDING_ACCEPTANCE , otherwise in ACCEPTED status. From this status bet can go to ACCEPTED or NOT_ACCEPTED status,   REFUNDED = When an event is cancelled or when the bet is settled as push, the bet will have REFUNDED status,   NOT_ACCEPTED = Bet was not accepted. Bet can be in this status only if it was previously in PENDING_ACCEPTANCE status,   WON = The bet is settled as won  
	BetStatus string `json:"betStatus"`
	// Bet type.
	BetType string `json:"betType"`
	// Win amount.
	Win float64 `json:"win"`
	// Risk amount.
	Risk float64 `json:"risk"`
	// Win-Loss for settled bets.
	WinLoss *float64 `json:"winLoss,omitempty"`
	OddsFormat OddsFormat `json:"oddsFormat"`
	// Client’s commission on the bet.
	CustomerCommission *float64 `json:"customerCommission,omitempty"`
	// временная замена
	CancellationReason *string `json:"cancellationReason,omitempty"`
	// Update Sequence
	UpdateSequence int64 `json:"updateSequence"`
	SportId *int `json:"sportId,omitempty"`
	LeagueId *int `json:"leagueId,omitempty"`
	EventId *int64 `json:"eventId,omitempty"`
	Handicap *float64 `json:"handicap,omitempty"`
	Price *float64 `json:"price,omitempty"`
	TeamName *string `json:"teamName,omitempty"`
	// Side type.
	Side *string `json:"side,omitempty"`
	// Pitcher name of team1. Only for bets on baseball.
	Pitcher1 *string `json:"pitcher1,omitempty"`
	// Pitcher name of team2. Only for bets on baseball.
	Pitcher2 *string `json:"pitcher2,omitempty"`
	// Baseball only. Refers to the pitcher for Team1.  This applicable only for MONEYLINE bet type, for all other bet types this has to be TRUE.
	Pitcher1MustStart *bool `json:"pitcher1MustStart,omitempty"`
	// Baseball only. Refers to the pitcher for Team2.  This applicable only for MONEYLINE bet type, for all other bet types this has to be TRUE.
	Pitcher2MustStart *bool `json:"pitcher2MustStart,omitempty"`
	Team1 *string `json:"team1,omitempty"`
	Team2 *string `json:"team2,omitempty"`
	PeriodNumber *int `json:"periodNumber,omitempty"`
	// Team 1 score that the bet was placed on, only for live bets.
	Team1Score *float64 `json:"team1Score,omitempty"`
	// Team 2 score that the bet was placed, only for live bets.
	Team2Score *float64 `json:"team2Score,omitempty"`
	// Full time team 1 score, only for settled bets.
	FtTeam1Score *float64 `json:"ftTeam1Score,omitempty"`
	// Full time team 2 score, only for settled bets.
	FtTeam2Score *float64 `json:"ftTeam2Score,omitempty"`
	// .End of period team 1 score, only for settled bets. If the bet was placed on Game period (periodNumber =0), this will be null . 
	PTeam1Score *float64 `json:"pTeam1Score,omitempty"`
	// End of period team 2 score, only for settled bets. If the bet was placed on Game period (periodNumber =0), this will be null
	PTeam2Score *float64 `json:"pTeam2Score,omitempty"`
	// Date time when the event starts.
	EventStartTime *time.Time `json:"eventStartTime,omitempty"`
}

// NewStraightBetV3 instantiates a new StraightBetV3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStraightBetV3(betId int64, wagerNumber int, placedAt time.Time, betStatus string, betType string, win float64, risk float64, oddsFormat OddsFormat, updateSequence int64, ) *StraightBetV3 {
	this := StraightBetV3{}
	this.BetId = betId
	this.WagerNumber = wagerNumber
	this.PlacedAt = placedAt
	this.BetStatus = betStatus
	this.BetType = betType
	this.Win = win
	this.Risk = risk
	this.OddsFormat = oddsFormat
	this.UpdateSequence = updateSequence
	return &this
}

// NewStraightBetV3WithDefaults instantiates a new StraightBetV3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStraightBetV3WithDefaults() *StraightBetV3 {
	this := StraightBetV3{}
	return &this
}

// GetBetId returns the BetId field value
func (o *StraightBetV3) GetBetId() int64 {
	if o == nil  {
		var ret int64
		return ret
	}

	return o.BetId
}

// GetBetIdOk returns a tuple with the BetId field value
// and a boolean to check if the value has been set.
func (o *StraightBetV3) GetBetIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BetId, true
}

// SetBetId sets field value
func (o *StraightBetV3) SetBetId(v int64) {
	o.BetId = v
}

// GetWagerNumber returns the WagerNumber field value
func (o *StraightBetV3) GetWagerNumber() int {
	if o == nil  {
		var ret int
		return ret
	}

	return o.WagerNumber
}

// GetWagerNumberOk returns a tuple with the WagerNumber field value
// and a boolean to check if the value has been set.
func (o *StraightBetV3) GetWagerNumberOk() (*int, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WagerNumber, true
}

// SetWagerNumber sets field value
func (o *StraightBetV3) SetWagerNumber(v int) {
	o.WagerNumber = v
}

// GetPlacedAt returns the PlacedAt field value
func (o *StraightBetV3) GetPlacedAt() time.Time {
	if o == nil  {
		var ret time.Time
		return ret
	}

	return o.PlacedAt
}

// GetPlacedAtOk returns a tuple with the PlacedAt field value
// and a boolean to check if the value has been set.
func (o *StraightBetV3) GetPlacedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PlacedAt, true
}

// SetPlacedAt sets field value
func (o *StraightBetV3) SetPlacedAt(v time.Time) {
	o.PlacedAt = v
}

// GetBetStatus returns the BetStatus field value
func (o *StraightBetV3) GetBetStatus() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.BetStatus
}

// GetBetStatusOk returns a tuple with the BetStatus field value
// and a boolean to check if the value has been set.
func (o *StraightBetV3) GetBetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BetStatus, true
}

// SetBetStatus sets field value
func (o *StraightBetV3) SetBetStatus(v string) {
	o.BetStatus = v
}

// GetBetType returns the BetType field value
func (o *StraightBetV3) GetBetType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.BetType
}

// GetBetTypeOk returns a tuple with the BetType field value
// and a boolean to check if the value has been set.
func (o *StraightBetV3) GetBetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BetType, true
}

// SetBetType sets field value
func (o *StraightBetV3) SetBetType(v string) {
	o.BetType = v
}

// GetWin returns the Win field value
func (o *StraightBetV3) GetWin() float64 {
	if o == nil  {
		var ret float64
		return ret
	}

	return o.Win
}

// GetWinOk returns a tuple with the Win field value
// and a boolean to check if the value has been set.
func (o *StraightBetV3) GetWinOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Win, true
}

// SetWin sets field value
func (o *StraightBetV3) SetWin(v float64) {
	o.Win = v
}

// GetRisk returns the Risk field value
func (o *StraightBetV3) GetRisk() float64 {
	if o == nil  {
		var ret float64
		return ret
	}

	return o.Risk
}

// GetRiskOk returns a tuple with the Risk field value
// and a boolean to check if the value has been set.
func (o *StraightBetV3) GetRiskOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Risk, true
}

// SetRisk sets field value
func (o *StraightBetV3) SetRisk(v float64) {
	o.Risk = v
}

// GetWinLoss returns the WinLoss field value if set, zero value otherwise.
func (o *StraightBetV3) GetWinLoss() float64 {
	if o == nil || o.WinLoss == nil {
		var ret float64
		return ret
	}
	return *o.WinLoss
}

// GetWinLossOk returns a tuple with the WinLoss field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StraightBetV3) GetWinLossOk() (*float64, bool) {
	if o == nil || o.WinLoss == nil {
		return nil, false
	}
	return o.WinLoss, true
}

// HasWinLoss returns a boolean if a field has been set.
func (o *StraightBetV3) HasWinLoss() bool {
	if o != nil && o.WinLoss != nil {
		return true
	}

	return false
}

// SetWinLoss gets a reference to the given float64 and assigns it to the WinLoss field.
func (o *StraightBetV3) SetWinLoss(v float64) {
	o.WinLoss = &v
}

// GetOddsFormat returns the OddsFormat field value
func (o *StraightBetV3) GetOddsFormat() OddsFormat {
	if o == nil  {
		var ret OddsFormat
		return ret
	}

	return o.OddsFormat
}

// GetOddsFormatOk returns a tuple with the OddsFormat field value
// and a boolean to check if the value has been set.
func (o *StraightBetV3) GetOddsFormatOk() (*OddsFormat, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OddsFormat, true
}

// SetOddsFormat sets field value
func (o *StraightBetV3) SetOddsFormat(v OddsFormat) {
	o.OddsFormat = v
}

// GetCustomerCommission returns the CustomerCommission field value if set, zero value otherwise.
func (o *StraightBetV3) GetCustomerCommission() float64 {
	if o == nil || o.CustomerCommission == nil {
		var ret float64
		return ret
	}
	return *o.CustomerCommission
}

// GetCustomerCommissionOk returns a tuple with the CustomerCommission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StraightBetV3) GetCustomerCommissionOk() (*float64, bool) {
	if o == nil || o.CustomerCommission == nil {
		return nil, false
	}
	return o.CustomerCommission, true
}

// HasCustomerCommission returns a boolean if a field has been set.
func (o *StraightBetV3) HasCustomerCommission() bool {
	if o != nil && o.CustomerCommission != nil {
		return true
	}

	return false
}

// SetCustomerCommission gets a reference to the given float64 and assigns it to the CustomerCommission field.
func (o *StraightBetV3) SetCustomerCommission(v float64) {
	o.CustomerCommission = &v
}

// GetCancellationReason returns the CancellationReason field value if set, zero value otherwise.
func (o *StraightBetV3) GetCancellationReason() string {
	if o == nil || o.CancellationReason == nil {
		var ret string
		return ret
	}
	return *o.CancellationReason
}

// GetCancellationReasonOk returns a tuple with the CancellationReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StraightBetV3) GetCancellationReasonOk() (*string, bool) {
	if o == nil || o.CancellationReason == nil {
		return nil, false
	}
	return o.CancellationReason, true
}

// HasCancellationReason returns a boolean if a field has been set.
func (o *StraightBetV3) HasCancellationReason() bool {
	if o != nil && o.CancellationReason != nil {
		return true
	}

	return false
}

// SetCancellationReason gets a reference to the given string and assigns it to the CancellationReason field.
func (o *StraightBetV3) SetCancellationReason(v string) {
	o.CancellationReason = &v
}

// GetUpdateSequence returns the UpdateSequence field value
func (o *StraightBetV3) GetUpdateSequence() int64 {
	if o == nil  {
		var ret int64
		return ret
	}

	return o.UpdateSequence
}

// GetUpdateSequenceOk returns a tuple with the UpdateSequence field value
// and a boolean to check if the value has been set.
func (o *StraightBetV3) GetUpdateSequenceOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UpdateSequence, true
}

// SetUpdateSequence sets field value
func (o *StraightBetV3) SetUpdateSequence(v int64) {
	o.UpdateSequence = v
}

// GetSportId returns the SportId field value if set, zero value otherwise.
func (o *StraightBetV3) GetSportId() int {
	if o == nil || o.SportId == nil {
		var ret int
		return ret
	}
	return *o.SportId
}

// GetSportIdOk returns a tuple with the SportId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StraightBetV3) GetSportIdOk() (*int, bool) {
	if o == nil || o.SportId == nil {
		return nil, false
	}
	return o.SportId, true
}

// HasSportId returns a boolean if a field has been set.
func (o *StraightBetV3) HasSportId() bool {
	if o != nil && o.SportId != nil {
		return true
	}

	return false
}

// SetSportId gets a reference to the given int and assigns it to the SportId field.
func (o *StraightBetV3) SetSportId(v int) {
	o.SportId = &v
}

// GetLeagueId returns the LeagueId field value if set, zero value otherwise.
func (o *StraightBetV3) GetLeagueId() int {
	if o == nil || o.LeagueId == nil {
		var ret int
		return ret
	}
	return *o.LeagueId
}

// GetLeagueIdOk returns a tuple with the LeagueId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StraightBetV3) GetLeagueIdOk() (*int, bool) {
	if o == nil || o.LeagueId == nil {
		return nil, false
	}
	return o.LeagueId, true
}

// HasLeagueId returns a boolean if a field has been set.
func (o *StraightBetV3) HasLeagueId() bool {
	if o != nil && o.LeagueId != nil {
		return true
	}

	return false
}

// SetLeagueId gets a reference to the given int and assigns it to the LeagueId field.
func (o *StraightBetV3) SetLeagueId(v int) {
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

// GetEventIdOk returns a tuple with the EventId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StraightBetV3) GetEventIdOk() (*int64, bool) {
	if o == nil || o.EventId == nil {
		return nil, false
	}
	return o.EventId, true
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
func (o *StraightBetV3) GetHandicap() float64 {
	if o == nil || o.Handicap == nil {
		var ret float64
		return ret
	}
	return *o.Handicap
}

// GetHandicapOk returns a tuple with the Handicap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StraightBetV3) GetHandicapOk() (*float64, bool) {
	if o == nil || o.Handicap == nil {
		return nil, false
	}
	return o.Handicap, true
}

// HasHandicap returns a boolean if a field has been set.
func (o *StraightBetV3) HasHandicap() bool {
	if o != nil && o.Handicap != nil {
		return true
	}

	return false
}

// SetHandicap gets a reference to the given float64 and assigns it to the Handicap field.
func (o *StraightBetV3) SetHandicap(v float64) {
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

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StraightBetV3) GetPriceOk() (*float64, bool) {
	if o == nil || o.Price == nil {
		return nil, false
	}
	return o.Price, true
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

// GetTeamNameOk returns a tuple with the TeamName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StraightBetV3) GetTeamNameOk() (*string, bool) {
	if o == nil || o.TeamName == nil {
		return nil, false
	}
	return o.TeamName, true
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
func (o *StraightBetV3) GetSide() string {
	if o == nil || o.Side == nil {
		var ret string
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StraightBetV3) GetSideOk() (*string, bool) {
	if o == nil || o.Side == nil {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *StraightBetV3) HasSide() bool {
	if o != nil && o.Side != nil {
		return true
	}

	return false
}

// SetSide gets a reference to the given string and assigns it to the Side field.
func (o *StraightBetV3) SetSide(v string) {
	o.Side = &v
}

// GetPitcher1 returns the Pitcher1 field value if set, zero value otherwise.
func (o *StraightBetV3) GetPitcher1() string {
	if o == nil || o.Pitcher1 == nil {
		var ret string
		return ret
	}
	return *o.Pitcher1
}

// GetPitcher1Ok returns a tuple with the Pitcher1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StraightBetV3) GetPitcher1Ok() (*string, bool) {
	if o == nil || o.Pitcher1 == nil {
		return nil, false
	}
	return o.Pitcher1, true
}

// HasPitcher1 returns a boolean if a field has been set.
func (o *StraightBetV3) HasPitcher1() bool {
	if o != nil && o.Pitcher1 != nil {
		return true
	}

	return false
}

// SetPitcher1 gets a reference to the given string and assigns it to the Pitcher1 field.
func (o *StraightBetV3) SetPitcher1(v string) {
	o.Pitcher1 = &v
}

// GetPitcher2 returns the Pitcher2 field value if set, zero value otherwise.
func (o *StraightBetV3) GetPitcher2() string {
	if o == nil || o.Pitcher2 == nil {
		var ret string
		return ret
	}
	return *o.Pitcher2
}

// GetPitcher2Ok returns a tuple with the Pitcher2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StraightBetV3) GetPitcher2Ok() (*string, bool) {
	if o == nil || o.Pitcher2 == nil {
		return nil, false
	}
	return o.Pitcher2, true
}

// HasPitcher2 returns a boolean if a field has been set.
func (o *StraightBetV3) HasPitcher2() bool {
	if o != nil && o.Pitcher2 != nil {
		return true
	}

	return false
}

// SetPitcher2 gets a reference to the given string and assigns it to the Pitcher2 field.
func (o *StraightBetV3) SetPitcher2(v string) {
	o.Pitcher2 = &v
}

// GetPitcher1MustStart returns the Pitcher1MustStart field value if set, zero value otherwise.
func (o *StraightBetV3) GetPitcher1MustStart() bool {
	if o == nil || o.Pitcher1MustStart == nil {
		var ret bool
		return ret
	}
	return *o.Pitcher1MustStart
}

// GetPitcher1MustStartOk returns a tuple with the Pitcher1MustStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StraightBetV3) GetPitcher1MustStartOk() (*bool, bool) {
	if o == nil || o.Pitcher1MustStart == nil {
		return nil, false
	}
	return o.Pitcher1MustStart, true
}

// HasPitcher1MustStart returns a boolean if a field has been set.
func (o *StraightBetV3) HasPitcher1MustStart() bool {
	if o != nil && o.Pitcher1MustStart != nil {
		return true
	}

	return false
}

// SetPitcher1MustStart gets a reference to the given bool and assigns it to the Pitcher1MustStart field.
func (o *StraightBetV3) SetPitcher1MustStart(v bool) {
	o.Pitcher1MustStart = &v
}

// GetPitcher2MustStart returns the Pitcher2MustStart field value if set, zero value otherwise.
func (o *StraightBetV3) GetPitcher2MustStart() bool {
	if o == nil || o.Pitcher2MustStart == nil {
		var ret bool
		return ret
	}
	return *o.Pitcher2MustStart
}

// GetPitcher2MustStartOk returns a tuple with the Pitcher2MustStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StraightBetV3) GetPitcher2MustStartOk() (*bool, bool) {
	if o == nil || o.Pitcher2MustStart == nil {
		return nil, false
	}
	return o.Pitcher2MustStart, true
}

// HasPitcher2MustStart returns a boolean if a field has been set.
func (o *StraightBetV3) HasPitcher2MustStart() bool {
	if o != nil && o.Pitcher2MustStart != nil {
		return true
	}

	return false
}

// SetPitcher2MustStart gets a reference to the given bool and assigns it to the Pitcher2MustStart field.
func (o *StraightBetV3) SetPitcher2MustStart(v bool) {
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

// GetTeam1Ok returns a tuple with the Team1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StraightBetV3) GetTeam1Ok() (*string, bool) {
	if o == nil || o.Team1 == nil {
		return nil, false
	}
	return o.Team1, true
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

// GetTeam2Ok returns a tuple with the Team2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StraightBetV3) GetTeam2Ok() (*string, bool) {
	if o == nil || o.Team2 == nil {
		return nil, false
	}
	return o.Team2, true
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
func (o *StraightBetV3) GetPeriodNumber() int {
	if o == nil || o.PeriodNumber == nil {
		var ret int
		return ret
	}
	return *o.PeriodNumber
}

// GetPeriodNumberOk returns a tuple with the PeriodNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StraightBetV3) GetPeriodNumberOk() (*int, bool) {
	if o == nil || o.PeriodNumber == nil {
		return nil, false
	}
	return o.PeriodNumber, true
}

// HasPeriodNumber returns a boolean if a field has been set.
func (o *StraightBetV3) HasPeriodNumber() bool {
	if o != nil && o.PeriodNumber != nil {
		return true
	}

	return false
}

// SetPeriodNumber gets a reference to the given int and assigns it to the PeriodNumber field.
func (o *StraightBetV3) SetPeriodNumber(v int) {
	o.PeriodNumber = &v
}

// GetTeam1Score returns the Team1Score field value if set, zero value otherwise.
func (o *StraightBetV3) GetTeam1Score() float64 {
	if o == nil || o.Team1Score == nil {
		var ret float64
		return ret
	}
	return *o.Team1Score
}

// GetTeam1ScoreOk returns a tuple with the Team1Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StraightBetV3) GetTeam1ScoreOk() (*float64, bool) {
	if o == nil || o.Team1Score == nil {
		return nil, false
	}
	return o.Team1Score, true
}

// HasTeam1Score returns a boolean if a field has been set.
func (o *StraightBetV3) HasTeam1Score() bool {
	if o != nil && o.Team1Score != nil {
		return true
	}

	return false
}

// SetTeam1Score gets a reference to the given float64 and assigns it to the Team1Score field.
func (o *StraightBetV3) SetTeam1Score(v float64) {
	o.Team1Score = &v
}

// GetTeam2Score returns the Team2Score field value if set, zero value otherwise.
func (o *StraightBetV3) GetTeam2Score() float64 {
	if o == nil || o.Team2Score == nil {
		var ret float64
		return ret
	}
	return *o.Team2Score
}

// GetTeam2ScoreOk returns a tuple with the Team2Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StraightBetV3) GetTeam2ScoreOk() (*float64, bool) {
	if o == nil || o.Team2Score == nil {
		return nil, false
	}
	return o.Team2Score, true
}

// HasTeam2Score returns a boolean if a field has been set.
func (o *StraightBetV3) HasTeam2Score() bool {
	if o != nil && o.Team2Score != nil {
		return true
	}

	return false
}

// SetTeam2Score gets a reference to the given float64 and assigns it to the Team2Score field.
func (o *StraightBetV3) SetTeam2Score(v float64) {
	o.Team2Score = &v
}

// GetFtTeam1Score returns the FtTeam1Score field value if set, zero value otherwise.
func (o *StraightBetV3) GetFtTeam1Score() float64 {
	if o == nil || o.FtTeam1Score == nil {
		var ret float64
		return ret
	}
	return *o.FtTeam1Score
}

// GetFtTeam1ScoreOk returns a tuple with the FtTeam1Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StraightBetV3) GetFtTeam1ScoreOk() (*float64, bool) {
	if o == nil || o.FtTeam1Score == nil {
		return nil, false
	}
	return o.FtTeam1Score, true
}

// HasFtTeam1Score returns a boolean if a field has been set.
func (o *StraightBetV3) HasFtTeam1Score() bool {
	if o != nil && o.FtTeam1Score != nil {
		return true
	}

	return false
}

// SetFtTeam1Score gets a reference to the given float64 and assigns it to the FtTeam1Score field.
func (o *StraightBetV3) SetFtTeam1Score(v float64) {
	o.FtTeam1Score = &v
}

// GetFtTeam2Score returns the FtTeam2Score field value if set, zero value otherwise.
func (o *StraightBetV3) GetFtTeam2Score() float64 {
	if o == nil || o.FtTeam2Score == nil {
		var ret float64
		return ret
	}
	return *o.FtTeam2Score
}

// GetFtTeam2ScoreOk returns a tuple with the FtTeam2Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StraightBetV3) GetFtTeam2ScoreOk() (*float64, bool) {
	if o == nil || o.FtTeam2Score == nil {
		return nil, false
	}
	return o.FtTeam2Score, true
}

// HasFtTeam2Score returns a boolean if a field has been set.
func (o *StraightBetV3) HasFtTeam2Score() bool {
	if o != nil && o.FtTeam2Score != nil {
		return true
	}

	return false
}

// SetFtTeam2Score gets a reference to the given float64 and assigns it to the FtTeam2Score field.
func (o *StraightBetV3) SetFtTeam2Score(v float64) {
	o.FtTeam2Score = &v
}

// GetPTeam1Score returns the PTeam1Score field value if set, zero value otherwise.
func (o *StraightBetV3) GetPTeam1Score() float64 {
	if o == nil || o.PTeam1Score == nil {
		var ret float64
		return ret
	}
	return *o.PTeam1Score
}

// GetPTeam1ScoreOk returns a tuple with the PTeam1Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StraightBetV3) GetPTeam1ScoreOk() (*float64, bool) {
	if o == nil || o.PTeam1Score == nil {
		return nil, false
	}
	return o.PTeam1Score, true
}

// HasPTeam1Score returns a boolean if a field has been set.
func (o *StraightBetV3) HasPTeam1Score() bool {
	if o != nil && o.PTeam1Score != nil {
		return true
	}

	return false
}

// SetPTeam1Score gets a reference to the given float64 and assigns it to the PTeam1Score field.
func (o *StraightBetV3) SetPTeam1Score(v float64) {
	o.PTeam1Score = &v
}

// GetPTeam2Score returns the PTeam2Score field value if set, zero value otherwise.
func (o *StraightBetV3) GetPTeam2Score() float64 {
	if o == nil || o.PTeam2Score == nil {
		var ret float64
		return ret
	}
	return *o.PTeam2Score
}

// GetPTeam2ScoreOk returns a tuple with the PTeam2Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StraightBetV3) GetPTeam2ScoreOk() (*float64, bool) {
	if o == nil || o.PTeam2Score == nil {
		return nil, false
	}
	return o.PTeam2Score, true
}

// HasPTeam2Score returns a boolean if a field has been set.
func (o *StraightBetV3) HasPTeam2Score() bool {
	if o != nil && o.PTeam2Score != nil {
		return true
	}

	return false
}

// SetPTeam2Score gets a reference to the given float64 and assigns it to the PTeam2Score field.
func (o *StraightBetV3) SetPTeam2Score(v float64) {
	o.PTeam2Score = &v
}

// GetEventStartTime returns the EventStartTime field value if set, zero value otherwise.
func (o *StraightBetV3) GetEventStartTime() time.Time {
	if o == nil || o.EventStartTime == nil {
		var ret time.Time
		return ret
	}
	return *o.EventStartTime
}

// GetEventStartTimeOk returns a tuple with the EventStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StraightBetV3) GetEventStartTimeOk() (*time.Time, bool) {
	if o == nil || o.EventStartTime == nil {
		return nil, false
	}
	return o.EventStartTime, true
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

func (o StraightBetV3) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["betId"] = o.BetId
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
	if o.WinLoss != nil {
		toSerialize["winLoss"] = o.WinLoss
	}
	if true {
		toSerialize["oddsFormat"] = o.OddsFormat
	}
	if o.CustomerCommission != nil {
		toSerialize["customerCommission"] = o.CustomerCommission
	}
	if o.CancellationReason != nil {
		toSerialize["cancellationReason"] = o.CancellationReason
	}
	if true {
		toSerialize["updateSequence"] = o.UpdateSequence
	}
	if o.SportId != nil {
		toSerialize["sportId"] = o.SportId
	}
	if o.LeagueId != nil {
		toSerialize["leagueId"] = o.LeagueId
	}
	if o.EventId != nil {
		toSerialize["eventId"] = o.EventId
	}
	if o.Handicap != nil {
		toSerialize["handicap"] = o.Handicap
	}
	if o.Price != nil {
		toSerialize["price"] = o.Price
	}
	if o.TeamName != nil {
		toSerialize["teamName"] = o.TeamName
	}
	if o.Side != nil {
		toSerialize["side"] = o.Side
	}
	if o.Pitcher1 != nil {
		toSerialize["pitcher1"] = o.Pitcher1
	}
	if o.Pitcher2 != nil {
		toSerialize["pitcher2"] = o.Pitcher2
	}
	if o.Pitcher1MustStart != nil {
		toSerialize["pitcher1MustStart"] = o.Pitcher1MustStart
	}
	if o.Pitcher2MustStart != nil {
		toSerialize["pitcher2MustStart"] = o.Pitcher2MustStart
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
	if o.Team1Score != nil {
		toSerialize["team1Score"] = o.Team1Score
	}
	if o.Team2Score != nil {
		toSerialize["team2Score"] = o.Team2Score
	}
	if o.FtTeam1Score != nil {
		toSerialize["ftTeam1Score"] = o.FtTeam1Score
	}
	if o.FtTeam2Score != nil {
		toSerialize["ftTeam2Score"] = o.FtTeam2Score
	}
	if o.PTeam1Score != nil {
		toSerialize["pTeam1Score"] = o.PTeam1Score
	}
	if o.PTeam2Score != nil {
		toSerialize["pTeam2Score"] = o.PTeam2Score
	}
	if o.EventStartTime != nil {
		toSerialize["eventStartTime"] = o.EventStartTime
	}
	return json.Marshal(toSerialize)
}

type NullableStraightBetV3 struct {
	value *StraightBetV3
	isSet bool
}

func (v NullableStraightBetV3) Get() *StraightBetV3 {
	return v.value
}

func (v *NullableStraightBetV3) Set(val *StraightBetV3) {
	v.value = val
	v.isSet = true
}

func (v NullableStraightBetV3) IsSet() bool {
	return v.isSet
}

func (v *NullableStraightBetV3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStraightBetV3(val *StraightBetV3) *NullableStraightBetV3 {
	return &NullableStraightBetV3{value: val, isSet: true}
}

func (v NullableStraightBetV3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStraightBetV3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
