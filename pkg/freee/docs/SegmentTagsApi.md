# \SegmentTagsApi

All URIs are relative to *https://api.freee.co.jp*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSegmentTag**](SegmentTagsApi.md#CreateSegmentTag) | **Post** /api/1/segments/{segment_id}/tags | セグメントの作成
[**DestroySegmentsTag**](SegmentTagsApi.md#DestroySegmentsTag) | **Delete** /api/1/segments/{segment_id}/tags/{id} | セグメントタグの削除
[**GetSegmentTags**](SegmentTagsApi.md#GetSegmentTags) | **Get** /api/1/segments/{segment_id}/tags | セグメントタグ一覧の取得
[**UpdateSegmentTag**](SegmentTagsApi.md#UpdateSegmentTag) | **Put** /api/1/segments/{segment_id}/tags/{id} | セグメントタグの更新



## CreateSegmentTag

> SegmentTagResponse CreateSegmentTag(ctx, segmentId).SegmentTagParams(segmentTagParams).Execute()

セグメントの作成



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
    segmentId := int32(56) // int32 | セグメントID（1,2,3のいずれか） 該当プラン以外で参照した場合にはエラーとなります。   1: 法人向けプロフェッショナル, 法人向けエンタープライズプラン   2,3: 法人向け エンタープライズプラン 
    segmentTagParams := *openapiclient.NewSegmentTagParams(int32(1), "プロジェクトA") // SegmentTagParams | セグメントタグの作成

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SegmentTagsApi.CreateSegmentTag(context.Background(), segmentId).SegmentTagParams(segmentTagParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SegmentTagsApi.CreateSegmentTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSegmentTag`: SegmentTagResponse
    fmt.Fprintf(os.Stdout, "Response from `SegmentTagsApi.CreateSegmentTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**segmentId** | **int32** | セグメントID（1,2,3のいずれか） 該当プラン以外で参照した場合にはエラーとなります。   1: 法人向けプロフェッショナル, 法人向けエンタープライズプラン   2,3: 法人向け エンタープライズプラン  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSegmentTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **segmentTagParams** | [**SegmentTagParams**](SegmentTagParams.md) | セグメントタグの作成 | 

### Return type

[**SegmentTagResponse**](segmentTagResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroySegmentsTag

> DestroySegmentsTag(ctx, segmentId, id).CompanyId(companyId).Execute()

セグメントタグの削除



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
    segmentId := int32(56) // int32 | セグメントID（1,2,3のいずれか） 該当プラン以外で参照した場合にはエラーとなります。   1: 法人向けプロフェッショナル, 法人向けエンタープライズプラン   2,3: 法人向け エンタープライズプラン 
    id := int32(56) // int32 | セグメントタグID
    companyId := int32(56) // int32 | 事業所ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SegmentTagsApi.DestroySegmentsTag(context.Background(), segmentId, id).CompanyId(companyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SegmentTagsApi.DestroySegmentsTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**segmentId** | **int32** | セグメントID（1,2,3のいずれか） 該当プラン以外で参照した場合にはエラーとなります。   1: 法人向けプロフェッショナル, 法人向けエンタープライズプラン   2,3: 法人向け エンタープライズプラン  | 
**id** | **int32** | セグメントタグID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDestroySegmentsTagRequest struct via the builder pattern


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


## GetSegmentTags

> InlineResponse20017 GetSegmentTags(ctx, segmentId).CompanyId(companyId).Offset(offset).Limit(limit).Execute()

セグメントタグ一覧の取得



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
    segmentId := int32(56) // int32 | セグメントID（1,2,3のいずれか） 該当プラン以外で参照した場合にはエラーとなります。   1: 法人向けプロフェッショナル, 法人向けエンタープライズプラン   2,3: 法人向け エンタープライズプラン 
    offset := int32(56) // int32 | 取得レコードのオフセット (デフォルト: 0) (optional)
    limit := int32(56) // int32 | 取得レコードの件数 (デフォルト: 20, 最小: 1, 最大: 500)  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SegmentTagsApi.GetSegmentTags(context.Background(), segmentId).CompanyId(companyId).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SegmentTagsApi.GetSegmentTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSegmentTags`: InlineResponse20017
    fmt.Fprintf(os.Stdout, "Response from `SegmentTagsApi.GetSegmentTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**segmentId** | **int32** | セグメントID（1,2,3のいずれか） 該当プラン以外で参照した場合にはエラーとなります。   1: 法人向けプロフェッショナル, 法人向けエンタープライズプラン   2,3: 法人向け エンタープライズプラン  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSegmentTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **int32** | 事業所ID | 

 **offset** | **int32** | 取得レコードのオフセット (デフォルト: 0) | 
 **limit** | **int32** | 取得レコードの件数 (デフォルト: 20, 最小: 1, 最大: 500)  | 

### Return type

[**InlineResponse20017**](inline_response_200_17.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSegmentTag

> SegmentTagResponse UpdateSegmentTag(ctx, segmentId, id).SegmentTagParams(segmentTagParams).Execute()

セグメントタグの更新



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
    segmentId := int32(56) // int32 | セグメントID（1,2,3のいずれか） 該当プラン以外で参照した場合にはエラーとなります。   1: 法人向けプロフェッショナル, 法人向けエンタープライズプラン   2,3: 法人向け エンタープライズプラン 
    id := int32(56) // int32 | セグメントタグID
    segmentTagParams := *openapiclient.NewSegmentTagParams(int32(1), "プロジェクトA") // SegmentTagParams | セグメントタグの作成

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SegmentTagsApi.UpdateSegmentTag(context.Background(), segmentId, id).SegmentTagParams(segmentTagParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SegmentTagsApi.UpdateSegmentTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSegmentTag`: SegmentTagResponse
    fmt.Fprintf(os.Stdout, "Response from `SegmentTagsApi.UpdateSegmentTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**segmentId** | **int32** | セグメントID（1,2,3のいずれか） 該当プラン以外で参照した場合にはエラーとなります。   1: 法人向けプロフェッショナル, 法人向けエンタープライズプラン   2,3: 法人向け エンタープライズプラン  | 
**id** | **int32** | セグメントタグID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSegmentTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **segmentTagParams** | [**SegmentTagParams**](SegmentTagParams.md) | セグメントタグの作成 | 

### Return type

[**SegmentTagResponse**](segmentTagResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

