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

// AccountInfo struct for AccountInfo
type AccountInfo struct {
	LoginUserName *string `json:"LoginUserName,omitempty"`
}

// NewAccountInfo instantiates a new AccountInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountInfo() *AccountInfo {
	this := AccountInfo{}
	return &this
}

// NewAccountInfoWithDefaults instantiates a new AccountInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountInfoWithDefaults() *AccountInfo {
	this := AccountInfo{}
	return &this
}

// GetLoginUserName returns the LoginUserName field value if set, zero value otherwise.
func (o *AccountInfo) GetLoginUserName() string {
	if o == nil || o.LoginUserName == nil {
		var ret string
		return ret
	}
	return *o.LoginUserName
}

// GetLoginUserNameOk returns a tuple with the LoginUserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInfo) GetLoginUserNameOk() (*string, bool) {
	if o == nil || o.LoginUserName == nil {
		return nil, false
	}
	return o.LoginUserName, true
}

// HasLoginUserName returns a boolean if a field has been set.
func (o *AccountInfo) HasLoginUserName() bool {
	if o != nil && o.LoginUserName != nil {
		return true
	}

	return false
}

// SetLoginUserName gets a reference to the given string and assigns it to the LoginUserName field.
func (o *AccountInfo) SetLoginUserName(v string) {
	o.LoginUserName = &v
}

func (o AccountInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LoginUserName != nil {
		toSerialize["LoginUserName"] = o.LoginUserName
	}
	return json.Marshal(toSerialize)
}

type NullableAccountInfo struct {
	value *AccountInfo
	isSet bool
}

func (v NullableAccountInfo) Get() *AccountInfo {
	return v.value
}

func (v *NullableAccountInfo) Set(val *AccountInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountInfo(val *AccountInfo) *NullableAccountInfo {
	return &NullableAccountInfo{value: val, isSet: true}
}

func (v NullableAccountInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
