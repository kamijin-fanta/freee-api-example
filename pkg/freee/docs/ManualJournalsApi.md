# \ManualJournalsApi

All URIs are relative to *https://api.freee.co.jp*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateManualJournal**](ManualJournalsApi.md#CreateManualJournal) | **Post** /api/1/manual_journals | 振替伝票の作成
[**DestroyManualJournal**](ManualJournalsApi.md#DestroyManualJournal) | **Delete** /api/1/manual_journals/{id} | 振替伝票の削除
[**GetManualJournal**](ManualJournalsApi.md#GetManualJournal) | **Get** /api/1/manual_journals/{id} | 振替伝票の取得
[**GetManualJournals**](ManualJournalsApi.md#GetManualJournals) | **Get** /api/1/manual_journals | 振替伝票一覧の取得
[**UpdateManualJournal**](ManualJournalsApi.md#UpdateManualJournal) | **Put** /api/1/manual_journals/{id} | 振替伝票の更新



## CreateManualJournal

> ManualJournalResponse CreateManualJournal(ctx).ManualJournalCreateParams(manualJournalCreateParams).Execute()

振替伝票の作成



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
    manualJournalCreateParams := *openapiclient.NewManualJournalCreateParams(int32(1), "2013-01-01", []openapiclient.ManualJournalCreateParamsDetails{*openapiclient.NewManualJournalCreateParamsDetails("debit", int32(1), int32(1), int32(10800))}) // ManualJournalCreateParams | 振替伝票の作成 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManualJournalsApi.CreateManualJournal(context.Background()).ManualJournalCreateParams(manualJournalCreateParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManualJournalsApi.CreateManualJournal``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateManualJournal`: ManualJournalResponse
    fmt.Fprintf(os.Stdout, "Response from `ManualJournalsApi.CreateManualJournal`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateManualJournalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **manualJournalCreateParams** | [**ManualJournalCreateParams**](ManualJournalCreateParams.md) | 振替伝票の作成 | 

### Return type

[**ManualJournalResponse**](manualJournalResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyManualJournal

> DestroyManualJournal(ctx, id).CompanyId(companyId).Execute()

振替伝票の削除



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
    resp, r, err := api_client.ManualJournalsApi.DestroyManualJournal(context.Background(), id).CompanyId(companyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManualJournalsApi.DestroyManualJournal``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDestroyManualJournalRequest struct via the builder pattern


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


## GetManualJournal

> ManualJournalResponse GetManualJournal(ctx, id).CompanyId(companyId).Execute()

振替伝票の取得



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
    id := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManualJournalsApi.GetManualJournal(context.Background(), id).CompanyId(companyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManualJournalsApi.GetManualJournal``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetManualJournal`: ManualJournalResponse
    fmt.Fprintf(os.Stdout, "Response from `ManualJournalsApi.GetManualJournal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetManualJournalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **int32** | 事業所ID | 


### Return type

[**ManualJournalResponse**](manualJournalResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetManualJournals

> InlineResponse2003 GetManualJournals(ctx).CompanyId(companyId).StartIssueDate(startIssueDate).EndIssueDate(endIssueDate).EntrySide(entrySide).AccountItemId(accountItemId).MinAmount(minAmount).MaxAmount(maxAmount).PartnerId(partnerId).PartnerCode(partnerCode).ItemId(itemId).SectionId(sectionId).Segment1TagId(segment1TagId).Segment2TagId(segment2TagId).Segment3TagId(segment3TagId).CommentStatus(commentStatus).CommentImportant(commentImportant).Adjustment(adjustment).TxnNumber(txnNumber).Offset(offset).Limit(limit).Execute()

振替伝票一覧の取得



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
    startIssueDate := "startIssueDate_example" // string | 発生日で絞込：開始日(yyyy-mm-dd) (optional)
    endIssueDate := "endIssueDate_example" // string | 発生日で絞込：終了日(yyyy-mm-dd) (optional)
    entrySide := "entrySide_example" // string | 貸借で絞込 (貸方: credit, 借方: debit) (optional)
    accountItemId := int32(56) // int32 | 勘定科目IDで絞込 (optional)
    minAmount := int32(56) // int32 | 金額で絞込：下限 (optional)
    maxAmount := int32(56) // int32 | 金額で絞込：上限 (optional)
    partnerId := int32(56) // int32 | 取引先IDで絞込（0を指定すると、取引先が未選択の貸借行を絞り込めます） (optional)
    partnerCode := "partnerCode_example" // string | 取引先コードで絞込 (optional)
    itemId := int32(56) // int32 | 品目IDで絞込（0を指定すると、品目が未選択の貸借行を絞り込めます） (optional)
    sectionId := int32(56) // int32 | 部門IDで絞込（0を指定すると、部門が未選択の貸借行を絞り込めます） (optional)
    segment1TagId := int32(56) // int32 | セグメント１IDで絞込（0を指定すると、セグメント１が未選択の貸借行を絞り込めます） (optional)
    segment2TagId := int32(56) // int32 | セグメント２IDで絞込（0を指定すると、セグメント２が未選択の貸借行を絞り込めます） (optional)
    segment3TagId := int32(56) // int32 | セグメント３IDで絞込（0を指定すると、セグメント３が未選択の貸借行を絞り込めます） (optional)
    commentStatus := "commentStatus_example" // string | コメント状態で絞込（自分宛のコメント: posted_with_mention, 自分宛のコメント-未解決: raised_with_mention, 自分宛のコメント-解決済: resolved_with_mention, コメントあり: posted, 未解決: raised, 解決済: resolved, コメントなし: none） (optional)
    commentImportant := true // bool | 重要コメント付きの振替伝票を絞込 (optional)
    adjustment := "adjustment_example" // string | 決算整理仕訳で絞込（決算整理仕訳のみ: only, 決算整理仕訳以外: without） (optional)
    txnNumber := "txnNumber_example" // string | 仕訳番号で絞込（事業所の仕訳番号形式が有効な場合のみ） (optional)
    offset := int32(56) // int32 | 取得レコードのオフセット (デフォルト: 0) (optional)
    limit := int32(56) // int32 | 取得レコードの件数 (デフォルト: 20, 最小: 1, 最大: 500)  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManualJournalsApi.GetManualJournals(context.Background()).CompanyId(companyId).StartIssueDate(startIssueDate).EndIssueDate(endIssueDate).EntrySide(entrySide).AccountItemId(accountItemId).MinAmount(minAmount).MaxAmount(maxAmount).PartnerId(partnerId).PartnerCode(partnerCode).ItemId(itemId).SectionId(sectionId).Segment1TagId(segment1TagId).Segment2TagId(segment2TagId).Segment3TagId(segment3TagId).CommentStatus(commentStatus).CommentImportant(commentImportant).Adjustment(adjustment).TxnNumber(txnNumber).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManualJournalsApi.GetManualJournals``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetManualJournals`: InlineResponse2003
    fmt.Fprintf(os.Stdout, "Response from `ManualJournalsApi.GetManualJournals`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetManualJournalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **int32** | 事業所ID | 
 **startIssueDate** | **string** | 発生日で絞込：開始日(yyyy-mm-dd) | 
 **endIssueDate** | **string** | 発生日で絞込：終了日(yyyy-mm-dd) | 
 **entrySide** | **string** | 貸借で絞込 (貸方: credit, 借方: debit) | 
 **accountItemId** | **int32** | 勘定科目IDで絞込 | 
 **minAmount** | **int32** | 金額で絞込：下限 | 
 **maxAmount** | **int32** | 金額で絞込：上限 | 
 **partnerId** | **int32** | 取引先IDで絞込（0を指定すると、取引先が未選択の貸借行を絞り込めます） | 
 **partnerCode** | **string** | 取引先コードで絞込 | 
 **itemId** | **int32** | 品目IDで絞込（0を指定すると、品目が未選択の貸借行を絞り込めます） | 
 **sectionId** | **int32** | 部門IDで絞込（0を指定すると、部門が未選択の貸借行を絞り込めます） | 
 **segment1TagId** | **int32** | セグメント１IDで絞込（0を指定すると、セグメント１が未選択の貸借行を絞り込めます） | 
 **segment2TagId** | **int32** | セグメント２IDで絞込（0を指定すると、セグメント２が未選択の貸借行を絞り込めます） | 
 **segment3TagId** | **int32** | セグメント３IDで絞込（0を指定すると、セグメント３が未選択の貸借行を絞り込めます） | 
 **commentStatus** | **string** | コメント状態で絞込（自分宛のコメント: posted_with_mention, 自分宛のコメント-未解決: raised_with_mention, 自分宛のコメント-解決済: resolved_with_mention, コメントあり: posted, 未解決: raised, 解決済: resolved, コメントなし: none） | 
 **commentImportant** | **bool** | 重要コメント付きの振替伝票を絞込 | 
 **adjustment** | **string** | 決算整理仕訳で絞込（決算整理仕訳のみ: only, 決算整理仕訳以外: without） | 
 **txnNumber** | **string** | 仕訳番号で絞込（事業所の仕訳番号形式が有効な場合のみ） | 
 **offset** | **int32** | 取得レコードのオフセット (デフォルト: 0) | 
 **limit** | **int32** | 取得レコードの件数 (デフォルト: 20, 最小: 1, 最大: 500)  | 

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateManualJournal

> ManualJournalResponse UpdateManualJournal(ctx, id).ManualJournalUpdateParams(manualJournalUpdateParams).Execute()

振替伝票の更新



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
    manualJournalUpdateParams := *openapiclient.NewManualJournalUpdateParams(int32(1), "2013-01-01", []openapiclient.ManualJournalUpdateParamsDetails{*openapiclient.NewManualJournalUpdateParamsDetails("debit", int32(1), int32(1), int32(10800))}) // ManualJournalUpdateParams | 振替伝票の更新 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManualJournalsApi.UpdateManualJournal(context.Background(), id).ManualJournalUpdateParams(manualJournalUpdateParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManualJournalsApi.UpdateManualJournal``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateManualJournal`: ManualJournalResponse
    fmt.Fprintf(os.Stdout, "Response from `ManualJournalsApi.UpdateManualJournal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateManualJournalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **manualJournalUpdateParams** | [**ManualJournalUpdateParams**](ManualJournalUpdateParams.md) | 振替伝票の更新 | 

### Return type

[**ManualJournalResponse**](manualJournalResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

