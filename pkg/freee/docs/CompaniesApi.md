# \CompaniesApi

All URIs are relative to *https://api.freee.co.jp*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCompanies**](CompaniesApi.md#GetCompanies) | **Get** /api/1/companies | 事業所一覧の取得
[**GetCompany**](CompaniesApi.md#GetCompany) | **Get** /api/1/companies/{id} | 事業所の詳細情報の取得
[**UpdateCompany**](CompaniesApi.md#UpdateCompany) | **Put** /api/1/companies/{id} | 事業所情報の更新



## GetCompanies

> CompanyIndexResponse GetCompanies(ctx).Execute()

事業所一覧の取得



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CompaniesApi.GetCompanies(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CompaniesApi.GetCompanies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCompanies`: CompanyIndexResponse
    fmt.Fprintf(os.Stdout, "Response from `CompaniesApi.GetCompanies`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompaniesRequest struct via the builder pattern


### Return type

[**CompanyIndexResponse**](companyIndexResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompany

> CompanyResponse GetCompany(ctx, id).Details(details).AccountItems(accountItems).Taxes(taxes).Items(items).Partners(partners).Sections(sections).Tags(tags).Walletables(walletables).Execute()

事業所の詳細情報の取得



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
    id := int32(56) // int32 | 事業所ID
    details := true // bool | 取得情報に勘定科目・税区分コード・税区分・品目・取引先・部門・メモタグ・口座の一覧を含める (optional)
    accountItems := true // bool | 取得情報に勘定科目一覧を含める (optional)
    taxes := true // bool | 取得情報に税区分コード・税区分一覧を含める (optional)
    items := true // bool | 取得情報に品目一覧を含める (optional)
    partners := true // bool | 取得情報に取引先一覧を含める (optional)
    sections := true // bool | 取得情報に部門一覧を含める (optional)
    tags := true // bool | 取得情報にメモタグ一覧を含める (optional)
    walletables := true // bool | 取得情報に口座一覧を含める (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CompaniesApi.GetCompany(context.Background(), id).Details(details).AccountItems(accountItems).Taxes(taxes).Items(items).Partners(partners).Sections(sections).Tags(tags).Walletables(walletables).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CompaniesApi.GetCompany``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCompany`: CompanyResponse
    fmt.Fprintf(os.Stdout, "Response from `CompaniesApi.GetCompany`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | 事業所ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **details** | **bool** | 取得情報に勘定科目・税区分コード・税区分・品目・取引先・部門・メモタグ・口座の一覧を含める | 
 **accountItems** | **bool** | 取得情報に勘定科目一覧を含める | 
 **taxes** | **bool** | 取得情報に税区分コード・税区分一覧を含める | 
 **items** | **bool** | 取得情報に品目一覧を含める | 
 **partners** | **bool** | 取得情報に取引先一覧を含める | 
 **sections** | **bool** | 取得情報に部門一覧を含める | 
 **tags** | **bool** | 取得情報にメモタグ一覧を含める | 
 **walletables** | **bool** | 取得情報に口座一覧を含める | 

### Return type

[**CompanyResponse**](companyResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCompany

> CompanyUpdateResponse UpdateCompany(ctx, id).CompanyParams(companyParams).Execute()

事業所情報の更新



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
    id := int32(56) // int32 | 事業所ID
    companyParams := *openapiclient.NewCompanyParams() // CompanyParams |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CompaniesApi.UpdateCompany(context.Background(), id).CompanyParams(companyParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CompaniesApi.UpdateCompany``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCompany`: CompanyUpdateResponse
    fmt.Fprintf(os.Stdout, "Response from `CompaniesApi.UpdateCompany`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | 事業所ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCompanyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **companyParams** | [**CompanyParams**](CompanyParams.md) |  | 

### Return type

[**CompanyUpdateResponse**](companyUpdateResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

