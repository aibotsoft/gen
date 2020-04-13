# TeaserGroupsPayout

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumberOfLegs** | Pointer to **int32** | Number of legs that must be bet and won to get the associated price. | [optional] 
**Price** | Pointer to **float64** | Price of the bet given the specified number of legs. | [optional] 

## Methods

### GetNumberOfLegs

`func (o *TeaserGroupsPayout) GetNumberOfLegs() int32`

GetNumberOfLegs returns the NumberOfLegs field if non-nil, zero value otherwise.

### GetNumberOfLegsOk

`func (o *TeaserGroupsPayout) GetNumberOfLegsOk() (int32, bool)`

GetNumberOfLegsOk returns a tuple with the NumberOfLegs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNumberOfLegs

`func (o *TeaserGroupsPayout) HasNumberOfLegs() bool`

HasNumberOfLegs returns a boolean if a field has been set.

### SetNumberOfLegs

`func (o *TeaserGroupsPayout) SetNumberOfLegs(v int32)`

SetNumberOfLegs gets a reference to the given int32 and assigns it to the NumberOfLegs field.

### GetPrice

`func (o *TeaserGroupsPayout) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *TeaserGroupsPayout) GetPriceOk() (float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPrice

`func (o *TeaserGroupsPayout) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPrice

`func (o *TeaserGroupsPayout) SetPrice(v float64)`

SetPrice gets a reference to the given float64 and assigns it to the Price field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


