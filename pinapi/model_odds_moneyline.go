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

// OddsMoneyline struct for OddsMoneyline
type OddsMoneyline struct {
	// Away team price
	Home *float64 `json:"home,omitempty"`
	// Away team price.
	Away *float64 `json:"away,omitempty"`
	// Draw price. This is present only for events we offer price for draw.
	Draw *float64 `json:"draw,omitempty"`
}

// NewOddsMoneyline instantiates a new OddsMoneyline object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOddsMoneyline() *OddsMoneyline {
	this := OddsMoneyline{}
	return &this
}

// NewOddsMoneylineWithDefaults instantiates a new OddsMoneyline object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOddsMoneylineWithDefaults() *OddsMoneyline {
	this := OddsMoneyline{}
	return &this
}

// GetHome returns the Home field value if set, zero value otherwise.
func (o *OddsMoneyline) GetHome() float64 {
	if o == nil || o.Home == nil {
		var ret float64
		return ret
	}
	return *o.Home
}

// GetHomeOk returns a tuple with the Home field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OddsMoneyline) GetHomeOk() (*float64, bool) {
	if o == nil || o.Home == nil {
		return nil, false
	}
	return o.Home, true
}

// HasHome returns a boolean if a field has been set.
func (o *OddsMoneyline) HasHome() bool {
	if o != nil && o.Home != nil {
		return true
	}

	return false
}

// SetHome gets a reference to the given float64 and assigns it to the Home field.
func (o *OddsMoneyline) SetHome(v float64) {
	o.Home = &v
}

// GetAway returns the Away field value if set, zero value otherwise.
func (o *OddsMoneyline) GetAway() float64 {
	if o == nil || o.Away == nil {
		var ret float64
		return ret
	}
	return *o.Away
}

// GetAwayOk returns a tuple with the Away field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OddsMoneyline) GetAwayOk() (*float64, bool) {
	if o == nil || o.Away == nil {
		return nil, false
	}
	return o.Away, true
}

// HasAway returns a boolean if a field has been set.
func (o *OddsMoneyline) HasAway() bool {
	if o != nil && o.Away != nil {
		return true
	}

	return false
}

// SetAway gets a reference to the given float64 and assigns it to the Away field.
func (o *OddsMoneyline) SetAway(v float64) {
	o.Away = &v
}

// GetDraw returns the Draw field value if set, zero value otherwise.
func (o *OddsMoneyline) GetDraw() float64 {
	if o == nil || o.Draw == nil {
		var ret float64
		return ret
	}
	return *o.Draw
}

// GetDrawOk returns a tuple with the Draw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OddsMoneyline) GetDrawOk() (*float64, bool) {
	if o == nil || o.Draw == nil {
		return nil, false
	}
	return o.Draw, true
}

// HasDraw returns a boolean if a field has been set.
func (o *OddsMoneyline) HasDraw() bool {
	if o != nil && o.Draw != nil {
		return true
	}

	return false
}

// SetDraw gets a reference to the given float64 and assigns it to the Draw field.
func (o *OddsMoneyline) SetDraw(v float64) {
	o.Draw = &v
}

func (o OddsMoneyline) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Home != nil {
		toSerialize["home"] = o.Home
	}
	if o.Away != nil {
		toSerialize["away"] = o.Away
	}
	if o.Draw != nil {
		toSerialize["draw"] = o.Draw
	}
	return json.Marshal(toSerialize)
}

type NullableOddsMoneyline struct {
	value *OddsMoneyline
	isSet bool
}

func (v NullableOddsMoneyline) Get() *OddsMoneyline {
	return v.value
}

func (v *NullableOddsMoneyline) Set(val *OddsMoneyline) {
	v.value = val
	v.isSet = true
}

func (v NullableOddsMoneyline) IsSet() bool {
	return v.isSet
}

func (v *NullableOddsMoneyline) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOddsMoneyline(val *OddsMoneyline) *NullableOddsMoneyline {
	return &NullableOddsMoneyline{value: val, isSet: true}
}

func (v NullableOddsMoneyline) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOddsMoneyline) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}