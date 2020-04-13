# GetBetsByTypeResponseV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MoreAvailable** | Pointer to **bool** | Whether there are more pages available. | [optional] 
**PageSize** | Pointer to **int32** | Page size. Default is 1000. | [optional] 
**FromRecord** | Pointer to **int32** | Starting record number of the result set. Records start at zero | [optional] 
**ToRecord** | Pointer to **int32** | Ending record number of the result set. | [optional] 
**StraightBets** | Pointer to [**[]StraightBetV3**](StraightBetV3.md) | A collection of placed straight bets. | [optional] 
**ParlayBets** | Pointer to [**[]ParlayBet**](ParlayBet.md) | A collection of placed parlay bets. | [optional] 
**TeaserBets** | Pointer to [**[]TeaserBet**](TeaserBet.md) | A collection of placed teaser bets. | [optional] 
**SpecialBets** | Pointer to [**[]SpecialBet**](SpecialBet.md) | A collection of placed special bets. | [optional] 
**ManualBets** | Pointer to [**[]ManualBet**](ManualBet.md) | A collection of placed manual bets. | [optional] 

## Methods

### GetMoreAvailable

`func (o *GetBetsByTypeResponseV3) GetMoreAvailable() bool`

GetMoreAvailable returns the MoreAvailable field if non-nil, zero value otherwise.

### GetMoreAvailableOk

`func (o *GetBetsByTypeResponseV3) GetMoreAvailableOk() (bool, bool)`

GetMoreAvailableOk returns a tuple with the MoreAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMoreAvailable

`func (o *GetBetsByTypeResponseV3) HasMoreAvailable() bool`

HasMoreAvailable returns a boolean if a field has been set.

### SetMoreAvailable

`func (o *GetBetsByTypeResponseV3) SetMoreAvailable(v bool)`

SetMoreAvailable gets a reference to the given bool and assigns it to the MoreAvailable field.

### GetPageSize

`func (o *GetBetsByTypeResponseV3) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *GetBetsByTypeResponseV3) GetPageSizeOk() (int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPageSize

`func (o *GetBetsByTypeResponseV3) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### SetPageSize

`func (o *GetBetsByTypeResponseV3) SetPageSize(v int32)`

SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.

### GetFromRecord

`func (o *GetBetsByTypeResponseV3) GetFromRecord() int32`

GetFromRecord returns the FromRecord field if non-nil, zero value otherwise.

### GetFromRecordOk

`func (o *GetBetsByTypeResponseV3) GetFromRecordOk() (int32, bool)`

GetFromRecordOk returns a tuple with the FromRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFromRecord

`func (o *GetBetsByTypeResponseV3) HasFromRecord() bool`

HasFromRecord returns a boolean if a field has been set.

### SetFromRecord

`func (o *GetBetsByTypeResponseV3) SetFromRecord(v int32)`

SetFromRecord gets a reference to the given int32 and assigns it to the FromRecord field.

### GetToRecord

`func (o *GetBetsByTypeResponseV3) GetToRecord() int32`

GetToRecord returns the ToRecord field if non-nil, zero value otherwise.

### GetToRecordOk

`func (o *GetBetsByTypeResponseV3) GetToRecordOk() (int32, bool)`

GetToRecordOk returns a tuple with the ToRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasToRecord

`func (o *GetBetsByTypeResponseV3) HasToRecord() bool`

HasToRecord returns a boolean if a field has been set.

### SetToRecord

`func (o *GetBetsByTypeResponseV3) SetToRecord(v int32)`

SetToRecord gets a reference to the given int32 and assigns it to the ToRecord field.

### GetStraightBets

`func (o *GetBetsByTypeResponseV3) GetStraightBets() []StraightBetV3`

GetStraightBets returns the StraightBets field if non-nil, zero value otherwise.

### GetStraightBetsOk

`func (o *GetBetsByTypeResponseV3) GetStraightBetsOk() ([]StraightBetV3, bool)`

GetStraightBetsOk returns a tuple with the StraightBets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStraightBets

`func (o *GetBetsByTypeResponseV3) HasStraightBets() bool`

HasStraightBets returns a boolean if a field has been set.

### SetStraightBets

`func (o *GetBetsByTypeResponseV3) SetStraightBets(v []StraightBetV3)`

SetStraightBets gets a reference to the given []StraightBetV3 and assigns it to the StraightBets field.

### GetParlayBets

`func (o *GetBetsByTypeResponseV3) GetParlayBets() []ParlayBet`

GetParlayBets returns the ParlayBets field if non-nil, zero value otherwise.

### GetParlayBetsOk

`func (o *GetBetsByTypeResponseV3) GetParlayBetsOk() ([]ParlayBet, bool)`

GetParlayBetsOk returns a tuple with the ParlayBets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasParlayBets

`func (o *GetBetsByTypeResponseV3) HasParlayBets() bool`

HasParlayBets returns a boolean if a field has been set.

### SetParlayBets

`func (o *GetBetsByTypeResponseV3) SetParlayBets(v []ParlayBet)`

SetParlayBets gets a reference to the given []ParlayBet and assigns it to the ParlayBets field.

### GetTeaserBets

`func (o *GetBetsByTypeResponseV3) GetTeaserBets() []TeaserBet`

GetTeaserBets returns the TeaserBets field if non-nil, zero value otherwise.

### GetTeaserBetsOk

`func (o *GetBetsByTypeResponseV3) GetTeaserBetsOk() ([]TeaserBet, bool)`

GetTeaserBetsOk returns a tuple with the TeaserBets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTeaserBets

`func (o *GetBetsByTypeResponseV3) HasTeaserBets() bool`

HasTeaserBets returns a boolean if a field has been set.

### SetTeaserBets

`func (o *GetBetsByTypeResponseV3) SetTeaserBets(v []TeaserBet)`

SetTeaserBets gets a reference to the given []TeaserBet and assigns it to the TeaserBets field.

### GetSpecialBets

`func (o *GetBetsByTypeResponseV3) GetSpecialBets() []SpecialBet`

GetSpecialBets returns the SpecialBets field if non-nil, zero value otherwise.

### GetSpecialBetsOk

`func (o *GetBetsByTypeResponseV3) GetSpecialBetsOk() ([]SpecialBet, bool)`

GetSpecialBetsOk returns a tuple with the SpecialBets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSpecialBets

`func (o *GetBetsByTypeResponseV3) HasSpecialBets() bool`

HasSpecialBets returns a boolean if a field has been set.

### SetSpecialBets

`func (o *GetBetsByTypeResponseV3) SetSpecialBets(v []SpecialBet)`

SetSpecialBets gets a reference to the given []SpecialBet and assigns it to the SpecialBets field.

### GetManualBets

`func (o *GetBetsByTypeResponseV3) GetManualBets() []ManualBet`

GetManualBets returns the ManualBets field if non-nil, zero value otherwise.

### GetManualBetsOk

`func (o *GetBetsByTypeResponseV3) GetManualBetsOk() ([]ManualBet, bool)`

GetManualBetsOk returns a tuple with the ManualBets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasManualBets

`func (o *GetBetsByTypeResponseV3) HasManualBets() bool`

HasManualBets returns a boolean if a field has been set.

### SetManualBets

`func (o *GetBetsByTypeResponseV3) SetManualBets(v []ManualBet)`

SetManualBets gets a reference to the given []ManualBet and assigns it to the ManualBets field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


