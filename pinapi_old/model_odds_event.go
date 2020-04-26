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

// OddsEvent struct for OddsEvent
type OddsEvent struct {
	// Event Id.
	Id *int64 `json:"id,omitempty" xml:"id"`
	// Away team score. Only for live soccer events.Supported only for full match period (number=0).
	AwayScore *float64 `json:"awayScore,omitempty" xml:"awayScore"`
	// Home team score. Only for live soccer events.Supported only for full match period (number=0).
	HomeScore *float64 `json:"homeScore,omitempty" xml:"homeScore"`
	// Away team red cards. Only for live soccer events. Supported only for full match period (number=0).
	AwayRedCards *int32 `json:"awayRedCards,omitempty" xml:"awayRedCards"`
	// Home team red cards. Only for live soccer events.Supported only for full match period (number=0).
	HomeRedCards *int32 `json:"homeRedCards,omitempty" xml:"homeRedCards"`
	// Contains a list of periods.
	Periods *[]OddsPeriod `json:"periods,omitempty" xml:"periods"`
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OddsEvent) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *OddsEvent) GetIdOk() (int64, bool) {
	if o == nil || o.Id == nil {
		var ret int64
		return ret, false
	}
	return *o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OddsEvent) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *OddsEvent) SetId(v int64) {
	o.Id = &v
}

// GetAwayScore returns the AwayScore field value if set, zero value otherwise.
func (o *OddsEvent) GetAwayScore() float64 {
	if o == nil || o.AwayScore == nil {
		var ret float64
		return ret
	}
	return *o.AwayScore
}

// GetAwayScoreOk returns a tuple with the AwayScore field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *OddsEvent) GetAwayScoreOk() (float64, bool) {
	if o == nil || o.AwayScore == nil {
		var ret float64
		return ret, false
	}
	return *o.AwayScore, true
}

// HasAwayScore returns a boolean if a field has been set.
func (o *OddsEvent) HasAwayScore() bool {
	if o != nil && o.AwayScore != nil {
		return true
	}

	return false
}

// SetAwayScore gets a reference to the given float64 and assigns it to the AwayScore field.
func (o *OddsEvent) SetAwayScore(v float64) {
	o.AwayScore = &v
}

// GetHomeScore returns the HomeScore field value if set, zero value otherwise.
func (o *OddsEvent) GetHomeScore() float64 {
	if o == nil || o.HomeScore == nil {
		var ret float64
		return ret
	}
	return *o.HomeScore
}

// GetHomeScoreOk returns a tuple with the HomeScore field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *OddsEvent) GetHomeScoreOk() (float64, bool) {
	if o == nil || o.HomeScore == nil {
		var ret float64
		return ret, false
	}
	return *o.HomeScore, true
}

// HasHomeScore returns a boolean if a field has been set.
func (o *OddsEvent) HasHomeScore() bool {
	if o != nil && o.HomeScore != nil {
		return true
	}

	return false
}

// SetHomeScore gets a reference to the given float64 and assigns it to the HomeScore field.
func (o *OddsEvent) SetHomeScore(v float64) {
	o.HomeScore = &v
}

// GetAwayRedCards returns the AwayRedCards field value if set, zero value otherwise.
func (o *OddsEvent) GetAwayRedCards() int32 {
	if o == nil || o.AwayRedCards == nil {
		var ret int32
		return ret
	}
	return *o.AwayRedCards
}

// GetAwayRedCardsOk returns a tuple with the AwayRedCards field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *OddsEvent) GetAwayRedCardsOk() (int32, bool) {
	if o == nil || o.AwayRedCards == nil {
		var ret int32
		return ret, false
	}
	return *o.AwayRedCards, true
}

// HasAwayRedCards returns a boolean if a field has been set.
func (o *OddsEvent) HasAwayRedCards() bool {
	if o != nil && o.AwayRedCards != nil {
		return true
	}

	return false
}

// SetAwayRedCards gets a reference to the given int32 and assigns it to the AwayRedCards field.
func (o *OddsEvent) SetAwayRedCards(v int32) {
	o.AwayRedCards = &v
}

// GetHomeRedCards returns the HomeRedCards field value if set, zero value otherwise.
func (o *OddsEvent) GetHomeRedCards() int32 {
	if o == nil || o.HomeRedCards == nil {
		var ret int32
		return ret
	}
	return *o.HomeRedCards
}

// GetHomeRedCardsOk returns a tuple with the HomeRedCards field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *OddsEvent) GetHomeRedCardsOk() (int32, bool) {
	if o == nil || o.HomeRedCards == nil {
		var ret int32
		return ret, false
	}
	return *o.HomeRedCards, true
}

// HasHomeRedCards returns a boolean if a field has been set.
func (o *OddsEvent) HasHomeRedCards() bool {
	if o != nil && o.HomeRedCards != nil {
		return true
	}

	return false
}

// SetHomeRedCards gets a reference to the given int32 and assigns it to the HomeRedCards field.
func (o *OddsEvent) SetHomeRedCards(v int32) {
	o.HomeRedCards = &v
}

// GetPeriods returns the Periods field value if set, zero value otherwise.
func (o *OddsEvent) GetPeriods() []OddsPeriod {
	if o == nil || o.Periods == nil {
		var ret []OddsPeriod
		return ret
	}
	return *o.Periods
}

// GetPeriodsOk returns a tuple with the Periods field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *OddsEvent) GetPeriodsOk() ([]OddsPeriod, bool) {
	if o == nil || o.Periods == nil {
		var ret []OddsPeriod
		return ret, false
	}
	return *o.Periods, true
}

// HasPeriods returns a boolean if a field has been set.
func (o *OddsEvent) HasPeriods() bool {
	if o != nil && o.Periods != nil {
		return true
	}

	return false
}

// SetPeriods gets a reference to the given []OddsPeriod and assigns it to the Periods field.
func (o *OddsEvent) SetPeriods(v []OddsPeriod) {
	o.Periods = &v
}

type NullableOddsEvent struct {
	Value        OddsEvent
	ExplicitNull bool
}

func (v NullableOddsEvent) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableOddsEvent) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
