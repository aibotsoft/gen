# OddsEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Event Id. | [optional] 
**AwayScore** | Pointer to **float64** | Away team score. Only for live soccer events.Supported only for full match period (number&#x3D;0). | [optional] 
**HomeScore** | Pointer to **float64** | Home team score. Only for live soccer events.Supported only for full match period (number&#x3D;0). | [optional] 
**AwayRedCards** | Pointer to **int32** | Away team red cards. Only for live soccer events. Supported only for full match period (number&#x3D;0). | [optional] 
**HomeRedCards** | Pointer to **int32** | Home team red cards. Only for live soccer events.Supported only for full match period (number&#x3D;0). | [optional] 
**Periods** | Pointer to [**[]OddsPeriod**](OddsPeriod.md) | Contains a list of periods. | [optional] 

## Methods

### GetId

`func (o *OddsEvent) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OddsEvent) GetIdOk() (int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *OddsEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *OddsEvent) SetId(v int64)`

SetId gets a reference to the given int64 and assigns it to the Id field.

### GetAwayScore

`func (o *OddsEvent) GetAwayScore() float64`

GetAwayScore returns the AwayScore field if non-nil, zero value otherwise.

### GetAwayScoreOk

`func (o *OddsEvent) GetAwayScoreOk() (float64, bool)`

GetAwayScoreOk returns a tuple with the AwayScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAwayScore

`func (o *OddsEvent) HasAwayScore() bool`

HasAwayScore returns a boolean if a field has been set.

### SetAwayScore

`func (o *OddsEvent) SetAwayScore(v float64)`

SetAwayScore gets a reference to the given float64 and assigns it to the AwayScore field.

### GetHomeScore

`func (o *OddsEvent) GetHomeScore() float64`

GetHomeScore returns the HomeScore field if non-nil, zero value otherwise.

### GetHomeScoreOk

`func (o *OddsEvent) GetHomeScoreOk() (float64, bool)`

GetHomeScoreOk returns a tuple with the HomeScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHomeScore

`func (o *OddsEvent) HasHomeScore() bool`

HasHomeScore returns a boolean if a field has been set.

### SetHomeScore

`func (o *OddsEvent) SetHomeScore(v float64)`

SetHomeScore gets a reference to the given float64 and assigns it to the HomeScore field.

### GetAwayRedCards

`func (o *OddsEvent) GetAwayRedCards() int32`

GetAwayRedCards returns the AwayRedCards field if non-nil, zero value otherwise.

### GetAwayRedCardsOk

`func (o *OddsEvent) GetAwayRedCardsOk() (int32, bool)`

GetAwayRedCardsOk returns a tuple with the AwayRedCards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAwayRedCards

`func (o *OddsEvent) HasAwayRedCards() bool`

HasAwayRedCards returns a boolean if a field has been set.

### SetAwayRedCards

`func (o *OddsEvent) SetAwayRedCards(v int32)`

SetAwayRedCards gets a reference to the given int32 and assigns it to the AwayRedCards field.

### GetHomeRedCards

`func (o *OddsEvent) GetHomeRedCards() int32`

GetHomeRedCards returns the HomeRedCards field if non-nil, zero value otherwise.

### GetHomeRedCardsOk

`func (o *OddsEvent) GetHomeRedCardsOk() (int32, bool)`

GetHomeRedCardsOk returns a tuple with the HomeRedCards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHomeRedCards

`func (o *OddsEvent) HasHomeRedCards() bool`

HasHomeRedCards returns a boolean if a field has been set.

### SetHomeRedCards

`func (o *OddsEvent) SetHomeRedCards(v int32)`

SetHomeRedCards gets a reference to the given int32 and assigns it to the HomeRedCards field.

### GetPeriods

`func (o *OddsEvent) GetPeriods() []OddsPeriod`

GetPeriods returns the Periods field if non-nil, zero value otherwise.

### GetPeriodsOk

`func (o *OddsEvent) GetPeriodsOk() ([]OddsPeriod, bool)`

GetPeriodsOk returns a tuple with the Periods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPeriods

`func (o *OddsEvent) HasPeriods() bool`

HasPeriods returns a boolean if a field has been set.

### SetPeriods

`func (o *OddsEvent) SetPeriods(v []OddsPeriod)`

SetPeriods gets a reference to the given []OddsPeriod and assigns it to the Periods field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


