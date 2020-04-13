# League

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | League Id. | [optional] 
**Name** | Pointer to **string** | Name of the league. | [optional] 
**HomeTeamType** | Pointer to **string** | Specifies whether the home team is team1 or team2. You need this information to place a bet. | [optional] 
**HasOfferings** | Pointer to **bool** | Whether the league currently has events or specials. | [optional] 
**Container** | Pointer to **string** | Represents grouping for the league, usually a region/country | [optional] 
**AllowRoundRobins** | Pointer to **bool** | Specifies whether you can place parlay round robins on events in this league. | [optional] 
**LeagueSpecialsCount** | Pointer to **int32** | Indicates how many specials are in the given league. | [optional] 
**EventSpecialsCount** | Pointer to **int32** | Indicates how many game specials are in the given league. | [optional] 
**EventCount** | Pointer to **int32** | Indicates how many events are in the given league. | [optional] 

## Methods

### GetId

`func (o *League) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *League) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *League) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *League) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetName

`func (o *League) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *League) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *League) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *League) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetHomeTeamType

`func (o *League) GetHomeTeamType() string`

GetHomeTeamType returns the HomeTeamType field if non-nil, zero value otherwise.

### GetHomeTeamTypeOk

`func (o *League) GetHomeTeamTypeOk() (string, bool)`

GetHomeTeamTypeOk returns a tuple with the HomeTeamType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHomeTeamType

`func (o *League) HasHomeTeamType() bool`

HasHomeTeamType returns a boolean if a field has been set.

### SetHomeTeamType

`func (o *League) SetHomeTeamType(v string)`

SetHomeTeamType gets a reference to the given string and assigns it to the HomeTeamType field.

### GetHasOfferings

`func (o *League) GetHasOfferings() bool`

GetHasOfferings returns the HasOfferings field if non-nil, zero value otherwise.

### GetHasOfferingsOk

`func (o *League) GetHasOfferingsOk() (bool, bool)`

GetHasOfferingsOk returns a tuple with the HasOfferings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHasOfferings

`func (o *League) HasHasOfferings() bool`

HasHasOfferings returns a boolean if a field has been set.

### SetHasOfferings

`func (o *League) SetHasOfferings(v bool)`

SetHasOfferings gets a reference to the given bool and assigns it to the HasOfferings field.

### GetContainer

`func (o *League) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *League) GetContainerOk() (string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasContainer

`func (o *League) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### SetContainer

`func (o *League) SetContainer(v string)`

SetContainer gets a reference to the given string and assigns it to the Container field.

### GetAllowRoundRobins

`func (o *League) GetAllowRoundRobins() bool`

GetAllowRoundRobins returns the AllowRoundRobins field if non-nil, zero value otherwise.

### GetAllowRoundRobinsOk

`func (o *League) GetAllowRoundRobinsOk() (bool, bool)`

GetAllowRoundRobinsOk returns a tuple with the AllowRoundRobins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAllowRoundRobins

`func (o *League) HasAllowRoundRobins() bool`

HasAllowRoundRobins returns a boolean if a field has been set.

### SetAllowRoundRobins

`func (o *League) SetAllowRoundRobins(v bool)`

SetAllowRoundRobins gets a reference to the given bool and assigns it to the AllowRoundRobins field.

### GetLeagueSpecialsCount

`func (o *League) GetLeagueSpecialsCount() int32`

GetLeagueSpecialsCount returns the LeagueSpecialsCount field if non-nil, zero value otherwise.

### GetLeagueSpecialsCountOk

`func (o *League) GetLeagueSpecialsCountOk() (int32, bool)`

GetLeagueSpecialsCountOk returns a tuple with the LeagueSpecialsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLeagueSpecialsCount

`func (o *League) HasLeagueSpecialsCount() bool`

HasLeagueSpecialsCount returns a boolean if a field has been set.

### SetLeagueSpecialsCount

`func (o *League) SetLeagueSpecialsCount(v int32)`

SetLeagueSpecialsCount gets a reference to the given int32 and assigns it to the LeagueSpecialsCount field.

### GetEventSpecialsCount

`func (o *League) GetEventSpecialsCount() int32`

GetEventSpecialsCount returns the EventSpecialsCount field if non-nil, zero value otherwise.

### GetEventSpecialsCountOk

`func (o *League) GetEventSpecialsCountOk() (int32, bool)`

GetEventSpecialsCountOk returns a tuple with the EventSpecialsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEventSpecialsCount

`func (o *League) HasEventSpecialsCount() bool`

HasEventSpecialsCount returns a boolean if a field has been set.

### SetEventSpecialsCount

`func (o *League) SetEventSpecialsCount(v int32)`

SetEventSpecialsCount gets a reference to the given int32 and assigns it to the EventSpecialsCount field.

### GetEventCount

`func (o *League) GetEventCount() int32`

GetEventCount returns the EventCount field if non-nil, zero value otherwise.

### GetEventCountOk

`func (o *League) GetEventCountOk() (int32, bool)`

GetEventCountOk returns a tuple with the EventCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEventCount

`func (o *League) HasEventCount() bool`

HasEventCount returns a boolean if a field has been set.

### SetEventCount

`func (o *League) SetEventCount(v int32)`

SetEventCount gets a reference to the given int32 and assigns it to the EventCount field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


