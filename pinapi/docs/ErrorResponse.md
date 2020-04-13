# ErrorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | INVALID_REQUEST_DATA  &#x3D; Invalid request parameters (http status 400)   INVALID_CREDENTIALS &#x3D; Authorization failed, invalid credentials (http status 401)   INVALID_AUTHORIZATION_HEADER &#x3D; HTTP Authorization header is missing (http status 401)  ACCOUNT_INACTIVE &#x3D; Client&#39;s account is not active (http status 403)   NO_API_ACCESS &#x3D; Account not permitted to access the API (http status 403)  RESUBMIT_REQUEST &#x3D; It can happen only when placing a bet (http status 400).  Unable to process the request but the request itself is valid. This happens more often on the live betting in situations when there is more than one place bet request at the same on the same line. When this happens, we don&#39;t keep the place bet request on the server until we know if we can accept or reject the bet, but instead we return the error. It&#39;s also very likely that the line will change after that. To reduce a chance of getting RESUBMIT_REQUEST client should try to place a bet as fast as possible.   | [optional] 
**Message** | Pointer to **string** |  | [optional] 

## Methods

### GetCode

`func (o *ErrorResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ErrorResponse) GetCodeOk() (string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCode

`func (o *ErrorResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCode

`func (o *ErrorResponse) SetCode(v string)`

SetCode gets a reference to the given string and assigns it to the Code field.

### GetMessage

`func (o *ErrorResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ErrorResponse) GetMessageOk() (string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMessage

`func (o *ErrorResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessage

`func (o *ErrorResponse) SetMessage(v string)`

SetMessage gets a reference to the given string and assigns it to the Message field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


