# SpecialLineResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | Status [SUCCESS &#x3D; OK, NOT_EXISTS &#x3D; Line not offered anymore] | [optional] 
**SpecialId** | Pointer to **int64** | Special Id. | [optional] 
**ContestantId** | Pointer to **int64** | Contestant Id. | [optional] 
**MinRiskStake** | Pointer to **float64** | Minimum bettable risk amount. | [optional] 
**MaxRiskStake** | Pointer to **float64** | Maximum bettable risk amount. | [optional] 
**MinWinStake** | Pointer to **float64** | Minimum bettable win amount. | [optional] 
**MaxWinStake** | Pointer to **float64** | Maximum bettable win amount. | [optional] 
**LineId** | Pointer to **int64** | Line identification needed to place a bet. | [optional] 
**Price** | Pointer to **float64** | Latest price. | [optional] 
**Handicap** | Pointer to **float64** | Handicap. | [optional] 

## Methods

### GetStatus

`func (o *SpecialLineResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SpecialLineResponse) GetStatusOk() (string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStatus

`func (o *SpecialLineResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatus

`func (o *SpecialLineResponse) SetStatus(v string)`

SetStatus gets a reference to the given string and assigns it to the Status field.

### GetSpecialId

`func (o *SpecialLineResponse) GetSpecialId() int64`

GetSpecialId returns the SpecialId field if non-nil, zero value otherwise.

### GetSpecialIdOk

`func (o *SpecialLineResponse) GetSpecialIdOk() (int64, bool)`

GetSpecialIdOk returns a tuple with the SpecialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSpecialId

`func (o *SpecialLineResponse) HasSpecialId() bool`

HasSpecialId returns a boolean if a field has been set.

### SetSpecialId

`func (o *SpecialLineResponse) SetSpecialId(v int64)`

SetSpecialId gets a reference to the given int64 and assigns it to the SpecialId field.

### GetContestantId

`func (o *SpecialLineResponse) GetContestantId() int64`

GetContestantId returns the ContestantId field if non-nil, zero value otherwise.

### GetContestantIdOk

`func (o *SpecialLineResponse) GetContestantIdOk() (int64, bool)`

GetContestantIdOk returns a tuple with the ContestantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasContestantId

`func (o *SpecialLineResponse) HasContestantId() bool`

HasContestantId returns a boolean if a field has been set.

### SetContestantId

`func (o *SpecialLineResponse) SetContestantId(v int64)`

SetContestantId gets a reference to the given int64 and assigns it to the ContestantId field.

### GetMinRiskStake

`func (o *SpecialLineResponse) GetMinRiskStake() float64`

GetMinRiskStake returns the MinRiskStake field if non-nil, zero value otherwise.

### GetMinRiskStakeOk

`func (o *SpecialLineResponse) GetMinRiskStakeOk() (float64, bool)`

GetMinRiskStakeOk returns a tuple with the MinRiskStake field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMinRiskStake

`func (o *SpecialLineResponse) HasMinRiskStake() bool`

HasMinRiskStake returns a boolean if a field has been set.

### SetMinRiskStake

`func (o *SpecialLineResponse) SetMinRiskStake(v float64)`

SetMinRiskStake gets a reference to the given float64 and assigns it to the MinRiskStake field.

### GetMaxRiskStake

`func (o *SpecialLineResponse) GetMaxRiskStake() float64`

GetMaxRiskStake returns the MaxRiskStake field if non-nil, zero value otherwise.

### GetMaxRiskStakeOk

`func (o *SpecialLineResponse) GetMaxRiskStakeOk() (float64, bool)`

GetMaxRiskStakeOk returns a tuple with the MaxRiskStake field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMaxRiskStake

`func (o *SpecialLineResponse) HasMaxRiskStake() bool`

HasMaxRiskStake returns a boolean if a field has been set.

### SetMaxRiskStake

`func (o *SpecialLineResponse) SetMaxRiskStake(v float64)`

SetMaxRiskStake gets a reference to the given float64 and assigns it to the MaxRiskStake field.

### GetMinWinStake

`func (o *SpecialLineResponse) GetMinWinStake() float64`

GetMinWinStake returns the MinWinStake field if non-nil, zero value otherwise.

### GetMinWinStakeOk

`func (o *SpecialLineResponse) GetMinWinStakeOk() (float64, bool)`

GetMinWinStakeOk returns a tuple with the MinWinStake field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMinWinStake

`func (o *SpecialLineResponse) HasMinWinStake() bool`

HasMinWinStake returns a boolean if a field has been set.

### SetMinWinStake

`func (o *SpecialLineResponse) SetMinWinStake(v float64)`

SetMinWinStake gets a reference to the given float64 and assigns it to the MinWinStake field.

### GetMaxWinStake

`func (o *SpecialLineResponse) GetMaxWinStake() float64`

GetMaxWinStake returns the MaxWinStake field if non-nil, zero value otherwise.

### GetMaxWinStakeOk

`func (o *SpecialLineResponse) GetMaxWinStakeOk() (float64, bool)`

GetMaxWinStakeOk returns a tuple with the MaxWinStake field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMaxWinStake

`func (o *SpecialLineResponse) HasMaxWinStake() bool`

HasMaxWinStake returns a boolean if a field has been set.

### SetMaxWinStake

`func (o *SpecialLineResponse) SetMaxWinStake(v float64)`

SetMaxWinStake gets a reference to the given float64 and assigns it to the MaxWinStake field.

### GetLineId

`func (o *SpecialLineResponse) GetLineId() int64`

GetLineId returns the LineId field if non-nil, zero value otherwise.

### GetLineIdOk

`func (o *SpecialLineResponse) GetLineIdOk() (int64, bool)`

GetLineIdOk returns a tuple with the LineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLineId

`func (o *SpecialLineResponse) HasLineId() bool`

HasLineId returns a boolean if a field has been set.

### SetLineId

`func (o *SpecialLineResponse) SetLineId(v int64)`

SetLineId gets a reference to the given int64 and assigns it to the LineId field.

### GetPrice

`func (o *SpecialLineResponse) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *SpecialLineResponse) GetPriceOk() (float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPrice

`func (o *SpecialLineResponse) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPrice

`func (o *SpecialLineResponse) SetPrice(v float64)`

SetPrice gets a reference to the given float64 and assigns it to the Price field.

### GetHandicap

`func (o *SpecialLineResponse) GetHandicap() float64`

GetHandicap returns the Handicap field if non-nil, zero value otherwise.

### GetHandicapOk

`func (o *SpecialLineResponse) GetHandicapOk() (float64, bool)`

GetHandicapOk returns a tuple with the Handicap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHandicap

`func (o *SpecialLineResponse) HasHandicap() bool`

HasHandicap returns a boolean if a field has been set.

### SetHandicap

`func (o *SpecialLineResponse) SetHandicap(v float64)`

SetHandicap gets a reference to the given float64 and assigns it to the Handicap field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


