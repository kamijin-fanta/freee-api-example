# \WalletTxnsApi

All URIs are relative to *https://api.freee.co.jp*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWalletTxn**](WalletTxnsApi.md#CreateWalletTxn) | **Post** /api/1/wallet_txns | 明細の作成
[**DestroyWalletTxn**](WalletTxnsApi.md#DestroyWalletTxn) | **Delete** /api/1/wallet_txns/{id} | 明細の削除
[**GetWalletTxn**](WalletTxnsApi.md#GetWalletTxn) | **Get** /api/1/wallet_txns/{id} | 明細の取得
[**GetWalletTxns**](WalletTxnsApi.md#GetWalletTxns) | **Get** /api/1/wallet_txns | 明細一覧の取得



## CreateWalletTxn

> WalletTxnResponse CreateWalletTxn(ctx).WalletTxnParams(walletTxnParams).Execute()

明細の作成



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
    walletTxnParams := *openapiclient.NewWalletTxnParams("income", int32(5000), int32(1), "bank_account", "2018-01-01", int32(1)) // WalletTxnParams | 明細の作成 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WalletTxnsApi.CreateWalletTxn(context.Background()).WalletTxnParams(walletTxnParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletTxnsApi.CreateWalletTxn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWalletTxn`: WalletTxnResponse
    fmt.Fprintf(os.Stdout, "Response from `WalletTxnsApi.CreateWalletTxn`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWalletTxnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **walletTxnParams** | [**WalletTxnParams**](WalletTxnParams.md) | 明細の作成 | 

### Return type

[**WalletTxnResponse**](walletTxnResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyWalletTxn

> DestroyWalletTxn(ctx, id).CompanyId(companyId).Execute()

明細の削除



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
    id := int32(56) // int32 | 明細ID
    companyId := int32(56) // int32 | 事業所ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WalletTxnsApi.DestroyWalletTxn(context.Background(), id).CompanyId(companyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletTxnsApi.DestroyWalletTxn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | 明細ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDestroyWalletTxnRequest struct via the builder pattern


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


## GetWalletTxn

> WalletTxnResponse GetWalletTxn(ctx, id).CompanyId(companyId).Execute()

明細の取得



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
    id := int32(56) // int32 | 明細ID
    companyId := int32(56) // int32 | 事業所ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WalletTxnsApi.GetWalletTxn(context.Background(), id).CompanyId(companyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletTxnsApi.GetWalletTxn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWalletTxn`: WalletTxnResponse
    fmt.Fprintf(os.Stdout, "Response from `WalletTxnsApi.GetWalletTxn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | 明細ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWalletTxnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **companyId** | **int32** | 事業所ID | 

### Return type

[**WalletTxnResponse**](walletTxnResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWalletTxns

> InlineResponse20013 GetWalletTxns(ctx).CompanyId(companyId).WalletableType(walletableType).WalletableId(walletableId).StartDate(startDate).EndDate(endDate).EntrySide(entrySide).Offset(offset).Limit(limit).Execute()

明細一覧の取得



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
    walletableType := "walletableType_example" // string | 口座区分 (銀行口座: bank_account, クレジットカード: credit_card, 現金: wallet) walletable_type、walletable_idは同時に指定が必要です。 (optional)
    walletableId := int32(56) // int32 | 口座ID walletable_type、walletable_idは同時に指定が必要です。 (optional)
    startDate := "startDate_example" // string | 取引日で絞込：開始日 (yyyy-mm-dd) (optional)
    endDate := "endDate_example" // string | 取引日で絞込：終了日 (yyyy-mm-dd) (optional)
    entrySide := "entrySide_example" // string | 入金／出金 (入金: income, 出金: expense) (optional)
    offset := int32(56) // int32 | 取得レコードのオフセット (デフォルト: 0) (optional)
    limit := int32(56) // int32 | 取得レコードの件数 (デフォルト: 20, 最小: 1, 最大: 100)  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WalletTxnsApi.GetWalletTxns(context.Background()).CompanyId(companyId).WalletableType(walletableType).WalletableId(walletableId).StartDate(startDate).EndDate(endDate).EntrySide(entrySide).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletTxnsApi.GetWalletTxns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWalletTxns`: InlineResponse20013
    fmt.Fprintf(os.Stdout, "Response from `WalletTxnsApi.GetWalletTxns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWalletTxnsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **int32** | 事業所ID | 
 **walletableType** | **string** | 口座区分 (銀行口座: bank_account, クレジットカード: credit_card, 現金: wallet) walletable_type、walletable_idは同時に指定が必要です。 | 
 **walletableId** | **int32** | 口座ID walletable_type、walletable_idは同時に指定が必要です。 | 
 **startDate** | **string** | 取引日で絞込：開始日 (yyyy-mm-dd) | 
 **endDate** | **string** | 取引日で絞込：終了日 (yyyy-mm-dd) | 
 **entrySide** | **string** | 入金／出金 (入金: income, 出金: expense) | 
 **offset** | **int32** | 取得レコードのオフセット (デフォルト: 0) | 
 **limit** | **int32** | 取得レコードの件数 (デフォルト: 20, 最小: 1, 最大: 100)  | 

### Return type

[**InlineResponse20013**](inline_response_200_13.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

