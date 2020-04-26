# SpecialBet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BetId** | Pointer to **int64** | Bet identification | 
**UniqueRequestId** | Pointer to **string** | Unique Request Id | [optional] 
**WagerNumber** | Pointer to **int32** | Wager identification. All bets placed thru the API will have value 1. Website Classic view supports multiple contest(special) bets placement in the same bet slip in that case the bet would have appropriate wager number, as well as all round robin parlay bets. | 
**PlacedAt** | Pointer to [**time.Time**](time.Time.md) | Date time when the bet was placed. | 
**BetStatus** | Pointer to **string** | Bet Status.  ACCEPTED &#x3D; Bet was accepted,  CANCELLED &#x3D; Bet is cancelled as per Pinnacle betting rules,  LOSE &#x3D; The bet is settled as lose, REFUNDED &#x3D; When an event is cancelled or when the bet is settled as push, the bet will have REFUNDED status,  WON &#x3D; The bet is settled as won   | 
**BetType** | Pointer to **string** |  | [default to SPECIAL]
**Win** | Pointer to **float64** | Win amount. | 
**Risk** | Pointer to **float64** | Risk amount. | 
**WinLoss** | Pointer to **NullableFloat64** | Win-Loss for settled bets. | [optional] 
**OddsFormat** | Pointer to [**OddsFormat**](OddsFormat.md) |  | 
**CustomerCommission** | Pointer to **NullableFloat64** | Client’s commission on the bet. | [optional] 
**CancellationReason** | Pointer to [**CancellationReason**](CancellationReason.md) |  | [optional] 
**UpdateSequence** | Pointer to **int64** | Update Sequence. It gets updated when the bet status change. | 
**SpecialId** | Pointer to **int64** |  | 
**SpecialName** | Pointer to **string** |  | 
**ContestantId** | Pointer to **int64** |  | 
**ContestantName** | Pointer to **string** |  | 
**Price** | Pointer to **float64** |  | 
**Handicap** | Pointer to **float64** |  | [optional] 
**Units** | Pointer to **string** |  | [optional] 
**SportId** | Pointer to **int32** |  | 
**LeagueId** | Pointer to **int32** |  | 
**EventId** | Pointer to **NullableInt64** | Populated if bet was placed on a special linked to the event. | [optional] 
**PeriodNumber** | Pointer to **NullableInt32** | Populated if bet was placed on a special linked to the event. | [optional] 
**Team1** | Pointer to **NullableString** | Populated if bet was placed on a special linked to the event. | [optional] 
**Team2** | Pointer to **NullableString** | Populated if bet was placed on a special linked to the event. | [optional] 
**EventStartTime** | Pointer to [**time.Time**](time.Time.md) | Date time when the event starts | [optional] 

## Methods

### GetBetId

`func (o *SpecialBet) GetBetId() int64`

GetBetId returns the BetId field if non-nil, zero value otherwise.

### GetBetIdOk

`func (o *SpecialBet) GetBetIdOk() (int64, bool)`

GetBetIdOk returns a tuple with the BetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBetId

`func (o *SpecialBet) HasBetId() bool`

HasBetId returns a boolean if a field has been set.

### SetBetId

`func (o *SpecialBet) SetBetId(v int64)`

SetBetId gets a reference to the given int64 and assigns it to the BetId field.

### GetUniqueRequestId

`func (o *SpecialBet) GetUniqueRequestId() string`

GetUniqueRequestId returns the UniqueRequestId field if non-nil, zero value otherwise.

### GetUniqueRequestIdOk

`func (o *SpecialBet) GetUniqueRequestIdOk() (string, bool)`

GetUniqueRequestIdOk returns a tuple with the UniqueRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUniqueRequestId

`func (o *SpecialBet) HasUniqueRequestId() bool`

HasUniqueRequestId returns a boolean if a field has been set.

### SetUniqueRequestId

`func (o *SpecialBet) SetUniqueRequestId(v string)`

SetUniqueRequestId gets a reference to the given string and assigns it to the UniqueRequestId field.

### GetWagerNumber

`func (o *SpecialBet) GetWagerNumber() int32`

GetWagerNumber returns the WagerNumber field if non-nil, zero value otherwise.

### GetWagerNumberOk

`func (o *SpecialBet) GetWagerNumberOk() (int32, bool)`

GetWagerNumberOk returns a tuple with the WagerNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWagerNumber

`func (o *SpecialBet) HasWagerNumber() bool`

HasWagerNumber returns a boolean if a field has been set.

### SetWagerNumber

`func (o *SpecialBet) SetWagerNumber(v int32)`

SetWagerNumber gets a reference to the given int32 and assigns it to the WagerNumber field.

### GetPlacedAt

`func (o *SpecialBet) GetPlacedAt() time.Time`

GetPlacedAt returns the PlacedAt field if non-nil, zero value otherwise.

### GetPlacedAtOk

`func (o *SpecialBet) GetPlacedAtOk() (time.Time, bool)`

GetPlacedAtOk returns a tuple with the PlacedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPlacedAt

`func (o *SpecialBet) HasPlacedAt() bool`

HasPlacedAt returns a boolean if a field has been set.

### SetPlacedAt

`func (o *SpecialBet) SetPlacedAt(v time.Time)`

SetPlacedAt gets a reference to the given time.Time and assigns it to the PlacedAt field.

### GetBetStatus

`func (o *SpecialBet) GetBetStatus() string`

GetBetStatus returns the BetStatus field if non-nil, zero value otherwise.

### GetBetStatusOk

`func (o *SpecialBet) GetBetStatusOk() (string, bool)`

GetBetStatusOk returns a tuple with the BetStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBetStatus

`func (o *SpecialBet) HasBetStatus() bool`

HasBetStatus returns a boolean if a field has been set.

### SetBetStatus

`func (o *SpecialBet) SetBetStatus(v string)`

SetBetStatus gets a reference to the given string and assigns it to the BetStatus field.

### GetBetType

`func (o *SpecialBet) GetBetType() string`

GetBetType returns the BetType field if non-nil, zero value otherwise.

### GetBetTypeOk

`func (o *SpecialBet) GetBetTypeOk() (string, bool)`

GetBetTypeOk returns a tuple with the BetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBetType

`func (o *SpecialBet) HasBetType() bool`

HasBetType returns a boolean if a field has been set.

### SetBetType

`func (o *SpecialBet) SetBetType(v string)`

SetBetType gets a reference to the given string and assigns it to the BetType field.

### GetWin

`func (o *SpecialBet) GetWin() float64`

GetWin returns the Win field if non-nil, zero value otherwise.

### GetWinOk

`func (o *SpecialBet) GetWinOk() (float64, bool)`

GetWinOk returns a tuple with the Win field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWin

`func (o *SpecialBet) HasWin() bool`

HasWin returns a boolean if a field has been set.

### SetWin

`func (o *SpecialBet) SetWin(v float64)`

SetWin gets a reference to the given float64 and assigns it to the Win field.

### GetRisk

`func (o *SpecialBet) GetRisk() float64`

GetRisk returns the Risk field if non-nil, zero value otherwise.

### GetRiskOk

`func (o *SpecialBet) GetRiskOk() (float64, bool)`

GetRiskOk returns a tuple with the Risk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRisk

`func (o *SpecialBet) HasRisk() bool`

HasRisk returns a boolean if a field has been set.

### SetRisk

`func (o *SpecialBet) SetRisk(v float64)`

SetRisk gets a reference to the given float64 and assigns it to the Risk field.

### GetWinLoss

`func (o *SpecialBet) GetWinLoss() NullableFloat64`

GetWinLoss returns the WinLoss field if non-nil, zero value otherwise.

### GetWinLossOk

`func (o *SpecialBet) GetWinLossOk() (NullableFloat64, bool)`

GetWinLossOk returns a tuple with the WinLoss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWinLoss

`func (o *SpecialBet) HasWinLoss() bool`

HasWinLoss returns a boolean if a field has been set.

### SetWinLoss

`func (o *SpecialBet) SetWinLoss(v NullableFloat64)`

SetWinLoss gets a reference to the given NullableFloat64 and assigns it to the WinLoss field.

### SetWinLossExplicitNull

`func (o *SpecialBet) SetWinLossExplicitNull(b bool)`

SetWinLossExplicitNull (un)sets WinLoss to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The WinLoss value is set to nil even if false is passed
### GetOddsFormat

`func (o *SpecialBet) GetOddsFormat() OddsFormat`

GetOddsFormat returns the OddsFormat field if non-nil, zero value otherwise.

### GetOddsFormatOk

`func (o *SpecialBet) GetOddsFormatOk() (OddsFormat, bool)`

GetOddsFormatOk returns a tuple with the OddsFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOddsFormat

`func (o *SpecialBet) HasOddsFormat() bool`

HasOddsFormat returns a boolean if a field has been set.

### SetOddsFormat

`func (o *SpecialBet) SetOddsFormat(v OddsFormat)`

SetOddsFormat gets a reference to the given OddsFormat and assigns it to the OddsFormat field.

### GetCustomerCommission

`func (o *SpecialBet) GetCustomerCommission() NullableFloat64`

GetCustomerCommission returns the CustomerCommission field if non-nil, zero value otherwise.

### GetCustomerCommissionOk

`func (o *SpecialBet) GetCustomerCommissionOk() (NullableFloat64, bool)`

GetCustomerCommissionOk returns a tuple with the CustomerCommission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCustomerCommission

`func (o *SpecialBet) HasCustomerCommission() bool`

HasCustomerCommission returns a boolean if a field has been set.

### SetCustomerCommission

`func (o *SpecialBet) SetCustomerCommission(v NullableFloat64)`

SetCustomerCommission gets a reference to the given NullableFloat64 and assigns it to the CustomerCommission field.

### SetCustomerCommissionExplicitNull

`func (o *SpecialBet) SetCustomerCommissionExplicitNull(b bool)`

SetCustomerCommissionExplicitNull (un)sets CustomerCommission to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CustomerCommission value is set to nil even if false is passed
### GetCancellationReason

`func (o *SpecialBet) GetCancellationReason() CancellationReason`

GetCancellationReason returns the CancellationReason field if non-nil, zero value otherwise.

### GetCancellationReasonOk

`func (o *SpecialBet) GetCancellationReasonOk() (CancellationReason, bool)`

GetCancellationReasonOk returns a tuple with the CancellationReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCancellationReason

`func (o *SpecialBet) HasCancellationReason() bool`

HasCancellationReason returns a boolean if a field has been set.

### SetCancellationReason

`func (o *SpecialBet) SetCancellationReason(v CancellationReason)`

SetCancellationReason gets a reference to the given CancellationReason and assigns it to the CancellationReason field.

### GetUpdateSequence

`func (o *SpecialBet) GetUpdateSequence() int64`

GetUpdateSequence returns the UpdateSequence field if non-nil, zero value otherwise.

### GetUpdateSequenceOk

`func (o *SpecialBet) GetUpdateSequenceOk() (int64, bool)`

GetUpdateSequenceOk returns a tuple with the UpdateSequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUpdateSequence

`func (o *SpecialBet) HasUpdateSequence() bool`

HasUpdateSequence returns a boolean if a field has been set.

### SetUpdateSequence

`func (o *SpecialBet) SetUpdateSequence(v int64)`

SetUpdateSequence gets a reference to the given int64 and assigns it to the UpdateSequence field.

### GetSpecialId

`func (o *SpecialBet) GetSpecialId() int64`

GetSpecialId returns the SpecialId field if non-nil, zero value otherwise.

### GetSpecialIdOk

`func (o *SpecialBet) GetSpecialIdOk() (int64, bool)`

GetSpecialIdOk returns a tuple with the SpecialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSpecialId

`func (o *SpecialBet) HasSpecialId() bool`

HasSpecialId returns a boolean if a field has been set.

### SetSpecialId

`func (o *SpecialBet) SetSpecialId(v int64)`

SetSpecialId gets a reference to the given int64 and assigns it to the SpecialId field.

### GetSpecialName

`func (o *SpecialBet) GetSpecialName() string`

GetSpecialName returns the SpecialName field if non-nil, zero value otherwise.

### GetSpecialNameOk

`func (o *SpecialBet) GetSpecialNameOk() (string, bool)`

GetSpecialNameOk returns a tuple with the SpecialName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSpecialName

`func (o *SpecialBet) HasSpecialName() bool`

HasSpecialName returns a boolean if a field has been set.

### SetSpecialName

`func (o *SpecialBet) SetSpecialName(v string)`

SetSpecialName gets a reference to the given string and assigns it to the SpecialName field.

### GetContestantId

`func (o *SpecialBet) GetContestantId() int64`

GetContestantId returns the ContestantId field if non-nil, zero value otherwise.

### GetContestantIdOk

`func (o *SpecialBet) GetContestantIdOk() (int64, bool)`

GetContestantIdOk returns a tuple with the ContestantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasContestantId

`func (o *SpecialBet) HasContestantId() bool`

HasContestantId returns a boolean if a field has been set.

### SetContestantId

`func (o *SpecialBet) SetContestantId(v int64)`

SetContestantId gets a reference to the given int64 and assigns it to the ContestantId field.

### GetContestantName

`func (o *SpecialBet) GetContestantName() string`

GetContestantName returns the ContestantName field if non-nil, zero value otherwise.

### GetContestantNameOk

`func (o *SpecialBet) GetContestantNameOk() (string, bool)`

GetContestantNameOk returns a tuple with the ContestantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasContestantName

`func (o *SpecialBet) HasContestantName() bool`

HasContestantName returns a boolean if a field has been set.

### SetContestantName

`func (o *SpecialBet) SetContestantName(v string)`

SetContestantName gets a reference to the given string and assigns it to the ContestantName field.

### GetPrice

`func (o *SpecialBet) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *SpecialBet) GetPriceOk() (float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPrice

`func (o *SpecialBet) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPrice

`func (o *SpecialBet) SetPrice(v float64)`

SetPrice gets a reference to the given float64 and assigns it to the Price field.

### GetHandicap

`func (o *SpecialBet) GetHandicap() float64`

GetHandicap returns the Handicap field if non-nil, zero value otherwise.

### GetHandicapOk

`func (o *SpecialBet) GetHandicapOk() (float64, bool)`

GetHandicapOk returns a tuple with the Handicap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHandicap

`func (o *SpecialBet) HasHandicap() bool`

HasHandicap returns a boolean if a field has been set.

### SetHandicap

`func (o *SpecialBet) SetHandicap(v float64)`

SetHandicap gets a reference to the given float64 and assigns it to the Handicap field.

### GetUnits

`func (o *SpecialBet) GetUnits() string`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *SpecialBet) GetUnitsOk() (string, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUnits

`func (o *SpecialBet) HasUnits() bool`

HasUnits returns a boolean if a field has been set.

### SetUnits

`func (o *SpecialBet) SetUnits(v string)`

SetUnits gets a reference to the given string and assigns it to the Units field.

### GetSportId

`func (o *SpecialBet) GetSportId() int32`

GetSportId returns the SportId field if non-nil, zero value otherwise.

### GetSportIdOk

`func (o *SpecialBet) GetSportIdOk() (int32, bool)`

GetSportIdOk returns a tuple with the SportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSportId

`func (o *SpecialBet) HasSportId() bool`

HasSportId returns a boolean if a field has been set.

### SetSportId

`func (o *SpecialBet) SetSportId(v int32)`

SetSportId gets a reference to the given int32 and assigns it to the SportId field.

### GetLeagueId

`func (o *SpecialBet) GetLeagueId() int32`

GetLeagueId returns the LeagueId field if non-nil, zero value otherwise.

### GetLeagueIdOk

`func (o *SpecialBet) GetLeagueIdOk() (int32, bool)`

GetLeagueIdOk returns a tuple with the LeagueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLeagueId

`func (o *SpecialBet) HasLeagueId() bool`

HasLeagueId returns a boolean if a field has been set.

### SetLeagueId

`func (o *SpecialBet) SetLeagueId(v int32)`

SetLeagueId gets a reference to the given int32 and assigns it to the LeagueId field.

### GetEventId

`func (o *SpecialBet) GetEventId() NullableInt64`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *SpecialBet) GetEventIdOk() (NullableInt64, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEventId

`func (o *SpecialBet) HasEventId() bool`

HasEventId returns a boolean if a field has been set.

### SetEventId

`func (o *SpecialBet) SetEventId(v NullableInt64)`

SetEventId gets a reference to the given NullableInt64 and assigns it to the EventId field.

### SetEventIdExplicitNull

`func (o *SpecialBet) SetEventIdExplicitNull(b bool)`

SetEventIdExplicitNull (un)sets EventId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The EventId value is set to nil even if false is passed
### GetPeriodNumber

`func (o *SpecialBet) GetPeriodNumber() NullableInt32`

GetPeriodNumber returns the PeriodNumber field if non-nil, zero value otherwise.

### GetPeriodNumberOk

`func (o *SpecialBet) GetPeriodNumberOk() (NullableInt32, bool)`

GetPeriodNumberOk returns a tuple with the PeriodNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPeriodNumber

`func (o *SpecialBet) HasPeriodNumber() bool`

HasPeriodNumber returns a boolean if a field has been set.

### SetPeriodNumber

`func (o *SpecialBet) SetPeriodNumber(v NullableInt32)`

SetPeriodNumber gets a reference to the given NullableInt32 and assigns it to the PeriodNumber field.

### SetPeriodNumberExplicitNull

`func (o *SpecialBet) SetPeriodNumberExplicitNull(b bool)`

SetPeriodNumberExplicitNull (un)sets PeriodNumber to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PeriodNumber value is set to nil even if false is passed
### GetTeam1

`func (o *SpecialBet) GetTeam1() NullableString`

GetTeam1 returns the Team1 field if non-nil, zero value otherwise.

### GetTeam1Ok

`func (o *SpecialBet) GetTeam1Ok() (NullableString, bool)`

GetTeam1Ok returns a tuple with the Team1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTeam1

`func (o *SpecialBet) HasTeam1() bool`

HasTeam1 returns a boolean if a field has been set.

### SetTeam1

`func (o *SpecialBet) SetTeam1(v NullableString)`

SetTeam1 gets a reference to the given NullableString and assigns it to the Team1 field.

### SetTeam1ExplicitNull

`func (o *SpecialBet) SetTeam1ExplicitNull(b bool)`

SetTeam1ExplicitNull (un)sets Team1 to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Team1 value is set to nil even if false is passed
### GetTeam2

`func (o *SpecialBet) GetTeam2() NullableString`

GetTeam2 returns the Team2 field if non-nil, zero value otherwise.

### GetTeam2Ok

`func (o *SpecialBet) GetTeam2Ok() (NullableString, bool)`

GetTeam2Ok returns a tuple with the Team2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTeam2

`func (o *SpecialBet) HasTeam2() bool`

HasTeam2 returns a boolean if a field has been set.

### SetTeam2

`func (o *SpecialBet) SetTeam2(v NullableString)`

SetTeam2 gets a reference to the given NullableString and assigns it to the Team2 field.

### SetTeam2ExplicitNull

`func (o *SpecialBet) SetTeam2ExplicitNull(b bool)`

SetTeam2ExplicitNull (un)sets Team2 to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Team2 value is set to nil even if false is passed
### GetEventStartTime

`func (o *SpecialBet) GetEventStartTime() time.Time`

GetEventStartTime returns the EventStartTime field if non-nil, zero value otherwise.

### GetEventStartTimeOk

`func (o *SpecialBet) GetEventStartTimeOk() (time.Time, bool)`

GetEventStartTimeOk returns a tuple with the EventStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEventStartTime

`func (o *SpecialBet) HasEventStartTime() bool`

HasEventStartTime returns a boolean if a field has been set.

### SetEventStartTime

`func (o *SpecialBet) SetEventStartTime(v time.Time)`

SetEventStartTime gets a reference to the given time.Time and assigns it to the EventStartTime field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


