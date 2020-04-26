# ParlayBet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BetId** | Pointer to **int64** | Bet identification | 
**UniqueRequestId** | Pointer to **string** | Unique Request Id | [optional] 
**WagerNumber** | Pointer to **int32** | Wager identification. All bets placed thru the API will have value 1. Website Classic view supports multiple contest(special) bets placement in the same bet slip in that case the bet would have appropriate wager number, as well as all round robin parlay bets. | 
**PlacedAt** | Pointer to [**time.Time**](time.Time.md) | Date time when the bet was placed. | 
**BetStatus** | Pointer to **string** | Bet Status.   ACCEPTED &#x3D; Bet was accepted,   CANCELLED &#x3D; Bet is cancelled as per Pinnacle betting rules,   LOSE &#x3D; The bet is settled as lose,   PENDING_ACCEPTANCE &#x3D; This status is reserved only for live bets. If a live bet is placed during danger zone or live delay is applied, it will be in PENDING_ACCEPTANCE , otherwise in ACCEPTED status. From this status bet can go to ACCEPTED or REJECTED status,   REFUNDED &#x3D; When an event is cancelled or when the bet is settled as push, the bet will have REFUNDED status,   NOT_ACCEPTED &#x3D; Bet was not accepted. Bet can be in this status only if it was previously in PENDING_ACCEPTANCE status,   WON &#x3D; The bet is settled as won  | 
**BetType** | Pointer to **string** |  | [default to PARLAY]
**Win** | Pointer to **float64** | Win amount. | 
**Risk** | Pointer to **float64** | Risk amount. | 
**WinLoss** | Pointer to **NullableFloat64** | Win-Loss for settled bets. | [optional] 
**OddsFormat** | Pointer to [**OddsFormat**](OddsFormat.md) |  | 
**CustomerCommission** | Pointer to **NullableFloat64** | Client’s commission on the bet. | [optional] 
**CancellationReason** | Pointer to [**CancellationReason**](CancellationReason.md) |  | [optional] 
**UpdateSequence** | Pointer to **int64** | Update Sequence | 
**Legs** | Pointer to [**[]ParlayLeg**](ParlayLeg.md) |  | 
**Price** | Pointer to **float64** |  | [optional] 
**FinalPrice** | Pointer to **float64** | Only for settled parlay. Final price may differ in case leg was cancelled or half won | [optional] 

## Methods

### GetBetId

`func (o *ParlayBet) GetBetId() int64`

GetBetId returns the BetId field if non-nil, zero value otherwise.

### GetBetIdOk

`func (o *ParlayBet) GetBetIdOk() (int64, bool)`

GetBetIdOk returns a tuple with the BetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBetId

`func (o *ParlayBet) HasBetId() bool`

HasBetId returns a boolean if a field has been set.

### SetBetId

`func (o *ParlayBet) SetBetId(v int64)`

SetBetId gets a reference to the given int64 and assigns it to the BetId field.

### GetUniqueRequestId

`func (o *ParlayBet) GetUniqueRequestId() string`

GetUniqueRequestId returns the UniqueRequestId field if non-nil, zero value otherwise.

### GetUniqueRequestIdOk

`func (o *ParlayBet) GetUniqueRequestIdOk() (string, bool)`

GetUniqueRequestIdOk returns a tuple with the UniqueRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUniqueRequestId

`func (o *ParlayBet) HasUniqueRequestId() bool`

HasUniqueRequestId returns a boolean if a field has been set.

### SetUniqueRequestId

`func (o *ParlayBet) SetUniqueRequestId(v string)`

SetUniqueRequestId gets a reference to the given string and assigns it to the UniqueRequestId field.

### GetWagerNumber

`func (o *ParlayBet) GetWagerNumber() int32`

GetWagerNumber returns the WagerNumber field if non-nil, zero value otherwise.

### GetWagerNumberOk

`func (o *ParlayBet) GetWagerNumberOk() (int32, bool)`

GetWagerNumberOk returns a tuple with the WagerNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWagerNumber

`func (o *ParlayBet) HasWagerNumber() bool`

HasWagerNumber returns a boolean if a field has been set.

### SetWagerNumber

`func (o *ParlayBet) SetWagerNumber(v int32)`

SetWagerNumber gets a reference to the given int32 and assigns it to the WagerNumber field.

### GetPlacedAt

`func (o *ParlayBet) GetPlacedAt() time.Time`

GetPlacedAt returns the PlacedAt field if non-nil, zero value otherwise.

### GetPlacedAtOk

`func (o *ParlayBet) GetPlacedAtOk() (time.Time, bool)`

GetPlacedAtOk returns a tuple with the PlacedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPlacedAt

`func (o *ParlayBet) HasPlacedAt() bool`

HasPlacedAt returns a boolean if a field has been set.

### SetPlacedAt

`func (o *ParlayBet) SetPlacedAt(v time.Time)`

SetPlacedAt gets a reference to the given time.Time and assigns it to the PlacedAt field.

### GetBetStatus

`func (o *ParlayBet) GetBetStatus() string`

GetBetStatus returns the BetStatus field if non-nil, zero value otherwise.

### GetBetStatusOk

`func (o *ParlayBet) GetBetStatusOk() (string, bool)`

GetBetStatusOk returns a tuple with the BetStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBetStatus

`func (o *ParlayBet) HasBetStatus() bool`

HasBetStatus returns a boolean if a field has been set.

### SetBetStatus

`func (o *ParlayBet) SetBetStatus(v string)`

SetBetStatus gets a reference to the given string and assigns it to the BetStatus field.

### GetBetType

`func (o *ParlayBet) GetBetType() string`

GetBetType returns the BetType field if non-nil, zero value otherwise.

### GetBetTypeOk

`func (o *ParlayBet) GetBetTypeOk() (string, bool)`

GetBetTypeOk returns a tuple with the BetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBetType

`func (o *ParlayBet) HasBetType() bool`

HasBetType returns a boolean if a field has been set.

### SetBetType

`func (o *ParlayBet) SetBetType(v string)`

SetBetType gets a reference to the given string and assigns it to the BetType field.

### GetWin

`func (o *ParlayBet) GetWin() float64`

GetWin returns the Win field if non-nil, zero value otherwise.

### GetWinOk

`func (o *ParlayBet) GetWinOk() (float64, bool)`

GetWinOk returns a tuple with the Win field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWin

`func (o *ParlayBet) HasWin() bool`

HasWin returns a boolean if a field has been set.

### SetWin

`func (o *ParlayBet) SetWin(v float64)`

SetWin gets a reference to the given float64 and assigns it to the Win field.

### GetRisk

`func (o *ParlayBet) GetRisk() float64`

GetRisk returns the Risk field if non-nil, zero value otherwise.

### GetRiskOk

`func (o *ParlayBet) GetRiskOk() (float64, bool)`

GetRiskOk returns a tuple with the Risk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRisk

`func (o *ParlayBet) HasRisk() bool`

HasRisk returns a boolean if a field has been set.

### SetRisk

`func (o *ParlayBet) SetRisk(v float64)`

SetRisk gets a reference to the given float64 and assigns it to the Risk field.

### GetWinLoss

`func (o *ParlayBet) GetWinLoss() NullableFloat64`

GetWinLoss returns the WinLoss field if non-nil, zero value otherwise.

### GetWinLossOk

`func (o *ParlayBet) GetWinLossOk() (NullableFloat64, bool)`

GetWinLossOk returns a tuple with the WinLoss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWinLoss

`func (o *ParlayBet) HasWinLoss() bool`

HasWinLoss returns a boolean if a field has been set.

### SetWinLoss

`func (o *ParlayBet) SetWinLoss(v NullableFloat64)`

SetWinLoss gets a reference to the given NullableFloat64 and assigns it to the WinLoss field.

### SetWinLossExplicitNull

`func (o *ParlayBet) SetWinLossExplicitNull(b bool)`

SetWinLossExplicitNull (un)sets WinLoss to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The WinLoss value is set to nil even if false is passed
### GetOddsFormat

`func (o *ParlayBet) GetOddsFormat() OddsFormat`

GetOddsFormat returns the OddsFormat field if non-nil, zero value otherwise.

### GetOddsFormatOk

`func (o *ParlayBet) GetOddsFormatOk() (OddsFormat, bool)`

GetOddsFormatOk returns a tuple with the OddsFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOddsFormat

`func (o *ParlayBet) HasOddsFormat() bool`

HasOddsFormat returns a boolean if a field has been set.

### SetOddsFormat

`func (o *ParlayBet) SetOddsFormat(v OddsFormat)`

SetOddsFormat gets a reference to the given OddsFormat and assigns it to the OddsFormat field.

### GetCustomerCommission

`func (o *ParlayBet) GetCustomerCommission() NullableFloat64`

GetCustomerCommission returns the CustomerCommission field if non-nil, zero value otherwise.

### GetCustomerCommissionOk

`func (o *ParlayBet) GetCustomerCommissionOk() (NullableFloat64, bool)`

GetCustomerCommissionOk returns a tuple with the CustomerCommission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCustomerCommission

`func (o *ParlayBet) HasCustomerCommission() bool`

HasCustomerCommission returns a boolean if a field has been set.

### SetCustomerCommission

`func (o *ParlayBet) SetCustomerCommission(v NullableFloat64)`

SetCustomerCommission gets a reference to the given NullableFloat64 and assigns it to the CustomerCommission field.

### SetCustomerCommissionExplicitNull

`func (o *ParlayBet) SetCustomerCommissionExplicitNull(b bool)`

SetCustomerCommissionExplicitNull (un)sets CustomerCommission to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CustomerCommission value is set to nil even if false is passed
### GetCancellationReason

`func (o *ParlayBet) GetCancellationReason() CancellationReason`

GetCancellationReason returns the CancellationReason field if non-nil, zero value otherwise.

### GetCancellationReasonOk

`func (o *ParlayBet) GetCancellationReasonOk() (CancellationReason, bool)`

GetCancellationReasonOk returns a tuple with the CancellationReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCancellationReason

`func (o *ParlayBet) HasCancellationReason() bool`

HasCancellationReason returns a boolean if a field has been set.

### SetCancellationReason

`func (o *ParlayBet) SetCancellationReason(v CancellationReason)`

SetCancellationReason gets a reference to the given CancellationReason and assigns it to the CancellationReason field.

### GetUpdateSequence

`func (o *ParlayBet) GetUpdateSequence() int64`

GetUpdateSequence returns the UpdateSequence field if non-nil, zero value otherwise.

### GetUpdateSequenceOk

`func (o *ParlayBet) GetUpdateSequenceOk() (int64, bool)`

GetUpdateSequenceOk returns a tuple with the UpdateSequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUpdateSequence

`func (o *ParlayBet) HasUpdateSequence() bool`

HasUpdateSequence returns a boolean if a field has been set.

### SetUpdateSequence

`func (o *ParlayBet) SetUpdateSequence(v int64)`

SetUpdateSequence gets a reference to the given int64 and assigns it to the UpdateSequence field.

### GetLegs

`func (o *ParlayBet) GetLegs() []ParlayLeg`

GetLegs returns the Legs field if non-nil, zero value otherwise.

### GetLegsOk

`func (o *ParlayBet) GetLegsOk() ([]ParlayLeg, bool)`

GetLegsOk returns a tuple with the Legs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLegs

`func (o *ParlayBet) HasLegs() bool`

HasLegs returns a boolean if a field has been set.

### SetLegs

`func (o *ParlayBet) SetLegs(v []ParlayLeg)`

SetLegs gets a reference to the given []ParlayLeg and assigns it to the Legs field.

### GetPrice

`func (o *ParlayBet) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ParlayBet) GetPriceOk() (float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPrice

`func (o *ParlayBet) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPrice

`func (o *ParlayBet) SetPrice(v float64)`

SetPrice gets a reference to the given float64 and assigns it to the Price field.

### GetFinalPrice

`func (o *ParlayBet) GetFinalPrice() float64`

GetFinalPrice returns the FinalPrice field if non-nil, zero value otherwise.

### GetFinalPriceOk

`func (o *ParlayBet) GetFinalPriceOk() (float64, bool)`

GetFinalPriceOk returns a tuple with the FinalPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFinalPrice

`func (o *ParlayBet) HasFinalPrice() bool`

HasFinalPrice returns a boolean if a field has been set.

### SetFinalPrice

`func (o *ParlayBet) SetFinalPrice(v float64)`

SetFinalPrice gets a reference to the given float64 and assigns it to the FinalPrice field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


