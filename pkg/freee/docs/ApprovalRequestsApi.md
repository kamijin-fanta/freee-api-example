# \ApprovalRequestsApi

All URIs are relative to *https://api.freee.co.jp*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApprovalRequest**](ApprovalRequestsApi.md#CreateApprovalRequest) | **Post** /api/1/approval_requests | 各種申請の作成
[**DestroyApprovalRequest**](ApprovalRequestsApi.md#DestroyApprovalRequest) | **Delete** /api/1/approval_requests/{id} | 各種申請の削除
[**GetApprovalRequest**](ApprovalRequestsApi.md#GetApprovalRequest) | **Get** /api/1/approval_requests/{id} | 各種申請の取得
[**GetApprovalRequestForm**](ApprovalRequestsApi.md#GetApprovalRequestForm) | **Get** /api/1/approval_requests/forms/{id} | 各種申請の申請フォームの取得
[**GetApprovalRequestForms**](ApprovalRequestsApi.md#GetApprovalRequestForms) | **Get** /api/1/approval_requests/forms | 各種申請の申請フォーム一覧の取得
[**GetApprovalRequests**](ApprovalRequestsApi.md#GetApprovalRequests) | **Get** /api/1/approval_requests | 各種申請の一覧
[**UpdateApprovalRequest**](ApprovalRequestsApi.md#UpdateApprovalRequest) | **Put** /api/1/approval_requests/{id} | 各種申請の更新
[**UpdateApprovalRequestAction**](ApprovalRequestsApi.md#UpdateApprovalRequestAction) | **Post** /api/1/approval_requests/{id}/actions | 各種申請の承認操作



## CreateApprovalRequest

> ApprovalRequestResponse CreateApprovalRequest(ctx).ApprovalRequestCreateParams(approvalRequestCreateParams).Execute()

各種申請の作成



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
    approvalRequestCreateParams := *openapiclient.NewApprovalRequestCreateParams(int32(1), "2020-01-01", int32(1), int32(1), true, []openapiclient.ApprovalRequestCreateParamsRequestItems{*openapiclient.NewApprovalRequestCreateParamsRequestItems()}) // ApprovalRequestCreateParams | 各種申請の作成 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApprovalRequestsApi.CreateApprovalRequest(context.Background()).ApprovalRequestCreateParams(approvalRequestCreateParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApprovalRequestsApi.CreateApprovalRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApprovalRequest`: ApprovalRequestResponse
    fmt.Fprintf(os.Stdout, "Response from `ApprovalRequestsApi.CreateApprovalRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApprovalRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **approvalRequestCreateParams** | [**ApprovalRequestCreateParams**](ApprovalRequestCreateParams.md) | 各種申請の作成 | 

### Return type

[**ApprovalRequestResponse**](approvalRequestResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyApprovalRequest

> DestroyApprovalRequest(ctx, id).CompanyId(companyId).Execute()

各種申請の削除



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
    id := int32(56) // int32 | 各種申請ID
    companyId := int32(56) // int32 | 事業所ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApprovalRequestsApi.DestroyApprovalRequest(context.Background(), id).CompanyId(companyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApprovalRequestsApi.DestroyApprovalRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | 各種申請ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDestroyApprovalRequestRequest struct via the builder pattern


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


## GetApprovalRequest

> ApprovalRequestResponse GetApprovalRequest(ctx, id).CompanyId(companyId).Execute()

各種申請の取得



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
    id := int32(56) // int32 | 各種申請ID
    companyId := int32(56) // int32 | 事業所ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApprovalRequestsApi.GetApprovalRequest(context.Background(), id).CompanyId(companyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApprovalRequestsApi.GetApprovalRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApprovalRequest`: ApprovalRequestResponse
    fmt.Fprintf(os.Stdout, "Response from `ApprovalRequestsApi.GetApprovalRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | 各種申請ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApprovalRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **companyId** | **int32** | 事業所ID | 

### Return type

[**ApprovalRequestResponse**](approvalRequestResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApprovalRequestForm

> ApprovalRequestFormResponse GetApprovalRequestForm(ctx, id).CompanyId(companyId).Execute()

各種申請の申請フォームの取得



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
    id := int32(56) // int32 | 申請フォームID
    companyId := int32(56) // int32 | 事業所ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApprovalRequestsApi.GetApprovalRequestForm(context.Background(), id).CompanyId(companyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApprovalRequestsApi.GetApprovalRequestForm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApprovalRequestForm`: ApprovalRequestFormResponse
    fmt.Fprintf(os.Stdout, "Response from `ApprovalRequestsApi.GetApprovalRequestForm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | 申請フォームID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApprovalRequestFormRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **companyId** | **int32** | 事業所ID | 

### Return type

[**ApprovalRequestFormResponse**](approvalRequestFormResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApprovalRequestForms

> InlineResponse20016 GetApprovalRequestForms(ctx).CompanyId(companyId).Execute()

各種申請の申請フォーム一覧の取得



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
    resp, r, err := api_client.ApprovalRequestsApi.GetApprovalRequestForms(context.Background()).CompanyId(companyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApprovalRequestsApi.GetApprovalRequestForms``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApprovalRequestForms`: InlineResponse20016
    fmt.Fprintf(os.Stdout, "Response from `ApprovalRequestsApi.GetApprovalRequestForms`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApprovalRequestFormsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **int32** | 事業所ID | 

### Return type

[**InlineResponse20016**](inline_response_200_16.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApprovalRequests

> ApprovalRequestsIndexResponse GetApprovalRequests(ctx).CompanyId(companyId).Status(status).ApplicationNumber(applicationNumber).Title(title).FormId(formId).StartApplicationDate(startApplicationDate).EndApplicationDate(endApplicationDate).ApplicantId(applicantId).ApproverId(approverId).Offset(offset).Limit(limit).Execute()

各種申請の一覧



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
    status := "draft" // string | 申請ステータス(draft:下書き, in_progress:申請中, approved:承認済, rejected:却下, feedback:差戻し) (optional)
    applicationNumber := int32(2) // int32 | 申請No. (optional)
    title := "大阪出張" // string | 申請タイトル (optional)
    formId := int32(56) // int32 | 申請フォームID (optional)
    startApplicationDate := "startApplicationDate_example" // string | 申請日で絞込：開始日(yyyy-mm-dd) (optional)
    endApplicationDate := "endApplicationDate_example" // string | 申請日で絞込：終了日(yyyy-mm-dd) (optional)
    applicantId := int32(56) // int32 | 申請者のユーザーID (optional)
    approverId := int32(1) // int32 | 承認者のユーザーID (optional)
    offset := int32(56) // int32 | 取得レコードのオフセット (デフォルト: 0) (optional)
    limit := int32(56) // int32 | 取得レコードの件数 (デフォルト: 50, 最小: 1, 最大: 500) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApprovalRequestsApi.GetApprovalRequests(context.Background()).CompanyId(companyId).Status(status).ApplicationNumber(applicationNumber).Title(title).FormId(formId).StartApplicationDate(startApplicationDate).EndApplicationDate(endApplicationDate).ApplicantId(applicantId).ApproverId(approverId).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApprovalRequestsApi.GetApprovalRequests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApprovalRequests`: ApprovalRequestsIndexResponse
    fmt.Fprintf(os.Stdout, "Response from `ApprovalRequestsApi.GetApprovalRequests`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApprovalRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **int32** | 事業所ID | 
 **status** | **string** | 申請ステータス(draft:下書き, in_progress:申請中, approved:承認済, rejected:却下, feedback:差戻し) | 
 **applicationNumber** | **int32** | 申請No. | 
 **title** | **string** | 申請タイトル | 
 **formId** | **int32** | 申請フォームID | 
 **startApplicationDate** | **string** | 申請日で絞込：開始日(yyyy-mm-dd) | 
 **endApplicationDate** | **string** | 申請日で絞込：終了日(yyyy-mm-dd) | 
 **applicantId** | **int32** | 申請者のユーザーID | 
 **approverId** | **int32** | 承認者のユーザーID | 
 **offset** | **int32** | 取得レコードのオフセット (デフォルト: 0) | 
 **limit** | **int32** | 取得レコードの件数 (デフォルト: 50, 最小: 1, 最大: 500) | 

### Return type

[**ApprovalRequestsIndexResponse**](approvalRequestsIndexResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApprovalRequest

> ApprovalRequestResponse UpdateApprovalRequest(ctx, id).ApprovalRequestUpdateParams(approvalRequestUpdateParams).Execute()

各種申請の更新



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
    id := int32(56) // int32 | 各種申請ID
    approvalRequestUpdateParams := *openapiclient.NewApprovalRequestUpdateParams(int32(1), "2020-01-01", int32(1), true, []openapiclient.ApprovalRequestCreateParamsRequestItems{*openapiclient.NewApprovalRequestCreateParamsRequestItems()}) // ApprovalRequestUpdateParams | 各種申請の更新

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApprovalRequestsApi.UpdateApprovalRequest(context.Background(), id).ApprovalRequestUpdateParams(approvalRequestUpdateParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApprovalRequestsApi.UpdateApprovalRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApprovalRequest`: ApprovalRequestResponse
    fmt.Fprintf(os.Stdout, "Response from `ApprovalRequestsApi.UpdateApprovalRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | 各種申請ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApprovalRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **approvalRequestUpdateParams** | [**ApprovalRequestUpdateParams**](ApprovalRequestUpdateParams.md) | 各種申請の更新 | 

### Return type

[**ApprovalRequestResponse**](approvalRequestResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApprovalRequestAction

> ApprovalRequestResponse UpdateApprovalRequestAction(ctx, id).ApprovalRequestActionCreateParams(approvalRequestActionCreateParams).Execute()

各種申請の承認操作



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
    id := int32(56) // int32 | 各種申請ID
    approvalRequestActionCreateParams := *openapiclient.NewApprovalRequestActionCreateParams(int32(1), "approve", int32(1), int32(1)) // ApprovalRequestActionCreateParams | 各種申請の承認操作

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApprovalRequestsApi.UpdateApprovalRequestAction(context.Background(), id).ApprovalRequestActionCreateParams(approvalRequestActionCreateParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApprovalRequestsApi.UpdateApprovalRequestAction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApprovalRequestAction`: ApprovalRequestResponse
    fmt.Fprintf(os.Stdout, "Response from `ApprovalRequestsApi.UpdateApprovalRequestAction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | 各種申請ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApprovalRequestActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **approvalRequestActionCreateParams** | [**ApprovalRequestActionCreateParams**](ApprovalRequestActionCreateParams.md) | 各種申請の承認操作 | 

### Return type

[**ApprovalRequestResponse**](approvalRequestResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

