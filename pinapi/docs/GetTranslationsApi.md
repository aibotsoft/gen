# \GetTranslationsApi

All URIs are relative to *https://api.pinnacle.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TranslationsV1Get**](GetTranslationsApi.md#TranslationsV1Get) | **Get** /v1/translations | Get Translations - v1



## TranslationsV1Get

> TranslationResponse TranslationsV1Get(ctx).CultureCodes(cultureCodes).BaseTexts(baseTexts).Execute()

Get Translations - v1



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTranslationsV1GetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cultureCodes** | [**[]string**](string.md) | Array of language cultures separated with |. | 
 **baseTexts** | [**[]string**](string.md) | Array of base texts to be translated separated with |. Each base text in the array must be URL encoded. Base texts are not case sensitive. | 

### Return type

[**TranslationResponse**](TranslationResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

