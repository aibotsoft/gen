/*
 * Pinnacle - Customer API Reference
 *
 *  # Authentication   API uses HTTP Basic access authentication. You need to send Authorization HTTP Request header:    `Authorization: Basic <Base64 value of UTF-8 encoded \"username:password\">`  Example:  `Authorization: Basic U03MyOT23YbzMDc6d3c3O1DQ1`
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pinapi

import (
	"bytes"
	"encoding/json"
)

// Currency struct for Currency
type Currency struct {
	// Currency code.
	Code *string `json:"code,omitempty" xml:"code"`
	// Currency name.
	Name *string `json:"name,omitempty" xml:"name"`
	// Exchange rate to USD.
	Rate *float64 `json:"rate,omitempty" xml:"rate"`
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *Currency) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Currency) GetCodeOk() (string, bool) {
	if o == nil || o.Code == nil {
		var ret string
		return ret, false
	}
	return *o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *Currency) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *Currency) SetCode(v string) {
	o.Code = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Currency) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Currency) GetNameOk() (string, bool) {
	if o == nil || o.Name == nil {
		var ret string
		return ret, false
	}
	return *o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Currency) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Currency) SetName(v string) {
	o.Name = &v
}

// GetRate returns the Rate field value if set, zero value otherwise.
func (o *Currency) GetRate() float64 {
	if o == nil || o.Rate == nil {
		var ret float64
		return ret
	}
	return *o.Rate
}

// GetRateOk returns a tuple with the Rate field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Currency) GetRateOk() (float64, bool) {
	if o == nil || o.Rate == nil {
		var ret float64
		return ret, false
	}
	return *o.Rate, true
}

// HasRate returns a boolean if a field has been set.
func (o *Currency) HasRate() bool {
	if o != nil && o.Rate != nil {
		return true
	}

	return false
}

// SetRate gets a reference to the given float64 and assigns it to the Rate field.
func (o *Currency) SetRate(v float64) {
	o.Rate = &v
}

type NullableCurrency struct {
	Value        Currency
	ExplicitNull bool
}

func (v NullableCurrency) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableCurrency) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
