# \TrialBalanceApi

All URIs are relative to *https://api.freee.co.jp*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTrialBs**](TrialBalanceApi.md#GetTrialBs) | **Get** /api/1/reports/trial_bs | 貸借対照表の取得
[**GetTrialBsThreeYears**](TrialBalanceApi.md#GetTrialBsThreeYears) | **Get** /api/1/reports/trial_bs_three_years | 貸借対照表(３期間比較)の取得
[**GetTrialBsTwoYears**](TrialBalanceApi.md#GetTrialBsTwoYears) | **Get** /api/1/reports/trial_bs_two_years | 貸借対照表(前年比較)の取得
[**GetTrialPl**](TrialBalanceApi.md#GetTrialPl) | **Get** /api/1/reports/trial_pl | 損益計算書の取得
[**GetTrialPlSections**](TrialBalanceApi.md#GetTrialPlSections) | **Get** /api/1/reports/trial_pl_sections | 損益計算書(部門比較)の取得
[**GetTrialPlThreeYears**](TrialBalanceApi.md#GetTrialPlThreeYears) | **Get** /api/1/reports/trial_pl_three_years | 損益計算書(３期間比較)の取得
[**GetTrialPlTwoYears**](TrialBalanceApi.md#GetTrialPlTwoYears) | **Get** /api/1/reports/trial_pl_two_years | 損益計算書(前年比較)の取得



## GetTrialBs

> TrialBsResponse GetTrialBs(ctx).CompanyId(companyId).FiscalYear(fiscalYear).StartMonth(startMonth).EndMonth(endMonth).StartDate(startDate).EndDate(endDate).AccountItemDisplayType(accountItemDisplayType).BreakdownDisplayType(breakdownDisplayType).PartnerId(partnerId).PartnerCode(partnerCode).ItemId(itemId).Adjustment(adjustment).Execute()

貸借対照表の取得



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
    fiscalYear := int32(56) // int32 | 会計年度 (optional)
    startMonth := int32(56) // int32 | 発生月で絞込：開始会計月(1-12) (optional)
    endMonth := int32(56) // int32 | 発生月で絞込：終了会計月(1-12) (optional)
    startDate := "startDate_example" // string | 発生日で絞込：開始日(yyyy-mm-dd) (optional)
    endDate := "endDate_example" // string | 発生日で絞込：終了日(yyyy-mm-dd) (optional)
    accountItemDisplayType := "accountItemDisplayType_example" // string | 勘定科目の表示（勘定科目: account_item, 決算書表示:group） (optional)
    breakdownDisplayType := "breakdownDisplayType_example" // string | 内訳の表示（取引先: partner, 品目: item, 勘定科目: account_item） ※勘定科目はaccount_item_display_typeが「group」の時のみ指定できます (optional)
    partnerId := int32(56) // int32 | 取引先IDで絞込（0を指定すると、取引先が未選択で絞り込めます） (optional)
    partnerCode := "partnerCode_example" // string | 取引先コードで絞込（事業所設定で取引先コードの利用を有効にしている場合のみ利用可能です） (optional)
    itemId := int32(56) // int32 | 品目IDで絞込（0を指定すると、品目が未選択で絞り込めます） (optional)
    adjustment := "adjustment_example" // string | 決算整理仕訳で絞込（決算整理仕訳のみ: only, 決算整理仕訳以外: without） (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TrialBalanceApi.GetTrialBs(context.Background()).CompanyId(companyId).FiscalYear(fiscalYear).StartMonth(startMonth).EndMonth(endMonth).StartDate(startDate).EndDate(endDate).AccountItemDisplayType(accountItemDisplayType).BreakdownDisplayType(breakdownDisplayType).PartnerId(partnerId).PartnerCode(partnerCode).ItemId(itemId).Adjustment(adjustment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrialBalanceApi.GetTrialBs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTrialBs`: TrialBsResponse
    fmt.Fprintf(os.Stdout, "Response from `TrialBalanceApi.GetTrialBs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTrialBsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **int32** | 事業所ID | 
 **fiscalYear** | **int32** | 会計年度 | 
 **startMonth** | **int32** | 発生月で絞込：開始会計月(1-12) | 
 **endMonth** | **int32** | 発生月で絞込：終了会計月(1-12) | 
 **startDate** | **string** | 発生日で絞込：開始日(yyyy-mm-dd) | 
 **endDate** | **string** | 発生日で絞込：終了日(yyyy-mm-dd) | 
 **accountItemDisplayType** | **string** | 勘定科目の表示（勘定科目: account_item, 決算書表示:group） | 
 **breakdownDisplayType** | **string** | 内訳の表示（取引先: partner, 品目: item, 勘定科目: account_item） ※勘定科目はaccount_item_display_typeが「group」の時のみ指定できます | 
 **partnerId** | **int32** | 取引先IDで絞込（0を指定すると、取引先が未選択で絞り込めます） | 
 **partnerCode** | **string** | 取引先コードで絞込（事業所設定で取引先コードの利用を有効にしている場合のみ利用可能です） | 
 **itemId** | **int32** | 品目IDで絞込（0を指定すると、品目が未選択で絞り込めます） | 
 **adjustment** | **string** | 決算整理仕訳で絞込（決算整理仕訳のみ: only, 決算整理仕訳以外: without） | 

### Return type

[**TrialBsResponse**](trialBsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrialBsThreeYears

> TrialBsThreeYearsResponse GetTrialBsThreeYears(ctx).CompanyId(companyId).FiscalYear(fiscalYear).StartMonth(startMonth).EndMonth(endMonth).StartDate(startDate).EndDate(endDate).AccountItemDisplayType(accountItemDisplayType).BreakdownDisplayType(breakdownDisplayType).PartnerId(partnerId).PartnerCode(partnerCode).ItemId(itemId).Adjustment(adjustment).Execute()

貸借対照表(３期間比較)の取得



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
    fiscalYear := int32(56) // int32 | 会計年度 (optional)
    startMonth := int32(56) // int32 | 発生月で絞込：開始会計月(1-12) (optional)
    endMonth := int32(56) // int32 | 発生月で絞込：終了会計月(1-12) (optional)
    startDate := "startDate_example" // string | 発生日で絞込：開始日(yyyy-mm-dd) (optional)
    endDate := "endDate_example" // string | 発生日で絞込：終了日(yyyy-mm-dd) (optional)
    accountItemDisplayType := "accountItemDisplayType_example" // string | 勘定科目の表示（勘定科目: account_item, 決算書表示:group） (optional)
    breakdownDisplayType := "breakdownDisplayType_example" // string | 内訳の表示（取引先: partner, 品目: item, 勘定科目: account_item） ※勘定科目はaccount_item_display_typeが「group」の時のみ指定できます (optional)
    partnerId := int32(56) // int32 | 取引先IDで絞込（0を指定すると、取引先が未選択で絞り込めます） (optional)
    partnerCode := "partnerCode_example" // string | 取引先コードで絞込（事業所設定で取引先コードの利用を有効にしている場合のみ利用可能です） (optional)
    itemId := int32(56) // int32 | 品目IDで絞込（0を指定すると、品目が未選択で絞り込めます） (optional)
    adjustment := "adjustment_example" // string | 決算整理仕訳で絞込（決算整理仕訳のみ: only, 決算整理仕訳以外: without） (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TrialBalanceApi.GetTrialBsThreeYears(context.Background()).CompanyId(companyId).FiscalYear(fiscalYear).StartMonth(startMonth).EndMonth(endMonth).StartDate(startDate).EndDate(endDate).AccountItemDisplayType(accountItemDisplayType).BreakdownDisplayType(breakdownDisplayType).PartnerId(partnerId).PartnerCode(partnerCode).ItemId(itemId).Adjustment(adjustment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrialBalanceApi.GetTrialBsThreeYears``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTrialBsThreeYears`: TrialBsThreeYearsResponse
    fmt.Fprintf(os.Stdout, "Response from `TrialBalanceApi.GetTrialBsThreeYears`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTrialBsThreeYearsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **int32** | 事業所ID | 
 **fiscalYear** | **int32** | 会計年度 | 
 **startMonth** | **int32** | 発生月で絞込：開始会計月(1-12) | 
 **endMonth** | **int32** | 発生月で絞込：終了会計月(1-12) | 
 **startDate** | **string** | 発生日で絞込：開始日(yyyy-mm-dd) | 
 **endDate** | **string** | 発生日で絞込：終了日(yyyy-mm-dd) | 
 **accountItemDisplayType** | **string** | 勘定科目の表示（勘定科目: account_item, 決算書表示:group） | 
 **breakdownDisplayType** | **string** | 内訳の表示（取引先: partner, 品目: item, 勘定科目: account_item） ※勘定科目はaccount_item_display_typeが「group」の時のみ指定できます | 
 **partnerId** | **int32** | 取引先IDで絞込（0を指定すると、取引先が未選択で絞り込めます） | 
 **partnerCode** | **string** | 取引先コードで絞込（事業所設定で取引先コードの利用を有効にしている場合のみ利用可能です） | 
 **itemId** | **int32** | 品目IDで絞込（0を指定すると、品目が未選択で絞り込めます） | 
 **adjustment** | **string** | 決算整理仕訳で絞込（決算整理仕訳のみ: only, 決算整理仕訳以外: without） | 

### Return type

[**TrialBsThreeYearsResponse**](trialBsThreeYearsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrialBsTwoYears

> TrialBsTwoYearsResponse GetTrialBsTwoYears(ctx).CompanyId(companyId).FiscalYear(fiscalYear).StartMonth(startMonth).EndMonth(endMonth).StartDate(startDate).EndDate(endDate).AccountItemDisplayType(accountItemDisplayType).BreakdownDisplayType(breakdownDisplayType).PartnerId(partnerId).PartnerCode(partnerCode).ItemId(itemId).Adjustment(adjustment).Execute()

貸借対照表(前年比較)の取得



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
    fiscalYear := int32(56) // int32 | 会計年度 (optional)
    startMonth := int32(56) // int32 | 発生月で絞込：開始会計月(1-12) (optional)
    endMonth := int32(56) // int32 | 発生月で絞込：終了会計月(1-12) (optional)
    startDate := "startDate_example" // string | 発生日で絞込：開始日(yyyy-mm-dd) (optional)
    endDate := "endDate_example" // string | 発生日で絞込：終了日(yyyy-mm-dd) (optional)
    accountItemDisplayType := "accountItemDisplayType_example" // string | 勘定科目の表示（勘定科目: account_item, 決算書表示:group） (optional)
    breakdownDisplayType := "breakdownDisplayType_example" // string | 内訳の表示（取引先: partner, 品目: item, 勘定科目: account_item） ※勘定科目はaccount_item_display_typeが「group」の時のみ指定できます (optional)
    partnerId := int32(56) // int32 | 取引先IDで絞込（0を指定すると、取引先が未選択で絞り込めます） (optional)
    partnerCode := "partnerCode_example" // string | 取引先コードで絞込（事業所設定で取引先コードの利用を有効にしている場合のみ利用可能です） (optional)
    itemId := int32(56) // int32 | 品目IDで絞込（0を指定すると、品目が未選択で絞り込めます） (optional)
    adjustment := "adjustment_example" // string | 決算整理仕訳で絞込（決算整理仕訳のみ: only, 決算整理仕訳以外: without） (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TrialBalanceApi.GetTrialBsTwoYears(context.Background()).CompanyId(companyId).FiscalYear(fiscalYear).StartMonth(startMonth).EndMonth(endMonth).StartDate(startDate).EndDate(endDate).AccountItemDisplayType(accountItemDisplayType).BreakdownDisplayType(breakdownDisplayType).PartnerId(partnerId).PartnerCode(partnerCode).ItemId(itemId).Adjustment(adjustment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrialBalanceApi.GetTrialBsTwoYears``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTrialBsTwoYears`: TrialBsTwoYearsResponse
    fmt.Fprintf(os.Stdout, "Response from `TrialBalanceApi.GetTrialBsTwoYears`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTrialBsTwoYearsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **int32** | 事業所ID | 
 **fiscalYear** | **int32** | 会計年度 | 
 **startMonth** | **int32** | 発生月で絞込：開始会計月(1-12) | 
 **endMonth** | **int32** | 発生月で絞込：終了会計月(1-12) | 
 **startDate** | **string** | 発生日で絞込：開始日(yyyy-mm-dd) | 
 **endDate** | **string** | 発生日で絞込：終了日(yyyy-mm-dd) | 
 **accountItemDisplayType** | **string** | 勘定科目の表示（勘定科目: account_item, 決算書表示:group） | 
 **breakdownDisplayType** | **string** | 内訳の表示（取引先: partner, 品目: item, 勘定科目: account_item） ※勘定科目はaccount_item_display_typeが「group」の時のみ指定できます | 
 **partnerId** | **int32** | 取引先IDで絞込（0を指定すると、取引先が未選択で絞り込めます） | 
 **partnerCode** | **string** | 取引先コードで絞込（事業所設定で取引先コードの利用を有効にしている場合のみ利用可能です） | 
 **itemId** | **int32** | 品目IDで絞込（0を指定すると、品目が未選択で絞り込めます） | 
 **adjustment** | **string** | 決算整理仕訳で絞込（決算整理仕訳のみ: only, 決算整理仕訳以外: without） | 

### Return type

[**TrialBsTwoYearsResponse**](trialBsTwoYearsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrialPl

> TrialPlResponse GetTrialPl(ctx).CompanyId(companyId).FiscalYear(fiscalYear).StartMonth(startMonth).EndMonth(endMonth).StartDate(startDate).EndDate(endDate).AccountItemDisplayType(accountItemDisplayType).BreakdownDisplayType(breakdownDisplayType).PartnerId(partnerId).PartnerCode(partnerCode).ItemId(itemId).SectionId(sectionId).Adjustment(adjustment).CostAllocation(costAllocation).Execute()

損益計算書の取得



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
    fiscalYear := int32(56) // int32 | 会計年度 (optional)
    startMonth := int32(56) // int32 | 発生月で絞込：開始会計月(1-12) (optional)
    endMonth := int32(56) // int32 | 発生月で絞込：終了会計月(1-12) (optional)
    startDate := "startDate_example" // string | 発生日で絞込：開始日(yyyy-mm-dd) (optional)
    endDate := "endDate_example" // string | 発生日で絞込：終了日(yyyy-mm-dd) (optional)
    accountItemDisplayType := "accountItemDisplayType_example" // string | 勘定科目の表示（勘定科目: account_item, 決算書表示:group） (optional)
    breakdownDisplayType := "breakdownDisplayType_example" // string | 内訳の表示（取引先: partner, 品目: item, 部門: section, 勘定科目: account_item） ※勘定科目はaccount_item_display_typeが「group」の時のみ指定できます (optional)
    partnerId := int32(56) // int32 | 取引先IDで絞込（0を指定すると、取引先が未選択で絞り込めます） (optional)
    partnerCode := "partnerCode_example" // string | 取引先コードで絞込（事業所設定で取引先コードの利用を有効にしている場合のみ利用可能です） (optional)
    itemId := int32(56) // int32 | 品目IDで絞込（0を指定すると、品目が未選択で絞り込めます） (optional)
    sectionId := int32(56) // int32 | 部門IDで絞込（0を指定すると、部門が未選択で絞り込めます） (optional)
    adjustment := "adjustment_example" // string | 決算整理仕訳で絞込（決算整理仕訳のみ: only, 決算整理仕訳以外: without） (optional)
    costAllocation := "costAllocation_example" // string | 配賦仕訳で絞込（配賦仕訳のみ：only,配賦仕訳以外：without） (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TrialBalanceApi.GetTrialPl(context.Background()).CompanyId(companyId).FiscalYear(fiscalYear).StartMonth(startMonth).EndMonth(endMonth).StartDate(startDate).EndDate(endDate).AccountItemDisplayType(accountItemDisplayType).BreakdownDisplayType(breakdownDisplayType).PartnerId(partnerId).PartnerCode(partnerCode).ItemId(itemId).SectionId(sectionId).Adjustment(adjustment).CostAllocation(costAllocation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrialBalanceApi.GetTrialPl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTrialPl`: TrialPlResponse
    fmt.Fprintf(os.Stdout, "Response from `TrialBalanceApi.GetTrialPl`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTrialPlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **int32** | 事業所ID | 
 **fiscalYear** | **int32** | 会計年度 | 
 **startMonth** | **int32** | 発生月で絞込：開始会計月(1-12) | 
 **endMonth** | **int32** | 発生月で絞込：終了会計月(1-12) | 
 **startDate** | **string** | 発生日で絞込：開始日(yyyy-mm-dd) | 
 **endDate** | **string** | 発生日で絞込：終了日(yyyy-mm-dd) | 
 **accountItemDisplayType** | **string** | 勘定科目の表示（勘定科目: account_item, 決算書表示:group） | 
 **breakdownDisplayType** | **string** | 内訳の表示（取引先: partner, 品目: item, 部門: section, 勘定科目: account_item） ※勘定科目はaccount_item_display_typeが「group」の時のみ指定できます | 
 **partnerId** | **int32** | 取引先IDで絞込（0を指定すると、取引先が未選択で絞り込めます） | 
 **partnerCode** | **string** | 取引先コードで絞込（事業所設定で取引先コードの利用を有効にしている場合のみ利用可能です） | 
 **itemId** | **int32** | 品目IDで絞込（0を指定すると、品目が未選択で絞り込めます） | 
 **sectionId** | **int32** | 部門IDで絞込（0を指定すると、部門が未選択で絞り込めます） | 
 **adjustment** | **string** | 決算整理仕訳で絞込（決算整理仕訳のみ: only, 決算整理仕訳以外: without） | 
 **costAllocation** | **string** | 配賦仕訳で絞込（配賦仕訳のみ：only,配賦仕訳以外：without） | 

### Return type

[**TrialPlResponse**](trialPlResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrialPlSections

> TrialPlSectionsResponse GetTrialPlSections(ctx).CompanyId(companyId).SectionIds(sectionIds).FiscalYear(fiscalYear).StartMonth(startMonth).EndMonth(endMonth).StartDate(startDate).EndDate(endDate).AccountItemDisplayType(accountItemDisplayType).BreakdownDisplayType(breakdownDisplayType).PartnerId(partnerId).PartnerCode(partnerCode).ItemId(itemId).Adjustment(adjustment).CostAllocation(costAllocation).Execute()

損益計算書(部門比較)の取得



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
    sectionIds := "sectionIds_example" // string | 出力する部門の指定（半角数字のidを半角カンマ区切りスペースなしで指定してください）
    fiscalYear := int32(56) // int32 | 会計年度 (optional)
    startMonth := int32(56) // int32 | 発生月で絞込：開始会計月(1-12) (optional)
    endMonth := int32(56) // int32 | 発生月で絞込：終了会計月(1-12) (optional)
    startDate := "startDate_example" // string | 発生日で絞込：開始日(yyyy-mm-dd) (optional)
    endDate := "endDate_example" // string | 発生日で絞込：終了日(yyyy-mm-dd) (optional)
    accountItemDisplayType := "accountItemDisplayType_example" // string | 勘定科目の表示（勘定科目: account_item, 決算書表示:group） (optional)
    breakdownDisplayType := "breakdownDisplayType_example" // string | 内訳の表示（取引先: partner, 品目: item, 勘定科目: account_item） ※勘定科目はaccount_item_display_typeが「group」の時のみ指定できます (optional)
    partnerId := int32(56) // int32 | 取引先IDで絞込（0を指定すると、取引先が未選択で絞り込めます） (optional)
    partnerCode := "partnerCode_example" // string | 取引先コードで絞込（事業所設定で取引先コードの利用を有効にしている場合のみ利用可能です） (optional)
    itemId := int32(56) // int32 | 品目IDで絞込（0を指定すると、品目が未選択で絞り込めます） (optional)
    adjustment := "adjustment_example" // string | 決算整理仕訳で絞込（決算整理仕訳のみ: only, 決算整理仕訳以外: without） (optional)
    costAllocation := "costAllocation_example" // string | 配賦仕訳で絞込（配賦仕訳のみ：only,配賦仕訳以外：without） (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TrialBalanceApi.GetTrialPlSections(context.Background()).CompanyId(companyId).SectionIds(sectionIds).FiscalYear(fiscalYear).StartMonth(startMonth).EndMonth(endMonth).StartDate(startDate).EndDate(endDate).AccountItemDisplayType(accountItemDisplayType).BreakdownDisplayType(breakdownDisplayType).PartnerId(partnerId).PartnerCode(partnerCode).ItemId(itemId).Adjustment(adjustment).CostAllocation(costAllocation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrialBalanceApi.GetTrialPlSections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTrialPlSections`: TrialPlSectionsResponse
    fmt.Fprintf(os.Stdout, "Response from `TrialBalanceApi.GetTrialPlSections`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTrialPlSectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **int32** | 事業所ID | 
 **sectionIds** | **string** | 出力する部門の指定（半角数字のidを半角カンマ区切りスペースなしで指定してください） | 
 **fiscalYear** | **int32** | 会計年度 | 
 **startMonth** | **int32** | 発生月で絞込：開始会計月(1-12) | 
 **endMonth** | **int32** | 発生月で絞込：終了会計月(1-12) | 
 **startDate** | **string** | 発生日で絞込：開始日(yyyy-mm-dd) | 
 **endDate** | **string** | 発生日で絞込：終了日(yyyy-mm-dd) | 
 **accountItemDisplayType** | **string** | 勘定科目の表示（勘定科目: account_item, 決算書表示:group） | 
 **breakdownDisplayType** | **string** | 内訳の表示（取引先: partner, 品目: item, 勘定科目: account_item） ※勘定科目はaccount_item_display_typeが「group」の時のみ指定できます | 
 **partnerId** | **int32** | 取引先IDで絞込（0を指定すると、取引先が未選択で絞り込めます） | 
 **partnerCode** | **string** | 取引先コードで絞込（事業所設定で取引先コードの利用を有効にしている場合のみ利用可能です） | 
 **itemId** | **int32** | 品目IDで絞込（0を指定すると、品目が未選択で絞り込めます） | 
 **adjustment** | **string** | 決算整理仕訳で絞込（決算整理仕訳のみ: only, 決算整理仕訳以外: without） | 
 **costAllocation** | **string** | 配賦仕訳で絞込（配賦仕訳のみ：only,配賦仕訳以外：without） | 

### Return type

[**TrialPlSectionsResponse**](trialPlSectionsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrialPlThreeYears

> TrialPlThreeYearsResponse GetTrialPlThreeYears(ctx).CompanyId(companyId).FiscalYear(fiscalYear).StartMonth(startMonth).EndMonth(endMonth).StartDate(startDate).EndDate(endDate).AccountItemDisplayType(accountItemDisplayType).BreakdownDisplayType(breakdownDisplayType).PartnerId(partnerId).PartnerCode(partnerCode).ItemId(itemId).SectionId(sectionId).Adjustment(adjustment).CostAllocation(costAllocation).Execute()

損益計算書(３期間比較)の取得



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
    fiscalYear := int32(56) // int32 | 会計年度 (optional)
    startMonth := int32(56) // int32 | 発生月で絞込：開始会計月(1-12) (optional)
    endMonth := int32(56) // int32 | 発生月で絞込：終了会計月(1-12) (optional)
    startDate := "startDate_example" // string | 発生日で絞込：開始日(yyyy-mm-dd) (optional)
    endDate := "endDate_example" // string | 発生日で絞込：終了日(yyyy-mm-dd) (optional)
    accountItemDisplayType := "accountItemDisplayType_example" // string | 勘定科目の表示（勘定科目: account_item, 決算書表示:group） (optional)
    breakdownDisplayType := "breakdownDisplayType_example" // string | 内訳の表示（取引先: partner, 品目: item, 部門: section, 勘定科目: account_item） ※勘定科目はaccount_item_display_typeが「group」の時のみ指定できます (optional)
    partnerId := int32(56) // int32 | 取引先IDで絞込（0を指定すると、取引先が未選択で絞り込めます） (optional)
    partnerCode := "partnerCode_example" // string | 取引先コードで絞込（事業所設定で取引先コードの利用を有効にしている場合のみ利用可能です） (optional)
    itemId := int32(56) // int32 | 品目IDで絞込（0を指定すると、品目が未選択で絞り込めます） (optional)
    sectionId := int32(56) // int32 | 部門IDで絞込（0を指定すると、部門が未選択で絞り込めます） (optional)
    adjustment := "adjustment_example" // string | 決算整理仕訳で絞込（決算整理仕訳のみ: only, 決算整理仕訳以外: without） (optional)
    costAllocation := "costAllocation_example" // string | 配賦仕訳で絞込（配賦仕訳のみ：only,配賦仕訳以外：without） (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TrialBalanceApi.GetTrialPlThreeYears(context.Background()).CompanyId(companyId).FiscalYear(fiscalYear).StartMonth(startMonth).EndMonth(endMonth).StartDate(startDate).EndDate(endDate).AccountItemDisplayType(accountItemDisplayType).BreakdownDisplayType(breakdownDisplayType).PartnerId(partnerId).PartnerCode(partnerCode).ItemId(itemId).SectionId(sectionId).Adjustment(adjustment).CostAllocation(costAllocation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrialBalanceApi.GetTrialPlThreeYears``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTrialPlThreeYears`: TrialPlThreeYearsResponse
    fmt.Fprintf(os.Stdout, "Response from `TrialBalanceApi.GetTrialPlThreeYears`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTrialPlThreeYearsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **int32** | 事業所ID | 
 **fiscalYear** | **int32** | 会計年度 | 
 **startMonth** | **int32** | 発生月で絞込：開始会計月(1-12) | 
 **endMonth** | **int32** | 発生月で絞込：終了会計月(1-12) | 
 **startDate** | **string** | 発生日で絞込：開始日(yyyy-mm-dd) | 
 **endDate** | **string** | 発生日で絞込：終了日(yyyy-mm-dd) | 
 **accountItemDisplayType** | **string** | 勘定科目の表示（勘定科目: account_item, 決算書表示:group） | 
 **breakdownDisplayType** | **string** | 内訳の表示（取引先: partner, 品目: item, 部門: section, 勘定科目: account_item） ※勘定科目はaccount_item_display_typeが「group」の時のみ指定できます | 
 **partnerId** | **int32** | 取引先IDで絞込（0を指定すると、取引先が未選択で絞り込めます） | 
 **partnerCode** | **string** | 取引先コードで絞込（事業所設定で取引先コードの利用を有効にしている場合のみ利用可能です） | 
 **itemId** | **int32** | 品目IDで絞込（0を指定すると、品目が未選択で絞り込めます） | 
 **sectionId** | **int32** | 部門IDで絞込（0を指定すると、部門が未選択で絞り込めます） | 
 **adjustment** | **string** | 決算整理仕訳で絞込（決算整理仕訳のみ: only, 決算整理仕訳以外: without） | 
 **costAllocation** | **string** | 配賦仕訳で絞込（配賦仕訳のみ：only,配賦仕訳以外：without） | 

### Return type

[**TrialPlThreeYearsResponse**](trialPlThreeYearsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrialPlTwoYears

> TrialPlTwoYearsResponse GetTrialPlTwoYears(ctx).CompanyId(companyId).FiscalYear(fiscalYear).StartMonth(startMonth).EndMonth(endMonth).StartDate(startDate).EndDate(endDate).AccountItemDisplayType(accountItemDisplayType).BreakdownDisplayType(breakdownDisplayType).PartnerId(partnerId).PartnerCode(partnerCode).ItemId(itemId).SectionId(sectionId).Adjustment(adjustment).CostAllocation(costAllocation).Execute()

損益計算書(前年比較)の取得



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
    fiscalYear := int32(56) // int32 | 会計年度 (optional)
    startMonth := int32(56) // int32 | 発生月で絞込：開始会計月(1-12) (optional)
    endMonth := int32(56) // int32 | 発生月で絞込：終了会計月(1-12) (optional)
    startDate := "startDate_example" // string | 発生日で絞込：開始日(yyyy-mm-dd) (optional)
    endDate := "endDate_example" // string | 発生日で絞込：終了日(yyyy-mm-dd) (optional)
    accountItemDisplayType := "accountItemDisplayType_example" // string | 勘定科目の表示（勘定科目: account_item, 決算書表示:group） (optional)
    breakdownDisplayType := "breakdownDisplayType_example" // string | 内訳の表示（取引先: partner, 品目: item, 部門: section, 勘定科目: account_item） ※勘定科目はaccount_item_display_typeが「group」の時のみ指定できます (optional)
    partnerId := int32(56) // int32 | 取引先IDで絞込（0を指定すると、取引先が未選択で絞り込めます） (optional)
    partnerCode := "partnerCode_example" // string | 取引先コードで絞込（事業所設定で取引先コードの利用を有効にしている場合のみ利用可能です） (optional)
    itemId := int32(56) // int32 | 品目IDで絞込（0を指定すると、品目が未選択で絞り込めます） (optional)
    sectionId := int32(56) // int32 | 部門IDで絞込（0を指定すると、部門が未選択で絞り込めます） (optional)
    adjustment := "adjustment_example" // string | 決算整理仕訳で絞込（決算整理仕訳のみ: only, 決算整理仕訳以外: without） (optional)
    costAllocation := "costAllocation_example" // string | 配賦仕訳で絞込（配賦仕訳のみ：only,配賦仕訳以外：without） (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TrialBalanceApi.GetTrialPlTwoYears(context.Background()).CompanyId(companyId).FiscalYear(fiscalYear).StartMonth(startMonth).EndMonth(endMonth).StartDate(startDate).EndDate(endDate).AccountItemDisplayType(accountItemDisplayType).BreakdownDisplayType(breakdownDisplayType).PartnerId(partnerId).PartnerCode(partnerCode).ItemId(itemId).SectionId(sectionId).Adjustment(adjustment).CostAllocation(costAllocation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrialBalanceApi.GetTrialPlTwoYears``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTrialPlTwoYears`: TrialPlTwoYearsResponse
    fmt.Fprintf(os.Stdout, "Response from `TrialBalanceApi.GetTrialPlTwoYears`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTrialPlTwoYearsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **int32** | 事業所ID | 
 **fiscalYear** | **int32** | 会計年度 | 
 **startMonth** | **int32** | 発生月で絞込：開始会計月(1-12) | 
 **endMonth** | **int32** | 発生月で絞込：終了会計月(1-12) | 
 **startDate** | **string** | 発生日で絞込：開始日(yyyy-mm-dd) | 
 **endDate** | **string** | 発生日で絞込：終了日(yyyy-mm-dd) | 
 **accountItemDisplayType** | **string** | 勘定科目の表示（勘定科目: account_item, 決算書表示:group） | 
 **breakdownDisplayType** | **string** | 内訳の表示（取引先: partner, 品目: item, 部門: section, 勘定科目: account_item） ※勘定科目はaccount_item_display_typeが「group」の時のみ指定できます | 
 **partnerId** | **int32** | 取引先IDで絞込（0を指定すると、取引先が未選択で絞り込めます） | 
 **partnerCode** | **string** | 取引先コードで絞込（事業所設定で取引先コードの利用を有効にしている場合のみ利用可能です） | 
 **itemId** | **int32** | 品目IDで絞込（0を指定すると、品目が未選択で絞り込めます） | 
 **sectionId** | **int32** | 部門IDで絞込（0を指定すると、部門が未選択で絞り込めます） | 
 **adjustment** | **string** | 決算整理仕訳で絞込（決算整理仕訳のみ: only, 決算整理仕訳以外: without） | 
 **costAllocation** | **string** | 配賦仕訳で絞込（配賦仕訳のみ：only,配賦仕訳以外：without） | 

### Return type

[**TrialPlTwoYearsResponse**](trialPlTwoYearsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

