/*
 * Pinnacle - Customer API Reference
 *
 *  # Authentication   API uses HTTP Basic access authentication. You need to send Authorization HTTP Request header:    `Authorization: Basic <Base64 value of UTF-8 encoded \"username:password\">`  Example:  `Authorization: Basic U03MyOT23YbzMDc6d3c3O1DQ1`
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pinapi

import (
	"bytes"
	"encoding/json"
)

// TranslationResponse struct for TranslationResponse
type TranslationResponse struct {
	// Collection of translations.
	Translations *[]Translation `json:"translations,omitempty" xml:"translations"`
}

// GetTranslations returns the Translations field value if set, zero value otherwise.
func (o *TranslationResponse) GetTranslations() []Translation {
	if o == nil || o.Translations == nil {
		var ret []Translation
		return ret
	}
	return *o.Translations
}

// GetTranslationsOk returns a tuple with the Translations field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *TranslationResponse) GetTranslationsOk() ([]Translation, bool) {
	if o == nil || o.Translations == nil {
		var ret []Translation
		return ret, false
	}
	return *o.Translations, true
}

// HasTranslations returns a boolean if a field has been set.
func (o *TranslationResponse) HasTranslations() bool {
	if o != nil && o.Translations != nil {
		return true
	}

	return false
}

// SetTranslations gets a reference to the given []Translation and assigns it to the Translations field.
func (o *TranslationResponse) SetTranslations(v []Translation) {
	o.Translations = &v
}

type NullableTranslationResponse struct {
	Value        TranslationResponse
	ExplicitNull bool
}

func (v NullableTranslationResponse) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableTranslationResponse) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
