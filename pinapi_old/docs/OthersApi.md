# \OthersApi

All URIs are relative to *https://api.pinnacle.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancellationReasonsV1Get**](OthersApi.md#CancellationReasonsV1Get) | **Get** /v1/cancellationreasons | Get Cancellation Reasons - v1
[**CurrenciesV2Get**](OthersApi.md#CurrenciesV2Get) | **Get** /v2/currencies | Get Currencies - v2
[**InRunningV1Get**](OthersApi.md#InRunningV1Get) | **Get** /v1/inrunning | Get In-Running - v1
[**LeaguesV2Get**](OthersApi.md#LeaguesV2Get) | **Get** /v2/leagues | Get Leagues - v2
[**PeriodsV1Get**](OthersApi.md#PeriodsV1Get) | **Get** /v1/periods | Get Periods - v1
[**SportsV2Get**](OthersApi.md#SportsV2Get) | **Get** /v2/sports | Get Sports - v2
[**TeaserGroupsV1Get**](OthersApi.md#TeaserGroupsV1Get) | **Get** /v1/teaser/groups | Get Teaser Groups - v1



## CancellationReasonsV1Get

> CancellationReasonResponse CancellationReasonsV1Get(ctx).Execute()

Get Cancellation Reasons - v1



### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCancellationReasonsV1GetRequest struct via the builder pattern


### Return type

[**CancellationReasonResponse**](CancellationReasonResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CurrenciesV2Get

> SuccessfulCurrenciesResponse CurrenciesV2Get(ctx).Execute()

Get Currencies - v2



### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCurrenciesV2GetRequest struct via the builder pattern


### Return type

[**SuccessfulCurrenciesResponse**](SuccessfulCurrenciesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InRunningV1Get

> InRunningResponse InRunningV1Get(ctx).Execute()

Get In-Running - v1



### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiInRunningV1GetRequest struct via the builder pattern


### Return type

[**InRunningResponse**](InRunningResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LeaguesV2Get

> Leagues LeaguesV2Get(ctx).SportId(sportId).Execute()

Get Leagues - v2



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLeaguesV2GetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sportId** | **string** | Sport id for which the leagues are requested. | 

### Return type

[**Leagues**](Leagues.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PeriodsV1Get

> SportPeriod PeriodsV1Get(ctx).SportId(sportId).Execute()

Get Periods - v1



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPeriodsV1GetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sportId** | **string** |  | 

### Return type

[**SportPeriod**](SportPeriod.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SportsV2Get

> SportsResponse SportsV2Get(ctx).Execute()

Get Sports - v2



### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSportsV2GetRequest struct via the builder pattern


### Return type

[**SportsResponse**](SportsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeaserGroupsV1Get

> TeaserGroupsResponse TeaserGroupsV1Get(ctx).OddsFormat(oddsFormat).Execute()

Get Teaser Groups - v1



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTeaserGroupsV1GetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oddsFormat** | **string** | Format the odds are returned in. [American, Decimal, HongKong, Indonesian, Malay] | 

### Return type

[**TeaserGroupsResponse**](TeaserGroupsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

