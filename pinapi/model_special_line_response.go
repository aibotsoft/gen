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

// SpecialLineResponse struct for SpecialLineResponse
type SpecialLineResponse struct {
	// Status [SUCCESS = OK, NOT_EXISTS = Line not offered anymore]
	Status *string `json:"status,omitempty"`
	// Special Id.
	SpecialId *int64 `json:"specialId,omitempty"`
	// Contestant Id.
	ContestantId *int64 `json:"contestantId,omitempty"`
	// Minimum bettable risk amount.
	MinRiskStake *float64 `json:"minRiskStake,omitempty"`
	// Maximum bettable risk amount.
	MaxRiskStake *float64 `json:"maxRiskStake,omitempty"`
	// Minimum bettable win amount.
	MinWinStake *float64 `json:"minWinStake,omitempty"`
	// Maximum bettable win amount.
	MaxWinStake *float64 `json:"maxWinStake,omitempty"`
	// Line identification needed to place a bet.
	LineId *int64 `json:"lineId,omitempty"`
	// Latest price.
	Price *float64 `json:"price,omitempty"`
	// Handicap.
	Handicap *float64 `json:"handicap,omitempty"`
}

// NewSpecialLineResponse instantiates a new SpecialLineResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpecialLineResponse() *SpecialLineResponse {
	this := SpecialLineResponse{}
	return &this
}

// NewSpecialLineResponseWithDefaults instantiates a new SpecialLineResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpecialLineResponseWithDefaults() *SpecialLineResponse {
	this := SpecialLineResponse{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SpecialLineResponse) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpecialLineResponse) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SpecialLineResponse) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *SpecialLineResponse) SetStatus(v string) {
	o.Status = &v
}

// GetSpecialId returns the SpecialId field value if set, zero value otherwise.
func (o *SpecialLineResponse) GetSpecialId() int64 {
	if o == nil || o.SpecialId == nil {
		var ret int64
		return ret
	}
	return *o.SpecialId
}

// GetSpecialIdOk returns a tuple with the SpecialId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpecialLineResponse) GetSpecialIdOk() (*int64, bool) {
	if o == nil || o.SpecialId == nil {
		return nil, false
	}
	return o.SpecialId, true
}

// HasSpecialId returns a boolean if a field has been set.
func (o *SpecialLineResponse) HasSpecialId() bool {
	if o != nil && o.SpecialId != nil {
		return true
	}

	return false
}

// SetSpecialId gets a reference to the given int64 and assigns it to the SpecialId field.
func (o *SpecialLineResponse) SetSpecialId(v int64) {
	o.SpecialId = &v
}

// GetContestantId returns the ContestantId field value if set, zero value otherwise.
func (o *SpecialLineResponse) GetContestantId() int64 {
	if o == nil || o.ContestantId == nil {
		var ret int64
		return ret
	}
	return *o.ContestantId
}

// GetContestantIdOk returns a tuple with the ContestantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpecialLineResponse) GetContestantIdOk() (*int64, bool) {
	if o == nil || o.ContestantId == nil {
		return nil, false
	}
	return o.ContestantId, true
}

// HasContestantId returns a boolean if a field has been set.
func (o *SpecialLineResponse) HasContestantId() bool {
	if o != nil && o.ContestantId != nil {
		return true
	}

	return false
}

// SetContestantId gets a reference to the given int64 and assigns it to the ContestantId field.
func (o *SpecialLineResponse) SetContestantId(v int64) {
	o.ContestantId = &v
}

// GetMinRiskStake returns the MinRiskStake field value if set, zero value otherwise.
func (o *SpecialLineResponse) GetMinRiskStake() float64 {
	if o == nil || o.MinRiskStake == nil {
		var ret float64
		return ret
	}
	return *o.MinRiskStake
}

// GetMinRiskStakeOk returns a tuple with the MinRiskStake field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpecialLineResponse) GetMinRiskStakeOk() (*float64, bool) {
	if o == nil || o.MinRiskStake == nil {
		return nil, false
	}
	return o.MinRiskStake, true
}

// HasMinRiskStake returns a boolean if a field has been set.
func (o *SpecialLineResponse) HasMinRiskStake() bool {
	if o != nil && o.MinRiskStake != nil {
		return true
	}

	return false
}

// SetMinRiskStake gets a reference to the given float64 and assigns it to the MinRiskStake field.
func (o *SpecialLineResponse) SetMinRiskStake(v float64) {
	o.MinRiskStake = &v
}

// GetMaxRiskStake returns the MaxRiskStake field value if set, zero value otherwise.
func (o *SpecialLineResponse) GetMaxRiskStake() float64 {
	if o == nil || o.MaxRiskStake == nil {
		var ret float64
		return ret
	}
	return *o.MaxRiskStake
}

// GetMaxRiskStakeOk returns a tuple with the MaxRiskStake field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpecialLineResponse) GetMaxRiskStakeOk() (*float64, bool) {
	if o == nil || o.MaxRiskStake == nil {
		return nil, false
	}
	return o.MaxRiskStake, true
}

// HasMaxRiskStake returns a boolean if a field has been set.
func (o *SpecialLineResponse) HasMaxRiskStake() bool {
	if o != nil && o.MaxRiskStake != nil {
		return true
	}

	return false
}

// SetMaxRiskStake gets a reference to the given float64 and assigns it to the MaxRiskStake field.
func (o *SpecialLineResponse) SetMaxRiskStake(v float64) {
	o.MaxRiskStake = &v
}

// GetMinWinStake returns the MinWinStake field value if set, zero value otherwise.
func (o *SpecialLineResponse) GetMinWinStake() float64 {
	if o == nil || o.MinWinStake == nil {
		var ret float64
		return ret
	}
	return *o.MinWinStake
}

// GetMinWinStakeOk returns a tuple with the MinWinStake field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpecialLineResponse) GetMinWinStakeOk() (*float64, bool) {
	if o == nil || o.MinWinStake == nil {
		return nil, false
	}
	return o.MinWinStake, true
}

// HasMinWinStake returns a boolean if a field has been set.
func (o *SpecialLineResponse) HasMinWinStake() bool {
	if o != nil && o.MinWinStake != nil {
		return true
	}

	return false
}

// SetMinWinStake gets a reference to the given float64 and assigns it to the MinWinStake field.
func (o *SpecialLineResponse) SetMinWinStake(v float64) {
	o.MinWinStake = &v
}

// GetMaxWinStake returns the MaxWinStake field value if set, zero value otherwise.
func (o *SpecialLineResponse) GetMaxWinStake() float64 {
	if o == nil || o.MaxWinStake == nil {
		var ret float64
		return ret
	}
	return *o.MaxWinStake
}

// GetMaxWinStakeOk returns a tuple with the MaxWinStake field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpecialLineResponse) GetMaxWinStakeOk() (*float64, bool) {
	if o == nil || o.MaxWinStake == nil {
		return nil, false
	}
	return o.MaxWinStake, true
}

// HasMaxWinStake returns a boolean if a field has been set.
func (o *SpecialLineResponse) HasMaxWinStake() bool {
	if o != nil && o.MaxWinStake != nil {
		return true
	}

	return false
}

// SetMaxWinStake gets a reference to the given float64 and assigns it to the MaxWinStake field.
func (o *SpecialLineResponse) SetMaxWinStake(v float64) {
	o.MaxWinStake = &v
}

// GetLineId returns the LineId field value if set, zero value otherwise.
func (o *SpecialLineResponse) GetLineId() int64 {
	if o == nil || o.LineId == nil {
		var ret int64
		return ret
	}
	return *o.LineId
}

// GetLineIdOk returns a tuple with the LineId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpecialLineResponse) GetLineIdOk() (*int64, bool) {
	if o == nil || o.LineId == nil {
		return nil, false
	}
	return o.LineId, true
}

// HasLineId returns a boolean if a field has been set.
func (o *SpecialLineResponse) HasLineId() bool {
	if o != nil && o.LineId != nil {
		return true
	}

	return false
}

// SetLineId gets a reference to the given int64 and assigns it to the LineId field.
func (o *SpecialLineResponse) SetLineId(v int64) {
	o.LineId = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *SpecialLineResponse) GetPrice() float64 {
	if o == nil || o.Price == nil {
		var ret float64
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpecialLineResponse) GetPriceOk() (*float64, bool) {
	if o == nil || o.Price == nil {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *SpecialLineResponse) HasPrice() bool {
	if o != nil && o.Price != nil {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float64 and assigns it to the Price field.
func (o *SpecialLineResponse) SetPrice(v float64) {
	o.Price = &v
}

// GetHandicap returns the Handicap field value if set, zero value otherwise.
func (o *SpecialLineResponse) GetHandicap() float64 {
	if o == nil || o.Handicap == nil {
		var ret float64
		return ret
	}
	return *o.Handicap
}

// GetHandicapOk returns a tuple with the Handicap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpecialLineResponse) GetHandicapOk() (*float64, bool) {
	if o == nil || o.Handicap == nil {
		return nil, false
	}
	return o.Handicap, true
}

// HasHandicap returns a boolean if a field has been set.
func (o *SpecialLineResponse) HasHandicap() bool {
	if o != nil && o.Handicap != nil {
		return true
	}

	return false
}

// SetHandicap gets a reference to the given float64 and assigns it to the Handicap field.
func (o *SpecialLineResponse) SetHandicap(v float64) {
	o.Handicap = &v
}

func (o SpecialLineResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.SpecialId != nil {
		toSerialize["specialId"] = o.SpecialId
	}
	if o.ContestantId != nil {
		toSerialize["contestantId"] = o.ContestantId
	}
	if o.MinRiskStake != nil {
		toSerialize["minRiskStake"] = o.MinRiskStake
	}
	if o.MaxRiskStake != nil {
		toSerialize["maxRiskStake"] = o.MaxRiskStake
	}
	if o.MinWinStake != nil {
		toSerialize["minWinStake"] = o.MinWinStake
	}
	if o.MaxWinStake != nil {
		toSerialize["maxWinStake"] = o.MaxWinStake
	}
	if o.LineId != nil {
		toSerialize["lineId"] = o.LineId
	}
	if o.Price != nil {
		toSerialize["price"] = o.Price
	}
	if o.Handicap != nil {
		toSerialize["handicap"] = o.Handicap
	}
	return json.Marshal(toSerialize)
}

type NullableSpecialLineResponse struct {
	value *SpecialLineResponse
	isSet bool
}

func (v NullableSpecialLineResponse) Get() *SpecialLineResponse {
	return v.value
}

func (v *NullableSpecialLineResponse) Set(val *SpecialLineResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSpecialLineResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSpecialLineResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpecialLineResponse(val *SpecialLineResponse) *NullableSpecialLineResponse {
	return &NullableSpecialLineResponse{value: val, isSet: true}
}

func (v NullableSpecialLineResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpecialLineResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
