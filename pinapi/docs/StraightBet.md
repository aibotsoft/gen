# StraightBet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BetId** | Pointer to **int64** | Bet identification | 
**WagerNumber** | Pointer to **int32** | Wager identification. All bets placed thru the API will have value 1. Website Classic view supports multiple contest(special) bets placement in the same bet slip in that case the bet would have appropriate wager number, as well as all round robin parlay bets. | 
**PlacedAt** | Pointer to [**time.Time**](time.Time.md) | Date time when the bet was placed. | 
**BetStatus** | Pointer to **string** | Bet Status.    ACCEPTED &#x3D; Bet was accepted,   CANCELLED &#x3D; Bet is cancelled as per Pinnacle betting rules,   LOSE &#x3D; The bet is settled as lose,   PENDING_ACCEPTANCE &#x3D; This status is reserved only for live bets. If a live bet is placed during danger zone or live delay is applied, it will be in PENDING_ACCEPTANCE , otherwise in ACCEPTED status. From this status bet can go to ACCEPTED or NOT_ACCEPTED status,   REFUNDED &#x3D; When an event is cancelled or when the bet is settled as push, the bet will have REFUNDED status,   NOT_ACCEPTED &#x3D; Bet was not accepted. Bet can be in this status only if it was previously in PENDING_ACCEPTANCE status,   WON &#x3D; The bet is settled as won   | 
**BetType** | Pointer to **string** | Bet type. | 
**Win** | Pointer to **float64** | Win amount. | 
**Risk** | Pointer to **float64** | Risk amount. | 
**WinLoss** | Pointer to **NullableFloat64** | Win-Loss for settled bets. | [optional] 
**OddsFormat** | Pointer to [**OddsFormat**](OddsFormat.md) |  | 
**CustomerCommission** | Pointer to **NullableFloat64** | Client’s commission on the bet. | [optional] 
**CancellationReason** | Pointer to [**CancellationReason**](CancellationReason.md) |  | [optional] 
**UpdateSequence** | Pointer to **int64** | Update Sequence | 
**SportId** | Pointer to **int32** |  | [optional] 
**LeagueId** | Pointer to **int32** |  | [optional] 
**EventId** | Pointer to **int64** |  | [optional] 
**Handicap** | Pointer to **NullableFloat64** |  | [optional] 
**Price** | Pointer to **float64** |  | [optional] 
**TeamName** | Pointer to **string** |  | [optional] 
**Side** | Pointer to **NullableString** | Side type. | [optional] 
**Pitcher1** | Pointer to **NullableString** | Pitcher name of team1. Only for bets on baseball. | [optional] 
**Pitcher2** | Pointer to **NullableString** | Pitcher name of team2. Only for bets on baseball. | [optional] 
**Pitcher1MustStart** | Pointer to **NullableString** | Whether the team1 pitcher must start. Only for bets on baseball. | [optional] 
**Pitcher2MustStart** | Pointer to **NullableString** | Whether the team1 pitcher must start. Only for bets on baseball. | [optional] 
**Team1** | Pointer to **string** |  | [optional] 
**Team2** | Pointer to **string** |  | [optional] 
**PeriodNumber** | Pointer to **int32** |  | [optional] 
**Team1Score** | Pointer to **NullableFloat64** | Team 1 score that the bet was placed on, only for live bets. | [optional] 
**Team2Score** | Pointer to **NullableFloat64** | Team 2 score that the bet was placed, only for live bets. | [optional] 
**FtTeam1Score** | Pointer to **NullableFloat64** | Full time team 1 score, only for settled bets. | [optional] 
**FtTeam2Score** | Pointer to **NullableFloat64** | Full time team 2 score, only for settled bets. | [optional] 
**PTeam1Score** | Pointer to **NullableFloat64** | .End of period team 1 score, only for settled bets. If the bet was placed on Game period (periodNumber &#x3D;0) , this will be null .  | [optional] 
**PTeam2Score** | Pointer to **NullableFloat64** | End of period team 2 score, only for settled bets. If the bet was placed on Game period (periodNumber &#x3D;0), this will be null | [optional] 
**IsLive** | Pointer to **string** | Whether the bet is on live event | [optional] 

## Methods

### GetBetId

`func (o *StraightBet) GetBetId() int64`

GetBetId returns the BetId field if non-nil, zero value otherwise.

### GetBetIdOk

`func (o *StraightBet) GetBetIdOk() (int64, bool)`

GetBetIdOk returns a tuple with the BetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBetId

`func (o *StraightBet) HasBetId() bool`

HasBetId returns a boolean if a field has been set.

### SetBetId

`func (o *StraightBet) SetBetId(v int64)`

SetBetId gets a reference to the given int64 and assigns it to the BetId field.

### GetWagerNumber

`func (o *StraightBet) GetWagerNumber() int32`

GetWagerNumber returns the WagerNumber field if non-nil, zero value otherwise.

### GetWagerNumberOk

`func (o *StraightBet) GetWagerNumberOk() (int32, bool)`

GetWagerNumberOk returns a tuple with the WagerNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWagerNumber

`func (o *StraightBet) HasWagerNumber() bool`

HasWagerNumber returns a boolean if a field has been set.

### SetWagerNumber

`func (o *StraightBet) SetWagerNumber(v int32)`

SetWagerNumber gets a reference to the given int32 and assigns it to the WagerNumber field.

### GetPlacedAt

`func (o *StraightBet) GetPlacedAt() time.Time`

GetPlacedAt returns the PlacedAt field if non-nil, zero value otherwise.

### GetPlacedAtOk

`func (o *StraightBet) GetPlacedAtOk() (time.Time, bool)`

GetPlacedAtOk returns a tuple with the PlacedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPlacedAt

`func (o *StraightBet) HasPlacedAt() bool`

HasPlacedAt returns a boolean if a field has been set.

### SetPlacedAt

`func (o *StraightBet) SetPlacedAt(v time.Time)`

SetPlacedAt gets a reference to the given time.Time and assigns it to the PlacedAt field.

### GetBetStatus

`func (o *StraightBet) GetBetStatus() string`

GetBetStatus returns the BetStatus field if non-nil, zero value otherwise.

### GetBetStatusOk

`func (o *StraightBet) GetBetStatusOk() (string, bool)`

GetBetStatusOk returns a tuple with the BetStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBetStatus

`func (o *StraightBet) HasBetStatus() bool`

HasBetStatus returns a boolean if a field has been set.

### SetBetStatus

`func (o *StraightBet) SetBetStatus(v string)`

SetBetStatus gets a reference to the given string and assigns it to the BetStatus field.

### GetBetType

`func (o *StraightBet) GetBetType() string`

GetBetType returns the BetType field if non-nil, zero value otherwise.

### GetBetTypeOk

`func (o *StraightBet) GetBetTypeOk() (string, bool)`

GetBetTypeOk returns a tuple with the BetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBetType

`func (o *StraightBet) HasBetType() bool`

HasBetType returns a boolean if a field has been set.

### SetBetType

`func (o *StraightBet) SetBetType(v string)`

SetBetType gets a reference to the given string and assigns it to the BetType field.

### GetWin

`func (o *StraightBet) GetWin() float64`

GetWin returns the Win field if non-nil, zero value otherwise.

### GetWinOk

`func (o *StraightBet) GetWinOk() (float64, bool)`

GetWinOk returns a tuple with the Win field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWin

`func (o *StraightBet) HasWin() bool`

HasWin returns a boolean if a field has been set.

### SetWin

`func (o *StraightBet) SetWin(v float64)`

SetWin gets a reference to the given float64 and assigns it to the Win field.

### GetRisk

`func (o *StraightBet) GetRisk() float64`

GetRisk returns the Risk field if non-nil, zero value otherwise.

### GetRiskOk

`func (o *StraightBet) GetRiskOk() (float64, bool)`

GetRiskOk returns a tuple with the Risk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRisk

`func (o *StraightBet) HasRisk() bool`

HasRisk returns a boolean if a field has been set.

### SetRisk

`func (o *StraightBet) SetRisk(v float64)`

SetRisk gets a reference to the given float64 and assigns it to the Risk field.

### GetWinLoss

`func (o *StraightBet) GetWinLoss() NullableFloat64`

GetWinLoss returns the WinLoss field if non-nil, zero value otherwise.

### GetWinLossOk

`func (o *StraightBet) GetWinLossOk() (NullableFloat64, bool)`

GetWinLossOk returns a tuple with the WinLoss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWinLoss

`func (o *StraightBet) HasWinLoss() bool`

HasWinLoss returns a boolean if a field has been set.

### SetWinLoss

`func (o *StraightBet) SetWinLoss(v NullableFloat64)`

SetWinLoss gets a reference to the given NullableFloat64 and assigns it to the WinLoss field.

### SetWinLossExplicitNull

`func (o *StraightBet) SetWinLossExplicitNull(b bool)`

SetWinLossExplicitNull (un)sets WinLoss to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The WinLoss value is set to nil even if false is passed
### GetOddsFormat

`func (o *StraightBet) GetOddsFormat() OddsFormat`

GetOddsFormat returns the OddsFormat field if non-nil, zero value otherwise.

### GetOddsFormatOk

`func (o *StraightBet) GetOddsFormatOk() (OddsFormat, bool)`

GetOddsFormatOk returns a tuple with the OddsFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOddsFormat

`func (o *StraightBet) HasOddsFormat() bool`

HasOddsFormat returns a boolean if a field has been set.

### SetOddsFormat

`func (o *StraightBet) SetOddsFormat(v OddsFormat)`

SetOddsFormat gets a reference to the given OddsFormat and assigns it to the OddsFormat field.

### GetCustomerCommission

`func (o *StraightBet) GetCustomerCommission() NullableFloat64`

GetCustomerCommission returns the CustomerCommission field if non-nil, zero value otherwise.

### GetCustomerCommissionOk

`func (o *StraightBet) GetCustomerCommissionOk() (NullableFloat64, bool)`

GetCustomerCommissionOk returns a tuple with the CustomerCommission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCustomerCommission

`func (o *StraightBet) HasCustomerCommission() bool`

HasCustomerCommission returns a boolean if a field has been set.

### SetCustomerCommission

`func (o *StraightBet) SetCustomerCommission(v NullableFloat64)`

SetCustomerCommission gets a reference to the given NullableFloat64 and assigns it to the CustomerCommission field.

### SetCustomerCommissionExplicitNull

`func (o *StraightBet) SetCustomerCommissionExplicitNull(b bool)`

SetCustomerCommissionExplicitNull (un)sets CustomerCommission to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CustomerCommission value is set to nil even if false is passed
### GetCancellationReason

`func (o *StraightBet) GetCancellationReason() CancellationReason`

GetCancellationReason returns the CancellationReason field if non-nil, zero value otherwise.

### GetCancellationReasonOk

`func (o *StraightBet) GetCancellationReasonOk() (CancellationReason, bool)`

GetCancellationReasonOk returns a tuple with the CancellationReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCancellationReason

`func (o *StraightBet) HasCancellationReason() bool`

HasCancellationReason returns a boolean if a field has been set.

### SetCancellationReason

`func (o *StraightBet) SetCancellationReason(v CancellationReason)`

SetCancellationReason gets a reference to the given CancellationReason and assigns it to the CancellationReason field.

### GetUpdateSequence

`func (o *StraightBet) GetUpdateSequence() int64`

GetUpdateSequence returns the UpdateSequence field if non-nil, zero value otherwise.

### GetUpdateSequenceOk

`func (o *StraightBet) GetUpdateSequenceOk() (int64, bool)`

GetUpdateSequenceOk returns a tuple with the UpdateSequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUpdateSequence

`func (o *StraightBet) HasUpdateSequence() bool`

HasUpdateSequence returns a boolean if a field has been set.

### SetUpdateSequence

`func (o *StraightBet) SetUpdateSequence(v int64)`

SetUpdateSequence gets a reference to the given int64 and assigns it to the UpdateSequence field.

### GetSportId

`func (o *StraightBet) GetSportId() int32`

GetSportId returns the SportId field if non-nil, zero value otherwise.

### GetSportIdOk

`func (o *StraightBet) GetSportIdOk() (int32, bool)`

GetSportIdOk returns a tuple with the SportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSportId

`func (o *StraightBet) HasSportId() bool`

HasSportId returns a boolean if a field has been set.

### SetSportId

`func (o *StraightBet) SetSportId(v int32)`

SetSportId gets a reference to the given int32 and assigns it to the SportId field.

### GetLeagueId

`func (o *StraightBet) GetLeagueId() int32`

GetLeagueId returns the LeagueId field if non-nil, zero value otherwise.

### GetLeagueIdOk

`func (o *StraightBet) GetLeagueIdOk() (int32, bool)`

GetLeagueIdOk returns a tuple with the LeagueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLeagueId

`func (o *StraightBet) HasLeagueId() bool`

HasLeagueId returns a boolean if a field has been set.

### SetLeagueId

`func (o *StraightBet) SetLeagueId(v int32)`

SetLeagueId gets a reference to the given int32 and assigns it to the LeagueId field.

### GetEventId

`func (o *StraightBet) GetEventId() int64`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *StraightBet) GetEventIdOk() (int64, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEventId

`func (o *StraightBet) HasEventId() bool`

HasEventId returns a boolean if a field has been set.

### SetEventId

`func (o *StraightBet) SetEventId(v int64)`

SetEventId gets a reference to the given int64 and assigns it to the EventId field.

### GetHandicap

`func (o *StraightBet) GetHandicap() NullableFloat64`

GetHandicap returns the Handicap field if non-nil, zero value otherwise.

### GetHandicapOk

`func (o *StraightBet) GetHandicapOk() (NullableFloat64, bool)`

GetHandicapOk returns a tuple with the Handicap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHandicap

`func (o *StraightBet) HasHandicap() bool`

HasHandicap returns a boolean if a field has been set.

### SetHandicap

`func (o *StraightBet) SetHandicap(v NullableFloat64)`

SetHandicap gets a reference to the given NullableFloat64 and assigns it to the Handicap field.

### SetHandicapExplicitNull

`func (o *StraightBet) SetHandicapExplicitNull(b bool)`

SetHandicapExplicitNull (un)sets Handicap to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Handicap value is set to nil even if false is passed
### GetPrice

`func (o *StraightBet) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *StraightBet) GetPriceOk() (float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPrice

`func (o *StraightBet) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPrice

`func (o *StraightBet) SetPrice(v float64)`

SetPrice gets a reference to the given float64 and assigns it to the Price field.

### GetTeamName

`func (o *StraightBet) GetTeamName() string`

GetTeamName returns the TeamName field if non-nil, zero value otherwise.

### GetTeamNameOk

`func (o *StraightBet) GetTeamNameOk() (string, bool)`

GetTeamNameOk returns a tuple with the TeamName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTeamName

`func (o *StraightBet) HasTeamName() bool`

HasTeamName returns a boolean if a field has been set.

### SetTeamName

`func (o *StraightBet) SetTeamName(v string)`

SetTeamName gets a reference to the given string and assigns it to the TeamName field.

### GetSide

`func (o *StraightBet) GetSide() NullableString`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *StraightBet) GetSideOk() (NullableString, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSide

`func (o *StraightBet) HasSide() bool`

HasSide returns a boolean if a field has been set.

### SetSide

`func (o *StraightBet) SetSide(v NullableString)`

SetSide gets a reference to the given NullableString and assigns it to the Side field.

### SetSideExplicitNull

`func (o *StraightBet) SetSideExplicitNull(b bool)`

SetSideExplicitNull (un)sets Side to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Side value is set to nil even if false is passed
### GetPitcher1

`func (o *StraightBet) GetPitcher1() NullableString`

GetPitcher1 returns the Pitcher1 field if non-nil, zero value otherwise.

### GetPitcher1Ok

`func (o *StraightBet) GetPitcher1Ok() (NullableString, bool)`

GetPitcher1Ok returns a tuple with the Pitcher1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPitcher1

`func (o *StraightBet) HasPitcher1() bool`

HasPitcher1 returns a boolean if a field has been set.

### SetPitcher1

`func (o *StraightBet) SetPitcher1(v NullableString)`

SetPitcher1 gets a reference to the given NullableString and assigns it to the Pitcher1 field.

### SetPitcher1ExplicitNull

`func (o *StraightBet) SetPitcher1ExplicitNull(b bool)`

SetPitcher1ExplicitNull (un)sets Pitcher1 to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Pitcher1 value is set to nil even if false is passed
### GetPitcher2

`func (o *StraightBet) GetPitcher2() NullableString`

GetPitcher2 returns the Pitcher2 field if non-nil, zero value otherwise.

### GetPitcher2Ok

`func (o *StraightBet) GetPitcher2Ok() (NullableString, bool)`

GetPitcher2Ok returns a tuple with the Pitcher2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPitcher2

`func (o *StraightBet) HasPitcher2() bool`

HasPitcher2 returns a boolean if a field has been set.

### SetPitcher2

`func (o *StraightBet) SetPitcher2(v NullableString)`

SetPitcher2 gets a reference to the given NullableString and assigns it to the Pitcher2 field.

### SetPitcher2ExplicitNull

`func (o *StraightBet) SetPitcher2ExplicitNull(b bool)`

SetPitcher2ExplicitNull (un)sets Pitcher2 to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Pitcher2 value is set to nil even if false is passed
### GetPitcher1MustStart

`func (o *StraightBet) GetPitcher1MustStart() NullableString`

GetPitcher1MustStart returns the Pitcher1MustStart field if non-nil, zero value otherwise.

### GetPitcher1MustStartOk

`func (o *StraightBet) GetPitcher1MustStartOk() (NullableString, bool)`

GetPitcher1MustStartOk returns a tuple with the Pitcher1MustStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPitcher1MustStart

`func (o *StraightBet) HasPitcher1MustStart() bool`

HasPitcher1MustStart returns a boolean if a field has been set.

### SetPitcher1MustStart

`func (o *StraightBet) SetPitcher1MustStart(v NullableString)`

SetPitcher1MustStart gets a reference to the given NullableString and assigns it to the Pitcher1MustStart field.

### SetPitcher1MustStartExplicitNull

`func (o *StraightBet) SetPitcher1MustStartExplicitNull(b bool)`

SetPitcher1MustStartExplicitNull (un)sets Pitcher1MustStart to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Pitcher1MustStart value is set to nil even if false is passed
### GetPitcher2MustStart

`func (o *StraightBet) GetPitcher2MustStart() NullableString`

GetPitcher2MustStart returns the Pitcher2MustStart field if non-nil, zero value otherwise.

### GetPitcher2MustStartOk

`func (o *StraightBet) GetPitcher2MustStartOk() (NullableString, bool)`

GetPitcher2MustStartOk returns a tuple with the Pitcher2MustStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPitcher2MustStart

`func (o *StraightBet) HasPitcher2MustStart() bool`

HasPitcher2MustStart returns a boolean if a field has been set.

### SetPitcher2MustStart

`func (o *StraightBet) SetPitcher2MustStart(v NullableString)`

SetPitcher2MustStart gets a reference to the given NullableString and assigns it to the Pitcher2MustStart field.

### SetPitcher2MustStartExplicitNull

`func (o *StraightBet) SetPitcher2MustStartExplicitNull(b bool)`

SetPitcher2MustStartExplicitNull (un)sets Pitcher2MustStart to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Pitcher2MustStart value is set to nil even if false is passed
### GetTeam1

`func (o *StraightBet) GetTeam1() string`

GetTeam1 returns the Team1 field if non-nil, zero value otherwise.

### GetTeam1Ok

`func (o *StraightBet) GetTeam1Ok() (string, bool)`

GetTeam1Ok returns a tuple with the Team1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTeam1

`func (o *StraightBet) HasTeam1() bool`

HasTeam1 returns a boolean if a field has been set.

### SetTeam1

`func (o *StraightBet) SetTeam1(v string)`

SetTeam1 gets a reference to the given string and assigns it to the Team1 field.

### GetTeam2

`func (o *StraightBet) GetTeam2() string`

GetTeam2 returns the Team2 field if non-nil, zero value otherwise.

### GetTeam2Ok

`func (o *StraightBet) GetTeam2Ok() (string, bool)`

GetTeam2Ok returns a tuple with the Team2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTeam2

`func (o *StraightBet) HasTeam2() bool`

HasTeam2 returns a boolean if a field has been set.

### SetTeam2

`func (o *StraightBet) SetTeam2(v string)`

SetTeam2 gets a reference to the given string and assigns it to the Team2 field.

### GetPeriodNumber

`func (o *StraightBet) GetPeriodNumber() int32`

GetPeriodNumber returns the PeriodNumber field if non-nil, zero value otherwise.

### GetPeriodNumberOk

`func (o *StraightBet) GetPeriodNumberOk() (int32, bool)`

GetPeriodNumberOk returns a tuple with the PeriodNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPeriodNumber

`func (o *StraightBet) HasPeriodNumber() bool`

HasPeriodNumber returns a boolean if a field has been set.

### SetPeriodNumber

`func (o *StraightBet) SetPeriodNumber(v int32)`

SetPeriodNumber gets a reference to the given int32 and assigns it to the PeriodNumber field.

### GetTeam1Score

`func (o *StraightBet) GetTeam1Score() NullableFloat64`

GetTeam1Score returns the Team1Score field if non-nil, zero value otherwise.

### GetTeam1ScoreOk

`func (o *StraightBet) GetTeam1ScoreOk() (NullableFloat64, bool)`

GetTeam1ScoreOk returns a tuple with the Team1Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTeam1Score

`func (o *StraightBet) HasTeam1Score() bool`

HasTeam1Score returns a boolean if a field has been set.

### SetTeam1Score

`func (o *StraightBet) SetTeam1Score(v NullableFloat64)`

SetTeam1Score gets a reference to the given NullableFloat64 and assigns it to the Team1Score field.

### SetTeam1ScoreExplicitNull

`func (o *StraightBet) SetTeam1ScoreExplicitNull(b bool)`

SetTeam1ScoreExplicitNull (un)sets Team1Score to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Team1Score value is set to nil even if false is passed
### GetTeam2Score

`func (o *StraightBet) GetTeam2Score() NullableFloat64`

GetTeam2Score returns the Team2Score field if non-nil, zero value otherwise.

### GetTeam2ScoreOk

`func (o *StraightBet) GetTeam2ScoreOk() (NullableFloat64, bool)`

GetTeam2ScoreOk returns a tuple with the Team2Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTeam2Score

`func (o *StraightBet) HasTeam2Score() bool`

HasTeam2Score returns a boolean if a field has been set.

### SetTeam2Score

`func (o *StraightBet) SetTeam2Score(v NullableFloat64)`

SetTeam2Score gets a reference to the given NullableFloat64 and assigns it to the Team2Score field.

### SetTeam2ScoreExplicitNull

`func (o *StraightBet) SetTeam2ScoreExplicitNull(b bool)`

SetTeam2ScoreExplicitNull (un)sets Team2Score to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Team2Score value is set to nil even if false is passed
### GetFtTeam1Score

`func (o *StraightBet) GetFtTeam1Score() NullableFloat64`

GetFtTeam1Score returns the FtTeam1Score field if non-nil, zero value otherwise.

### GetFtTeam1ScoreOk

`func (o *StraightBet) GetFtTeam1ScoreOk() (NullableFloat64, bool)`

GetFtTeam1ScoreOk returns a tuple with the FtTeam1Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFtTeam1Score

`func (o *StraightBet) HasFtTeam1Score() bool`

HasFtTeam1Score returns a boolean if a field has been set.

### SetFtTeam1Score

`func (o *StraightBet) SetFtTeam1Score(v NullableFloat64)`

SetFtTeam1Score gets a reference to the given NullableFloat64 and assigns it to the FtTeam1Score field.

### SetFtTeam1ScoreExplicitNull

`func (o *StraightBet) SetFtTeam1ScoreExplicitNull(b bool)`

SetFtTeam1ScoreExplicitNull (un)sets FtTeam1Score to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The FtTeam1Score value is set to nil even if false is passed
### GetFtTeam2Score

`func (o *StraightBet) GetFtTeam2Score() NullableFloat64`

GetFtTeam2Score returns the FtTeam2Score field if non-nil, zero value otherwise.

### GetFtTeam2ScoreOk

`func (o *StraightBet) GetFtTeam2ScoreOk() (NullableFloat64, bool)`

GetFtTeam2ScoreOk returns a tuple with the FtTeam2Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFtTeam2Score

`func (o *StraightBet) HasFtTeam2Score() bool`

HasFtTeam2Score returns a boolean if a field has been set.

### SetFtTeam2Score

`func (o *StraightBet) SetFtTeam2Score(v NullableFloat64)`

SetFtTeam2Score gets a reference to the given NullableFloat64 and assigns it to the FtTeam2Score field.

### SetFtTeam2ScoreExplicitNull

`func (o *StraightBet) SetFtTeam2ScoreExplicitNull(b bool)`

SetFtTeam2ScoreExplicitNull (un)sets FtTeam2Score to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The FtTeam2Score value is set to nil even if false is passed
### GetPTeam1Score

`func (o *StraightBet) GetPTeam1Score() NullableFloat64`

GetPTeam1Score returns the PTeam1Score field if non-nil, zero value otherwise.

### GetPTeam1ScoreOk

`func (o *StraightBet) GetPTeam1ScoreOk() (NullableFloat64, bool)`

GetPTeam1ScoreOk returns a tuple with the PTeam1Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPTeam1Score

`func (o *StraightBet) HasPTeam1Score() bool`

HasPTeam1Score returns a boolean if a field has been set.

### SetPTeam1Score

`func (o *StraightBet) SetPTeam1Score(v NullableFloat64)`

SetPTeam1Score gets a reference to the given NullableFloat64 and assigns it to the PTeam1Score field.

### SetPTeam1ScoreExplicitNull

`func (o *StraightBet) SetPTeam1ScoreExplicitNull(b bool)`

SetPTeam1ScoreExplicitNull (un)sets PTeam1Score to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PTeam1Score value is set to nil even if false is passed
### GetPTeam2Score

`func (o *StraightBet) GetPTeam2Score() NullableFloat64`

GetPTeam2Score returns the PTeam2Score field if non-nil, zero value otherwise.

### GetPTeam2ScoreOk

`func (o *StraightBet) GetPTeam2ScoreOk() (NullableFloat64, bool)`

GetPTeam2ScoreOk returns a tuple with the PTeam2Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPTeam2Score

`func (o *StraightBet) HasPTeam2Score() bool`

HasPTeam2Score returns a boolean if a field has been set.

### SetPTeam2Score

`func (o *StraightBet) SetPTeam2Score(v NullableFloat64)`

SetPTeam2Score gets a reference to the given NullableFloat64 and assigns it to the PTeam2Score field.

### SetPTeam2ScoreExplicitNull

`func (o *StraightBet) SetPTeam2ScoreExplicitNull(b bool)`

SetPTeam2ScoreExplicitNull (un)sets PTeam2Score to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PTeam2Score value is set to nil even if false is passed
### GetIsLive

`func (o *StraightBet) GetIsLive() string`

GetIsLive returns the IsLive field if non-nil, zero value otherwise.

### GetIsLiveOk

`func (o *StraightBet) GetIsLiveOk() (string, bool)`

GetIsLiveOk returns a tuple with the IsLive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsLive

`func (o *StraightBet) HasIsLive() bool`

HasIsLive returns a boolean if a field has been set.

### SetIsLive

`func (o *StraightBet) SetIsLive(v string)`

SetIsLive gets a reference to the given string and assigns it to the IsLive field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


