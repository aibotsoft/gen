/*
 * Pinnacle - Customer API Reference
 *
 *  # Authentication   API uses HTTP Basic access authentication. You need to send Authorization HTTP Request header:    `Authorization: Basic <Base64 value of UTF-8 encoded \"username:password\">`  Example:  `Authorization: Basic U03MyOT23YbzMDc6d3c3O1DQ1` 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"encoding/json"
)

// OddsTotal struct for OddsTotal
type OddsTotal struct {
	// This is present only if it’s alternative line.
	AltLineId *int64 `json:"altLineId,omitempty"`
	// Total points.
	Points *float64 `json:"points,omitempty"`
	// Over price.
	Over *float64 `json:"over,omitempty"`
	// Under price.
	Under *float64 `json:"under,omitempty"`
}

// NewOddsTotal instantiates a new OddsTotal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOddsTotal() *OddsTotal {
	this := OddsTotal{}
	return &this
}

// NewOddsTotalWithDefaults instantiates a new OddsTotal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOddsTotalWithDefaults() *OddsTotal {
	this := OddsTotal{}
	return &this
}

// GetAltLineId returns the AltLineId field value if set, zero value otherwise.
func (o *OddsTotal) GetAltLineId() int64 {
	if o == nil || o.AltLineId == nil {
		var ret int64
		return ret
	}
	return *o.AltLineId
}

// GetAltLineIdOk returns a tuple with the AltLineId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OddsTotal) GetAltLineIdOk() (*int64, bool) {
	if o == nil || o.AltLineId == nil {
		return nil, false
	}
	return o.AltLineId, true
}

// HasAltLineId returns a boolean if a field has been set.
func (o *OddsTotal) HasAltLineId() bool {
	if o != nil && o.AltLineId != nil {
		return true
	}

	return false
}

// SetAltLineId gets a reference to the given int64 and assigns it to the AltLineId field.
func (o *OddsTotal) SetAltLineId(v int64) {
	o.AltLineId = &v
}

// GetPoints returns the Points field value if set, zero value otherwise.
func (o *OddsTotal) GetPoints() float64 {
	if o == nil || o.Points == nil {
		var ret float64
		return ret
	}
	return *o.Points
}

// GetPointsOk returns a tuple with the Points field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OddsTotal) GetPointsOk() (*float64, bool) {
	if o == nil || o.Points == nil {
		return nil, false
	}
	return o.Points, true
}

// HasPoints returns a boolean if a field has been set.
func (o *OddsTotal) HasPoints() bool {
	if o != nil && o.Points != nil {
		return true
	}

	return false
}

// SetPoints gets a reference to the given float64 and assigns it to the Points field.
func (o *OddsTotal) SetPoints(v float64) {
	o.Points = &v
}

// GetOver returns the Over field value if set, zero value otherwise.
func (o *OddsTotal) GetOver() float64 {
	if o == nil || o.Over == nil {
		var ret float64
		return ret
	}
	return *o.Over
}

// GetOverOk returns a tuple with the Over field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OddsTotal) GetOverOk() (*float64, bool) {
	if o == nil || o.Over == nil {
		return nil, false
	}
	return o.Over, true
}

// HasOver returns a boolean if a field has been set.
func (o *OddsTotal) HasOver() bool {
	if o != nil && o.Over != nil {
		return true
	}

	return false
}

// SetOver gets a reference to the given float64 and assigns it to the Over field.
func (o *OddsTotal) SetOver(v float64) {
	o.Over = &v
}

// GetUnder returns the Under field value if set, zero value otherwise.
func (o *OddsTotal) GetUnder() float64 {
	if o == nil || o.Under == nil {
		var ret float64
		return ret
	}
	return *o.Under
}

// GetUnderOk returns a tuple with the Under field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OddsTotal) GetUnderOk() (*float64, bool) {
	if o == nil || o.Under == nil {
		return nil, false
	}
	return o.Under, true
}

// HasUnder returns a boolean if a field has been set.
func (o *OddsTotal) HasUnder() bool {
	if o != nil && o.Under != nil {
		return true
	}

	return false
}

// SetUnder gets a reference to the given float64 and assigns it to the Under field.
func (o *OddsTotal) SetUnder(v float64) {
	o.Under = &v
}

func (o OddsTotal) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AltLineId != nil {
		toSerialize["altLineId"] = o.AltLineId
	}
	if o.Points != nil {
		toSerialize["points"] = o.Points
	}
	if o.Over != nil {
		toSerialize["over"] = o.Over
	}
	if o.Under != nil {
		toSerialize["under"] = o.Under
	}
	return json.Marshal(toSerialize)
}

type NullableOddsTotal struct {
	value *OddsTotal
	isSet bool
}

func (v NullableOddsTotal) Get() *OddsTotal {
	return v.value
}

func (v *NullableOddsTotal) Set(val *OddsTotal) {
	v.value = val
	v.isSet = true
}

func (v NullableOddsTotal) IsSet() bool {
	return v.isSet
}

func (v *NullableOddsTotal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOddsTotal(val *OddsTotal) *NullableOddsTotal {
	return &NullableOddsTotal{value: val, isSet: true}
}

func (v NullableOddsTotal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOddsTotal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
