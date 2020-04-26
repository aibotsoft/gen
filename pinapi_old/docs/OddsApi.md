# \OddsApi

All URIs are relative to *https://api.pinnacle.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OddsSpecialV1Get**](OddsApi.md#OddsSpecialV1Get) | **Get** /v1/odds/special | Get Special Odds - v1
[**OddsStraightV1Get**](OddsApi.md#OddsStraightV1Get) | **Get** /v1/odds | Get Straight Odds - v1



## OddsSpecialV1Get

> SpecialOddsResponse OddsSpecialV1Get(ctx).SportId(sportId).OddsFormat(oddsFormat).LeagueIds(leagueIds).Since(since).SpecialId(specialId).Execute()

Get Special Odds - v1



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOddsSpecialV1GetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sportId** | **int32** | Id of a sport for which to retrieve the specials. | 
 **oddsFormat** | **string** | Format the odds are returned in. [American, Decimal, HongKong, Indonesian, Malay] | 
 **leagueIds** | [**[]int32**](int32.md) | The leagueIds array may contain a list of comma separated league ids. | 
 **since** | **int64** | This is used to receive incremental updates. Use the value of last from previous response. When since parameter is not provided, the fixtures are delayed up to 1 min to encourage the use of the parameter. | 
 **specialId** | **int64** | Id of the special. This is an optional argument. | 

### Return type

[**SpecialOddsResponse**](SpecialOddsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OddsStraightV1Get

> OddsResponse OddsStraightV1Get(ctx).SportId(sportId).LeagueIds(leagueIds).OddsFormat(oddsFormat).Since(since).IsLive(isLive).EventIds(eventIds).ToCurrencyCode(toCurrencyCode).Execute()

Get Straight Odds - v1



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOddsStraightV1GetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sportId** | **int32** | The sportid for which to retrieve the odds. | 
 **leagueIds** | [**[]int32**](int32.md) | The leagueIds array may contain a list of comma separated league ids. | 
 **oddsFormat** | **string** | Format in which we return the odds. Default is American. [American, Decimal, HongKong, Indonesian, Malay] | 
 **since** | **int64** | This is used to receive incremental updates. Use the value of last from previous odds response. When since parameter is not provided, the odds are delayed up to 1 min to encourage the use of the parameter. Please note that when using since parameter you will get in the response ONLY changed periods. If a period did not have any changes it will not be in the response. | 
 **isLive** | **bool** | To retrieve ONLY live odds set the value to 1 (isLive&#x3D;1). Otherwise response will have all odds. | 
 **eventIds** | [**[]int64**](int64.md) | Filter by EventIds | 
 **toCurrencyCode** | **string** | 3 letter currency code as in the [/currency](https://pinnacleapi.github.io/linesapi#operation/Currencies_V2_Get) response. Limits will be returned in the requested currency. Default is USD. | 

### Return type

[**OddsResponse**](OddsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

