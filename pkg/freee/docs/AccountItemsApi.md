# \AccountItemsApi

All URIs are relative to *https://api.freee.co.jp*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccountItem**](AccountItemsApi.md#CreateAccountItem) | **Post** /api/1/account_items | 勘定科目の作成
[**DestroyAccountItem**](AccountItemsApi.md#DestroyAccountItem) | **Delete** /api/1/account_items/{id} | 勘定科目の削除
[**GetAccountItem**](AccountItemsApi.md#GetAccountItem) | **Get** /api/1/account_items/{id} | 勘定科目の詳細情報の取得
[**GetAccountItems**](AccountItemsApi.md#GetAccountItems) | **Get** /api/1/account_items | 勘定科目一覧の取得
[**UpdateAccountItem**](AccountItemsApi.md#UpdateAccountItem) | **Put** /api/1/account_items/{id} | 勘定科目の更新



## CreateAccountItem

> AccountItemResponse CreateAccountItem(ctx).AccountItemParams(accountItemParams).Execute()

勘定科目の作成



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
    accountItemParams := *openapiclient.NewAccountItemParams(int32(1), *openapiclient.NewAccountItemParamsAccountItem("新しい勘定科目", int32(1), "その他預金", int32(1), int32(1), int32(1))) // AccountItemParams | 勘定科目の作成

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountItemsApi.CreateAccountItem(context.Background()).AccountItemParams(accountItemParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountItemsApi.CreateAccountItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAccountItem`: AccountItemResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountItemsApi.CreateAccountItem`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccountItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountItemParams** | [**AccountItemParams**](AccountItemParams.md) | 勘定科目の作成 | 

### Return type

[**AccountItemResponse**](accountItemResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyAccountItem

> DestroyAccountItem(ctx, id).CompanyId(companyId).Execute()

勘定科目の削除



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
    id := int32(56) // int32 | 
    companyId := int32(56) // int32 | 事業所ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountItemsApi.DestroyAccountItem(context.Background(), id).CompanyId(companyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountItemsApi.DestroyAccountItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDestroyAccountItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **companyId** | **int32** | 事業所ID | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountItem

> AccountItemResponse GetAccountItem(ctx, id).CompanyId(companyId).Execute()

勘定科目の詳細情報の取得



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
    companyId := int32(56) // int32 | 事業所ID
    id := int32(56) // int32 | 勘定科目ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountItemsApi.GetAccountItem(context.Background(), id).CompanyId(companyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountItemsApi.GetAccountItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccountItem`: AccountItemResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountItemsApi.GetAccountItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | 勘定科目ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **int32** | 事業所ID | 


### Return type

[**AccountItemResponse**](accountItemResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountItems

> AccountItemsResponse GetAccountItems(ctx).CompanyId(companyId).BaseDate(baseDate).Execute()

勘定科目一覧の取得



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
    companyId := int32(56) // int32 | 事業所ID
    baseDate := "baseDate_example" // string | 基準日:指定した場合、勘定科目に紐づく税区分(default_tax_code)が、基準日の税率に基づいて返ります。 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountItemsApi.GetAccountItems(context.Background()).CompanyId(companyId).BaseDate(baseDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountItemsApi.GetAccountItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccountItems`: AccountItemsResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountItemsApi.GetAccountItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **int32** | 事業所ID | 
 **baseDate** | **string** | 基準日:指定した場合、勘定科目に紐づく税区分(default_tax_code)が、基準日の税率に基づいて返ります。 | 

### Return type

[**AccountItemsResponse**](accountItemsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccountItem

> AccountItemResponse UpdateAccountItem(ctx, id).AccountItemParams(accountItemParams).Execute()

勘定科目の更新



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
    id := int32(56) // int32 | 
    accountItemParams := *openapiclient.NewAccountItemParams(int32(1), *openapiclient.NewAccountItemParamsAccountItem("新しい勘定科目", int32(1), "その他預金", int32(1), int32(1), int32(1))) // AccountItemParams | 勘定科目の更新

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountItemsApi.UpdateAccountItem(context.Background(), id).AccountItemParams(accountItemParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountItemsApi.UpdateAccountItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAccountItem`: AccountItemResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountItemsApi.UpdateAccountItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccountItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountItemParams** | [**AccountItemParams**](AccountItemParams.md) | 勘定科目の更新 | 

### Return type

[**AccountItemResponse**](accountItemResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

