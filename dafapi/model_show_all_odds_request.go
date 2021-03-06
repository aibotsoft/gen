/*
 * Sample API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package dafapi

import (
	"encoding/json"
)

// ShowAllOddsRequest struct for ShowAllOddsRequest
type ShowAllOddsRequest struct {
	GameId int64 `json:"GameId"`
	DateType DateTypeEnum `json:"DateType"`
	BetTypeClass BetTypeClassEnum `json:"BetTypeClass"`
	Matchid *int64 `json:"Matchid,omitempty"`
}

// NewShowAllOddsRequest instantiates a new ShowAllOddsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShowAllOddsRequest(gameId int64, dateType DateTypeEnum, betTypeClass BetTypeClassEnum, ) *ShowAllOddsRequest {
	this := ShowAllOddsRequest{}
	this.GameId = gameId
	this.DateType = dateType
	this.BetTypeClass = betTypeClass
	return &this
}

// NewShowAllOddsRequestWithDefaults instantiates a new ShowAllOddsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShowAllOddsRequestWithDefaults() *ShowAllOddsRequest {
	this := ShowAllOddsRequest{}
	return &this
}

// GetGameId returns the GameId field value
func (o *ShowAllOddsRequest) GetGameId() int64 {
	if o == nil  {
		var ret int64
		return ret
	}

	return o.GameId
}

// GetGameIdOk returns a tuple with the GameId field value
// and a boolean to check if the value has been set.
func (o *ShowAllOddsRequest) GetGameIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.GameId, true
}

// SetGameId sets field value
func (o *ShowAllOddsRequest) SetGameId(v int64) {
	o.GameId = v
}

// GetDateType returns the DateType field value
func (o *ShowAllOddsRequest) GetDateType() DateTypeEnum {
	if o == nil  {
		var ret DateTypeEnum
		return ret
	}

	return o.DateType
}

// GetDateTypeOk returns a tuple with the DateType field value
// and a boolean to check if the value has been set.
func (o *ShowAllOddsRequest) GetDateTypeOk() (*DateTypeEnum, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DateType, true
}

// SetDateType sets field value
func (o *ShowAllOddsRequest) SetDateType(v DateTypeEnum) {
	o.DateType = v
}

// GetBetTypeClass returns the BetTypeClass field value
func (o *ShowAllOddsRequest) GetBetTypeClass() BetTypeClassEnum {
	if o == nil  {
		var ret BetTypeClassEnum
		return ret
	}

	return o.BetTypeClass
}

// GetBetTypeClassOk returns a tuple with the BetTypeClass field value
// and a boolean to check if the value has been set.
func (o *ShowAllOddsRequest) GetBetTypeClassOk() (*BetTypeClassEnum, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BetTypeClass, true
}

// SetBetTypeClass sets field value
func (o *ShowAllOddsRequest) SetBetTypeClass(v BetTypeClassEnum) {
	o.BetTypeClass = v
}

// GetMatchid returns the Matchid field value if set, zero value otherwise.
func (o *ShowAllOddsRequest) GetMatchid() int64 {
	if o == nil || o.Matchid == nil {
		var ret int64
		return ret
	}
	return *o.Matchid
}

// GetMatchidOk returns a tuple with the Matchid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShowAllOddsRequest) GetMatchidOk() (*int64, bool) {
	if o == nil || o.Matchid == nil {
		return nil, false
	}
	return o.Matchid, true
}

// HasMatchid returns a boolean if a field has been set.
func (o *ShowAllOddsRequest) HasMatchid() bool {
	if o != nil && o.Matchid != nil {
		return true
	}

	return false
}

// SetMatchid gets a reference to the given int64 and assigns it to the Matchid field.
func (o *ShowAllOddsRequest) SetMatchid(v int64) {
	o.Matchid = &v
}

func (o ShowAllOddsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["GameId"] = o.GameId
	}
	if true {
		toSerialize["DateType"] = o.DateType
	}
	if true {
		toSerialize["BetTypeClass"] = o.BetTypeClass
	}
	if o.Matchid != nil {
		toSerialize["Matchid"] = o.Matchid
	}
	return json.Marshal(toSerialize)
}

type NullableShowAllOddsRequest struct {
	value *ShowAllOddsRequest
	isSet bool
}

func (v NullableShowAllOddsRequest) Get() *ShowAllOddsRequest {
	return v.value
}

func (v *NullableShowAllOddsRequest) Set(val *ShowAllOddsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableShowAllOddsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableShowAllOddsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShowAllOddsRequest(val *ShowAllOddsRequest) *NullableShowAllOddsRequest {
	return &NullableShowAllOddsRequest{value: val, isSet: true}
}

func (v NullableShowAllOddsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShowAllOddsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
