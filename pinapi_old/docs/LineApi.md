# \LineApi

All URIs are relative to *https://api.pinnacle.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LineSpecialV1Get**](LineApi.md#LineSpecialV1Get) | **Get** /v1/line/special | Get Special Line - v1
[**LineStraightV1Get**](LineApi.md#LineStraightV1Get) | **Get** /v1/line | Get Straight Line - v1



## LineSpecialV1Get

> SpecialLineResponse LineSpecialV1Get(ctx).OddsFormat(oddsFormat).SpecialId(specialId).ContestantId(contestantId).Handicap(handicap).Execute()

Get Special Line - v1



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLineSpecialV1GetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oddsFormat** | **string** | Format the odds are returned in. [American, Decimal, HongKong, Indonesian, Malay] | 
 **specialId** | **int64** | Id of the special. | 
 **contestantId** | **int64** | Id of the contestant. | 
 **handicap** | **int32** | handicap of the contestant. As contestant&#39;s handicap is a mutable property, it may happened that line/special returns status:SUCCESS, but with the different handicap from the one that client had at the moment of calling the line/special. One can specify handicap parameter in the request and if the contestant&#39;s handicap changed, it would return status:NOT_EXISTS. This way line/special is more aligned to how /line works. | 

### Return type

[**SpecialLineResponse**](SpecialLineResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LineStraightV1Get

> LineResponse LineStraightV1Get(ctx).LeagueId(leagueId).Handicap(handicap).OddsFormat(oddsFormat).SportId(sportId).EventId(eventId).PeriodNumber(periodNumber).BetType(betType).Team(team).Side(side).Execute()

Get Straight Line - v1



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLineStraightV1GetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **leagueId** | **int32** | League Id. | 
 **handicap** | **float64** | This is needed for SPREAD, TOTAL_POINTS and TEAM_TOTAL_POINTS bet types | 
 **oddsFormat** | **string** | Format in which we return the odds. Default is American. | 
 **sportId** | **int32** | Sport identification | 
 **eventId** | **int64** | Event identification | 
 **periodNumber** | **int32** | This represents the period of the match. For example, for soccer we have 0 (Game),  1 (1st Half) &amp; 2 (2nd Half) | 
 **betType** | **string** | Bet Type | 
 **team** | **string** | Chosen team type. This is needed only for SPREAD, MONEYLINE and TEAM_TOTAL_POINTS bet types | 
 **side** | **string** | Chosen side. This is needed only for TOTAL_POINTS and TEAM_TOTAL_POINTS | 

### Return type

[**LineResponse**](LineResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

