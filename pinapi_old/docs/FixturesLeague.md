# FixturesLeague

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | League ID. | [optional] 
**Name** | Pointer to **string** | League Name. | [optional] 
**Events** | Pointer to [**[]Fixture**](Fixture.md) | Contains a list of events. | [optional] 

## Methods

### GetId

`func (o *FixturesLeague) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FixturesLeague) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *FixturesLeague) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *FixturesLeague) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetName

`func (o *FixturesLeague) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FixturesLeague) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *FixturesLeague) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *FixturesLeague) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetEvents

`func (o *FixturesLeague) GetEvents() []Fixture`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *FixturesLeague) GetEventsOk() ([]Fixture, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEvents

`func (o *FixturesLeague) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### SetEvents

`func (o *FixturesLeague) SetEvents(v []Fixture)`

SetEvents gets a reference to the given []Fixture and assigns it to the Events field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


