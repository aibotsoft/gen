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

// LoginData struct for LoginData
type LoginData struct {
	Whitelabel *string `json:"whitelabel,omitempty"`
	Username *string `json:"username,omitempty"`
	CustomerId *int64 `json:"customer_id,omitempty"`
	Lang *string `json:"lang,omitempty"`
	SessionId *string `json:"session_id,omitempty"`
}

// NewLoginData instantiates a new LoginData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoginData() *LoginData {
	this := LoginData{}
	return &this
}

// NewLoginDataWithDefaults instantiates a new LoginData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoginDataWithDefaults() *LoginData {
	this := LoginData{}
	return &this
}

// GetWhitelabel returns the Whitelabel field value if set, zero value otherwise.
func (o *LoginData) GetWhitelabel() string {
	if o == nil || o.Whitelabel == nil {
		var ret string
		return ret
	}
	return *o.Whitelabel
}

// GetWhitelabelOk returns a tuple with the Whitelabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginData) GetWhitelabelOk() (*string, bool) {
	if o == nil || o.Whitelabel == nil {
		return nil, false
	}
	return o.Whitelabel, true
}

// HasWhitelabel returns a boolean if a field has been set.
func (o *LoginData) HasWhitelabel() bool {
	if o != nil && o.Whitelabel != nil {
		return true
	}

	return false
}

// SetWhitelabel gets a reference to the given string and assigns it to the Whitelabel field.
func (o *LoginData) SetWhitelabel(v string) {
	o.Whitelabel = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *LoginData) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginData) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *LoginData) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *LoginData) SetUsername(v string) {
	o.Username = &v
}

// GetCustomerId returns the CustomerId field value if set, zero value otherwise.
func (o *LoginData) GetCustomerId() int64 {
	if o == nil || o.CustomerId == nil {
		var ret int64
		return ret
	}
	return *o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginData) GetCustomerIdOk() (*int64, bool) {
	if o == nil || o.CustomerId == nil {
		return nil, false
	}
	return o.CustomerId, true
}

// HasCustomerId returns a boolean if a field has been set.
func (o *LoginData) HasCustomerId() bool {
	if o != nil && o.CustomerId != nil {
		return true
	}

	return false
}

// SetCustomerId gets a reference to the given int64 and assigns it to the CustomerId field.
func (o *LoginData) SetCustomerId(v int64) {
	o.CustomerId = &v
}

// GetLang returns the Lang field value if set, zero value otherwise.
func (o *LoginData) GetLang() string {
	if o == nil || o.Lang == nil {
		var ret string
		return ret
	}
	return *o.Lang
}

// GetLangOk returns a tuple with the Lang field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginData) GetLangOk() (*string, bool) {
	if o == nil || o.Lang == nil {
		return nil, false
	}
	return o.Lang, true
}

// HasLang returns a boolean if a field has been set.
func (o *LoginData) HasLang() bool {
	if o != nil && o.Lang != nil {
		return true
	}

	return false
}

// SetLang gets a reference to the given string and assigns it to the Lang field.
func (o *LoginData) SetLang(v string) {
	o.Lang = &v
}

// GetSessionId returns the SessionId field value if set, zero value otherwise.
func (o *LoginData) GetSessionId() string {
	if o == nil || o.SessionId == nil {
		var ret string
		return ret
	}
	return *o.SessionId
}

// GetSessionIdOk returns a tuple with the SessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginData) GetSessionIdOk() (*string, bool) {
	if o == nil || o.SessionId == nil {
		return nil, false
	}
	return o.SessionId, true
}

// HasSessionId returns a boolean if a field has been set.
func (o *LoginData) HasSessionId() bool {
	if o != nil && o.SessionId != nil {
		return true
	}

	return false
}

// SetSessionId gets a reference to the given string and assigns it to the SessionId field.
func (o *LoginData) SetSessionId(v string) {
	o.SessionId = &v
}

func (o LoginData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Whitelabel != nil {
		toSerialize["whitelabel"] = o.Whitelabel
	}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	if o.CustomerId != nil {
		toSerialize["customer_id"] = o.CustomerId
	}
	if o.Lang != nil {
		toSerialize["lang"] = o.Lang
	}
	if o.SessionId != nil {
		toSerialize["session_id"] = o.SessionId
	}
	return json.Marshal(toSerialize)
}

type NullableLoginData struct {
	value *LoginData
	isSet bool
}

func (v NullableLoginData) Get() *LoginData {
	return v.value
}

func (v *NullableLoginData) Set(val *LoginData) {
	v.value = val
	v.isSet = true
}

func (v NullableLoginData) IsSet() bool {
	return v.isSet
}

func (v *NullableLoginData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoginData(val *LoginData) *NullableLoginData {
	return &NullableLoginData{value: val, isSet: true}
}

func (v NullableLoginData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoginData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
