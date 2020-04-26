# SpecialBetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UniqueRequestId** | Pointer to **string** | This unique id of the place bet requests. This is to support idempotent requests. | [optional] 
**AcceptBetterLine** | Pointer to **bool** | Whether or not to accept a bet when there is a line change in favor of the client. | [optional] 
**OddsFormat** | Pointer to [**OddsFormat**](OddsFormat.md) |  | [optional] 
**Stake** | Pointer to **float64** | amount in client’s currency. | [optional] 
**WinRiskStake** | Pointer to **string** | Whether the stake amount is risk or win amount. | [optional] 
**LineId** | Pointer to **int64** | Line identification. | [optional] 
**SpecialId** | Pointer to **int64** | Special identification. | [optional] 
**ContestantId** | Pointer to **int64** | Contestant identification. | [optional] 

## Methods

### GetUniqueRequestId

`func (o *SpecialBetRequest) GetUniqueRequestId() string`

GetUniqueRequestId returns the UniqueRequestId field if non-nil, zero value otherwise.

### GetUniqueRequestIdOk

`func (o *SpecialBetRequest) GetUniqueRequestIdOk() (string, bool)`

GetUniqueRequestIdOk returns a tuple with the UniqueRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUniqueRequestId

`func (o *SpecialBetRequest) HasUniqueRequestId() bool`

HasUniqueRequestId returns a boolean if a field has been set.

### SetUniqueRequestId

`func (o *SpecialBetRequest) SetUniqueRequestId(v string)`

SetUniqueRequestId gets a reference to the given string and assigns it to the UniqueRequestId field.

### GetAcceptBetterLine

`func (o *SpecialBetRequest) GetAcceptBetterLine() bool`

GetAcceptBetterLine returns the AcceptBetterLine field if non-nil, zero value otherwise.

### GetAcceptBetterLineOk

`func (o *SpecialBetRequest) GetAcceptBetterLineOk() (bool, bool)`

GetAcceptBetterLineOk returns a tuple with the AcceptBetterLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAcceptBetterLine

`func (o *SpecialBetRequest) HasAcceptBetterLine() bool`

HasAcceptBetterLine returns a boolean if a field has been set.

### SetAcceptBetterLine

`func (o *SpecialBetRequest) SetAcceptBetterLine(v bool)`

SetAcceptBetterLine gets a reference to the given bool and assigns it to the AcceptBetterLine field.

### GetOddsFormat

`func (o *SpecialBetRequest) GetOddsFormat() OddsFormat`

GetOddsFormat returns the OddsFormat field if non-nil, zero value otherwise.

### GetOddsFormatOk

`func (o *SpecialBetRequest) GetOddsFormatOk() (OddsFormat, bool)`

GetOddsFormatOk returns a tuple with the OddsFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOddsFormat

`func (o *SpecialBetRequest) HasOddsFormat() bool`

HasOddsFormat returns a boolean if a field has been set.

### SetOddsFormat

`func (o *SpecialBetRequest) SetOddsFormat(v OddsFormat)`

SetOddsFormat gets a reference to the given OddsFormat and assigns it to the OddsFormat field.

### GetStake

`func (o *SpecialBetRequest) GetStake() float64`

GetStake returns the Stake field if non-nil, zero value otherwise.

### GetStakeOk

`func (o *SpecialBetRequest) GetStakeOk() (float64, bool)`

GetStakeOk returns a tuple with the Stake field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStake

`func (o *SpecialBetRequest) HasStake() bool`

HasStake returns a boolean if a field has been set.

### SetStake

`func (o *SpecialBetRequest) SetStake(v float64)`

SetStake gets a reference to the given float64 and assigns it to the Stake field.

### GetWinRiskStake

`func (o *SpecialBetRequest) GetWinRiskStake() string`

GetWinRiskStake returns the WinRiskStake field if non-nil, zero value otherwise.

### GetWinRiskStakeOk

`func (o *SpecialBetRequest) GetWinRiskStakeOk() (string, bool)`

GetWinRiskStakeOk returns a tuple with the WinRiskStake field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWinRiskStake

`func (o *SpecialBetRequest) HasWinRiskStake() bool`

HasWinRiskStake returns a boolean if a field has been set.

### SetWinRiskStake

`func (o *SpecialBetRequest) SetWinRiskStake(v string)`

SetWinRiskStake gets a reference to the given string and assigns it to the WinRiskStake field.

### GetLineId

`func (o *SpecialBetRequest) GetLineId() int64`

GetLineId returns the LineId field if non-nil, zero value otherwise.

### GetLineIdOk

`func (o *SpecialBetRequest) GetLineIdOk() (int64, bool)`

GetLineIdOk returns a tuple with the LineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLineId

`func (o *SpecialBetRequest) HasLineId() bool`

HasLineId returns a boolean if a field has been set.

### SetLineId

`func (o *SpecialBetRequest) SetLineId(v int64)`

SetLineId gets a reference to the given int64 and assigns it to the LineId field.

### GetSpecialId

`func (o *SpecialBetRequest) GetSpecialId() int64`

GetSpecialId returns the SpecialId field if non-nil, zero value otherwise.

### GetSpecialIdOk

`func (o *SpecialBetRequest) GetSpecialIdOk() (int64, bool)`

GetSpecialIdOk returns a tuple with the SpecialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSpecialId

`func (o *SpecialBetRequest) HasSpecialId() bool`

HasSpecialId returns a boolean if a field has been set.

### SetSpecialId

`func (o *SpecialBetRequest) SetSpecialId(v int64)`

SetSpecialId gets a reference to the given int64 and assigns it to the SpecialId field.

### GetContestantId

`func (o *SpecialBetRequest) GetContestantId() int64`

GetContestantId returns the ContestantId field if non-nil, zero value otherwise.

### GetContestantIdOk

`func (o *SpecialBetRequest) GetContestantIdOk() (int64, bool)`

GetContestantIdOk returns a tuple with the ContestantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasContestantId

`func (o *SpecialBetRequest) HasContestantId() bool`

HasContestantId returns a boolean if a field has been set.

### SetContestantId

`func (o *SpecialBetRequest) SetContestantId(v int64)`

SetContestantId gets a reference to the given int64 and assigns it to the ContestantId field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


