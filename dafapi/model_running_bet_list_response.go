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

// RunningBetListResponse struct for RunningBetListResponse
type RunningBetListResponse struct {
	StakeCount string `json:"StakeCount"`
	TicketCount int64 `json:"TicketCount"`
	Ticket string `json:"Ticket"`
}

// NewRunningBetListResponse instantiates a new RunningBetListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRunningBetListResponse(stakeCount string, ticketCount int64, ticket string, ) *RunningBetListResponse {
	this := RunningBetListResponse{}
	this.StakeCount = stakeCount
	this.TicketCount = ticketCount
	this.Ticket = ticket
	return &this
}

// NewRunningBetListResponseWithDefaults instantiates a new RunningBetListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRunningBetListResponseWithDefaults() *RunningBetListResponse {
	this := RunningBetListResponse{}
	return &this
}

// GetStakeCount returns the StakeCount field value
func (o *RunningBetListResponse) GetStakeCount() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.StakeCount
}

// GetStakeCountOk returns a tuple with the StakeCount field value
// and a boolean to check if the value has been set.
func (o *RunningBetListResponse) GetStakeCountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.StakeCount, true
}

// SetStakeCount sets field value
func (o *RunningBetListResponse) SetStakeCount(v string) {
	o.StakeCount = v
}

// GetTicketCount returns the TicketCount field value
func (o *RunningBetListResponse) GetTicketCount() int64 {
	if o == nil  {
		var ret int64
		return ret
	}

	return o.TicketCount
}

// GetTicketCountOk returns a tuple with the TicketCount field value
// and a boolean to check if the value has been set.
func (o *RunningBetListResponse) GetTicketCountOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TicketCount, true
}

// SetTicketCount sets field value
func (o *RunningBetListResponse) SetTicketCount(v int64) {
	o.TicketCount = v
}

// GetTicket returns the Ticket field value
func (o *RunningBetListResponse) GetTicket() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Ticket
}

// GetTicketOk returns a tuple with the Ticket field value
// and a boolean to check if the value has been set.
func (o *RunningBetListResponse) GetTicketOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Ticket, true
}

// SetTicket sets field value
func (o *RunningBetListResponse) SetTicket(v string) {
	o.Ticket = v
}

func (o RunningBetListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["StakeCount"] = o.StakeCount
	}
	if true {
		toSerialize["TicketCount"] = o.TicketCount
	}
	if true {
		toSerialize["Ticket"] = o.Ticket
	}
	return json.Marshal(toSerialize)
}

type NullableRunningBetListResponse struct {
	value *RunningBetListResponse
	isSet bool
}

func (v NullableRunningBetListResponse) Get() *RunningBetListResponse {
	return v.value
}

func (v *NullableRunningBetListResponse) Set(val *RunningBetListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRunningBetListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRunningBetListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRunningBetListResponse(val *RunningBetListResponse) *NullableRunningBetListResponse {
	return &NullableRunningBetListResponse{value: val, isSet: true}
}

func (v NullableRunningBetListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRunningBetListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}