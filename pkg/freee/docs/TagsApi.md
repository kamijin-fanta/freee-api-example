# \TagsApi

All URIs are relative to *https://api.freee.co.jp*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTag**](TagsApi.md#CreateTag) | **Post** /api/1/tags | メモタグの作成
[**DestroyTag**](TagsApi.md#DestroyTag) | **Delete** /api/1/tags/{id} | メモタグの削除
[**GetTag**](TagsApi.md#GetTag) | **Get** /api/1/tags/{id} | メモタグの詳細情報の取得
[**GetTags**](TagsApi.md#GetTags) | **Get** /api/1/tags | メモタグ一覧の取得
[**UpdateTag**](TagsApi.md#UpdateTag) | **Put** /api/1/tags/{id} | メモタグの更新



## CreateTag

> TagResponse CreateTag(ctx).TagParams(tagParams).Execute()

メモタグの作成



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
    tagParams := *openapiclient.NewTagParams(int32(1), "メモタグ1") // TagParams | メモタグの作成

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TagsApi.CreateTag(context.Background()).TagParams(tagParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.CreateTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTag`: TagResponse
    fmt.Fprintf(os.Stdout, "Response from `TagsApi.CreateTag`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tagParams** | [**TagParams**](TagParams.md) | メモタグの作成 | 

### Return type

[**TagResponse**](tagResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyTag

> DestroyTag(ctx, id).CompanyId(companyId).Execute()

メモタグの削除



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
    id := int32(56) // int32 | タグID
    companyId := int32(56) // int32 | 事業所ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TagsApi.DestroyTag(context.Background(), id).CompanyId(companyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.DestroyTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | タグID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDestroyTagRequest struct via the builder pattern


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


## GetTag

> TagResponse GetTag(ctx, id).CompanyId(companyId).Execute()

メモタグの詳細情報の取得



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
    id := int32(56) // int32 | タグID
    companyId := int32(56) // int32 | 事業所ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TagsApi.GetTag(context.Background(), id).CompanyId(companyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.GetTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTag`: TagResponse
    fmt.Fprintf(os.Stdout, "Response from `TagsApi.GetTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | タグID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **companyId** | **int32** | 事業所ID | 

### Return type

[**TagResponse**](tagResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTags

> InlineResponse200 GetTags(ctx).CompanyId(companyId).Offset(offset).Limit(limit).Execute()

メモタグ一覧の取得



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
    limit := int32(56) // int32 | 取得レコードの件数 (デフォルト: 50, 最小: 1, 最大: 3000) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TagsApi.GetTags(context.Background()).CompanyId(companyId).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.GetTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTags`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `TagsApi.GetTags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **int32** | 事業所ID | 
 **offset** | **int32** | 取得レコードのオフセット (デフォルト: 0) | 
 **limit** | **int32** | 取得レコードの件数 (デフォルト: 50, 最小: 1, 最大: 3000) | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTag

> TagResponse UpdateTag(ctx, id).TagParams(tagParams).Execute()

メモタグの更新



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
    id := int32(56) // int32 | メモタグID
    tagParams := *openapiclient.NewTagParams(int32(1), "メモタグ1") // TagParams | メモタグの更新 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TagsApi.UpdateTag(context.Background(), id).TagParams(tagParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.UpdateTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTag`: TagResponse
    fmt.Fprintf(os.Stdout, "Response from `TagsApi.UpdateTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | メモタグID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tagParams** | [**TagParams**](TagParams.md) | メモタグの更新 | 

### Return type

[**TagResponse**](tagResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

