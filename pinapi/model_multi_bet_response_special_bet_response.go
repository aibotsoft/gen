/*
 * Pinnacle - Customer API Reference
 *
 *  # Authentication   API uses HTTP Basic access authentication. You need to send Authorization HTTP Request header:    `Authorization: Basic <Base64 value of UTF-8 encoded \"username:password\">`  Example:  `Authorization: Basic U03MyOT23YbzMDc6d3c3O1DQ1` 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package pinapi

import (
	"encoding/json"
)

// MultiBetResponseSpecialBetResponse struct for MultiBetResponseSpecialBetResponse
type MultiBetResponseSpecialBetResponse struct {
	// The individual bets.
	Bets *[]SpecialBetResponse `json:"bets,omitempty"`
}

// NewMultiBetResponseSpecialBetResponse instantiates a new MultiBetResponseSpecialBetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultiBetResponseSpecialBetResponse() *MultiBetResponseSpecialBetResponse {
	this := MultiBetResponseSpecialBetResponse{}
	return &this
}

// NewMultiBetResponseSpecialBetResponseWithDefaults instantiates a new MultiBetResponseSpecialBetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultiBetResponseSpecialBetResponseWithDefaults() *MultiBetResponseSpecialBetResponse {
	this := MultiBetResponseSpecialBetResponse{}
	return &this
}

// GetBets returns the Bets field value if set, zero value otherwise.
func (o *MultiBetResponseSpecialBetResponse) GetBets() []SpecialBetResponse {
	if o == nil || o.Bets == nil {
		var ret []SpecialBetResponse
		return ret
	}
	return *o.Bets
}

// GetBetsOk returns a tuple with the Bets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiBetResponseSpecialBetResponse) GetBetsOk() (*[]SpecialBetResponse, bool) {
	if o == nil || o.Bets == nil {
		return nil, false
	}
	return o.Bets, true
}

// HasBets returns a boolean if a field has been set.
func (o *MultiBetResponseSpecialBetResponse) HasBets() bool {
	if o != nil && o.Bets != nil {
		return true
	}

	return false
}

// SetBets gets a reference to the given []SpecialBetResponse and assigns it to the Bets field.
func (o *MultiBetResponseSpecialBetResponse) SetBets(v []SpecialBetResponse) {
	o.Bets = &v
}

func (o MultiBetResponseSpecialBetResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Bets != nil {
		toSerialize["bets"] = o.Bets
	}
	return json.Marshal(toSerialize)
}

type NullableMultiBetResponseSpecialBetResponse struct {
	value *MultiBetResponseSpecialBetResponse
	isSet bool
}

func (v NullableMultiBetResponseSpecialBetResponse) Get() *MultiBetResponseSpecialBetResponse {
	return v.value
}

func (v *NullableMultiBetResponseSpecialBetResponse) Set(val *MultiBetResponseSpecialBetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMultiBetResponseSpecialBetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMultiBetResponseSpecialBetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultiBetResponseSpecialBetResponse(val *MultiBetResponseSpecialBetResponse) *NullableMultiBetResponseSpecialBetResponse {
	return &NullableMultiBetResponseSpecialBetResponse{value: val, isSet: true}
}

func (v NullableMultiBetResponseSpecialBetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultiBetResponseSpecialBetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}