# Translation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | Pointer to **string** | Original requested text to be translated. | [optional] 
**Cultures** | Pointer to [**[]TranslationByCulture**](TranslationByCulture.md) | Collection of translations by culture. | [optional] 

## Methods

### GetText

`func (o *Translation) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *Translation) GetTextOk() (string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasText

`func (o *Translation) HasText() bool`

HasText returns a boolean if a field has been set.

### SetText

`func (o *Translation) SetText(v string)`

SetText gets a reference to the given string and assigns it to the Text field.

### GetCultures

`func (o *Translation) GetCultures() []TranslationByCulture`

GetCultures returns the Cultures field if non-nil, zero value otherwise.

### GetCulturesOk

`func (o *Translation) GetCulturesOk() ([]TranslationByCulture, bool)`

GetCulturesOk returns a tuple with the Cultures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCultures

`func (o *Translation) HasCultures() bool`

HasCultures returns a boolean if a field has been set.

### SetCultures

`func (o *Translation) SetCultures(v []TranslationByCulture)`

SetCultures gets a reference to the given []TranslationByCulture and assigns it to the Cultures field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


