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

// AuthenticateRequest struct for AuthenticateRequest
type AuthenticateRequest struct {
	Username *string `json:"username,omitempty"`
	Password *string `json:"password,omitempty"`
}

// NewAuthenticateRequest instantiates a new AuthenticateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticateRequest() *AuthenticateRequest {
	this := AuthenticateRequest{}
	return &this
}

// NewAuthenticateRequestWithDefaults instantiates a new AuthenticateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticateRequestWithDefaults() *AuthenticateRequest {
	this := AuthenticateRequest{}
	return &this
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *AuthenticateRequest) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticateRequest) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *AuthenticateRequest) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *AuthenticateRequest) SetUsername(v string) {
	o.Username = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *AuthenticateRequest) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticateRequest) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *AuthenticateRequest) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *AuthenticateRequest) SetPassword(v string) {
	o.Password = &v
}

func (o AuthenticateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	return json.Marshal(toSerialize)
}

type NullableAuthenticateRequest struct {
	value *AuthenticateRequest
	isSet bool
}

func (v NullableAuthenticateRequest) Get() *AuthenticateRequest {
	return v.value
}

func (v *NullableAuthenticateRequest) Set(val *AuthenticateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticateRequest(val *AuthenticateRequest) *NullableAuthenticateRequest {
	return &NullableAuthenticateRequest{value: val, isSet: true}
}

func (v NullableAuthenticateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
