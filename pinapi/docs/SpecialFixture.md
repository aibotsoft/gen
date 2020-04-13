# SpecialFixture

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Unique Id | [optional] 
**BetType** | Pointer to **string** | The type [MULTI_WAY_HEAD_TO_HEAD, SPREAD, OVER_UNDER] | [optional] 
**Name** | Pointer to **string** | Name of the special. | [optional] 
**Date** | Pointer to [**time.Time**](time.Time.md) | Date of the special in UTC. | [optional] 
**Cutoff** | Pointer to [**time.Time**](time.Time.md) | Wagering cutoff date in UTC. | [optional] 
**Category** | Pointer to **string** | The category that the special falls under. | [optional] 
**Units** | Pointer to **string** | Measurment in the context of the special. This is applicable to specials bet type spead and over/under. In a hockey special this could be goals. | [optional] 
**Status** | Pointer to **string** | Status of the Special    O &#x3D; This is the starting status. It means that the lines  are open for betting,    H &#x3D; This status indicates that the lines are temporarily unavailable  for betting,    I &#x3D; This status indicates that one or more lines have a red circle  (a lower maximum bet amount)  | [optional] 
**Event** | Pointer to [**SpecialsFixturesEvent**](SpecialsFixturesEvent.md) |  | [optional] 
**Contestants** | Pointer to [**[]SpecialsFixturesContestant**](SpecialsFixturesContestant.md) | ContestantLines available for wagering. | [optional] 
**LiveStatus** | Pointer to **int32** | When a special is linked to an event, we will return live status of the event, otherwise it will be 0.  0 &#x3D; No live betting will be offered on this event,  1 &#x3D; Live betting event,  2 &#x3D; Live betting will be offered on this match, but on a different event.   Please note that live delay is applied when placing bets on special with LiveStatus&#x3D;1   | [optional] 

## Methods

### GetId

`func (o *SpecialFixture) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SpecialFixture) GetIdOk() (int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *SpecialFixture) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *SpecialFixture) SetId(v int64)`

SetId gets a reference to the given int64 and assigns it to the Id field.

### GetBetType

`func (o *SpecialFixture) GetBetType() string`

GetBetType returns the BetType field if non-nil, zero value otherwise.

### GetBetTypeOk

`func (o *SpecialFixture) GetBetTypeOk() (string, bool)`

GetBetTypeOk returns a tuple with the BetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBetType

`func (o *SpecialFixture) HasBetType() bool`

HasBetType returns a boolean if a field has been set.

### SetBetType

`func (o *SpecialFixture) SetBetType(v string)`

SetBetType gets a reference to the given string and assigns it to the BetType field.

### GetName

`func (o *SpecialFixture) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SpecialFixture) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *SpecialFixture) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *SpecialFixture) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetDate

`func (o *SpecialFixture) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *SpecialFixture) GetDateOk() (time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDate

`func (o *SpecialFixture) HasDate() bool`

HasDate returns a boolean if a field has been set.

### SetDate

`func (o *SpecialFixture) SetDate(v time.Time)`

SetDate gets a reference to the given time.Time and assigns it to the Date field.

### GetCutoff

`func (o *SpecialFixture) GetCutoff() time.Time`

GetCutoff returns the Cutoff field if non-nil, zero value otherwise.

### GetCutoffOk

`func (o *SpecialFixture) GetCutoffOk() (time.Time, bool)`

GetCutoffOk returns a tuple with the Cutoff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCutoff

`func (o *SpecialFixture) HasCutoff() bool`

HasCutoff returns a boolean if a field has been set.

### SetCutoff

`func (o *SpecialFixture) SetCutoff(v time.Time)`

SetCutoff gets a reference to the given time.Time and assigns it to the Cutoff field.

### GetCategory

`func (o *SpecialFixture) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *SpecialFixture) GetCategoryOk() (string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCategory

`func (o *SpecialFixture) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategory

`func (o *SpecialFixture) SetCategory(v string)`

SetCategory gets a reference to the given string and assigns it to the Category field.

### GetUnits

`func (o *SpecialFixture) GetUnits() string`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *SpecialFixture) GetUnitsOk() (string, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUnits

`func (o *SpecialFixture) HasUnits() bool`

HasUnits returns a boolean if a field has been set.

### SetUnits

`func (o *SpecialFixture) SetUnits(v string)`

SetUnits gets a reference to the given string and assigns it to the Units field.

### GetStatus

`func (o *SpecialFixture) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SpecialFixture) GetStatusOk() (string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStatus

`func (o *SpecialFixture) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatus

`func (o *SpecialFixture) SetStatus(v string)`

SetStatus gets a reference to the given string and assigns it to the Status field.

### GetEvent

`func (o *SpecialFixture) GetEvent() SpecialsFixturesEvent`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *SpecialFixture) GetEventOk() (SpecialsFixturesEvent, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEvent

`func (o *SpecialFixture) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### SetEvent

`func (o *SpecialFixture) SetEvent(v SpecialsFixturesEvent)`

SetEvent gets a reference to the given SpecialsFixturesEvent and assigns it to the Event field.

### GetContestants

`func (o *SpecialFixture) GetContestants() []SpecialsFixturesContestant`

GetContestants returns the Contestants field if non-nil, zero value otherwise.

### GetContestantsOk

`func (o *SpecialFixture) GetContestantsOk() ([]SpecialsFixturesContestant, bool)`

GetContestantsOk returns a tuple with the Contestants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasContestants

`func (o *SpecialFixture) HasContestants() bool`

HasContestants returns a boolean if a field has been set.

### SetContestants

`func (o *SpecialFixture) SetContestants(v []SpecialsFixturesContestant)`

SetContestants gets a reference to the given []SpecialsFixturesContestant and assigns it to the Contestants field.

### GetLiveStatus

`func (o *SpecialFixture) GetLiveStatus() int32`

GetLiveStatus returns the LiveStatus field if non-nil, zero value otherwise.

### GetLiveStatusOk

`func (o *SpecialFixture) GetLiveStatusOk() (int32, bool)`

GetLiveStatusOk returns a tuple with the LiveStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLiveStatus

`func (o *SpecialFixture) HasLiveStatus() bool`

HasLiveStatus returns a boolean if a field has been set.

### SetLiveStatus

`func (o *SpecialFixture) SetLiveStatus(v int32)`

SetLiveStatus gets a reference to the given int32 and assigns it to the LiveStatus field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


