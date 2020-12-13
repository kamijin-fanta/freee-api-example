# \RenewsApi

All URIs are relative to *https://api.freee.co.jp*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDealRenew**](RenewsApi.md#CreateDealRenew) | **Post** /api/1/deals/{id}/renews | 取引（収入／支出）に対する+更新の作成
[**DeleteDealRenew**](RenewsApi.md#DeleteDealRenew) | **Delete** /api/1/deals/{id}/renews/{renew_id} | 取引（収入／支出）の+更新の削除
[**UpdateDealRenew**](RenewsApi.md#UpdateDealRenew) | **Put** /api/1/deals/{id}/renews/{renew_id} | 取引（収入／支出）の+更新の更新



## CreateDealRenew

> DealResponse CreateDealRenew(ctx, id).RenewCreateParams(renewCreateParams).Execute()

取引（収入／支出）に対する+更新の作成



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
    id := int32(56) // int32 | 取引ID
    renewCreateParams := *openapiclient.NewRenewCreateParams(int32(1), "2014-01-01", int32(1), []openapiclient.RenewCreateParamsDetails{*openapiclient.NewRenewCreateParamsDetails(int32(1), int32(1), int32(1080))}) // RenewCreateParams | 取引（収入／支出）に対する+更新の情報

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RenewsApi.CreateDealRenew(context.Background(), id).RenewCreateParams(renewCreateParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RenewsApi.CreateDealRenew``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDealRenew`: DealResponse
    fmt.Fprintf(os.Stdout, "Response from `RenewsApi.CreateDealRenew`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | 取引ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDealRenewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **renewCreateParams** | [**RenewCreateParams**](RenewCreateParams.md) | 取引（収入／支出）に対する+更新の情報 | 

### Return type

[**DealResponse**](dealResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDealRenew

> DealResponse DeleteDealRenew(ctx, id, renewId).CompanyId(companyId).Execute()

取引（収入／支出）の+更新の削除



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
    id := int32(56) // int32 | 取引ID
    renewId := int32(56) // int32 | +更新ID
    companyId := int32(56) // int32 | 事業所ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RenewsApi.DeleteDealRenew(context.Background(), id, renewId).CompanyId(companyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RenewsApi.DeleteDealRenew``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteDealRenew`: DealResponse
    fmt.Fprintf(os.Stdout, "Response from `RenewsApi.DeleteDealRenew`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | 取引ID | 
**renewId** | **int32** | +更新ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDealRenewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **companyId** | **int32** | 事業所ID | 

### Return type

[**DealResponse**](dealResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDealRenew

> DealResponse UpdateDealRenew(ctx, id, renewId).RenewUpdateParams(renewUpdateParams).Execute()

取引（収入／支出）の+更新の更新



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
    id := int32(56) // int32 | 取引ID
    renewId := int32(56) // int32 | +更新ID
    renewUpdateParams := *openapiclient.NewRenewUpdateParams(int32(1), "2014-01-01", []openapiclient.RenewUpdateParamsDetails{*openapiclient.NewRenewUpdateParamsDetails(int32(1), int32(1), int32(1080))}) // RenewUpdateParams | +更新の更新情報

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RenewsApi.UpdateDealRenew(context.Background(), id, renewId).RenewUpdateParams(renewUpdateParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RenewsApi.UpdateDealRenew``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDealRenew`: DealResponse
    fmt.Fprintf(os.Stdout, "Response from `RenewsApi.UpdateDealRenew`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | 取引ID | 
**renewId** | **int32** | +更新ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDealRenewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **renewUpdateParams** | [**RenewUpdateParams**](RenewUpdateParams.md) | +更新の更新情報 | 

### Return type

[**DealResponse**](dealResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

