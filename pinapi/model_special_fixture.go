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
	"time"
)

// SpecialFixture struct for SpecialFixture
type SpecialFixture struct {
	// Unique Id
	Id *int64 `json:"id,omitempty"`
	// The type [MULTI_WAY_HEAD_TO_HEAD, SPREAD, OVER_UNDER]
	BetType *string `json:"betType,omitempty"`
	// Name of the special.
	Name *string `json:"name,omitempty"`
	// Date of the special in UTC.
	Date *time.Time `json:"date,omitempty"`
	// Wagering cutoff date in UTC.
	Cutoff *time.Time `json:"cutoff,omitempty"`
	// The category that the special falls under.
	Category *string `json:"category,omitempty"`
	// Measurment in the context of the special. This is applicable to specials bet type spead and over/under. In a hockey special this could be goals.
	Units *string `json:"units,omitempty"`
	// Status of the Special    O = This is the starting status. It means that the lines  are open for betting,    H = This status indicates that the lines are temporarily unavailable  for betting,    I = This status indicates that one or more lines have a red circle  (a lower maximum bet amount) 
	Status *string `json:"status,omitempty"`
	Event *SpecialsFixturesEvent `json:"event,omitempty"`
	// ContestantLines available for wagering.
	Contestants *[]SpecialsFixturesContestant `json:"contestants,omitempty"`
	// When a special is linked to an event, we will return live status of the event, otherwise it will be 0.  0 = No live betting will be offered on this event,  1 = Live betting event,  2 = Live betting will be offered on this match, but on a different event.   Please note that live delay is applied when placing bets on special with LiveStatus=1  
	LiveStatus *int `json:"liveStatus,omitempty"`
}

// NewSpecialFixture instantiates a new SpecialFixture object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpecialFixture() *SpecialFixture {
	this := SpecialFixture{}
	return &this
}

// NewSpecialFixtureWithDefaults instantiates a new SpecialFixture object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpecialFixtureWithDefaults() *SpecialFixture {
	this := SpecialFixture{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SpecialFixture) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpecialFixture) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SpecialFixture) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *SpecialFixture) SetId(v int64) {
	o.Id = &v
}

// GetBetType returns the BetType field value if set, zero value otherwise.
func (o *SpecialFixture) GetBetType() string {
	if o == nil || o.BetType == nil {
		var ret string
		return ret
	}
	return *o.BetType
}

// GetBetTypeOk returns a tuple with the BetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpecialFixture) GetBetTypeOk() (*string, bool) {
	if o == nil || o.BetType == nil {
		return nil, false
	}
	return o.BetType, true
}

// HasBetType returns a boolean if a field has been set.
func (o *SpecialFixture) HasBetType() bool {
	if o != nil && o.BetType != nil {
		return true
	}

	return false
}

// SetBetType gets a reference to the given string and assigns it to the BetType field.
func (o *SpecialFixture) SetBetType(v string) {
	o.BetType = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SpecialFixture) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpecialFixture) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SpecialFixture) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SpecialFixture) SetName(v string) {
	o.Name = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *SpecialFixture) GetDate() time.Time {
	if o == nil || o.Date == nil {
		var ret time.Time
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpecialFixture) GetDateOk() (*time.Time, bool) {
	if o == nil || o.Date == nil {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *SpecialFixture) HasDate() bool {
	if o != nil && o.Date != nil {
		return true
	}

	return false
}

// SetDate gets a reference to the given time.Time and assigns it to the Date field.
func (o *SpecialFixture) SetDate(v time.Time) {
	o.Date = &v
}

// GetCutoff returns the Cutoff field value if set, zero value otherwise.
func (o *SpecialFixture) GetCutoff() time.Time {
	if o == nil || o.Cutoff == nil {
		var ret time.Time
		return ret
	}
	return *o.Cutoff
}

// GetCutoffOk returns a tuple with the Cutoff field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpecialFixture) GetCutoffOk() (*time.Time, bool) {
	if o == nil || o.Cutoff == nil {
		return nil, false
	}
	return o.Cutoff, true
}

// HasCutoff returns a boolean if a field has been set.
func (o *SpecialFixture) HasCutoff() bool {
	if o != nil && o.Cutoff != nil {
		return true
	}

	return false
}

// SetCutoff gets a reference to the given time.Time and assigns it to the Cutoff field.
func (o *SpecialFixture) SetCutoff(v time.Time) {
	o.Cutoff = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *SpecialFixture) GetCategory() string {
	if o == nil || o.Category == nil {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpecialFixture) GetCategoryOk() (*string, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *SpecialFixture) HasCategory() bool {
	if o != nil && o.Category != nil {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *SpecialFixture) SetCategory(v string) {
	o.Category = &v
}

// GetUnits returns the Units field value if set, zero value otherwise.
func (o *SpecialFixture) GetUnits() string {
	if o == nil || o.Units == nil {
		var ret string
		return ret
	}
	return *o.Units
}

// GetUnitsOk returns a tuple with the Units field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpecialFixture) GetUnitsOk() (*string, bool) {
	if o == nil || o.Units == nil {
		return nil, false
	}
	return o.Units, true
}

// HasUnits returns a boolean if a field has been set.
func (o *SpecialFixture) HasUnits() bool {
	if o != nil && o.Units != nil {
		return true
	}

	return false
}

// SetUnits gets a reference to the given string and assigns it to the Units field.
func (o *SpecialFixture) SetUnits(v string) {
	o.Units = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SpecialFixture) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpecialFixture) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SpecialFixture) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *SpecialFixture) SetStatus(v string) {
	o.Status = &v
}

// GetEvent returns the Event field value if set, zero value otherwise.
func (o *SpecialFixture) GetEvent() SpecialsFixturesEvent {
	if o == nil || o.Event == nil {
		var ret SpecialsFixturesEvent
		return ret
	}
	return *o.Event
}

// GetEventOk returns a tuple with the Event field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpecialFixture) GetEventOk() (*SpecialsFixturesEvent, bool) {
	if o == nil || o.Event == nil {
		return nil, false
	}
	return o.Event, true
}

// HasEvent returns a boolean if a field has been set.
func (o *SpecialFixture) HasEvent() bool {
	if o != nil && o.Event != nil {
		return true
	}

	return false
}

// SetEvent gets a reference to the given SpecialsFixturesEvent and assigns it to the Event field.
func (o *SpecialFixture) SetEvent(v SpecialsFixturesEvent) {
	o.Event = &v
}

// GetContestants returns the Contestants field value if set, zero value otherwise.
func (o *SpecialFixture) GetContestants() []SpecialsFixturesContestant {
	if o == nil || o.Contestants == nil {
		var ret []SpecialsFixturesContestant
		return ret
	}
	return *o.Contestants
}

// GetContestantsOk returns a tuple with the Contestants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpecialFixture) GetContestantsOk() (*[]SpecialsFixturesContestant, bool) {
	if o == nil || o.Contestants == nil {
		return nil, false
	}
	return o.Contestants, true
}

// HasContestants returns a boolean if a field has been set.
func (o *SpecialFixture) HasContestants() bool {
	if o != nil && o.Contestants != nil {
		return true
	}

	return false
}

// SetContestants gets a reference to the given []SpecialsFixturesContestant and assigns it to the Contestants field.
func (o *SpecialFixture) SetContestants(v []SpecialsFixturesContestant) {
	o.Contestants = &v
}

// GetLiveStatus returns the LiveStatus field value if set, zero value otherwise.
func (o *SpecialFixture) GetLiveStatus() int {
	if o == nil || o.LiveStatus == nil {
		var ret int
		return ret
	}
	return *o.LiveStatus
}

// GetLiveStatusOk returns a tuple with the LiveStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpecialFixture) GetLiveStatusOk() (*int, bool) {
	if o == nil || o.LiveStatus == nil {
		return nil, false
	}
	return o.LiveStatus, true
}

// HasLiveStatus returns a boolean if a field has been set.
func (o *SpecialFixture) HasLiveStatus() bool {
	if o != nil && o.LiveStatus != nil {
		return true
	}

	return false
}

// SetLiveStatus gets a reference to the given int and assigns it to the LiveStatus field.
func (o *SpecialFixture) SetLiveStatus(v int) {
	o.LiveStatus = &v
}

func (o SpecialFixture) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.BetType != nil {
		toSerialize["betType"] = o.BetType
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Date != nil {
		toSerialize["date"] = o.Date
	}
	if o.Cutoff != nil {
		toSerialize["cutoff"] = o.Cutoff
	}
	if o.Category != nil {
		toSerialize["category"] = o.Category
	}
	if o.Units != nil {
		toSerialize["units"] = o.Units
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Event != nil {
		toSerialize["event"] = o.Event
	}
	if o.Contestants != nil {
		toSerialize["contestants"] = o.Contestants
	}
	if o.LiveStatus != nil {
		toSerialize["liveStatus"] = o.LiveStatus
	}
	return json.Marshal(toSerialize)
}

type NullableSpecialFixture struct {
	value *SpecialFixture
	isSet bool
}

func (v NullableSpecialFixture) Get() *SpecialFixture {
	return v.value
}

func (v *NullableSpecialFixture) Set(val *SpecialFixture) {
	v.value = val
	v.isSet = true
}

func (v NullableSpecialFixture) IsSet() bool {
	return v.isSet
}

func (v *NullableSpecialFixture) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpecialFixture(val *SpecialFixture) *NullableSpecialFixture {
	return &NullableSpecialFixture{value: val, isSet: true}
}

func (v NullableSpecialFixture) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpecialFixture) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
