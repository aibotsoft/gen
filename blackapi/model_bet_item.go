/*
 * Black API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package blackapi

import (
	"encoding/json"
)

// BetItem struct for BetItem
type BetItem struct {
	BetId *int64 `json:"bet_id,omitempty"`
	OrderId *int64 `json:"order_id,omitempty"`
	Sport *string `json:"sport,omitempty"`
	EventId *string `json:"event_id,omitempty"`
	Bookie *string `json:"bookie,omitempty"`
	Username *string `json:"username,omitempty"`
	BetType *string `json:"bet_type,omitempty"`
	CcyRate *float64 `json:"ccy_rate,omitempty"`
	Reconciled *bool `json:"reconciled,omitempty"`
	BookieBetId *string `json:"bookie_bet_id,omitempty"`
	Status *BetItemStatus `json:"status,omitempty"`
	WantPrice *float64 `json:"want_price,omitempty"`
	GotPrice *float64 `json:"got_price,omitempty"`
	WantStake *[]interface{} `json:"want_stake,omitempty"`
	GotStake *[]interface{} `json:"got_stake,omitempty"`
	ProfitLoss *[]interface{} `json:"profit_loss,omitempty"`
}

// NewBetItem instantiates a new BetItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBetItem() *BetItem {
	this := BetItem{}
	return &this
}

// NewBetItemWithDefaults instantiates a new BetItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBetItemWithDefaults() *BetItem {
	this := BetItem{}
	return &this
}

// GetBetId returns the BetId field value if set, zero value otherwise.
func (o *BetItem) GetBetId() int64 {
	if o == nil || o.BetId == nil {
		var ret int64
		return ret
	}
	return *o.BetId
}

// GetBetIdOk returns a tuple with the BetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetItem) GetBetIdOk() (*int64, bool) {
	if o == nil || o.BetId == nil {
		return nil, false
	}
	return o.BetId, true
}

// HasBetId returns a boolean if a field has been set.
func (o *BetItem) HasBetId() bool {
	if o != nil && o.BetId != nil {
		return true
	}

	return false
}

// SetBetId gets a reference to the given int64 and assigns it to the BetId field.
func (o *BetItem) SetBetId(v int64) {
	o.BetId = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *BetItem) GetOrderId() int64 {
	if o == nil || o.OrderId == nil {
		var ret int64
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetItem) GetOrderIdOk() (*int64, bool) {
	if o == nil || o.OrderId == nil {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *BetItem) HasOrderId() bool {
	if o != nil && o.OrderId != nil {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given int64 and assigns it to the OrderId field.
func (o *BetItem) SetOrderId(v int64) {
	o.OrderId = &v
}

// GetSport returns the Sport field value if set, zero value otherwise.
func (o *BetItem) GetSport() string {
	if o == nil || o.Sport == nil {
		var ret string
		return ret
	}
	return *o.Sport
}

// GetSportOk returns a tuple with the Sport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetItem) GetSportOk() (*string, bool) {
	if o == nil || o.Sport == nil {
		return nil, false
	}
	return o.Sport, true
}

// HasSport returns a boolean if a field has been set.
func (o *BetItem) HasSport() bool {
	if o != nil && o.Sport != nil {
		return true
	}

	return false
}

// SetSport gets a reference to the given string and assigns it to the Sport field.
func (o *BetItem) SetSport(v string) {
	o.Sport = &v
}

// GetEventId returns the EventId field value if set, zero value otherwise.
func (o *BetItem) GetEventId() string {
	if o == nil || o.EventId == nil {
		var ret string
		return ret
	}
	return *o.EventId
}

// GetEventIdOk returns a tuple with the EventId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetItem) GetEventIdOk() (*string, bool) {
	if o == nil || o.EventId == nil {
		return nil, false
	}
	return o.EventId, true
}

// HasEventId returns a boolean if a field has been set.
func (o *BetItem) HasEventId() bool {
	if o != nil && o.EventId != nil {
		return true
	}

	return false
}

// SetEventId gets a reference to the given string and assigns it to the EventId field.
func (o *BetItem) SetEventId(v string) {
	o.EventId = &v
}

// GetBookie returns the Bookie field value if set, zero value otherwise.
func (o *BetItem) GetBookie() string {
	if o == nil || o.Bookie == nil {
		var ret string
		return ret
	}
	return *o.Bookie
}

// GetBookieOk returns a tuple with the Bookie field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetItem) GetBookieOk() (*string, bool) {
	if o == nil || o.Bookie == nil {
		return nil, false
	}
	return o.Bookie, true
}

// HasBookie returns a boolean if a field has been set.
func (o *BetItem) HasBookie() bool {
	if o != nil && o.Bookie != nil {
		return true
	}

	return false
}

// SetBookie gets a reference to the given string and assigns it to the Bookie field.
func (o *BetItem) SetBookie(v string) {
	o.Bookie = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *BetItem) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetItem) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *BetItem) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *BetItem) SetUsername(v string) {
	o.Username = &v
}

// GetBetType returns the BetType field value if set, zero value otherwise.
func (o *BetItem) GetBetType() string {
	if o == nil || o.BetType == nil {
		var ret string
		return ret
	}
	return *o.BetType
}

// GetBetTypeOk returns a tuple with the BetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetItem) GetBetTypeOk() (*string, bool) {
	if o == nil || o.BetType == nil {
		return nil, false
	}
	return o.BetType, true
}

// HasBetType returns a boolean if a field has been set.
func (o *BetItem) HasBetType() bool {
	if o != nil && o.BetType != nil {
		return true
	}

	return false
}

// SetBetType gets a reference to the given string and assigns it to the BetType field.
func (o *BetItem) SetBetType(v string) {
	o.BetType = &v
}

// GetCcyRate returns the CcyRate field value if set, zero value otherwise.
func (o *BetItem) GetCcyRate() float64 {
	if o == nil || o.CcyRate == nil {
		var ret float64
		return ret
	}
	return *o.CcyRate
}

// GetCcyRateOk returns a tuple with the CcyRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetItem) GetCcyRateOk() (*float64, bool) {
	if o == nil || o.CcyRate == nil {
		return nil, false
	}
	return o.CcyRate, true
}

// HasCcyRate returns a boolean if a field has been set.
func (o *BetItem) HasCcyRate() bool {
	if o != nil && o.CcyRate != nil {
		return true
	}

	return false
}

// SetCcyRate gets a reference to the given float64 and assigns it to the CcyRate field.
func (o *BetItem) SetCcyRate(v float64) {
	o.CcyRate = &v
}

// GetReconciled returns the Reconciled field value if set, zero value otherwise.
func (o *BetItem) GetReconciled() bool {
	if o == nil || o.Reconciled == nil {
		var ret bool
		return ret
	}
	return *o.Reconciled
}

// GetReconciledOk returns a tuple with the Reconciled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetItem) GetReconciledOk() (*bool, bool) {
	if o == nil || o.Reconciled == nil {
		return nil, false
	}
	return o.Reconciled, true
}

// HasReconciled returns a boolean if a field has been set.
func (o *BetItem) HasReconciled() bool {
	if o != nil && o.Reconciled != nil {
		return true
	}

	return false
}

// SetReconciled gets a reference to the given bool and assigns it to the Reconciled field.
func (o *BetItem) SetReconciled(v bool) {
	o.Reconciled = &v
}

// GetBookieBetId returns the BookieBetId field value if set, zero value otherwise.
func (o *BetItem) GetBookieBetId() string {
	if o == nil || o.BookieBetId == nil {
		var ret string
		return ret
	}
	return *o.BookieBetId
}

// GetBookieBetIdOk returns a tuple with the BookieBetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetItem) GetBookieBetIdOk() (*string, bool) {
	if o == nil || o.BookieBetId == nil {
		return nil, false
	}
	return o.BookieBetId, true
}

// HasBookieBetId returns a boolean if a field has been set.
func (o *BetItem) HasBookieBetId() bool {
	if o != nil && o.BookieBetId != nil {
		return true
	}

	return false
}

// SetBookieBetId gets a reference to the given string and assigns it to the BookieBetId field.
func (o *BetItem) SetBookieBetId(v string) {
	o.BookieBetId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *BetItem) GetStatus() BetItemStatus {
	if o == nil || o.Status == nil {
		var ret BetItemStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetItem) GetStatusOk() (*BetItemStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *BetItem) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given BetItemStatus and assigns it to the Status field.
func (o *BetItem) SetStatus(v BetItemStatus) {
	o.Status = &v
}

// GetWantPrice returns the WantPrice field value if set, zero value otherwise.
func (o *BetItem) GetWantPrice() float64 {
	if o == nil || o.WantPrice == nil {
		var ret float64
		return ret
	}
	return *o.WantPrice
}

// GetWantPriceOk returns a tuple with the WantPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetItem) GetWantPriceOk() (*float64, bool) {
	if o == nil || o.WantPrice == nil {
		return nil, false
	}
	return o.WantPrice, true
}

// HasWantPrice returns a boolean if a field has been set.
func (o *BetItem) HasWantPrice() bool {
	if o != nil && o.WantPrice != nil {
		return true
	}

	return false
}

// SetWantPrice gets a reference to the given float64 and assigns it to the WantPrice field.
func (o *BetItem) SetWantPrice(v float64) {
	o.WantPrice = &v
}

// GetGotPrice returns the GotPrice field value if set, zero value otherwise.
func (o *BetItem) GetGotPrice() float64 {
	if o == nil || o.GotPrice == nil {
		var ret float64
		return ret
	}
	return *o.GotPrice
}

// GetGotPriceOk returns a tuple with the GotPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetItem) GetGotPriceOk() (*float64, bool) {
	if o == nil || o.GotPrice == nil {
		return nil, false
	}
	return o.GotPrice, true
}

// HasGotPrice returns a boolean if a field has been set.
func (o *BetItem) HasGotPrice() bool {
	if o != nil && o.GotPrice != nil {
		return true
	}

	return false
}

// SetGotPrice gets a reference to the given float64 and assigns it to the GotPrice field.
func (o *BetItem) SetGotPrice(v float64) {
	o.GotPrice = &v
}

// GetWantStake returns the WantStake field value if set, zero value otherwise.
func (o *BetItem) GetWantStake() []interface{} {
	if o == nil || o.WantStake == nil {
		var ret []interface{}
		return ret
	}
	return *o.WantStake
}

// GetWantStakeOk returns a tuple with the WantStake field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetItem) GetWantStakeOk() (*[]interface{}, bool) {
	if o == nil || o.WantStake == nil {
		return nil, false
	}
	return o.WantStake, true
}

// HasWantStake returns a boolean if a field has been set.
func (o *BetItem) HasWantStake() bool {
	if o != nil && o.WantStake != nil {
		return true
	}

	return false
}

// SetWantStake gets a reference to the given []interface{} and assigns it to the WantStake field.
func (o *BetItem) SetWantStake(v []interface{}) {
	o.WantStake = &v
}

// GetGotStake returns the GotStake field value if set, zero value otherwise.
func (o *BetItem) GetGotStake() []interface{} {
	if o == nil || o.GotStake == nil {
		var ret []interface{}
		return ret
	}
	return *o.GotStake
}

// GetGotStakeOk returns a tuple with the GotStake field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetItem) GetGotStakeOk() (*[]interface{}, bool) {
	if o == nil || o.GotStake == nil {
		return nil, false
	}
	return o.GotStake, true
}

// HasGotStake returns a boolean if a field has been set.
func (o *BetItem) HasGotStake() bool {
	if o != nil && o.GotStake != nil {
		return true
	}

	return false
}

// SetGotStake gets a reference to the given []interface{} and assigns it to the GotStake field.
func (o *BetItem) SetGotStake(v []interface{}) {
	o.GotStake = &v
}

// GetProfitLoss returns the ProfitLoss field value if set, zero value otherwise.
func (o *BetItem) GetProfitLoss() []interface{} {
	if o == nil || o.ProfitLoss == nil {
		var ret []interface{}
		return ret
	}
	return *o.ProfitLoss
}

// GetProfitLossOk returns a tuple with the ProfitLoss field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetItem) GetProfitLossOk() (*[]interface{}, bool) {
	if o == nil || o.ProfitLoss == nil {
		return nil, false
	}
	return o.ProfitLoss, true
}

// HasProfitLoss returns a boolean if a field has been set.
func (o *BetItem) HasProfitLoss() bool {
	if o != nil && o.ProfitLoss != nil {
		return true
	}

	return false
}

// SetProfitLoss gets a reference to the given []interface{} and assigns it to the ProfitLoss field.
func (o *BetItem) SetProfitLoss(v []interface{}) {
	o.ProfitLoss = &v
}

func (o BetItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BetId != nil {
		toSerialize["bet_id"] = o.BetId
	}
	if o.OrderId != nil {
		toSerialize["order_id"] = o.OrderId
	}
	if o.Sport != nil {
		toSerialize["sport"] = o.Sport
	}
	if o.EventId != nil {
		toSerialize["event_id"] = o.EventId
	}
	if o.Bookie != nil {
		toSerialize["bookie"] = o.Bookie
	}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	if o.BetType != nil {
		toSerialize["bet_type"] = o.BetType
	}
	if o.CcyRate != nil {
		toSerialize["ccy_rate"] = o.CcyRate
	}
	if o.Reconciled != nil {
		toSerialize["reconciled"] = o.Reconciled
	}
	if o.BookieBetId != nil {
		toSerialize["bookie_bet_id"] = o.BookieBetId
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.WantPrice != nil {
		toSerialize["want_price"] = o.WantPrice
	}
	if o.GotPrice != nil {
		toSerialize["got_price"] = o.GotPrice
	}
	if o.WantStake != nil {
		toSerialize["want_stake"] = o.WantStake
	}
	if o.GotStake != nil {
		toSerialize["got_stake"] = o.GotStake
	}
	if o.ProfitLoss != nil {
		toSerialize["profit_loss"] = o.ProfitLoss
	}
	return json.Marshal(toSerialize)
}

type NullableBetItem struct {
	value *BetItem
	isSet bool
}

func (v NullableBetItem) Get() *BetItem {
	return v.value
}

func (v *NullableBetItem) Set(val *BetItem) {
	v.value = val
	v.isSet = true
}

func (v NullableBetItem) IsSet() bool {
	return v.isSet
}

func (v *NullableBetItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetItem(val *BetItem) *NullableBetItem {
	return &NullableBetItem{value: val, isSet: true}
}

func (v NullableBetItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
