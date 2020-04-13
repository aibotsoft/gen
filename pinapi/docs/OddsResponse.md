# OddsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SportId** | Pointer to **int32** | Same as requested sport Id. | [optional] 
**Last** | Pointer to **int64** | Use this value for the subsequent requests for since query parameter to get just the changes since previous response. | [optional] 
**Leagues** | Pointer to [**[]OddsLeague**](OddsLeague.md) | Contains a list of Leagues. | [optional] 

## Methods

### GetSportId

`func (o *OddsResponse) GetSportId() int32`

GetSportId returns the SportId field if non-nil, zero value otherwise.

### GetSportIdOk

`func (o *OddsResponse) GetSportIdOk() (int32, bool)`

GetSportIdOk returns a tuple with the SportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSportId

`func (o *OddsResponse) HasSportId() bool`

HasSportId returns a boolean if a field has been set.

### SetSportId

`func (o *OddsResponse) SetSportId(v int32)`

SetSportId gets a reference to the given int32 and assigns it to the SportId field.

### GetLast

`func (o *OddsResponse) GetLast() int64`

GetLast returns the Last field if non-nil, zero value otherwise.

### GetLastOk

`func (o *OddsResponse) GetLastOk() (int64, bool)`

GetLastOk returns a tuple with the Last field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLast

`func (o *OddsResponse) HasLast() bool`

HasLast returns a boolean if a field has been set.

### SetLast

`func (o *OddsResponse) SetLast(v int64)`

SetLast gets a reference to the given int64 and assigns it to the Last field.

### GetLeagues

`func (o *OddsResponse) GetLeagues() []OddsLeague`

GetLeagues returns the Leagues field if non-nil, zero value otherwise.

### GetLeaguesOk

`func (o *OddsResponse) GetLeaguesOk() ([]OddsLeague, bool)`

GetLeaguesOk returns a tuple with the Leagues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLeagues

`func (o *OddsResponse) HasLeagues() bool`

HasLeagues returns a boolean if a field has been set.

### SetLeagues

`func (o *OddsResponse) SetLeagues(v []OddsLeague)`

SetLeagues gets a reference to the given []OddsLeague and assigns it to the Leagues field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


