# \InvoicesApi

All URIs are relative to *https://api.freee.co.jp*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateInvoice**](InvoicesApi.md#CreateInvoice) | **Post** /api/1/invoices | 請求書の作成
[**DestroyInvoice**](InvoicesApi.md#DestroyInvoice) | **Delete** /api/1/invoices/{id} | 請求書の削除
[**GetInvoice**](InvoicesApi.md#GetInvoice) | **Get** /api/1/invoices/{id} | 請求書の取得
[**GetInvoices**](InvoicesApi.md#GetInvoices) | **Get** /api/1/invoices | 請求書一覧の取得
[**UpdateInvoice**](InvoicesApi.md#UpdateInvoice) | **Put** /api/1/invoices/{id} | 請求書の更新



## CreateInvoice

> InvoiceResponse CreateInvoice(ctx).InvoiceCreateParams(invoiceCreateParams).Execute()

請求書の作成



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
    invoiceCreateParams := *openapiclient.NewInvoiceCreateParams(int32(1), "株式会社freeeパートナー", "御中") // InvoiceCreateParams | 請求書の作成 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InvoicesApi.CreateInvoice(context.Background()).InvoiceCreateParams(invoiceCreateParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoicesApi.CreateInvoice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateInvoice`: InvoiceResponse
    fmt.Fprintf(os.Stdout, "Response from `InvoicesApi.CreateInvoice`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **invoiceCreateParams** | [**InvoiceCreateParams**](InvoiceCreateParams.md) | 請求書の作成 | 

### Return type

[**InvoiceResponse**](invoiceResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyInvoice

> DestroyInvoice(ctx, id).CompanyId(companyId).Execute()

請求書の削除



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
    resp, r, err := api_client.InvoicesApi.DestroyInvoice(context.Background(), id).CompanyId(companyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoicesApi.DestroyInvoice``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDestroyInvoiceRequest struct via the builder pattern


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


## GetInvoice

> InvoiceResponse GetInvoice(ctx, id).CompanyId(companyId).Execute()

請求書の取得



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
    id := int32(56) // int32 | 請求書ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InvoicesApi.GetInvoice(context.Background(), id).CompanyId(companyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoicesApi.GetInvoice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInvoice`: InvoiceResponse
    fmt.Fprintf(os.Stdout, "Response from `InvoicesApi.GetInvoice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | 請求書ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **int32** | 事業所ID | 


### Return type

[**InvoiceResponse**](invoiceResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInvoices

> InvoiceIndexResponse GetInvoices(ctx).CompanyId(companyId).PartnerId(partnerId).PartnerCode(partnerCode).StartIssueDate(startIssueDate).EndIssueDate(endIssueDate).StartDueDate(startDueDate).EndDueDate(endDueDate).InvoiceNumber(invoiceNumber).Description(description).InvoiceStatus(invoiceStatus).PaymentStatus(paymentStatus).Offset(offset).Limit(limit).Execute()

請求書一覧の取得



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
    partnerId := int32(56) // int32 | 取引先IDで絞込 (optional)
    partnerCode := "partnerCode_example" // string | 取引先コードで絞込 (optional)
    startIssueDate := "startIssueDate_example" // string | 請求日の開始日(yyyy-mm-dd) (optional)
    endIssueDate := "endIssueDate_example" // string | 請求日の終了日(yyyy-mm-dd) (optional)
    startDueDate := "startDueDate_example" // string | 期日の開始日(yyyy-mm-dd) (optional)
    endDueDate := "endDueDate_example" // string | 期日の終了日(yyyy-mm-dd) (optional)
    invoiceNumber := "invoiceNumber_example" // string | 請求書番号 (optional)
    description := "description_example" // string | 概要 (optional)
    invoiceStatus := "invoiceStatus_example" // string | 請求書ステータス  (draft: 下書き, applying: 申請中, remanded: 差し戻し, rejected: 却下, approved: 承認済み, unsubmitted: 送付待ち, submitted: 送付済み) (optional)
    paymentStatus := "paymentStatus_example" // string | 入金ステータス  (unsettled: 入金待ち, settled: 入金済み) (optional)
    offset := int32(56) // int32 | 取得レコードのオフセット (デフォルト: 0) (optional)
    limit := int32(56) // int32 | 取得レコードの件数 (デフォルト: 20, 最大: 100)  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InvoicesApi.GetInvoices(context.Background()).CompanyId(companyId).PartnerId(partnerId).PartnerCode(partnerCode).StartIssueDate(startIssueDate).EndIssueDate(endIssueDate).StartDueDate(startDueDate).EndDueDate(endDueDate).InvoiceNumber(invoiceNumber).Description(description).InvoiceStatus(invoiceStatus).PaymentStatus(paymentStatus).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoicesApi.GetInvoices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInvoices`: InvoiceIndexResponse
    fmt.Fprintf(os.Stdout, "Response from `InvoicesApi.GetInvoices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInvoicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **int32** | 事業所ID | 
 **partnerId** | **int32** | 取引先IDで絞込 | 
 **partnerCode** | **string** | 取引先コードで絞込 | 
 **startIssueDate** | **string** | 請求日の開始日(yyyy-mm-dd) | 
 **endIssueDate** | **string** | 請求日の終了日(yyyy-mm-dd) | 
 **startDueDate** | **string** | 期日の開始日(yyyy-mm-dd) | 
 **endDueDate** | **string** | 期日の終了日(yyyy-mm-dd) | 
 **invoiceNumber** | **string** | 請求書番号 | 
 **description** | **string** | 概要 | 
 **invoiceStatus** | **string** | 請求書ステータス  (draft: 下書き, applying: 申請中, remanded: 差し戻し, rejected: 却下, approved: 承認済み, unsubmitted: 送付待ち, submitted: 送付済み) | 
 **paymentStatus** | **string** | 入金ステータス  (unsettled: 入金待ち, settled: 入金済み) | 
 **offset** | **int32** | 取得レコードのオフセット (デフォルト: 0) | 
 **limit** | **int32** | 取得レコードの件数 (デフォルト: 20, 最大: 100)  | 

### Return type

[**InvoiceIndexResponse**](invoiceIndexResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInvoice

> InvoiceResponse UpdateInvoice(ctx, id).InvoiceUpdateParams(invoiceUpdateParams).Execute()

請求書の更新



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
    id := int32(56) // int32 | 請求書ID
    invoiceUpdateParams := *openapiclient.NewInvoiceUpdateParams(int32(1), "株式会社freeeパートナー", "御中") // InvoiceUpdateParams | 請求書の更新 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InvoicesApi.UpdateInvoice(context.Background(), id).InvoiceUpdateParams(invoiceUpdateParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoicesApi.UpdateInvoice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateInvoice`: InvoiceResponse
    fmt.Fprintf(os.Stdout, "Response from `InvoicesApi.UpdateInvoice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | 請求書ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **invoiceUpdateParams** | [**InvoiceUpdateParams**](InvoiceUpdateParams.md) | 請求書の更新 | 

### Return type

[**InvoiceResponse**](invoiceResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

