# LinesErrorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Error** | Pointer to [**ErrorResponse**](ErrorResponse.md) |  | [optional] 
**Code** | Pointer to **int32** | Code identifying an error that occurred. | [optional] 

## Methods

### GetStatus

`func (o *LinesErrorResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LinesErrorResponse) GetStatusOk() (string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStatus

`func (o *LinesErrorResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatus

`func (o *LinesErrorResponse) SetStatus(v string)`

SetStatus gets a reference to the given string and assigns it to the Status field.

### GetError

`func (o *LinesErrorResponse) GetError() ErrorResponse`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *LinesErrorResponse) GetErrorOk() (ErrorResponse, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasError

`func (o *LinesErrorResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### SetError

`func (o *LinesErrorResponse) SetError(v ErrorResponse)`

SetError gets a reference to the given ErrorResponse and assigns it to the Error field.

### GetCode

`func (o *LinesErrorResponse) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *LinesErrorResponse) GetCodeOk() (int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCode

`func (o *LinesErrorResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCode

`func (o *LinesErrorResponse) SetCode(v int32)`

SetCode gets a reference to the given int32 and assigns it to the Code field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


