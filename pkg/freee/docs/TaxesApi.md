# \TaxesApi

All URIs are relative to *https://api.freee.co.jp*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTaxCode**](TaxesApi.md#GetTaxCode) | **Get** /api/1/taxes/codes/{code} | 税区分コードの取得
[**GetTaxCodes**](TaxesApi.md#GetTaxCodes) | **Get** /api/1/taxes/codes | 税区分コード一覧の取得
[**GetTaxesCompanies**](TaxesApi.md#GetTaxesCompanies) | **Get** /api/1/taxes/companies/{company_id} | 税区分コード詳細一覧の取得



## GetTaxCode

> TaxResponse GetTaxCode(ctx, code).Execute()

税区分コードの取得



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
    code := int32(56) // int32 | 税区分コード

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TaxesApi.GetTaxCode(context.Background(), code).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxesApi.GetTaxCode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTaxCode`: TaxResponse
    fmt.Fprintf(os.Stdout, "Response from `TaxesApi.GetTaxCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **int32** | 税区分コード | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaxCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TaxResponse**](taxResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTaxCodes

> InlineResponse2007 GetTaxCodes(ctx).Execute()

税区分コード一覧の取得



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
    resp, r, err := api_client.TaxesApi.GetTaxCodes(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxesApi.GetTaxCodes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTaxCodes`: InlineResponse2007
    fmt.Fprintf(os.Stdout, "Response from `TaxesApi.GetTaxCodes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaxCodesRequest struct via the builder pattern


### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTaxesCompanies

> InlineResponse2008 GetTaxesCompanies(ctx, companyId).Execute()

税区分コード詳細一覧の取得

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TaxesApi.GetTaxesCompanies(context.Background(), companyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxesApi.GetTaxesCompanies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTaxesCompanies`: InlineResponse2008
    fmt.Fprintf(os.Stdout, "Response from `TaxesApi.GetTaxesCompanies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32** | 事業所ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaxesCompaniesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse2008**](inline_response_200_8.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

