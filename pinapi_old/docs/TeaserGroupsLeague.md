# TeaserGroupsLeague

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Unique identifier. League details can be retrieved from a call to v2/leagues endpoint. | [optional] 
**Spread** | Pointer to [**TeaserGroupsBetType**](TeaserGroupsBetType.md) |  | [optional] 
**Total** | Pointer to [**TeaserGroupsBetType**](TeaserGroupsBetType.md) |  | [optional] 

## Methods

### GetId

`func (o *TeaserGroupsLeague) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TeaserGroupsLeague) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *TeaserGroupsLeague) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *TeaserGroupsLeague) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetSpread

`func (o *TeaserGroupsLeague) GetSpread() TeaserGroupsBetType`

GetSpread returns the Spread field if non-nil, zero value otherwise.

### GetSpreadOk

`func (o *TeaserGroupsLeague) GetSpreadOk() (TeaserGroupsBetType, bool)`

GetSpreadOk returns a tuple with the Spread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSpread

`func (o *TeaserGroupsLeague) HasSpread() bool`

HasSpread returns a boolean if a field has been set.

### SetSpread

`func (o *TeaserGroupsLeague) SetSpread(v TeaserGroupsBetType)`

SetSpread gets a reference to the given TeaserGroupsBetType and assigns it to the Spread field.

### GetTotal

`func (o *TeaserGroupsLeague) GetTotal() TeaserGroupsBetType`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *TeaserGroupsLeague) GetTotalOk() (TeaserGroupsBetType, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTotal

`func (o *TeaserGroupsLeague) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### SetTotal

`func (o *TeaserGroupsLeague) SetTotal(v TeaserGroupsBetType)`

SetTotal gets a reference to the given TeaserGroupsBetType and assigns it to the Total field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


