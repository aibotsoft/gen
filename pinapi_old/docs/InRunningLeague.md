# InRunningLeague

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | League Id | [optional] 
**Events** | Pointer to [**[]InRunningEvent**](InRunningEvent.md) | Events container | [optional] 

## Methods

### GetId

`func (o *InRunningLeague) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InRunningLeague) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *InRunningLeague) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *InRunningLeague) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetEvents

`func (o *InRunningLeague) GetEvents() []InRunningEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *InRunningLeague) GetEventsOk() ([]InRunningEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEvents

`func (o *InRunningLeague) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### SetEvents

`func (o *InRunningLeague) SetEvents(v []InRunningEvent)`

SetEvents gets a reference to the given []InRunningEvent and assigns it to the Events field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


