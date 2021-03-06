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

// SettledFixturesPeriod struct for SettledFixturesPeriod
type SettledFixturesPeriod struct {
	// This represents the period of the match. For example, for soccer we have 0 (Game), 1 (1st Half) & 2 (2nd Half)
	Number *int `json:"number,omitempty"`
	// Period settlement status.   1 = Event period is settled,  2 = Event period is re-settled,  3 = Event period is cancelled,  4 = Event period is re-settled as cancelled,  5 = Event is deleted 
	Status *int `json:"status,omitempty"`
	// Unique id of the settlement. In case of a re-settlement, a new settlementId and settledAt will be generated.
	SettlementId *int64 `json:"settlementId,omitempty"`
	// Date and time in UTC when the period was settled.
	SettledAt *time.Time `json:"settledAt,omitempty"`
	// Team1 score.
	Team1Score *int `json:"team1Score,omitempty"`
	// Team2 score.
	Team2Score *int `json:"team2Score,omitempty"`
	// Team1 sets score. Supported for tennis only.
	Team1ScoreSets *int `json:"team1ScoreSets,omitempty"`
	// Team2 sets score. Supported for tennis only.
	Team2ScoreSets *int `json:"team2ScoreSets,omitempty"`
	CancellationReason *CancellationReasonType `json:"cancellationReason,omitempty"`
}

// NewSettledFixturesPeriod instantiates a new SettledFixturesPeriod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSettledFixturesPeriod() *SettledFixturesPeriod {
	this := SettledFixturesPeriod{}
	return &this
}

// NewSettledFixturesPeriodWithDefaults instantiates a new SettledFixturesPeriod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSettledFixturesPeriodWithDefaults() *SettledFixturesPeriod {
	this := SettledFixturesPeriod{}
	return &this
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *SettledFixturesPeriod) GetNumber() int {
	if o == nil || o.Number == nil {
		var ret int
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettledFixturesPeriod) GetNumberOk() (*int, bool) {
	if o == nil || o.Number == nil {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *SettledFixturesPeriod) HasNumber() bool {
	if o != nil && o.Number != nil {
		return true
	}

	return false
}

// SetNumber gets a reference to the given int and assigns it to the Number field.
func (o *SettledFixturesPeriod) SetNumber(v int) {
	o.Number = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SettledFixturesPeriod) GetStatus() int {
	if o == nil || o.Status == nil {
		var ret int
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettledFixturesPeriod) GetStatusOk() (*int, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SettledFixturesPeriod) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int and assigns it to the Status field.
func (o *SettledFixturesPeriod) SetStatus(v int) {
	o.Status = &v
}

// GetSettlementId returns the SettlementId field value if set, zero value otherwise.
func (o *SettledFixturesPeriod) GetSettlementId() int64 {
	if o == nil || o.SettlementId == nil {
		var ret int64
		return ret
	}
	return *o.SettlementId
}

// GetSettlementIdOk returns a tuple with the SettlementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettledFixturesPeriod) GetSettlementIdOk() (*int64, bool) {
	if o == nil || o.SettlementId == nil {
		return nil, false
	}
	return o.SettlementId, true
}

// HasSettlementId returns a boolean if a field has been set.
func (o *SettledFixturesPeriod) HasSettlementId() bool {
	if o != nil && o.SettlementId != nil {
		return true
	}

	return false
}

// SetSettlementId gets a reference to the given int64 and assigns it to the SettlementId field.
func (o *SettledFixturesPeriod) SetSettlementId(v int64) {
	o.SettlementId = &v
}

// GetSettledAt returns the SettledAt field value if set, zero value otherwise.
func (o *SettledFixturesPeriod) GetSettledAt() time.Time {
	if o == nil || o.SettledAt == nil {
		var ret time.Time
		return ret
	}
	return *o.SettledAt
}

// GetSettledAtOk returns a tuple with the SettledAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettledFixturesPeriod) GetSettledAtOk() (*time.Time, bool) {
	if o == nil || o.SettledAt == nil {
		return nil, false
	}
	return o.SettledAt, true
}

// HasSettledAt returns a boolean if a field has been set.
func (o *SettledFixturesPeriod) HasSettledAt() bool {
	if o != nil && o.SettledAt != nil {
		return true
	}

	return false
}

// SetSettledAt gets a reference to the given time.Time and assigns it to the SettledAt field.
func (o *SettledFixturesPeriod) SetSettledAt(v time.Time) {
	o.SettledAt = &v
}

// GetTeam1Score returns the Team1Score field value if set, zero value otherwise.
func (o *SettledFixturesPeriod) GetTeam1Score() int {
	if o == nil || o.Team1Score == nil {
		var ret int
		return ret
	}
	return *o.Team1Score
}

// GetTeam1ScoreOk returns a tuple with the Team1Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettledFixturesPeriod) GetTeam1ScoreOk() (*int, bool) {
	if o == nil || o.Team1Score == nil {
		return nil, false
	}
	return o.Team1Score, true
}

// HasTeam1Score returns a boolean if a field has been set.
func (o *SettledFixturesPeriod) HasTeam1Score() bool {
	if o != nil && o.Team1Score != nil {
		return true
	}

	return false
}

// SetTeam1Score gets a reference to the given int and assigns it to the Team1Score field.
func (o *SettledFixturesPeriod) SetTeam1Score(v int) {
	o.Team1Score = &v
}

// GetTeam2Score returns the Team2Score field value if set, zero value otherwise.
func (o *SettledFixturesPeriod) GetTeam2Score() int {
	if o == nil || o.Team2Score == nil {
		var ret int
		return ret
	}
	return *o.Team2Score
}

// GetTeam2ScoreOk returns a tuple with the Team2Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettledFixturesPeriod) GetTeam2ScoreOk() (*int, bool) {
	if o == nil || o.Team2Score == nil {
		return nil, false
	}
	return o.Team2Score, true
}

// HasTeam2Score returns a boolean if a field has been set.
func (o *SettledFixturesPeriod) HasTeam2Score() bool {
	if o != nil && o.Team2Score != nil {
		return true
	}

	return false
}

// SetTeam2Score gets a reference to the given int and assigns it to the Team2Score field.
func (o *SettledFixturesPeriod) SetTeam2Score(v int) {
	o.Team2Score = &v
}

// GetTeam1ScoreSets returns the Team1ScoreSets field value if set, zero value otherwise.
func (o *SettledFixturesPeriod) GetTeam1ScoreSets() int {
	if o == nil || o.Team1ScoreSets == nil {
		var ret int
		return ret
	}
	return *o.Team1ScoreSets
}

// GetTeam1ScoreSetsOk returns a tuple with the Team1ScoreSets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettledFixturesPeriod) GetTeam1ScoreSetsOk() (*int, bool) {
	if o == nil || o.Team1ScoreSets == nil {
		return nil, false
	}
	return o.Team1ScoreSets, true
}

// HasTeam1ScoreSets returns a boolean if a field has been set.
func (o *SettledFixturesPeriod) HasTeam1ScoreSets() bool {
	if o != nil && o.Team1ScoreSets != nil {
		return true
	}

	return false
}

// SetTeam1ScoreSets gets a reference to the given int and assigns it to the Team1ScoreSets field.
func (o *SettledFixturesPeriod) SetTeam1ScoreSets(v int) {
	o.Team1ScoreSets = &v
}

// GetTeam2ScoreSets returns the Team2ScoreSets field value if set, zero value otherwise.
func (o *SettledFixturesPeriod) GetTeam2ScoreSets() int {
	if o == nil || o.Team2ScoreSets == nil {
		var ret int
		return ret
	}
	return *o.Team2ScoreSets
}

// GetTeam2ScoreSetsOk returns a tuple with the Team2ScoreSets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettledFixturesPeriod) GetTeam2ScoreSetsOk() (*int, bool) {
	if o == nil || o.Team2ScoreSets == nil {
		return nil, false
	}
	return o.Team2ScoreSets, true
}

// HasTeam2ScoreSets returns a boolean if a field has been set.
func (o *SettledFixturesPeriod) HasTeam2ScoreSets() bool {
	if o != nil && o.Team2ScoreSets != nil {
		return true
	}

	return false
}

// SetTeam2ScoreSets gets a reference to the given int and assigns it to the Team2ScoreSets field.
func (o *SettledFixturesPeriod) SetTeam2ScoreSets(v int) {
	o.Team2ScoreSets = &v
}

// GetCancellationReason returns the CancellationReason field value if set, zero value otherwise.
func (o *SettledFixturesPeriod) GetCancellationReason() CancellationReasonType {
	if o == nil || o.CancellationReason == nil {
		var ret CancellationReasonType
		return ret
	}
	return *o.CancellationReason
}

// GetCancellationReasonOk returns a tuple with the CancellationReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettledFixturesPeriod) GetCancellationReasonOk() (*CancellationReasonType, bool) {
	if o == nil || o.CancellationReason == nil {
		return nil, false
	}
	return o.CancellationReason, true
}

// HasCancellationReason returns a boolean if a field has been set.
func (o *SettledFixturesPeriod) HasCancellationReason() bool {
	if o != nil && o.CancellationReason != nil {
		return true
	}

	return false
}

// SetCancellationReason gets a reference to the given CancellationReasonType and assigns it to the CancellationReason field.
func (o *SettledFixturesPeriod) SetCancellationReason(v CancellationReasonType) {
	o.CancellationReason = &v
}

func (o SettledFixturesPeriod) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Number != nil {
		toSerialize["number"] = o.Number
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
	if o.Team1Score != nil {
		toSerialize["team1Score"] = o.Team1Score
	}
	if o.Team2Score != nil {
		toSerialize["team2Score"] = o.Team2Score
	}
	if o.Team1ScoreSets != nil {
		toSerialize["team1ScoreSets"] = o.Team1ScoreSets
	}
	if o.Team2ScoreSets != nil {
		toSerialize["team2ScoreSets"] = o.Team2ScoreSets
	}
	if o.CancellationReason != nil {
		toSerialize["cancellationReason"] = o.CancellationReason
	}
	return json.Marshal(toSerialize)
}

type NullableSettledFixturesPeriod struct {
	value *SettledFixturesPeriod
	isSet bool
}

func (v NullableSettledFixturesPeriod) Get() *SettledFixturesPeriod {
	return v.value
}

func (v *NullableSettledFixturesPeriod) Set(val *SettledFixturesPeriod) {
	v.value = val
	v.isSet = true
}

func (v NullableSettledFixturesPeriod) IsSet() bool {
	return v.isSet
}

func (v *NullableSettledFixturesPeriod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSettledFixturesPeriod(val *SettledFixturesPeriod) *NullableSettledFixturesPeriod {
	return &NullableSettledFixturesPeriod{value: val, isSet: true}
}

func (v NullableSettledFixturesPeriod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSettledFixturesPeriod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
