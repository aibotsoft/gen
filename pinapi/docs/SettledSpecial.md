# SettledSpecial

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Id for the Settled Special | [optional] 
**Status** | Pointer to **int32** | Status of the settled special. | [optional] 
**SettlementId** | Pointer to **int64** | Id for the Settled Special | [optional] 
**SettledAt** | Pointer to [**time.Time**](time.Time.md) | Settled DateTime | [optional] 
**CancellationReason** | Pointer to [**CancellationReason**](CancellationReason.md) |  | [optional] 

## Methods

### GetId

`func (o *SettledSpecial) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SettledSpecial) GetIdOk() (int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *SettledSpecial) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *SettledSpecial) SetId(v int64)`

SetId gets a reference to the given int64 and assigns it to the Id field.

### GetStatus

`func (o *SettledSpecial) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SettledSpecial) GetStatusOk() (int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStatus

`func (o *SettledSpecial) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatus

`func (o *SettledSpecial) SetStatus(v int32)`

SetStatus gets a reference to the given int32 and assigns it to the Status field.

### GetSettlementId

`func (o *SettledSpecial) GetSettlementId() int64`

GetSettlementId returns the SettlementId field if non-nil, zero value otherwise.

### GetSettlementIdOk

`func (o *SettledSpecial) GetSettlementIdOk() (int64, bool)`

GetSettlementIdOk returns a tuple with the SettlementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSettlementId

`func (o *SettledSpecial) HasSettlementId() bool`

HasSettlementId returns a boolean if a field has been set.

### SetSettlementId

`func (o *SettledSpecial) SetSettlementId(v int64)`

SetSettlementId gets a reference to the given int64 and assigns it to the SettlementId field.

### GetSettledAt

`func (o *SettledSpecial) GetSettledAt() time.Time`

GetSettledAt returns the SettledAt field if non-nil, zero value otherwise.

### GetSettledAtOk

`func (o *SettledSpecial) GetSettledAtOk() (time.Time, bool)`

GetSettledAtOk returns a tuple with the SettledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSettledAt

`func (o *SettledSpecial) HasSettledAt() bool`

HasSettledAt returns a boolean if a field has been set.

### SetSettledAt

`func (o *SettledSpecial) SetSettledAt(v time.Time)`

SetSettledAt gets a reference to the given time.Time and assigns it to the SettledAt field.

### GetCancellationReason

`func (o *SettledSpecial) GetCancellationReason() CancellationReason`

GetCancellationReason returns the CancellationReason field if non-nil, zero value otherwise.

### GetCancellationReasonOk

`func (o *SettledSpecial) GetCancellationReasonOk() (CancellationReason, bool)`

GetCancellationReasonOk returns a tuple with the CancellationReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCancellationReason

`func (o *SettledSpecial) HasCancellationReason() bool`

HasCancellationReason returns a boolean if a field has been set.

### SetCancellationReason

`func (o *SettledSpecial) SetCancellationReason(v CancellationReason)`

SetCancellationReason gets a reference to the given CancellationReason and assigns it to the CancellationReason field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


