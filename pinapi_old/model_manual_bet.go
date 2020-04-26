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
	"time"
)

// ManualBet struct for ManualBet
type ManualBet struct {
	// Bet identification
	BetId int64 `json:"betId" xml:"betId"`
	// Wager identification. All bets placed thru the API will have value 1. Website Classic view supports multiple contest(special) bets placement in the same bet slip in that case the bet would have appropriate wager number, as well as all round robin parlay bets.
	WagerNumber int32 `json:"wagerNumber" xml:"wagerNumber"`
	// Date time when the bet was placed.
	PlacedAt time.Time `json:"placedAt" xml:"placedAt"`
	// Bet Status.   ACCEPTED = Bet was accepted,   CANCELLED = Bet is cancelled as per Pinnacle betting rules,   LOSE = The bet is settled as lose,   REFUNDED = When an event is cancelled or when the bet is settled as push, the bet will have REFUNDED status,   WON = The bet is settled as won
	BetStatus string `json:"betStatus" xml:"betStatus"`
	BetType   string `json:"betType" xml:"betType"`
	// Win amount.
	Win float64 `json:"win" xml:"win"`
	// Risk amount.
	Risk float64 `json:"risk" xml:"risk"`
	// Win-Loss for settled bets.
	WinLoss *float64 `json:"winLoss,omitempty" xml:"winLoss"`
	// Update Sequence
	UpdateSequence int64 `json:"updateSequence" xml:"updateSequence"`
	// Manual bet description.
	Description string `json:"description" xml:"description"`
	// Referenced original bet id.
	ReferenceBetId *NullableInt64 `json:"referenceBetId,omitempty" xml:"referenceBetId"`
}

// GetBetId returns the BetId field value
func (o *ManualBet) GetBetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.BetId
}

// SetBetId sets field value
func (o *ManualBet) SetBetId(v int64) {
	o.BetId = v
}

// GetWagerNumber returns the WagerNumber field value
func (o *ManualBet) GetWagerNumber() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.WagerNumber
}

// SetWagerNumber sets field value
func (o *ManualBet) SetWagerNumber(v int32) {
	o.WagerNumber = v
}

// GetPlacedAt returns the PlacedAt field value
func (o *ManualBet) GetPlacedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.PlacedAt
}

// SetPlacedAt sets field value
func (o *ManualBet) SetPlacedAt(v time.Time) {
	o.PlacedAt = v
}

// GetBetStatus returns the BetStatus field value
func (o *ManualBet) GetBetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BetStatus
}

// SetBetStatus sets field value
func (o *ManualBet) SetBetStatus(v string) {
	o.BetStatus = v
}

// GetBetType returns the BetType field value
func (o *ManualBet) GetBetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BetType
}

// SetBetType sets field value
func (o *ManualBet) SetBetType(v string) {
	o.BetType = v
}

// GetWin returns the Win field value
func (o *ManualBet) GetWin() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Win
}

// SetWin sets field value
func (o *ManualBet) SetWin(v float64) {
	o.Win = v
}

// GetRisk returns the Risk field value
func (o *ManualBet) GetRisk() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Risk
}

// SetRisk sets field value
func (o *ManualBet) SetRisk(v float64) {
	o.Risk = v
}

// GetWinLoss returns the WinLoss field value if set, zero value otherwise.
func (o *ManualBet) GetWinLoss() float64 {
	if o == nil || o.WinLoss == nil {
		var ret float64
		return ret
	}
	return *o.WinLoss
}

// GetWinLossOk returns a tuple with the WinLoss field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ManualBet) GetWinLossOk() (float64, bool) {
	if o == nil || o.WinLoss == nil {
		var ret float64
		return ret, false
	}
	return *o.WinLoss, true
}

// HasWinLoss returns a boolean if a field has been set.
func (o *ManualBet) HasWinLoss() bool {
	if o != nil && o.WinLoss != nil {
		return true
	}

	return false
}

// SetWinLoss gets a reference to the given float64 and assigns it to the WinLoss field.
func (o *ManualBet) SetWinLoss(v float64) {
	o.WinLoss = &v
}

// GetUpdateSequence returns the UpdateSequence field value
func (o *ManualBet) GetUpdateSequence() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UpdateSequence
}

// SetUpdateSequence sets field value
func (o *ManualBet) SetUpdateSequence(v int64) {
	o.UpdateSequence = v
}

// GetDescription returns the Description field value
func (o *ManualBet) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// SetDescription sets field value
func (o *ManualBet) SetDescription(v string) {
	o.Description = v
}

// GetReferenceBetId returns the ReferenceBetId field value if set, zero value otherwise.
func (o *ManualBet) GetReferenceBetId() NullableInt64 {
	if o == nil || o.ReferenceBetId == nil {
		var ret NullableInt64
		return ret
	}
	return *o.ReferenceBetId
}

// GetReferenceBetIdOk returns a tuple with the ReferenceBetId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ManualBet) GetReferenceBetIdOk() (NullableInt64, bool) {
	if o == nil || o.ReferenceBetId == nil {
		var ret NullableInt64
		return ret, false
	}
	return *o.ReferenceBetId, true
}

// HasReferenceBetId returns a boolean if a field has been set.
func (o *ManualBet) HasReferenceBetId() bool {
	if o != nil && o.ReferenceBetId != nil {
		return true
	}

	return false
}

// SetReferenceBetId gets a reference to the given NullableInt64 and assigns it to the ReferenceBetId field.
func (o *ManualBet) SetReferenceBetId(v NullableInt64) {
	o.ReferenceBetId = &v
}

type NullableManualBet struct {
	Value        ManualBet
	ExplicitNull bool
}

func (v NullableManualBet) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableManualBet) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
