# \SelectablesApi

All URIs are relative to *https://api.freee.co.jp*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFormsSelectables**](SelectablesApi.md#GetFormsSelectables) | **Get** /api/1/forms/selectables | フォーム用選択項目情報の取得



## GetFormsSelectables

> SelectablesIndexResponse GetFormsSelectables(ctx).CompanyId(companyId).Includes(includes).Execute()

フォーム用選択項目情報の取得



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
    includes := "includes_example" // string | 取得する項目(項目: account_item) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SelectablesApi.GetFormsSelectables(context.Background()).CompanyId(companyId).Includes(includes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelectablesApi.GetFormsSelectables``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFormsSelectables`: SelectablesIndexResponse
    fmt.Fprintf(os.Stdout, "Response from `SelectablesApi.GetFormsSelectables`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFormsSelectablesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **int32** | 事業所ID | 
 **includes** | **string** | 取得する項目(項目: account_item) | 

### Return type

[**SelectablesIndexResponse**](selectablesIndexResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

