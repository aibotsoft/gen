# ParlayLeg

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SportId** | Pointer to **int32** |  | [optional] 
**LegBetType** | Pointer to **string** | Parlay leg type. | [optional] 
**LegBetStatus** | Pointer to **string** | Parlay Leg status. CANCELLED &#x3D; The leg is canceled- the stake on this leg will be transferred to the next one. In this case the leg will be ignored when calculating the winLoss,   LOSE &#x3D; The leg is a loss or a push-lose. When Push-lose happens, the half of the stake on the leg will be pushed to the next leg, and the other half will be a lose. This can happen only when the leg is placed on a quarter points handicap,   PUSH &#x3D; The leg is a push - the stake on this leg will be transferred to the next one. In this case the leg will be ignored when calculating the winLoss,   REFUNDED &#x3D; The leg is refunded - the stake on this leg will be transferred to the next one. In this case the leg will be ignored when calculating the winLoss,   WON &#x3D; The leg is a won or a push-won. When Push-won happens, the half of the stake on the leg will be pushed to the next leg, and the other half is won. This can happen only when the leg is placed on a quarter points handicap   | [optional] 
**LeagueId** | Pointer to **int32** |  | [optional] 
**EventId** | Pointer to **int64** |  | [optional] 
**EventStartTime** | Pointer to [**time.Time**](time.Time.md) | Date time when the event starts. | [optional] 
**Handicap** | Pointer to **NullableFloat64** |  | [optional] 
**Price** | Pointer to **float64** |  | [optional] 
**TeamName** | Pointer to **string** |  | [optional] 
**Side** | Pointer to **NullableString** | Side type. | [optional] 
**Pitcher1** | Pointer to **NullableString** |  | [optional] 
**Pitcher2** | Pointer to **NullableString** |  | [optional] 
**Pitcher1MustStart** | Pointer to **bool** |  | [optional] 
**Pitcher2MustStart** | Pointer to **bool** |  | [optional] 
**Team1** | Pointer to **string** | Wellington Phoenix | [optional] 
**Team2** | Pointer to **string** | Adelaide United | [optional] 
**PeriodNumber** | Pointer to **int32** |  | [optional] 
**FtTeam1Score** | Pointer to **NullableFloat64** | Full time team 1 score | [optional] 
**FtTeam2Score** | Pointer to **NullableFloat64** | Full time team 2 score | [optional] 
**PTeam1Score** | Pointer to **NullableFloat64** | End of period team 1 score. If the bet was placed on Game period (periodNumber &#x3D;0) , this will be null | [optional] 
**PTeam2Score** | Pointer to **NullableFloat64** | End of period team 2 score. If the bet was placed on Game period (periodNumber &#x3D;0) , this will be null | [optional] 
**CancellationReason** | Pointer to [**CancellationReason**](CancellationReason.md) |  | [optional] 

## Methods

### GetSportId

`func (o *ParlayLeg) GetSportId() int32`

GetSportId returns the SportId field if non-nil, zero value otherwise.

### GetSportIdOk

`func (o *ParlayLeg) GetSportIdOk() (int32, bool)`

GetSportIdOk returns a tuple with the SportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSportId

`func (o *ParlayLeg) HasSportId() bool`

HasSportId returns a boolean if a field has been set.

### SetSportId

`func (o *ParlayLeg) SetSportId(v int32)`

SetSportId gets a reference to the given int32 and assigns it to the SportId field.

### GetLegBetType

`func (o *ParlayLeg) GetLegBetType() string`

GetLegBetType returns the LegBetType field if non-nil, zero value otherwise.

### GetLegBetTypeOk

`func (o *ParlayLeg) GetLegBetTypeOk() (string, bool)`

GetLegBetTypeOk returns a tuple with the LegBetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLegBetType

`func (o *ParlayLeg) HasLegBetType() bool`

HasLegBetType returns a boolean if a field has been set.

### SetLegBetType

`func (o *ParlayLeg) SetLegBetType(v string)`

SetLegBetType gets a reference to the given string and assigns it to the LegBetType field.

### GetLegBetStatus

`func (o *ParlayLeg) GetLegBetStatus() string`

GetLegBetStatus returns the LegBetStatus field if non-nil, zero value otherwise.

### GetLegBetStatusOk

`func (o *ParlayLeg) GetLegBetStatusOk() (string, bool)`

GetLegBetStatusOk returns a tuple with the LegBetStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLegBetStatus

`func (o *ParlayLeg) HasLegBetStatus() bool`

HasLegBetStatus returns a boolean if a field has been set.

### SetLegBetStatus

`func (o *ParlayLeg) SetLegBetStatus(v string)`

SetLegBetStatus gets a reference to the given string and assigns it to the LegBetStatus field.

### GetLeagueId

`func (o *ParlayLeg) GetLeagueId() int32`

GetLeagueId returns the LeagueId field if non-nil, zero value otherwise.

### GetLeagueIdOk

`func (o *ParlayLeg) GetLeagueIdOk() (int32, bool)`

GetLeagueIdOk returns a tuple with the LeagueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLeagueId

`func (o *ParlayLeg) HasLeagueId() bool`

HasLeagueId returns a boolean if a field has been set.

### SetLeagueId

`func (o *ParlayLeg) SetLeagueId(v int32)`

SetLeagueId gets a reference to the given int32 and assigns it to the LeagueId field.

### GetEventId

`func (o *ParlayLeg) GetEventId() int64`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *ParlayLeg) GetEventIdOk() (int64, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEventId

`func (o *ParlayLeg) HasEventId() bool`

HasEventId returns a boolean if a field has been set.

### SetEventId

`func (o *ParlayLeg) SetEventId(v int64)`

SetEventId gets a reference to the given int64 and assigns it to the EventId field.

### GetEventStartTime

`func (o *ParlayLeg) GetEventStartTime() time.Time`

GetEventStartTime returns the EventStartTime field if non-nil, zero value otherwise.

### GetEventStartTimeOk

`func (o *ParlayLeg) GetEventStartTimeOk() (time.Time, bool)`

GetEventStartTimeOk returns a tuple with the EventStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEventStartTime

`func (o *ParlayLeg) HasEventStartTime() bool`

HasEventStartTime returns a boolean if a field has been set.

### SetEventStartTime

`func (o *ParlayLeg) SetEventStartTime(v time.Time)`

SetEventStartTime gets a reference to the given time.Time and assigns it to the EventStartTime field.

### GetHandicap

`func (o *ParlayLeg) GetHandicap() NullableFloat64`

GetHandicap returns the Handicap field if non-nil, zero value otherwise.

### GetHandicapOk

`func (o *ParlayLeg) GetHandicapOk() (NullableFloat64, bool)`

GetHandicapOk returns a tuple with the Handicap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHandicap

`func (o *ParlayLeg) HasHandicap() bool`

HasHandicap returns a boolean if a field has been set.

### SetHandicap

`func (o *ParlayLeg) SetHandicap(v NullableFloat64)`

SetHandicap gets a reference to the given NullableFloat64 and assigns it to the Handicap field.

### SetHandicapExplicitNull

`func (o *ParlayLeg) SetHandicapExplicitNull(b bool)`

SetHandicapExplicitNull (un)sets Handicap to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Handicap value is set to nil even if false is passed
### GetPrice

`func (o *ParlayLeg) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ParlayLeg) GetPriceOk() (float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPrice

`func (o *ParlayLeg) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPrice

`func (o *ParlayLeg) SetPrice(v float64)`

SetPrice gets a reference to the given float64 and assigns it to the Price field.

### GetTeamName

`func (o *ParlayLeg) GetTeamName() string`

GetTeamName returns the TeamName field if non-nil, zero value otherwise.

### GetTeamNameOk

`func (o *ParlayLeg) GetTeamNameOk() (string, bool)`

GetTeamNameOk returns a tuple with the TeamName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTeamName

`func (o *ParlayLeg) HasTeamName() bool`

HasTeamName returns a boolean if a field has been set.

### SetTeamName

`func (o *ParlayLeg) SetTeamName(v string)`

SetTeamName gets a reference to the given string and assigns it to the TeamName field.

### GetSide

`func (o *ParlayLeg) GetSide() NullableString`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *ParlayLeg) GetSideOk() (NullableString, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSide

`func (o *ParlayLeg) HasSide() bool`

HasSide returns a boolean if a field has been set.

### SetSide

`func (o *ParlayLeg) SetSide(v NullableString)`

SetSide gets a reference to the given NullableString and assigns it to the Side field.

### SetSideExplicitNull

`func (o *ParlayLeg) SetSideExplicitNull(b bool)`

SetSideExplicitNull (un)sets Side to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Side value is set to nil even if false is passed
### GetPitcher1

`func (o *ParlayLeg) GetPitcher1() NullableString`

GetPitcher1 returns the Pitcher1 field if non-nil, zero value otherwise.

### GetPitcher1Ok

`func (o *ParlayLeg) GetPitcher1Ok() (NullableString, bool)`

GetPitcher1Ok returns a tuple with the Pitcher1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPitcher1

`func (o *ParlayLeg) HasPitcher1() bool`

HasPitcher1 returns a boolean if a field has been set.

### SetPitcher1

`func (o *ParlayLeg) SetPitcher1(v NullableString)`

SetPitcher1 gets a reference to the given NullableString and assigns it to the Pitcher1 field.

### SetPitcher1ExplicitNull

`func (o *ParlayLeg) SetPitcher1ExplicitNull(b bool)`

SetPitcher1ExplicitNull (un)sets Pitcher1 to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Pitcher1 value is set to nil even if false is passed
### GetPitcher2

`func (o *ParlayLeg) GetPitcher2() NullableString`

GetPitcher2 returns the Pitcher2 field if non-nil, zero value otherwise.

### GetPitcher2Ok

`func (o *ParlayLeg) GetPitcher2Ok() (NullableString, bool)`

GetPitcher2Ok returns a tuple with the Pitcher2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPitcher2

`func (o *ParlayLeg) HasPitcher2() bool`

HasPitcher2 returns a boolean if a field has been set.

### SetPitcher2

`func (o *ParlayLeg) SetPitcher2(v NullableString)`

SetPitcher2 gets a reference to the given NullableString and assigns it to the Pitcher2 field.

### SetPitcher2ExplicitNull

`func (o *ParlayLeg) SetPitcher2ExplicitNull(b bool)`

SetPitcher2ExplicitNull (un)sets Pitcher2 to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Pitcher2 value is set to nil even if false is passed
### GetPitcher1MustStart

`func (o *ParlayLeg) GetPitcher1MustStart() bool`

GetPitcher1MustStart returns the Pitcher1MustStart field if non-nil, zero value otherwise.

### GetPitcher1MustStartOk

`func (o *ParlayLeg) GetPitcher1MustStartOk() (bool, bool)`

GetPitcher1MustStartOk returns a tuple with the Pitcher1MustStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPitcher1MustStart

`func (o *ParlayLeg) HasPitcher1MustStart() bool`

HasPitcher1MustStart returns a boolean if a field has been set.

### SetPitcher1MustStart

`func (o *ParlayLeg) SetPitcher1MustStart(v bool)`

SetPitcher1MustStart gets a reference to the given bool and assigns it to the Pitcher1MustStart field.

### GetPitcher2MustStart

`func (o *ParlayLeg) GetPitcher2MustStart() bool`

GetPitcher2MustStart returns the Pitcher2MustStart field if non-nil, zero value otherwise.

### GetPitcher2MustStartOk

`func (o *ParlayLeg) GetPitcher2MustStartOk() (bool, bool)`

GetPitcher2MustStartOk returns a tuple with the Pitcher2MustStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPitcher2MustStart

`func (o *ParlayLeg) HasPitcher2MustStart() bool`

HasPitcher2MustStart returns a boolean if a field has been set.

### SetPitcher2MustStart

`func (o *ParlayLeg) SetPitcher2MustStart(v bool)`

SetPitcher2MustStart gets a reference to the given bool and assigns it to the Pitcher2MustStart field.

### GetTeam1

`func (o *ParlayLeg) GetTeam1() string`

GetTeam1 returns the Team1 field if non-nil, zero value otherwise.

### GetTeam1Ok

`func (o *ParlayLeg) GetTeam1Ok() (string, bool)`

GetTeam1Ok returns a tuple with the Team1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTeam1

`func (o *ParlayLeg) HasTeam1() bool`

HasTeam1 returns a boolean if a field has been set.

### SetTeam1

`func (o *ParlayLeg) SetTeam1(v string)`

SetTeam1 gets a reference to the given string and assigns it to the Team1 field.

### GetTeam2

`func (o *ParlayLeg) GetTeam2() string`

GetTeam2 returns the Team2 field if non-nil, zero value otherwise.

### GetTeam2Ok

`func (o *ParlayLeg) GetTeam2Ok() (string, bool)`

GetTeam2Ok returns a tuple with the Team2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTeam2

`func (o *ParlayLeg) HasTeam2() bool`

HasTeam2 returns a boolean if a field has been set.

### SetTeam2

`func (o *ParlayLeg) SetTeam2(v string)`

SetTeam2 gets a reference to the given string and assigns it to the Team2 field.

### GetPeriodNumber

`func (o *ParlayLeg) GetPeriodNumber() int32`

GetPeriodNumber returns the PeriodNumber field if non-nil, zero value otherwise.

### GetPeriodNumberOk

`func (o *ParlayLeg) GetPeriodNumberOk() (int32, bool)`

GetPeriodNumberOk returns a tuple with the PeriodNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPeriodNumber

`func (o *ParlayLeg) HasPeriodNumber() bool`

HasPeriodNumber returns a boolean if a field has been set.

### SetPeriodNumber

`func (o *ParlayLeg) SetPeriodNumber(v int32)`

SetPeriodNumber gets a reference to the given int32 and assigns it to the PeriodNumber field.

### GetFtTeam1Score

`func (o *ParlayLeg) GetFtTeam1Score() NullableFloat64`

GetFtTeam1Score returns the FtTeam1Score field if non-nil, zero value otherwise.

### GetFtTeam1ScoreOk

`func (o *ParlayLeg) GetFtTeam1ScoreOk() (NullableFloat64, bool)`

GetFtTeam1ScoreOk returns a tuple with the FtTeam1Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFtTeam1Score

`func (o *ParlayLeg) HasFtTeam1Score() bool`

HasFtTeam1Score returns a boolean if a field has been set.

### SetFtTeam1Score

`func (o *ParlayLeg) SetFtTeam1Score(v NullableFloat64)`

SetFtTeam1Score gets a reference to the given NullableFloat64 and assigns it to the FtTeam1Score field.

### SetFtTeam1ScoreExplicitNull

`func (o *ParlayLeg) SetFtTeam1ScoreExplicitNull(b bool)`

SetFtTeam1ScoreExplicitNull (un)sets FtTeam1Score to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The FtTeam1Score value is set to nil even if false is passed
### GetFtTeam2Score

`func (o *ParlayLeg) GetFtTeam2Score() NullableFloat64`

GetFtTeam2Score returns the FtTeam2Score field if non-nil, zero value otherwise.

### GetFtTeam2ScoreOk

`func (o *ParlayLeg) GetFtTeam2ScoreOk() (NullableFloat64, bool)`

GetFtTeam2ScoreOk returns a tuple with the FtTeam2Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFtTeam2Score

`func (o *ParlayLeg) HasFtTeam2Score() bool`

HasFtTeam2Score returns a boolean if a field has been set.

### SetFtTeam2Score

`func (o *ParlayLeg) SetFtTeam2Score(v NullableFloat64)`

SetFtTeam2Score gets a reference to the given NullableFloat64 and assigns it to the FtTeam2Score field.

### SetFtTeam2ScoreExplicitNull

`func (o *ParlayLeg) SetFtTeam2ScoreExplicitNull(b bool)`

SetFtTeam2ScoreExplicitNull (un)sets FtTeam2Score to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The FtTeam2Score value is set to nil even if false is passed
### GetPTeam1Score

`func (o *ParlayLeg) GetPTeam1Score() NullableFloat64`

GetPTeam1Score returns the PTeam1Score field if non-nil, zero value otherwise.

### GetPTeam1ScoreOk

`func (o *ParlayLeg) GetPTeam1ScoreOk() (NullableFloat64, bool)`

GetPTeam1ScoreOk returns a tuple with the PTeam1Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPTeam1Score

`func (o *ParlayLeg) HasPTeam1Score() bool`

HasPTeam1Score returns a boolean if a field has been set.

### SetPTeam1Score

`func (o *ParlayLeg) SetPTeam1Score(v NullableFloat64)`

SetPTeam1Score gets a reference to the given NullableFloat64 and assigns it to the PTeam1Score field.

### SetPTeam1ScoreExplicitNull

`func (o *ParlayLeg) SetPTeam1ScoreExplicitNull(b bool)`

SetPTeam1ScoreExplicitNull (un)sets PTeam1Score to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PTeam1Score value is set to nil even if false is passed
### GetPTeam2Score

`func (o *ParlayLeg) GetPTeam2Score() NullableFloat64`

GetPTeam2Score returns the PTeam2Score field if non-nil, zero value otherwise.

### GetPTeam2ScoreOk

`func (o *ParlayLeg) GetPTeam2ScoreOk() (NullableFloat64, bool)`

GetPTeam2ScoreOk returns a tuple with the PTeam2Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPTeam2Score

`func (o *ParlayLeg) HasPTeam2Score() bool`

HasPTeam2Score returns a boolean if a field has been set.

### SetPTeam2Score

`func (o *ParlayLeg) SetPTeam2Score(v NullableFloat64)`

SetPTeam2Score gets a reference to the given NullableFloat64 and assigns it to the PTeam2Score field.

### SetPTeam2ScoreExplicitNull

`func (o *ParlayLeg) SetPTeam2ScoreExplicitNull(b bool)`

SetPTeam2ScoreExplicitNull (un)sets PTeam2Score to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PTeam2Score value is set to nil even if false is passed
### GetCancellationReason

`func (o *ParlayLeg) GetCancellationReason() CancellationReason`

GetCancellationReason returns the CancellationReason field if non-nil, zero value otherwise.

### GetCancellationReasonOk

`func (o *ParlayLeg) GetCancellationReasonOk() (CancellationReason, bool)`

GetCancellationReasonOk returns a tuple with the CancellationReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCancellationReason

`func (o *ParlayLeg) HasCancellationReason() bool`

HasCancellationReason returns a boolean if a field has been set.

### SetCancellationReason

`func (o *ParlayLeg) SetCancellationReason(v CancellationReason)`

SetCancellationReason gets a reference to the given CancellationReason and assigns it to the CancellationReason field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


