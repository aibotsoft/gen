/*
 * Pinnacle - Customer API Reference
 *
 *  # Authentication   API uses HTTP Basic access authentication. You need to send Authorization HTTP Request header:    `Authorization: Basic <Base64 value of UTF-8 encoded \"username:password\">`  Example:  `Authorization: Basic U03MyOT23YbzMDc6d3c3O1DQ1`
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pinapi_old

import (
	"bytes"
	"encoding/json"
)

// OddsSpread struct for OddsSpread
type OddsSpread struct {
	// This is present only if it’s alternative line.
	AltLineId *int64 `json:"altLineId,omitempty" xml:"altLineId"`
	// Home team handicap.
	Hdp *float64 `json:"hdp,omitempty" xml:"hdp"`
	// Home team price.
	Home *float64 `json:"home,omitempty" xml:"home"`
	// Away team price.
	Away *float64 `json:"away,omitempty" xml:"away"`
}

// GetAltLineId returns the AltLineId field value if set, zero value otherwise.
func (o *OddsSpread) GetAltLineId() int64 {
	if o == nil || o.AltLineId == nil {
		var ret int64
		return ret
	}
	return *o.AltLineId
}

// GetAltLineIdOk returns a tuple with the AltLineId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *OddsSpread) GetAltLineIdOk() (int64, bool) {
	if o == nil || o.AltLineId == nil {
		var ret int64
		return ret, false
	}
	return *o.AltLineId, true
}

// HasAltLineId returns a boolean if a field has been set.
func (o *OddsSpread) HasAltLineId() bool {
	if o != nil && o.AltLineId != nil {
		return true
	}

	return false
}

// SetAltLineId gets a reference to the given int64 and assigns it to the AltLineId field.
func (o *OddsSpread) SetAltLineId(v int64) {
	o.AltLineId = &v
}

// GetHdp returns the Hdp field value if set, zero value otherwise.
func (o *OddsSpread) GetHdp() float64 {
	if o == nil || o.Hdp == nil {
		var ret float64
		return ret
	}
	return *o.Hdp
}

// GetHdpOk returns a tuple with the Hdp field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *OddsSpread) GetHdpOk() (float64, bool) {
	if o == nil || o.Hdp == nil {
		var ret float64
		return ret, false
	}
	return *o.Hdp, true
}

// HasHdp returns a boolean if a field has been set.
func (o *OddsSpread) HasHdp() bool {
	if o != nil && o.Hdp != nil {
		return true
	}

	return false
}

// SetHdp gets a reference to the given float64 and assigns it to the Hdp field.
func (o *OddsSpread) SetHdp(v float64) {
	o.Hdp = &v
}

// GetHome returns the Home field value if set, zero value otherwise.
func (o *OddsSpread) GetHome() float64 {
	if o == nil || o.Home == nil {
		var ret float64
		return ret
	}
	return *o.Home
}

// GetHomeOk returns a tuple with the Home field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *OddsSpread) GetHomeOk() (float64, bool) {
	if o == nil || o.Home == nil {
		var ret float64
		return ret, false
	}
	return *o.Home, true
}

// HasHome returns a boolean if a field has been set.
func (o *OddsSpread) HasHome() bool {
	if o != nil && o.Home != nil {
		return true
	}

	return false
}

// SetHome gets a reference to the given float64 and assigns it to the Home field.
func (o *OddsSpread) SetHome(v float64) {
	o.Home = &v
}

// GetAway returns the Away field value if set, zero value otherwise.
func (o *OddsSpread) GetAway() float64 {
	if o == nil || o.Away == nil {
		var ret float64
		return ret
	}
	return *o.Away
}

// GetAwayOk returns a tuple with the Away field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *OddsSpread) GetAwayOk() (float64, bool) {
	if o == nil || o.Away == nil {
		var ret float64
		return ret, false
	}
	return *o.Away, true
}

// HasAway returns a boolean if a field has been set.
func (o *OddsSpread) HasAway() bool {
	if o != nil && o.Away != nil {
		return true
	}

	return false
}

// SetAway gets a reference to the given float64 and assigns it to the Away field.
func (o *OddsSpread) SetAway(v float64) {
	o.Away = &v
}

type NullableOddsSpread struct {
	Value        OddsSpread
	ExplicitNull bool
}

func (v NullableOddsSpread) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableOddsSpread) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
