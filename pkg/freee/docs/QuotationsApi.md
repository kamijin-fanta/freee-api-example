# \QuotationsApi

All URIs are relative to *https://api.freee.co.jp*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateQuotation**](QuotationsApi.md#CreateQuotation) | **Post** /api/1/quotations | 見積書の作成
[**DestroyQuotation**](QuotationsApi.md#DestroyQuotation) | **Delete** /api/1/quotations/{id} | 見積書の削除
[**GetQuotation**](QuotationsApi.md#GetQuotation) | **Get** /api/1/quotations/{id} | 見積書の取得
[**GetQuotations**](QuotationsApi.md#GetQuotations) | **Get** /api/1/quotations | 見積書一覧の取得
[**UpdateQuotation**](QuotationsApi.md#UpdateQuotation) | **Put** /api/1/quotations/{id} | 見積書の更新



## CreateQuotation

> QuotationResponse CreateQuotation(ctx).QuotationCreateParams(quotationCreateParams).Execute()

見積書の作成



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
    quotationCreateParams := *openapiclient.NewQuotationCreateParams(int32(1), "freeeパートナー株式会社", "御中") // QuotationCreateParams | 見積書の作成 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.QuotationsApi.CreateQuotation(context.Background()).QuotationCreateParams(quotationCreateParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QuotationsApi.CreateQuotation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateQuotation`: QuotationResponse
    fmt.Fprintf(os.Stdout, "Response from `QuotationsApi.CreateQuotation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateQuotationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **quotationCreateParams** | [**QuotationCreateParams**](QuotationCreateParams.md) | 見積書の作成 | 

### Return type

[**QuotationResponse**](quotationResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyQuotation

> DestroyQuotation(ctx, id).CompanyId(companyId).Execute()

見積書の削除



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
    resp, r, err := api_client.QuotationsApi.DestroyQuotation(context.Background(), id).CompanyId(companyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QuotationsApi.DestroyQuotation``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDestroyQuotationRequest struct via the builder pattern


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


## GetQuotation

> QuotationResponse GetQuotation(ctx, id).CompanyId(companyId).Execute()

見積書の取得



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
    id := int32(56) // int32 | 見積書ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.QuotationsApi.GetQuotation(context.Background(), id).CompanyId(companyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QuotationsApi.GetQuotation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetQuotation`: QuotationResponse
    fmt.Fprintf(os.Stdout, "Response from `QuotationsApi.GetQuotation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | 見積書ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetQuotationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **int32** | 事業所ID | 


### Return type

[**QuotationResponse**](quotationResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQuotations

> QuotationIndexResponse GetQuotations(ctx).CompanyId(companyId).PartnerId(partnerId).PartnerCode(partnerCode).StartIssueDate(startIssueDate).EndIssueDate(endIssueDate).QuotationNumber(quotationNumber).Description(description).QuotationStatus(quotationStatus).Offset(offset).Limit(limit).Execute()

見積書一覧の取得



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
    startIssueDate := "startIssueDate_example" // string | 見積日の開始日(yyyy-mm-dd) (optional)
    endIssueDate := "endIssueDate_example" // string | 見積日の終了日(yyyy-mm-dd) (optional)
    quotationNumber := "quotationNumber_example" // string | 見積書番号 (optional)
    description := "description_example" // string | 概要 (optional)
    quotationStatus := "quotationStatus_example" // string | 見積書ステータス  (unsubmitted: 送付待ち, submitted: 送付済み, all: 全て) (optional)
    offset := int32(56) // int32 | 取得レコードのオフセット (デフォルト: 0) (optional)
    limit := int32(56) // int32 | 取得レコードの件数 (デフォルト: 20, 最大: 100)  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.QuotationsApi.GetQuotations(context.Background()).CompanyId(companyId).PartnerId(partnerId).PartnerCode(partnerCode).StartIssueDate(startIssueDate).EndIssueDate(endIssueDate).QuotationNumber(quotationNumber).Description(description).QuotationStatus(quotationStatus).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QuotationsApi.GetQuotations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetQuotations`: QuotationIndexResponse
    fmt.Fprintf(os.Stdout, "Response from `QuotationsApi.GetQuotations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetQuotationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **int32** | 事業所ID | 
 **partnerId** | **int32** | 取引先IDで絞込 | 
 **partnerCode** | **string** | 取引先コードで絞込 | 
 **startIssueDate** | **string** | 見積日の開始日(yyyy-mm-dd) | 
 **endIssueDate** | **string** | 見積日の終了日(yyyy-mm-dd) | 
 **quotationNumber** | **string** | 見積書番号 | 
 **description** | **string** | 概要 | 
 **quotationStatus** | **string** | 見積書ステータス  (unsubmitted: 送付待ち, submitted: 送付済み, all: 全て) | 
 **offset** | **int32** | 取得レコードのオフセット (デフォルト: 0) | 
 **limit** | **int32** | 取得レコードの件数 (デフォルト: 20, 最大: 100)  | 

### Return type

[**QuotationIndexResponse**](quotationIndexResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateQuotation

> QuotationResponse UpdateQuotation(ctx, id).QuotationUpdateParams(quotationUpdateParams).Execute()

見積書の更新



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
    id := int32(56) // int32 | 見積書ID
    quotationUpdateParams := *openapiclient.NewQuotationUpdateParams(int32(1), "freeeパートナー株式会社", "御中") // QuotationUpdateParams | 見積書の更新 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.QuotationsApi.UpdateQuotation(context.Background(), id).QuotationUpdateParams(quotationUpdateParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QuotationsApi.UpdateQuotation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateQuotation`: QuotationResponse
    fmt.Fprintf(os.Stdout, "Response from `QuotationsApi.UpdateQuotation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | 見積書ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateQuotationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **quotationUpdateParams** | [**QuotationUpdateParams**](QuotationUpdateParams.md) | 見積書の更新 | 

### Return type

[**QuotationResponse**](quotationResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

