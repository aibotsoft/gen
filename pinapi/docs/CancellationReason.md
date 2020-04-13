# CancellationReason

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** |  | 
**Details** | Pointer to [**[]CancellationDetailsItem**](CancellationDetailsItem.md) |  | [optional] 

## Methods

### GetCode

`func (o *CancellationReason) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CancellationReason) GetCodeOk() (string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCode

`func (o *CancellationReason) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCode

`func (o *CancellationReason) SetCode(v string)`

SetCode gets a reference to the given string and assigns it to the Code field.

### GetDetails

`func (o *CancellationReason) GetDetails() []CancellationDetailsItem`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *CancellationReason) GetDetailsOk() ([]CancellationDetailsItem, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDetails

`func (o *CancellationReason) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### SetDetails

`func (o *CancellationReason) SetDetails(v []CancellationDetailsItem)`

SetDetails gets a reference to the given []CancellationDetailsItem and assigns it to the Details field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


