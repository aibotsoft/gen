# \FavoritesApi

All URIs are relative to *https://ismart.dafabet.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FavoritesFuckOff**](FavoritesApi.md#FavoritesFuckOff) | **Post** /Favorites | Returns a list of users.



## FavoritesFuckOff

> FavoritesResponse FavoritesFuckOff(ctx).Execute()

Returns a list of users.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FavoritesApi.FavoritesFuckOff(context.Background(), ).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FavoritesApi.FavoritesFuckOff``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FavoritesFuckOff`: FavoritesResponse
    fmt.Fprintf(os.Stdout, "Response from `FavoritesApi.FavoritesFuckOff`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFavoritesFuckOffRequest struct via the builder pattern


### Return type

[**FavoritesResponse**](FavoritesResponse.md)

### Authorization

[cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

