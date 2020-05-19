# FavoritesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCode** | Pointer to **int** |  | [optional] 
**ErrorMsg** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**FavoritesResponseData**](FavoritesResponse_Data.md) |  | [optional] 

## Methods

### NewFavoritesResponse

`func NewFavoritesResponse() *FavoritesResponse`

NewFavoritesResponse instantiates a new FavoritesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFavoritesResponseWithDefaults

`func NewFavoritesResponseWithDefaults() *FavoritesResponse`

NewFavoritesResponseWithDefaults instantiates a new FavoritesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *FavoritesResponse) GetErrorCode() int`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *FavoritesResponse) GetErrorCodeOk() (*int, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *FavoritesResponse) SetErrorCode(v int)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *FavoritesResponse) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorMsg

`func (o *FavoritesResponse) GetErrorMsg() string`

GetErrorMsg returns the ErrorMsg field if non-nil, zero value otherwise.

### GetErrorMsgOk

`func (o *FavoritesResponse) GetErrorMsgOk() (*string, bool)`

GetErrorMsgOk returns a tuple with the ErrorMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMsg

`func (o *FavoritesResponse) SetErrorMsg(v string)`

SetErrorMsg sets ErrorMsg field to given value.

### HasErrorMsg

`func (o *FavoritesResponse) HasErrorMsg() bool`

HasErrorMsg returns a boolean if a field has been set.

### GetData

`func (o *FavoritesResponse) GetData() FavoritesResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *FavoritesResponse) GetDataOk() (*FavoritesResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *FavoritesResponse) SetData(v FavoritesResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *FavoritesResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


