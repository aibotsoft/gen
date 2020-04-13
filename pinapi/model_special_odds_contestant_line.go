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

// SpecialOddsContestantLine struct for SpecialOddsContestantLine
type SpecialOddsContestantLine struct {
	// ContestantLine Id.
	Id *int64 `json:"id,omitempty" xml:"id"`
	// Line identifier required for placing a bet.
	LineId *int64 `json:"lineId,omitempty" xml:"lineId"`
	// Price of the line.
	Price *float64 `json:"price,omitempty" xml:"price"`
	// A number indicating the spread, over/under etc.
	Handicap *float64 `json:"handicap,omitempty" xml:"handicap"`
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SpecialOddsContestantLine) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SpecialOddsContestantLine) GetIdOk() (int64, bool) {
	if o == nil || o.Id == nil {
		var ret int64
		return ret, false
	}
	return *o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SpecialOddsContestantLine) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *SpecialOddsContestantLine) SetId(v int64) {
	o.Id = &v
}

// GetLineId returns the LineId field value if set, zero value otherwise.
func (o *SpecialOddsContestantLine) GetLineId() int64 {
	if o == nil || o.LineId == nil {
		var ret int64
		return ret
	}
	return *o.LineId
}

// GetLineIdOk returns a tuple with the LineId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SpecialOddsContestantLine) GetLineIdOk() (int64, bool) {
	if o == nil || o.LineId == nil {
		var ret int64
		return ret, false
	}
	return *o.LineId, true
}

// HasLineId returns a boolean if a field has been set.
func (o *SpecialOddsContestantLine) HasLineId() bool {
	if o != nil && o.LineId != nil {
		return true
	}

	return false
}

// SetLineId gets a reference to the given int64 and assigns it to the LineId field.
func (o *SpecialOddsContestantLine) SetLineId(v int64) {
	o.LineId = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *SpecialOddsContestantLine) GetPrice() float64 {
	if o == nil || o.Price == nil {
		var ret float64
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SpecialOddsContestantLine) GetPriceOk() (float64, bool) {
	if o == nil || o.Price == nil {
		var ret float64
		return ret, false
	}
	return *o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *SpecialOddsContestantLine) HasPrice() bool {
	if o != nil && o.Price != nil {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float64 and assigns it to the Price field.
func (o *SpecialOddsContestantLine) SetPrice(v float64) {
	o.Price = &v
}

// GetHandicap returns the Handicap field value if set, zero value otherwise.
func (o *SpecialOddsContestantLine) GetHandicap() float64 {
	if o == nil || o.Handicap == nil {
		var ret float64
		return ret
	}
	return *o.Handicap
}

// GetHandicapOk returns a tuple with the Handicap field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SpecialOddsContestantLine) GetHandicapOk() (float64, bool) {
	if o == nil || o.Handicap == nil {
		var ret float64
		return ret, false
	}
	return *o.Handicap, true
}

// HasHandicap returns a boolean if a field has been set.
func (o *SpecialOddsContestantLine) HasHandicap() bool {
	if o != nil && o.Handicap != nil {
		return true
	}

	return false
}

// SetHandicap gets a reference to the given float64 and assigns it to the Handicap field.
func (o *SpecialOddsContestantLine) SetHandicap(v float64) {
	o.Handicap = &v
}

type NullableSpecialOddsContestantLine struct {
	Value        SpecialOddsContestantLine
	ExplicitNull bool
}

func (v NullableSpecialOddsContestantLine) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableSpecialOddsContestantLine) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
