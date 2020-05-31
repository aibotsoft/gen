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
	"time"
)

// SettledSpecial Settled Special
type SettledSpecial struct {
	// Id for the Settled Special
	Id *int64 `json:"id,omitempty"`
	// Status of the settled special.
	Status *int `json:"status,omitempty"`
	// Id for the Settled Special
	SettlementId *int64 `json:"settlementId,omitempty"`
	// Settled DateTime
	SettledAt *time.Time `json:"settledAt,omitempty"`
	CancellationReason *CancellationReason `json:"cancellationReason,omitempty"`
}

// NewSettledSpecial instantiates a new SettledSpecial object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSettledSpecial() *SettledSpecial {
	this := SettledSpecial{}
	return &this
}

// NewSettledSpecialWithDefaults instantiates a new SettledSpecial object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSettledSpecialWithDefaults() *SettledSpecial {
	this := SettledSpecial{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SettledSpecial) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettledSpecial) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
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
func (o *SettledSpecial) GetStatus() int {
	if o == nil || o.Status == nil {
		var ret int
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettledSpecial) GetStatusOk() (*int, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SettledSpecial) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int and assigns it to the Status field.
func (o *SettledSpecial) SetStatus(v int) {
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

// GetSettlementIdOk returns a tuple with the SettlementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettledSpecial) GetSettlementIdOk() (*int64, bool) {
	if o == nil || o.SettlementId == nil {
		return nil, false
	}
	return o.SettlementId, true
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

// GetSettledAtOk returns a tuple with the SettledAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettledSpecial) GetSettledAtOk() (*time.Time, bool) {
	if o == nil || o.SettledAt == nil {
		return nil, false
	}
	return o.SettledAt, true
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

// GetCancellationReasonOk returns a tuple with the CancellationReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettledSpecial) GetCancellationReasonOk() (*CancellationReason, bool) {
	if o == nil || o.CancellationReason == nil {
		return nil, false
	}
	return o.CancellationReason, true
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

func (o SettledSpecial) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.SettlementId != nil {
		toSerialize["settlementId"] = o.SettlementId
	}
	if o.SettledAt != nil {
		toSerialize["settledAt"] = o.SettledAt
	}
	if o.CancellationReason != nil {
		toSerialize["cancellationReason"] = o.CancellationReason
	}
	return json.Marshal(toSerialize)
}

type NullableSettledSpecial struct {
	value *SettledSpecial
	isSet bool
}

func (v NullableSettledSpecial) Get() *SettledSpecial {
	return v.value
}

func (v *NullableSettledSpecial) Set(val *SettledSpecial) {
	v.value = val
	v.isSet = true
}

func (v NullableSettledSpecial) IsSet() bool {
	return v.isSet
}

func (v *NullableSettledSpecial) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSettledSpecial(val *SettledSpecial) *NullableSettledSpecial {
	return &NullableSettledSpecial{value: val, isSet: true}
}

func (v NullableSettledSpecial) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSettledSpecial) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
