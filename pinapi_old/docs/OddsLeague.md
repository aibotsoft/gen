# OddsLeague

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | League Id. | [optional] 
**Events** | Pointer to [**[]OddsEvent**](OddsEvent.md) | Contains a list of events. | [optional] 

## Methods

### GetId

`func (o *OddsLeague) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OddsLeague) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *OddsLeague) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *OddsLeague) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetEvents

`func (o *OddsLeague) GetEvents() []OddsEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *OddsLeague) GetEventsOk() ([]OddsEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEvents

`func (o *OddsLeague) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### SetEvents

`func (o *OddsLeague) SetEvents(v []OddsEvent)`

SetEvents gets a reference to the given []OddsEvent and assigns it to the Events field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


