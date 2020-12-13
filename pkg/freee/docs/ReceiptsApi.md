# \ReceiptsApi

All URIs are relative to *https://api.freee.co.jp*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateReceipt**](ReceiptsApi.md#CreateReceipt) | **Post** /api/1/receipts | ファイルボックス 証憑ファイルアップロード
[**DestroyReceipt**](ReceiptsApi.md#DestroyReceipt) | **Delete** /api/1/receipts/{id} | ファイルボックス 証憑ファイルを削除する
[**DownloadReceipt**](ReceiptsApi.md#DownloadReceipt) | **Get** /api/1/receipts/{id}/download | ファイルボックス 証憑ファイルのダウンロード
[**GetReceipt**](ReceiptsApi.md#GetReceipt) | **Get** /api/1/receipts/{id} | ファイルボックス 証憑ファイルの取得
[**GetReceipts**](ReceiptsApi.md#GetReceipts) | **Get** /api/1/receipts | ファイルボックス 証憑ファイル一覧の取得
[**UpdateReceipt**](ReceiptsApi.md#UpdateReceipt) | **Put** /api/1/receipts/{id} | ファイルボックス 証憑ファイル情報更新



## CreateReceipt

> ReceiptResponse CreateReceipt(ctx).CompanyId(companyId).Receipt(receipt).Description(description).IssueDate(issueDate).Execute()

ファイルボックス 証憑ファイルアップロード



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
    receipt := os.NewFile(1234, "some_file") // *os.File | 証憑ファイル
    description := "description_example" // string | メモ (255文字以内) (optional)
    issueDate := "issueDate_example" // string | 取引日 (yyyy-mm-dd) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReceiptsApi.CreateReceipt(context.Background()).CompanyId(companyId).Receipt(receipt).Description(description).IssueDate(issueDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReceiptsApi.CreateReceipt``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateReceipt`: ReceiptResponse
    fmt.Fprintf(os.Stdout, "Response from `ReceiptsApi.CreateReceipt`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateReceiptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **int32** | 事業所ID | 
 **receipt** | ***os.File** | 証憑ファイル | 
 **description** | **string** | メモ (255文字以内) | 
 **issueDate** | **string** | 取引日 (yyyy-mm-dd) | 

### Return type

[**ReceiptResponse**](receiptResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyReceipt

> DestroyReceipt(ctx, id).CompanyId(companyId).Execute()

ファイルボックス 証憑ファイルを削除する



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
    id := int32(56) // int32 | 証憑ID
    companyId := int32(56) // int32 | 事業所ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReceiptsApi.DestroyReceipt(context.Background(), id).CompanyId(companyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReceiptsApi.DestroyReceipt``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | 証憑ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDestroyReceiptRequest struct via the builder pattern


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


## DownloadReceipt

> string DownloadReceipt(ctx, id).CompanyId(companyId).Execute()

ファイルボックス 証憑ファイルのダウンロード



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
    id := int32(56) // int32 | 証憑ID
    companyId := int32(56) // int32 | 事業所ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReceiptsApi.DownloadReceipt(context.Background(), id).CompanyId(companyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReceiptsApi.DownloadReceipt``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadReceipt`: string
    fmt.Fprintf(os.Stdout, "Response from `ReceiptsApi.DownloadReceipt`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | 証憑ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadReceiptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **companyId** | **int32** | 事業所ID | 

### Return type

**string**

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/csv, application/pdf, image/_*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReceipt

> ReceiptResponse GetReceipt(ctx, id).CompanyId(companyId).Execute()

ファイルボックス 証憑ファイルの取得



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
    id := int32(56) // int32 | 証憑ID
    companyId := int32(56) // int32 | 事業所ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReceiptsApi.GetReceipt(context.Background(), id).CompanyId(companyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReceiptsApi.GetReceipt``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReceipt`: ReceiptResponse
    fmt.Fprintf(os.Stdout, "Response from `ReceiptsApi.GetReceipt`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | 証憑ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReceiptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **companyId** | **int32** | 事業所ID | 

### Return type

[**ReceiptResponse**](receiptResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReceipts

> InlineResponse20014 GetReceipts(ctx).CompanyId(companyId).StartDate(startDate).EndDate(endDate).UserName(userName).Number(number).CommentType(commentType).CommentImportant(commentImportant).Category(category).Offset(offset).Limit(limit).Execute()

ファイルボックス 証憑ファイル一覧の取得



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
    startDate := "startDate_example" // string | アップロード日 (yyyy-mm-dd)
    endDate := "endDate_example" // string | アップロード日 (yyyy-mm-dd)
    userName := "userName_example" // string | アップロードしたユーザー名、メールアドレス (optional)
    number := int32(56) // int32 | アップロードファイルNo (optional)
    commentType := "commentType_example" // string | posted:コメントあり, raised:未解決, resolved:解決済 (optional)
    commentImportant := true // bool | trueの時、重要コメント付きが対象 (optional)
    category := "category_example" // string | all:すべて、without_deal:未登録、with_expense_application_line:経費申請中, with_deal:登録済み、ignored:無視 (optional)
    offset := int32(56) // int32 | 取得レコードのオフセット (デフォルト: 0) (optional)
    limit := int32(56) // int32 | 取得レコードの件数 (デフォルト: 50, 最小: 1, 最大: 3000) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReceiptsApi.GetReceipts(context.Background()).CompanyId(companyId).StartDate(startDate).EndDate(endDate).UserName(userName).Number(number).CommentType(commentType).CommentImportant(commentImportant).Category(category).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReceiptsApi.GetReceipts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReceipts`: InlineResponse20014
    fmt.Fprintf(os.Stdout, "Response from `ReceiptsApi.GetReceipts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetReceiptsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **int32** | 事業所ID | 
 **startDate** | **string** | アップロード日 (yyyy-mm-dd) | 
 **endDate** | **string** | アップロード日 (yyyy-mm-dd) | 
 **userName** | **string** | アップロードしたユーザー名、メールアドレス | 
 **number** | **int32** | アップロードファイルNo | 
 **commentType** | **string** | posted:コメントあり, raised:未解決, resolved:解決済 | 
 **commentImportant** | **bool** | trueの時、重要コメント付きが対象 | 
 **category** | **string** | all:すべて、without_deal:未登録、with_expense_application_line:経費申請中, with_deal:登録済み、ignored:無視 | 
 **offset** | **int32** | 取得レコードのオフセット (デフォルト: 0) | 
 **limit** | **int32** | 取得レコードの件数 (デフォルト: 50, 最小: 1, 最大: 3000) | 

### Return type

[**InlineResponse20014**](inline_response_200_14.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateReceipt

> ReceiptResponse UpdateReceipt(ctx, id).ReceiptUpdateParams(receiptUpdateParams).Execute()

ファイルボックス 証憑ファイル情報更新



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
    id := int32(56) // int32 | 証憑ID
    receiptUpdateParams := *openapiclient.NewReceiptUpdateParams(int32(1), "2018-01-01") // ReceiptUpdateParams | 経費申請の更新

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReceiptsApi.UpdateReceipt(context.Background(), id).ReceiptUpdateParams(receiptUpdateParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReceiptsApi.UpdateReceipt``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateReceipt`: ReceiptResponse
    fmt.Fprintf(os.Stdout, "Response from `ReceiptsApi.UpdateReceipt`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | 証憑ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateReceiptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **receiptUpdateParams** | [**ReceiptUpdateParams**](ReceiptUpdateParams.md) | 経費申請の更新 | 

### Return type

[**ReceiptResponse**](receiptResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

