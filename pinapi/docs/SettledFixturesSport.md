# SettledFixturesSport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SportId** | Pointer to **int32** | Same as requested sport Id. | [optional] 
**Last** | Pointer to **int64** | Use this value for the subsequent requests for since query parameter to get just the changes since previous response. | [optional] 
**Leagues** | Pointer to [**[]SettledFixturesLeague**](SettledFixturesLeague.md) | Contains a list of Leagues. | [optional] 

## Methods

### GetSportId

`func (o *SettledFixturesSport) GetSportId() int32`

GetSportId returns the SportId field if non-nil, zero value otherwise.

### GetSportIdOk

`func (o *SettledFixturesSport) GetSportIdOk() (int32, bool)`

GetSportIdOk returns a tuple with the SportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSportId

`func (o *SettledFixturesSport) HasSportId() bool`

HasSportId returns a boolean if a field has been set.

### SetSportId

`func (o *SettledFixturesSport) SetSportId(v int32)`

SetSportId gets a reference to the given int32 and assigns it to the SportId field.

### GetLast

`func (o *SettledFixturesSport) GetLast() int64`

GetLast returns the Last field if non-nil, zero value otherwise.

### GetLastOk

`func (o *SettledFixturesSport) GetLastOk() (int64, bool)`

GetLastOk returns a tuple with the Last field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLast

`func (o *SettledFixturesSport) HasLast() bool`

HasLast returns a boolean if a field has been set.

### SetLast

`func (o *SettledFixturesSport) SetLast(v int64)`

SetLast gets a reference to the given int64 and assigns it to the Last field.

### GetLeagues

`func (o *SettledFixturesSport) GetLeagues() []SettledFixturesLeague`

GetLeagues returns the Leagues field if non-nil, zero value otherwise.

### GetLeaguesOk

`func (o *SettledFixturesSport) GetLeaguesOk() ([]SettledFixturesLeague, bool)`

GetLeaguesOk returns a tuple with the Leagues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLeagues

`func (o *SettledFixturesSport) HasLeagues() bool`

HasLeagues returns a boolean if a field has been set.

### SetLeagues

`func (o *SettledFixturesSport) SetLeagues(v []SettledFixturesLeague)`

SetLeagues gets a reference to the given []SettledFixturesLeague and assigns it to the Leagues field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


