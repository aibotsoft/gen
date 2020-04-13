# SpecialOddsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SportId** | Pointer to **int32** | Id of a sport for which to retrieve the odds. | [optional] 
**Last** | Pointer to **int64** | Used for retrieving changes only on subsequent requests. Provide this value as the Since paramter in subsequent calls to only retrieve changes. | [optional] 
**Leagues** | Pointer to [**[]SpecialOddsLeague**](SpecialOddsLeague.md) | Contains a list of Leagues. | [optional] 

## Methods

### GetSportId

`func (o *SpecialOddsResponse) GetSportId() int32`

GetSportId returns the SportId field if non-nil, zero value otherwise.

### GetSportIdOk

`func (o *SpecialOddsResponse) GetSportIdOk() (int32, bool)`

GetSportIdOk returns a tuple with the SportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSportId

`func (o *SpecialOddsResponse) HasSportId() bool`

HasSportId returns a boolean if a field has been set.

### SetSportId

`func (o *SpecialOddsResponse) SetSportId(v int32)`

SetSportId gets a reference to the given int32 and assigns it to the SportId field.

### GetLast

`func (o *SpecialOddsResponse) GetLast() int64`

GetLast returns the Last field if non-nil, zero value otherwise.

### GetLastOk

`func (o *SpecialOddsResponse) GetLastOk() (int64, bool)`

GetLastOk returns a tuple with the Last field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLast

`func (o *SpecialOddsResponse) HasLast() bool`

HasLast returns a boolean if a field has been set.

### SetLast

`func (o *SpecialOddsResponse) SetLast(v int64)`

SetLast gets a reference to the given int64 and assigns it to the Last field.

### GetLeagues

`func (o *SpecialOddsResponse) GetLeagues() []SpecialOddsLeague`

GetLeagues returns the Leagues field if non-nil, zero value otherwise.

### GetLeaguesOk

`func (o *SpecialOddsResponse) GetLeaguesOk() ([]SpecialOddsLeague, bool)`

GetLeaguesOk returns a tuple with the Leagues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLeagues

`func (o *SpecialOddsResponse) HasLeagues() bool`

HasLeagues returns a boolean if a field has been set.

### SetLeagues

`func (o *SpecialOddsResponse) SetLeagues(v []SpecialOddsLeague)`

SetLeagues gets a reference to the given []SpecialOddsLeague and assigns it to the Leagues field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


