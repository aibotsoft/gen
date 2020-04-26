# PlaceBetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OddsFormat** | Pointer to [**OddsFormat**](OddsFormat.md) |  | [optional] 
**UniqueRequestId** | Pointer to **string** | This is a Unique ID for PlaceBet requests. This is to support idempotent requests. | [optional] 
**AcceptBetterLine** | Pointer to **bool** | Whether or not to accept a bet when there is a line change in favor of the client. | [optional] 
**Stake** | Pointer to **float64** | amount in client’s currency. | [optional] 
**WinRiskStake** | Pointer to **string** | Whether the stake amount is risk or win amount. | [optional] 
**LineId** | Pointer to **int64** | Line identification. | [optional] 
**AltLineId** | Pointer to **NullableInt64** | Alternate line identification. | [optional] 
**Pitcher1MustStart** | Pointer to **bool** | Baseball only. Refers to the pitcher for Team1. This applicable only for MONEYLINE bet type, for all other bet types this has to be TRUE. | [optional] 
**Pitcher2MustStart** | Pointer to **bool** | Baseball only. Refers to the pitcher for Team2. This applicable only for MONEYLINE bet type, for all other bet types this has to be TRUE. | [optional] 
**FillType** | Pointer to **string** | NORMAL - bet will be placed on specified stake.   FILLANDKILL - If the stake is over the max limit, bet will be placed on max limit, otherwise it will be placed on specified stake.   FILLMAXLIMIT - bet will be places on max limit, stake amount will be ignored. Please note that maximum limits can change at any moment, which may result in risking more than anticipated. This option is replacement of isMaxStakeBet from v1/bets/place&#39;  | [optional] [default to FILL_TYPE_NORMAL]
**SportId** | Pointer to **int32** |  | [optional] 
**EventId** | Pointer to **int64** |  | [optional] 
**PeriodNumber** | Pointer to **int32** |  | [optional] 
**BetType** | Pointer to **string** | Bet type. | [optional] 
**Team** | Pointer to **string** | Team type. | [optional] 
**Side** | Pointer to **NullableString** | Side type. | [optional] 

## Methods

### GetOddsFormat

`func (o *PlaceBetRequest) GetOddsFormat() OddsFormat`

GetOddsFormat returns the OddsFormat field if non-nil, zero value otherwise.

### GetOddsFormatOk

`func (o *PlaceBetRequest) GetOddsFormatOk() (OddsFormat, bool)`

GetOddsFormatOk returns a tuple with the OddsFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOddsFormat

`func (o *PlaceBetRequest) HasOddsFormat() bool`

HasOddsFormat returns a boolean if a field has been set.

### SetOddsFormat

`func (o *PlaceBetRequest) SetOddsFormat(v OddsFormat)`

SetOddsFormat gets a reference to the given OddsFormat and assigns it to the OddsFormat field.

### GetUniqueRequestId

`func (o *PlaceBetRequest) GetUniqueRequestId() string`

GetUniqueRequestId returns the UniqueRequestId field if non-nil, zero value otherwise.

### GetUniqueRequestIdOk

`func (o *PlaceBetRequest) GetUniqueRequestIdOk() (string, bool)`

GetUniqueRequestIdOk returns a tuple with the UniqueRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUniqueRequestId

`func (o *PlaceBetRequest) HasUniqueRequestId() bool`

HasUniqueRequestId returns a boolean if a field has been set.

### SetUniqueRequestId

`func (o *PlaceBetRequest) SetUniqueRequestId(v string)`

SetUniqueRequestId gets a reference to the given string and assigns it to the UniqueRequestId field.

### GetAcceptBetterLine

`func (o *PlaceBetRequest) GetAcceptBetterLine() bool`

GetAcceptBetterLine returns the AcceptBetterLine field if non-nil, zero value otherwise.

### GetAcceptBetterLineOk

`func (o *PlaceBetRequest) GetAcceptBetterLineOk() (bool, bool)`

GetAcceptBetterLineOk returns a tuple with the AcceptBetterLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAcceptBetterLine

`func (o *PlaceBetRequest) HasAcceptBetterLine() bool`

HasAcceptBetterLine returns a boolean if a field has been set.

### SetAcceptBetterLine

`func (o *PlaceBetRequest) SetAcceptBetterLine(v bool)`

SetAcceptBetterLine gets a reference to the given bool and assigns it to the AcceptBetterLine field.

### GetStake

`func (o *PlaceBetRequest) GetStake() float64`

GetStake returns the Stake field if non-nil, zero value otherwise.

### GetStakeOk

`func (o *PlaceBetRequest) GetStakeOk() (float64, bool)`

GetStakeOk returns a tuple with the Stake field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStake

`func (o *PlaceBetRequest) HasStake() bool`

HasStake returns a boolean if a field has been set.

### SetStake

`func (o *PlaceBetRequest) SetStake(v float64)`

SetStake gets a reference to the given float64 and assigns it to the Stake field.

### GetWinRiskStake

`func (o *PlaceBetRequest) GetWinRiskStake() string`

GetWinRiskStake returns the WinRiskStake field if non-nil, zero value otherwise.

### GetWinRiskStakeOk

`func (o *PlaceBetRequest) GetWinRiskStakeOk() (string, bool)`

GetWinRiskStakeOk returns a tuple with the WinRiskStake field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWinRiskStake

`func (o *PlaceBetRequest) HasWinRiskStake() bool`

HasWinRiskStake returns a boolean if a field has been set.

### SetWinRiskStake

`func (o *PlaceBetRequest) SetWinRiskStake(v string)`

SetWinRiskStake gets a reference to the given string and assigns it to the WinRiskStake field.

### GetLineId

`func (o *PlaceBetRequest) GetLineId() int64`

GetLineId returns the LineId field if non-nil, zero value otherwise.

### GetLineIdOk

`func (o *PlaceBetRequest) GetLineIdOk() (int64, bool)`

GetLineIdOk returns a tuple with the LineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLineId

`func (o *PlaceBetRequest) HasLineId() bool`

HasLineId returns a boolean if a field has been set.

### SetLineId

`func (o *PlaceBetRequest) SetLineId(v int64)`

SetLineId gets a reference to the given int64 and assigns it to the LineId field.

### GetAltLineId

`func (o *PlaceBetRequest) GetAltLineId() NullableInt64`

GetAltLineId returns the AltLineId field if non-nil, zero value otherwise.

### GetAltLineIdOk

`func (o *PlaceBetRequest) GetAltLineIdOk() (NullableInt64, bool)`

GetAltLineIdOk returns a tuple with the AltLineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAltLineId

`func (o *PlaceBetRequest) HasAltLineId() bool`

HasAltLineId returns a boolean if a field has been set.

### SetAltLineId

`func (o *PlaceBetRequest) SetAltLineId(v NullableInt64)`

SetAltLineId gets a reference to the given NullableInt64 and assigns it to the AltLineId field.

### SetAltLineIdExplicitNull

`func (o *PlaceBetRequest) SetAltLineIdExplicitNull(b bool)`

SetAltLineIdExplicitNull (un)sets AltLineId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The AltLineId value is set to nil even if false is passed
### GetPitcher1MustStart

`func (o *PlaceBetRequest) GetPitcher1MustStart() bool`

GetPitcher1MustStart returns the Pitcher1MustStart field if non-nil, zero value otherwise.

### GetPitcher1MustStartOk

`func (o *PlaceBetRequest) GetPitcher1MustStartOk() (bool, bool)`

GetPitcher1MustStartOk returns a tuple with the Pitcher1MustStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPitcher1MustStart

`func (o *PlaceBetRequest) HasPitcher1MustStart() bool`

HasPitcher1MustStart returns a boolean if a field has been set.

### SetPitcher1MustStart

`func (o *PlaceBetRequest) SetPitcher1MustStart(v bool)`

SetPitcher1MustStart gets a reference to the given bool and assigns it to the Pitcher1MustStart field.

### GetPitcher2MustStart

`func (o *PlaceBetRequest) GetPitcher2MustStart() bool`

GetPitcher2MustStart returns the Pitcher2MustStart field if non-nil, zero value otherwise.

### GetPitcher2MustStartOk

`func (o *PlaceBetRequest) GetPitcher2MustStartOk() (bool, bool)`

GetPitcher2MustStartOk returns a tuple with the Pitcher2MustStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPitcher2MustStart

`func (o *PlaceBetRequest) HasPitcher2MustStart() bool`

HasPitcher2MustStart returns a boolean if a field has been set.

### SetPitcher2MustStart

`func (o *PlaceBetRequest) SetPitcher2MustStart(v bool)`

SetPitcher2MustStart gets a reference to the given bool and assigns it to the Pitcher2MustStart field.

### GetFillType

`func (o *PlaceBetRequest) GetFillType() string`

GetFillType returns the FillType field if non-nil, zero value otherwise.

### GetFillTypeOk

`func (o *PlaceBetRequest) GetFillTypeOk() (string, bool)`

GetFillTypeOk returns a tuple with the FillType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFillType

`func (o *PlaceBetRequest) HasFillType() bool`

HasFillType returns a boolean if a field has been set.

### SetFillType

`func (o *PlaceBetRequest) SetFillType(v string)`

SetFillType gets a reference to the given string and assigns it to the FillType field.

### GetSportId

`func (o *PlaceBetRequest) GetSportId() int32`

GetSportId returns the SportId field if non-nil, zero value otherwise.

### GetSportIdOk

`func (o *PlaceBetRequest) GetSportIdOk() (int32, bool)`

GetSportIdOk returns a tuple with the SportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSportId

`func (o *PlaceBetRequest) HasSportId() bool`

HasSportId returns a boolean if a field has been set.

### SetSportId

`func (o *PlaceBetRequest) SetSportId(v int32)`

SetSportId gets a reference to the given int32 and assigns it to the SportId field.

### GetEventId

`func (o *PlaceBetRequest) GetEventId() int64`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *PlaceBetRequest) GetEventIdOk() (int64, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEventId

`func (o *PlaceBetRequest) HasEventId() bool`

HasEventId returns a boolean if a field has been set.

### SetEventId

`func (o *PlaceBetRequest) SetEventId(v int64)`

SetEventId gets a reference to the given int64 and assigns it to the EventId field.

### GetPeriodNumber

`func (o *PlaceBetRequest) GetPeriodNumber() int32`

GetPeriodNumber returns the PeriodNumber field if non-nil, zero value otherwise.

### GetPeriodNumberOk

`func (o *PlaceBetRequest) GetPeriodNumberOk() (int32, bool)`

GetPeriodNumberOk returns a tuple with the PeriodNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPeriodNumber

`func (o *PlaceBetRequest) HasPeriodNumber() bool`

HasPeriodNumber returns a boolean if a field has been set.

### SetPeriodNumber

`func (o *PlaceBetRequest) SetPeriodNumber(v int32)`

SetPeriodNumber gets a reference to the given int32 and assigns it to the PeriodNumber field.

### GetBetType

`func (o *PlaceBetRequest) GetBetType() string`

GetBetType returns the BetType field if non-nil, zero value otherwise.

### GetBetTypeOk

`func (o *PlaceBetRequest) GetBetTypeOk() (string, bool)`

GetBetTypeOk returns a tuple with the BetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBetType

`func (o *PlaceBetRequest) HasBetType() bool`

HasBetType returns a boolean if a field has been set.

### SetBetType

`func (o *PlaceBetRequest) SetBetType(v string)`

SetBetType gets a reference to the given string and assigns it to the BetType field.

### GetTeam

`func (o *PlaceBetRequest) GetTeam() string`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *PlaceBetRequest) GetTeamOk() (string, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTeam

`func (o *PlaceBetRequest) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### SetTeam

`func (o *PlaceBetRequest) SetTeam(v string)`

SetTeam gets a reference to the given string and assigns it to the Team field.

### GetSide

`func (o *PlaceBetRequest) GetSide() NullableString`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *PlaceBetRequest) GetSideOk() (NullableString, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSide

`func (o *PlaceBetRequest) HasSide() bool`

HasSide returns a boolean if a field has been set.

### SetSide

`func (o *PlaceBetRequest) SetSide(v NullableString)`

SetSide gets a reference to the given NullableString and assigns it to the Side field.

### SetSideExplicitNull

`func (o *PlaceBetRequest) SetSideExplicitNull(b bool)`

SetSideExplicitNull (un)sets Side to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Side value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


