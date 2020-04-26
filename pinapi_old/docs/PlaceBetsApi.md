# \PlaceBetsApi

All URIs are relative to *https://api.pinnacle.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BetsSpecial**](PlaceBetsApi.md#BetsSpecial) | **Post** /v1/bets/special | Place specials bet.
[**BetsStraightV2**](PlaceBetsApi.md#BetsStraightV2) | **Post** /v2/bets/straight | Place straight bet  - v2



## BetsSpecial

> MultiBetResponseSpecialBetResponse BetsSpecial(ctx).Request(request).Execute()

Place specials bet.



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBetsSpecialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**MultiBetRequestSpecialBetRequest**](MultiBetRequestSpecialBetRequest.md) | The SpecialBet request. | 

### Return type

[**MultiBetResponseSpecialBetResponse**](MultiBetResponse.SpecialBetResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BetsStraightV2

> PlaceBetResponseV2 BetsStraightV2(ctx).Request(request).Execute()

Place straight bet  - v2



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBetsStraightV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**PlaceBetRequest**](PlaceBetRequest.md) |  | 

### Return type

[**PlaceBetResponseV2**](PlaceBetResponseV2.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

