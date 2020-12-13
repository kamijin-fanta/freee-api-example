# \DealsApi

All URIs are relative to *https://api.freee.co.jp*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDeal**](DealsApi.md#CreateDeal) | **Post** /api/1/deals | 取引（収入／支出）の作成
[**DestroyDeal**](DealsApi.md#DestroyDeal) | **Delete** /api/1/deals/{id} | 取引（収入／支出）の削除
[**GetDeal**](DealsApi.md#GetDeal) | **Get** /api/1/deals/{id} | 取引（収入／支出）の取得
[**GetDeals**](DealsApi.md#GetDeals) | **Get** /api/1/deals | 取引（収入／支出）一覧の取得
[**UpdateDeal**](DealsApi.md#UpdateDeal) | **Put** /api/1/deals/{id} | 取引（収入／支出）の更新



## CreateDeal

> DealCreateResponse CreateDeal(ctx).DealCreateParams(dealCreateParams).Execute()

取引（収入／支出）の作成



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
    dealCreateParams := *openapiclient.NewDealCreateParams("2013-01-01", "income", int32(1), []openapiclient.DealCreateParamsDetails{*openapiclient.NewDealCreateParamsDetails(int32(1), int32(1), int32(1))}) // DealCreateParams | 取引（収入／支出）の作成 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DealsApi.CreateDeal(context.Background()).DealCreateParams(dealCreateParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DealsApi.CreateDeal``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDeal`: DealCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `DealsApi.CreateDeal`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDealRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dealCreateParams** | [**DealCreateParams**](DealCreateParams.md) | 取引（収入／支出）の作成 | 

### Return type

[**DealCreateResponse**](dealCreateResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyDeal

> DestroyDeal(ctx, id).CompanyId(companyId).Execute()

取引（収入／支出）の削除

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
    companyId := int32(56) // int32 | 事業所ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DealsApi.DestroyDeal(context.Background(), id).CompanyId(companyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DealsApi.DestroyDeal``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | 取引ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDestroyDealRequest struct via the builder pattern


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


## GetDeal

> DealResponse GetDeal(ctx, id).CompanyId(companyId).Accruals(accruals).Execute()

取引（収入／支出）の取得



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
    accruals := "accruals_example" // string | 取引の債権債務行の表示（without: 表示しない(デフォルト), with: 表示する） (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DealsApi.GetDeal(context.Background(), id).CompanyId(companyId).Accruals(accruals).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DealsApi.GetDeal``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeal`: DealResponse
    fmt.Fprintf(os.Stdout, "Response from `DealsApi.GetDeal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDealRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **int32** | 事業所ID | 

 **accruals** | **string** | 取引の債権債務行の表示（without: 表示しない(デフォルト), with: 表示する） | 

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


## GetDeals

> InlineResponse2002 GetDeals(ctx).CompanyId(companyId).PartnerId(partnerId).AccountItemId(accountItemId).PartnerCode(partnerCode).Status(status).Type_(type_).StartIssueDate(startIssueDate).EndIssueDate(endIssueDate).StartDueDate(startDueDate).EndDueDate(endDueDate).StartRenewDate(startRenewDate).EndRenewDate(endRenewDate).Offset(offset).Limit(limit).RegisteredFrom(registeredFrom).Accruals(accruals).Execute()

取引（収入／支出）一覧の取得



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
    accountItemId := int32(56) // int32 | 勘定科目IDで絞込 (optional)
    partnerCode := "partnerCode_example" // string | 取引先コードで絞込 (optional)
    status := "status_example" // string | 決済状況で絞込 (未決済: unsettled, 完了: settled) (optional)
    type_ := "type__example" // string | 収支区分 (収入: income, 支出: expense) (optional)
    startIssueDate := "startIssueDate_example" // string | 発生日で絞込：開始日(yyyy-mm-dd) (optional)
    endIssueDate := "endIssueDate_example" // string | 発生日で絞込：終了日(yyyy-mm-dd) (optional)
    startDueDate := "startDueDate_example" // string | 支払期日で絞込：開始日(yyyy-mm-dd) (optional)
    endDueDate := "endDueDate_example" // string | 支払期日で絞込：終了日(yyyy-mm-dd) (optional)
    startRenewDate := "startRenewDate_example" // string | +更新日で絞込：開始日(yyyy-mm-dd) (optional)
    endRenewDate := "endRenewDate_example" // string | +更新日で絞込：終了日(yyyy-mm-dd) (optional)
    offset := int32(56) // int32 | 取得レコードのオフセット (デフォルト: 0) (optional)
    limit := int32(56) // int32 | 取得レコードの件数 (デフォルト: 20, 最大: 100)  (optional)
    registeredFrom := "registeredFrom_example" // string | 取引登録元アプリで絞込（me: 本APIを叩くアプリ自身から登録した取引のみ） (optional)
    accruals := "accruals_example" // string | 取引の債権債務行の表示（without: 表示しない(デフォルト), with: 表示する） (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DealsApi.GetDeals(context.Background()).CompanyId(companyId).PartnerId(partnerId).AccountItemId(accountItemId).PartnerCode(partnerCode).Status(status).Type_(type_).StartIssueDate(startIssueDate).EndIssueDate(endIssueDate).StartDueDate(startDueDate).EndDueDate(endDueDate).StartRenewDate(startRenewDate).EndRenewDate(endRenewDate).Offset(offset).Limit(limit).RegisteredFrom(registeredFrom).Accruals(accruals).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DealsApi.GetDeals``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeals`: InlineResponse2002
    fmt.Fprintf(os.Stdout, "Response from `DealsApi.GetDeals`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDealsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **int32** | 事業所ID | 
 **partnerId** | **int32** | 取引先IDで絞込 | 
 **accountItemId** | **int32** | 勘定科目IDで絞込 | 
 **partnerCode** | **string** | 取引先コードで絞込 | 
 **status** | **string** | 決済状況で絞込 (未決済: unsettled, 完了: settled) | 
 **type_** | **string** | 収支区分 (収入: income, 支出: expense) | 
 **startIssueDate** | **string** | 発生日で絞込：開始日(yyyy-mm-dd) | 
 **endIssueDate** | **string** | 発生日で絞込：終了日(yyyy-mm-dd) | 
 **startDueDate** | **string** | 支払期日で絞込：開始日(yyyy-mm-dd) | 
 **endDueDate** | **string** | 支払期日で絞込：終了日(yyyy-mm-dd) | 
 **startRenewDate** | **string** | +更新日で絞込：開始日(yyyy-mm-dd) | 
 **endRenewDate** | **string** | +更新日で絞込：終了日(yyyy-mm-dd) | 
 **offset** | **int32** | 取得レコードのオフセット (デフォルト: 0) | 
 **limit** | **int32** | 取得レコードの件数 (デフォルト: 20, 最大: 100)  | 
 **registeredFrom** | **string** | 取引登録元アプリで絞込（me: 本APIを叩くアプリ自身から登録した取引のみ） | 
 **accruals** | **string** | 取引の債権債務行の表示（without: 表示しない(デフォルト), with: 表示する） | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeal

> DealResponse UpdateDeal(ctx, id).DealUpdateParams(dealUpdateParams).Execute()

取引（収入／支出）の更新



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
    dealUpdateParams := *openapiclient.NewDealUpdateParams("2013-01-01", "income", int32(1), []openapiclient.DealUpdateParamsDetails{*openapiclient.NewDealUpdateParamsDetails(int32(1), int32(1), int32(1))}) // DealUpdateParams | 取引（収入／支出）の更新 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DealsApi.UpdateDeal(context.Background(), id).DealUpdateParams(dealUpdateParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DealsApi.UpdateDeal``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDeal`: DealResponse
    fmt.Fprintf(os.Stdout, "Response from `DealsApi.UpdateDeal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | 取引ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDealRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dealUpdateParams** | [**DealUpdateParams**](DealUpdateParams.md) | 取引（収入／支出）の更新 | 

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

