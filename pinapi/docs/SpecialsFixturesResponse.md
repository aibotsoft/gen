# SpecialsFixturesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SportId** | Pointer to **int32** | Id of a sport for which to retrieve the odds. | [optional] 
**Last** | Pointer to **int64** | Used for retrieving changes only on subsequent requests. Provide this value as the Since paramter in subsequent calls to only retrieve changes. | [optional] 
**Leagues** | Pointer to [**[]SpecialsFixturesLeague**](SpecialsFixturesLeague.md) | Contains a list of Leagues. | [optional] 

## Methods

### GetSportId

`func (o *SpecialsFixturesResponse) GetSportId() int32`

GetSportId returns the SportId field if non-nil, zero value otherwise.

### GetSportIdOk

`func (o *SpecialsFixturesResponse) GetSportIdOk() (int32, bool)`

GetSportIdOk returns a tuple with the SportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSportId

`func (o *SpecialsFixturesResponse) HasSportId() bool`

HasSportId returns a boolean if a field has been set.

### SetSportId

`func (o *SpecialsFixturesResponse) SetSportId(v int32)`

SetSportId gets a reference to the given int32 and assigns it to the SportId field.

### GetLast

`func (o *SpecialsFixturesResponse) GetLast() int64`

GetLast returns the Last field if non-nil, zero value otherwise.

### GetLastOk

`func (o *SpecialsFixturesResponse) GetLastOk() (int64, bool)`

GetLastOk returns a tuple with the Last field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLast

`func (o *SpecialsFixturesResponse) HasLast() bool`

HasLast returns a boolean if a field has been set.

### SetLast

`func (o *SpecialsFixturesResponse) SetLast(v int64)`

SetLast gets a reference to the given int64 and assigns it to the Last field.

### GetLeagues

`func (o *SpecialsFixturesResponse) GetLeagues() []SpecialsFixturesLeague`

GetLeagues returns the Leagues field if non-nil, zero value otherwise.

### GetLeaguesOk

`func (o *SpecialsFixturesResponse) GetLeaguesOk() ([]SpecialsFixturesLeague, bool)`

GetLeaguesOk returns a tuple with the Leagues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLeagues

`func (o *SpecialsFixturesResponse) HasLeagues() bool`

HasLeagues returns a boolean if a field has been set.

### SetLeagues

`func (o *SpecialsFixturesResponse) SetLeagues(v []SpecialsFixturesLeague)`

SetLeagues gets a reference to the given []SpecialsFixturesLeague and assigns it to the Leagues field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


