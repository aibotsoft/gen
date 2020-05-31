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

// BetTypeClassEnum the model 'BetTypeClassEnum'
type BetTypeClassEnum string

// List of BetTypeClassEnum
const (
	OU BetTypeClassEnum = "OU"
	MORE BetTypeClassEnum = "more"
)

// Ptr returns reference to BetTypeClassEnum value
func (v BetTypeClassEnum) Ptr() *BetTypeClassEnum {
	return &v
}


type NullableBetTypeClassEnum struct {
	value *BetTypeClassEnum
	isSet bool
}

func (v NullableBetTypeClassEnum) Get() *BetTypeClassEnum {
	return v.value
}

func (v *NullableBetTypeClassEnum) Set(val *BetTypeClassEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableBetTypeClassEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableBetTypeClassEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetTypeClassEnum(val *BetTypeClassEnum) *NullableBetTypeClassEnum {
	return &NullableBetTypeClassEnum{value: val, isSet: true}
}

func (v NullableBetTypeClassEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetTypeClassEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}