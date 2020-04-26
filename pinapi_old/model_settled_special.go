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

// SettledSpecial Settled Special
type SettledSpecial struct {
	// Id for the Settled Special
	Id *int64 `json:"id,omitempty" xml:"id"`
	// Status of the settled special.
	Status *int32 `json:"status,omitempty" xml:"status"`
	// Id for the Settled Special
	SettlementId *int64 `json:"settlementId,omitempty" xml:"settlementId"`
	// Settled DateTime
	SettledAt          *time.Time          `json:"settledAt,omitempty" xml:"settledAt"`
	CancellationReason *CancellationReason `json:"cancellationReason,omitempty" xml:"cancellationReason"`
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SettledSpecial) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SettledSpecial) GetIdOk() (int64, bool) {
	if o == nil || o.Id == nil {
		var ret int64
		return ret, false
	}
	return *o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SettledSpecial) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *SettledSpecial) SetId(v int64) {
	o.Id = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SettledSpecial) GetStatus() int32 {
	if o == nil || o.Status == nil {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SettledSpecial) GetStatusOk() (int32, bool) {
	if o == nil || o.Status == nil {
		var ret int32
		return ret, false
	}
	return *o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SettledSpecial) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *SettledSpecial) SetStatus(v int32) {
	o.Status = &v
}

// GetSettlementId returns the SettlementId field value if set, zero value otherwise.
func (o *SettledSpecial) GetSettlementId() int64 {
	if o == nil || o.SettlementId == nil {
		var ret int64
		return ret
	}
	return *o.SettlementId
}

// GetSettlementIdOk returns a tuple with the SettlementId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SettledSpecial) GetSettlementIdOk() (int64, bool) {
	if o == nil || o.SettlementId == nil {
		var ret int64
		return ret, false
	}
	return *o.SettlementId, true
}

// HasSettlementId returns a boolean if a field has been set.
func (o *SettledSpecial) HasSettlementId() bool {
	if o != nil && o.SettlementId != nil {
		return true
	}

	return false
}

// SetSettlementId gets a reference to the given int64 and assigns it to the SettlementId field.
func (o *SettledSpecial) SetSettlementId(v int64) {
	o.SettlementId = &v
}

// GetSettledAt returns the SettledAt field value if set, zero value otherwise.
func (o *SettledSpecial) GetSettledAt() time.Time {
	if o == nil || o.SettledAt == nil {
		var ret time.Time
		return ret
	}
	return *o.SettledAt
}

// GetSettledAtOk returns a tuple with the SettledAt field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SettledSpecial) GetSettledAtOk() (time.Time, bool) {
	if o == nil || o.SettledAt == nil {
		var ret time.Time
		return ret, false
	}
	return *o.SettledAt, true
}

// HasSettledAt returns a boolean if a field has been set.
func (o *SettledSpecial) HasSettledAt() bool {
	if o != nil && o.SettledAt != nil {
		return true
	}

	return false
}

// SetSettledAt gets a reference to the given time.Time and assigns it to the SettledAt field.
func (o *SettledSpecial) SetSettledAt(v time.Time) {
	o.SettledAt = &v
}

// GetCancellationReason returns the CancellationReason field value if set, zero value otherwise.
func (o *SettledSpecial) GetCancellationReason() CancellationReason {
	if o == nil || o.CancellationReason == nil {
		var ret CancellationReason
		return ret
	}
	return *o.CancellationReason
}

// GetCancellationReasonOk returns a tuple with the CancellationReason field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SettledSpecial) GetCancellationReasonOk() (CancellationReason, bool) {
	if o == nil || o.CancellationReason == nil {
		var ret CancellationReason
		return ret, false
	}
	return *o.CancellationReason, true
}

// HasCancellationReason returns a boolean if a field has been set.
func (o *SettledSpecial) HasCancellationReason() bool {
	if o != nil && o.CancellationReason != nil {
		return true
	}

	return false
}

// SetCancellationReason gets a reference to the given CancellationReason and assigns it to the CancellationReason field.
func (o *SettledSpecial) SetCancellationReason(v CancellationReason) {
	o.CancellationReason = &v
}

type NullableSettledSpecial struct {
	Value        SettledSpecial
	ExplicitNull bool
}

func (v NullableSettledSpecial) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableSettledSpecial) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
