# Fixture

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Event id. | [optional] 
**ParentId** | Pointer to **int64** | If event is linked to another event, parentId will be populated.  Live event would have pre game event as parent id. | [optional] 
**Starts** | Pointer to [**time.Time**](time.Time.md) | Start time of the event in UTC. | [optional] 
**Home** | Pointer to **string** | Home team name. | [optional] 
**Away** | Pointer to **string** | Away team name. | [optional] 
**RotNum** | Pointer to **string** | Team1 rotation number. Please note that in the next version of /fixtures, rotNum property will be decommissioned. ParentId can be used instead to group the related events. | [optional] 
**LiveStatus** | Pointer to **int32** | Indicates live status of the event.   0 &#x3D; No live betting will be offered on this event,  1 &#x3D; Live betting event,  2 &#x3D; Live betting will be offered on this match, but on a different event. Please note that [pre-game and live events are different](https://github.com/pinnacleapi/pinnacleapi-documentation/blob/master/FAQ.md#how-to-find-associated-events) .  | [optional] 
**HomePitcher** | Pointer to **string** | Home team pitcher. Only for Baseball. | [optional] 
**AwayPitcher** | Pointer to **string** | Away team pitcher. Only for Baseball. | [optional] 
**Status** | Pointer to **string** | This is deprecated parameter, please check period&#39;s &#x60;status&#x60; in the &#x60;/odds&#x60; endpoint to see if it&#39;s open for betting.   O &#x3D; This is the starting status of a game.    H &#x3D; This status indicates that the lines are temporarily unavailable for betting,   I &#x3D; This status indicates that one or more lines have a red circle (lower maximum bet amount).  | [optional] 
**ParlayRestriction** | Pointer to **int32** |  Parlay status of the event.   0 &#x3D; Allowed to parlay, without restrictions,  1 &#x3D; Not allowed to parlay this event,  2 &#x3D; Allowed to parlay with the restrictions. You cannot have more than one leg from the same event in the parlay. All events with the same rotation number are treated as same event.  | [optional] 
**AltTeaser** | Pointer to **bool** | Whether an event is offer with alternative teaser points. Events with alternative teaser points may vary from teaser definition. | [optional] 
**ResultingUnit** | Pointer to **string** | Specifies based on what the event will be resulted, e.g. Corners, Bookings   | [optional] 

## Methods

### GetId

`func (o *Fixture) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Fixture) GetIdOk() (int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *Fixture) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *Fixture) SetId(v int64)`

SetId gets a reference to the given int64 and assigns it to the Id field.

### GetParentId

`func (o *Fixture) GetParentId() int64`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *Fixture) GetParentIdOk() (int64, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasParentId

`func (o *Fixture) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentId

`func (o *Fixture) SetParentId(v int64)`

SetParentId gets a reference to the given int64 and assigns it to the ParentId field.

### GetStarts

`func (o *Fixture) GetStarts() time.Time`

GetStarts returns the Starts field if non-nil, zero value otherwise.

### GetStartsOk

`func (o *Fixture) GetStartsOk() (time.Time, bool)`

GetStartsOk returns a tuple with the Starts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStarts

`func (o *Fixture) HasStarts() bool`

HasStarts returns a boolean if a field has been set.

### SetStarts

`func (o *Fixture) SetStarts(v time.Time)`

SetStarts gets a reference to the given time.Time and assigns it to the Starts field.

### GetHome

`func (o *Fixture) GetHome() string`

GetHome returns the Home field if non-nil, zero value otherwise.

### GetHomeOk

`func (o *Fixture) GetHomeOk() (string, bool)`

GetHomeOk returns a tuple with the Home field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHome

`func (o *Fixture) HasHome() bool`

HasHome returns a boolean if a field has been set.

### SetHome

`func (o *Fixture) SetHome(v string)`

SetHome gets a reference to the given string and assigns it to the Home field.

### GetAway

`func (o *Fixture) GetAway() string`

GetAway returns the Away field if non-nil, zero value otherwise.

### GetAwayOk

`func (o *Fixture) GetAwayOk() (string, bool)`

GetAwayOk returns a tuple with the Away field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAway

`func (o *Fixture) HasAway() bool`

HasAway returns a boolean if a field has been set.

### SetAway

`func (o *Fixture) SetAway(v string)`

SetAway gets a reference to the given string and assigns it to the Away field.

### GetRotNum

`func (o *Fixture) GetRotNum() string`

GetRotNum returns the RotNum field if non-nil, zero value otherwise.

### GetRotNumOk

`func (o *Fixture) GetRotNumOk() (string, bool)`

GetRotNumOk returns a tuple with the RotNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRotNum

`func (o *Fixture) HasRotNum() bool`

HasRotNum returns a boolean if a field has been set.

### SetRotNum

`func (o *Fixture) SetRotNum(v string)`

SetRotNum gets a reference to the given string and assigns it to the RotNum field.

### GetLiveStatus

`func (o *Fixture) GetLiveStatus() int32`

GetLiveStatus returns the LiveStatus field if non-nil, zero value otherwise.

### GetLiveStatusOk

`func (o *Fixture) GetLiveStatusOk() (int32, bool)`

GetLiveStatusOk returns a tuple with the LiveStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLiveStatus

`func (o *Fixture) HasLiveStatus() bool`

HasLiveStatus returns a boolean if a field has been set.

### SetLiveStatus

`func (o *Fixture) SetLiveStatus(v int32)`

SetLiveStatus gets a reference to the given int32 and assigns it to the LiveStatus field.

### GetHomePitcher

`func (o *Fixture) GetHomePitcher() string`

GetHomePitcher returns the HomePitcher field if non-nil, zero value otherwise.

### GetHomePitcherOk

`func (o *Fixture) GetHomePitcherOk() (string, bool)`

GetHomePitcherOk returns a tuple with the HomePitcher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHomePitcher

`func (o *Fixture) HasHomePitcher() bool`

HasHomePitcher returns a boolean if a field has been set.

### SetHomePitcher

`func (o *Fixture) SetHomePitcher(v string)`

SetHomePitcher gets a reference to the given string and assigns it to the HomePitcher field.

### GetAwayPitcher

`func (o *Fixture) GetAwayPitcher() string`

GetAwayPitcher returns the AwayPitcher field if non-nil, zero value otherwise.

### GetAwayPitcherOk

`func (o *Fixture) GetAwayPitcherOk() (string, bool)`

GetAwayPitcherOk returns a tuple with the AwayPitcher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAwayPitcher

`func (o *Fixture) HasAwayPitcher() bool`

HasAwayPitcher returns a boolean if a field has been set.

### SetAwayPitcher

`func (o *Fixture) SetAwayPitcher(v string)`

SetAwayPitcher gets a reference to the given string and assigns it to the AwayPitcher field.

### GetStatus

`func (o *Fixture) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Fixture) GetStatusOk() (string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStatus

`func (o *Fixture) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatus

`func (o *Fixture) SetStatus(v string)`

SetStatus gets a reference to the given string and assigns it to the Status field.

### GetParlayRestriction

`func (o *Fixture) GetParlayRestriction() int32`

GetParlayRestriction returns the ParlayRestriction field if non-nil, zero value otherwise.

### GetParlayRestrictionOk

`func (o *Fixture) GetParlayRestrictionOk() (int32, bool)`

GetParlayRestrictionOk returns a tuple with the ParlayRestriction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasParlayRestriction

`func (o *Fixture) HasParlayRestriction() bool`

HasParlayRestriction returns a boolean if a field has been set.

### SetParlayRestriction

`func (o *Fixture) SetParlayRestriction(v int32)`

SetParlayRestriction gets a reference to the given int32 and assigns it to the ParlayRestriction field.

### GetAltTeaser

`func (o *Fixture) GetAltTeaser() bool`

GetAltTeaser returns the AltTeaser field if non-nil, zero value otherwise.

### GetAltTeaserOk

`func (o *Fixture) GetAltTeaserOk() (bool, bool)`

GetAltTeaserOk returns a tuple with the AltTeaser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAltTeaser

`func (o *Fixture) HasAltTeaser() bool`

HasAltTeaser returns a boolean if a field has been set.

### SetAltTeaser

`func (o *Fixture) SetAltTeaser(v bool)`

SetAltTeaser gets a reference to the given bool and assigns it to the AltTeaser field.

### GetResultingUnit

`func (o *Fixture) GetResultingUnit() string`

GetResultingUnit returns the ResultingUnit field if non-nil, zero value otherwise.

### GetResultingUnitOk

`func (o *Fixture) GetResultingUnitOk() (string, bool)`

GetResultingUnitOk returns a tuple with the ResultingUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResultingUnit

`func (o *Fixture) HasResultingUnit() bool`

HasResultingUnit returns a boolean if a field has been set.

### SetResultingUnit

`func (o *Fixture) SetResultingUnit(v string)`

SetResultingUnit gets a reference to the given string and assigns it to the ResultingUnit field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


