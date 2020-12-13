# \WalletablesApi

All URIs are relative to *https://api.freee.co.jp*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWalletable**](WalletablesApi.md#CreateWalletable) | **Post** /api/1/walletables | 口座の作成
[**DestroyWalletable**](WalletablesApi.md#DestroyWalletable) | **Delete** /api/1/walletables/{type}/{id} | 口座の削除
[**GetWalletable**](WalletablesApi.md#GetWalletable) | **Get** /api/1/walletables/{type}/{id} | 口座情報の取得
[**GetWalletables**](WalletablesApi.md#GetWalletables) | **Get** /api/1/walletables | 口座一覧の取得
[**UpdateWalletable**](WalletablesApi.md#UpdateWalletable) | **Put** /api/1/walletables/{type}/{id} | 口座の更新



## CreateWalletable

> WalletableCreateResponse CreateWalletable(ctx).WalletableCreateParams(walletableCreateParams).Execute()

口座の作成



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
    walletableCreateParams := *openapiclient.NewWalletableCreateParams("ＸＸ銀行", "bank_account", int32(1)) // WalletableCreateParams | 口座の作成 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WalletablesApi.CreateWalletable(context.Background()).WalletableCreateParams(walletableCreateParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletablesApi.CreateWalletable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWalletable`: WalletableCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `WalletablesApi.CreateWalletable`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWalletableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **walletableCreateParams** | [**WalletableCreateParams**](WalletableCreateParams.md) | 口座の作成 | 

### Return type

[**WalletableCreateResponse**](walletableCreateResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyWalletable

> DestroyWalletable(ctx, id, type_).CompanyId(companyId).Execute()

口座の削除



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
    id := int32(56) // int32 | 口座ID
    type_ := "type__example" // string | 口座種別（bank_account : 銀行口座, credit_card : クレジットカード, wallet : その他の決済口座）
    companyId := int32(56) // int32 | 事業所ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WalletablesApi.DestroyWalletable(context.Background(), id, type_).CompanyId(companyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletablesApi.DestroyWalletable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | 口座ID | 
**type_** | **string** | 口座種別（bank_account : 銀行口座, credit_card : クレジットカード, wallet : その他の決済口座） | 

### Other Parameters

Other parameters are passed through a pointer to a apiDestroyWalletableRequest struct via the builder pattern


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


## GetWalletable

> InlineResponse20010 GetWalletable(ctx, id, type_).CompanyId(companyId).Execute()

口座情報の取得



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
    id := int32(56) // int32 | 口座ID
    type_ := "type__example" // string | 口座種別（bank_account : 銀行口座, credit_card : クレジットカード, wallet : その他の決済口座）
    companyId := int32(56) // int32 | 事業所ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WalletablesApi.GetWalletable(context.Background(), id, type_).CompanyId(companyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletablesApi.GetWalletable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWalletable`: InlineResponse20010
    fmt.Fprintf(os.Stdout, "Response from `WalletablesApi.GetWalletable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | 口座ID | 
**type_** | **string** | 口座種別（bank_account : 銀行口座, credit_card : クレジットカード, wallet : その他の決済口座） | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWalletableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **companyId** | **int32** | 事業所ID | 

### Return type

[**InlineResponse20010**](inline_response_200_10.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWalletables

> InlineResponse2009 GetWalletables(ctx).CompanyId(companyId).WithBalance(withBalance).Type_(type_).Execute()

口座一覧の取得



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
    withBalance := true // bool | 残高情報を含める (optional)
    type_ := "type__example" // string | 口座種別（bank_account : 銀行口座, credit_card : クレジットカード, wallet : その他の決済口座） (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WalletablesApi.GetWalletables(context.Background()).CompanyId(companyId).WithBalance(withBalance).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletablesApi.GetWalletables``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWalletables`: InlineResponse2009
    fmt.Fprintf(os.Stdout, "Response from `WalletablesApi.GetWalletables`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWalletablesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **int32** | 事業所ID | 
 **withBalance** | **bool** | 残高情報を含める | 
 **type_** | **string** | 口座種別（bank_account : 銀行口座, credit_card : クレジットカード, wallet : その他の決済口座） | 

### Return type

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWalletable

> InlineResponse20010 UpdateWalletable(ctx, id, type_).WalletableUpdateParams(walletableUpdateParams).Execute()

口座の更新



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
    type_ := "type__example" // string | 口座種別（bank_account : 銀行口座, credit_card : クレジットカード, wallet : その他の決済口座）
    walletableUpdateParams := *openapiclient.NewWalletableUpdateParams("ＸＸ銀行", int32(1)) // WalletableUpdateParams | 口座の作成 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WalletablesApi.UpdateWalletable(context.Background(), id, type_).WalletableUpdateParams(walletableUpdateParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletablesApi.UpdateWalletable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWalletable`: InlineResponse20010
    fmt.Fprintf(os.Stdout, "Response from `WalletablesApi.UpdateWalletable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 
**type_** | **string** | 口座種別（bank_account : 銀行口座, credit_card : クレジットカード, wallet : その他の決済口座） | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWalletableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **walletableUpdateParams** | [**WalletableUpdateParams**](WalletableUpdateParams.md) | 口座の作成 | 

### Return type

[**InlineResponse20010**](inline_response_200_10.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

