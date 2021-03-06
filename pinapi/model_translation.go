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

// Translation struct for Translation
type Translation struct {
	// Original requested text to be translated.
	Text *string `json:"text,omitempty"`
	// Collection of translations by culture.
	Cultures *[]TranslationByCulture `json:"cultures,omitempty"`
}

// NewTranslation instantiates a new Translation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTranslation() *Translation {
	this := Translation{}
	return &this
}

// NewTranslationWithDefaults instantiates a new Translation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTranslationWithDefaults() *Translation {
	this := Translation{}
	return &this
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *Translation) GetText() string {
	if o == nil || o.Text == nil {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Translation) GetTextOk() (*string, bool) {
	if o == nil || o.Text == nil {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *Translation) HasText() bool {
	if o != nil && o.Text != nil {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *Translation) SetText(v string) {
	o.Text = &v
}

// GetCultures returns the Cultures field value if set, zero value otherwise.
func (o *Translation) GetCultures() []TranslationByCulture {
	if o == nil || o.Cultures == nil {
		var ret []TranslationByCulture
		return ret
	}
	return *o.Cultures
}

// GetCulturesOk returns a tuple with the Cultures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Translation) GetCulturesOk() (*[]TranslationByCulture, bool) {
	if o == nil || o.Cultures == nil {
		return nil, false
	}
	return o.Cultures, true
}

// HasCultures returns a boolean if a field has been set.
func (o *Translation) HasCultures() bool {
	if o != nil && o.Cultures != nil {
		return true
	}

	return false
}

// SetCultures gets a reference to the given []TranslationByCulture and assigns it to the Cultures field.
func (o *Translation) SetCultures(v []TranslationByCulture) {
	o.Cultures = &v
}

func (o Translation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Text != nil {
		toSerialize["text"] = o.Text
	}
	if o.Cultures != nil {
		toSerialize["cultures"] = o.Cultures
	}
	return json.Marshal(toSerialize)
}

type NullableTranslation struct {
	value *Translation
	isSet bool
}

func (v NullableTranslation) Get() *Translation {
	return v.value
}

func (v *NullableTranslation) Set(val *Translation) {
	v.value = val
	v.isSet = true
}

func (v NullableTranslation) IsSet() bool {
	return v.isSet
}

func (v *NullableTranslation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTranslation(val *Translation) *NullableTranslation {
	return &NullableTranslation{value: val, isSet: true}
}

func (v NullableTranslation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTranslation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
