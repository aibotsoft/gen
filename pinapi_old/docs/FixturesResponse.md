# FixturesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SportId** | Pointer to **int32** | Same as requested sport Id. | [optional] 
**Last** | Pointer to **int64** | Use this value for the subsequent requests for since query parameter to get just the changes since previous response. | [optional] 
**League** | Pointer to [**[]FixturesLeague**](FixturesLeague.md) | Contains a list of Leagues. | [optional] 

## Methods

### GetSportId

`func (o *FixturesResponse) GetSportId() int32`

GetSportId returns the SportId field if non-nil, zero value otherwise.

### GetSportIdOk

`func (o *FixturesResponse) GetSportIdOk() (int32, bool)`

GetSportIdOk returns a tuple with the SportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSportId

`func (o *FixturesResponse) HasSportId() bool`

HasSportId returns a boolean if a field has been set.

### SetSportId

`func (o *FixturesResponse) SetSportId(v int32)`

SetSportId gets a reference to the given int32 and assigns it to the SportId field.

### GetLast

`func (o *FixturesResponse) GetLast() int64`

GetLast returns the Last field if non-nil, zero value otherwise.

### GetLastOk

`func (o *FixturesResponse) GetLastOk() (int64, bool)`

GetLastOk returns a tuple with the Last field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLast

`func (o *FixturesResponse) HasLast() bool`

HasLast returns a boolean if a field has been set.

### SetLast

`func (o *FixturesResponse) SetLast(v int64)`

SetLast gets a reference to the given int64 and assigns it to the Last field.

### GetLeague

`func (o *FixturesResponse) GetLeague() []FixturesLeague`

GetLeague returns the League field if non-nil, zero value otherwise.

### GetLeagueOk

`func (o *FixturesResponse) GetLeagueOk() ([]FixturesLeague, bool)`

GetLeagueOk returns a tuple with the League field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLeague

`func (o *FixturesResponse) HasLeague() bool`

HasLeague returns a boolean if a field has been set.

### SetLeague

`func (o *FixturesResponse) SetLeague(v []FixturesLeague)`

SetLeague gets a reference to the given []FixturesLeague and assigns it to the League field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


