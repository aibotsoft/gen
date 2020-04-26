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

// GetBetsByTypeResponseV3 struct for GetBetsByTypeResponseV3
type GetBetsByTypeResponseV3 struct {
	// Whether there are more pages available.
	MoreAvailable *bool `json:"moreAvailable,omitempty"`
	// Page size. Default is 1000.
	PageSize *int `json:"pageSize,omitempty"`
	// Starting record number of the result set. Records start at zero
	FromRecord *int `json:"fromRecord,omitempty"`
	// Ending record number of the result set.
	ToRecord *int `json:"toRecord,omitempty"`
	// A collection of placed straight bets.
	StraightBets *[]StraightBetV3 `json:"straightBets,omitempty"`
	// A collection of placed parlay bets.
	ParlayBets *[]ParlayBet `json:"parlayBets,omitempty"`
	// A collection of placed teaser bets.
	TeaserBets *[]TeaserBet `json:"teaserBets,omitempty"`
	// A collection of placed special bets.
	SpecialBets *[]SpecialBet `json:"specialBets,omitempty"`
	// A collection of placed manual bets.
	ManualBets *[]ManualBet `json:"manualBets,omitempty"`
}

// NewGetBetsByTypeResponseV3 instantiates a new GetBetsByTypeResponseV3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBetsByTypeResponseV3() *GetBetsByTypeResponseV3 {
	this := GetBetsByTypeResponseV3{}
	return &this
}

// NewGetBetsByTypeResponseV3WithDefaults instantiates a new GetBetsByTypeResponseV3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBetsByTypeResponseV3WithDefaults() *GetBetsByTypeResponseV3 {
	this := GetBetsByTypeResponseV3{}
	return &this
}

// GetMoreAvailable returns the MoreAvailable field value if set, zero value otherwise.
func (o *GetBetsByTypeResponseV3) GetMoreAvailable() bool {
	if o == nil || o.MoreAvailable == nil {
		var ret bool
		return ret
	}
	return *o.MoreAvailable
}

// GetMoreAvailableOk returns a tuple with the MoreAvailable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBetsByTypeResponseV3) GetMoreAvailableOk() (*bool, bool) {
	if o == nil || o.MoreAvailable == nil {
		return nil, false
	}
	return o.MoreAvailable, true
}

// HasMoreAvailable returns a boolean if a field has been set.
func (o *GetBetsByTypeResponseV3) HasMoreAvailable() bool {
	if o != nil && o.MoreAvailable != nil {
		return true
	}

	return false
}

// SetMoreAvailable gets a reference to the given bool and assigns it to the MoreAvailable field.
func (o *GetBetsByTypeResponseV3) SetMoreAvailable(v bool) {
	o.MoreAvailable = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *GetBetsByTypeResponseV3) GetPageSize() int {
	if o == nil || o.PageSize == nil {
		var ret int
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBetsByTypeResponseV3) GetPageSizeOk() (*int, bool) {
	if o == nil || o.PageSize == nil {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *GetBetsByTypeResponseV3) HasPageSize() bool {
	if o != nil && o.PageSize != nil {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int and assigns it to the PageSize field.
func (o *GetBetsByTypeResponseV3) SetPageSize(v int) {
	o.PageSize = &v
}

// GetFromRecord returns the FromRecord field value if set, zero value otherwise.
func (o *GetBetsByTypeResponseV3) GetFromRecord() int {
	if o == nil || o.FromRecord == nil {
		var ret int
		return ret
	}
	return *o.FromRecord
}

// GetFromRecordOk returns a tuple with the FromRecord field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBetsByTypeResponseV3) GetFromRecordOk() (*int, bool) {
	if o == nil || o.FromRecord == nil {
		return nil, false
	}
	return o.FromRecord, true
}

// HasFromRecord returns a boolean if a field has been set.
func (o *GetBetsByTypeResponseV3) HasFromRecord() bool {
	if o != nil && o.FromRecord != nil {
		return true
	}

	return false
}

// SetFromRecord gets a reference to the given int and assigns it to the FromRecord field.
func (o *GetBetsByTypeResponseV3) SetFromRecord(v int) {
	o.FromRecord = &v
}

// GetToRecord returns the ToRecord field value if set, zero value otherwise.
func (o *GetBetsByTypeResponseV3) GetToRecord() int {
	if o == nil || o.ToRecord == nil {
		var ret int
		return ret
	}
	return *o.ToRecord
}

// GetToRecordOk returns a tuple with the ToRecord field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBetsByTypeResponseV3) GetToRecordOk() (*int, bool) {
	if o == nil || o.ToRecord == nil {
		return nil, false
	}
	return o.ToRecord, true
}

// HasToRecord returns a boolean if a field has been set.
func (o *GetBetsByTypeResponseV3) HasToRecord() bool {
	if o != nil && o.ToRecord != nil {
		return true
	}

	return false
}

// SetToRecord gets a reference to the given int and assigns it to the ToRecord field.
func (o *GetBetsByTypeResponseV3) SetToRecord(v int) {
	o.ToRecord = &v
}

// GetStraightBets returns the StraightBets field value if set, zero value otherwise.
func (o *GetBetsByTypeResponseV3) GetStraightBets() []StraightBetV3 {
	if o == nil || o.StraightBets == nil {
		var ret []StraightBetV3
		return ret
	}
	return *o.StraightBets
}

// GetStraightBetsOk returns a tuple with the StraightBets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBetsByTypeResponseV3) GetStraightBetsOk() (*[]StraightBetV3, bool) {
	if o == nil || o.StraightBets == nil {
		return nil, false
	}
	return o.StraightBets, true
}

// HasStraightBets returns a boolean if a field has been set.
func (o *GetBetsByTypeResponseV3) HasStraightBets() bool {
	if o != nil && o.StraightBets != nil {
		return true
	}

	return false
}

// SetStraightBets gets a reference to the given []StraightBetV3 and assigns it to the StraightBets field.
func (o *GetBetsByTypeResponseV3) SetStraightBets(v []StraightBetV3) {
	o.StraightBets = &v
}

// GetParlayBets returns the ParlayBets field value if set, zero value otherwise.
func (o *GetBetsByTypeResponseV3) GetParlayBets() []ParlayBet {
	if o == nil || o.ParlayBets == nil {
		var ret []ParlayBet
		return ret
	}
	return *o.ParlayBets
}

// GetParlayBetsOk returns a tuple with the ParlayBets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBetsByTypeResponseV3) GetParlayBetsOk() (*[]ParlayBet, bool) {
	if o == nil || o.ParlayBets == nil {
		return nil, false
	}
	return o.ParlayBets, true
}

// HasParlayBets returns a boolean if a field has been set.
func (o *GetBetsByTypeResponseV3) HasParlayBets() bool {
	if o != nil && o.ParlayBets != nil {
		return true
	}

	return false
}

// SetParlayBets gets a reference to the given []ParlayBet and assigns it to the ParlayBets field.
func (o *GetBetsByTypeResponseV3) SetParlayBets(v []ParlayBet) {
	o.ParlayBets = &v
}

// GetTeaserBets returns the TeaserBets field value if set, zero value otherwise.
func (o *GetBetsByTypeResponseV3) GetTeaserBets() []TeaserBet {
	if o == nil || o.TeaserBets == nil {
		var ret []TeaserBet
		return ret
	}
	return *o.TeaserBets
}

// GetTeaserBetsOk returns a tuple with the TeaserBets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBetsByTypeResponseV3) GetTeaserBetsOk() (*[]TeaserBet, bool) {
	if o == nil || o.TeaserBets == nil {
		return nil, false
	}
	return o.TeaserBets, true
}

// HasTeaserBets returns a boolean if a field has been set.
func (o *GetBetsByTypeResponseV3) HasTeaserBets() bool {
	if o != nil && o.TeaserBets != nil {
		return true
	}

	return false
}

// SetTeaserBets gets a reference to the given []TeaserBet and assigns it to the TeaserBets field.
func (o *GetBetsByTypeResponseV3) SetTeaserBets(v []TeaserBet) {
	o.TeaserBets = &v
}

// GetSpecialBets returns the SpecialBets field value if set, zero value otherwise.
func (o *GetBetsByTypeResponseV3) GetSpecialBets() []SpecialBet {
	if o == nil || o.SpecialBets == nil {
		var ret []SpecialBet
		return ret
	}
	return *o.SpecialBets
}

// GetSpecialBetsOk returns a tuple with the SpecialBets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBetsByTypeResponseV3) GetSpecialBetsOk() (*[]SpecialBet, bool) {
	if o == nil || o.SpecialBets == nil {
		return nil, false
	}
	return o.SpecialBets, true
}

// HasSpecialBets returns a boolean if a field has been set.
func (o *GetBetsByTypeResponseV3) HasSpecialBets() bool {
	if o != nil && o.SpecialBets != nil {
		return true
	}

	return false
}

// SetSpecialBets gets a reference to the given []SpecialBet and assigns it to the SpecialBets field.
func (o *GetBetsByTypeResponseV3) SetSpecialBets(v []SpecialBet) {
	o.SpecialBets = &v
}

// GetManualBets returns the ManualBets field value if set, zero value otherwise.
func (o *GetBetsByTypeResponseV3) GetManualBets() []ManualBet {
	if o == nil || o.ManualBets == nil {
		var ret []ManualBet
		return ret
	}
	return *o.ManualBets
}

// GetManualBetsOk returns a tuple with the ManualBets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBetsByTypeResponseV3) GetManualBetsOk() (*[]ManualBet, bool) {
	if o == nil || o.ManualBets == nil {
		return nil, false
	}
	return o.ManualBets, true
}

// HasManualBets returns a boolean if a field has been set.
func (o *GetBetsByTypeResponseV3) HasManualBets() bool {
	if o != nil && o.ManualBets != nil {
		return true
	}

	return false
}

// SetManualBets gets a reference to the given []ManualBet and assigns it to the ManualBets field.
func (o *GetBetsByTypeResponseV3) SetManualBets(v []ManualBet) {
	o.ManualBets = &v
}

func (o GetBetsByTypeResponseV3) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MoreAvailable != nil {
		toSerialize["moreAvailable"] = o.MoreAvailable
	}
	if o.PageSize != nil {
		toSerialize["pageSize"] = o.PageSize
	}
	if o.FromRecord != nil {
		toSerialize["fromRecord"] = o.FromRecord
	}
	if o.ToRecord != nil {
		toSerialize["toRecord"] = o.ToRecord
	}
	if o.StraightBets != nil {
		toSerialize["straightBets"] = o.StraightBets
	}
	if o.ParlayBets != nil {
		toSerialize["parlayBets"] = o.ParlayBets
	}
	if o.TeaserBets != nil {
		toSerialize["teaserBets"] = o.TeaserBets
	}
	if o.SpecialBets != nil {
		toSerialize["specialBets"] = o.SpecialBets
	}
	if o.ManualBets != nil {
		toSerialize["manualBets"] = o.ManualBets
	}
	return json.Marshal(toSerialize)
}

type NullableGetBetsByTypeResponseV3 struct {
	value *GetBetsByTypeResponseV3
	isSet bool
}

func (v NullableGetBetsByTypeResponseV3) Get() *GetBetsByTypeResponseV3 {
	return v.value
}

func (v *NullableGetBetsByTypeResponseV3) Set(val *GetBetsByTypeResponseV3) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBetsByTypeResponseV3) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBetsByTypeResponseV3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBetsByTypeResponseV3(val *GetBetsByTypeResponseV3) *NullableGetBetsByTypeResponseV3 {
	return &NullableGetBetsByTypeResponseV3{value: val, isSet: true}
}

func (v NullableGetBetsByTypeResponseV3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBetsByTypeResponseV3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
