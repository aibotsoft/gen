/*
 * EpinApi
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package epinapi

import (
	"encoding/json"
	"time"
)

// LoginResponse struct for LoginResponse
type LoginResponse struct {
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	ExpiresAt *time.Time `json:"expiresAt,omitempty"`
	IsExpired *bool `json:"isExpired,omitempty"`
	IsSuperseded *bool `json:"isSuperseded,omitempty"`
	IsTimedOut *bool `json:"isTimedOut,omitempty"`
	LastUsedAt *time.Time `json:"lastUsedAt,omitempty"`
	License *string `json:"license,omitempty"`
	PreferredView *string `json:"preferredView,omitempty"`
	Token *string `json:"token,omitempty"`
	TransactionId *string `json:"transactionId,omitempty"`
	TrustCode *string `json:"trustCode,omitempty"`
	Username *string `json:"username,omitempty"`
}

// NewLoginResponse instantiates a new LoginResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoginResponse() *LoginResponse {
	this := LoginResponse{}
	return &this
}

// NewLoginResponseWithDefaults instantiates a new LoginResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoginResponseWithDefaults() *LoginResponse {
	this := LoginResponse{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *LoginResponse) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *LoginResponse) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *LoginResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *LoginResponse) GetExpiresAt() time.Time {
	if o == nil || o.ExpiresAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginResponse) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || o.ExpiresAt == nil {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *LoginResponse) HasExpiresAt() bool {
	if o != nil && o.ExpiresAt != nil {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given time.Time and assigns it to the ExpiresAt field.
func (o *LoginResponse) SetExpiresAt(v time.Time) {
	o.ExpiresAt = &v
}

// GetIsExpired returns the IsExpired field value if set, zero value otherwise.
func (o *LoginResponse) GetIsExpired() bool {
	if o == nil || o.IsExpired == nil {
		var ret bool
		return ret
	}
	return *o.IsExpired
}

// GetIsExpiredOk returns a tuple with the IsExpired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginResponse) GetIsExpiredOk() (*bool, bool) {
	if o == nil || o.IsExpired == nil {
		return nil, false
	}
	return o.IsExpired, true
}

// HasIsExpired returns a boolean if a field has been set.
func (o *LoginResponse) HasIsExpired() bool {
	if o != nil && o.IsExpired != nil {
		return true
	}

	return false
}

// SetIsExpired gets a reference to the given bool and assigns it to the IsExpired field.
func (o *LoginResponse) SetIsExpired(v bool) {
	o.IsExpired = &v
}

// GetIsSuperseded returns the IsSuperseded field value if set, zero value otherwise.
func (o *LoginResponse) GetIsSuperseded() bool {
	if o == nil || o.IsSuperseded == nil {
		var ret bool
		return ret
	}
	return *o.IsSuperseded
}

// GetIsSupersededOk returns a tuple with the IsSuperseded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginResponse) GetIsSupersededOk() (*bool, bool) {
	if o == nil || o.IsSuperseded == nil {
		return nil, false
	}
	return o.IsSuperseded, true
}

// HasIsSuperseded returns a boolean if a field has been set.
func (o *LoginResponse) HasIsSuperseded() bool {
	if o != nil && o.IsSuperseded != nil {
		return true
	}

	return false
}

// SetIsSuperseded gets a reference to the given bool and assigns it to the IsSuperseded field.
func (o *LoginResponse) SetIsSuperseded(v bool) {
	o.IsSuperseded = &v
}

// GetIsTimedOut returns the IsTimedOut field value if set, zero value otherwise.
func (o *LoginResponse) GetIsTimedOut() bool {
	if o == nil || o.IsTimedOut == nil {
		var ret bool
		return ret
	}
	return *o.IsTimedOut
}

// GetIsTimedOutOk returns a tuple with the IsTimedOut field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginResponse) GetIsTimedOutOk() (*bool, bool) {
	if o == nil || o.IsTimedOut == nil {
		return nil, false
	}
	return o.IsTimedOut, true
}

// HasIsTimedOut returns a boolean if a field has been set.
func (o *LoginResponse) HasIsTimedOut() bool {
	if o != nil && o.IsTimedOut != nil {
		return true
	}

	return false
}

// SetIsTimedOut gets a reference to the given bool and assigns it to the IsTimedOut field.
func (o *LoginResponse) SetIsTimedOut(v bool) {
	o.IsTimedOut = &v
}

// GetLastUsedAt returns the LastUsedAt field value if set, zero value otherwise.
func (o *LoginResponse) GetLastUsedAt() time.Time {
	if o == nil || o.LastUsedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUsedAt
}

// GetLastUsedAtOk returns a tuple with the LastUsedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginResponse) GetLastUsedAtOk() (*time.Time, bool) {
	if o == nil || o.LastUsedAt == nil {
		return nil, false
	}
	return o.LastUsedAt, true
}

// HasLastUsedAt returns a boolean if a field has been set.
func (o *LoginResponse) HasLastUsedAt() bool {
	if o != nil && o.LastUsedAt != nil {
		return true
	}

	return false
}

// SetLastUsedAt gets a reference to the given time.Time and assigns it to the LastUsedAt field.
func (o *LoginResponse) SetLastUsedAt(v time.Time) {
	o.LastUsedAt = &v
}

// GetLicense returns the License field value if set, zero value otherwise.
func (o *LoginResponse) GetLicense() string {
	if o == nil || o.License == nil {
		var ret string
		return ret
	}
	return *o.License
}

// GetLicenseOk returns a tuple with the License field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginResponse) GetLicenseOk() (*string, bool) {
	if o == nil || o.License == nil {
		return nil, false
	}
	return o.License, true
}

// HasLicense returns a boolean if a field has been set.
func (o *LoginResponse) HasLicense() bool {
	if o != nil && o.License != nil {
		return true
	}

	return false
}

// SetLicense gets a reference to the given string and assigns it to the License field.
func (o *LoginResponse) SetLicense(v string) {
	o.License = &v
}

// GetPreferredView returns the PreferredView field value if set, zero value otherwise.
func (o *LoginResponse) GetPreferredView() string {
	if o == nil || o.PreferredView == nil {
		var ret string
		return ret
	}
	return *o.PreferredView
}

// GetPreferredViewOk returns a tuple with the PreferredView field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginResponse) GetPreferredViewOk() (*string, bool) {
	if o == nil || o.PreferredView == nil {
		return nil, false
	}
	return o.PreferredView, true
}

// HasPreferredView returns a boolean if a field has been set.
func (o *LoginResponse) HasPreferredView() bool {
	if o != nil && o.PreferredView != nil {
		return true
	}

	return false
}

// SetPreferredView gets a reference to the given string and assigns it to the PreferredView field.
func (o *LoginResponse) SetPreferredView(v string) {
	o.PreferredView = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *LoginResponse) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginResponse) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *LoginResponse) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *LoginResponse) SetToken(v string) {
	o.Token = &v
}

// GetTransactionId returns the TransactionId field value if set, zero value otherwise.
func (o *LoginResponse) GetTransactionId() string {
	if o == nil || o.TransactionId == nil {
		var ret string
		return ret
	}
	return *o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginResponse) GetTransactionIdOk() (*string, bool) {
	if o == nil || o.TransactionId == nil {
		return nil, false
	}
	return o.TransactionId, true
}

// HasTransactionId returns a boolean if a field has been set.
func (o *LoginResponse) HasTransactionId() bool {
	if o != nil && o.TransactionId != nil {
		return true
	}

	return false
}

// SetTransactionId gets a reference to the given string and assigns it to the TransactionId field.
func (o *LoginResponse) SetTransactionId(v string) {
	o.TransactionId = &v
}

// GetTrustCode returns the TrustCode field value if set, zero value otherwise.
func (o *LoginResponse) GetTrustCode() string {
	if o == nil || o.TrustCode == nil {
		var ret string
		return ret
	}
	return *o.TrustCode
}

// GetTrustCodeOk returns a tuple with the TrustCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginResponse) GetTrustCodeOk() (*string, bool) {
	if o == nil || o.TrustCode == nil {
		return nil, false
	}
	return o.TrustCode, true
}

// HasTrustCode returns a boolean if a field has been set.
func (o *LoginResponse) HasTrustCode() bool {
	if o != nil && o.TrustCode != nil {
		return true
	}

	return false
}

// SetTrustCode gets a reference to the given string and assigns it to the TrustCode field.
func (o *LoginResponse) SetTrustCode(v string) {
	o.TrustCode = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *LoginResponse) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginResponse) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *LoginResponse) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *LoginResponse) SetUsername(v string) {
	o.Username = &v
}

func (o LoginResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.ExpiresAt != nil {
		toSerialize["expiresAt"] = o.ExpiresAt
	}
	if o.IsExpired != nil {
		toSerialize["isExpired"] = o.IsExpired
	}
	if o.IsSuperseded != nil {
		toSerialize["isSuperseded"] = o.IsSuperseded
	}
	if o.IsTimedOut != nil {
		toSerialize["isTimedOut"] = o.IsTimedOut
	}
	if o.LastUsedAt != nil {
		toSerialize["lastUsedAt"] = o.LastUsedAt
	}
	if o.License != nil {
		toSerialize["license"] = o.License
	}
	if o.PreferredView != nil {
		toSerialize["preferredView"] = o.PreferredView
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.TransactionId != nil {
		toSerialize["transactionId"] = o.TransactionId
	}
	if o.TrustCode != nil {
		toSerialize["trustCode"] = o.TrustCode
	}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	return json.Marshal(toSerialize)
}

type NullableLoginResponse struct {
	value *LoginResponse
	isSet bool
}

func (v NullableLoginResponse) Get() *LoginResponse {
	return v.value
}

func (v *NullableLoginResponse) Set(val *LoginResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLoginResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLoginResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoginResponse(val *LoginResponse) *NullableLoginResponse {
	return &NullableLoginResponse{value: val, isSet: true}
}

func (v NullableLoginResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoginResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
