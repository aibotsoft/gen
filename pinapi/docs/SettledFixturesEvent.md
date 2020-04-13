# SettledFixturesEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Event Id. | [optional] 
**Periods** | Pointer to [**[]SettledFixturesPeriod**](SettledFixturesPeriod.md) | Contains a list of periods. | [optional] 

## Methods

### GetId

`func (o *SettledFixturesEvent) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SettledFixturesEvent) GetIdOk() (int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *SettledFixturesEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *SettledFixturesEvent) SetId(v int64)`

SetId gets a reference to the given int64 and assigns it to the Id field.

### GetPeriods

`func (o *SettledFixturesEvent) GetPeriods() []SettledFixturesPeriod`

GetPeriods returns the Periods field if non-nil, zero value otherwise.

### GetPeriodsOk

`func (o *SettledFixturesEvent) GetPeriodsOk() ([]SettledFixturesPeriod, bool)`

GetPeriodsOk returns a tuple with the Periods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPeriods

`func (o *SettledFixturesEvent) HasPeriods() bool`

HasPeriods returns a boolean if a field has been set.

### SetPeriods

`func (o *SettledFixturesEvent) SetPeriods(v []SettledFixturesPeriod)`

SetPeriods gets a reference to the given []SettledFixturesPeriod and assigns it to the Periods field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


