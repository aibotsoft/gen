# SettledSpecialsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SportId** | Pointer to **int32** | Id of a sport for which to retrieve the odds. | [optional] 
**Last** | Pointer to **int64** | Last index for the settled fixture | [optional] 
**Leagues** | Pointer to [**[]SettledSpecialsLeague**](SettledSpecialsLeague.md) | List of Leagues. | [optional] 

## Methods

### GetSportId

`func (o *SettledSpecialsResponse) GetSportId() int32`

GetSportId returns the SportId field if non-nil, zero value otherwise.

### GetSportIdOk

`func (o *SettledSpecialsResponse) GetSportIdOk() (int32, bool)`

GetSportIdOk returns a tuple with the SportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSportId

`func (o *SettledSpecialsResponse) HasSportId() bool`

HasSportId returns a boolean if a field has been set.

### SetSportId

`func (o *SettledSpecialsResponse) SetSportId(v int32)`

SetSportId gets a reference to the given int32 and assigns it to the SportId field.

### GetLast

`func (o *SettledSpecialsResponse) GetLast() int64`

GetLast returns the Last field if non-nil, zero value otherwise.

### GetLastOk

`func (o *SettledSpecialsResponse) GetLastOk() (int64, bool)`

GetLastOk returns a tuple with the Last field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLast

`func (o *SettledSpecialsResponse) HasLast() bool`

HasLast returns a boolean if a field has been set.

### SetLast

`func (o *SettledSpecialsResponse) SetLast(v int64)`

SetLast gets a reference to the given int64 and assigns it to the Last field.

### GetLeagues

`func (o *SettledSpecialsResponse) GetLeagues() []SettledSpecialsLeague`

GetLeagues returns the Leagues field if non-nil, zero value otherwise.

### GetLeaguesOk

`func (o *SettledSpecialsResponse) GetLeaguesOk() ([]SettledSpecialsLeague, bool)`

GetLeaguesOk returns a tuple with the Leagues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLeagues

`func (o *SettledSpecialsResponse) HasLeagues() bool`

HasLeagues returns a boolean if a field has been set.

### SetLeagues

`func (o *SettledSpecialsResponse) SetLeagues(v []SettledSpecialsLeague)`

SetLeagues gets a reference to the given []SettledSpecialsLeague and assigns it to the Leagues field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


