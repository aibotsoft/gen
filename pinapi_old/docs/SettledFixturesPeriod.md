# SettledFixturesPeriod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | Pointer to **int32** | This represents the period of the match. For example, for soccer we have 0 (Game), 1 (1st Half) &amp; 2 (2nd Half) | [optional] 
**Status** | Pointer to **int32** | Period settlement status.   1 &#x3D; Event period is settled,  2 &#x3D; Event period is re-settled,  3 &#x3D; Event period is cancelled,  4 &#x3D; Event period is re-settled as cancelled,  5 &#x3D; Event is deleted  | [optional] 
**SettlementId** | Pointer to **int64** | Unique id of the settlement. In case of a re-settlement, a new settlementId and settledAt will be generated. | [optional] 
**SettledAt** | Pointer to [**time.Time**](time.Time.md) | Date and time in UTC when the period was settled. | [optional] 
**Team1Score** | Pointer to **int32** | Team1 score. | [optional] 
**Team2Score** | Pointer to **int32** | Team2 score. | [optional] 
**Team1ScoreSets** | Pointer to **int32** | Team1 sets score. Supported for tennis only. | [optional] 
**Team2ScoreSets** | Pointer to **int32** | Team2 sets score. Supported for tennis only. | [optional] 
**CancellationReason** | Pointer to [**CancellationReasonType**](CancellationReasonType.md) |  | [optional] 

## Methods

### GetNumber

`func (o *SettledFixturesPeriod) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *SettledFixturesPeriod) GetNumberOk() (int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNumber

`func (o *SettledFixturesPeriod) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### SetNumber

`func (o *SettledFixturesPeriod) SetNumber(v int32)`

SetNumber gets a reference to the given int32 and assigns it to the Number field.

### GetStatus

`func (o *SettledFixturesPeriod) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SettledFixturesPeriod) GetStatusOk() (int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStatus

`func (o *SettledFixturesPeriod) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatus

`func (o *SettledFixturesPeriod) SetStatus(v int32)`

SetStatus gets a reference to the given int32 and assigns it to the Status field.

### GetSettlementId

`func (o *SettledFixturesPeriod) GetSettlementId() int64`

GetSettlementId returns the SettlementId field if non-nil, zero value otherwise.

### GetSettlementIdOk

`func (o *SettledFixturesPeriod) GetSettlementIdOk() (int64, bool)`

GetSettlementIdOk returns a tuple with the SettlementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSettlementId

`func (o *SettledFixturesPeriod) HasSettlementId() bool`

HasSettlementId returns a boolean if a field has been set.

### SetSettlementId

`func (o *SettledFixturesPeriod) SetSettlementId(v int64)`

SetSettlementId gets a reference to the given int64 and assigns it to the SettlementId field.

### GetSettledAt

`func (o *SettledFixturesPeriod) GetSettledAt() time.Time`

GetSettledAt returns the SettledAt field if non-nil, zero value otherwise.

### GetSettledAtOk

`func (o *SettledFixturesPeriod) GetSettledAtOk() (time.Time, bool)`

GetSettledAtOk returns a tuple with the SettledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSettledAt

`func (o *SettledFixturesPeriod) HasSettledAt() bool`

HasSettledAt returns a boolean if a field has been set.

### SetSettledAt

`func (o *SettledFixturesPeriod) SetSettledAt(v time.Time)`

SetSettledAt gets a reference to the given time.Time and assigns it to the SettledAt field.

### GetTeam1Score

`func (o *SettledFixturesPeriod) GetTeam1Score() int32`

GetTeam1Score returns the Team1Score field if non-nil, zero value otherwise.

### GetTeam1ScoreOk

`func (o *SettledFixturesPeriod) GetTeam1ScoreOk() (int32, bool)`

GetTeam1ScoreOk returns a tuple with the Team1Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTeam1Score

`func (o *SettledFixturesPeriod) HasTeam1Score() bool`

HasTeam1Score returns a boolean if a field has been set.

### SetTeam1Score

`func (o *SettledFixturesPeriod) SetTeam1Score(v int32)`

SetTeam1Score gets a reference to the given int32 and assigns it to the Team1Score field.

### GetTeam2Score

`func (o *SettledFixturesPeriod) GetTeam2Score() int32`

GetTeam2Score returns the Team2Score field if non-nil, zero value otherwise.

### GetTeam2ScoreOk

`func (o *SettledFixturesPeriod) GetTeam2ScoreOk() (int32, bool)`

GetTeam2ScoreOk returns a tuple with the Team2Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTeam2Score

`func (o *SettledFixturesPeriod) HasTeam2Score() bool`

HasTeam2Score returns a boolean if a field has been set.

### SetTeam2Score

`func (o *SettledFixturesPeriod) SetTeam2Score(v int32)`

SetTeam2Score gets a reference to the given int32 and assigns it to the Team2Score field.

### GetTeam1ScoreSets

`func (o *SettledFixturesPeriod) GetTeam1ScoreSets() int32`

GetTeam1ScoreSets returns the Team1ScoreSets field if non-nil, zero value otherwise.

### GetTeam1ScoreSetsOk

`func (o *SettledFixturesPeriod) GetTeam1ScoreSetsOk() (int32, bool)`

GetTeam1ScoreSetsOk returns a tuple with the Team1ScoreSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTeam1ScoreSets

`func (o *SettledFixturesPeriod) HasTeam1ScoreSets() bool`

HasTeam1ScoreSets returns a boolean if a field has been set.

### SetTeam1ScoreSets

`func (o *SettledFixturesPeriod) SetTeam1ScoreSets(v int32)`

SetTeam1ScoreSets gets a reference to the given int32 and assigns it to the Team1ScoreSets field.

### GetTeam2ScoreSets

`func (o *SettledFixturesPeriod) GetTeam2ScoreSets() int32`

GetTeam2ScoreSets returns the Team2ScoreSets field if non-nil, zero value otherwise.

### GetTeam2ScoreSetsOk

`func (o *SettledFixturesPeriod) GetTeam2ScoreSetsOk() (int32, bool)`

GetTeam2ScoreSetsOk returns a tuple with the Team2ScoreSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTeam2ScoreSets

`func (o *SettledFixturesPeriod) HasTeam2ScoreSets() bool`

HasTeam2ScoreSets returns a boolean if a field has been set.

### SetTeam2ScoreSets

`func (o *SettledFixturesPeriod) SetTeam2ScoreSets(v int32)`

SetTeam2ScoreSets gets a reference to the given int32 and assigns it to the Team2ScoreSets field.

### GetCancellationReason

`func (o *SettledFixturesPeriod) GetCancellationReason() CancellationReasonType`

GetCancellationReason returns the CancellationReason field if non-nil, zero value otherwise.

### GetCancellationReasonOk

`func (o *SettledFixturesPeriod) GetCancellationReasonOk() (CancellationReasonType, bool)`

GetCancellationReasonOk returns a tuple with the CancellationReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCancellationReason

`func (o *SettledFixturesPeriod) HasCancellationReason() bool`

HasCancellationReason returns a boolean if a field has been set.

### SetCancellationReason

`func (o *SettledFixturesPeriod) SetCancellationReason(v CancellationReasonType)`

SetCancellationReason gets a reference to the given CancellationReasonType and assigns it to the CancellationReason field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


