# ManualBet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BetId** | Pointer to **int64** | Bet identification | 
**WagerNumber** | Pointer to **int32** | Wager identification. All bets placed thru the API will have value 1. Website Classic view supports multiple contest(special) bets placement in the same bet slip in that case the bet would have appropriate wager number, as well as all round robin parlay bets. | 
**PlacedAt** | Pointer to [**time.Time**](time.Time.md) | Date time when the bet was placed. | 
**BetStatus** | Pointer to **string** | Bet Status.   ACCEPTED &#x3D; Bet was accepted,   CANCELLED &#x3D; Bet is cancelled as per Pinnacle betting rules,   LOSE &#x3D; The bet is settled as lose,   REFUNDED &#x3D; When an event is cancelled or when the bet is settled as push, the bet will have REFUNDED status,   WON &#x3D; The bet is settled as won   | 
**BetType** | Pointer to **string** |  | [default to MANUAL]
**Win** | Pointer to **float64** | Win amount. | 
**Risk** | Pointer to **float64** | Risk amount. | 
**WinLoss** | Pointer to **float64** | Win-Loss for settled bets. | [optional] 
**UpdateSequence** | Pointer to **int64** | Update Sequence | 
**Description** | Pointer to **string** | Manual bet description. | 
**ReferenceBetId** | Pointer to **NullableInt64** | Referenced original bet id. | [optional] 

## Methods

### GetBetId

`func (o *ManualBet) GetBetId() int64`

GetBetId returns the BetId field if non-nil, zero value otherwise.

### GetBetIdOk

`func (o *ManualBet) GetBetIdOk() (int64, bool)`

GetBetIdOk returns a tuple with the BetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBetId

`func (o *ManualBet) HasBetId() bool`

HasBetId returns a boolean if a field has been set.

### SetBetId

`func (o *ManualBet) SetBetId(v int64)`

SetBetId gets a reference to the given int64 and assigns it to the BetId field.

### GetWagerNumber

`func (o *ManualBet) GetWagerNumber() int32`

GetWagerNumber returns the WagerNumber field if non-nil, zero value otherwise.

### GetWagerNumberOk

`func (o *ManualBet) GetWagerNumberOk() (int32, bool)`

GetWagerNumberOk returns a tuple with the WagerNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWagerNumber

`func (o *ManualBet) HasWagerNumber() bool`

HasWagerNumber returns a boolean if a field has been set.

### SetWagerNumber

`func (o *ManualBet) SetWagerNumber(v int32)`

SetWagerNumber gets a reference to the given int32 and assigns it to the WagerNumber field.

### GetPlacedAt

`func (o *ManualBet) GetPlacedAt() time.Time`

GetPlacedAt returns the PlacedAt field if non-nil, zero value otherwise.

### GetPlacedAtOk

`func (o *ManualBet) GetPlacedAtOk() (time.Time, bool)`

GetPlacedAtOk returns a tuple with the PlacedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPlacedAt

`func (o *ManualBet) HasPlacedAt() bool`

HasPlacedAt returns a boolean if a field has been set.

### SetPlacedAt

`func (o *ManualBet) SetPlacedAt(v time.Time)`

SetPlacedAt gets a reference to the given time.Time and assigns it to the PlacedAt field.

### GetBetStatus

`func (o *ManualBet) GetBetStatus() string`

GetBetStatus returns the BetStatus field if non-nil, zero value otherwise.

### GetBetStatusOk

`func (o *ManualBet) GetBetStatusOk() (string, bool)`

GetBetStatusOk returns a tuple with the BetStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBetStatus

`func (o *ManualBet) HasBetStatus() bool`

HasBetStatus returns a boolean if a field has been set.

### SetBetStatus

`func (o *ManualBet) SetBetStatus(v string)`

SetBetStatus gets a reference to the given string and assigns it to the BetStatus field.

### GetBetType

`func (o *ManualBet) GetBetType() string`

GetBetType returns the BetType field if non-nil, zero value otherwise.

### GetBetTypeOk

`func (o *ManualBet) GetBetTypeOk() (string, bool)`

GetBetTypeOk returns a tuple with the BetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBetType

`func (o *ManualBet) HasBetType() bool`

HasBetType returns a boolean if a field has been set.

### SetBetType

`func (o *ManualBet) SetBetType(v string)`

SetBetType gets a reference to the given string and assigns it to the BetType field.

### GetWin

`func (o *ManualBet) GetWin() float64`

GetWin returns the Win field if non-nil, zero value otherwise.

### GetWinOk

`func (o *ManualBet) GetWinOk() (float64, bool)`

GetWinOk returns a tuple with the Win field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWin

`func (o *ManualBet) HasWin() bool`

HasWin returns a boolean if a field has been set.

### SetWin

`func (o *ManualBet) SetWin(v float64)`

SetWin gets a reference to the given float64 and assigns it to the Win field.

### GetRisk

`func (o *ManualBet) GetRisk() float64`

GetRisk returns the Risk field if non-nil, zero value otherwise.

### GetRiskOk

`func (o *ManualBet) GetRiskOk() (float64, bool)`

GetRiskOk returns a tuple with the Risk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRisk

`func (o *ManualBet) HasRisk() bool`

HasRisk returns a boolean if a field has been set.

### SetRisk

`func (o *ManualBet) SetRisk(v float64)`

SetRisk gets a reference to the given float64 and assigns it to the Risk field.

### GetWinLoss

`func (o *ManualBet) GetWinLoss() float64`

GetWinLoss returns the WinLoss field if non-nil, zero value otherwise.

### GetWinLossOk

`func (o *ManualBet) GetWinLossOk() (float64, bool)`

GetWinLossOk returns a tuple with the WinLoss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWinLoss

`func (o *ManualBet) HasWinLoss() bool`

HasWinLoss returns a boolean if a field has been set.

### SetWinLoss

`func (o *ManualBet) SetWinLoss(v float64)`

SetWinLoss gets a reference to the given float64 and assigns it to the WinLoss field.

### GetUpdateSequence

`func (o *ManualBet) GetUpdateSequence() int64`

GetUpdateSequence returns the UpdateSequence field if non-nil, zero value otherwise.

### GetUpdateSequenceOk

`func (o *ManualBet) GetUpdateSequenceOk() (int64, bool)`

GetUpdateSequenceOk returns a tuple with the UpdateSequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUpdateSequence

`func (o *ManualBet) HasUpdateSequence() bool`

HasUpdateSequence returns a boolean if a field has been set.

### SetUpdateSequence

`func (o *ManualBet) SetUpdateSequence(v int64)`

SetUpdateSequence gets a reference to the given int64 and assigns it to the UpdateSequence field.

### GetDescription

`func (o *ManualBet) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ManualBet) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *ManualBet) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *ManualBet) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetReferenceBetId

`func (o *ManualBet) GetReferenceBetId() NullableInt64`

GetReferenceBetId returns the ReferenceBetId field if non-nil, zero value otherwise.

### GetReferenceBetIdOk

`func (o *ManualBet) GetReferenceBetIdOk() (NullableInt64, bool)`

GetReferenceBetIdOk returns a tuple with the ReferenceBetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReferenceBetId

`func (o *ManualBet) HasReferenceBetId() bool`

HasReferenceBetId returns a boolean if a field has been set.

### SetReferenceBetId

`func (o *ManualBet) SetReferenceBetId(v NullableInt64)`

SetReferenceBetId gets a reference to the given NullableInt64 and assigns it to the ReferenceBetId field.

### SetReferenceBetIdExplicitNull

`func (o *ManualBet) SetReferenceBetIdExplicitNull(b bool)`

SetReferenceBetIdExplicitNull (un)sets ReferenceBetId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ReferenceBetId value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


