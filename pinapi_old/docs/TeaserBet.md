# TeaserBet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BetId** | Pointer to **int64** | Bet identification | 
**UniqueRequestId** | Pointer to **string** | Unique Request Id | [optional] 
**WagerNumber** | Pointer to **int32** | Wager identification. All bets placed thru the API will have value 1. Website Classic view supports multiple contest(special) bets placement in the same bet slip in that case the bet would have appropriate wager number, as well as all round robin parlay bets. | 
**PlacedAt** | Pointer to [**time.Time**](time.Time.md) | Date time when the bet was placed. | 
**BetStatus** | Pointer to **string** | Bet Status.   ACCEPTED &#x3D; Bet was accepted,   CANCELLED &#x3D; Bet is cancelled as per Pinnacle betting rules,   LOSE &#x3D; The bet is settled as lose,   REFUNDED &#x3D; When an event is cancelled or when the bet is settled as push, the bet will have REFUNDED status,   WON &#x3D; The bet is settled as won   | 
**BetType** | Pointer to **string** |  | [default to TEASER]
**Win** | Pointer to **float64** | Win amount. | 
**Risk** | Pointer to **float64** | Risk amount. | 
**WinLoss** | Pointer to **float64** | Win-Loss for settled bets. | [optional] 
**OddsFormat** | Pointer to [**OddsFormat**](OddsFormat.md) |  | 
**CustomerCommission** | Pointer to **float64** | Client’s commission on the bet. | [optional] 
**CancellationReason** | Pointer to [**CancellationReason**](CancellationReason.md) |  | [optional] 
**UpdateSequence** | Pointer to **int64** | Update Sequence | 
**TeaserName** | Pointer to **string** |  | 
**IsSameEventOnly** | Pointer to **bool** |  | 
**MinPicks** | Pointer to **float64** |  | 
**MaxPicks** | Pointer to **float64** |  | 
**Price** | Pointer to **float64** | Populated for all teaser bets and will be the original price at the time of the placement. | [optional] 
**FinalPrice** | Pointer to **float64** | Only for settled parlay. Final price may differ in case leg was cancelled or half won. | [optional] 
**TeaserId** | Pointer to **float32** | Reference to the teaser id | [optional] 
**TeaserGroupId** | Pointer to **float32** | Reference to the teaser group id | [optional] 
**Legs** | Pointer to [**[]TeaserLeg**](TeaserLeg.md) |  | 

## Methods

### GetBetId

`func (o *TeaserBet) GetBetId() int64`

GetBetId returns the BetId field if non-nil, zero value otherwise.

### GetBetIdOk

`func (o *TeaserBet) GetBetIdOk() (int64, bool)`

GetBetIdOk returns a tuple with the BetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBetId

`func (o *TeaserBet) HasBetId() bool`

HasBetId returns a boolean if a field has been set.

### SetBetId

`func (o *TeaserBet) SetBetId(v int64)`

SetBetId gets a reference to the given int64 and assigns it to the BetId field.

### GetUniqueRequestId

`func (o *TeaserBet) GetUniqueRequestId() string`

GetUniqueRequestId returns the UniqueRequestId field if non-nil, zero value otherwise.

### GetUniqueRequestIdOk

`func (o *TeaserBet) GetUniqueRequestIdOk() (string, bool)`

GetUniqueRequestIdOk returns a tuple with the UniqueRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUniqueRequestId

`func (o *TeaserBet) HasUniqueRequestId() bool`

HasUniqueRequestId returns a boolean if a field has been set.

### SetUniqueRequestId

`func (o *TeaserBet) SetUniqueRequestId(v string)`

SetUniqueRequestId gets a reference to the given string and assigns it to the UniqueRequestId field.

### GetWagerNumber

`func (o *TeaserBet) GetWagerNumber() int32`

GetWagerNumber returns the WagerNumber field if non-nil, zero value otherwise.

### GetWagerNumberOk

`func (o *TeaserBet) GetWagerNumberOk() (int32, bool)`

GetWagerNumberOk returns a tuple with the WagerNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWagerNumber

`func (o *TeaserBet) HasWagerNumber() bool`

HasWagerNumber returns a boolean if a field has been set.

### SetWagerNumber

`func (o *TeaserBet) SetWagerNumber(v int32)`

SetWagerNumber gets a reference to the given int32 and assigns it to the WagerNumber field.

### GetPlacedAt

`func (o *TeaserBet) GetPlacedAt() time.Time`

GetPlacedAt returns the PlacedAt field if non-nil, zero value otherwise.

### GetPlacedAtOk

`func (o *TeaserBet) GetPlacedAtOk() (time.Time, bool)`

GetPlacedAtOk returns a tuple with the PlacedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPlacedAt

`func (o *TeaserBet) HasPlacedAt() bool`

HasPlacedAt returns a boolean if a field has been set.

### SetPlacedAt

`func (o *TeaserBet) SetPlacedAt(v time.Time)`

SetPlacedAt gets a reference to the given time.Time and assigns it to the PlacedAt field.

### GetBetStatus

`func (o *TeaserBet) GetBetStatus() string`

GetBetStatus returns the BetStatus field if non-nil, zero value otherwise.

### GetBetStatusOk

`func (o *TeaserBet) GetBetStatusOk() (string, bool)`

GetBetStatusOk returns a tuple with the BetStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBetStatus

`func (o *TeaserBet) HasBetStatus() bool`

HasBetStatus returns a boolean if a field has been set.

### SetBetStatus

`func (o *TeaserBet) SetBetStatus(v string)`

SetBetStatus gets a reference to the given string and assigns it to the BetStatus field.

### GetBetType

`func (o *TeaserBet) GetBetType() string`

GetBetType returns the BetType field if non-nil, zero value otherwise.

### GetBetTypeOk

`func (o *TeaserBet) GetBetTypeOk() (string, bool)`

GetBetTypeOk returns a tuple with the BetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBetType

`func (o *TeaserBet) HasBetType() bool`

HasBetType returns a boolean if a field has been set.

### SetBetType

`func (o *TeaserBet) SetBetType(v string)`

SetBetType gets a reference to the given string and assigns it to the BetType field.

### GetWin

`func (o *TeaserBet) GetWin() float64`

GetWin returns the Win field if non-nil, zero value otherwise.

### GetWinOk

`func (o *TeaserBet) GetWinOk() (float64, bool)`

GetWinOk returns a tuple with the Win field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWin

`func (o *TeaserBet) HasWin() bool`

HasWin returns a boolean if a field has been set.

### SetWin

`func (o *TeaserBet) SetWin(v float64)`

SetWin gets a reference to the given float64 and assigns it to the Win field.

### GetRisk

`func (o *TeaserBet) GetRisk() float64`

GetRisk returns the Risk field if non-nil, zero value otherwise.

### GetRiskOk

`func (o *TeaserBet) GetRiskOk() (float64, bool)`

GetRiskOk returns a tuple with the Risk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRisk

`func (o *TeaserBet) HasRisk() bool`

HasRisk returns a boolean if a field has been set.

### SetRisk

`func (o *TeaserBet) SetRisk(v float64)`

SetRisk gets a reference to the given float64 and assigns it to the Risk field.

### GetWinLoss

`func (o *TeaserBet) GetWinLoss() float64`

GetWinLoss returns the WinLoss field if non-nil, zero value otherwise.

### GetWinLossOk

`func (o *TeaserBet) GetWinLossOk() (float64, bool)`

GetWinLossOk returns a tuple with the WinLoss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWinLoss

`func (o *TeaserBet) HasWinLoss() bool`

HasWinLoss returns a boolean if a field has been set.

### SetWinLoss

`func (o *TeaserBet) SetWinLoss(v float64)`

SetWinLoss gets a reference to the given float64 and assigns it to the WinLoss field.

### GetOddsFormat

`func (o *TeaserBet) GetOddsFormat() OddsFormat`

GetOddsFormat returns the OddsFormat field if non-nil, zero value otherwise.

### GetOddsFormatOk

`func (o *TeaserBet) GetOddsFormatOk() (OddsFormat, bool)`

GetOddsFormatOk returns a tuple with the OddsFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOddsFormat

`func (o *TeaserBet) HasOddsFormat() bool`

HasOddsFormat returns a boolean if a field has been set.

### SetOddsFormat

`func (o *TeaserBet) SetOddsFormat(v OddsFormat)`

SetOddsFormat gets a reference to the given OddsFormat and assigns it to the OddsFormat field.

### GetCustomerCommission

`func (o *TeaserBet) GetCustomerCommission() float64`

GetCustomerCommission returns the CustomerCommission field if non-nil, zero value otherwise.

### GetCustomerCommissionOk

`func (o *TeaserBet) GetCustomerCommissionOk() (float64, bool)`

GetCustomerCommissionOk returns a tuple with the CustomerCommission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCustomerCommission

`func (o *TeaserBet) HasCustomerCommission() bool`

HasCustomerCommission returns a boolean if a field has been set.

### SetCustomerCommission

`func (o *TeaserBet) SetCustomerCommission(v float64)`

SetCustomerCommission gets a reference to the given float64 and assigns it to the CustomerCommission field.

### GetCancellationReason

`func (o *TeaserBet) GetCancellationReason() CancellationReason`

GetCancellationReason returns the CancellationReason field if non-nil, zero value otherwise.

### GetCancellationReasonOk

`func (o *TeaserBet) GetCancellationReasonOk() (CancellationReason, bool)`

GetCancellationReasonOk returns a tuple with the CancellationReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCancellationReason

`func (o *TeaserBet) HasCancellationReason() bool`

HasCancellationReason returns a boolean if a field has been set.

### SetCancellationReason

`func (o *TeaserBet) SetCancellationReason(v CancellationReason)`

SetCancellationReason gets a reference to the given CancellationReason and assigns it to the CancellationReason field.

### GetUpdateSequence

`func (o *TeaserBet) GetUpdateSequence() int64`

GetUpdateSequence returns the UpdateSequence field if non-nil, zero value otherwise.

### GetUpdateSequenceOk

`func (o *TeaserBet) GetUpdateSequenceOk() (int64, bool)`

GetUpdateSequenceOk returns a tuple with the UpdateSequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUpdateSequence

`func (o *TeaserBet) HasUpdateSequence() bool`

HasUpdateSequence returns a boolean if a field has been set.

### SetUpdateSequence

`func (o *TeaserBet) SetUpdateSequence(v int64)`

SetUpdateSequence gets a reference to the given int64 and assigns it to the UpdateSequence field.

### GetTeaserName

`func (o *TeaserBet) GetTeaserName() string`

GetTeaserName returns the TeaserName field if non-nil, zero value otherwise.

### GetTeaserNameOk

`func (o *TeaserBet) GetTeaserNameOk() (string, bool)`

GetTeaserNameOk returns a tuple with the TeaserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTeaserName

`func (o *TeaserBet) HasTeaserName() bool`

HasTeaserName returns a boolean if a field has been set.

### SetTeaserName

`func (o *TeaserBet) SetTeaserName(v string)`

SetTeaserName gets a reference to the given string and assigns it to the TeaserName field.

### GetIsSameEventOnly

`func (o *TeaserBet) GetIsSameEventOnly() bool`

GetIsSameEventOnly returns the IsSameEventOnly field if non-nil, zero value otherwise.

### GetIsSameEventOnlyOk

`func (o *TeaserBet) GetIsSameEventOnlyOk() (bool, bool)`

GetIsSameEventOnlyOk returns a tuple with the IsSameEventOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsSameEventOnly

`func (o *TeaserBet) HasIsSameEventOnly() bool`

HasIsSameEventOnly returns a boolean if a field has been set.

### SetIsSameEventOnly

`func (o *TeaserBet) SetIsSameEventOnly(v bool)`

SetIsSameEventOnly gets a reference to the given bool and assigns it to the IsSameEventOnly field.

### GetMinPicks

`func (o *TeaserBet) GetMinPicks() float64`

GetMinPicks returns the MinPicks field if non-nil, zero value otherwise.

### GetMinPicksOk

`func (o *TeaserBet) GetMinPicksOk() (float64, bool)`

GetMinPicksOk returns a tuple with the MinPicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMinPicks

`func (o *TeaserBet) HasMinPicks() bool`

HasMinPicks returns a boolean if a field has been set.

### SetMinPicks

`func (o *TeaserBet) SetMinPicks(v float64)`

SetMinPicks gets a reference to the given float64 and assigns it to the MinPicks field.

### GetMaxPicks

`func (o *TeaserBet) GetMaxPicks() float64`

GetMaxPicks returns the MaxPicks field if non-nil, zero value otherwise.

### GetMaxPicksOk

`func (o *TeaserBet) GetMaxPicksOk() (float64, bool)`

GetMaxPicksOk returns a tuple with the MaxPicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMaxPicks

`func (o *TeaserBet) HasMaxPicks() bool`

HasMaxPicks returns a boolean if a field has been set.

### SetMaxPicks

`func (o *TeaserBet) SetMaxPicks(v float64)`

SetMaxPicks gets a reference to the given float64 and assigns it to the MaxPicks field.

### GetPrice

`func (o *TeaserBet) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *TeaserBet) GetPriceOk() (float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPrice

`func (o *TeaserBet) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPrice

`func (o *TeaserBet) SetPrice(v float64)`

SetPrice gets a reference to the given float64 and assigns it to the Price field.

### GetFinalPrice

`func (o *TeaserBet) GetFinalPrice() float64`

GetFinalPrice returns the FinalPrice field if non-nil, zero value otherwise.

### GetFinalPriceOk

`func (o *TeaserBet) GetFinalPriceOk() (float64, bool)`

GetFinalPriceOk returns a tuple with the FinalPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFinalPrice

`func (o *TeaserBet) HasFinalPrice() bool`

HasFinalPrice returns a boolean if a field has been set.

### SetFinalPrice

`func (o *TeaserBet) SetFinalPrice(v float64)`

SetFinalPrice gets a reference to the given float64 and assigns it to the FinalPrice field.

### GetTeaserId

`func (o *TeaserBet) GetTeaserId() float32`

GetTeaserId returns the TeaserId field if non-nil, zero value otherwise.

### GetTeaserIdOk

`func (o *TeaserBet) GetTeaserIdOk() (float32, bool)`

GetTeaserIdOk returns a tuple with the TeaserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTeaserId

`func (o *TeaserBet) HasTeaserId() bool`

HasTeaserId returns a boolean if a field has been set.

### SetTeaserId

`func (o *TeaserBet) SetTeaserId(v float32)`

SetTeaserId gets a reference to the given float32 and assigns it to the TeaserId field.

### GetTeaserGroupId

`func (o *TeaserBet) GetTeaserGroupId() float32`

GetTeaserGroupId returns the TeaserGroupId field if non-nil, zero value otherwise.

### GetTeaserGroupIdOk

`func (o *TeaserBet) GetTeaserGroupIdOk() (float32, bool)`

GetTeaserGroupIdOk returns a tuple with the TeaserGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTeaserGroupId

`func (o *TeaserBet) HasTeaserGroupId() bool`

HasTeaserGroupId returns a boolean if a field has been set.

### SetTeaserGroupId

`func (o *TeaserBet) SetTeaserGroupId(v float32)`

SetTeaserGroupId gets a reference to the given float32 and assigns it to the TeaserGroupId field.

### GetLegs

`func (o *TeaserBet) GetLegs() []TeaserLeg`

GetLegs returns the Legs field if non-nil, zero value otherwise.

### GetLegsOk

`func (o *TeaserBet) GetLegsOk() ([]TeaserLeg, bool)`

GetLegsOk returns a tuple with the Legs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLegs

`func (o *TeaserBet) HasLegs() bool`

HasLegs returns a boolean if a field has been set.

### SetLegs

`func (o *TeaserBet) SetLegs(v []TeaserLeg)`

SetLegs gets a reference to the given []TeaserLeg and assigns it to the Legs field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


