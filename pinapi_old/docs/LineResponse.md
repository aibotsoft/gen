# LineResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | If the value is NOT_EXISTS, than this will be the only parameter in the response. All other params would be empty. [SUCCESS &#x3D; OK, NOT_EXISTS &#x3D; Line not offered anymore] | [optional] 
**Price** | Pointer to **float64** | Latest price. | [optional] 
**LineId** | Pointer to **int64** | Line identification needed to place a bet. | [optional] 
**AltLineId** | Pointer to **int64** | This would be needed to place the bet if the handicap is on alternate line, otherwise it will not be populated in the response. | [optional] 
**Team1Score** | Pointer to **int32** | Away team score. Applicable to soccer only. | [optional] 
**Team2Score** | Pointer to **int32** | Home team score. Applicable to soccer only. | [optional] 
**Team1RedCards** | Pointer to **int32** | Team 1 red cards. Applicable to soccer only. | [optional] 
**Team2RedCards** | Pointer to **int32** | Team 2 red cards. Applicable to soccer only. | [optional] 
**MaxRiskStake** | Pointer to **float64** | Maximum bettable risk amount. | [optional] 
**MinRiskStake** | Pointer to **float64** | Minimum bettable risk amount. | [optional] 
**MaxWinStake** | Pointer to **float64** | Maximum bettable win amount. | [optional] 
**MinWinStake** | Pointer to **float64** | Minimum bettable win amount. | [optional] 
**EffectiveAsOf** | Pointer to **string** | Line is effective as of this date and time in UTC. | [optional] 

## Methods

### GetStatus

`func (o *LineResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LineResponse) GetStatusOk() (string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStatus

`func (o *LineResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatus

`func (o *LineResponse) SetStatus(v string)`

SetStatus gets a reference to the given string and assigns it to the Status field.

### GetPrice

`func (o *LineResponse) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *LineResponse) GetPriceOk() (float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPrice

`func (o *LineResponse) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPrice

`func (o *LineResponse) SetPrice(v float64)`

SetPrice gets a reference to the given float64 and assigns it to the Price field.

### GetLineId

`func (o *LineResponse) GetLineId() int64`

GetLineId returns the LineId field if non-nil, zero value otherwise.

### GetLineIdOk

`func (o *LineResponse) GetLineIdOk() (int64, bool)`

GetLineIdOk returns a tuple with the LineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLineId

`func (o *LineResponse) HasLineId() bool`

HasLineId returns a boolean if a field has been set.

### SetLineId

`func (o *LineResponse) SetLineId(v int64)`

SetLineId gets a reference to the given int64 and assigns it to the LineId field.

### GetAltLineId

`func (o *LineResponse) GetAltLineId() int64`

GetAltLineId returns the AltLineId field if non-nil, zero value otherwise.

### GetAltLineIdOk

`func (o *LineResponse) GetAltLineIdOk() (int64, bool)`

GetAltLineIdOk returns a tuple with the AltLineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAltLineId

`func (o *LineResponse) HasAltLineId() bool`

HasAltLineId returns a boolean if a field has been set.

### SetAltLineId

`func (o *LineResponse) SetAltLineId(v int64)`

SetAltLineId gets a reference to the given int64 and assigns it to the AltLineId field.

### GetTeam1Score

`func (o *LineResponse) GetTeam1Score() int32`

GetTeam1Score returns the Team1Score field if non-nil, zero value otherwise.

### GetTeam1ScoreOk

`func (o *LineResponse) GetTeam1ScoreOk() (int32, bool)`

GetTeam1ScoreOk returns a tuple with the Team1Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTeam1Score

`func (o *LineResponse) HasTeam1Score() bool`

HasTeam1Score returns a boolean if a field has been set.

### SetTeam1Score

`func (o *LineResponse) SetTeam1Score(v int32)`

SetTeam1Score gets a reference to the given int32 and assigns it to the Team1Score field.

### GetTeam2Score

`func (o *LineResponse) GetTeam2Score() int32`

GetTeam2Score returns the Team2Score field if non-nil, zero value otherwise.

### GetTeam2ScoreOk

`func (o *LineResponse) GetTeam2ScoreOk() (int32, bool)`

GetTeam2ScoreOk returns a tuple with the Team2Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTeam2Score

`func (o *LineResponse) HasTeam2Score() bool`

HasTeam2Score returns a boolean if a field has been set.

### SetTeam2Score

`func (o *LineResponse) SetTeam2Score(v int32)`

SetTeam2Score gets a reference to the given int32 and assigns it to the Team2Score field.

### GetTeam1RedCards

`func (o *LineResponse) GetTeam1RedCards() int32`

GetTeam1RedCards returns the Team1RedCards field if non-nil, zero value otherwise.

### GetTeam1RedCardsOk

`func (o *LineResponse) GetTeam1RedCardsOk() (int32, bool)`

GetTeam1RedCardsOk returns a tuple with the Team1RedCards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTeam1RedCards

`func (o *LineResponse) HasTeam1RedCards() bool`

HasTeam1RedCards returns a boolean if a field has been set.

### SetTeam1RedCards

`func (o *LineResponse) SetTeam1RedCards(v int32)`

SetTeam1RedCards gets a reference to the given int32 and assigns it to the Team1RedCards field.

### GetTeam2RedCards

`func (o *LineResponse) GetTeam2RedCards() int32`

GetTeam2RedCards returns the Team2RedCards field if non-nil, zero value otherwise.

### GetTeam2RedCardsOk

`func (o *LineResponse) GetTeam2RedCardsOk() (int32, bool)`

GetTeam2RedCardsOk returns a tuple with the Team2RedCards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTeam2RedCards

`func (o *LineResponse) HasTeam2RedCards() bool`

HasTeam2RedCards returns a boolean if a field has been set.

### SetTeam2RedCards

`func (o *LineResponse) SetTeam2RedCards(v int32)`

SetTeam2RedCards gets a reference to the given int32 and assigns it to the Team2RedCards field.

### GetMaxRiskStake

`func (o *LineResponse) GetMaxRiskStake() float64`

GetMaxRiskStake returns the MaxRiskStake field if non-nil, zero value otherwise.

### GetMaxRiskStakeOk

`func (o *LineResponse) GetMaxRiskStakeOk() (float64, bool)`

GetMaxRiskStakeOk returns a tuple with the MaxRiskStake field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMaxRiskStake

`func (o *LineResponse) HasMaxRiskStake() bool`

HasMaxRiskStake returns a boolean if a field has been set.

### SetMaxRiskStake

`func (o *LineResponse) SetMaxRiskStake(v float64)`

SetMaxRiskStake gets a reference to the given float64 and assigns it to the MaxRiskStake field.

### GetMinRiskStake

`func (o *LineResponse) GetMinRiskStake() float64`

GetMinRiskStake returns the MinRiskStake field if non-nil, zero value otherwise.

### GetMinRiskStakeOk

`func (o *LineResponse) GetMinRiskStakeOk() (float64, bool)`

GetMinRiskStakeOk returns a tuple with the MinRiskStake field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMinRiskStake

`func (o *LineResponse) HasMinRiskStake() bool`

HasMinRiskStake returns a boolean if a field has been set.

### SetMinRiskStake

`func (o *LineResponse) SetMinRiskStake(v float64)`

SetMinRiskStake gets a reference to the given float64 and assigns it to the MinRiskStake field.

### GetMaxWinStake

`func (o *LineResponse) GetMaxWinStake() float64`

GetMaxWinStake returns the MaxWinStake field if non-nil, zero value otherwise.

### GetMaxWinStakeOk

`func (o *LineResponse) GetMaxWinStakeOk() (float64, bool)`

GetMaxWinStakeOk returns a tuple with the MaxWinStake field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMaxWinStake

`func (o *LineResponse) HasMaxWinStake() bool`

HasMaxWinStake returns a boolean if a field has been set.

### SetMaxWinStake

`func (o *LineResponse) SetMaxWinStake(v float64)`

SetMaxWinStake gets a reference to the given float64 and assigns it to the MaxWinStake field.

### GetMinWinStake

`func (o *LineResponse) GetMinWinStake() float64`

GetMinWinStake returns the MinWinStake field if non-nil, zero value otherwise.

### GetMinWinStakeOk

`func (o *LineResponse) GetMinWinStakeOk() (float64, bool)`

GetMinWinStakeOk returns a tuple with the MinWinStake field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMinWinStake

`func (o *LineResponse) HasMinWinStake() bool`

HasMinWinStake returns a boolean if a field has been set.

### SetMinWinStake

`func (o *LineResponse) SetMinWinStake(v float64)`

SetMinWinStake gets a reference to the given float64 and assigns it to the MinWinStake field.

### GetEffectiveAsOf

`func (o *LineResponse) GetEffectiveAsOf() string`

GetEffectiveAsOf returns the EffectiveAsOf field if non-nil, zero value otherwise.

### GetEffectiveAsOfOk

`func (o *LineResponse) GetEffectiveAsOfOk() (string, bool)`

GetEffectiveAsOfOk returns a tuple with the EffectiveAsOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEffectiveAsOf

`func (o *LineResponse) HasEffectiveAsOf() bool`

HasEffectiveAsOf returns a boolean if a field has been set.

### SetEffectiveAsOf

`func (o *LineResponse) SetEffectiveAsOf(v string)`

SetEffectiveAsOf gets a reference to the given string and assigns it to the EffectiveAsOf field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


