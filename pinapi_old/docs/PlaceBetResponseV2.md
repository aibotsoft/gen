# PlaceBetResponseV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | Status of the response. | [optional] 
**ErrorCode** | Pointer to **NullableString** | If Status is PROCESSED_WITH_ERROR, errorCode will be in the response.   ALL_BETTING_CLOSED &#x3D; Betting is not allowed at this moment. This may happen during system maintenance,   ALL_LIVE_BETTING_CLOSED &#x3D; Live betting is not allowed at this moment. This may happen during system maintenance,   ABOVE_EVENT_MAX &#x3D; Bet cannot be placed because client exceeded allowed maximum of risk on a line,   ABOVE_MAX_BET_AMOUNT &#x3D; Stake is above allowed maximum amount,    BELOW_MIN_BET_AMOUNT &#x3D; Stake is below allowed minimum amount,   BLOCKED_BETTING &#x3D; Betting is suspended for the client,   BLOCKED_CLIENT &#x3D; Client is no longer active,    INSUFFICIENT_FUNDS &#x3D; Bet is submitted by a client with insufficient funds,   INVALID_COUNTRY &#x3D; Client country is not allowed for betting,   INVALID_EVENT &#x3D; Invalid eventid,   INVALID_ODDS_FORMAT &#x3D; If a bet was submitted with the odds format that is not allowed for the client,   LINE_CHANGED &#x3D; Bet is submitted on a line that has changed,   LISTED_PITCHERS_SELECTION_ERROR &#x3D; If bet was submitted with pitcher1MustStart and/or pitcher2MustStart parameters in Place Bet request with values that are not allowed,   OFFLINE_EVENT &#x3D; Bet is submitted on an event that is offline or the submitted line is not offered at the moment due to points/handicap change or the submitted bet type is just not offered at the moment,   PAST_CUTOFFTIME &#x3D; Bet is submitted on a game after the betting cutoff time,   RED_CARDS_CHANGED &#x3D; Bet is submitted on a live soccer event with changed red card count,   SCORE_CHANGED &#x3D; Bet is submitted on a live soccer event with changed score,   TIME_RESTRICTION &#x3D; Bet is submitted within too short of a period from the same bet previously placed by a client,   DUPLICATE_UNIQUE_REQUEST_ID &#x3D; Request with the same uniqueRequestId was already processed. Please set the new value if you still want the request to be processed,   INCOMPLETE_CUSTOMER_BETTING_PROFILE &#x3D; System configuration issue,   INVALID_CUSTOMER_PROFILE &#x3D; System configuration issue,   LIMITS_CONFIGURATION_ISSUE &#x3D; System configuration issue,   RESPONSIBLE_BETTING_LOSS_LIMIT_EXCEEDED &#x3D; Client has reached his total loss limit,   RESPONSIBLE_BETTING_RISK_LIMIT_EXCEEDED &#x3D; Client has reached his total risk limit,   SYSTEM_ERROR_3 &#x3D; Unexpected error,   LICENCE_RESTRICTION_LIVE_BETTING_BLOCKED - Live betting blocked due to licence restrictions  | [optional] 
**UniqueRequestId** | Pointer to **string** | Echo of the uniqueRequestId from the request. | [optional] 
**StraightBet** | Pointer to [**StraightBet**](StraightBet.md) |  | [optional] 

## Methods

### GetStatus

`func (o *PlaceBetResponseV2) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PlaceBetResponseV2) GetStatusOk() (string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStatus

`func (o *PlaceBetResponseV2) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatus

`func (o *PlaceBetResponseV2) SetStatus(v string)`

SetStatus gets a reference to the given string and assigns it to the Status field.

### GetErrorCode

`func (o *PlaceBetResponseV2) GetErrorCode() NullableString`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *PlaceBetResponseV2) GetErrorCodeOk() (NullableString, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasErrorCode

`func (o *PlaceBetResponseV2) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### SetErrorCode

`func (o *PlaceBetResponseV2) SetErrorCode(v NullableString)`

SetErrorCode gets a reference to the given NullableString and assigns it to the ErrorCode field.

### SetErrorCodeExplicitNull

`func (o *PlaceBetResponseV2) SetErrorCodeExplicitNull(b bool)`

SetErrorCodeExplicitNull (un)sets ErrorCode to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ErrorCode value is set to nil even if false is passed
### GetUniqueRequestId

`func (o *PlaceBetResponseV2) GetUniqueRequestId() string`

GetUniqueRequestId returns the UniqueRequestId field if non-nil, zero value otherwise.

### GetUniqueRequestIdOk

`func (o *PlaceBetResponseV2) GetUniqueRequestIdOk() (string, bool)`

GetUniqueRequestIdOk returns a tuple with the UniqueRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUniqueRequestId

`func (o *PlaceBetResponseV2) HasUniqueRequestId() bool`

HasUniqueRequestId returns a boolean if a field has been set.

### SetUniqueRequestId

`func (o *PlaceBetResponseV2) SetUniqueRequestId(v string)`

SetUniqueRequestId gets a reference to the given string and assigns it to the UniqueRequestId field.

### GetStraightBet

`func (o *PlaceBetResponseV2) GetStraightBet() StraightBet`

GetStraightBet returns the StraightBet field if non-nil, zero value otherwise.

### GetStraightBetOk

`func (o *PlaceBetResponseV2) GetStraightBetOk() (StraightBet, bool)`

GetStraightBetOk returns a tuple with the StraightBet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStraightBet

`func (o *PlaceBetResponseV2) HasStraightBet() bool`

HasStraightBet returns a boolean if a field has been set.

### SetStraightBet

`func (o *PlaceBetResponseV2) SetStraightBet(v StraightBet)`

SetStraightBet gets a reference to the given StraightBet and assigns it to the StraightBet field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


