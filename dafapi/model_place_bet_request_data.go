/*
 * Sample API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package dafapi

import (
	"encoding/json"
)

// PlaceBetRequestData struct for PlaceBetRequestData
type PlaceBetRequestData struct {
	ItemList *[]PlaceBetItem `json:"ItemList,omitempty"`
	ErrorMsg *string `json:"ErrorMsg,omitempty"`
	Common *CommonItem `json:"Common,omitempty"`
}

// NewPlaceBetRequestData instantiates a new PlaceBetRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlaceBetRequestData() *PlaceBetRequestData {
	this := PlaceBetRequestData{}
	return &this
}

// NewPlaceBetRequestDataWithDefaults instantiates a new PlaceBetRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlaceBetRequestDataWithDefaults() *PlaceBetRequestData {
	this := PlaceBetRequestData{}
	return &this
}

// GetItemList returns the ItemList field value if set, zero value otherwise.
func (o *PlaceBetRequestData) GetItemList() []PlaceBetItem {
	if o == nil || o.ItemList == nil {
		var ret []PlaceBetItem
		return ret
	}
	return *o.ItemList
}

// GetItemListOk returns a tuple with the ItemList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaceBetRequestData) GetItemListOk() (*[]PlaceBetItem, bool) {
	if o == nil || o.ItemList == nil {
		return nil, false
	}
	return o.ItemList, true
}

// HasItemList returns a boolean if a field has been set.
func (o *PlaceBetRequestData) HasItemList() bool {
	if o != nil && o.ItemList != nil {
		return true
	}

	return false
}

// SetItemList gets a reference to the given []PlaceBetItem and assigns it to the ItemList field.
func (o *PlaceBetRequestData) SetItemList(v []PlaceBetItem) {
	o.ItemList = &v
}

// GetErrorMsg returns the ErrorMsg field value if set, zero value otherwise.
func (o *PlaceBetRequestData) GetErrorMsg() string {
	if o == nil || o.ErrorMsg == nil {
		var ret string
		return ret
	}
	return *o.ErrorMsg
}

// GetErrorMsgOk returns a tuple with the ErrorMsg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaceBetRequestData) GetErrorMsgOk() (*string, bool) {
	if o == nil || o.ErrorMsg == nil {
		return nil, false
	}
	return o.ErrorMsg, true
}

// HasErrorMsg returns a boolean if a field has been set.
func (o *PlaceBetRequestData) HasErrorMsg() bool {
	if o != nil && o.ErrorMsg != nil {
		return true
	}

	return false
}

// SetErrorMsg gets a reference to the given string and assigns it to the ErrorMsg field.
func (o *PlaceBetRequestData) SetErrorMsg(v string) {
	o.ErrorMsg = &v
}

// GetCommon returns the Common field value if set, zero value otherwise.
func (o *PlaceBetRequestData) GetCommon() CommonItem {
	if o == nil || o.Common == nil {
		var ret CommonItem
		return ret
	}
	return *o.Common
}

// GetCommonOk returns a tuple with the Common field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaceBetRequestData) GetCommonOk() (*CommonItem, bool) {
	if o == nil || o.Common == nil {
		return nil, false
	}
	return o.Common, true
}

// HasCommon returns a boolean if a field has been set.
func (o *PlaceBetRequestData) HasCommon() bool {
	if o != nil && o.Common != nil {
		return true
	}

	return false
}

// SetCommon gets a reference to the given CommonItem and assigns it to the Common field.
func (o *PlaceBetRequestData) SetCommon(v CommonItem) {
	o.Common = &v
}

func (o PlaceBetRequestData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ItemList != nil {
		toSerialize["ItemList"] = o.ItemList
	}
	if o.ErrorMsg != nil {
		toSerialize["ErrorMsg"] = o.ErrorMsg
	}
	if o.Common != nil {
		toSerialize["Common"] = o.Common
	}
	return json.Marshal(toSerialize)
}

type NullablePlaceBetRequestData struct {
	value *PlaceBetRequestData
	isSet bool
}

func (v NullablePlaceBetRequestData) Get() *PlaceBetRequestData {
	return v.value
}

func (v *NullablePlaceBetRequestData) Set(val *PlaceBetRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullablePlaceBetRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullablePlaceBetRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlaceBetRequestData(val *PlaceBetRequestData) *NullablePlaceBetRequestData {
	return &NullablePlaceBetRequestData{value: val, isSet: true}
}

func (v NullablePlaceBetRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlaceBetRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
