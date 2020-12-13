# \JournalsApi

All URIs are relative to *https://api.freee.co.jp*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DownloadJournal**](JournalsApi.md#DownloadJournal) | **Get** /api/1/journals/reports/{id}/download | ダウンロード実行
[**GetJournalStatus**](JournalsApi.md#GetJournalStatus) | **Get** /api/1/journals/reports/{id}/status | ステータス確認
[**GetJournals**](JournalsApi.md#GetJournals) | **Get** /api/1/journals | ダウンロード要求



## DownloadJournal

> string DownloadJournal(ctx, id).CompanyId(companyId).Execute()

ダウンロード実行



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
    id := int32(56) // int32 | 受け付けID
    companyId := int32(56) // int32 | 事業所ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.JournalsApi.DownloadJournal(context.Background(), id).CompanyId(companyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JournalsApi.DownloadJournal``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadJournal`: string
    fmt.Fprintf(os.Stdout, "Response from `JournalsApi.DownloadJournal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | 受け付けID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadJournalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **companyId** | **int32** | 事業所ID | 

### Return type

**string**

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/csv, application/pdf, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJournalStatus

> JournalStatusResponse GetJournalStatus(ctx, id).CompanyId(companyId).Execute()

ステータス確認



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
    id := int32(56) // int32 | 受け付けID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.JournalsApi.GetJournalStatus(context.Background(), id).CompanyId(companyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JournalsApi.GetJournalStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetJournalStatus`: JournalStatusResponse
    fmt.Fprintf(os.Stdout, "Response from `JournalsApi.GetJournalStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | 受け付けID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJournalStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **int32** | 事業所ID | 


### Return type

[**JournalStatusResponse**](journalStatusResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJournals

> JournalsResponse GetJournals(ctx).DownloadType(downloadType).CompanyId(companyId).VisibleTags(visibleTags).StartDate(startDate).EndDate(endDate).Execute()

ダウンロード要求



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
    downloadType := "downloadType_example" // string | ダウンロード形式
    companyId := int32(56) // int32 | 事業所ID
    visibleTags := []string{"VisibleTags_example"} // []string | 補助科目やコメントとして出力する項目 (optional)
    startDate := "startDate_example" // string | 取得開始日 (yyyy-mm-dd) (optional)
    endDate := "endDate_example" // string | 取得終了日 (yyyy-mm-dd) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.JournalsApi.GetJournals(context.Background()).DownloadType(downloadType).CompanyId(companyId).VisibleTags(visibleTags).StartDate(startDate).EndDate(endDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JournalsApi.GetJournals``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetJournals`: JournalsResponse
    fmt.Fprintf(os.Stdout, "Response from `JournalsApi.GetJournals`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetJournalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **downloadType** | **string** | ダウンロード形式 | 
 **companyId** | **int32** | 事業所ID | 
 **visibleTags** | **[]string** | 補助科目やコメントとして出力する項目 | 
 **startDate** | **string** | 取得開始日 (yyyy-mm-dd) | 
 **endDate** | **string** | 取得終了日 (yyyy-mm-dd) | 

### Return type

[**JournalsResponse**](journalsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

