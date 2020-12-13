# \ExpenseApplicationLineTemplatesApi

All URIs are relative to *https://api.freee.co.jp*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateExpenseApplicationLineTemplate**](ExpenseApplicationLineTemplatesApi.md#CreateExpenseApplicationLineTemplate) | **Post** /api/1/expense_application_line_templates | 経費科目の作成
[**DestroyExpenseApplicationLineTemplate**](ExpenseApplicationLineTemplatesApi.md#DestroyExpenseApplicationLineTemplate) | **Delete** /api/1/expense_application_line_templates/{id} | 経費科目の削除
[**GetExpenseApplicationLineTemplate**](ExpenseApplicationLineTemplatesApi.md#GetExpenseApplicationLineTemplate) | **Get** /api/1/expense_application_line_templates/{id} | 経費科目の取得
[**GetExpenseApplicationLineTemplates**](ExpenseApplicationLineTemplatesApi.md#GetExpenseApplicationLineTemplates) | **Get** /api/1/expense_application_line_templates | 経費科目一覧の取得
[**UpdateExpenseApplicationLineTemplate**](ExpenseApplicationLineTemplatesApi.md#UpdateExpenseApplicationLineTemplate) | **Put** /api/1/expense_application_line_templates/{id} | 経費科目の更新



## CreateExpenseApplicationLineTemplate

> ExpenseApplicationLineTemplateResponse CreateExpenseApplicationLineTemplate(ctx).ExpenseApplicationLineTemplateParams(expenseApplicationLineTemplateParams).Execute()

経費科目の作成

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
    expenseApplicationLineTemplateParams := *openapiclient.NewExpenseApplicationLineTemplateParams(int32(1), "交通費", int32(1), int32(1)) // ExpenseApplicationLineTemplateParams | 経費科目の作成

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExpenseApplicationLineTemplatesApi.CreateExpenseApplicationLineTemplate(context.Background()).ExpenseApplicationLineTemplateParams(expenseApplicationLineTemplateParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExpenseApplicationLineTemplatesApi.CreateExpenseApplicationLineTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateExpenseApplicationLineTemplate`: ExpenseApplicationLineTemplateResponse
    fmt.Fprintf(os.Stdout, "Response from `ExpenseApplicationLineTemplatesApi.CreateExpenseApplicationLineTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateExpenseApplicationLineTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **expenseApplicationLineTemplateParams** | [**ExpenseApplicationLineTemplateParams**](ExpenseApplicationLineTemplateParams.md) | 経費科目の作成 | 

### Return type

[**ExpenseApplicationLineTemplateResponse**](expenseApplicationLineTemplateResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyExpenseApplicationLineTemplate

> DestroyExpenseApplicationLineTemplate(ctx, id).CompanyId(companyId).Execute()

経費科目の削除

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
    id := int32(56) // int32 | 経費科目ID
    companyId := int32(56) // int32 | 事業所ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExpenseApplicationLineTemplatesApi.DestroyExpenseApplicationLineTemplate(context.Background(), id).CompanyId(companyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExpenseApplicationLineTemplatesApi.DestroyExpenseApplicationLineTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | 経費科目ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDestroyExpenseApplicationLineTemplateRequest struct via the builder pattern


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


## GetExpenseApplicationLineTemplate

> ExpenseApplicationLineTemplateResponse GetExpenseApplicationLineTemplate(ctx, id).CompanyId(companyId).Execute()

経費科目の取得

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
    id := int32(56) // int32 | 経費科目ID
    companyId := int32(56) // int32 | 事業所ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExpenseApplicationLineTemplatesApi.GetExpenseApplicationLineTemplate(context.Background(), id).CompanyId(companyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExpenseApplicationLineTemplatesApi.GetExpenseApplicationLineTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExpenseApplicationLineTemplate`: ExpenseApplicationLineTemplateResponse
    fmt.Fprintf(os.Stdout, "Response from `ExpenseApplicationLineTemplatesApi.GetExpenseApplicationLineTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | 経費科目ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExpenseApplicationLineTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **companyId** | **int32** | 事業所ID | 

### Return type

[**ExpenseApplicationLineTemplateResponse**](expenseApplicationLineTemplateResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExpenseApplicationLineTemplates

> InlineResponse20015 GetExpenseApplicationLineTemplates(ctx).CompanyId(companyId).Offset(offset).Limit(limit).Execute()

経費科目一覧の取得



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
    offset := int32(56) // int32 | 取得レコードのオフセット (デフォルト: 0) (optional)
    limit := int32(56) // int32 | 取得レコードの件数 (デフォルト: 20, 最小: 1, 最大: 100) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExpenseApplicationLineTemplatesApi.GetExpenseApplicationLineTemplates(context.Background()).CompanyId(companyId).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExpenseApplicationLineTemplatesApi.GetExpenseApplicationLineTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExpenseApplicationLineTemplates`: InlineResponse20015
    fmt.Fprintf(os.Stdout, "Response from `ExpenseApplicationLineTemplatesApi.GetExpenseApplicationLineTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetExpenseApplicationLineTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **int32** | 事業所ID | 
 **offset** | **int32** | 取得レコードのオフセット (デフォルト: 0) | 
 **limit** | **int32** | 取得レコードの件数 (デフォルト: 20, 最小: 1, 最大: 100) | 

### Return type

[**InlineResponse20015**](inline_response_200_15.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateExpenseApplicationLineTemplate

> ExpenseApplicationLineTemplateResponse UpdateExpenseApplicationLineTemplate(ctx, id).ExpenseApplicationLineTemplateParams(expenseApplicationLineTemplateParams).Execute()

経費科目の更新

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
    id := int32(56) // int32 | 経費科目ID
    expenseApplicationLineTemplateParams := *openapiclient.NewExpenseApplicationLineTemplateParams(int32(1), "交通費", int32(1), int32(1)) // ExpenseApplicationLineTemplateParams | 経費科目の更新

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExpenseApplicationLineTemplatesApi.UpdateExpenseApplicationLineTemplate(context.Background(), id).ExpenseApplicationLineTemplateParams(expenseApplicationLineTemplateParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExpenseApplicationLineTemplatesApi.UpdateExpenseApplicationLineTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateExpenseApplicationLineTemplate`: ExpenseApplicationLineTemplateResponse
    fmt.Fprintf(os.Stdout, "Response from `ExpenseApplicationLineTemplatesApi.UpdateExpenseApplicationLineTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | 経費科目ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateExpenseApplicationLineTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expenseApplicationLineTemplateParams** | [**ExpenseApplicationLineTemplateParams**](ExpenseApplicationLineTemplateParams.md) | 経費科目の更新 | 

### Return type

[**ExpenseApplicationLineTemplateResponse**](expenseApplicationLineTemplateResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

