# TeaserGroupsTeaser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Unique identifier. | [optional] 
**Description** | Pointer to **string** | Description for the Teaser. | [optional] 
**SportId** | Pointer to **int32** | Unique Sport identifier. Sport details can be retrieved from a call to v2/sports endpoint. | [optional] 
**MinLegs** | Pointer to **int32** | Minimum number of legs that must be selected. | [optional] 
**MaxLegs** | Pointer to **int32** | Maximum number of legs that can be selected. | [optional] 
**SameEventOnly** | Pointer to **bool** | If &#39;true&#39; then all legs must be from the same event, otherwise legs can be from different events. | [optional] 
**Payouts** | Pointer to [**[]TeaserGroupsPayout**](TeaserGroupsPayout.md) | A collection of Payout indicating all possible payout combinations. | [optional] 
**Leagues** | Pointer to [**[]TeaserGroupsLeague**](TeaserGroupsLeague.md) | A collection of Leagues available to the teaser. | [optional] 

## Methods

### GetId

`func (o *TeaserGroupsTeaser) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TeaserGroupsTeaser) GetIdOk() (int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *TeaserGroupsTeaser) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *TeaserGroupsTeaser) SetId(v int64)`

SetId gets a reference to the given int64 and assigns it to the Id field.

### GetDescription

`func (o *TeaserGroupsTeaser) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TeaserGroupsTeaser) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *TeaserGroupsTeaser) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *TeaserGroupsTeaser) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetSportId

`func (o *TeaserGroupsTeaser) GetSportId() int32`

GetSportId returns the SportId field if non-nil, zero value otherwise.

### GetSportIdOk

`func (o *TeaserGroupsTeaser) GetSportIdOk() (int32, bool)`

GetSportIdOk returns a tuple with the SportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSportId

`func (o *TeaserGroupsTeaser) HasSportId() bool`

HasSportId returns a boolean if a field has been set.

### SetSportId

`func (o *TeaserGroupsTeaser) SetSportId(v int32)`

SetSportId gets a reference to the given int32 and assigns it to the SportId field.

### GetMinLegs

`func (o *TeaserGroupsTeaser) GetMinLegs() int32`

GetMinLegs returns the MinLegs field if non-nil, zero value otherwise.

### GetMinLegsOk

`func (o *TeaserGroupsTeaser) GetMinLegsOk() (int32, bool)`

GetMinLegsOk returns a tuple with the MinLegs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMinLegs

`func (o *TeaserGroupsTeaser) HasMinLegs() bool`

HasMinLegs returns a boolean if a field has been set.

### SetMinLegs

`func (o *TeaserGroupsTeaser) SetMinLegs(v int32)`

SetMinLegs gets a reference to the given int32 and assigns it to the MinLegs field.

### GetMaxLegs

`func (o *TeaserGroupsTeaser) GetMaxLegs() int32`

GetMaxLegs returns the MaxLegs field if non-nil, zero value otherwise.

### GetMaxLegsOk

`func (o *TeaserGroupsTeaser) GetMaxLegsOk() (int32, bool)`

GetMaxLegsOk returns a tuple with the MaxLegs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMaxLegs

`func (o *TeaserGroupsTeaser) HasMaxLegs() bool`

HasMaxLegs returns a boolean if a field has been set.

### SetMaxLegs

`func (o *TeaserGroupsTeaser) SetMaxLegs(v int32)`

SetMaxLegs gets a reference to the given int32 and assigns it to the MaxLegs field.

### GetSameEventOnly

`func (o *TeaserGroupsTeaser) GetSameEventOnly() bool`

GetSameEventOnly returns the SameEventOnly field if non-nil, zero value otherwise.

### GetSameEventOnlyOk

`func (o *TeaserGroupsTeaser) GetSameEventOnlyOk() (bool, bool)`

GetSameEventOnlyOk returns a tuple with the SameEventOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSameEventOnly

`func (o *TeaserGroupsTeaser) HasSameEventOnly() bool`

HasSameEventOnly returns a boolean if a field has been set.

### SetSameEventOnly

`func (o *TeaserGroupsTeaser) SetSameEventOnly(v bool)`

SetSameEventOnly gets a reference to the given bool and assigns it to the SameEventOnly field.

### GetPayouts

`func (o *TeaserGroupsTeaser) GetPayouts() []TeaserGroupsPayout`

GetPayouts returns the Payouts field if non-nil, zero value otherwise.

### GetPayoutsOk

`func (o *TeaserGroupsTeaser) GetPayoutsOk() ([]TeaserGroupsPayout, bool)`

GetPayoutsOk returns a tuple with the Payouts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPayouts

`func (o *TeaserGroupsTeaser) HasPayouts() bool`

HasPayouts returns a boolean if a field has been set.

### SetPayouts

`func (o *TeaserGroupsTeaser) SetPayouts(v []TeaserGroupsPayout)`

SetPayouts gets a reference to the given []TeaserGroupsPayout and assigns it to the Payouts field.

### GetLeagues

`func (o *TeaserGroupsTeaser) GetLeagues() []TeaserGroupsLeague`

GetLeagues returns the Leagues field if non-nil, zero value otherwise.

### GetLeaguesOk

`func (o *TeaserGroupsTeaser) GetLeaguesOk() ([]TeaserGroupsLeague, bool)`

GetLeaguesOk returns a tuple with the Leagues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLeagues

`func (o *TeaserGroupsTeaser) HasLeagues() bool`

HasLeagues returns a boolean if a field has been set.

### SetLeagues

`func (o *TeaserGroupsTeaser) SetLeagues(v []TeaserGroupsLeague)`

SetLeagues gets a reference to the given []TeaserGroupsLeague and assigns it to the Leagues field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


