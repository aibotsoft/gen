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

// Period struct for Period
type Period struct {
	// Period Number
	Number *int `json:"number,omitempty"`
	// Description for the period
	Description *string `json:"description,omitempty"`
	// Short description for the period
	ShortDescription *string `json:"shortDescription,omitempty"`
	// Description for the Spread
	SpreadDescription *string `json:"spreadDescription,omitempty"`
	// Description for the Moneyline
	MoneylineDescription *string `json:"moneylineDescription,omitempty"`
	// Description for the Totals
	TotalDescription *string `json:"totalDescription,omitempty"`
	// Description for Team1 Totals
	Team1TotalDescription *string `json:"team1TotalDescription,omitempty"`
	// Description for Team2 Totals
	Team2TotalDescription *string `json:"team2TotalDescription,omitempty"`
	// Short description for the Spread
	SpreadShortDescription *string `json:"spreadShortDescription,omitempty"`
	// Short description for the Moneyline
	MoneylineShortDescription *string `json:"moneylineShortDescription,omitempty"`
	// Short description for the Totals
	TotalShortDescription *string `json:"totalShortDescription,omitempty"`
	// Short description for Team1 Totals
	Team1TotalShortDescription *string `json:"team1TotalShortDescription,omitempty"`
	// Short description for Team2 Totals
	Team2TotalShortDescription *string `json:"team2TotalShortDescription,omitempty"`
}

// NewPeriod instantiates a new Period object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPeriod() *Period {
	this := Period{}
	return &this
}

// NewPeriodWithDefaults instantiates a new Period object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPeriodWithDefaults() *Period {
	this := Period{}
	return &this
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *Period) GetNumber() int {
	if o == nil || o.Number == nil {
		var ret int
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Period) GetNumberOk() (*int, bool) {
	if o == nil || o.Number == nil {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *Period) HasNumber() bool {
	if o != nil && o.Number != nil {
		return true
	}

	return false
}

// SetNumber gets a reference to the given int and assigns it to the Number field.
func (o *Period) SetNumber(v int) {
	o.Number = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Period) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Period) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Period) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Period) SetDescription(v string) {
	o.Description = &v
}

// GetShortDescription returns the ShortDescription field value if set, zero value otherwise.
func (o *Period) GetShortDescription() string {
	if o == nil || o.ShortDescription == nil {
		var ret string
		return ret
	}
	return *o.ShortDescription
}

// GetShortDescriptionOk returns a tuple with the ShortDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Period) GetShortDescriptionOk() (*string, bool) {
	if o == nil || o.ShortDescription == nil {
		return nil, false
	}
	return o.ShortDescription, true
}

// HasShortDescription returns a boolean if a field has been set.
func (o *Period) HasShortDescription() bool {
	if o != nil && o.ShortDescription != nil {
		return true
	}

	return false
}

// SetShortDescription gets a reference to the given string and assigns it to the ShortDescription field.
func (o *Period) SetShortDescription(v string) {
	o.ShortDescription = &v
}

// GetSpreadDescription returns the SpreadDescription field value if set, zero value otherwise.
func (o *Period) GetSpreadDescription() string {
	if o == nil || o.SpreadDescription == nil {
		var ret string
		return ret
	}
	return *o.SpreadDescription
}

// GetSpreadDescriptionOk returns a tuple with the SpreadDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Period) GetSpreadDescriptionOk() (*string, bool) {
	if o == nil || o.SpreadDescription == nil {
		return nil, false
	}
	return o.SpreadDescription, true
}

// HasSpreadDescription returns a boolean if a field has been set.
func (o *Period) HasSpreadDescription() bool {
	if o != nil && o.SpreadDescription != nil {
		return true
	}

	return false
}

// SetSpreadDescription gets a reference to the given string and assigns it to the SpreadDescription field.
func (o *Period) SetSpreadDescription(v string) {
	o.SpreadDescription = &v
}

// GetMoneylineDescription returns the MoneylineDescription field value if set, zero value otherwise.
func (o *Period) GetMoneylineDescription() string {
	if o == nil || o.MoneylineDescription == nil {
		var ret string
		return ret
	}
	return *o.MoneylineDescription
}

// GetMoneylineDescriptionOk returns a tuple with the MoneylineDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Period) GetMoneylineDescriptionOk() (*string, bool) {
	if o == nil || o.MoneylineDescription == nil {
		return nil, false
	}
	return o.MoneylineDescription, true
}

// HasMoneylineDescription returns a boolean if a field has been set.
func (o *Period) HasMoneylineDescription() bool {
	if o != nil && o.MoneylineDescription != nil {
		return true
	}

	return false
}

// SetMoneylineDescription gets a reference to the given string and assigns it to the MoneylineDescription field.
func (o *Period) SetMoneylineDescription(v string) {
	o.MoneylineDescription = &v
}

// GetTotalDescription returns the TotalDescription field value if set, zero value otherwise.
func (o *Period) GetTotalDescription() string {
	if o == nil || o.TotalDescription == nil {
		var ret string
		return ret
	}
	return *o.TotalDescription
}

// GetTotalDescriptionOk returns a tuple with the TotalDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Period) GetTotalDescriptionOk() (*string, bool) {
	if o == nil || o.TotalDescription == nil {
		return nil, false
	}
	return o.TotalDescription, true
}

// HasTotalDescription returns a boolean if a field has been set.
func (o *Period) HasTotalDescription() bool {
	if o != nil && o.TotalDescription != nil {
		return true
	}

	return false
}

// SetTotalDescription gets a reference to the given string and assigns it to the TotalDescription field.
func (o *Period) SetTotalDescription(v string) {
	o.TotalDescription = &v
}

// GetTeam1TotalDescription returns the Team1TotalDescription field value if set, zero value otherwise.
func (o *Period) GetTeam1TotalDescription() string {
	if o == nil || o.Team1TotalDescription == nil {
		var ret string
		return ret
	}
	return *o.Team1TotalDescription
}

// GetTeam1TotalDescriptionOk returns a tuple with the Team1TotalDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Period) GetTeam1TotalDescriptionOk() (*string, bool) {
	if o == nil || o.Team1TotalDescription == nil {
		return nil, false
	}
	return o.Team1TotalDescription, true
}

// HasTeam1TotalDescription returns a boolean if a field has been set.
func (o *Period) HasTeam1TotalDescription() bool {
	if o != nil && o.Team1TotalDescription != nil {
		return true
	}

	return false
}

// SetTeam1TotalDescription gets a reference to the given string and assigns it to the Team1TotalDescription field.
func (o *Period) SetTeam1TotalDescription(v string) {
	o.Team1TotalDescription = &v
}

// GetTeam2TotalDescription returns the Team2TotalDescription field value if set, zero value otherwise.
func (o *Period) GetTeam2TotalDescription() string {
	if o == nil || o.Team2TotalDescription == nil {
		var ret string
		return ret
	}
	return *o.Team2TotalDescription
}

// GetTeam2TotalDescriptionOk returns a tuple with the Team2TotalDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Period) GetTeam2TotalDescriptionOk() (*string, bool) {
	if o == nil || o.Team2TotalDescription == nil {
		return nil, false
	}
	return o.Team2TotalDescription, true
}

// HasTeam2TotalDescription returns a boolean if a field has been set.
func (o *Period) HasTeam2TotalDescription() bool {
	if o != nil && o.Team2TotalDescription != nil {
		return true
	}

	return false
}

// SetTeam2TotalDescription gets a reference to the given string and assigns it to the Team2TotalDescription field.
func (o *Period) SetTeam2TotalDescription(v string) {
	o.Team2TotalDescription = &v
}

// GetSpreadShortDescription returns the SpreadShortDescription field value if set, zero value otherwise.
func (o *Period) GetSpreadShortDescription() string {
	if o == nil || o.SpreadShortDescription == nil {
		var ret string
		return ret
	}
	return *o.SpreadShortDescription
}

// GetSpreadShortDescriptionOk returns a tuple with the SpreadShortDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Period) GetSpreadShortDescriptionOk() (*string, bool) {
	if o == nil || o.SpreadShortDescription == nil {
		return nil, false
	}
	return o.SpreadShortDescription, true
}

// HasSpreadShortDescription returns a boolean if a field has been set.
func (o *Period) HasSpreadShortDescription() bool {
	if o != nil && o.SpreadShortDescription != nil {
		return true
	}

	return false
}

// SetSpreadShortDescription gets a reference to the given string and assigns it to the SpreadShortDescription field.
func (o *Period) SetSpreadShortDescription(v string) {
	o.SpreadShortDescription = &v
}

// GetMoneylineShortDescription returns the MoneylineShortDescription field value if set, zero value otherwise.
func (o *Period) GetMoneylineShortDescription() string {
	if o == nil || o.MoneylineShortDescription == nil {
		var ret string
		return ret
	}
	return *o.MoneylineShortDescription
}

// GetMoneylineShortDescriptionOk returns a tuple with the MoneylineShortDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Period) GetMoneylineShortDescriptionOk() (*string, bool) {
	if o == nil || o.MoneylineShortDescription == nil {
		return nil, false
	}
	return o.MoneylineShortDescription, true
}

// HasMoneylineShortDescription returns a boolean if a field has been set.
func (o *Period) HasMoneylineShortDescription() bool {
	if o != nil && o.MoneylineShortDescription != nil {
		return true
	}

	return false
}

// SetMoneylineShortDescription gets a reference to the given string and assigns it to the MoneylineShortDescription field.
func (o *Period) SetMoneylineShortDescription(v string) {
	o.MoneylineShortDescription = &v
}

// GetTotalShortDescription returns the TotalShortDescription field value if set, zero value otherwise.
func (o *Period) GetTotalShortDescription() string {
	if o == nil || o.TotalShortDescription == nil {
		var ret string
		return ret
	}
	return *o.TotalShortDescription
}

// GetTotalShortDescriptionOk returns a tuple with the TotalShortDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Period) GetTotalShortDescriptionOk() (*string, bool) {
	if o == nil || o.TotalShortDescription == nil {
		return nil, false
	}
	return o.TotalShortDescription, true
}

// HasTotalShortDescription returns a boolean if a field has been set.
func (o *Period) HasTotalShortDescription() bool {
	if o != nil && o.TotalShortDescription != nil {
		return true
	}

	return false
}

// SetTotalShortDescription gets a reference to the given string and assigns it to the TotalShortDescription field.
func (o *Period) SetTotalShortDescription(v string) {
	o.TotalShortDescription = &v
}

// GetTeam1TotalShortDescription returns the Team1TotalShortDescription field value if set, zero value otherwise.
func (o *Period) GetTeam1TotalShortDescription() string {
	if o == nil || o.Team1TotalShortDescription == nil {
		var ret string
		return ret
	}
	return *o.Team1TotalShortDescription
}

// GetTeam1TotalShortDescriptionOk returns a tuple with the Team1TotalShortDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Period) GetTeam1TotalShortDescriptionOk() (*string, bool) {
	if o == nil || o.Team1TotalShortDescription == nil {
		return nil, false
	}
	return o.Team1TotalShortDescription, true
}

// HasTeam1TotalShortDescription returns a boolean if a field has been set.
func (o *Period) HasTeam1TotalShortDescription() bool {
	if o != nil && o.Team1TotalShortDescription != nil {
		return true
	}

	return false
}

// SetTeam1TotalShortDescription gets a reference to the given string and assigns it to the Team1TotalShortDescription field.
func (o *Period) SetTeam1TotalShortDescription(v string) {
	o.Team1TotalShortDescription = &v
}

// GetTeam2TotalShortDescription returns the Team2TotalShortDescription field value if set, zero value otherwise.
func (o *Period) GetTeam2TotalShortDescription() string {
	if o == nil || o.Team2TotalShortDescription == nil {
		var ret string
		return ret
	}
	return *o.Team2TotalShortDescription
}

// GetTeam2TotalShortDescriptionOk returns a tuple with the Team2TotalShortDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Period) GetTeam2TotalShortDescriptionOk() (*string, bool) {
	if o == nil || o.Team2TotalShortDescription == nil {
		return nil, false
	}
	return o.Team2TotalShortDescription, true
}

// HasTeam2TotalShortDescription returns a boolean if a field has been set.
func (o *Period) HasTeam2TotalShortDescription() bool {
	if o != nil && o.Team2TotalShortDescription != nil {
		return true
	}

	return false
}

// SetTeam2TotalShortDescription gets a reference to the given string and assigns it to the Team2TotalShortDescription field.
func (o *Period) SetTeam2TotalShortDescription(v string) {
	o.Team2TotalShortDescription = &v
}

func (o Period) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Number != nil {
		toSerialize["number"] = o.Number
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.ShortDescription != nil {
		toSerialize["shortDescription"] = o.ShortDescription
	}
	if o.SpreadDescription != nil {
		toSerialize["spreadDescription"] = o.SpreadDescription
	}
	if o.MoneylineDescription != nil {
		toSerialize["moneylineDescription"] = o.MoneylineDescription
	}
	if o.TotalDescription != nil {
		toSerialize["totalDescription"] = o.TotalDescription
	}
	if o.Team1TotalDescription != nil {
		toSerialize["team1TotalDescription"] = o.Team1TotalDescription
	}
	if o.Team2TotalDescription != nil {
		toSerialize["team2TotalDescription"] = o.Team2TotalDescription
	}
	if o.SpreadShortDescription != nil {
		toSerialize["spreadShortDescription"] = o.SpreadShortDescription
	}
	if o.MoneylineShortDescription != nil {
		toSerialize["moneylineShortDescription"] = o.MoneylineShortDescription
	}
	if o.TotalShortDescription != nil {
		toSerialize["totalShortDescription"] = o.TotalShortDescription
	}
	if o.Team1TotalShortDescription != nil {
		toSerialize["team1TotalShortDescription"] = o.Team1TotalShortDescription
	}
	if o.Team2TotalShortDescription != nil {
		toSerialize["team2TotalShortDescription"] = o.Team2TotalShortDescription
	}
	return json.Marshal(toSerialize)
}

type NullablePeriod struct {
	value *Period
	isSet bool
}

func (v NullablePeriod) Get() *Period {
	return v.value
}

func (v *NullablePeriod) Set(val *Period) {
	v.value = val
	v.isSet = true
}

func (v NullablePeriod) IsSet() bool {
	return v.isSet
}

func (v *NullablePeriod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePeriod(val *Period) *NullablePeriod {
	return &NullablePeriod{value: val, isSet: true}
}

func (v NullablePeriod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePeriod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
