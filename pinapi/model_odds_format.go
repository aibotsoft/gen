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

// OddsFormat Bet odds format.   AMERICAN = American odds format,   DECIMAL = Decimal (European) odds format,   HONGKONG = Hong Kong odds format,   INDONESIAN = Indonesian odds format,   MALAY = Malaysian odds format  
type OddsFormat string

// List of OddsFormat
const (
	AMERICAN OddsFormat = "AMERICAN"
	DECIMAL OddsFormat = "DECIMAL"
	HONGKONG OddsFormat = "HONGKONG"
	INDONESIAN OddsFormat = "INDONESIAN"
	MALAY OddsFormat = "MALAY"
)

// Ptr returns reference to OddsFormat value
func (v OddsFormat) Ptr() *OddsFormat {
	return &v
}


type NullableOddsFormat struct {
	value *OddsFormat
	isSet bool
}

func (v NullableOddsFormat) Get() *OddsFormat {
	return v.value
}

func (v *NullableOddsFormat) Set(val *OddsFormat) {
	v.value = val
	v.isSet = true
}

func (v NullableOddsFormat) IsSet() bool {
	return v.isSet
}

func (v *NullableOddsFormat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOddsFormat(val *OddsFormat) *NullableOddsFormat {
	return &NullableOddsFormat{value: val, isSet: true}
}

func (v NullableOddsFormat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOddsFormat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
