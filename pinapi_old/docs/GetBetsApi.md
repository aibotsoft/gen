# \GetBetsApi

All URIs are relative to *https://api.pinnacle.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BetsGetBetsByTypeV3**](GetBetsApi.md#BetsGetBetsByTypeV3) | **Get** /v3/bets | Get Bets - v3



## BetsGetBetsByTypeV3

> GetBetsByTypeResponseV3 BetsGetBetsByTypeV3(ctx).Betlist(betlist).BetStatuses(betStatuses).FromDate(fromDate).ToDate(toDate).SortDir(sortDir).PageSize(pageSize).FromRecord(fromRecord).Betids(betids).UniqueRequestIds(uniqueRequestIds).Execute()

Get Bets - v3



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBetsGetBetsByTypeV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **betlist** | **string** | Type of bet list to return. Not needed when betids is submitted. | 
 **betStatuses** | [**[]string**](string.md) | Type of bet statues to return. This works only in conjustion with betlist, as additional filter. | 
 **fromDate** | **string** | Start date of the requested period. Required when betlist parameter is submitted. Start date can be up to 30 days in the past. Expected format is ISO8601 - can be set to just date or date and time.  | 
 **toDate** | **string** | End date of the requested period. Required when betlist parameter is submitted. Expected format is ISO8601 - can be set to just date or date and time.  toDate value is exclusive, meaning it cannot be equal to fromDate.  | 
 **sortDir** | **string** | Sort direction by postedAt/settledAt. Respected only when querying by date range. | [default to ASC]
 **pageSize** | **int32** | Page size in case. Max is 1000. Respected only when querying by date range. | [default to 1000]
 **fromRecord** | **int32** | Starting record (inclusive) of the result. Respected only when querying by date range. To fetch next page set it to toRecord+1  | [default to 0]
 **betids** | [**[]int64**](int64.md) | A comma separated list of bet ids. When betids is submitted, no other parameter is necessary. Maximum is 100 ids. Works for all non settled bets and all bets settled in the last 30 days. | 
 **uniqueRequestIds** | [**[]string**](string.md) | A comma separated list of &#x60;uniqueRequestId&#x60; from the place bet request. If specified, it&#39;s highest priority, all other parameters are ignored. Maximum is 10 ids. If client has bet id, preferred way is to use &#x60;betIds&#x60; query parameter, you can use &#x60;uniqueRequestIds&#x60; when you do not  have bet id.  That are 2 cases when client may not have a bet id:  1. When you bet on live event with live delay, place bet response in that case does not return bet id, so client can query bet status by &#x60;uniqueRequestIds&#x60;. 2. In case of any network issues when client is not sure what happened with his place bet request. Empty response means that the bet was not placed. Please check [Deduplication section](https://www.pinnacle.com/de/api/manual#overview) for more details.  Note that there is a restriction: querying by uniqueRequestIds  is supported for straight and  special bets and only up to 30 min from the moment the bet was place.   | 

### Return type

[**GetBetsByTypeResponseV3**](GetBetsByTypeResponseV3.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

