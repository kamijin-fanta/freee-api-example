# \SectionsApi

All URIs are relative to *https://api.freee.co.jp*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSection**](SectionsApi.md#CreateSection) | **Post** /api/1/sections | 部門の作成
[**DestroySection**](SectionsApi.md#DestroySection) | **Delete** /api/1/sections/{id} | 部門の削除
[**GetSection**](SectionsApi.md#GetSection) | **Get** /api/1/sections/{id} | 
[**GetSections**](SectionsApi.md#GetSections) | **Get** /api/1/sections | 部門一覧の取得
[**UpdateSection**](SectionsApi.md#UpdateSection) | **Put** /api/1/sections/{id} | 部門の更新



## CreateSection

> SectionResponse CreateSection(ctx).SectionParams(sectionParams).Execute()

部門の作成



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
    sectionParams := *openapiclient.NewSectionParams(int32(1), "開発部門") // SectionParams | 部門の作成 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SectionsApi.CreateSection(context.Background()).SectionParams(sectionParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SectionsApi.CreateSection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSection`: SectionResponse
    fmt.Fprintf(os.Stdout, "Response from `SectionsApi.CreateSection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sectionParams** | [**SectionParams**](SectionParams.md) | 部門の作成 | 

### Return type

[**SectionResponse**](sectionResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroySection

> DestroySection(ctx, id).CompanyId(companyId).Execute()

部門の削除



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
    resp, r, err := api_client.SectionsApi.DestroySection(context.Background(), id).CompanyId(companyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SectionsApi.DestroySection``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDestroySectionRequest struct via the builder pattern


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


## GetSection

> SectionResponse GetSection(ctx, id).CompanyId(companyId).Execute()





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
    id := int32(56) // int32 | 部門ID
    companyId := int32(56) // int32 | 事業所ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SectionsApi.GetSection(context.Background(), id).CompanyId(companyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SectionsApi.GetSection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSection`: SectionResponse
    fmt.Fprintf(os.Stdout, "Response from `SectionsApi.GetSection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | 部門ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **companyId** | **int32** | 事業所ID | 

### Return type

[**SectionResponse**](sectionResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSections

> InlineResponse2001 GetSections(ctx).CompanyId(companyId).Execute()

部門一覧の取得



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
    resp, r, err := api_client.SectionsApi.GetSections(context.Background()).CompanyId(companyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SectionsApi.GetSections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSections`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `SectionsApi.GetSections`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **int32** | 事業所ID | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSection

> SectionResponse UpdateSection(ctx, id).SectionParams(sectionParams).Execute()

部門の更新



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
    sectionParams := *openapiclient.NewSectionParams(int32(1), "開発部門") // SectionParams | 部門の更新 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SectionsApi.UpdateSection(context.Background(), id).SectionParams(sectionParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SectionsApi.UpdateSection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSection`: SectionResponse
    fmt.Fprintf(os.Stdout, "Response from `SectionsApi.UpdateSection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sectionParams** | [**SectionParams**](SectionParams.md) | 部門の更新 | 

### Return type

[**SectionResponse**](sectionResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

