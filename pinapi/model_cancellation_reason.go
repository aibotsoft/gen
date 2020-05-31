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

// CancellationReason Possible keys \\:   * correctTeam1Id * correctTeam2Id * correctListedPitcher1 * correctListedPitcher2 * correctSpread * correctTotalPoints * correctTeam1TotalPoints * correctTeam2TotalPoints * correctTeam1Score * correctTeam2Score * correctTeam1TennisSetsScore * correctTeam2TennisSetsScore 
type CancellationReason struct {
	Code string `json:"code"`
	Details *[]CancellationDetailsItem `json:"details,omitempty"`
}

// NewCancellationReason instantiates a new CancellationReason object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCancellationReason(code string, ) *CancellationReason {
	this := CancellationReason{}
	this.Code = code
	return &this
}

// NewCancellationReasonWithDefaults instantiates a new CancellationReason object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCancellationReasonWithDefaults() *CancellationReason {
	this := CancellationReason{}
	return &this
}

// GetCode returns the Code field value
func (o *CancellationReason) GetCode() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *CancellationReason) GetCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *CancellationReason) SetCode(v string) {
	o.Code = v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *CancellationReason) GetDetails() []CancellationDetailsItem {
	if o == nil || o.Details == nil {
		var ret []CancellationDetailsItem
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancellationReason) GetDetailsOk() (*[]CancellationDetailsItem, bool) {
	if o == nil || o.Details == nil {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *CancellationReason) HasDetails() bool {
	if o != nil && o.Details != nil {
		return true
	}

	return false
}

// SetDetails gets a reference to the given []CancellationDetailsItem and assigns it to the Details field.
func (o *CancellationReason) SetDetails(v []CancellationDetailsItem) {
	o.Details = &v
}

func (o CancellationReason) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["code"] = o.Code
	}
	if o.Details != nil {
		toSerialize["details"] = o.Details
	}
	return json.Marshal(toSerialize)
}

type NullableCancellationReason struct {
	value *CancellationReason
	isSet bool
}

func (v NullableCancellationReason) Get() *CancellationReason {
	return v.value
}

func (v *NullableCancellationReason) Set(val *CancellationReason) {
	v.value = val
	v.isSet = true
}

func (v NullableCancellationReason) IsSet() bool {
	return v.isSet
}

func (v *NullableCancellationReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCancellationReason(val *CancellationReason) *NullableCancellationReason {
	return &NullableCancellationReason{value: val, isSet: true}
}

func (v NullableCancellationReason) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCancellationReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
