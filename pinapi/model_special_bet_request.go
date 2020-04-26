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

// SpecialBetRequest struct for SpecialBetRequest
type SpecialBetRequest struct {
	// This unique id of the place bet requests. This is to support idempotent requests.
	UniqueRequestId *string `json:"uniqueRequestId,omitempty"`
	// Whether or not to accept a bet when there is a line change in favor of the client.
	AcceptBetterLine *bool `json:"acceptBetterLine,omitempty"`
	OddsFormat *OddsFormat `json:"oddsFormat,omitempty"`
	// amount in client’s currency.
	Stake *float64 `json:"stake,omitempty"`
	// Whether the stake amount is risk or win amount.
	WinRiskStake *string `json:"winRiskStake,omitempty"`
	// Line identification.
	LineId *int64 `json:"lineId,omitempty"`
	// Special identification.
	SpecialId *int64 `json:"specialId,omitempty"`
	// Contestant identification.
	ContestantId *int64 `json:"contestantId,omitempty"`
}

// NewSpecialBetRequest instantiates a new SpecialBetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpecialBetRequest() *SpecialBetRequest {
	this := SpecialBetRequest{}
	return &this
}

// NewSpecialBetRequestWithDefaults instantiates a new SpecialBetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpecialBetRequestWithDefaults() *SpecialBetRequest {
	this := SpecialBetRequest{}
	return &this
}

// GetUniqueRequestId returns the UniqueRequestId field value if set, zero value otherwise.
func (o *SpecialBetRequest) GetUniqueRequestId() string {
	if o == nil || o.UniqueRequestId == nil {
		var ret string
		return ret
	}
	return *o.UniqueRequestId
}

// GetUniqueRequestIdOk returns a tuple with the UniqueRequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpecialBetRequest) GetUniqueRequestIdOk() (*string, bool) {
	if o == nil || o.UniqueRequestId == nil {
		return nil, false
	}
	return o.UniqueRequestId, true
}

// HasUniqueRequestId returns a boolean if a field has been set.
func (o *SpecialBetRequest) HasUniqueRequestId() bool {
	if o != nil && o.UniqueRequestId != nil {
		return true
	}

	return false
}

// SetUniqueRequestId gets a reference to the given string and assigns it to the UniqueRequestId field.
func (o *SpecialBetRequest) SetUniqueRequestId(v string) {
	o.UniqueRequestId = &v
}

// GetAcceptBetterLine returns the AcceptBetterLine field value if set, zero value otherwise.
func (o *SpecialBetRequest) GetAcceptBetterLine() bool {
	if o == nil || o.AcceptBetterLine == nil {
		var ret bool
		return ret
	}
	return *o.AcceptBetterLine
}

// GetAcceptBetterLineOk returns a tuple with the AcceptBetterLine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpecialBetRequest) GetAcceptBetterLineOk() (*bool, bool) {
	if o == nil || o.AcceptBetterLine == nil {
		return nil, false
	}
	return o.AcceptBetterLine, true
}

// HasAcceptBetterLine returns a boolean if a field has been set.
func (o *SpecialBetRequest) HasAcceptBetterLine() bool {
	if o != nil && o.AcceptBetterLine != nil {
		return true
	}

	return false
}

// SetAcceptBetterLine gets a reference to the given bool and assigns it to the AcceptBetterLine field.
func (o *SpecialBetRequest) SetAcceptBetterLine(v bool) {
	o.AcceptBetterLine = &v
}

// GetOddsFormat returns the OddsFormat field value if set, zero value otherwise.
func (o *SpecialBetRequest) GetOddsFormat() OddsFormat {
	if o == nil || o.OddsFormat == nil {
		var ret OddsFormat
		return ret
	}
	return *o.OddsFormat
}

// GetOddsFormatOk returns a tuple with the OddsFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpecialBetRequest) GetOddsFormatOk() (*OddsFormat, bool) {
	if o == nil || o.OddsFormat == nil {
		return nil, false
	}
	return o.OddsFormat, true
}

// HasOddsFormat returns a boolean if a field has been set.
func (o *SpecialBetRequest) HasOddsFormat() bool {
	if o != nil && o.OddsFormat != nil {
		return true
	}

	return false
}

// SetOddsFormat gets a reference to the given OddsFormat and assigns it to the OddsFormat field.
func (o *SpecialBetRequest) SetOddsFormat(v OddsFormat) {
	o.OddsFormat = &v
}

// GetStake returns the Stake field value if set, zero value otherwise.
func (o *SpecialBetRequest) GetStake() float64 {
	if o == nil || o.Stake == nil {
		var ret float64
		return ret
	}
	return *o.Stake
}

// GetStakeOk returns a tuple with the Stake field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpecialBetRequest) GetStakeOk() (*float64, bool) {
	if o == nil || o.Stake == nil {
		return nil, false
	}
	return o.Stake, true
}

// HasStake returns a boolean if a field has been set.
func (o *SpecialBetRequest) HasStake() bool {
	if o != nil && o.Stake != nil {
		return true
	}

	return false
}

// SetStake gets a reference to the given float64 and assigns it to the Stake field.
func (o *SpecialBetRequest) SetStake(v float64) {
	o.Stake = &v
}

// GetWinRiskStake returns the WinRiskStake field value if set, zero value otherwise.
func (o *SpecialBetRequest) GetWinRiskStake() string {
	if o == nil || o.WinRiskStake == nil {
		var ret string
		return ret
	}
	return *o.WinRiskStake
}

// GetWinRiskStakeOk returns a tuple with the WinRiskStake field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpecialBetRequest) GetWinRiskStakeOk() (*string, bool) {
	if o == nil || o.WinRiskStake == nil {
		return nil, false
	}
	return o.WinRiskStake, true
}

// HasWinRiskStake returns a boolean if a field has been set.
func (o *SpecialBetRequest) HasWinRiskStake() bool {
	if o != nil && o.WinRiskStake != nil {
		return true
	}

	return false
}

// SetWinRiskStake gets a reference to the given string and assigns it to the WinRiskStake field.
func (o *SpecialBetRequest) SetWinRiskStake(v string) {
	o.WinRiskStake = &v
}

// GetLineId returns the LineId field value if set, zero value otherwise.
func (o *SpecialBetRequest) GetLineId() int64 {
	if o == nil || o.LineId == nil {
		var ret int64
		return ret
	}
	return *o.LineId
}

// GetLineIdOk returns a tuple with the LineId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpecialBetRequest) GetLineIdOk() (*int64, bool) {
	if o == nil || o.LineId == nil {
		return nil, false
	}
	return o.LineId, true
}

// HasLineId returns a boolean if a field has been set.
func (o *SpecialBetRequest) HasLineId() bool {
	if o != nil && o.LineId != nil {
		return true
	}

	return false
}

// SetLineId gets a reference to the given int64 and assigns it to the LineId field.
func (o *SpecialBetRequest) SetLineId(v int64) {
	o.LineId = &v
}

// GetSpecialId returns the SpecialId field value if set, zero value otherwise.
func (o *SpecialBetRequest) GetSpecialId() int64 {
	if o == nil || o.SpecialId == nil {
		var ret int64
		return ret
	}
	return *o.SpecialId
}

// GetSpecialIdOk returns a tuple with the SpecialId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpecialBetRequest) GetSpecialIdOk() (*int64, bool) {
	if o == nil || o.SpecialId == nil {
		return nil, false
	}
	return o.SpecialId, true
}

// HasSpecialId returns a boolean if a field has been set.
func (o *SpecialBetRequest) HasSpecialId() bool {
	if o != nil && o.SpecialId != nil {
		return true
	}

	return false
}

// SetSpecialId gets a reference to the given int64 and assigns it to the SpecialId field.
func (o *SpecialBetRequest) SetSpecialId(v int64) {
	o.SpecialId = &v
}

// GetContestantId returns the ContestantId field value if set, zero value otherwise.
func (o *SpecialBetRequest) GetContestantId() int64 {
	if o == nil || o.ContestantId == nil {
		var ret int64
		return ret
	}
	return *o.ContestantId
}

// GetContestantIdOk returns a tuple with the ContestantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpecialBetRequest) GetContestantIdOk() (*int64, bool) {
	if o == nil || o.ContestantId == nil {
		return nil, false
	}
	return o.ContestantId, true
}

// HasContestantId returns a boolean if a field has been set.
func (o *SpecialBetRequest) HasContestantId() bool {
	if o != nil && o.ContestantId != nil {
		return true
	}

	return false
}

// SetContestantId gets a reference to the given int64 and assigns it to the ContestantId field.
func (o *SpecialBetRequest) SetContestantId(v int64) {
	o.ContestantId = &v
}

func (o SpecialBetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UniqueRequestId != nil {
		toSerialize["uniqueRequestId"] = o.UniqueRequestId
	}
	if o.AcceptBetterLine != nil {
		toSerialize["acceptBetterLine"] = o.AcceptBetterLine
	}
	if o.OddsFormat != nil {
		toSerialize["oddsFormat"] = o.OddsFormat
	}
	if o.Stake != nil {
		toSerialize["stake"] = o.Stake
	}
	if o.WinRiskStake != nil {
		toSerialize["winRiskStake"] = o.WinRiskStake
	}
	if o.LineId != nil {
		toSerialize["lineId"] = o.LineId
	}
	if o.SpecialId != nil {
		toSerialize["specialId"] = o.SpecialId
	}
	if o.ContestantId != nil {
		toSerialize["contestantId"] = o.ContestantId
	}
	return json.Marshal(toSerialize)
}

type NullableSpecialBetRequest struct {
	value *SpecialBetRequest
	isSet bool
}

func (v NullableSpecialBetRequest) Get() *SpecialBetRequest {
	return v.value
}

func (v *NullableSpecialBetRequest) Set(val *SpecialBetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSpecialBetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSpecialBetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpecialBetRequest(val *SpecialBetRequest) *NullableSpecialBetRequest {
	return &NullableSpecialBetRequest{value: val, isSet: true}
}

func (v NullableSpecialBetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpecialBetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
