# \FixturesApi

All URIs are relative to *https://api.pinnacle.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FixturesSettledV1Get**](FixturesApi.md#FixturesSettledV1Get) | **Get** /v1/fixtures/settled | Get Settled Fixtures - v1
[**FixturesSpecialV1Get**](FixturesApi.md#FixturesSpecialV1Get) | **Get** /v1/fixtures/special | Get Special Fixtures - v1
[**FixturesSpecialsSettledV1Get**](FixturesApi.md#FixturesSpecialsSettledV1Get) | **Get** /v1/fixtures/special/settled | Get Settled Special Fixtures - v1
[**FixturesV1Get**](FixturesApi.md#FixturesV1Get) | **Get** /v1/fixtures | Get Fixtures - v1



## FixturesSettledV1Get

> SettledFixturesSport FixturesSettledV1Get(ctx).SportId(sportId).LeagueIds(leagueIds).Since(since).Execute()

Get Settled Fixtures - v1



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFixturesSettledV1GetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sportId** | **int32** |  | 
 **leagueIds** | [**[]int32**](int32.md) |  | 
 **since** | **int32** |  | 

### Return type

[**SettledFixturesSport**](SettledFixturesSport.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FixturesSpecialV1Get

> SpecialsFixturesResponse FixturesSpecialV1Get(ctx).SportId(sportId).LeagueIds(leagueIds).Since(since).Category(category).EventId(eventId).SpecialId(specialId).Execute()

Get Special Fixtures - v1



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFixturesSpecialV1GetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sportId** | **int32** | Id of a sport for which to retrieve the specials. | 
 **leagueIds** | [**[]int32**](int32.md) | The leagueIds array may contain a list of comma separated league ids. | 
 **since** | **int64** | This is used to receive incremental updates. Use the value of last field from the previous response. When since parameter is not provided, the fixtures are delayed up to 1 min to encourage the use of the parameter. | 
 **category** | **string** | The category the special falls under. | 
 **eventId** | **int64** | Id of an event associated with a special. | 
 **specialId** | **int64** | Id of the special. | 

### Return type

[**SpecialsFixturesResponse**](SpecialsFixturesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FixturesSpecialsSettledV1Get

> SettledSpecialsResponse FixturesSpecialsSettledV1Get(ctx).SportId(sportId).LeagueIds(leagueIds).Since(since).Execute()

Get Settled Special Fixtures - v1



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFixturesSpecialsSettledV1GetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sportId** | **int32** | Id of the sport for which to retrieve the settled specials. | 
 **leagueIds** | [**[]int32**](int32.md) | Array of leagueIds. This is optional parameter. | 
 **since** | **int64** | This is used to receive incremental updates. Use the value of last from previous response. | 

### Return type

[**SettledSpecialsResponse**](SettledSpecialsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FixturesV1Get

> FixturesResponse FixturesV1Get(ctx).SportId(sportId).LeagueIds(leagueIds).IsLive(isLive).Since(since).EventIds(eventIds).Execute()

Get Fixtures - v1



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFixturesV1GetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sportId** | **int32** | The sport id to retrieve the fixtures for. | 
 **leagueIds** | [**[]int32**](int32.md) | The leagueIds array may contain a list of comma separated league ids. | 
 **isLive** | **bool** | To retrieve ONLY live events set the value to 1 (isLive&#x3D;1). Missing or any other value will result in retrieval of events regardless of their Live status. | 
 **since** | **int64** | This is used to receive incremental updates. Use the value of last from previous fixtures response. When since parameter is not provided, the fixtures are delayed up to 1 minute to encourage the use of the parameter. | 
 **eventIds** | [**[]int32**](int32.md) | Comma separated list of event ids to filter by | 

### Return type

[**FixturesResponse**](FixturesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

