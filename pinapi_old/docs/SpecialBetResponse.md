# SpecialBetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | Status of the request. | [optional] 
**ErrorCode** | Pointer to **NullableString** | When Status is PROCESSED_WITH_ERROR, provides a code indicating the specific problem.  ALL_BETTING_CLOSED &#x3D; Betting is not allowed at this moment. This may happen during system maintenance.    ABOVE_MAX_BET_AMOUNT &#x3D; Stake is above allowed maximum amount,    BELOW_MIN_BET_AMOUNT &#x3D; Stake is below allowed minimum amount,    BLOCKED_BETTING &#x3D; Betting is suspended for the client,    BLOCKED_CLIENT &#x3D; Client is no longer active,    CONTEST_NOT_FOUND &#x3D; Incorrect contest id provided or contest is no longer available,    DUPLICATE_UNIQUE_REQUEST_ID &#x3D; UniqueRequestId must be unique for each bet,    INCOMPLETE_CUSTOMER_BETTING_PROFILE &#x3D; Customer profile could not be loaded,     INSUFFICIENT_FUNDS &#x3D; Bet is submitted by a client with insufficient funds,    INVALID_COUNTRY &#x3D; Client country is not allowed for betting,    INVALID_REQUEST &#x3D; Special bet request is not valid,    LINE_CHANGED &#x3D; Bet is submitted on a line that has changed,    PAST_CUTOFFTIME &#x3D; Bet is submitted on a game after the betting cutoff time,    RESPONSIBLE_BETTING_LOSS_LIMIT_EXCEEDED &#x3D; Self-imposed loss limit exceeded,    RESPONSIBLE_BETTING_RISK_LIMIT_EXCEEDED &#x3D; Self-imposed risk limit exceeded,   SYSTEM_ERROR_1 &#x3D; Unexpected error,    SYSTEM_ERROR_2 &#x3D; Unexpected error,    UNIQUE_REQUEST_ID_REQUIRED &#x3D; UniqueRequestId is missing,    INVALID_CUSTOMER_PROFILE  | [optional] 
**BetId** | Pointer to **int64** | Id of a newly created bet. | [optional] 
**UniqueRequestId** | Pointer to **string** | Unique identifier provided in the request. | [optional] 
**BetterLineWasAccepted** | Pointer to **bool** | Whether or not the bet was accepted on the line that changed in favour of client. This can be true only if acceptBetterLine in the Place Bet request is set to TRUE. | [optional] 
**SpecialBet** | Pointer to [**SpecialBet**](SpecialBet.md) |  | [optional] 

## Methods

### GetStatus

`func (o *SpecialBetResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SpecialBetResponse) GetStatusOk() (string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStatus

`func (o *SpecialBetResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatus

`func (o *SpecialBetResponse) SetStatus(v string)`

SetStatus gets a reference to the given string and assigns it to the Status field.

### GetErrorCode

`func (o *SpecialBetResponse) GetErrorCode() NullableString`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *SpecialBetResponse) GetErrorCodeOk() (NullableString, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasErrorCode

`func (o *SpecialBetResponse) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### SetErrorCode

`func (o *SpecialBetResponse) SetErrorCode(v NullableString)`

SetErrorCode gets a reference to the given NullableString and assigns it to the ErrorCode field.

### SetErrorCodeExplicitNull

`func (o *SpecialBetResponse) SetErrorCodeExplicitNull(b bool)`

SetErrorCodeExplicitNull (un)sets ErrorCode to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ErrorCode value is set to nil even if false is passed
### GetBetId

`func (o *SpecialBetResponse) GetBetId() int64`

GetBetId returns the BetId field if non-nil, zero value otherwise.

### GetBetIdOk

`func (o *SpecialBetResponse) GetBetIdOk() (int64, bool)`

GetBetIdOk returns a tuple with the BetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBetId

`func (o *SpecialBetResponse) HasBetId() bool`

HasBetId returns a boolean if a field has been set.

### SetBetId

`func (o *SpecialBetResponse) SetBetId(v int64)`

SetBetId gets a reference to the given int64 and assigns it to the BetId field.

### GetUniqueRequestId

`func (o *SpecialBetResponse) GetUniqueRequestId() string`

GetUniqueRequestId returns the UniqueRequestId field if non-nil, zero value otherwise.

### GetUniqueRequestIdOk

`func (o *SpecialBetResponse) GetUniqueRequestIdOk() (string, bool)`

GetUniqueRequestIdOk returns a tuple with the UniqueRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUniqueRequestId

`func (o *SpecialBetResponse) HasUniqueRequestId() bool`

HasUniqueRequestId returns a boolean if a field has been set.

### SetUniqueRequestId

`func (o *SpecialBetResponse) SetUniqueRequestId(v string)`

SetUniqueRequestId gets a reference to the given string and assigns it to the UniqueRequestId field.

### GetBetterLineWasAccepted

`func (o *SpecialBetResponse) GetBetterLineWasAccepted() bool`

GetBetterLineWasAccepted returns the BetterLineWasAccepted field if non-nil, zero value otherwise.

### GetBetterLineWasAcceptedOk

`func (o *SpecialBetResponse) GetBetterLineWasAcceptedOk() (bool, bool)`

GetBetterLineWasAcceptedOk returns a tuple with the BetterLineWasAccepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBetterLineWasAccepted

`func (o *SpecialBetResponse) HasBetterLineWasAccepted() bool`

HasBetterLineWasAccepted returns a boolean if a field has been set.

### SetBetterLineWasAccepted

`func (o *SpecialBetResponse) SetBetterLineWasAccepted(v bool)`

SetBetterLineWasAccepted gets a reference to the given bool and assigns it to the BetterLineWasAccepted field.

### GetSpecialBet

`func (o *SpecialBetResponse) GetSpecialBet() SpecialBet`

GetSpecialBet returns the SpecialBet field if non-nil, zero value otherwise.

### GetSpecialBetOk

`func (o *SpecialBetResponse) GetSpecialBetOk() (SpecialBet, bool)`

GetSpecialBetOk returns a tuple with the SpecialBet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSpecialBet

`func (o *SpecialBetResponse) HasSpecialBet() bool`

HasSpecialBet returns a boolean if a field has been set.

### SetSpecialBet

`func (o *SpecialBetResponse) SetSpecialBet(v SpecialBet)`

SetSpecialBet gets a reference to the given SpecialBet and assigns it to the SpecialBet field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


