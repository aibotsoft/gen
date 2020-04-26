# CancellationReasonType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | Cancellation Reason Code | [optional] 
**Details** | Pointer to [**CancellationReasonDetailsType**](CancellationReasonDetailsType.md) |  | [optional] 

## Methods

### GetCode

`func (o *CancellationReasonType) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CancellationReasonType) GetCodeOk() (string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCode

`func (o *CancellationReasonType) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCode

`func (o *CancellationReasonType) SetCode(v string)`

SetCode gets a reference to the given string and assigns it to the Code field.

### GetDetails

`func (o *CancellationReasonType) GetDetails() CancellationReasonDetailsType`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *CancellationReasonType) GetDetailsOk() (CancellationReasonDetailsType, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDetails

`func (o *CancellationReasonType) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### SetDetails

`func (o *CancellationReasonType) SetDetails(v CancellationReasonDetailsType)`

SetDetails gets a reference to the given CancellationReasonDetailsType and assigns it to the Details field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


