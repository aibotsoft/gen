# InRunningEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Game Id | [optional] 
**State** | Pointer to **int32** | State of the game.  1 &#x3D; First half in progress,  2 &#x3D; Half time in progress,  3 &#x3D; Second half in progress,  4 &#x3D; End of regular time, 5 &#x3D; First half extra time in progress,  6 &#x3D; Extra time half time in progress,  7 &#x3D; Second half extra time in progress,  8 &#x3D; End of extra time,  9 &#x3D; End of Game,  10 &#x3D; Game is temporary suspended,  11 &#x3D; Penalties in progress  | [optional] 
**Elapsed** | Pointer to **int32** | Elapsed minutes | [optional] 

## Methods

### GetId

`func (o *InRunningEvent) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InRunningEvent) GetIdOk() (int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *InRunningEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *InRunningEvent) SetId(v int64)`

SetId gets a reference to the given int64 and assigns it to the Id field.

### GetState

`func (o *InRunningEvent) GetState() int32`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *InRunningEvent) GetStateOk() (int32, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasState

`func (o *InRunningEvent) HasState() bool`

HasState returns a boolean if a field has been set.

### SetState

`func (o *InRunningEvent) SetState(v int32)`

SetState gets a reference to the given int32 and assigns it to the State field.

### GetElapsed

`func (o *InRunningEvent) GetElapsed() int32`

GetElapsed returns the Elapsed field if non-nil, zero value otherwise.

### GetElapsedOk

`func (o *InRunningEvent) GetElapsedOk() (int32, bool)`

GetElapsedOk returns a tuple with the Elapsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasElapsed

`func (o *InRunningEvent) HasElapsed() bool`

HasElapsed returns a boolean if a field has been set.

### SetElapsed

`func (o *InRunningEvent) SetElapsed(v int32)`

SetElapsed gets a reference to the given int32 and assigns it to the Elapsed field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


