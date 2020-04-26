# SpecialOddsSpecial

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Special Id. | [optional] 
**MaxBet** | Pointer to **float64** | Maximum bet volume amount. See [How to calculate max risk from the max volume](https://github.com/pinnacleapi/pinnacleapi-documentation/blob/master/FAQ.md#how-to-calculate-max-risk-from-the-max-volume-limits-in-odds) | [optional] 
**ContestantLines** | Pointer to [**[]SpecialOddsContestantLine**](SpecialOddsContestantLine.md) | ContestantLines available for wagering on. | [optional] 

## Methods

### GetId

`func (o *SpecialOddsSpecial) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SpecialOddsSpecial) GetIdOk() (int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *SpecialOddsSpecial) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *SpecialOddsSpecial) SetId(v int64)`

SetId gets a reference to the given int64 and assigns it to the Id field.

### GetMaxBet

`func (o *SpecialOddsSpecial) GetMaxBet() float64`

GetMaxBet returns the MaxBet field if non-nil, zero value otherwise.

### GetMaxBetOk

`func (o *SpecialOddsSpecial) GetMaxBetOk() (float64, bool)`

GetMaxBetOk returns a tuple with the MaxBet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMaxBet

`func (o *SpecialOddsSpecial) HasMaxBet() bool`

HasMaxBet returns a boolean if a field has been set.

### SetMaxBet

`func (o *SpecialOddsSpecial) SetMaxBet(v float64)`

SetMaxBet gets a reference to the given float64 and assigns it to the MaxBet field.

### GetContestantLines

`func (o *SpecialOddsSpecial) GetContestantLines() []SpecialOddsContestantLine`

GetContestantLines returns the ContestantLines field if non-nil, zero value otherwise.

### GetContestantLinesOk

`func (o *SpecialOddsSpecial) GetContestantLinesOk() ([]SpecialOddsContestantLine, bool)`

GetContestantLinesOk returns a tuple with the ContestantLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasContestantLines

`func (o *SpecialOddsSpecial) HasContestantLines() bool`

HasContestantLines returns a boolean if a field has been set.

### SetContestantLines

`func (o *SpecialOddsSpecial) SetContestantLines(v []SpecialOddsContestantLine)`

SetContestantLines gets a reference to the given []SpecialOddsContestantLine and assigns it to the ContestantLines field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


