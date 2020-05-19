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

// GetMarketsResponseDataMarkets struct for GetMarketsResponseDataMarkets
type GetMarketsResponseDataMarkets struct {
	NewOdds *[]MarketItem `json:"NewOdds,omitempty"`
}

// NewGetMarketsResponseDataMarkets instantiates a new GetMarketsResponseDataMarkets object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMarketsResponseDataMarkets() *GetMarketsResponseDataMarkets {
	this := GetMarketsResponseDataMarkets{}
	return &this
}

// NewGetMarketsResponseDataMarketsWithDefaults instantiates a new GetMarketsResponseDataMarkets object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMarketsResponseDataMarketsWithDefaults() *GetMarketsResponseDataMarkets {
	this := GetMarketsResponseDataMarkets{}
	return &this
}

// GetNewOdds returns the NewOdds field value if set, zero value otherwise.
func (o *GetMarketsResponseDataMarkets) GetNewOdds() []MarketItem {
	if o == nil || o.NewOdds == nil {
		var ret []MarketItem
		return ret
	}
	return *o.NewOdds
}

// GetNewOddsOk returns a tuple with the NewOdds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketsResponseDataMarkets) GetNewOddsOk() (*[]MarketItem, bool) {
	if o == nil || o.NewOdds == nil {
		return nil, false
	}
	return o.NewOdds, true
}

// HasNewOdds returns a boolean if a field has been set.
func (o *GetMarketsResponseDataMarkets) HasNewOdds() bool {
	if o != nil && o.NewOdds != nil {
		return true
	}

	return false
}

// SetNewOdds gets a reference to the given []MarketItem and assigns it to the NewOdds field.
func (o *GetMarketsResponseDataMarkets) SetNewOdds(v []MarketItem) {
	o.NewOdds = &v
}

func (o GetMarketsResponseDataMarkets) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NewOdds != nil {
		toSerialize["NewOdds"] = o.NewOdds
	}
	return json.Marshal(toSerialize)
}

type NullableGetMarketsResponseDataMarkets struct {
	value *GetMarketsResponseDataMarkets
	isSet bool
}

func (v NullableGetMarketsResponseDataMarkets) Get() *GetMarketsResponseDataMarkets {
	return v.value
}

func (v *NullableGetMarketsResponseDataMarkets) Set(val *GetMarketsResponseDataMarkets) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMarketsResponseDataMarkets) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMarketsResponseDataMarkets) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMarketsResponseDataMarkets(val *GetMarketsResponseDataMarkets) *NullableGetMarketsResponseDataMarkets {
	return &NullableGetMarketsResponseDataMarkets{value: val, isSet: true}
}

func (v NullableGetMarketsResponseDataMarkets) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMarketsResponseDataMarkets) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
