# TeaserGroups

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Unique identifier. | [optional] 
**Name** | Pointer to **string** | Friendly name for the Teaser Group | [optional] 
**Teasers** | Pointer to [**[]TeaserGroupsTeaser**](TeaserGroupsTeaser.md) | A collection of Teaser. | [optional] 

## Methods

### GetId

`func (o *TeaserGroups) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TeaserGroups) GetIdOk() (int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *TeaserGroups) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *TeaserGroups) SetId(v int64)`

SetId gets a reference to the given int64 and assigns it to the Id field.

### GetName

`func (o *TeaserGroups) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TeaserGroups) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *TeaserGroups) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *TeaserGroups) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetTeasers

`func (o *TeaserGroups) GetTeasers() []TeaserGroupsTeaser`

GetTeasers returns the Teasers field if non-nil, zero value otherwise.

### GetTeasersOk

`func (o *TeaserGroups) GetTeasersOk() ([]TeaserGroupsTeaser, bool)`

GetTeasersOk returns a tuple with the Teasers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTeasers

`func (o *TeaserGroups) HasTeasers() bool`

HasTeasers returns a boolean if a field has been set.

### SetTeasers

`func (o *TeaserGroups) SetTeasers(v []TeaserGroupsTeaser)`

SetTeasers gets a reference to the given []TeaserGroupsTeaser and assigns it to the Teasers field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


