# TeaserLeg

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SportId** | Pointer to **int32** |  | [optional] 
**LegBetType** | Pointer to **string** | Teaser leg type. | [optional] 
**LegBetStatus** | Pointer to **string** | CANCELLED &#x3D; The leg is canceled- the stake on this leg will be transferred to the next one. In this case the leg will be ignored when calculating the winLoss,   LOSE &#x3D; The leg is a loss or a push-lose. When Push-lose happens, the half of the stake on the leg will be pushed to the next leg, and the other half will be a lose. This can happen only when the leg is placed on a quarter points handicap,   PUSH &#x3D; The leg is a push - the stake on this leg will be transferred to the next one. In this case the leg will be ignored when calculating the winLoss,   REFUNDED &#x3D; The leg is refunded - the stake on this leg will be transferred to the next one. In this case the leg will be ignored when calculating the winLoss,   WON &#x3D; The leg is a won or a push-won. When Push-won happens, the half of the stake on the leg will be pushed to the next leg, and the other half is won. This can happen only when the leg is placed on a quarter points handicap    | [optional] 
**LeagueId** | Pointer to **int32** |  | [optional] 
**EventId** | Pointer to **int64** |  | [optional] 
**EventStartTime** | Pointer to [**time.Time**](time.Time.md) | Date time when the event starts. | [optional] 
**Handicap** | Pointer to **float64** |  | [optional] 
**TeamName** | Pointer to **string** |  | [optional] 
**Side** | Pointer to **string** | Side type. | [optional] 
**Team1** | Pointer to **string** |  | [optional] 
**Team2** | Pointer to **string** |  | [optional] 
**PeriodNumber** | Pointer to **int32** |  | [optional] 

## Methods

### GetSportId

`func (o *TeaserLeg) GetSportId() int32`

GetSportId returns the SportId field if non-nil, zero value otherwise.

### GetSportIdOk

`func (o *TeaserLeg) GetSportIdOk() (int32, bool)`

GetSportIdOk returns a tuple with the SportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSportId

`func (o *TeaserLeg) HasSportId() bool`

HasSportId returns a boolean if a field has been set.

### SetSportId

`func (o *TeaserLeg) SetSportId(v int32)`

SetSportId gets a reference to the given int32 and assigns it to the SportId field.

### GetLegBetType

`func (o *TeaserLeg) GetLegBetType() string`

GetLegBetType returns the LegBetType field if non-nil, zero value otherwise.

### GetLegBetTypeOk

`func (o *TeaserLeg) GetLegBetTypeOk() (string, bool)`

GetLegBetTypeOk returns a tuple with the LegBetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLegBetType

`func (o *TeaserLeg) HasLegBetType() bool`

HasLegBetType returns a boolean if a field has been set.

### SetLegBetType

`func (o *TeaserLeg) SetLegBetType(v string)`

SetLegBetType gets a reference to the given string and assigns it to the LegBetType field.

### GetLegBetStatus

`func (o *TeaserLeg) GetLegBetStatus() string`

GetLegBetStatus returns the LegBetStatus field if non-nil, zero value otherwise.

### GetLegBetStatusOk

`func (o *TeaserLeg) GetLegBetStatusOk() (string, bool)`

GetLegBetStatusOk returns a tuple with the LegBetStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLegBetStatus

`func (o *TeaserLeg) HasLegBetStatus() bool`

HasLegBetStatus returns a boolean if a field has been set.

### SetLegBetStatus

`func (o *TeaserLeg) SetLegBetStatus(v string)`

SetLegBetStatus gets a reference to the given string and assigns it to the LegBetStatus field.

### GetLeagueId

`func (o *TeaserLeg) GetLeagueId() int32`

GetLeagueId returns the LeagueId field if non-nil, zero value otherwise.

### GetLeagueIdOk

`func (o *TeaserLeg) GetLeagueIdOk() (int32, bool)`

GetLeagueIdOk returns a tuple with the LeagueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLeagueId

`func (o *TeaserLeg) HasLeagueId() bool`

HasLeagueId returns a boolean if a field has been set.

### SetLeagueId

`func (o *TeaserLeg) SetLeagueId(v int32)`

SetLeagueId gets a reference to the given int32 and assigns it to the LeagueId field.

### GetEventId

`func (o *TeaserLeg) GetEventId() int64`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *TeaserLeg) GetEventIdOk() (int64, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEventId

`func (o *TeaserLeg) HasEventId() bool`

HasEventId returns a boolean if a field has been set.

### SetEventId

`func (o *TeaserLeg) SetEventId(v int64)`

SetEventId gets a reference to the given int64 and assigns it to the EventId field.

### GetEventStartTime

`func (o *TeaserLeg) GetEventStartTime() time.Time`

GetEventStartTime returns the EventStartTime field if non-nil, zero value otherwise.

### GetEventStartTimeOk

`func (o *TeaserLeg) GetEventStartTimeOk() (time.Time, bool)`

GetEventStartTimeOk returns a tuple with the EventStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEventStartTime

`func (o *TeaserLeg) HasEventStartTime() bool`

HasEventStartTime returns a boolean if a field has been set.

### SetEventStartTime

`func (o *TeaserLeg) SetEventStartTime(v time.Time)`

SetEventStartTime gets a reference to the given time.Time and assigns it to the EventStartTime field.

### GetHandicap

`func (o *TeaserLeg) GetHandicap() float64`

GetHandicap returns the Handicap field if non-nil, zero value otherwise.

### GetHandicapOk

`func (o *TeaserLeg) GetHandicapOk() (float64, bool)`

GetHandicapOk returns a tuple with the Handicap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHandicap

`func (o *TeaserLeg) HasHandicap() bool`

HasHandicap returns a boolean if a field has been set.

### SetHandicap

`func (o *TeaserLeg) SetHandicap(v float64)`

SetHandicap gets a reference to the given float64 and assigns it to the Handicap field.

### GetTeamName

`func (o *TeaserLeg) GetTeamName() string`

GetTeamName returns the TeamName field if non-nil, zero value otherwise.

### GetTeamNameOk

`func (o *TeaserLeg) GetTeamNameOk() (string, bool)`

GetTeamNameOk returns a tuple with the TeamName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTeamName

`func (o *TeaserLeg) HasTeamName() bool`

HasTeamName returns a boolean if a field has been set.

### SetTeamName

`func (o *TeaserLeg) SetTeamName(v string)`

SetTeamName gets a reference to the given string and assigns it to the TeamName field.

### GetSide

`func (o *TeaserLeg) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *TeaserLeg) GetSideOk() (string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSide

`func (o *TeaserLeg) HasSide() bool`

HasSide returns a boolean if a field has been set.

### SetSide

`func (o *TeaserLeg) SetSide(v string)`

SetSide gets a reference to the given string and assigns it to the Side field.

### GetTeam1

`func (o *TeaserLeg) GetTeam1() string`

GetTeam1 returns the Team1 field if non-nil, zero value otherwise.

### GetTeam1Ok

`func (o *TeaserLeg) GetTeam1Ok() (string, bool)`

GetTeam1Ok returns a tuple with the Team1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTeam1

`func (o *TeaserLeg) HasTeam1() bool`

HasTeam1 returns a boolean if a field has been set.

### SetTeam1

`func (o *TeaserLeg) SetTeam1(v string)`

SetTeam1 gets a reference to the given string and assigns it to the Team1 field.

### GetTeam2

`func (o *TeaserLeg) GetTeam2() string`

GetTeam2 returns the Team2 field if non-nil, zero value otherwise.

### GetTeam2Ok

`func (o *TeaserLeg) GetTeam2Ok() (string, bool)`

GetTeam2Ok returns a tuple with the Team2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTeam2

`func (o *TeaserLeg) HasTeam2() bool`

HasTeam2 returns a boolean if a field has been set.

### SetTeam2

`func (o *TeaserLeg) SetTeam2(v string)`

SetTeam2 gets a reference to the given string and assigns it to the Team2 field.

### GetPeriodNumber

`func (o *TeaserLeg) GetPeriodNumber() int32`

GetPeriodNumber returns the PeriodNumber field if non-nil, zero value otherwise.

### GetPeriodNumberOk

`func (o *TeaserLeg) GetPeriodNumberOk() (int32, bool)`

GetPeriodNumberOk returns a tuple with the PeriodNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPeriodNumber

`func (o *TeaserLeg) HasPeriodNumber() bool`

HasPeriodNumber returns a boolean if a field has been set.

### SetPeriodNumber

`func (o *TeaserLeg) SetPeriodNumber(v int32)`

SetPeriodNumber gets a reference to the given int32 and assigns it to the PeriodNumber field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


